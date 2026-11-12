package views

import (
	"mime"
	"path"
	"strings"

	"golang_app/golangApp/static"

	"github.com/gofiber/fiber/v3"
)

// AssetHandler serves the embedded bundle at `/public/*` in production.
//
// Development keeps using `app.Static("/public", "./static")`, because the Vite
// watcher rewrites those files on every save and they are not part of the binary.
func AssetHandler() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		name := strings.TrimPrefix(path.Clean(ctx.Path()), "/public/")
		if name == "." || strings.HasPrefix(name, "..") {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		// Only the compiled bundle and the shared styles are embedded.
		if !strings.HasPrefix(name, "dist/") && !strings.HasPrefix(name, "styles/") {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		data, err := static.Assets.ReadFile(name)
		if err != nil {
			return ctx.SendStatus(fiber.StatusNotFound)
		}

		if contentType := mime.TypeByExtension(path.Ext(name)); contentType != "" {
			ctx.Set(fiber.HeaderContentType, contentType)
		}

		if strings.HasPrefix(name, "dist/") {
			// Hashed filenames: safe to cache forever (Rails: config.assets.digest).
			ctx.Set(fiber.HeaderCacheControl, "public, max-age=31536000, immutable")
		} else {
			ctx.Set(fiber.HeaderCacheControl, "public, max-age=86400")
		}

		return ctx.Send(data)
	}
}
