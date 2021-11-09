package domain

type QuestionRepository interface {
	GetAll() ([]*Question, error)
	Store(question *Question) error
	GetByNumber(number string) (Question, error)
	Destroy(number string) error
}

type QuestionUsecase interface {
	Store(args []string) error
	GetAll() ([]*Question, error)
	GetByNumber(number string) (Question, error)
	AnswerQuestion(args []string) error
	Destroy(number string) error
}

type Question struct {
	ID       int    `json:"id"`
	Number   string `json:"number" validate:"required,numeric"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required,numeric"`
}
