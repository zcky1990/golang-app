package controllers

// The per-action HTML templates that used to live here (`IndexPage`, `HomePage`,
// the `metaData`/`ImportPath` plumbing and `ctx.Render(...)`) were removed when
// the view layer became a single embedded SPA shell.
//
// See:
//   - app/controllers/spa_controller.go -> routes that answer with the shell
//   - app/views/shell.go                -> the renderer (layout + hashed assets)
//   - app/javascript/                   -> the Vue application that owns the markup
