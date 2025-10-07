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
	r.HandleFunc("/personas", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/personas/{id}", controllers.GetPersona).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
