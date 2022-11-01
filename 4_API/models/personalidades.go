package models

import "4_API/database"

type Personalidade struct {
	Id       int    `json: "id"`
	Nome     string `json: "nome"`
	Historia string `json: "historia"`
}

var Personalidades []Personalidade

func Adicionando(Id int, Nome, Historia string) {
	db := database.ConectDB()
	defer db.Close()
	InsereDados, err := db.Prepare("insert into personalidades (Id, Nome, Historia) values ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	InsereDados.Exec(Id, Nome, Historia)

}
