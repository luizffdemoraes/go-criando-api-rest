package main

import (
	"pizzaria/models"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Calabresa", Preco: 30.00},
	{ID: 2, Nome: "Mussarela", Preco: 25.00},
	{ID: 3, Nome: "Portuguesa", Preco: 35.00},
}

// https://gin-gonic.com/docs/quickstart/
func main() {
	router := gin.Default()

	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)

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
	pizzas = append(pizzas, newPizza)
	c.JSON(200, gin.H{
		"message": newPizza,
	})

}
