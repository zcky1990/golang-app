package controllers

import (
	"strings"

	"golang_app/golangApp/app/services"
	"golang_app/golangApp/app/views"
	env "golang_app/golangApp/config/environments"
	"golang_app/golangApp/config/localize"
	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v3"
)

// SpaController answers every HTML route with the embedded SPA shell.
//
// Rails equivalent: `ApplicationController` + `render layout: true`. There is no
// per-action template anymore - the browser boots Vue and the client router takes
// over (app/javascript/router.js must mirror config/routes.go).
type SpaController struct {
	translation *localize.Localization
	envConf     env.EnvironmentConfiguration
}

func NewSpaController(envConf env.EnvironmentConfiguration, translation *localize.Localization) *SpaController {
	return &SpaController{envConf: envConf, translation: translation}
}

// Index is the root route (Rails: `root "home#index"`).
func (ctrl *SpaController) Index() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		locale := ctrl.locale(ctx)
		return views.RenderShell(ctx, views.Props{
			Title:       ctrl.message("PAGE_TITLE_HOME", locale),
			Description: ctrl.message("PAGE_DESCRIPTION_HOME", locale),
			Locale:      locale,
			Route:       "home",
		})
	}
}

// WeddingShow is the public invitation page (Rails: `resources :weddings, only: [:show]`).
func (ctrl *SpaController) WeddingShow() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		id := ctx.Params("id")
		locale := ctrl.locale(ctx)

		// The id ends up in window props, so it is validated before rendering;
		// an invalid id is a 404 rather than a warning inside the SPA.
		if !services.ValidUUID(id) {
			return ctrl.NotFound()(ctx)
		}

		return views.RenderShell(ctx, views.Props{
			Title:       ctrl.message("PAGE_TITLE_WEDDING", locale),
			Description: ctrl.message("PAGE_DESCRIPTION_WEDDING", locale),
			Locale:      locale,
			Route:       "wedding",
			Params:      map[string]string{"id": strings.ToLower(id)},
		})
	}
}

// NotFound is the catch-all (Rails: `get "*path", to: "errors#not_found"`).
// JSON is kept for API clients so a SPA fallback never answers an API call with
// HTML.
func (ctrl *SpaController) NotFound() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		locale := ctrl.locale(ctx)

		if !views.IsHTMLRequest(ctx) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				c.STATUS:        c.FAILED,
				c.ERROR_MESSAGE: ctrl.message("NOT_FOUND", locale),
			})
		}

		return views.RenderShell(ctx, views.Props{
			Title:       ctrl.message("PAGE_TITLE_NOT_FOUND", locale),
			Description: ctrl.message("PAGE_DESCRIPTION_NOT_FOUND", locale),
			Locale:      locale,
			Route:       "not_found",
			Status:      fiber.StatusNotFound,
		})
	}
}

// locale mirrors the Accept-Language handling used by the JSON controllers.
func (ctrl *SpaController) locale(ctx fiber.Ctx) string {
	accept := strings.ToLower(ctx.Get("Accept-Language"))
	switch {
	case strings.HasPrefix(accept, c.LOCALE_INDONESIA):
		return c.LOCALE_INDONESIA
	default:
		return c.LOCALE_ENGLISH
	}
}

func (ctrl *SpaController) message(key, locale string) string {
	if ctrl.translation == nil {
		return ""
	}
	return ctrl.translation.GetMessage(key, locale)
}
