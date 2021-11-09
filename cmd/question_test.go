package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"quiz_master/builder"
	"quiz_master/domain"
	"quiz_master/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQuestion_Success(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestion := builder.NewQuestion(
		builder.SetNumber("1"),
		builder.SetQuestion("lorem ipsum dolor?"),
		builder.SetAnswer("2"),
	)
	mockQuestionUsecase.On("GetByNumber", mock.Anything).Return(*mockQuestion, nil).Once()

	cmd := NewQuestionCmd(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{mockQuestion.Number})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "Q : "+mockQuestion.Question+"\nA : "+mockQuestion.Answer+"\n")
}

func TestQuestion_Fail(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("GetByNumber", mock.Anything).Return(domain.Question{}, fmt.Errorf("some error")).Once()

	cmd := NewQuestionCmd(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"1"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "some error\n")
}

func TestAnswerQuestion_Success(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("AnswerQuestion", mock.Anything).Return(nil).Once()

	cmd := NewAnswerQuestionCmd(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"1", "1"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "Correct!\n")
}

func TestAnswerQuestion_Fail(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("AnswerQuestion", mock.Anything).Return(fmt.Errorf("Wrong Answer!")).Once()

	cmd := NewAnswerQuestionCmd(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"1", "1"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "Wrong Answer!\n")
}

func TestCreateQuestion_Success(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("Store", mock.Anything).Return(nil).Once()
	mockQuestion := builder.NewQuestion(
		builder.SetNumber("1"),
		builder.SetQuestion("lorem ipsum dolor?"),
		builder.SetAnswer("2"),
	)
	cmd := NewCreateQuestion(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{mockQuestion.Number, mockQuestion.Question, mockQuestion.Answer})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "Question no "+mockQuestion.Number+" created :\nQ : "+mockQuestion.Question+"\nA : "+mockQuestion.Answer+"\n")
}

func TestCreateQuestion_Fail(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("Store", mock.Anything).Return(fmt.Errorf("some error")).Once()
	mockQuestion := builder.NewQuestion(
		builder.SetNumber("1"),
		builder.SetQuestion("lorem ipsum dolor?"),
		builder.SetAnswer("2"),
	)
	cmd := NewCreateQuestion(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{mockQuestion.Number, mockQuestion.Question, mockQuestion.Answer})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "some error\n")
}

func TestDeleteQuestion_Success(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("Destroy", mock.Anything).Return(nil).Once()
	cmd := NewDeleteQuestionCmd(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"1"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "Question no 1 was deleted!\n")
}

func TestDeleteQuestion_Fail(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("Destroy", mock.Anything).Return(fmt.Errorf("some error")).Once()
	cmd := NewDeleteQuestionCmd(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"1"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "some error\n")
}

func TestListQuestion_Success(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("GetAll", mock.Anything).Return([]*domain.Question{}, nil).Once()
	cmd := NewListQuestion(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{})
	cmd.Execute()
	_, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListQuestion_Fail(t *testing.T) {
	mockQuestionUsecase := new(mocks.QuestionUsecase)
	mockQuestionUsecase.On("GetAll", mock.Anything).Return([]*domain.Question{}, fmt.Errorf("some error")).Once()
	cmd := NewListQuestion(mockQuestionUsecase)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, string(out), "some error\n")
}
