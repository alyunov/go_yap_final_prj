package main

import (
	"log"

	"github.com/alyunov/go_yap_final_prj/pkg/server"
)

func main() {

	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
