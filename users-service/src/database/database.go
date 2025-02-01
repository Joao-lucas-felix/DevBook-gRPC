package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Driver SQLite3
)

// SetupSqlLite3 create a connection with the database and create the table if not exists
func SetupSQLite3(dbPath string) (*sql.DB, error) {
	// Conectar ao banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error while trying to open the database: %w", err)
	}

	if err := db.Ping(); err != nil{
		return nil, err
	}
	// create table `user`
	createTableUsersQuery := `
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(250) NOT NULL,
			nick VARCHAR(250) NOT NULL UNIQUE,
			email VARCHAR(250) NOT NULL UNIQUE,
			password_hash VARCHAR(500) NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP

		);
	`
	createTableFollowersQuery := `
		CREATE TABLE followers (
			user_id INT NOT NULL,
			follower_id INT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
			PRIMARY KEY (user_id, follower_id)
		);
	`

	_, err = db.Exec(createTableUsersQuery)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("erro while trying to create the table users: %w", err)
	}
	_, err = db.Exec(createTableFollowersQuery)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("erro while trying to create the table followers: %w", err)
	}
	

	log.Println("Data base are successfully configured")
	return db, nil
}

// GenConn connect with the database
func GenConn(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("erro while trying to open the database: %w", err)
	}

	if err := db.Ping(); err != nil{
		return nil, err
	}
	return db, nil
}