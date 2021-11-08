package mocks

import (
	"quiz_master/domain"

	mock "github.com/stretchr/testify/mock"
)

type QuestionRepository struct {
	mock.Mock
}

func (m *QuestionRepository) GetAll() ([]*domain.Question, error) {
	ret := m.Called()

	var r0 []*domain.Question
	if rf, ok := ret.Get(0).(func() []*domain.Question); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Question)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *QuestionRepository) GetByNumber(number string) (*domain.Question, error) {
	ret := m.Called(number)

	var r0 *domain.Question
	if rf, ok := ret.Get(0).(func(string) *domain.Question); ok {
		r0 = rf(number)
	} else {
		r0 = ret.Get(0).(*domain.Question)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *QuestionRepository) Store(q *domain.Question) error {
	ret := m.Called(q)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Question) error); ok {
		r0 = rf(q)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *QuestionRepository) Destroy(number string) error {
	ret := m.Called(number)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(number)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
