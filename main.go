package main

import (
	"log"
	"net/http"
	"reddit/route"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/board", route.GetBoards).Methods("GET")
	r.HandleFunc("/api/board/{id}", route.GetBoard).Methods("GET")
	r.HandleFunc("/api/board", route.CreateBook).Methods("POST")
	r.HandleFunc("/api/board/{id}", route.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/board/{id}", route.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}
