package main

import (
	"log"
	"net/http"
	"os"

	"library/internal/httpapi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h := httpapi.NewServer()
	log.Printf("Backend listening on :%s", port)
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal(err)
	}
}
