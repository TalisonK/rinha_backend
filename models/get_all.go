package models

import (
	"github.com/talisonk/rinha/database"
)

func GetAll() (pessoas []Pessoa, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM pessoas`)

	if err != nil {
		return
	}

	for rows.Next() {
		var pessoa Pessoa

		err = rows.Scan(&pessoa.Id, &pessoa.Apelido, &pessoa.Nascimento, &pessoa.Nome, &pessoa.Stack)

		if err != nil {
			continue
		}

		pessoas = append(pessoas, pessoa)
	}

	return

}
