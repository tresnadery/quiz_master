package builder

import (
	"qtest/domain"
)

type Option func(*domain.Question)

func NewQuestion(options ...Option) *domain.Question {
	q := &domain.Question{}
	for _, o := range options {
		o(q)
	}
	return q
}

func SetNumber(number string) Option {
	return func(q *domain.Question) {
		q.Number = number
	}
}

func SetQuestion(question string) Option {
	return func(q *domain.Question) {
		q.Question = question
	}
}

func SetAnswer(answer string) Option {
	return func(q *domain.Question) {
		q.Answer = answer
	}
}
