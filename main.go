package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

// https://gin-gonic.com/docs/quickstart/
func main() {
	loadPizzas()
	router := gin.Default()

	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:id", getPizzasByID)
	router.DELETE("/pizzas/:id", deletePizzasByID)
	router.PUT("/pizzas/:id", updatePizzasByID)
	router.Run() // listen and serve on 0.0.0.0:8080
}

func getPizzas(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	// Cria uma nova instância de Pizza
	newPizza := models.Pizza{}
	// Faz o binding do JSON recebido para a struct Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error()})
		return
	}
	// Salva as pizzas no arquivo pizzas.json
	savePizzas()
	// Define um ID único para a nova pizza baseado no tamanho da lista
	newPizza.ID = len(pizzas) + 1
	// Adiciona a nova pizza à lista de pizzas
	pizzas = append(pizzas, newPizza)
	c.JSON(200, gin.H{
		"message": newPizza,
	})
}

func getPizzasByID(c *gin.Context) {
	// Obtém o ID da pizza a partir do parâmetro da URL
	idParam := c.Param("id")
	// Converte o ID de string para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	// Percorre a lista de pizzas para encontrar a que possui o ID correspondente
	for _, pizza := range pizzas {
		// Verifica se o ID da pizza atual corresponde ao ID buscado retorna a pizza encontrada
		if pizza.ID == id {
			c.JSON(200, pizza)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza not found",
	})
}

func deletePizzasByID(c *gin.Context) {
	// Obtém o ID da pizza a partir do parâmetro da URL
	idParam := c.Param("id")
	// Converte o ID de string para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	// Percorre a lista de pizzas para encontrar a que possui o ID correspondente
	for i, pizza := range pizzas {
		if pizza.ID == id {
			// Remove a pizza da lista, mantendo os elementos antes e depois dela
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			savePizzas()
			c.JSON(200, gin.H{"message": "Pizza deleted", "id": id})
			return
		}
	}

	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func updatePizzasByID(c *gin.Context) {
	// Obtém o parâmetro "id" da URL
	idParam := c.Param("id")
	// Converte o ID de string para inteiro
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	// Cria uma variável para armazenar os dados da pizza atualizada
	updatedPizza := models.Pizza{}
	// Faz o binding do JSON recebido para a estrutura updatedPizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error()})
		return
	}

	// Percorre a lista de pizzas para encontrar a que possui o ID correspondente
	for i, pizza := range pizzas {
		if pizza.ID == id {
			// Atualiza os dados da pizza encontrada
			pizzas[i] = updatedPizza
			pizzas[i].ID = id // Mantém o ID original
			savePizzas()
			c.JSON(200, pizzas[i])
			return
		}
	}

	c.JSON(404, gin.H{"message": "Pizza not found"})
}

func loadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decode:", err)
		return
	}
}

func savePizzas() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encode:", err)
		return
	}
}
