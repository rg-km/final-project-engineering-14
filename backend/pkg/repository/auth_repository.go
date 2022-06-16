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

func (repo *AuthRepositorySQLite) GetUser(ctx context.Context, email, password string) (domain.UserDomain, error) {
	tx, err := repo.DB.Begin()
	helper.PanicIfError(err)

	querySelectALl := ` 
	SELECT id, username, email, password, phone, role, is_login 
	FROM users 
	WHERE email = ? AND password = ?`

	var user domain.UserDomain
	stmt, err := tx.PrepareContext(ctx, querySelectALl)
	helper.PanicIfError(err)
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, email, password)
	err = row.Scan(
		&user.Id, &user.Username, &user.Email, &user.Password,
		&user.Phone, &user.Role, &user.IsLogin,
	)
	if err != nil {
		tx.Rollback()
		return domain.UserDomain{}, err
	}

	queryUpdate := `UPDATE users SET is_login = true WHERE id = ?`
	stmt, err = tx.PrepareContext(ctx, queryUpdate)
	helper.PanicIfError(err)
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Id)
	if err != nil {
		tx.Rollback()
		return domain.UserDomain{}, err
	}

	return user, tx.Commit()
}

func (repo *AuthRepositorySQLite) Logout(ctx context.Context, userId uint32) (bool, error) {
	tx, err := repo.DB.Begin()
	helper.PanicIfError(err)

	query := `UPDATE users SET is_login = false WHERE id = ?`

	stmt, err := tx.PrepareContext(ctx, query)
	helper.PanicIfError(err)
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit()
}
