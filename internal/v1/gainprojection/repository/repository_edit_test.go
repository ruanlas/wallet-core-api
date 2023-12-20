package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestEditGainProjectionSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainPMock := NewGainProjectionBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddPayIn(now).
		AddIsPassive(true).
		AddIsAlreadyDone(false).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain_projection SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ?, is_already_done = ? 
		WHERE id = ?`).
		ExpectExec().
		WithArgs(
			gainPMock.PayIn,
			gainPMock.Description,
			gainPMock.Value,
			gainPMock.IsPassive,
			gainPMock.Category.Id,
			gainPMock.IsAlreadyDone,
			gainPMock.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	gainPEdited, err := _repository.Edit(context.Background(), *gainPMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", gainPEdited.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainProjectionBeginFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainPMock := NewGainProjectionBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddIsAlreadyDone(false).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainProjectionPrepareFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainPMock := NewGainProjectionBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddIsAlreadyDone(false).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain_projection SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ?, is_already_done = ? 
		WHERE id = ?`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainProjectionExecFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainPMock := NewGainProjectionBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddIsAlreadyDone(false).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain_projection SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ?, is_already_done = ? 
		WHERE id = ?`).
		ExpectExec().
		WithArgs(
			gainPMock.PayIn,
			gainPMock.Description,
			gainPMock.Value,
			gainPMock.IsPassive,
			gainPMock.Category.Id,
			gainPMock.IsAlreadyDone,
			gainPMock.Id).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainProjectionCommitFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainPMock := NewGainProjectionBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddIsAlreadyDone(false).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain_projection SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ?, is_already_done = ? 
		WHERE id = ?`).
		ExpectExec().
		WithArgs(
			gainPMock.PayIn,
			gainPMock.Description,
			gainPMock.Value,
			gainPMock.IsPassive,
			gainPMock.Category.Id,
			gainPMock.IsAlreadyDone,
			gainPMock.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
