// Package views renders the Single Page Application shell.
//
// Rails equivalent: `app/views/layouts/application.html.erb`. There are no
// per-controller templates anymore - every HTML route answers with the same
// shell and Vue owns the markup (see app/javascript). The only dynamic parts are
// the <head> metadata and the JSON props handed to the client.
//
// Templates come from the binary (`//go:embed layouts`) in production and from
// disk in development, so editing the shell while `go run .` is running does not
// require a rebuild.
package views

import (
	"embed"
)

//go:embed layouts
var layoutFS embed.FS

const (
	// ViewsRoot is where the shell lives on disk (development hot reload).
	ViewsRoot = "./app/views"

	// LayoutName is the template executed for every HTML response.
	LayoutName = "application.html"
)
