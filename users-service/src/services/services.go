package services

import (
	"time"

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
// Update receives the gRPC message and update the user data 
func Update (u *users.UpdateUserRequest) (models.User, error) {
	userSaved, err := GetById(u.Id)
	if err != nil {
		return models.User{}, err
	}
	var user = models.User{
		ID: uint64(u.Id),
		Name: u.Name,
		Nick: u.Nick,
		Email: u.Email,
	}
	db, err := database.GenConn(DBPATH)
	if err != nil {
		return models.User{}, err
	}
	r := repository.NewUserRepository(db)
	userReturn := models.User{
		ID: userSaved.ID,
		Name: user.Name,
		Nick: user.Nick,
		Email: user.Email,
		CreatedAt: userSaved.CreatedAt,
		UpdatedAt: time.Now(),
	}
	if err := r.Update(user); err != nil {
		return models.User{}, err
	}

	return userReturn, nil
}

// GetAll get all users 
func GetAll(nickOrEmail string) ([]models.User, error) {
	db, err := database.GenConn(DBPATH)
	if err != nil {
		return nil, err
	}
	r := repository.NewUserRepository(db)
	return r.FindAll(nickOrEmail)
}
// GetById get a specific user
func GetById(userID int64) (models.User, error) {
	db, err := database.GenConn(DBPATH)
	if err != nil {
		return models.User{}, err
	}
	r := repository.NewUserRepository(db)
	return r.FindById(userID)
}

