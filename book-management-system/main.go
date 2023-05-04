package main

import (
	"book-management-system/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.Handle("/", r)
	log.Println("Starting server at port 8080:")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
