package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
)

func main() {
	finish := make(chan bool)
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}",HelloWorld).Methods("GET")
	
	go func() {
	    log.Fatal(http.ListenAndServe(":8080",router))
	}
	
	router2 := mux.NewRouter()
	router2.HandleFunc("/health", GetHealth).Methods("GET")
	go func() {
		log.Fatal(http.ListenAndServe(":9000", router2))
	}
	<- finish
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HelloWorld (w http.ResponseWriter, r *http.Request) {
	log.Println("Responding to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,"Hello:",name);
}
