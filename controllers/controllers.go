package controllers

import (
	"encoding/json"
	"fmt"
	"go-api-rest/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personas)
}
