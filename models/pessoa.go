package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "gorm.io/gorm"
)

type Pessoa struct {
	Id         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Apelido    string         `json:"apelido" gorm:"type:text;not null;unique;idx_pessoa"`
	Nome       string         `json:"nome" gorm:"type:text;not null;unique;idx_pessoa"`
	Nascimento string         `json:"nascimento" gorm:"type:text;not null"`
	Stack      pq.StringArray `json:"stack" gorm:"type:text[];idx_pessoa"`
}
