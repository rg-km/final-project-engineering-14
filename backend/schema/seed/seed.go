package seed

import (
	"context"
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
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

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

	_, err := db.ExecContext(ctx, query, args...)
	helper.PanicIfErrorWithMessage("Error when seed with error:", err)

	log.Println("Seeding success")
}
