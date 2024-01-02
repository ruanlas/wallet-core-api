package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	queryParams := NewQueryParamsBuilder().
		AddMonth(10).
		AddYear(2024).
		AddUserId("User1").
		AddLimit(10).
		AddOffset(0).
		Build()

	now := time.Now()
	gainPMock := NewGainBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddIsPassive(true).
		AddCategory(GainCategory{Id: 1, Category: "Salário"}).
		AddDescription("Description de teste").
		AddValue(500.50).
		AddUserId("User1").
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
	}).AddRow(
		gainPMock.Id,
		gainPMock.CreatedAt.Unix(),
		gainPMock.PayIn,
		gainPMock.Description,
		gainPMock.Value,
		gainPMock.IsPassive,
		gainPMock.UserId,
		gainPMock.Category.Id,
		gainPMock.Category.Category,
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
			gc.category
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE 
			MONTH(g.pay_in) = ? AND YEAR(g.pay_in) = ? AND g.user_id = ?
		LIMIT ? OFFSET ?`).
		WithArgs(queryParams.month, queryParams.year, queryParams.userId, queryParams.limit, queryParams.offset).
		WillReturnRows(rowsGainMock)

	listGain, err := _repository.GetAll(context.Background(), queryParams)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", (*listGain)[0].Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllQueryFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	queryParams := NewQueryParamsBuilder().
		AddMonth(10).
		AddYear(2024).
		AddUserId("User1").
		AddLimit(10).
		AddOffset(0).
		Build()

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
			gc.category
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE 
			MONTH(g.pay_in) = ? AND YEAR(g.pay_in) = ? AND g.user_id = ?
		LIMIT ? OFFSET ?`).
		WithArgs(queryParams.month, queryParams.year, queryParams.userId, queryParams.limit, queryParams.offset).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.GetAll(context.Background(), queryParams)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllScanFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	queryParams := NewQueryParamsBuilder().
		AddMonth(10).
		AddYear(2024).
		AddUserId("User1").
		AddLimit(10).
		AddOffset(0).
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
	}).AddRow(
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	).RowError(1, errors.New("An error has been ocurred"))

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
			gc.category
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE 
			MONTH(g.pay_in) = ? AND YEAR(g.pay_in) = ? AND g.user_id = ?
		LIMIT ? OFFSET ?`).
		WithArgs(queryParams.month, queryParams.year, queryParams.userId, queryParams.limit, queryParams.offset).
		WillReturnRows(rowsGainMock)

	_, err = _repository.GetAll(context.Background(), queryParams)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
