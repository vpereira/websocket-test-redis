package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Job)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo() {
	for {
		job := <-broadcast
		jobJson, _ := json.Marshal(job)
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(string(jobJson)))
			if err != nil {
				log.Fatal(err)
				client.Close()
				log.Fatal("client removed")
				delete(clients, client)
			}
		}
	}
}

func writer(j Job) {
	broadcast <- j
}
