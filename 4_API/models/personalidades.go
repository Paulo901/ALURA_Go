package models

import "4_API/db"

type Personalidade struct {
	Id       int    `json: "id"`
	Nome     string `json: "nome"`
	Historia string `json: "historia"`
}

var Personalidades []Personalidade

func Adicionando(Id int, Nome, Historia string) {
	db := db.ConectDB()
	defer db.Close()
	InsereDados, err := db.Prepare("insert into personalidades (Id, Nome, Historia) values ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	InsereDados.Exec(Id, Nome, Historia)

}
