package main

import (
	"log"

	"github.com/UmairAhmedImran/gobank/internal/storage"
)


func main() {
	store, err := storage.NewPostgresStore()

	if err != nil {
    		log.Fatal(err)
  	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
