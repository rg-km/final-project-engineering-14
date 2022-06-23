package migrate

import (
	"database/sql"
	"log"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
)

func Migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(150) NOT NULL,
		password VARCHAR(255) NOT NULL,
		phone VARCHAR(50) NOT NULL,
		role VARCHAR(15) NOT NULL,
		is_login BOOLEAN NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP 
	);

	CREATE TABLE IF NOT EXISTS programming_languanges (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(15) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		question VARCHAR(255) NOT NULL,
		proglang_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (proglang_id) REFERENCES programming_languanges(id)
	);

	CREATE TABLE IF NOT EXISTS answers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		answer VARCHAR(50) NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS answers_attempts (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		answer VARCHAR(50) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
	_, err := db.Exec(query)
	helper.PanicIfErrorWithMessage("Error when migrate with error:", err)

	log.Println("Migration success")
}
