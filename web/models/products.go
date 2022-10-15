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

		err = pegaProdutos.Scan(&Id, &Nome, &Preco, &Quantidade)

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

func CriandoProdutos(nome string, preco float64, quantidade int) {

	db := db.ConectorBD()

	InsereDados, err := db.Prepare("insert into produtos(nome, preco, quantidade) values ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	InsereDados.Exec(nome, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {

	db := db.ConectorBD()

	Deletando, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	Deletando.Exec(id)

	defer db.Close()
}
