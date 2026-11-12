// Package static bundles the compiled front-end assets into the Go binary.
//
// Rails equivalent: the `public/` directory shipped with the app. Because the
// binary is the deployment unit here, everything the SPA shell links to must be
// available at compile time:
//
//	dist/   -> Vite production build (hashed JS/CSS, .vite/manifest.json)
//	styles/ -> fonts referenced by absolute URLs from the compiled stylesheet
//
// The source assets (static/css, static/javascript) stay on disk; only build
// output is embedded. `all:` is used so dotted files (like
// dist/.vite/manifest.json) are included - without it `go:embed` silently skips
// them and the shell would render without CSS/JS.
package static

import (
	"embed"
	"io/fs"
)

//go:embed all:dist all:styles
var embedded embed.FS

// Assets is the embedded asset tree (root contains "dist/" and "styles/").
var Assets embed.FS = embedded

// Dist returns the embedded production bundle (the `dist/` sub-tree).
func Dist() (fs.FS, error) {
	return fs.Sub(embedded, "dist")
}
