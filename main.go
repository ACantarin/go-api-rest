package main

import (
	"fmt"
	"go-api-rest/models"
	"go-api-rest/routes"
)

func main() {
	models.Personas = []models.Persona{
		{
			Id:    1,
			Name:  "André",
			Story: "André é um programador",
		},
		{
			Id:    2,
			Name:  "João",
			Story: "João é um designer",
		},
	}
	fmt.Println("Iniciando o server Rest com Go...")
	routes.HandleRequests()
}
