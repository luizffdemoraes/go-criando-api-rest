package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {

	/***
	var Declaration
	var nomePizzaria string
	***/
	// := Short Assignment Statement
	nomePizzaria := "Pizzaria Go"
	instagram, telefone := "@pizzaria_go", 11951
	fmt.Printf("Nome da pizzaria: %s (instagram: %s) - Telefone: %d", nomePizzaria, instagram, telefone)
}
