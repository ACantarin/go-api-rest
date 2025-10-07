package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {
	database.Connect()

	fmt.Println("Iniciando o server Rest com Go...")
	routes.HandleRequests()
}
