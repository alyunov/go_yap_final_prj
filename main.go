package main

import (
	"log"
	"net/http"
)

func main() {
	webDir := "./web"

	http.Handle("/", http.FileServer(http.Dir(webDir)))

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
