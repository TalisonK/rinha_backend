package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/talisonk/rinha/database"
	"github.com/talisonk/rinha/models"
)

func Create(c *fiber.Ctx) error {

	pessoa := new(models.Pessoa)

	if err := c.BodyParser(pessoa); err != nil {
		log.Printf("Erro ao fazer a inserção: %v", err)
		return c.Status(http.StatusBadRequest).JSON(nil)
	}

	ret := database.DB.Db.Create(&pessoa).Select("id")

	if ret.Error != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(nil)
	}

	c.Set("Location", fmt.Sprintf("/pessoas/%v", pessoa.Id))

	return c.SendStatus(http.StatusCreated)

}

func GetPessoa(c *fiber.Ctx) error {

	uuid, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(nil)
	}

	pessoa := new(models.Pessoa)
	pessoa.Id = uuid

	ret := database.DB.Db.Select("id", "apelido", "nome", "nascimento", "stack").Find(&pessoa)

	if ret.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"Error":   true,
			"Message": fmt.Sprintf("Pessoa não encontrada %v", ret.Error),
		})
	}

	return c.Status(http.StatusOK).JSON(pessoa)

}

func FiltroPessoas(c *fiber.Ctx) error {

	query := c.Query("t")

	pessoas := new([]models.Pessoa)

	ret := database.DB.Db.Where("nome LIKE ?", "%"+query+"%").
		Or("apelido LIKE ?", "%"+query+"%").
		Or("stack @> ARRAY[?]", query).
		Limit(50).
		Find(&pessoas)

	if ret.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(nil)
	}

	return c.Status(http.StatusOK).JSON(pessoas)

}

func ContagemPessoas(c *fiber.Ctx) error {

	var count int64

	ret := database.DB.Db.Model(&models.Pessoa{}).Count(&count)

	if ret.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar contar as pessoas %v", ret.Error),
		})
	}

	return c.Status(http.StatusOK).JSON(count)
}
