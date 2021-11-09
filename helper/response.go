package helper

import (
	"fmt"
	"quiz_master/domain"
)

func ListQuestionResponse(questions []*domain.Question) {
	fmt.Println("No |\tQuestion\t|\tAnswer")
	for _, q := range questions {
		fmt.Printf("%s\t%s\t\t\t%s\n", q.Number, q.Question, q.Answer)
	}
	fmt.Printf("\n")
}
