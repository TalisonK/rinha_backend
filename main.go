package main

import (
	"fmt"
	"log"
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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		log.Println("Get received")
	})

	log.Println("Server running on port", configs.GetServerPort())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r))

}
