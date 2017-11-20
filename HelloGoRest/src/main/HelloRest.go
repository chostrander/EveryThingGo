package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}",HelloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",router))

}

func HelloWorld (w http.ResponseWriter, r *http.Request) {
	log.Println("Responding to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,"Hello:",name);
}
