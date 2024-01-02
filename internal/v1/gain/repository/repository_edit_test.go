package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestEditGainSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainMock := NewGainBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddPayIn(now).
		AddIsPassive(true).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			gainMock.PayIn,
			gainMock.Description,
			gainMock.Value,
			gainMock.IsPassive,
			gainMock.Category.Id,
			gainMock.Id,
			gainMock.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	gainEdited, err := _repository.Edit(context.Background(), *gainMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", gainEdited.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainBeginFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainMock := NewGainBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainPrepareFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainMock := NewGainBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ? 
		WHERE id = ? AND user_id = ?`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainExecFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainMock := NewGainBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			gainMock.PayIn,
			gainMock.Description,
			gainMock.Value,
			gainMock.IsPassive,
			gainMock.Category.Id,
			gainMock.Id,
			gainMock.UserId).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditGainCommitFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	gainMock := NewGainBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddCategory(GainCategory{Id: 1}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE gain SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			gainMock.PayIn,
			gainMock.Description,
			gainMock.Value,
			gainMock.IsPassive,
			gainMock.Category.Id,
			gainMock.Id,
			gainMock.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
