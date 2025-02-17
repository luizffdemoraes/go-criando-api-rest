package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"pizzaria/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// https://pkg.go.dev/net/http
func GetPizzas(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": data.Pizzas,
	})
}

func PostPizzas(c *gin.Context) {
	// Cria uma nova instância de Pizza
	newPizza := models.Pizza{}
	// Faz o binding do JSON recebido para a struct Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&newPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// Salva as pizzas no arquivo pizzas.json
	data.SavePizzas()
	// Define um ID único para a nova pizza baseado no tamanho da lista
	newPizza.ID = len(data.Pizzas) + 1
	// Adiciona a nova pizza à lista de pizzas
	data.Pizzas = append(data.Pizzas, newPizza)
	c.JSON(http.StatusCreated, gin.H{
		"message": newPizza,
	})
}

func GetPizzasByID(c *gin.Context) {
	// Obtém o ID da pizza a partir do parâmetro da URL
	idParam := c.Param("id")
	// Converte o ID de string para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	// Percorre a lista de pizzas para encontrar a que possui o ID correspondente
	for _, pizza := range data.Pizzas {
		// Verifica se o ID da pizza atual corresponde ao ID buscado retorna a pizza encontrada
		if pizza.ID == id {
			c.JSON(http.StatusOK, pizza)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Pizza not found",
	})
}

func DeletePizzasByID(c *gin.Context) {
	// Obtém o ID da pizza a partir do parâmetro da URL
	idParam := c.Param("id")
	// Converte o ID de string para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	// Percorre a lista de pizzas para encontrar a que possui o ID correspondente
	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			// Remove a pizza da lista, mantendo os elementos antes e depois dela
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizzas()
			c.JSON(http.StatusOK, gin.H{"message": "Pizza deleted", "id": id})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}

func UpdatePizzasByID(c *gin.Context) {
	// Obtém o parâmetro "id" da URL
	idParam := c.Param("id")
	// Converte o ID de string para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	// Cria uma variável para armazenar os dados da pizza atualizada
	updatedPizza := models.Pizza{}
	// Faz o binding do JSON recebido para a estrutura updatedPizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error()})
		return
	}

	if err := service.ValidatePizzaPrice(&updatedPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Percorre a lista de pizzas para encontrar a que possui o ID correspondente
	for i, pizza := range data.Pizzas {
		if pizza.ID == id {
			// Atualiza os dados da pizza encontrada
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id // Mantém o ID original
			data.SavePizzas()
			c.JSON(http.StatusOK, data.Pizzas[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
}
