package controllers

import (
	"4_API/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "HOME")
}
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(models.Personalidades)

}
