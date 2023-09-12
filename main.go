package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/talisonk/rinha/configs"
	"github.com/talisonk/rinha/database"
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

	Router(app)

	app.Listen(":3033")

}
