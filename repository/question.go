package repository

import (
	"database/sql"
	"fmt"
	"quiz_master/domain"

	_ "github.com/go-sql-driver/mysql"
)

type QuestionRepository interface {
	GetAll() ([]*domain.Question, error)
	Store(question *domain.Question) error
	GetByNumber(number string) (*domain.Question, error)
	Destroy(number string) error
}

type questionRepository struct {
	conn *sql.DB
}

func NewQuestionRepository(conn *sql.DB) QuestionRepository {
	return &questionRepository{conn}
}

func (r questionRepository) GetAll() ([]*domain.Question, error) {
	questions := []*domain.Question{}
	rows, err := r.conn.Query("SELECT number,question,answer FROM questions WHERE deleted_at IS NULL ORDER BY number ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		question := &domain.Question{}
		err := rows.Scan(&question.Number, &question.Question, &question.Answer)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}

func (r *questionRepository) Store(question *domain.Question) error {

	stmt, err := r.conn.Prepare("INSERT INTO questions(number, question, answer) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(question.Number, question.Question, question.Answer)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}

	return nil
}

func (r *questionRepository) GetByNumber(number string) (*domain.Question, error) {
	q := domain.Question{}
	stmt, err := r.conn.Prepare("SELECT id,number,question,answer FROM questions WHERE number = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(number).Scan(&q.ID, &q.Number, &q.Question, &q.Answer)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Question not found")
	}

	return &q, nil
}

func (r *questionRepository) Destroy(number string) error {
	stmt, err := r.conn.Prepare("DELETE FROM questions WHERE number = ? AND deleted_at IS NULL")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(number)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}

	return nil
}
