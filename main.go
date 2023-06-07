package main

import (
	"log"
)

func main() {

	store, err := NewPostgresStore()
	isAccDropEnable := false

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(isAccDropEnable); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
