package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app := Setup()

	log.Fatal(app.Listen(":8080"))
}

func Setup() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Quizlog API")
	})
	return app
}
