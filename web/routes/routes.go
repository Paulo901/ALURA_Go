package routes

import (
	"net/http"
	"web/controller"
)

func PegaRota() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
}
