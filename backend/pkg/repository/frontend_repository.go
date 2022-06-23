package repository

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-14/backend/helper"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
)

type FrontendRepositorySQLite struct {
	DB *sql.DB
}

func NewFrontendRepository(db *sql.DB) *FrontendRepositorySQLite {
	return &FrontendRepositorySQLite{
		DB: db,
	}
}

func (repo *FrontendRepositorySQLite) GetProgrammingLanguanges() ([]domain.ProgrammingLanguangeDomain, error) {
	query := `SELECT pl.id, pl.name FROM programming_languanges AS pl`

	rows, err := repo.DB.Query(query)
	helper.PanicIfError(err)

	var programmingLanguages []domain.ProgrammingLanguangeDomain
	for rows.Next() {
		var programmingLanguage domain.ProgrammingLanguangeDomain
		err = rows.Scan(&programmingLanguage.Id, &programmingLanguage.Name)
		helper.PanicIfError(err)

		programmingLanguages = append(programmingLanguages, programmingLanguage)
	}

	if err := rows.Err(); err != nil {
		helper.PanicIfError(err)
	}

	return programmingLanguages, nil
}

func (repo *FrontendRepositorySQLite) GetProgrammingLanguangeByName(progLangName string) (domain.ProgrammingLanguangeDomain, error) {
	query := `
	SELECT pl.id, pl.name, pl.created_at, pl.updated_at 
	FROM programming_languanges AS pl 
	WHERE name = ?`

	var programmingLanguage domain.ProgrammingLanguangeDomain
	row := repo.DB.QueryRow(query, progLangName)
	err := row.Scan(
		&programmingLanguage.Id,
		&programmingLanguage.Name,
		&programmingLanguage.CreatedAt,
		&programmingLanguage.UpdatedAt,
	)
	helper.PanicIfError(err)

	return programmingLanguage, nil
}

func (repo *FrontendRepositorySQLite) GetQuestionByProgrammingLanguange(progLangId, perPage, page int32) ([]web.QuestionRequest, error) {
	query := `
	SELECT DISTINCT q.question, pl.name
	FROM questions AS q
	INNER JOIN programming_languanges AS pl ON pl.id = q.proglang_id
	WHERE pl.id = ?
	ORDER BY q.id
	LIMIT ? OFFSET ?`

	rows, err := repo.DB.Query(query, progLangId, perPage, (page-1)*perPage)
	helper.PanicIfError(err)
	defer rows.Close()

	var questions []web.QuestionRequest
	for rows.Next() {
		var question web.QuestionRequest
		err := rows.Scan(&question.Question, &question.ProgrammingLanguange)
		helper.PanicIfError(err)

		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		helper.PanicIfError(err)
	}

	return questions, nil
}

func (repo *FrontendRepositorySQLite) SaveAnswersAttempt(userId int32, answers []domain.AnswersDomain) (bool, error) {
	query := `DELETE FROM answers_attempts;`
	_, err := repo.DB.Exec(query)
	helper.PanicIfError(err)

	query = `INSERT INTO answers_attempts (user_id, answer) VALUES (?, ?)`

	for _, answer := range answers {
		result, err := repo.DB.Exec(query, userId, answer.Answer)
		helper.PanicIfError(err)

		_, err = result.LastInsertId()
		helper.PanicIfError(err)
	}

	return true, nil
}

func (repo *FrontendRepositorySQLite) CountAnswersAttempt(userId int32) (int32, error) {
	query := `
	SELECT
		SUM(CASE at.answer 
        WHEN 'Really Understand' THEN 10
        WHEN 'Understand' THEN 5
        ELSE 0 END) AS TOTAL
	FROM answers_attempts AS at
	WHERE at.user_id = ?;`

	var total int32
	row := repo.DB.QueryRow(query, userId)
	err := row.Scan(&total)
	helper.PanicIfError(err)

	return total, nil
}
