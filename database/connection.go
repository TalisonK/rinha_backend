package database

import (
	"database/sql"
	"fmt"

	"github.com/talisonk/rinha/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {

	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Host, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err

}
