package config

import (
	"embed"
	"fmt"
	cntrl "golang_app/golangApp/app/controllers"
	api "golang_app/golangApp/app/controllers/api"
	mid "golang_app/golangApp/app/middlewares"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"

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
}

var viewsfs embed.FS

func RoutesNew(mongodb *mongo.Database, translation *localize.Localization, redis *redis.RedisClient) *Routes {
	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/application",
	})
	return &Routes{
		App:         app,
		Database:    mongodb,
		Translation: translation,
		Redis:       redis,
	}
}

func (r *Routes) AddViewRoutes() {
	// Serve static files from the "public" directory
	r.App.Static("/public", "./static")

	homeController := cntrl.NewHomeController(r.Translation, r.Redis)
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

func (r *Routes) SetUpRoutes() {
	r.AddViewRoutes()
	r.AddAPIRoutes()
}

func (r *Routes) StartServer() {
	r.App.Listen(fmt.Sprintf(":%s", os.Getenv(c.PORT)))
}
