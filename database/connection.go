package database

import (
	"fmt"
	"log"

	"github.com/talisonk/rinha/configs"
	"github.com/talisonk/rinha/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func OpenConnection() error {

	conf := configs.GetDB()

	//Configuração para conectar no banco de dados local
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", conf.Host, conf.User, conf.Pass, conf.Database, conf.Port)

	//Configuração para conectar no banco de dados via docker-compose
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=%s sslmode=disable", conf.User, conf.Pass, conf.Database, conf.Port)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("Failed to connect to database.", err)
		return err
	}

	log.Println("Connected to database!")
	conn.Logger = logger.Default.LogMode(logger.Info)

	conn.AutoMigrate(&models.Pessoa{})

	DB = Dbinstance{
		Db: conn,
	}

	return nil

}

func CloseConnection() {
	db, err := DB.Db.DB()
	if err != nil {
		log.Fatalln("Failed to close database.", err)
	}
	db.Close()
}
