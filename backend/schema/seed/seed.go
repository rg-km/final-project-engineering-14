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
	(name, url_images, created_at, updated_at)
	VALUES
	('Go', 'https://www.linkpicture.com/q/go_1.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Python', 'https://www.linkpicture.com/q/python_1.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Java', 'https://www.linkpicture.com/q/java_3.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('C#', 'https://www.linkpicture.com/q/cs_2.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Ruby', 'https://www.linkpicture.com/q/ruby_1.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('PHP', 'https://www.linkpicture.com/q/php_2.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Rust', 'https://www.linkpicture.com/q/rust_5.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('JavaScript', 'https://www.linkpicture.com/q/js_16.png', '2022-06-01 00:00:00', '2022-06-01 00:00:00');

	INSERT INTO questions
	(question, proglang_id, created_at, updated_at)
	VALUES
	('How well do you understand about variable scope in Golang ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('How well do you understand using the for loop in Golang ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Are you familiar with how array, slice, and maps work and use in Golang ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('How well do you know about packages in Go programs ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('How well do you understand about variable scope in Golang ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Are you familiar with how struct work and use in Golang ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Are you familiar with how interface work and use in Golang ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Do you have an understanding of unit test ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Do you have an understanding of Web Servers in Golang ?', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00');
	
	INSERT INTO answers
	(answer, created_at, updated_at)
	VALUES
	('Really Understand', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Understand', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Do Not Understand', '2022-06-01 00:00:00', '2022-06-01 00:00:00');

	INSERT INTO levels 
	(level, created_at, updated_at)
	VALUES
	('Beginner', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('Intermediate', '2022-06-01 00:00:00', '2022-06-01 00:00:00');

	INSERT INTO recommendations
	(image_url, level_id, created_at, updated_at)
	VALUES
	('https://www.linkpicture.com/q/Basic_1.png', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
	('https://www.linkpicture.com/q/Intermediate.png', 2, '2022-06-01 00:00:00', '2022-06-01 00:00:00');
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
