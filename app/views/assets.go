package views

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang_app/golangApp/static"
)

// Mode decides where the compiled bundle is read from.
type Mode string

const (
	// ModeDisk reads `static/dev` from the filesystem (development and test).
	// `npm run dev` (vite build --watch) rewrites that folder, and the renderer
	// picks the new manifest up on the next request without a Go restart.
	ModeDisk Mode = "disk"

	// ModeEmbed reads the bundle compiled into the binary (production).
	ModeEmbed Mode = "embed"
)

const (
	// EntryPoint is the Vite entry, relative to the project root.
	EntryPoint = "app/javascript/main.js"

	manifestFile = ".vite/manifest.json"
)

// Bundle is the set of public URLs the shell has to link for one entry point.
type Bundle struct {
	JS  string
	CSS []string
}

type assetConfig struct {
	mode         Mode
	staticRoot   string
	publicPrefix string
	bundleDir    string
}

var (
	cfgMu sync.RWMutex
	cfg   = assetConfig{
		mode:         ModeDisk,
		staticRoot:   "./static",
		publicPrefix: "/public",
		bundleDir:    "dev",
	}
)

// Configure wires the asset resolver for an environment.
//
// Production reads the bundle embedded into the binary (`static/dist`), every
// other environment reads `static/dev` from disk. This is decided once, at boot,
// so a production build can never accidentally serve development assets that do
// not exist in the binary.
func Configure(env, staticRoot, publicPrefix string) {
	mode := ModeDisk
	if env == "production" {
		mode = ModeEmbed
	}
	if staticRoot == "" {
		staticRoot = "./static"
	}
	if publicPrefix == "" {
		publicPrefix = "/public"
	}
	ConfigureWith(mode, staticRoot, publicPrefix)
}

// ConfigureWith is the explicit form of Configure (used by tests).
func ConfigureWith(mode Mode, staticRoot, publicPrefix string) {
	bundleDir := "dev"
	if mode == ModeEmbed {
		bundleDir = "dist"
	}

	cfgMu.Lock()
	cfg = assetConfig{
		mode:         mode,
		staticRoot:   staticRoot,
		publicPrefix: publicPrefix,
		bundleDir:    bundleDir,
	}
	cfgMu.Unlock()

	// Invalidate the caches: they are keyed by config, but a mode switch must not
	// keep serving a bundle resolved for the previous one.
	bundleMu.Lock()
	bundleCache = map[string]cachedBundle{}
	bundleMu.Unlock()

	shellMu.Lock()
	shellCache = nil
	shellMu.Unlock()
}

// CurrentMode reports the resolved asset mode (used by logging and tests).
func CurrentMode() Mode {
	cfgMu.RLock()
	defer cfgMu.RUnlock()
	return cfg.mode
}

// PublicURL builds an absolute URL below the public prefix.
func PublicURL(elem ...string) string {
	cfgMu.RLock()
	prefix := strings.Trim(cfg.publicPrefix, "/")
	cfgMu.RUnlock()

	parts := append([]string{"/", prefix}, elem...)
	return path.Join(parts...)
}

