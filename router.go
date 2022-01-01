package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/handlers"
	"net/http"
)

func InitRouting() {
	api := mux.NewRouter()

	api.HandleFunc("/", handlers.SayHello).Methods(http.MethodGet)

	log.Println("Listening on :80")
	log.Fatal(http.ListenAndServe(":80", api))
}
