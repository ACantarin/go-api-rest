package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/personas", controllers.AllPersonas)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
