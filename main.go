package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/talisonk/rinha/configs"
	"github.com/talisonk/rinha/handlers"
)

func main() {

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)

	http.ListenAndServe(configs.GetServerPort(), r)

}
