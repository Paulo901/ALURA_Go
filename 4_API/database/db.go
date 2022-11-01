package database

import (
	"database/sql"
	"log"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"log"
)

var (
	DB  *sql.DB //DB  *gorm.DB
	err error
)

func ConectDB() *sql.DB {

	conexao := "user = postgres  dbname = Personalidades password = 2270 host = localhost  sslmode = false"
	DB, err = sql.Open("postgres", conexao) //gorm.Open(postgres.Open(conexao)) -> com docker
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	return DB
}
