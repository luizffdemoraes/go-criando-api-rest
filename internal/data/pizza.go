package data

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/internal/models"
)

var Pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Println("Error decode:", err)
		return
	}
}

func SavePizzas() {
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Pizzas); err != nil {
		fmt.Println("Error encode:", err)
		return
	}
}
