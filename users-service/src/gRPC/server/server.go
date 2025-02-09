package server

import (
	"context"
	"errors"
	"log"

	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/gRPC/pb/users"
	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/services"
)

type Server struct {
	users.UnimplementedUsersServicesServer
}

// Ping respond with a message to check if the server are up
func (s Server) Ping(context context.Context, msg *users.PingMsg) (*users.PongMsg, error) {
	log.Println("Ping")
	return &users.PongMsg{Ping: "Pong"}, nil
}

// Create receive a message that contains the user data, persist and returns the same data if the operation are successfully.
func (s Server) Create(c context.Context, u *users.CreateUserRequest) (*users.UserResponse, error) {
	userPersisted, err := services.Create(u)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &users.UserResponse{
		Id:        int64(userPersisted.ID),
		Name:      userPersisted.Name,
		Nick:      userPersisted.Nick,
		Email:     userPersisted.Email,
		CreatedAt: userPersisted.CreatedAt.String(),
		UpdatedAt: userPersisted.UpdatedAt.String(),
	}, nil

}

// GetAll receive a message with the name or nick of the user and return with a stream of all users where the name or nick are like the received.
func (s Server) GetAll(req *users.GetAllUsersRequest, srv users.UsersServices_GetAllServer) error {
	nameOrNick := req.GetNameOrNick()
	usersGet, err := services.GetAll(nameOrNick)
	if err != nil {
		return err
	}

	for i := range usersGet {
		userI := &users.UserResponse{
			Id:        int64(usersGet[i].ID),
			Name:      usersGet[i].Name,
			Nick:      usersGet[i].Nick,
			Email:     usersGet[i].Email,
			CreatedAt: usersGet[i].CreatedAt.String(),
			UpdatedAt: usersGet[i].UpdatedAt.String(),
		}

		if err := srv.Send(userI); err != nil {
			return err
		}
	}
	return nil
}

// GetById receive a message with the ID of the user and return the data of this user if exists. 
func (s Server) GetById(ctx context.Context, req *users.UserIdRequest) (*users.UserResponse, error) {
	userID := req.GetUserId()
	usersGet, err := services.GetById(userID)
	if err != nil {
		return nil, err
	}
	if usersGet.ID == 0 {
		return nil, errors.New("the user does not exists")
	}
	return &users.UserResponse{
		Id:        int64(usersGet.ID),
		Name:      usersGet.Name,
		Nick:      usersGet.Nick,
		Email:     usersGet.Email,
		CreatedAt: usersGet.CreatedAt.String(),
		UpdatedAt: usersGet.UpdatedAt.String(),
	}, nil
}

// Update receive a message with the data to update a specific user
func (s Server) Update(ctx context.Context, req *users.UpdateUserRequest) (*users.UserResponse, error){
	user, err :=  services.Update(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &users.UserResponse{
		Id:        int64(user.ID),
		Name:      user.Name,
		Nick:      user.Nick,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}