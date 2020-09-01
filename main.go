package main

import (
	"fmt"
	"log"

	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := NewRouter()
	
	http.Handle("/", r)
	
	// serve and check for error
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Choose the folder to serve
	staticDir := "/"

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	return router
}