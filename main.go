package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "demo/handler"
)

// @schemes http
func main() {
	log.Print("starting server...")

	http.HandleFunc("/", handler.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
