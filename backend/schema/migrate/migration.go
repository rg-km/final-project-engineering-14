package migrate

import (
	"context"
	"database/sql"
	"log"
	"time"

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
	`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, query)
	helper.PanicIfErrorWithMessage("Error when migrate with error:", err)

	log.Println("Migration success")
}
