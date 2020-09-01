package main

import (
	"fmt"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my simple website!\n")
}

func AddHandler(w http.ResponseWriter, r *http.Request) {

	// get url variables
    vars := mux.Vars(r)
    num1, err1 := strconv.Atoi(vars["num1"])
	num2, err2 := strconv.Atoi(vars["num2"])
	
	// check for error
	if err1 != nil || err2 != nil {
		fmt.Fprintf(w, "Sorry, there was an error. :(")
	} else {
		fmt.Fprintf(w, "%d\n", num1 + num2)
	}
}

func main() {
	r := mux.NewRouter()
	
    r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/add/{num1}/{num2}", AddHandler)
	
	http.Handle("/", r)
	
	// serve and check for error
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}