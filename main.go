package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/talisonk/rinha/configs"
	"github.com/talisonk/rinha/database"
	"github.com/talisonk/rinha/handlers"
)

func main() {

	err := configs.Load()

	if err != nil {
		log.Fatal("Erro ao carregar as configurações")
		os.Exit(2)
	}

	err = database.OpenConnection()

	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
		os.Exit(2)
	}

	defer database.CloseConnection()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/", handlers.Create)

	app.Listen(":3033")

}
