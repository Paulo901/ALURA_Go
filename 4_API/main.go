package main

import (
	"4_API/models"
	"4_API/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "n1", Historia: "h1"},
		{Id: 2, Nome: "n2", Historia: "h2"},
	}
	routes.HandleRequest()
}
