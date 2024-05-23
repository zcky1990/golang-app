package config

import (
	"fmt"
	cntrl "golang_app/golangApp/app/controllers"
	api "golang_app/golangApp/app/controllers/api"
	mid "golang_app/golangApp/app/middlewares"
	env "golang_app/golangApp/config/environments"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"

	db "golang_app/golangApp/config/mongo"
	"log"

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

func RoutesNew() *Routes {
	appCnf := GetApplicationInstance()

	mongoDB, err := db.NewMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect()

	translation := localize.NewLocalization()
	redisClient := redis.NewRedisClient()
	defer redisClient.Close()

	config := appCnf.EnvConfig.GetConfiguration()
	engine := html.New(config.EngineHtmlPath, config.EnginePageType)

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: config.EngineViewsLayout,
	})

	return &Routes{
		App:         app,
		Database:    mongoDB.Db,
		Translation: translation,
		Redis:       redisClient,
		Config:      config,
		EnvCnfg:     appCnf.EnvConfig,
	}
}

func (r *Routes) AddViewRoutes() {
	r.App.Static(r.Config.StaticAssetPublicPath, r.Config.StaticAssetPath)

	homeController := cntrl.NewHomeController(r.EnvCnfg, r.Translation, r.Redis)
	r.App.Get("/", homeController.IndexPage())
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
