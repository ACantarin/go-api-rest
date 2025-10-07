package routes

import (
	"go-api-rest/controllers"
	"go-api-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalidades", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/personalidades/{id}", controllers.GetPersona).Methods("Get")
	r.HandleFunc("/personalidades", controllers.CreatePersona).Methods("Post")
	r.HandleFunc("/personalidades/{id}", controllers.DeletePersona).Methods("Delete")
	r.HandleFunc("/personalidades/{id}", controllers.UpdatePersona).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
