package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Vitor", Historia: "Lutador"},
		{Id: 2, Nome: "Franklin", Historia: "Cantor"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor rest com Go")
	routes.HandleRequest()
}
