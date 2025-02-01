package services

import (
	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/database"
	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/gRPC/pb/users"
	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/models"
	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/repository"
)
 
const DBPATH = "./users-service/src/database/users.db"

// Create receives the gRPC message and persist the data in the DB
func Create( u *users.CreateUserRequest ) (models.User, error) {
	var user = models.User{
		Name: u.Name,
		Nick: u.Nick,
		Email: u.Email,
		Password: u.Password,
	}
	db, err := database.GenConn(DBPATH)
	if err != nil {
		return models.User{}, err
	}
	r  := repository.NewUserRepository(db)
	return r.Create(user)
}