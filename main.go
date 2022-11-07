package main

import (
	"context"
	"fmt"
	"github.com/bkRyusim/quizlog-go/config"
	"github.com/bkRyusim/quizlog-go/controller"
	"github.com/bkRyusim/quizlog-go/ent"
	"github.com/facebookgo/inject"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	fmt.Println(config.Config)

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

	err = g.Provide(
		&inject.Object{Value: &userController},
		&inject.Object{Value: &authController},
		&inject.Object{Value: client.User, Name: "userClient"})

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

	return app
}
