package main

import (
	"log"
	"fmt"

	"github.com/UmairAhmedImran/gobank/internal/storage"
)


func main() {
	store, err := storage.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)
	
	//server := NewAPIServer(":3000", store)
	//server.Run()
}
