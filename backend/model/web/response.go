package web

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type LoginResponse struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type ProgrammingLanguageResponse struct {
	Id        int32  `json:"id"`
	Name      string `json:"name"`
	UrlImages string `json:"url_images"`
}

type QuestionCreateResponse struct {
	Id                   int32            `json:"id"`
	ProgrammingLanguange string           `json:"programming_languange"`
	Question             string           `json:"question"`
	Answer               []AnswerResponse `json:"answer"`
}

type QuestionResponse struct {
	Id                   int32  `json:"id"`
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
	LevelId  int32  `json:"level_id"`
	Username string `json:"username"`
	Level    string `json:"level"`
	Score    int32  `json:"score"`
}

type RecommendationResponse struct {
	ImageUrl string `json:"image_url"`
}
