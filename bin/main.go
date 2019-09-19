package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "9040"
	}

	// Check for the important env vars to be present
	EnvVarsCheck()

	listener, err := net.Listen("tcp", ":"+PORT)

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	registerServices(srv)
	reflection.Register(srv)

	log.Printf("Protobuf Server running @ http://localhost:%s", PORT)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

// EnvVarsCheck : Checks that important ENV vars are set
func EnvVarsCheck() {
	if os.Getenv("DB_HOST") == "" {
		panic("DB_HOST env var is not set")
	}

	if os.Getenv("DB_NAME") == "" {
		panic("DB_NAME env var is not set")
	}

	if os.Getenv("DB_USERNAME") == "" {
		panic("DB_USERNAME env var is not set")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		panic("DB_PASSWORD env var is not set")
	}
}
