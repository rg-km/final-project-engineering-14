package domain

import "time"

type UserDomain struct {
	Id        uint32    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	IsLogin   bool      `json:"is_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
