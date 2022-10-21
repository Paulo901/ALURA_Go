package routes

import (
	"4_API/controllers"
	"log"
	"net/http"
)

func HandleRequest() {

	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":5040", nil))
}
