package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api-rest/database"
	"go-api-rest/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade

	database.DB.Find(&personalidades)

	json.NewEncoder(w).Encode(personalidades)
}

func GetPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var personalidade models.Personalidade
	result := database.DB.First(&personalidade, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Registro não encontrado",
		})
		return
	}

	json.NewEncoder(w).Encode(personalidade)
}

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}
