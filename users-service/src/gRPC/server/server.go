package server

import (
	"context"
	"log"

	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/gRPC/pb/users"
	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/services"
)

type Server struct {
	users.UnimplementedUsersServicesServer
}

func (s Server) Ping(context context.Context, msg *users.PingMsg) (*users.PongMsg, error) {
	log.Println("Ping")
	return &users.PongMsg{Ping: "Pong"}, nil
}

func (s Server)  Create(c context.Context,  u *users.CreateUserRequest) (*users.UserResponse, error){
	userPersisted, err := services.Create(u)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	

	return &users.UserResponse{
		Id: int64(userPersisted.ID),
		Name: userPersisted.Name,
		Nick: userPersisted.Nick,
		Email: userPersisted.Email,
		CreatedAt: userPersisted.CreatedAt.String(),
		UpdatedAt: userPersisted.UpdatedAt.String(),
	}, nil 

}