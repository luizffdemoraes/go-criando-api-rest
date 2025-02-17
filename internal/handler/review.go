package handler

import (
	"net/http"
	"pizzaria/internal/data"
	"pizzaria/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostReviewsByIDPizza(c *gin.Context) {
	pizzaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var newReview models.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == pizzaID {
			p.Reviews = append(p.Reviews, newReview)
			data.Pizzas[i] = p
			data.SavePizzas()
			c.JSON(http.StatusCreated, gin.H{"message": data.Pizzas[i]})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})
	}

}
