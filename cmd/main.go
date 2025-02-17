package main

import (
	"pizzaria/internal/data"
	"pizzaria/internal/handler"

	"github.com/gin-gonic/gin"
)

// https://gin-gonic.com/docs/quickstart/
func main() {
	data.LoadPizzas()
	router := gin.Default()

	router.GET("/pizzas", handler.GetPizzas)
	router.POST("/pizzas", handler.PostPizzas)
	router.GET("/pizzas/:id", handler.GetPizzasByID)
	router.DELETE("/pizzas/:id", handler.DeletePizzasByID)
	router.PUT("/pizzas/:id", handler.UpdatePizzasByID)
	router.POST("/pizzas/:id/reviews", handler.PostReviewsByIDPizza)
	router.Run() // listen and serve on 0.0.0.0:8080
}
