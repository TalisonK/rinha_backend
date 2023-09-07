package models

import (
	"log"

	"github.com/google/uuid"
	"github.com/talisonk/rinha/database"
)

func Insert(pessoa Pessoa) (id uuid.UUID, err error) {

	conn, err := database.OpenConnection()

	log.Println("Inserting", pessoa)

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO pessoas (apelido, nome, nascimento, stack) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pessoa.Stack).Scan(&id)

	return
}
