package main

import (
	"log"

	"github.com/alyunov/go_yap_final_prj/pkg/db"
	"github.com/alyunov/go_yap_final_prj/pkg/server"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Failed to create bd: %v", err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
