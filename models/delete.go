package models

import (
	"github.com/google/uuid"
	"github.com/talisonk/rinha/database"
)

func Delete(id uuid.UUID) (int64, error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM pessoas WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}
