package main

import (
	"log"
	"net/http"
)

func main() {

	log.Printf("Server started...")
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatalf("Error to start server: %v", err)
	}
}
