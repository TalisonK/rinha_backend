package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lib/pq"
	"github.com/talisonk/rinha/database"
	"github.com/talisonk/rinha/models"
)

func GetAll(res http.ResponseWriter, req *http.Request) {

	conn, err := database.OpenConnection()

	if err != nil {
		log.Printf("Erro ao fazer a inserção: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM pessoas`)

	if err != nil {
		return
	}

	var pessoas []models.Pessoa

	for rows.Next() {
		var pessoa models.Pessoa

		err = rows.Scan(&pessoa.Id, &pessoa.Apelido, &pessoa.Nascimento, &pessoa.Nome, pq.Array(&pessoa.Stack))

		if err != nil {
			continue
		}

		pessoas = append(pessoas, pessoa)
	}

	res.WriteHeader(http.StatusCreated)
	res.Header().Add("Location", "/pessoas/123")
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(pessoas)
}
