package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode("Hello World!!")
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", HelloWorld()).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
