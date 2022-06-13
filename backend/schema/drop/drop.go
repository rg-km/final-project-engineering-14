package drop

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
)

func DropAllTable(db *sql.DB) {
	query := `
		DROP TABLE IF EXISTS users;
	`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, query)
	helper.PanicIfErrorWithMessage("Error when dropping table:", err)

	log.Println("Dropped table successfully")
}
