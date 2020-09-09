package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serverAddress := ":3000"
	NewRedisClient()
	router := mux.NewRouter()

	router.HandleFunc("/ws", wsHandler)
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/update.json", updateHandler).Methods("POST")
	router.HandleFunc("/list.json", listHandler).Methods("GET")
	router.HandleFunc("/create.json", createHandler).Methods("POST")
	router.HandleFunc("/delete.json", deleteHandler).Methods("DELETE")

	log.Printf("server started at %s\n", serverAddress)
	go echo()
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
