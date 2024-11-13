package main

import (
	"endpointlab/api"
	"log"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	server.Run(":8080")
}


