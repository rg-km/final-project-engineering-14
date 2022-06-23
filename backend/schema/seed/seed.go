package seed

import (
	"database/sql"
	"log"
	"time"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/security"
)

func Seed(db *sql.DB) {
	query := `
	INSERT INTO users 
	(username, email, password, phone, role, is_login, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);

	INSERT INTO programming_languanges
	(name, created_at, updated_at)
	VALUES
	('Go', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Python', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Java', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('C#', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Ruby', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('PHP', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Kotlin', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Rust', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Scala', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('JavaScript', '2020-01-01 00:00:00', '2020-01-01 00:00:00');
	
	INSERT INTO answers
	(answer, created_at, updated_at)
	VALUES
	('Really Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
	('Do Not Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00');
	`

	hashed := security.GeneratePasswordHash("admin123")
	formatHours := "2006-01-02 15:04:05"

	args := []interface{}{
		"admin",
		"admin@gmail.com",
		hashed,
		"081234567890",
		"admin",
		false,
		time.Now().Format(formatHours),
		time.Now().Format(formatHours),
	}

	_, err := db.Exec(query, args...)
	helper.PanicIfErrorWithMessage("Error when seed with error:", err)

	log.Println("Seeding success")
}
