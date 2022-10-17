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
		var id, quantidade int
		var nome string
		var preco float64

		err = pegaProdutos.Scan(&id, &nome, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		prod.Id = id
		prod.Quantidade = quantidade
		prod.Nome = nome
		prod.Preco = preco

		produtos = append(produtos, prod)

	}
	defer db.Close()
	return produtos
}

func CriandoProdutos(id int, nome string, preco float64, quantidade int) {

	db := db.ConectorBD()

	InsereDados, err := db.Prepare("insert into produtos(id, nome, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	InsereDados.Exec(id, nome, preco, quantidade)
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
