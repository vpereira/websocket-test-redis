package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	resp "github.com/nicklaw5/go-respond"
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// register client
	clients[ws] = true
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	render := resp.NewResponse(w)
	createJobList()
	render.Ok("{status: ok}")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	render := resp.NewResponse(w)
	FlushAllKeys()
	render.Ok("{status: ok}")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var job Job
	render := resp.NewResponse(w)

	err := json.NewDecoder(r.Body).Decode(&job)

	if err != nil {
		render.BadRequest(err.Error())
	}

	UpdateKV(job.ID.String(), job.Status)

	go writer(job)
	alljobs, _ := GetAllKeysValues()
	render.Ok(alljobs)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	render := resp.NewResponse(w)
	alljobs, _ := GetAllKeysValues()
	render.Ok(alljobs)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.tmpl"))
	alljobs, _ := GetAllKeysValues()
	data := IndexData{Jobs: alljobs}
	tmpl.Execute(w, data)
}
