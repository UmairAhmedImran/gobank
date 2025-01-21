package main

import (
	"log"

	"github.com/UmairAhmedImran/gobank/internal/storage"
)


func main() {
	store, err := storage.NewPostgresStore()
	
	// only run when to Drop the table
	//if err := store.DropTable("account"); err != nil {
	//	log.Fatal(err)
	//}

	if err != nil {
    		log.Fatal(err)
  	}
	
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
