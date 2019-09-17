package main

import (
	"fmt"
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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", PORT))

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	registerServices(srv)
	reflection.Register(srv)

	log.Printf(fmt.Sprintf("Protobuf Server running @ http://localhost:%s", PORT))

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
