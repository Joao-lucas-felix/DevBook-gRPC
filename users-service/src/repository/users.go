package repository

import (
	"database/sql"
	"fmt"

	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/models"
	"github.com/Joao-lucas-felix/DevBook-gRPC/users-service/src/security"
)

// Repository is the struct to manipulate the data in the database
type Repository struct {
	db *sql.DB
}

var queries = map[string]string{

	"create": " INSERT INTO users (name, nick, email, password_hash) VALUES (?, ?, ?, ?); ",

	"update": " UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ? ",

	"delete": " DELETE FROM users u WHERE u.id = ? ",

	"findAll": " SELECT id, name, nick ,email, created_at, updated_at FROM users u WHERE u.name LIKE ? OR u.nick LIKE ? ", //format the string to %% s %%

	"findById": " SELECT id, name, nick ,email, created_at, updated_at FROM users u WHERE u.id = ? ",

	"findByNick": " SELECT id, name, nick ,email, created_at, updated_at FROM users u WHERE u.nick LIKE ?  ",

	"findByEmail": " SELECT id, name, nick ,email, created_at, updated_at FROM users u WHERE u.email = ? ",
}

// NewUserRepository generate a news user repository using the db
func NewUserRepository(db *sql.DB) *Repository {

	return &Repository{db: db}
}

// Create persists a user in the database
func (repository Repository) Create(user models.User) (models.User, error) {
	statement, err := repository.db.Prepare(queries["create"])
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	hashedBytes, err := security.Hash(user.Password)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedBytes)

	_, err = statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return models.User{}, err
	}
	row, err := repository.db.Query(queries["findByEmail"], user.Email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()
	var userToReturn models.User
	if row.Next() {
		if err := row.Scan(
			&userToReturn.ID,
			&userToReturn.Name,
			&userToReturn.Nick,
			&userToReturn.Email,
			&userToReturn.CreatedAt,
			&userToReturn.UpdatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return userToReturn, nil
}

// Update persists a user in the database
func (repository Repository) Update(user models.User) (models.User, error) {
	statement, err := repository.db.Prepare(queries["create"])
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	return user, nil
}

// Delete persists a user in the database
func (repository Repository) Delete(id int64) (models.User, error) {
	statement, err := repository.db.Prepare(queries["create"])
	if err != nil {
		return models.User{}, err
	}
	defer statement.Close()

	return models.User{}, nil
}

// FindAll persists a user in the database
func (repository Repository) FindAll(nameOrNick string) ([]models.User, error) {

	//format the string to %% s %%
	querieFormatedString := fmt.Sprintf("%%%s%%", nameOrNick)
	rows, err := repository.db.Query(queries["findAll"], querieFormatedString, querieFormatedString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		// SELECT id, name, nick ,email, created_at, updated_at
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository Repository) FindById(userID int64) (models.User, error) {

	rows, err := repository.db.Query(queries["findById"], userID)
	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()
	var user models.User
	// SELECT id, name, nick ,email, created_at, updated_at
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
