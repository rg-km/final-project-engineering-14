package domain

import "time"

type UserDomain struct {
	Id        int32
	Username  string
	Email     string
	Password  string
	Phone     string
	Role      string
	IsLogin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProgrammingLanguangeDomain struct {
	Id        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type QuestionDomain struct {
	Id         int32
	Question   string
	ProglangId int32
	AnswerId   int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type AnswersDomain struct {
	Id        int32
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
