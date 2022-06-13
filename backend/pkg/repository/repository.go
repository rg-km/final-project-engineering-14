package repository

import (
	"context"
	"database/sql"

	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
)

type AuthRepository interface {
	Save(ctx context.Context, user domain.UserDomain) (domain.UserDomain, error)
}

type Repository struct {
	AuthRepository AuthRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
	}
}
