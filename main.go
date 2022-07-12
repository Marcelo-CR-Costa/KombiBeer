package main

import (
	"github.com/marcelocosta/kombibeer/database"
	"github.com/marcelocosta/kombibeer/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()

}
