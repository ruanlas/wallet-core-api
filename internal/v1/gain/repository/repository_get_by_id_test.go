package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetGainByIdSuccess(t *testing.T) {
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
		AddCategory(GainCategory{Id: 1, Category: "Salário"}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		AddGainProjectionId("7172a75e-f41e-47df-a514-12580f34bd09").
		Build()

	rowsGainMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"description",
		"value",
		"is_passive",
		"user_id",
		"category_id",
		"category",
		"gain_projection_id",
	}).AddRow(
		gainMock.Id,
		gainMock.CreatedAt.Unix(),
		gainMock.PayIn,
		gainMock.Description,
		gainMock.Value,
		gainMock.IsPassive,
		gainMock.UserId,
		gainMock.Category.Id,
		gainMock.Category.Category,
		gainMock.GainProjectionId,
	)

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			g.id,
			g.created_at,
			g.pay_in,
			g.description,
			g.value,
			g.is_passive,
			g.user_id,
			gc.id,
			gc.category,
			g.gain_projection_id
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE g.id = ? AND g.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnRows(rowsGainMock)

	gainPSaved, err := _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", gainPSaved.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetGainByIdQueryFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			g.id,
			g.created_at,
			g.pay_in,
			g.description,
			g.value,
			g.is_passive,
			g.user_id,
			gc.id,
			gc.category,
			g.gain_projection_id
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE g.id = ? AND g.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetGainByIdRowEmpty(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	rowsGainMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"description",
		"value",
		"is_passive",
		"user_id",
		"category_id",
		"category",
		"gain_projection_id",
	})

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			g.id,
			g.created_at,
			g.pay_in,
			g.description,
			g.value,
			g.is_passive,
			g.user_id,
			gc.id,
			gc.category,
			g.gain_projection_id
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE g.id = ? AND g.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnRows(rowsGainMock)

	gainPSaved, err := _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.NoError(t, err)
	assert.Empty(t, gainPSaved)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetGainByIdRowScanFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	rowsGainMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"description",
		"value",
		"is_passive",
		"user_id",
		"category_id",
		"category",
		"gain_projection_id",
	}).AddRow(
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil).
		RowError(1, errors.New("An error has been ocurred"))

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			g.id,
			g.created_at,
			g.pay_in,
			g.description,
			g.value,
			g.is_passive,
			g.user_id,
			gc.id,
			gc.category,
			g.gain_projection_id
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE g.id = ? AND g.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnRows(rowsGainMock)

	_, err = _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
