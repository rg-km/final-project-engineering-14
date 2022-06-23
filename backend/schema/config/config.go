package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rg-km/final-project-engineering-14/backend/helper"
)

const (
	dbname = "fpe_14"
)

func DBConnect() (*sql.DB, error) {
	os.RemoveAll("schema/" + dbname + ".db")

	log.Println("Creating " + dbname + ".db...")
	file, err := os.Create("schema/" + dbname + ".db")
	helper.PanicIfErrorWithMessage("Error creating database file: ", err)
	file.Close()

	log.Println(dbname + ".db created")

	db, err := sql.Open("sqlite3", "schema/"+dbname+".db")
	helper.PanicIfErrorWithMessage("Error when connecting to database:", err)

	// Sett DB Pool
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	helper.PanicIfErrorWithMessage("Erorr when ping database : ", err)

	log.Printf("Connected to DB %s successfully\n", dbname)

	return db, nil
}
