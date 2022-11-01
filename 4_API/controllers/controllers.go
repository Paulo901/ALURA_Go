package controllers

import (
	"4_API/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "HOME")
}
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	/*
		com Docker
		var tp []models.Personalidade
		database.DB.Find(&tp)
	*/

	json.NewEncoder(w).Encode(models.Personalidades)

}
func OnePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("Nome")
		historia := r.FormValue("Historia")

		Id := 1
		Id++
		models.Adicionando(Id, nome, historia)

	}
	http.Redirect(w, r, "/", 301)
}
