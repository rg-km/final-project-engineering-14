package web

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ProgrammingLanguageResponse struct {
	Name string `json:"name"`
}

type QuestionCreateResponse struct {
	ProgrammingLanguange string           `json:"programming_languange"`
	Question             string           `json:"question"`
	Answer               []AnswerResponse `json:"answer"`
}

type QuestionResponse struct {
	Question             string `json:"question"`
	ProgrammingLanguange string `json:"programming_languange"`
}

type AnswerResponse struct {
	Answer string `json:"answer"`
}

type CountQuestionResponse struct {
	ProgrammingLanguage string `json:"programming_languange"`
	TotalQuestion       uint32 `json:"total_question"`
}

type TotalAnswerAttemptResponse struct {
	Username string `json:"username"`
	Level    string `json:"level"`
	Score    int32  `json:"total"`
}
