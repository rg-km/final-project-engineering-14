package drop

import (
	"database/sql"
	"log"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
)

func DropAllTable(db *sql.DB) {
	query := `
		DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS programming_languanges;
		DROP TABLE IF EXISTS questions;
		DROP TABLE IF EXISTS answers;
	`
	_, err := db.Exec(query)
	helper.PanicIfErrorWithMessage("Error when dropping table:", err)

	log.Println("Dropped table successfully")
}
