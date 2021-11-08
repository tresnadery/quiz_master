package repository

import (
	"database/sql"
	"fmt"
	"log"
	"quiz_master/domain"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var q = &domain.Question{
	ID:       1,
	Number:   "1",
	Question: "lorem ipsum dolor sit amet?",
	Answer:   "1",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetAll_Success(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := "SELECT number,question,answer FROM questions WHERE deleted_at IS NULL ORDER BY number ASC"

	rows := sqlmock.NewRows([]string{"number", "question", "answer"}).
		AddRow(q.Number, q.Question, q.Answer)
	mock.ExpectQuery(query).WillReturnRows(rows)

	questions, err := questionRepo.GetAll()
	assert.NotEmpty(t, questions)
	assert.NoError(t, err)
	assert.Len(t, questions, 1)
}

func TestGetAll_FailBecauseErrorQuery(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := "SELECT number,question,answer FROM questions WHERE deleted_at IS NULL ORDER BY number ASC"

	mock.ExpectQuery(query).WillReturnError(fmt.Errorf("some error"))

	questions, err := questionRepo.GetAll()
	assert.Empty(t, questions)
	assert.Error(t, err)
}

func TestGetAll_FailBecauseDataTypeIsWrong(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := "SELECT number,question,answer FROM questions WHERE deleted_at IS NULL ORDER BY number ASC"

	rows := sqlmock.NewRows([]string{"number", "question", "answer"}).
		AddRow(q.Number, q.Question, nil)
	mock.ExpectQuery(query).WillReturnRows(rows)

	questions, err := questionRepo.GetAll()
	assert.Empty(t, questions)
	assert.Error(t, err)
}

func TestStore_Success(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("INSERT INTO questions(number, question, answer) VALUES(?, ?, ?)")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(q.Number, q.Question, q.Answer).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := questionRepo.Store(q)
	assert.NoError(t, err)
}

func TestStore_FailQueryNotMatch(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("INSERT INTO questions(id,number, question, answer) VALUES(?, ?, ?)")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(q.Number, q.Question, q.Answer).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := questionRepo.Store(q)
	assert.Error(t, err)
}

func TestStore_FailErrorOnQuery(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("INSERT INTO questions(number, question, answer) VALUES(?, ?, ?)")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(q.Number, q.Question, q.Answer).
		WillReturnError(fmt.Errorf("some error"))

	err := questionRepo.Store(q)
	assert.Error(t, err)
}

func TestStore_FailNoRowAffected(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("INSERT INTO questions(number, question, answer) VALUES(?, ?, ?)")
	prep := mock.ExpectPrepare(query)

	prep.ExpectExec().
		WithArgs(q.Number, q.Question, q.Answer).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := questionRepo.Store(q)
	assert.Error(t, err)
}

func TestGetByNumber_Success(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("SELECT id,number,question,answer FROM questions WHERE number = ?")
	prep := mock.ExpectPrepare(query)

	rows := sqlmock.NewRows([]string{"id", "number", "question", "answer"}).
		AddRow(q.ID, q.Number, q.Question, q.Answer)
	prep.ExpectQuery().WithArgs(q.Number).WillReturnRows(rows)

	question, err := questionRepo.GetByNumber(q.Number)
	assert.NotEmpty(t, question)
	assert.NoError(t, err)
}

func TestGetByNumber_FailQueryNotMatch(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("SELECT id,number,question FROM questions WHERE number = ?")
	prep := mock.ExpectPrepare(query)

	rows := sqlmock.NewRows([]string{"id", "number", "question", "answer"}).
		AddRow(q.ID, q.Number, q.Question, q.Answer)
	prep.ExpectQuery().WithArgs(q.Number).WillReturnRows(rows)

	question, err := questionRepo.GetByNumber(q.Number)
	assert.Empty(t, question)
	assert.Error(t, err)
}

func TestGetByNumber_FailErrorQuery(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("SELECT id,number,question,answer FROM questions WHERE number = ?")
	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WithArgs(q.Number).WillReturnError(fmt.Errorf("some error"))

	question, err := questionRepo.GetByNumber(q.Number)
	assert.Empty(t, question)
	assert.Error(t, err)
}

func TestGetByNumber_FailRecordNotFound(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("SELECT id,number,question,answer FROM questions WHERE number = ?")
	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WithArgs(q.Number).WillReturnError(sql.ErrNoRows)

	question, err := questionRepo.GetByNumber(q.Number)
	assert.Empty(t, question)
	assert.Error(t, err)
}

func TestDestroy_Success(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("DELETE FROM questions WHERE number = ? AND deleted_at IS NULL")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(q.Number).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := questionRepo.Destroy(q.Number)
	assert.NoError(t, err)
}

func TestDestroy_FailQueryNotMatch(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("DELETE FROM questions WHERE id = ? AND deleted_at IS NULL")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(q.Number).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := questionRepo.Destroy(q.Number)
	assert.Error(t, err)
}

func TestDestroy_FailErrorQuery(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("DELETE FROM questions WHERE number = ? AND deleted_at IS NULL")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(q.Number).
		WillReturnError(fmt.Errorf("some error"))

	err := questionRepo.Destroy(q.Number)
	assert.Error(t, err)
}

func TestDestroy_FailNoRowAffected(t *testing.T) {
	db, mock := NewMock()
	questionRepo := NewQuestionRepository(db)

	query := regexp.QuoteMeta("DELETE FROM questions WHERE number = ? AND deleted_at IS NULL")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().
		WithArgs(q.Number).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := questionRepo.Destroy(q.Number)
	assert.Error(t, err)
}
