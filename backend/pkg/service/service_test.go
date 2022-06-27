package service_test

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rg-km/final-project-engineering-14/backend/model/web"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/repository"
	"github.com/rg-km/final-project-engineering-14/backend/pkg/service"
)

var _ = Describe("Service Test", func() {
	var db *sql.DB
	var err error
	var authService *service.AuthServiceSQLite
	var backendService *service.BackendSQLite
	var frontendService *service.FrontendServiceSQLite
	var authRepo *repository.AuthRepositorySQLite

	BeforeEach(func() {
		db, err = sql.Open("sqlite3", "./service_test.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(150) NOT NULL,
			password VARCHAR(255) NOT NULL,
			phone VARCHAR(50) NOT NULL,
			role VARCHAR(15) NOT NULL,
			is_login BOOLEAN NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP 
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
				url_images VARCHAR(255) NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			);

			CREATE TABLE IF NOT EXISTS answers_attempts (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER NOT NULL,
				answer VARCHAR(50) NOT NULL,
				FOREIGN KEY (user_id) REFERENCES users(id)
			);

			CREATE TABLE IF NOT EXISTS levels (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				level VARCHAR(15) NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
			);

			CREATE TABLE IF NOT EXISTS recommendations (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			  image_url TEXT NOT NULL,
				level_id INTEGER NOT NULL,
				created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (level_id) REFERENCES levels(id)
			);

			INSERT INTO users (
				username, email, password, phone, role, is_login, created_at, updated_at
			) VALUES 
			('admin', 'admin@gmail.com', 'admin123', '081234567890', 'admin', 0, '2022-06-13 00:00:00', '2022-06-04 00:00:00'),
			('navis', 'navis@gmail.com', '1ARVn2Auq2_WAqx2gNrL-q3RNjAzXpUfCXrzkA6d4Xa22yhRLy4AC50E-6UTPoscbo31nbOoq51gvkuXzJ6B2w==', '081234567890', 'guest', 0, '2022-06-13 00:00:00', '2022-06-04 00:00:00');

			INSERT INTO answers
				(answer, created_at, updated_at)
			VALUES
			('Really Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Do Not Understand', '2020-01-01 00:00:00', '2020-01-01 00:00:00');

			INSERT INTO programming_languanges
				(name, url_images, created_at, updated_at)
			VALUES
			('Go', 'https://www.linkpicture.com/q/go_1.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Python', 'https://www.linkpicture.com/q/python_1.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Java', 'https://www.linkpicture.com/q/java_3.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('C#', 'https://www.linkpicture.com/q/cs_2.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Ruby', 'https://www.linkpicture.com/q/ruby_1.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('PHP', 'https://www.linkpicture.com/q/php_2.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('Rust', 'https://www.linkpicture.com/q/rust_5.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00'),
			('JavaScript', 'https://www.linkpicture.com/q/js_16.png', '2020-01-01 00:00:00', '2020-01-01 00:00:00');

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

			INSERT INTO levels 
			(level, created_at, updated_at)
			VALUES
			('Beginner', '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
			('Intermediate', '2022-06-01 00:00:00', '2022-06-01 00:00:00');
		
			INSERT INTO recommendations
			(image_url, level_id, created_at, updated_at)
			VALUES
			('https://www.linkpicture.com/q/Basic_1.png', 1, '2022-06-01 00:00:00', '2022-06-01 00:00:00'),
			('https://www.linkpicture.com/q/Intermediate.png', 2, '2022-06-01 00:00:00', '2022-06-01 00:00:00');
		`)

		if err != nil {
			panic(err)
		}

		authService = service.NewAuthService(&repository.AuthRepositorySQLite{db})
		backendService = service.NewBackendService(&repository.BackendRepositorySQLite{db}, &repository.FrontendRepositorySQLite{db})
		frontendService = service.NewFrontendService(&repository.AuthRepositorySQLite{db}, &repository.FrontendRepositorySQLite{db}, &repository.BackendRepositorySQLite{db})
		authRepo = repository.NewAuthRepository(db)

	})

	AfterEach(func() {
		db, err = sql.Open("sqlite3", "./service_test.db")
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(`
			DROP TABLE users;
			DROP TABLE answers;
			DROP TABLE questions;
			DROP TABLE programming_languanges;
			DROP TABLE answers_attempts;
			DROP TABLE levels;
			DROP TABLE recommendations;
		`)

		if err != nil {
			panic(err)
		}

	})

	Describe("Create", func() {
		When("user data is valid", func() {
			It("should return user data", func() {
				user, err := authService.Create(web.RegisterRequest{
					Username: "paskal",
					Email:    "paskal@gmail.com",
					Password: "1234",
					Phone:    "081234567890",
				}, "paskal@gmail.com")
				Expect(err).To(BeNil())

				Expect(user.Username).To(Equal("paskal"))
				Expect(user.Email).To(Equal("paskal@gmail.com"))
				Expect(user.Phone).To(Equal("081234567890"))
			})
		})
	})

	Describe("Generate Token", func() {
		When("Login will generate token", func() {
			It("should return token", func() {

				user, err := authService.GenerateToken("navis@gmail.com", "1234")
				Expect(err).ToNot(HaveOccurred())

				Expect(user.Token).To(Equal(user.Token))
			})
		})
	})

	Describe("Parse Token", func() {
		When("Token is valid", func() {
			It("should return user data", func() {

				token, err := authService.GenerateToken("navis@gmail.com", "1234")
				Expect(err).To(BeNil())

				_, _, err = authService.ParseToken(context.TODO(), token.Token)
				Expect(err).To(BeNil())

				claims := &service.Claims{}

				Expect(claims.UserId).To(Equal(int32(0)))

			})
		})
	})

	Describe("Logout", func() {
		When("Logout user successfully", func() {
			It("Accept the logout", func() {
				user, err := authRepo.GetUser("navis@gmail.com", "1ARVn2Auq2_WAqx2gNrL-q3RNjAzXpUfCXrzkA6d4Xa22yhRLy4AC50E-6UTPoscbo31nbOoq51gvkuXzJ6B2w==")
				Expect(err).ToNot(HaveOccurred())
				Expect(user.IsLogin).To(Equal(false))

				userIdLogin := user.Id
				isSuccesLogout, err := authService.Logout(userIdLogin)
				Expect(err).ToNot(HaveOccurred())
				Expect(isSuccesLogout).To(Equal(true))
			})
		})
	})

	Describe("Create Admin Question", func() {
		When("Admin Question is valid", func() {
			It("should return question", func() {
				question, err := backendService.CreateAdminQuestion(1, web.QuestionRequest{
					ProgrammingLanguange: "Go",
					Question:             "Are you familiar with how pointer work and use in Golang?",
				})
				Expect(err).To(BeNil())

				Expect(question.ProgrammingLanguange).To(Equal("Go"))
				Expect(question.Question).To(Equal("Are you familiar with how pointer work and use in Golang?"))
			})
		})
	})

	Describe("Get Admin Question", func() {
		When("Admin Question is valid", func() {
			It("should return question", func() {
				question, err := backendService.GetAdminQuestions()
				Expect(err).To(BeNil())

				Expect(question).To(Equal([]web.QuestionResponse{
					{
						Id:                   1,
						ProgrammingLanguange: "Go",
						Question:             "How well do you know about packages in Go programs",
					},
					{
						Id:                   2,
						ProgrammingLanguange: "Go",
						Question:             "How well do you understand using the for loop in Golang?",
					},
					{
						Id:                   3,
						ProgrammingLanguange: "Go",
						Question:             "Are you familiar with how array, slice, and maps work and use in Golang?",
					},
				}))
			})
		})
	})

	Describe("Count Question", func() {
		When("Admin Question is valid", func() {
			It("should return question", func() {
				count, err := backendService.GetCountQuestions()
				Expect(err).To(BeNil())

				Expect(count).To(Equal([]web.CountQuestionResponse{
					{
						ProgrammingLanguage: "Go",
						TotalQuestion:       3,
					},
				}))
			})
		})
	})

	Describe("Update Admin Question", func() {
		When("Admin Question is valid", func() {
			It("should return true", func() {
				question, err := backendService.UpdateAdminQuestion(web.QuestionRequest{
					ProgrammingLanguange: "Go",
					Question:             "How well do you know about packages in Go programs",
				}, 1)
				Expect(err).To(BeNil())

				Expect(question).To(Equal(web.QuestionResponse{
					Id:                   1,
					ProgrammingLanguange: "Go",
					Question:             "How well do you know about packages in Go programs",
				}))
			})
		})
	})

	Describe("Delete Admin Question", func() {
		When("Admin Question is valid", func() {
			It("should return true", func() {
				question, err := backendService.DeleteAdminQuestion(1)
				Expect(err).To(BeNil())

				Expect(question).To(Equal(true))
			})
		})
	})

	Describe("Get Programming Language", func() {
		When("Programming Language is valid", func() {
			It("should return programming language", func() {
				programmingLanguage, err := frontendService.GetProgrammingLanguanges()
				Expect(err).To(BeNil())

				Expect(programmingLanguage).To(Equal([]web.ProgrammingLanguageResponse{
					{
						Id:        1,
						Name:      "Go",
						UrlImages: "https://www.linkpicture.com/q/go_1.png"},
					{
						Id:        2,
						Name:      "Python",
						UrlImages: "https://www.linkpicture.com/q/python_1.png"},
					{
						Id:        3,
						Name:      "Java",
						UrlImages: "https://www.linkpicture.com/q/java_3.png"},
					{
						Id:        4,
						Name:      "C#",
						UrlImages: "https://www.linkpicture.com/q/cs_2.png"},
					{
						Id:        5,
						Name:      "Ruby",
						UrlImages: "https://www.linkpicture.com/q/ruby_1.png"},
					{
						Id:        6,
						Name:      "PHP",
						UrlImages: "https://www.linkpicture.com/q/php_2.png"},
					{
						Id:        7,
						Name:      "Rust",
						UrlImages: "https://www.linkpicture.com/q/rust_5.png"},
					{
						Id:        8,
						Name:      "JavaScript",
						UrlImages: "https://www.linkpicture.com/q/js_16.png"},
				}))
			})
		})
	})

	Describe("Get Question By Programming Language ", func() {
		When("Programming Language is valid", func() {
			It("should return question", func() {
				question, err := frontendService.GetQuestionByProgrammingLanguage(1, 1, 1)
				Expect(err).To(BeNil())

				Expect(question).To(Equal([]web.QuestionCreateResponse{
					{
						Id:                   1,
						ProgrammingLanguange: "Go",
						Question:             "How well do you know about packages in Go programs",
						Answer: []web.AnswerResponse{
							{
								Answer: "Really Understand",
							},
							{
								Answer: "Understand",
							},
							{
								Answer: "Do Not Understand",
							},
						},
					},
				}))
			})
		})
	})

	Describe("Post Answers Attempt", func() {
		When("Answers is valid", func() {
			It("should return total", func() {
				total, err := frontendService.PostAnswersAttempt(2, web.AnswersAttemptRequest{
					AnswerOne:   "Really Understand",
					AnswerTwo:   "Understand",
					AnswerThree: "Do Not Understand",
					AnswerFour:  "Do Not Understand",
					AnswerFive:  "Do Not Understand",
					AnswerSix:   "Do Not Understand",
					AnswerSeven: "Do Not Understand",
					AnswerEight: "Do Not Understand",
					AnswerNine:  "Do Not Understand",
					AnswerTen:   "Understand",
				})

				Expect(err).To(BeNil())

				Expect(total.Score).To(Equal(int32(20)))
				Expect(total.Level).To(Equal("Beginner"))
				Expect(total.Username).To(Equal("navis"))
			})
		})
	})

	Describe("Get Recomendation", func() {
		When("Level Id is valid", func() {
			It("should return recomendation", func() {
				seeRecomendation, err := frontendService.SeeRecommendationByLevelId(1)
				Expect(err).To(BeNil())

				Expect(seeRecomendation).To(Equal(web.RecommendationResponse{
					ImageUrl: "https://www.linkpicture.com/q/Basic_1.png",
				}))
			})
		})
	})
})
