package models

import (
	"github.com/google/uuid"
)

type Pessoa struct {
	id         uuid.UUID  `json: "id"; default:uuid_generate_v4()`
	apelido    string     `json: "apelido"`
	nome       string     `json: "nome"`
	nascimento string     `json: "nascimento"`
	stack      ([]string) `json: "stack"`
}
