package server

import (
	"context"
	"log"

	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/gRPC/pb/users"
)

type Server struct {
	users.UnimplementedUsersServicesServer
}

func (s Server) Ping(context context.Context, msg *users.PingMsg) (*users.PongMsg, error) {
	log.Println("Ping")
	return &users.PongMsg{Ping: "Pong"}, nil
}
