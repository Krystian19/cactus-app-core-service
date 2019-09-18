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
