package repository_test

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rg-km/final-project-engineering-14/backend/model/domain"
	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
)

var _ = Describe("Repository Test", func() {
	var db *sql.DB
	var err error
	var authRepo *repository.AuthRepositorySQLite
	var backendRepo *repository.BackendRepositorySQLite
	var frontendRepo *repository.FrontendRepositorySQLite

	BeforeEach(func() {
		db, err = sql.Open("sqlite3", "./repository_test.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			CREATE TABLE users (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				username TEXT,
				email TEXT,
				password TEXT,
				phone TEXT,
				role TEXT,
				is_login BOOLEAN,
				created_at DATETIME,
				updated_at DATETIME
			);
			
			CREATE TABLE IF NOT EXISTS answers (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				answer VARCHAR(50) NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			);

			CREATE TABLE IF NOT EXISTS questions (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				question VARCHAR(255) NOT NULL,
				proglang_id INTEGER NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (proglang_id) REFERENCES programming_languanges(id)
			);

			CREATE TABLE IF NOT EXISTS programming_languanges (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(15) NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			);

			CREATE TABLE IF NOT EXISTS answers_attempts (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER NOT NULL,
				answer VARCHAR(50) NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id)
			);

			INSERT INTO users (
				username, email, password, phone, role, is_login, created_at, updated_at
			) VALUES 
			('admin', 'admin@gmail.com', 'admin123', '081234567890', 'admin', 0, '2022-06-13 00:00:00', '2022-06-04 00:00:00'),
			('navis', 'navis@gmail.com', '1234', '081234567890', 'guest', 0, '2022-06-13 00:00:00', '2022-06-04 00:00:00');

			INSERT INTO answers
				(answer, created_at, updated_at)
			VALUES
			('Really Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Do Not Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00');

			INSERT INTO programming_languanges
				(name, created_at, updated_at)
			VALUES
			('Go', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Python', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Java', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('C#', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Ruby', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('PHP', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Kotlin', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Rust', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Scala', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('JavaScript', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('SQL', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Solidity', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Perl', '2020-01-01 00:00:00', '2020-01-01 00:00:00');

			INSERT INTO questions
				(question, proglang_id, created_at, updated_at)
			VALUES
			('How well do you know about packages in Go programs', 1, '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('How well do you understand using the for loop in Golang?', 1, '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Are you familiar with how array, slice, and maps work and use in Golang?', 1, '2020-01-01 00:00:00', '2020-01-01 00:00:00');

			INSERT INTO answers_attempts
				(user_id, answer)
			VALUES
			(1, 'Do Not Understand'),
			(1, 'Understand'),
			(1, 'Really Understand');

		`)

		if err != nil {
			panic(err)
		}

		authRepo = repository.NewAuthRepository(db)
		backendRepo = repository.NewBackendRepository(db)
		frontendRepo = repository.NewFrontendRepository(db)

	})

	AfterEach(func() {
		db, err = sql.Open("sqlite3", "./repository_test.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			DROP TABLE users;
			DROP TABLE answers;
			DROP TABLE questions;
			DROP TABLE programming_languanges;
			DROP TABLE answers_attempts;
		`)

		if err != nil {
			panic(err)
		}
	})

	Describe("GetUser", func() {
		When("email is valid", func() {
			It("should return user", func() {
				user, err := authRepo.GetUser("admin@gmail.com", "admin123")
				Expect(err).ToNot(HaveOccurred())

				fmt.Println("user: ", user)
				Expect(user.Username).To(Equal("admin"))
				Expect(user.Email).To(Equal("admin@gmail.com"))
				Expect(user.Password).To(Equal("admin123"))
				Expect(user.Phone).To(Equal("081234567890"))
				Expect(user.Role).To(Equal("admin"))
				Expect(user.IsLogin).To(Equal(false))
			})
		})

		When("email is invalid", func() {
			It("should return error", func() {
				_, err := authRepo.GetUser("adminhhua@gmail.com", "1234")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("login failed"))
			})
		})
	})

	Describe("Get User By Id", func() {
		When("id is valid", func() {
			It("should return user", func() {
				user, err := authRepo.GetUserById(1)
				Expect(err).ToNot(HaveOccurred())

				Expect(user.Username).To(Equal("admin"))
				Expect(user.Email).To(Equal("admin@gmail.com"))
			})
		})
	})

	Describe("Save", func() {
		When("Save user successfully", func() {
			It("should return user", func() {
				user, err := authRepo.Save(domain.UserDomain{
					Id:        2,
					Username:  "dwi",
					Email:     "dwi@gmail.com",
					Password:  "1234",
					Phone:     "081234567890",
					Role:      "guest",
					IsLogin:   false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, "dwi@gmail.com")
				Expect(err).ToNot(HaveOccurred())

				Expect(user.Username).To(Equal("dwi"))
				Expect(user.Email).To(Equal("dwi@gmail.com"))
				Expect(user.Password).To(Equal("1234"))
				Expect(user.Phone).To(Equal("081234567890"))
				Expect(user.Role).To(Equal("guest"))
				Expect(user.IsLogin).To(Equal(false))
			})
		})
	})

	Describe("Logout", func() {
		When("Logout user successfully", func() {
			It("Accept the logout", func() {
				user, err := authRepo.GetUser("admin@gmail.com", "admin123")
				Expect(err).ToNot(HaveOccurred())
				Expect(user.IsLogin).To(Equal(false))

				userIdLogin := user.Id
				isSuccesLogout, err := authRepo.Logout(userIdLogin)
				Expect(err).ToNot(HaveOccurred())
				Expect(isSuccesLogout).To(Equal(true))
			})
		})
	})

	Describe("Save Question", func() {
		When("Save Question successfully", func() {
			It("should return question", func() {
				question, err := backendRepo.SaveQuestion(domain.QuestionDomain{
					Id:         4,
					Question:   "Are you familiar with how pointer work and use in Golang?",
					ProglangId: 1,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, domain.ProgrammingLanguangeDomain{
					Id:        1,
					Name:      "Go",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, 1)

				Expect(err).ToNot(HaveOccurred())

				Expect(question.Question).To(Equal("Are you familiar with how pointer work and use in Golang?"))
				Expect(question.ProglangId).To(Equal(int32(1)))
			})
		})
	})

	Describe("Get Question", func() {
		When("Get Question successfully", func() {
			It("should return question", func() {
				question, err := backendRepo.GetQuestions()
				Expect(err).ToNot(HaveOccurred())

				Expect(question[0].Question).To(Equal("How well do you know about packages in Go programs"))
			})
		})
	})

	Describe("Get Count Question", func() {
		When("Get Count Question successfully", func() {
			It("should return count question", func() {
				countQuestion, err := backendRepo.GetCountQuestions()
				Expect(err).ToNot(HaveOccurred())

				Expect(countQuestion).To(Equal([]web.CountQuestionResponse{
					{
						ProgrammingLanguage: "Go",
						TotalQuestion:       3,
					},
				}))
			})
		})
	})

	Describe("Update Question", func() {
		When("Update Question successfully", func() {
			It("should return true", func() {
				updateSuccess, err := backendRepo.UpdateQuestion(domain.ProgrammingLanguangeDomain{
					Id:        1,
					Name:      "Go",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, web.QuestionRequest{
					ProgrammingLanguange: "Go",
					Question:             "Are you familiar with how pointer work and use in Golang?",
				}, int32(4))
				Expect(err).ToNot(HaveOccurred())

				Expect(updateSuccess).To(Equal(true))
			})
		})
	})

	Describe("Delete Question", func() {
		When("Delete Question successfully", func() {
			It("should return true", func() {
				deleteSuccess, err := backendRepo.DeleteQuestion(int32(3))
				Expect(err).ToNot(HaveOccurred())

				Expect(deleteSuccess).To(Equal(true))
			})
		})
	})

	Describe("GetAnswer", func() {
		When("GetAnswer successfully", func() {
			It("should return answer", func() {
				answer, err := backendRepo.GetAnswers()
				Expect(err).ToNot(HaveOccurred())

				Expect(answer[0].Answer).To(Equal("Really Understand"))
				Expect(answer[1].Answer).To(Equal("Understand"))
				Expect(answer[2].Answer).To(Equal("Do Not Understand"))
			})
		})
	})

	Describe("Get Programming Language", func() {
		When("Get Programming Language successfully", func() {
			It("should return programming language", func() {
				proglang, err := frontendRepo.GetProgrammingLanguanges()
				Expect(err).ToNot(HaveOccurred())

				Expect(proglang).To(Equal([]domain.ProgrammingLanguangeDomain{
					{Id: 1,
						Name:      "Go",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 2,
						Name:      "Python",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 3,
						Name:      "Java",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 4,
						Name:      "C#",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 5,
						Name:      "Ruby",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 6,
						Name:      "PHP",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 7,
						Name:      "Kotlin",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 8,
						Name:      "Rust",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 9,
						Name:      "Scala",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 10,
						Name:      "JavaScript",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 11,
						Name:      "SQL",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 12,
						Name:      "Solidity",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
					{Id: 13,
						Name:      "Perl",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{}},
				}))
			})
		})
	})

	Describe("Get Programming Language By Name", func() {
		When("Get Programming Language By Name successfully", func() {
			It("should return programming language", func() {
				proglang, err := frontendRepo.GetProgrammingLanguangeByName("Go")
				Expect(err).ToNot(HaveOccurred())

				Expect(proglang).To(Equal(domain.ProgrammingLanguangeDomain{
					Id:        1,
					Name:      "Go",
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				}))
			})
		})
	})

	Describe("Get Question By Programming Language", func() {
		When("Get Question By Programming Language successfully", func() {
			It("should return question", func() {
				question, err := frontendRepo.GetQuestionByProgrammingLanguange(int32(1), int32(1), int32(1))
				Expect(err).ToNot(HaveOccurred())

				Expect(question).To(Equal([]web.QuestionRequest{
					{ProgrammingLanguange: "Go",
						Question: "How well do you know about packages in Go programs"},
				}))
			})
		})
	})

	Describe("Save answer attempt", func() {
		When("Save answer attempt successfully", func() {
			It("should return true", func() {
				saveSuccess, err := frontendRepo.SaveAnswersAttempt(1, []domain.AnswersDomain{
					{Answer: "Really Understand",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					}, {Answer: "Understand",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					}, {Answer: "Do Not Understand",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					}})
				Expect(err).ToNot(HaveOccurred())

				Expect(saveSuccess).To(Equal(true))
			})
		})
	})

	Describe("Count answer attempt", func() {
		When("Count answer attempt successfully", func() {
			It("should return count", func() {
				_, err := frontendRepo.SaveAnswersAttempt(1, []domain.AnswersDomain{
					{Answer: "Really Understand",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					}, {Answer: "Understand",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					}, {Answer: "Do Not Understand",
						CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					}})
				Expect(err).ToNot(HaveOccurred())

				count, err := frontendRepo.CountAnswersAttempt(1)
				Expect(err).ToNot(HaveOccurred())

				Expect(count).To(Equal(int32(15)))
			})
		})
	})
})
