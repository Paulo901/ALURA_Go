package models

import "3_WEB/db"

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

func EditaProduto(id string) Produto {

	db := db.ConectorBD()

	DadosProduto, err := db.Query("select * from produtos while id=$1" + id)
	if err != nil {
		panic(err.Error())
	}
	ProdutoAtualizar := Produto{}

	for DadosProduto.Next() {

		var nome string
		var preco float64
		var quantidade, id int

		err = DadosProduto.Scan(&id, &nome, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		ProdutoAtualizar.Id = id
		ProdutoAtualizar.Nome = nome
		ProdutoAtualizar.Preco = preco
		ProdutoAtualizar.Quantidade = quantidade

	}
	defer db.Close()
	return ProdutoAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectorBD()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
