package main

import (
	"net/http"
	"web/routes"
)

func main() {

	routes.PegaRota()
	http.ListenAndServe(":5040", nil)
}
