package models

import (
	"github.com/google/uuid"
	"github.com/talisonk/rinha/database"
)

func Insert(pessoa Pessoa) (id uuid.UUID, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO pessoas (apelido, nome, nascimento, stack) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, pessoa.apelido, pessoa.nome, pessoa.nascimento, pessoa.stack).Scan(&id)

	return
}
