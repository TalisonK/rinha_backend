package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/talisonk/rinha/handlers"
)

func Router(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/pessoas", handlers.Create)

	app.Get("/pessoas/:id", handlers.GetPessoa)

	app.Get("/pessoas?t", handlers.FiltroPessoas)

	app.Get("/contagem-pessoas", handlers.ContagemPessoas)

}
