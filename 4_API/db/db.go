package db

import "database/sql"

func ConectDB() *sql.DB {

	conexao := "user = postgres  dbname = Personalidades password = 2270 host = localhost  sslmode = false"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
