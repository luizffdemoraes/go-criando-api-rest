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
	newPizza := models.Pizza{}
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error()})
		return
	}
	savePizzas()
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	c.JSON(200, gin.H{
		"message": newPizza,
	})
}

func getPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	for _, pizza := range pizzas {
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
	c.JSON(200, gin.H{"method": "DELETE"})
}

func updatePizzasByID(c *gin.Context) {
	c.JSON(200, gin.H{"method": "PUT"})
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
