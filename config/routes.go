package config

import (
	"fmt"
	cntrl "golang_app/golangApp/app/controllers"
	"golang_app/golangApp/app/middlewares"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"
	"golang_app/golangApp/config/session"
	c "golang_app/golangApp/constants"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routes struct {
	App         *fiber.App
	Database    *mongo.Database
	Translation *localize.Localization
	Redis       *redis.RedisClient
}

func RoutesNew(mongodb *mongo.Database, translation *localize.Localization, redis *redis.RedisClient) *Routes {
	app := fiber.New()
	ssn := session.SessionStoreNew()
	app.Use(func(ctx *fiber.Ctx) error {
		ses, err := ssn.Session.Get(ctx)
		if err != nil {
			return err
		}
		defer ses.Save()
		return ctx.Next()
	})
	return &Routes{
		App:         app,
		Database:    mongodb,
		Translation: translation,
		Redis:       redis,
	}
}

func (r *Routes) SetUpRoutes() {
	userController := cntrl.NewUserController(r.Database, r.Translation, r.Redis)
	imageController := cntrl.NewCloudinaryController(r.Translation, r.Redis)
	weddingController := cntrl.NewWeddingController(r.Database, r.Translation, r.Redis)

	api := r.App.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/users/sign-up", userController.Signup())
	v1.Post("/users/login", userController.Login())
	v1.Post("/users/update-user", middlewares.JWTMiddleware(), userController.UpdateUser())
	v1.Post("/upload/image", middlewares.JWTMiddleware(), imageController.UploadFile())
	v1.Post("/wedding/create", middlewares.JWTMiddleware(), weddingController.CreateWeddingData())
	v1.Get("/wedding/:id", middlewares.JWTMiddleware(), weddingController.GetWeddingData())
}

func (r *Routes) StartServer() {
	r.App.Listen(fmt.Sprintf(":%s", os.Getenv(c.PORT)))
}
