package repository

import (
	"database/sql"
	"strings"
	"time"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

type BackendRepositorySQLite struct {
	DB *sql.DB
}

func NewBackendRepository(db *sql.DB) *BackendRepositorySQLite {
	return &BackendRepositorySQLite{
		DB: db,
	}
}

func (repo *BackendRepositorySQLite) SaveQuestion(question domain.QuestionDomain, proglang domain.ProgrammingLanguangeDomain, userId int32) (domain.QuestionDomain, error) {
	query := `
	INSERT INTO questions 
	(question, created_at, updated_at, proglang_id) 
	VALUES (?, ?, ?, ?);`

	result, err := repo.DB.Exec(query,
		question.Question, question.CreatedAt,
		question.UpdatedAt, proglang.Id,
	)
	helper.PanicIfError(err)

	questionId, err := result.LastInsertId()
	helper.PanicIfError(err)

	question.Id = int32(questionId)

	return question, nil
}

func (repo *BackendRepositorySQLite) GetQuestions() ([]web.QuestionRequest, error) {
	query := `
	SELECT q.id, q.question AS question, pl.name AS programming_lang
	FROM questions AS q
	INNER JOIN programming_languanges AS pl ON pl.id = q.proglang_id;`

	rows, err := repo.DB.Query(query)
	helper.PanicIfError(err)
	defer rows.Close()

	var questions []web.QuestionRequest
	for rows.Next() {
		var question web.QuestionRequest
		err := rows.Scan(
			&question.Id, &question.Question, &question.ProgrammingLanguange,
		)
		helper.PanicIfError(err)

		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		helper.PanicIfError(err)
	}

	return questions, nil
}

func (repo *BackendRepositorySQLite) GetCountQuestions() ([]web.CountQuestionResponse, error) {
	query := `
	SELECT pl.name AS proglang_name, COUNT(q.id) AS Total
	FROM questions AS q
	INNER JOIN programming_languanges AS pl ON pl.id = q.proglang_id
	GROUP BY pl.id;`

	rows, err := repo.DB.Query(query)
	helper.PanicIfError(err)
	defer rows.Close()

	var dashboardAdmin []web.CountQuestionResponse
	for rows.Next() {
		var dashboard web.CountQuestionResponse
		err := rows.Scan(&dashboard.ProgrammingLanguage, &dashboard.TotalQuestion)
		helper.PanicIfError(err)

		dashboardAdmin = append(dashboardAdmin, dashboard)
	}

	if err := rows.Err(); err != nil {
		helper.PanicIfError(err)
	}

	return dashboardAdmin, nil
}

func (repo *BackendRepositorySQLite) UpdateQuestion(proglang domain.ProgrammingLanguangeDomain, question web.QuestionRequest, questionId int32) (web.QuestionResponse, error) {
	query := "UPDATE questions SET "
	qParts := []string{}
	args := []interface{}{}

	if question.Question != "" {
		qParts = append(qParts, "question = ?")
		args = append(args, question.Question)
	}

	if question.ProgrammingLanguange != "" {
		qParts = append(qParts, "proglang_id = ?")
		args = append(args, proglang.Id)
	}

	query += strings.Join(qParts, ", ")
	query += ", updated_at = ? WHERE id = ?"
	args = append(args, time.Now(), questionId)

	result, err := repo.DB.Exec(query, args...)
	helper.PanicIfError(err)

	_, err = result.LastInsertId()
	helper.PanicIfError(err)

	var questionResponse web.QuestionResponse

	// select question
	query = `SELECT q.id, q.question, pl.name
	FROM questions AS q
	INNER JOIN programming_languanges AS pl ON pl.id = q.proglang_id
	WHERE q.id = ?;`

	row := repo.DB.QueryRow(query, questionId)
	helper.PanicIfError(err)

	err = row.Scan(
		&questionResponse.Id, &questionResponse.Question,
		&questionResponse.ProgrammingLanguange,
	)
	helper.PanicIfError(err)

	return questionResponse, nil
}

func (repo *BackendRepositorySQLite) DeleteQuestion(questionId int32) (bool, error) {
	query := `DELETE FROM questions WHERE id = ?`

	_, err := repo.DB.Exec(query, questionId)
	helper.PanicIfError(err)

	return true, nil
}

func (repo *BackendRepositorySQLite) GetAnswers() ([]domain.AnswersDomain, error) {
	query := `SELECT a.id, a.answer, a.created_at, a.updated_at FROM answers AS a`

	rows, err := repo.DB.Query(query)
	helper.PanicIfError(err)
	defer rows.Close()

	var answers []domain.AnswersDomain
	for rows.Next() {
		var answer domain.AnswersDomain
		err := rows.Scan(
			&answer.Id, &answer.Answer, &answer.CreatedAt, &answer.UpdatedAt,
		)
		helper.PanicIfError(err)

		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
		helper.PanicIfError(err)
	}

	return answers, nil
}
