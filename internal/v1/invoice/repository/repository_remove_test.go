package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRemoveInvoiceSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM invoice WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.NoError(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveInvoiceBeginFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin().WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveInvoicePrepareFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM invoice WHERE id = ? AND user_id = ?`).
		WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveInvoiceExecFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM invoice WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveInvoiceCommitFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM invoice WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().
		WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
