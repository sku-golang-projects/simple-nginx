package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func userHomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to user page."))
}

func adminHomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to admin page."))
}

func main() {
	router := mux.NewRouter()

	basicRoute := router.PathPrefix("/v1/user").Subrouter()
	basicRoute.HandleFunc("", userHomePage).Methods("GET")

	adminRoute := router.PathPrefix("/v1/admin").Subrouter()
	adminRoute.HandleFunc("", adminHomePage).Methods("GET")

	log.Println("Running server on port 8000.")
	log.Fatalln(http.ListenAndServe(":8000", router))
}
