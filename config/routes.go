package config

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	ctrl "golang_app/golangApp/app/controllers"
	api "golang_app/golangApp/app/controllers/api"
	mid "golang_app/golangApp/app/middlewares"
	"golang_app/golangApp/app/views"
	env "golang_app/golangApp/config/environments"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/logger"
	"golang_app/golangApp/config/redis"
	"golang_app/golangApp/config/session"

	c "golang_app/golangApp/constants"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Routes struct {
	App         *fiber.App
	Database    *gorm.DB
	Translation *localize.Localization
	Redis       *redis.RedisClient
	Config      *env.Config
	EnvCnfg     env.EnvironmentConfiguration
	Spa         *ctrl.SpaController
}

func RoutesNew(appCnf *Application, database *gorm.DB, translation *localize.Localization, redisClient *redis.RedisClient) *Routes {
	config := appCnf.EnvConfig.GetConfiguration()

	// The shell renderer resolves the compiled bundle once, at boot: embedded
	// `static/dist` in production, `static/dev` on disk everywhere else
	// (Rails: config/initializers/assets.rb).
	views.Configure(config.Env, config.StaticAssetPath, config.StaticAssetPublicPath)
	views.LogShellResolution()

	app := fiber.New(fiber.Config{
		AppName: "Golang App",
		// Handlers may return an error instead of writing a response; this decides
		// between the HTML shell and a JSON envelope.
		ErrorHandler: errorHandler,
		// Uploads are proxied to Cloudinary, so the 4MB Fiber default is tight.
		BodyLimit: 10 * 1024 * 1024,
	})

	// request logging (Rails: log/<environment>.log)
	logger.LogFormatterNew("", logType(), "", "", logFileName()).SetLoggerFormat(app)

	// Middleware to manage sessions. Fiber v3 passes fiber.Ctx (an interface),
	// not a *fiber.Ctx pointer.
	sess := session.SessionStoreNew(redisClient.Client())
	app.Use(func(ctx fiber.Ctx) error {
		// Save session for future middleware in local context
		ctx.Locals("session", sess)
		return ctx.Next()
	})

	return &Routes{
		App:         app,
		Database:    database,
		Translation: translation,
		Redis:       redisClient,
		Config:      config,
		EnvCnfg:     appCnf.EnvConfig,
		Spa:         ctrl.NewSpaController(appCnf.EnvConfig, translation),
	}
}

// AddViewRoutes is the "HTML part" of config/routes.rb. Every route answers with
// the SPA shell; the Vue router (app/javascript/router.js) mirrors this table.
func (r *Routes) AddViewRoutes() {
	if r.Config.Env == "production" {
		// The bundle ships inside the binary, so it is served from the embed FS.
		r.App.Get("/public/*", views.AssetHandler())
	} else {
		// Development: Vite rewrites ./static on every save.
		r.App.Static(r.Config.StaticAssetPublicPath, r.Config.StaticAssetPath)
	}

	// root "home#index"
	r.App.Get("/", r.Spa.Index())

	// resources :weddings, only: [:show]
	r.App.Get("/wedding/:id", r.Spa.WeddingShow())
}

// AddAPIRoutes is the `namespace :api` block: JSON only, no shell rendering.
func (r *Routes) AddAPIRoutes() {
	userController := api.NewUserController(r.Database, r.Translation, r.Redis)
	imageController := api.NewCloudinaryController(r.Translation, r.Redis)
	weddingController := api.NewWeddingController(r.Database, r.Translation, r.Redis)
	apiRouter := r.App.Group("/api")

	// CSRF protects the cookie-authenticated SPA calls; bearer-token clients and
	// the pre-auth endpoints (sign-up/login) are unaffected.
	v1 := apiRouter.Group("/v1", mid.CSRF())
	v1.Post("/users/sign-up", userController.Signup())
	v1.Post("/users/login", userController.Login())
	v1.Post("/users/logout", userController.Logout())
	v1.Post("/users/update-user", mid.CheckJWTandSessionMiddleware(), userController.UpdateUser())
	v1.Post("/upload/image", mid.CheckJWTandSessionMiddleware(), imageController.UploadFile())
	v1.Post("/wedding/create", mid.CheckJWTandSessionMiddleware(), weddingController.CreateWeddingData())
	v1.Get("/wedding/:id", mid.CheckJWTandSessionMiddleware(), weddingController.GetWeddingData())
}

// StartServer listens until SIGINT/SIGTERM (Fiber v3 shuts down gracefully
// instead of killing the process) and returns the reason it stopped.
func (r *Routes) StartServer() error {
	r.AddViewRoutes()
	r.AddAPIRoutes()

	// Catch-all, registered last (Rails: `get "*path", to: "errors#not_found"`).
	r.App.Use(r.Spa.NotFound())

	port := os.Getenv(c.PORT)
	if port == "" {
		port = "3000"
	}

	shutdownCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Fiber v3's Listen returns an error only when the server fails to start or
	// stops unexpectedly, so it is reported back to main instead of being lost.
	errCh := make(chan error, 1)
	go func() {
		if err := r.App.Listen(":" + port); err != nil {
			errCh <- err
		}
	}()

	log.Printf("Server listening on port %s", port)

	select {
	case err := <-errCh:
		return err
	case <-shutdownCtx.Done():
		log.Println("Shutting down gracefully, press Ctrl+C again to force")
		return r.App.ShutdownWithTimeout(10 * time.Second)
	}
}

// errorHandler renders unhandled errors: the SPA shell for browser requests,
// JSON for API clients (the way a Rails app splits HTML and JSON error pages).
func errorHandler(ctx fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
	}

	if views.IsHTMLRequest(ctx) {
		props := views.Props{
			Title:       "Something went wrong",
			Description: "An unexpected error occurred.",
			Route:       "error",
			Status:      code,
		}
		if code == fiber.StatusNotFound {
			props.Title = "Page not found"
		}
		if renderErr := views.RenderShell(ctx, props); renderErr == nil {
			return nil
		}
	}

	return ctx.Status(code).JSON(fiber.Map{
		c.STATUS:        c.FAILED,
		c.ERROR_MESSAGE: err.Error(),
	})
}

// logType mirrors Rails' per-environment logger: stdout in development,
// log/production.log in production.
func logType() string {
	if os.Getenv("ENV") == "production" {
		return "file"
	}
	return "out"
}

func logFileName() string {
	if os.Getenv("ENV") == "production" {
		return "production"
	}
	return ""
}
