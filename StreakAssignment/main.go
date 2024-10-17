package main

import (
	"log"
	"net/http"
	"task/findpairs/handlers"
)

func main() {
	http.HandleFunc("/find-pairs", handlers.FindPairs)

	// Starting the server
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
   }
