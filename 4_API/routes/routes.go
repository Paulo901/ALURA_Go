package routes

import (
	"4_API/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	r.HandleFunc("/api/personalidades/{id}", controllers.OnePersonalidade)
	r.HandleFunc("/cria/personalidade", controllers.Insert)

	log.Fatal(http.ListenAndServe(":8000", r))
}
