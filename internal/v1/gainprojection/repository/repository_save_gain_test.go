package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSaveGainSuccess(t *testing.T) {
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
		AddGainProjectionId("7a494375-53a1-41e4-a9db-6bb30eaf23c2").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		INSERT INTO gain_done (id, created_at, pay_in, description, value, is_passive, user_id, category_id, gain_projection_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			gainMock.Id,
			gainMock.CreatedAt.Unix(),
			gainMock.PayIn,
			gainMock.Description,
			gainMock.Value,
			gainMock.IsPassive,
			gainMock.UserId,
			gainMock.Category.Id,
			gainMock.GainProjectionId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	gainSaved, err := _repository.SaveGain(context.Background(), *gainMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", gainSaved.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainBeginFail(t *testing.T) {
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
		AddGainProjectionId("7a494375-53a1-41e4-a9db-6bb30eaf23c2").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.SaveGain(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainPrepareFail(t *testing.T) {
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
		AddGainProjectionId("7a494375-53a1-41e4-a9db-6bb30eaf23c2").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		INSERT INTO gain_done (id, created_at, pay_in, description, value, is_passive, user_id, category_id, gain_projection_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.SaveGain(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainExecFail(t *testing.T) {
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
		AddGainProjectionId("7a494375-53a1-41e4-a9db-6bb30eaf23c2").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		INSERT INTO gain_done (id, created_at, pay_in, description, value, is_passive, user_id, category_id, gain_projection_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			gainMock.Id,
			gainMock.CreatedAt.Unix(),
			gainMock.PayIn,
			gainMock.Description,
			gainMock.Value,
			gainMock.IsPassive,
			gainMock.UserId,
			gainMock.Category.Id,
			gainMock.GainProjectionId).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.SaveGain(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveGainCommitFail(t *testing.T) {
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
		AddGainProjectionId("7a494375-53a1-41e4-a9db-6bb30eaf23c2").
		Build()
	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		INSERT INTO gain_done (id, created_at, pay_in, description, value, is_passive, user_id, category_id, gain_projection_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			gainMock.Id,
			gainMock.CreatedAt.Unix(),
			gainMock.PayIn,
			gainMock.Description,
			gainMock.Value,
			gainMock.IsPassive,
			gainMock.UserId,
			gainMock.Category.Id,
			gainMock.GainProjectionId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.SaveGain(context.Background(), *gainMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
