package main

import (
	"os"

	"github.com/marcelocosta/kombibeer/database"
	"github.com/marcelocosta/kombibeer/routes"
)

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"

	}

	database.ConectaComBancoDeDados(host)

	routes.HandleRequests()

}

// Sugestões: README, aplicação em conteiner, testes unitários
