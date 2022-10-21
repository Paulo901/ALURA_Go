package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "HOME")
}
func HandleRequest() {

	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":5040", nil))
}

func main() {

	HandleRequest()
}
