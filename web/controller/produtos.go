package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoStr, err := strconv.ParseFloat(preco, 64)
		quantidadeStr, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro de Formatação", err)
		}
		models.CriandoProdutos(nome, precoStr, quantidadeStr)

	}
	http.Redirect(w, r, "/", 301)
}
