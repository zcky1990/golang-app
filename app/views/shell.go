package views

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
)

// Props is everything the shell is allowed to know about a request.
//
// Rails equivalent: the locals passed to `render layout: true`. Secrets and
// tokens never reach the HTML - authentication lives in an httpOnly cookie.
type Props struct {
	Title       string
	Description string
	Locale      string
	Route       string
	Params      map[string]string
	Status      int
}

// pageProps is serialised into the page and read back by app/javascript/main.js.
type pageProps struct {
	Route  string            `json:"route"`
	Params map[string]string `json:"params,omitempty"`
	Locale string            `json:"locale"`
	Env    string            `json:"env"`
}

type shellData struct {
	Title        string
	Description  string
	Locale       string
	Route        string
	PropsJSON    template.JS
	CSS          []string
	JS           string
	PublicPrefix string
	Env          string
	Year         int
}

type shellTemplate struct {
	tmpl  *template.Template
	stamp string
}

var (
	shellMu    sync.Mutex
	shellCache *shellTemplate
)

// RenderShell writes the SPA shell with a 200 (or props.Status) response.
func RenderShell(ctx fiber.Ctx, props Props) error {
	shell, err := shell()
	if err != nil {
		return err
	}

	if props.Locale == "" {
		props.Locale = "en"
	}
	if props.Title == "" {
		props.Title = "Golang App"
	}

	bundle := BundleFor(EntryPoint)

	payload, err := json.Marshal(pageProps{
		Route:  props.Route,
		Params: props.Params,
		Locale: props.Locale,
		Env:    envName(),
	})
	if err != nil {
		return err
	}

	data := shellData{
		Title:        props.Title,
		Description:  props.Description,
		Locale:       props.Locale,
		Route:        props.Route,
		PropsJSON:    template.JS(payload),
		CSS:          bundle.CSS,
		JS:           bundle.JS,
		PublicPrefix: PublicURL(),
		Env:          envName(),
		Year:         time.Now().Year(),
	}

	var buf bytes.Buffer
	if err := shell.ExecuteTemplate(&buf, LayoutName, data); err != nil {
		return err
	}

	status := props.Status
	if status == 0 {
		status = fiber.StatusOK
	}

	ctx.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
	// The shell links hashed assets, so it must never be cached itself.
	ctx.Set(fiber.HeaderCacheControl, "no-store")

	return ctx.Status(status).Send(buf.Bytes())
}

// shell parses the layout once and re-parses it only when the file changed on
// disk (development). In production the embedded copy is parsed a single time.
func shell() (*template.Template, error) {
	cfgMu.RLock()
	c := cfg
	cfgMu.RUnlock()

	stamp := shellStamp(c)

	shellMu.Lock()
	defer shellMu.Unlock()

	if shellCache != nil && shellCache.stamp == stamp {
		return shellCache.tmpl, nil
	}

	tmpl, err := parseShell(c)
	if err != nil {
		return nil, err
	}

	shellCache = &shellTemplate{tmpl: tmpl, stamp: stamp}
	return tmpl, nil
}

func parseShell(c assetConfig) (*template.Template, error) {
	// Development: read the layout from disk so edits show up on the next request.
	if c.mode == ModeDisk {
		if info, err := os.Stat(filepath.Join(ViewsRoot, "layouts")); err == nil && info.IsDir() {
			return template.New(LayoutName).ParseFS(os.DirFS(ViewsRoot), path.Join("layouts", "*.html"))
		}
	}

	// Production (and fallback): the embedded copy shipped in the binary.
	return template.New(LayoutName).ParseFS(layoutFS, path.Join("layouts", "*.html"))
}

func shellStamp(c assetConfig) string {
	if c.mode == ModeEmbed {
		return "embedded"
	}

	info, err := os.Stat(filepath.Join(ViewsRoot, "layouts", LayoutName))
	if err != nil {
		return "embedded"
	}
	return info.ModTime().UTC().Format(time.RFC3339Nano)
}

// IsHTMLRequest reports whether a request expects HTML rather than JSON. It is
// how the error handler decides between the shell and a JSON envelope.
func IsHTMLRequest(ctx fiber.Ctx) bool {
	if strings.HasPrefix(ctx.Path(), "/api") {
		return false
	}

	accept := ctx.Get(fiber.HeaderAccept)
	if strings.Contains(accept, "application/json") && !strings.Contains(accept, "text/html") {
		return false
	}
	return true
}

func envName() string {
	if env := os.Getenv("ENV"); env != "" {
		return env
	}
	return "development"
}

// LogShellResolution prints where the shell and the bundle were resolved from.
// Called once at boot, so a misconfigured deployment fails loudly in the log
// instead of silently serving an unstyled page.
func LogShellResolution() {
	bundle := BundleFor(EntryPoint)
	log.Printf("views: mode=%s shell=%s js=%q css=%v", CurrentMode(), shellSource(), bundle.JS, bundle.CSS)
	if bundle.JS == "" {
		log.Printf("views: no JavaScript bundle found for %s; run `npm run build:dev` (development) or `npm run build:prod` (production)", EntryPoint)
	}
}

func shellSource() string {
	cfgMu.RLock()
	mode := cfg.mode
	cfgMu.RUnlock()

	if mode == ModeDisk {
		if info, err := os.Stat(filepath.Join(ViewsRoot, "layouts", LayoutName)); err == nil {
			return info.ModTime().UTC().Format(time.RFC3339)
		}
	}
	return "embedded"
}
