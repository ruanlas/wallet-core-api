package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetGainProjectionByIdSuccess(t *testing.T) {
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
		AddCategory(GainCategory{Id: 1, Category: "Salário"}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
		Build()

	rowsGainProjectionMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"description",
		"value",
		"is_passive",
		"is_already_done",
		"user_id",
		"category_id",
		"category",
	}).AddRow(
		gainPMock.Id,
		gainPMock.CreatedAt.Unix(),
		gainPMock.PayIn,
		gainPMock.Description,
		gainPMock.Value,
		gainPMock.IsPassive,
		gainPMock.IsAlreadyDone,
		gainPMock.UserId,
		gainPMock.Category.Id,
		gainPMock.Category.Category,
	)

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			gp.id,
			gp.created_at,
			gp.pay_in,
			gp.description,
			gp.value,
			gp.is_passive,
			gp.is_already_done,
			gp.user_id,
			gc.id,
			gc.category
		FROM
			gain_projection gp
		INNER JOIN gain_category gc ON 
			gc.id = gp.category_id
		WHERE gp.id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8").
		WillReturnRows(rowsGainProjectionMock)

	gainPSaved, err := _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", gainPSaved.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetGainProjectionByIdQueryFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			gp.id,
			gp.created_at,
			gp.pay_in,
			gp.description,
			gp.value,
			gp.is_passive,
			gp.is_already_done,
			gp.user_id,
			gc.id,
			gc.category
		FROM
			gain_projection gp
		INNER JOIN gain_category gc ON 
			gc.id = gp.category_id
		WHERE gp.id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8").
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetGainProjectionByIdRowEmpty(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	rowsGainProjectionMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"description",
		"value",
		"is_passive",
		"is_already_done",
		"user_id",
		"category_id",
		"category",
	})

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			gp.id,
			gp.created_at,
			gp.pay_in,
			gp.description,
			gp.value,
			gp.is_passive,
			gp.is_already_done,
			gp.user_id,
			gc.id,
			gc.category
		FROM
			gain_projection gp
		INNER JOIN gain_category gc ON 
			gc.id = gp.category_id
		WHERE gp.id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8").
		WillReturnRows(rowsGainProjectionMock)

	gainPSaved, err := _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.NoError(t, err)
	assert.Empty(t, gainPSaved)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetGainProjectionByIdRowScanFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	rowsGainProjectionMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"description",
		"value",
		"is_passive",
		"is_already_done",
		"user_id",
		"category_id",
		"category",
	}).AddRow(
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil).
		RowError(1, errors.New("An error has been ocurred"))

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			gp.id,
			gp.created_at,
			gp.pay_in,
			gp.description,
			gp.value,
			gp.is_passive,
			gp.is_already_done,
			gp.user_id,
			gc.id,
			gc.category
		FROM
			gain_projection gp
		INNER JOIN gain_category gc ON 
			gc.id = gp.category_id
		WHERE gp.id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8").
		WillReturnRows(rowsGainProjectionMock)

	_, err = _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
