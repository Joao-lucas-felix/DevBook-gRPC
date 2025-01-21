package server

import (
	"log"
	"net"

	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/gRPC/pb/users"
	"google.golang.org/grpc"
)

func BootstrapServer() error {
	
	log.Println("Starting gRPC server in the prot :9090")
	listner, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	users.RegisterUsersServicesServer(s, &Server{})
	
	return s.Serve(listner)
}
