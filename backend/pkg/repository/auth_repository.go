package repository

import (
	"database/sql"
	"errors"

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

func (repo *AuthRepositorySQLite) Save(userReq domain.UserDomain, email string) (domain.UserDomain, error) {
	var user domain.UserDomain

	query := `SELECT id, email FROM users WHERE email = ?`
	row := repo.DB.QueryRow(query, email)
	err := row.Scan(&user.Id, &user.Email)
	if err == nil {
		return user, errors.New("email already exists")
	}

	query = `
	INSERT INTO users 
	(username, email, password, phone, role, is_login, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	result, err := repo.DB.Exec(query,
		userReq.Username, email, userReq.Password, userReq.Phone, userReq.Role,
		userReq.IsLogin, userReq.CreatedAt, userReq.UpdatedAt,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int32(id)
	return userReq, nil
}

func (repo *AuthRepositorySQLite) GetUser(email, password string) (domain.UserDomain, error) {
	query := ` 
	SELECT id, username, email, password, phone, role, is_login 
	FROM users 
	WHERE email = ? AND password = ?`

	var user domain.UserDomain
	row := repo.DB.QueryRow(query, email, password)
	err := row.Scan(
		&user.Id, &user.Username, &user.Email, &user.Password,
		&user.Phone, &user.Role, &user.IsLogin,
	)
	if err != nil {
		return user, errors.New("login failed")
	}

	query = `UPDATE users SET is_login = true WHERE id = ?`
	_, err = repo.DB.Exec(query, user.Id)
	helper.PanicIfError(err)

	return user, nil
}

func (repo *AuthRepositorySQLite) GetUserById(userId int32) (domain.UserDomain, error) {
	query := `
	SELECT id, username, email, phone, role, is_login, created_at, updated_at 
	FROM users WHERE id = ?`

	var user domain.UserDomain
	row := repo.DB.QueryRow(query, userId)
	err := row.Scan(
		&user.Id, &user.Username, &user.Email, &user.Phone, &user.Role, &user.IsLogin, &user.CreatedAt, &user.UpdatedAt,
	)
	helper.PanicIfError(err)

	return user, nil
}

func (repo *AuthRepositorySQLite) Logout(userId int32) (bool, error) {
	query := `UPDATE users SET is_login = false WHERE id = ?`

	_, err := repo.DB.Exec(query, userId)
	helper.PanicIfError(err)

	return true, nil
}
