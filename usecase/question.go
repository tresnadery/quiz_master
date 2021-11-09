package usecase

import (
	"errors"
	"fmt"
	"quiz_master/builder"
	"quiz_master/domain"
	"strconv"
	"strings"

	helper "quiz_master/helper"

	ntw "moul.io/number-to-words"
)

type questionUsecase struct {
	questionRepository domain.QuestionRepository
}

func NewQuestionUsecase(repo domain.QuestionRepository) domain.QuestionUsecase {
	return &questionUsecase{repo}
}

func (u *questionUsecase) Store(args []string) error {
	q := builder.NewQuestion(
		builder.SetNumber(args[0]),
		builder.SetQuestion(args[1]),
		builder.SetAnswer(args[2]),
	)

	if err := helper.Validate(q); err != nil {
		return err
	}

	existedQuestion, _ := u.questionRepository.GetByNumber(args[0])
	if existedQuestion != (domain.Question{}) {
		return fmt.Errorf("Question no " + args[0] + " already existed!")
	}

	return u.questionRepository.Store(q)
}

func (u *questionUsecase) GetAll() ([]*domain.Question, error) {
	return u.questionRepository.GetAll()
}

func (u *questionUsecase) GetByNumber(number string) (domain.Question, error) {
	return u.questionRepository.GetByNumber(number)
}

func (u *questionUsecase) AnswerQuestion(args []string) error {
	answer := args[1]
	question, err := u.questionRepository.GetByNumber(args[0])
	if err != nil {
		return err
	}
	convAnswer, _ := strconv.Atoi(question.Answer)
	if strings.ToLower(answer) != ntw.IntegerToEnUs(convAnswer) && question.Answer != answer {
		return errors.New("Wrong Answer!")
	}

	return nil
}

func (u *questionUsecase) Destroy(number string) error {
	_, err := u.questionRepository.GetByNumber(number)
	if err != nil {
		return err
	}
	return u.questionRepository.Destroy(number)
}
