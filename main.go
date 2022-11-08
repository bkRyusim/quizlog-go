package main

import (
	"context"
	"github.com/bkRyusim/quizlog-go/config"
	"github.com/bkRyusim/quizlog-go/controller"
	"github.com/bkRyusim/quizlog-go/ent"
	"github.com/facebookgo/inject"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"log"
)

func main() {
	app := Setup()

	log.Fatal(app.Listen(":8080"))
}

func Setup() *fiber.App {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config.Config)
	if err != nil {
		panic(err)
	}

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err != nil {
		log.Fatal(err)
	}

	var g inject.Graph
	var userController controller.UserController
	var authController controller.AuthController
	var postController controller.PostController
	var quizController controller.QuizController

	err = g.Provide(
		&inject.Object{Value: &userController},
		&inject.Object{Value: &authController},
		&inject.Object{Value: &postController},
		&inject.Object{Value: &quizController},
		&inject.Object{Value: client.User, Name: "userClient"},
		&inject.Object{Value: client.Quiz, Name: "quizClient"})

	if err != nil {
		panic(err)
	}

	if err := g.Populate(); err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Quizlog API")
	})

	app.Post("/auth", authController.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Config.AuthConfig.SecretKey),
	}))

	app.Post("/join", userController.NewUser)
	app.Get("/posts", postController.GetAllPosts)
	app.Post("/quiz", quizController.NewQuiz)
	app.Get("/quiz", quizController.GetAllQuiz)

	return app
}
