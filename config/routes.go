package config

import (
	"fmt"
	ctrl "golang_app/golangApp/app/controllers"
	api "golang_app/golangApp/app/controllers/api"
	mid "golang_app/golangApp/app/middlewares"
	env "golang_app/golangApp/config/environments"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	"golang_app/golangApp/config/session"

	c "golang_app/golangApp/constants"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	App         *fiber.App
	Database    *mongo.Database
	Translation *localize.Localization
	Redis       *redis.RedisClient
	Config      *env.Config
	EnvCnfg     env.EnvironmentConfiguration
}

func RoutesNew(appCnf *Application, mongo *mongo.Database, translation *localize.Localization, redis *redis.RedisClient) *Routes {
	// get our configuration and setup fibers engine so our application can use fibers template
	config := appCnf.EnvConfig.GetConfiguration()
	engine := html.New(config.EngineHtmlPath, config.EnginePageType)

	// initiate fibers with option enable fibers template
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: config.EngineViewsLayout,
	})

	// Middleware to manage sessions
	sess := session.SessionStoreNew()
	app.Use(func(c *fiber.Ctx) error {
		// Save session for future middleware in local context
		c.Locals("session", sess)
		return c.Next()
	})

	return &Routes{
		App:         app,
		Database:    mongo,
		Translation: translation,
		Redis:       redis,
		Config:      config,
		EnvCnfg:     appCnf.EnvConfig,
	}
}

func (r *Routes) AddViewRoutes() {
	r.App.Static(r.Config.StaticAssetPublicPath, r.Config.StaticAssetPath)

	homeController := ctrl.NewHomeController(r.EnvCnfg, r.Translation, r.Redis)
	r.App.Get("/", homeController.IndexPage())
	r.App.Get("/home", homeController.HomePage())
}

func (r *Routes) AddAPIRoutes() {
	userController := api.NewUserController(r.Database, r.Translation, r.Redis)
	imageController := api.NewCloudinaryController(r.Translation, r.Redis)
	weddingController := api.NewWeddingController(r.Database, r.Translation, r.Redis)
	api := r.App.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/users/sign-up", userController.Signup())
	v1.Post("/users/login", userController.Login())
	v1.Post("/users/update-user", mid.JWTMiddleware(), userController.UpdateUser())
	v1.Post("/upload/image", mid.JWTMiddleware(), imageController.UploadFile())
	v1.Post("/wedding/create", mid.JWTMiddleware(), weddingController.CreateWeddingData())
	v1.Get("/wedding/:id", mid.JWTMiddleware(), weddingController.GetWeddingData())
}

func (r *Routes) StartServer() {
	r.AddViewRoutes()
	r.AddAPIRoutes()
	r.App.Listen(fmt.Sprintf(":%s", os.Getenv(c.PORT)))
}
