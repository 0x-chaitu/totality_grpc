package main

import (
	"log"
	"net"
	"totality_corp_kv/internal/repository"
	"totality_corp_kv/internal/service"
	"totality_corp_kv/pkg/kvpair"

	"google.golang.org/grpc"
)

func main() {
	const addr = "0.0.0.0:50051"

	// create a TCP listener on the specified port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server instance
	server := grpc.NewServer()

	c := kvpair.NewCache()
	repos := repository.NewRepositories(c)

	service.NewServices(service.Deps{
		Repos:  repos,
		Server: server,
	})

	// start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
