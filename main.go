package main

import (
	"log"
	"exonova-go/core/server/http"
)

func main() {
	server := http.NewServer(":6969")
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}