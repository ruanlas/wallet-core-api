package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSaveGainProjectionSuccess(t *testing.T) {
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
		INSERT INTO gain_projection (id, created_at, pay_in, description, value, is_passive, is_already_done, user_id, category_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			gainPMock.Id,
			gainPMock.CreatedAt.Unix(),
			gainPMock.PayIn,
			gainPMock.Description,
			gainPMock.Value,
			gainPMock.IsPassive,
			gainPMock.IsAlreadyDone,
			gainPMock.UserId,
			gainPMock.Category.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	gainPSaved, err := _repository.Save(context.Background(), *gainPMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", gainPSaved.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainProjectionBeginFail(t *testing.T) {
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

	_, err = _repository.Save(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainProjectionPrepareFail(t *testing.T) {
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
		INSERT INTO gain_projection (id, created_at, pay_in, description, value, is_passive, is_already_done, user_id, category_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainProjectionExecFail(t *testing.T) {
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
		INSERT INTO gain_projection (id, created_at, pay_in, description, value, is_passive, is_already_done, user_id, category_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			gainPMock.Id,
			gainPMock.CreatedAt.Unix(),
			gainPMock.PayIn,
			gainPMock.Description,
			gainPMock.Value,
			gainPMock.IsPassive,
			gainPMock.IsAlreadyDone,
			gainPMock.UserId,
			gainPMock.Category.Id).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainProjectionCommitFail(t *testing.T) {
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
		INSERT INTO gain_projection (id, created_at, pay_in, description, value, is_passive, is_already_done, user_id, category_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			gainPMock.Id,
			gainPMock.CreatedAt.Unix(),
			gainPMock.PayIn,
			gainPMock.Description,
			gainPMock.Value,
			gainPMock.IsPassive,
			gainPMock.IsAlreadyDone,
			gainPMock.UserId,
			gainPMock.Category.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *gainPMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
