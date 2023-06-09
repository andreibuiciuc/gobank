package main

import (
	"log"
)

func main() {

	store, err := NewPostgresStore()
	isAccDropEnabled := false

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(isAccDropEnabled); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3001", store)
	server.Run()
}
