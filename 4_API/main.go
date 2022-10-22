package main

import (
	"4_API/models"
	"4_API/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "n1", Historia: "h1"},
		{Nome: "n2", Historia: "h2"},
	}
	routes.HandleRequest()
}