type manifestEntry struct {
	File    string   `json:"file"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	CSS     []string `json:"css"`
}

type cachedBundle struct {
	stamp string
	value Bundle
}

var (
	bundleMu    sync.Mutex
	bundleCache = map[string]cachedBundle{}
)

// BundleFor resolves the hashed filenames produced by Vite for the given entry.
func BundleFor(entry string) Bundle {
	cfgMu.RLock()
	c := cfg
	cfgMu.RUnlock()

	stamp := bundleStamp(c)
	key := entry + "|" + string(c.mode) + "|" + c.staticRoot + "|" + c.bundleDir

	bundleMu.Lock()
	if cached, ok := bundleCache[key]; ok && cached.stamp == stamp {
		value := cached.value
		bundleMu.Unlock()
		return value
	}
	bundleMu.Unlock()

	value := resolveBundle(c, entry)

	bundleMu.Lock()
	bundleCache[key] = cachedBundle{stamp: stamp, value: value}
	bundleMu.Unlock()

	return value
}

// bundleStamp changes whenever Vite rewrote the manifest, which is how the dev
// watcher is picked up without restarting the server.
func bundleStamp(c assetConfig) string {
	if c.mode == ModeEmbed {
		return "embedded"
	}

	p := filepath.Join(c.staticRoot, c.bundleDir, filepath.FromSlash(manifestFile))
	info, err := os.Stat(p)
	if err != nil {
		return "missing"
	}
	return info.ModTime().UTC().Format(time.RFC3339Nano) + ":" + strconv.FormatInt(info.Size(), 10)
}

func bundleFS(c assetConfig) (fs.FS, error) {
	if c.mode == ModeEmbed {
		return static.Dist()
	}
	sub, err := fs.Sub(os.DirFS(c.staticRoot), c.bundleDir)
	if err != nil {
		return nil, err
	}
	return sub, nil
}

func resolveBundle(c assetConfig, entry string) Bundle {
	fsys, err := bundleFS(c)
	if err != nil {
		log.Printf("views: cannot open bundle directory %q: %v", c.bundleDir, err)
		return Bundle{}
	}

	data, err := fs.ReadFile(fsys, manifestFile)
	if err != nil {
		log.Printf("views: %s/%s not found (%v); run `npm run build:dev` or `npm run build:prod`", c.bundleDir, manifestFile, err)
		return scanBundle(fsys, c)
	}

	var manifest map[string]manifestEntry
	if err := json.Unmarshal(data, &manifest); err != nil {
		log.Printf("views: invalid %s/%s: %v", c.bundleDir, manifestFile, err)
		return scanBundle(fsys, c)
	}

	selected, ok := pickEntry(manifest, entry)
	if !ok {
		log.Printf("views: entry %q not present in %s/%s", entry, c.bundleDir, manifestFile)
		return scanBundle(fsys, c)
	}

	bundle := Bundle{JS: publicURL(c, selected.File)}
	for _, css := range selected.CSS {
		bundle.CSS = append(bundle.CSS, publicURL(c, css))
	}
	return bundle
}

// pickEntry tolerates the different manifest keys Vite generates depending on how
// the entry was declared (module id, "/" prefixed id, or a named chunk).
func pickEntry(manifest map[string]manifestEntry, entry string) (manifestEntry, bool) {
	if item, ok := manifest[entry]; ok {
		return item, true
	}

	keys := make([]string, 0, len(manifest))
	for key := range manifest {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		item := manifest[key]
		if !item.IsEntry {
			continue
		}
		trimmed := strings.TrimPrefix(key, "/")
		if trimmed == entry || strings.HasSuffix(trimmed, entry) || path.Base(trimmed) == path.Base(entry) {
			return item, true
		}
	}

	for _, key := range keys {
		if manifest[key].IsEntry {
			return manifest[key], true
		}
	}

	return manifestEntry{}, false
}

// scanBundle is the last resort when the manifest is missing or unreadable: it
// looks for the emitted files so a plain build without a manifest still renders.
func scanBundle(fsys fs.FS, c assetConfig) Bundle {
	var bundle Bundle

	_ = fs.WalkDir(fsys, "javascript", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		name := d.Name()
		if strings.HasPrefix(name, "main") && strings.HasSuffix(name, ".js") {
			bundle.JS = publicURL(c, p)
		}
		return nil
	})

	var css []string
	_ = fs.WalkDir(fsys, "assets", func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if strings.HasSuffix(d.Name(), ".css") {
			css = append(css, publicURL(c, p))
		}
		return nil
	})
	sort.Strings(css)
	bundle.CSS = css

	return bundle
}

func publicURL(c assetConfig, file string) string {
	prefix := strings.Trim(c.publicPrefix, "/")
	return path.Join("/", prefix, c.bundleDir, strings.TrimPrefix(file, "/"))
}
