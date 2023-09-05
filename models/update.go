package models

import (
	"github.com/google/uuid"
	"github.com/talisonk/rinha/database"
)

func Upgrade(id uuid.UUID, pessoa Pessoa) (int64, error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`UPDATE pessoas SET apelido=$1 nome=$2 nascimento=$3 stack=$4 WHERE id=$5`, pessoa.apelido, pessoa.nome, pessoa.nascimento, pessoa.stack, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
