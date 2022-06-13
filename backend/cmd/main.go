package main

import (
	"log"

	"github.com/rg-km/final-project-engineering-14/backend"
	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/handler"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/service"
	"github.com/rg-km/final-project-engineering-14/backend/schema/config"
	"github.com/rg-km/final-project-engineering-14/backend/schema/drop"
	"github.com/rg-km/final-project-engineering-14/backend/schema/migrate"
	"github.com/rg-km/final-project-engineering-14/backend/schema/seed"
)

func main() {
	db, err := config.DBConnect()
	helper.PanicIfErrorWithMessage("Error when connecting to database:", err)
	defer db.Close()

	drop.DropAllTable(db)
	migrate.Migrate(db)
	seed.Seed(db)

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(backend.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("Error occured while running server: ", err.Error())
	}
}
