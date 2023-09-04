package models

import (
	"github.com/google/uuid"
	"github.com/talisonk/rinha/database"
)

func Get(id uuid.UUID) (pessoa Pessoa, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM pessoas WHERE id=$1`, id)

	err = row.Scan(&pessoa.id, &pessoa.apelido, &pessoa.nascimento, &pessoa.nome, &pessoa.stack)

	return

}
