/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"quiz_master/database"
	"quiz_master/helper"
	"quiz_master/repository"
	"quiz_master/usecase"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

// questionCmd represents the question command
func NewQuestionCmd(u usecase.QuestionUsecase) *cobra.Command {
	return &cobra.Command{
		Use:   "question <number>",
		Short: "This command is use to show detail question",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			question, err := u.GetByNumber(args[0])
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), err.Error()+"\n")
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Q : "+question.Question+"\nA : "+question.Answer+"\n")
		},
	}
}

func NewAnswerQuestionCmd(u usecase.QuestionUsecase) *cobra.Command {
	return &cobra.Command{
		Use:   "answer_question <number> <answer>",
		Short: "This command to answer the question",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := u.AnswerQuestion(args); err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), err.Error()+"\n")
				return
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Correct!\n")
		},
	}
}

func NewCreateQuestion(u usecase.QuestionUsecase) *cobra.Command {
	return &cobra.Command{
		Use:   "create_question <number> <question> <answer>",
		Args:  cobra.ExactArgs(3),
		Short: "This command use to create question",
		Run: func(cmd *cobra.Command, args []string) {
			if err := u.Store(args); err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), err.Error()+"\n")
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Question no "+args[0]+" created :\nQ : "+args[1]+"\nA : "+args[2]+"\n")
		},
	}
}

// createQuestionCmd represents the createQuestion command
func NewDeleteQuestionCmd(u usecase.QuestionUsecase) *cobra.Command {
	return &cobra.Command{
		Use:   "delete_question <number>",
		Short: "This command is use to delete question",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := u.Destroy(args[0]); err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), err.Error()+"\n")
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Question no "+args[0]+" was deleted!\n")
		},
	}
}

func NewListQuestion(u usecase.QuestionUsecase) *cobra.Command {
	return &cobra.Command{
		Use:   "list_question",
		Short: "This command is use to list question",
		Run: func(cmd *cobra.Command, args []string) {
			questions, err := u.GetAll()
			if err != nil {
				fmt.Fprintf(cmd.OutOrStdout(), err.Error()+"\n")
			}

			helper.ListQuestionResponse(questions)
		},
	}
}

func InitCmd() {

	db := database.InitDB()
	repository := repository.NewQuestionRepository(db)
	ucase := usecase.NewQuestionUsecase(repository)
	rootCmd.AddCommand(NewQuestionCmd(ucase))
	rootCmd.AddCommand(NewAnswerQuestionCmd(ucase))
	rootCmd.AddCommand(NewCreateQuestion(ucase))
	rootCmd.AddCommand(NewDeleteQuestionCmd(ucase))
	rootCmd.AddCommand(NewListQuestion(ucase))
}
