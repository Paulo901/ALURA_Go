package main

import (
	"3_WEB/routes"
	"net/http"
)

func main() {

	routes.PegaRota()
	http.ListenAndServe(":5040", nil)
}
