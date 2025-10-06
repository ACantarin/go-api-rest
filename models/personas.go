package models

type Persona struct {
	Name  string `json:"name"`
	Story string `json:"story"`
}

var Personas []Persona
