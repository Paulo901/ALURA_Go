package controller

import (
	"3_WEB/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))
var id = 2

//apenas porque não está incrementando

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
		if err != nil {
			log.Println("Erro de Formatação", err)
		}
		quantidadeStr, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro de Formatação", err)
		}
		id++
		models.CriandoProdutos(id, nome, precoStr, quantidadeStr)

	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	GetId := r.URL.Query().Get("id")
	models.DeletaProduto(GetId)
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na convesão da quantidade para int:", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	//idDoProduto := r.URL.Query().Get("Id")
	//produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", nil)
}
