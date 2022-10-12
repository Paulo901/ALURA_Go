package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":5040", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{{"Sapato", 200, 20}, {"camisa", 50, 30}}
	temp.ExecuteTemplate(w, "Index", produtos)
}
