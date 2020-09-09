package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type IndexData struct {
	Jobs map[string]string
}

func createJobList() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 15; i++ {
		j := Job{ID: uuid.New(), Status: randomStatus()}
		InsertKV(string(j.ID.String()), j.Status)
	}
}

func randomStatus() string {
	statuses := []string{"waiting", "running", "failed", "succeeded"}
	return statuses[rand.Intn(len(statuses))]
}
