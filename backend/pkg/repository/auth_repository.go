package repository

import (
	"context"
	"database/sql"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
)

type AuthRepositorySQLite struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositorySQLite {
	return &AuthRepositorySQLite{
		DB: db,
	}
}

func (repo *AuthRepositorySQLite) Save(ctx context.Context, user domain.UserDomain) (domain.UserDomain, error) {
	query := `
	INSERT INTO users 
	(username, email, password, phone, role, is_login, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	done := make(chan bool, 1)
	go func(ch chan<- bool) {
		defer close(ch)

		stmt, err := repo.DB.PrepareContext(ctx, query)
		if err != nil {
			ch <- false
			return
		}
		defer stmt.Close()

		result, err := stmt.ExecContext(ctx,
			user.Username, user.Email, user.Password, user.Phone,
			user.Role, user.IsLogin, user.CreatedAt, user.UpdatedAt,
		)
		if err != nil {
			ch <- false
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			ch <- false
			return
		}

		user.Id = uint32(id)
		ch <- true
	}(done)

	if helper.OK(done) {
		return user, nil
	}

	return domain.UserDomain{}, nil
}
