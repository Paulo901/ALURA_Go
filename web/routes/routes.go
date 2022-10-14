package routes

import (
	"net/http"
	"web/controller"
)

func PegaRota() {
	http.HandleFunc("/", controller.Index)
}
