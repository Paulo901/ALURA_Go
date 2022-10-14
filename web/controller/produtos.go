package controller

import (
	"net/http"
	"text/template"
	"web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}
