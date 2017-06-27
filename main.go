package main

import (
	"net/http"

	"log"

	"github.com/restful/handlers"
	"github.com/restful/storage"
)

func main() {

	// Create DB instance
	db := storage.NewInMemoryDB()

	// Servemux is used to register route handlers
	mux := http.NewServeMux()
	mux.Handle("/get", handlers.GetKey(db))
	mux.Handle("/set", handlers.PutKey(db))

	log.Printf("Serving on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
