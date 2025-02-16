package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/docs/quickstart/
func main() {
	router := gin.Default()

	router.GET("/pizzas", getPizzas)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func getPizzas(c *gin.Context) {

	pizzas := []models.Pizza{
		{ID: 1, Nome: "Calabresa", Preco: 30.00},
		{ID: 2, Nome: "Mussarela", Preco: 25.00},
		{ID: 3, Nome: "Portuguesa", Preco: 35.00},
	}

	c.JSON(200, gin.H{
		"message": pizzas,
	})
}
