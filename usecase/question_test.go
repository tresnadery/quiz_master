package usecase

import (
	"database/sql"
	"fmt"
	"quiz_master/builder"
	"quiz_master/domain"
	"quiz_master/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	mockQuestion := builder.NewQuestion(
		builder.SetNumber("1"),
		builder.SetQuestion("lorem ipsum dolor?"),
		builder.SetAnswer("2"),
	)

	mockListQuestion := make([]*domain.Question, 0)
	mockListQuestion = append(mockListQuestion, mockQuestion)
	t.Run("success", func(t *testing.T) {
		mockQuestionRepo.On("GetAll", mock.Anything).Return(mockListQuestion, nil).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		questions, err := u.GetAll()
		assert.NoError(t, err)
		assert.Len(t, questions, len(mockListQuestion))

		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestStore_FailNumberValidation(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("error-failed", func(t *testing.T) {
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.Store([]string{"abc", "lorem ipsum", "1"})
		assert.Error(t, err)
		assert.Equal(t, err, fmt.Errorf("Number must be a valid numeric value"))

		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestStore_FailAnswerValidation(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("error-failed", func(t *testing.T) {
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.Store([]string{"100", "lorem ipsum", "ac"})
		assert.Error(t, err)
		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestStore_FailQuestionAlreadyExisted(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("error-failed", func(t *testing.T) {
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(domain.Question{ID: 1}, fmt.Errorf("Question no 1 already existed!")).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.Store([]string{"1", "lorem ipsum", "1"})
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "Question no 1 already existed!")
		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestStore_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(domain.Question{}, nil).Once()
		mockQuestionRepo.On("Store", mock.AnythingOfType("*domain.Question")).Return(nil).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.Store([]string{"1", "lorem ipsum", "1"})
		assert.NoError(t, err)
		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestGetByNumber_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestion := builder.NewQuestion(
			builder.SetNumber("1"),
			builder.SetQuestion("lorem ipsum dolor?"),
			builder.SetAnswer("2"),
		)
		mockQuestion.ID = 1
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(*mockQuestion, nil).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		question, err := u.GetByNumber("1")
		assert.NoError(t, err)
		assert.Equal(t, question, *mockQuestion)
		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestAnswerQuestion_FailQuestionNotFound(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(domain.Question{}, sql.ErrNoRows).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.AnswerQuestion([]string{"1", "1"})
		assert.Error(t, err)

		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestAnswerQuestion_FailAnswerIsWrong(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestion := builder.NewQuestion(
			builder.SetNumber("1"),
			builder.SetQuestion("lorem ipsum dolor?"),
			builder.SetAnswer("2"),
		)
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(*mockQuestion, nil).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.AnswerQuestion([]string{"1", "3"})
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "Wrong Answer!")

		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestAnswerQuestion_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestion := builder.NewQuestion(
			builder.SetNumber("1"),
			builder.SetQuestion("lorem ipsum dolor?"),
			builder.SetAnswer("2"),
		)
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(*mockQuestion, nil).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.AnswerQuestion([]string{"1", "2"})
		assert.NoError(t, err)

		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestAnswerQuestion_SuccessWithWordAnswer(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestion := builder.NewQuestion(
			builder.SetNumber("1"),
			builder.SetQuestion("lorem ipsum dolor?"),
			builder.SetAnswer("2"),
		)
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(*mockQuestion, nil).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.AnswerQuestion([]string{"1", "Two"})
		assert.NoError(t, err)

		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestDestroyQuestion_FailQuestionNotFound(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(domain.Question{}, fmt.Errorf("Question not found")).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.Destroy("1")
		assert.Error(t, err)

		mockQuestionRepo.AssertExpectations(t)
	})
}

func TestDestroyQuestion_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.QuestionRepository)
	t.Run("success", func(t *testing.T) {
		mockQuestion := builder.NewQuestion(
			builder.SetNumber("1"),
			builder.SetQuestion("lorem ipsum dolor?"),
			builder.SetAnswer("2"),
		)
		mockQuestion.ID = 1
		mockQuestionRepo.On("GetByNumber", mock.Anything).Return(*mockQuestion, nil).Once()
		mockQuestionRepo.On("Destroy", mock.Anything).Return(nil).Once()
		u := NewQuestionUsecase(mockQuestionRepo)
		err := u.Destroy("1")
		assert.NoError(t, err)

		mockQuestionRepo.AssertExpectations(t)
	})
}
