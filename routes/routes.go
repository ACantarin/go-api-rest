package routes

import (
	"go-api-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersona).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
