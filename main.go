package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	pizzas := []Pizza{
		{ID: 1, nome: "Toscana", preco: 49.50},
		{ID: 2, nome: "Marguerita", preco: 79.50},
		{ID: 3, nome: "Atum com queijo", preco: 69.50},
	}
	fmt.Println(pizzas)
}
