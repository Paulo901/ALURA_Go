package models

import "web/db"

type Produto struct {
	Id, Quantidade int
	Nome           string
	Preco          float64
}

func BuscaProdutos() []Produto {

	db := db.ConectorBD()
	pegaProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	prod := Produto{}
	produtos := []Produto{}

	for pegaProdutos.Next() {
		var Id, Quantidade int
		var Nome string
		var Preco float64

		err = pegaProdutos.Scan(&Id, &Quantidade, &Nome, &Preco)

		if err != nil {
			panic(err.Error())
		}
		prod.Id = Id
		prod.Quantidade = Quantidade
		prod.Nome = Nome
		prod.Preco = Preco

		produtos = append(produtos, prod)

	}
	defer db.Close()
	return produtos
}
