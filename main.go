package main

import (
	"entgo.io/ent/dialect"
	"github.com/bkRyusim/quizlog-go/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	app, client := Setup()
	defer client.Close()

	log.Fatal(app.Listen(":8080"))
}

func Setup() (*fiber.App, *ent.Client) {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Quizlog API")
	})
	return app, client
}
