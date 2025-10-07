package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personas)
}

func GetPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, persona := range models.Personas {
		if persona.Id == id {
			json.NewEncoder(w).Encode(persona)
		}
	}
}
