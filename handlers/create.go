package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/talisonk/rinha/database"
	"github.com/talisonk/rinha/models"
)

func Create(c *fiber.Ctx) error {

	pessoa := new(models.Pessoa)

	if err := c.BodyParser(pessoa); err != nil {
		log.Printf("Erro ao fazer a inserção: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir %v", err),
		})
	}

	ret := database.DB.Db.Create(&pessoa)

	log.Println(ret)

	return c.Status(http.StatusCreated).JSON(pessoa)

}

// func create(res http.ResponseWriter, req *http.Request) {

// 	var pessoa models.Pessoa

// 	err := json.NewDecoder(req.Body).Decode(&pessoa)
// 	if err != nil {
// 		log.Printf("Erro ao fazer a inserção: %v", err)
// 		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 		return
// 	}

// 	id, err := models.Insert(pessoa)

// 	var resp map[string]any

// 	if err != nil {
// 		resp = map[string]any{
// 			"Error":   true,
// 			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir %v", err),
// 		}
// 	} else {
// 		resp = map[string]any{
// 			"Error":   false,
// 			"Message": "Pessoa inserida com sucesso",
// 		}
// 	}

// 	res.WriteHeader(http.StatusCreated)
// 	res.Header().Add("Location", fmt.Sprintf("/pessoas/%v", id))
// 	res.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(res).Encode(resp)

// }
