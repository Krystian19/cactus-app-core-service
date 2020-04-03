package main

import (
	"log"
	"os"
)

const (
	// Default port for the GRPC Server
	defaultPort = "9040"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = defaultPort
	}

	log.Printf("GRPC Server running @ http://localhost:%s", PORT)
	if e := Server(PORT); e != nil {
		log.Fatal(e)
	}
}
