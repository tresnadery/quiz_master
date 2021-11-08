package domain

type Question struct {
	ID       int    `json:"id"`
	Number   string `json:"number" validate:"required,numeric"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required,numeric"`
}
