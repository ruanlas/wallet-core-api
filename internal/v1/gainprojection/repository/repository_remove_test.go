package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRemoveGainProjectionSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM gain_projection WHERE id = ?`).
		ExpectExec().
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8").
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.NoError(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveGainProjectionBeginFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin().WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveGainProjectionPrepareFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM gain_projection WHERE id = ?`).
		WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveGainProjectionExecFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM gain_projection WHERE id = ?`).
		ExpectExec().
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8").
		WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRemoveGainProjectionCommitFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`DELETE FROM gain_projection WHERE id = ?`).
		ExpectExec().
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8").
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().
		WillReturnError(errors.New("An error has been ocurred"))

	err = _repository.Remove(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
