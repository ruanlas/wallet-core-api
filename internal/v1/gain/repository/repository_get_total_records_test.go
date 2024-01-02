package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTotalRecordsSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	queryParams := NewQueryParamsBuilder().
		AddMonth(10).
		AddYear(2024).AddUserId("User1").Build()

	totalRecordsMock := sqlMock.NewRows([]string{
		"total_records",
	}).AddRow(5)

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT COUNT(*) as total_records FROM gain WHERE MONTH(pay_in) = ? AND YEAR(pay_in) = ? AND user_id = ?`).
		WithArgs(queryParams.month, queryParams.year, queryParams.userId).
		WillReturnRows(totalRecordsMock)

	totalRecords, err := _repository.GetTotalRecords(context.Background(), queryParams)
	assert.NoError(t, err)
	assert.Equal(t, uint(5), *totalRecords)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTotalRecordsScanFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	queryParams := NewQueryParamsBuilder().
		AddMonth(10).
		AddYear(2024).AddUserId("User1").Build()

	totalRecordsMock := sqlMock.NewRows([]string{
		"total_records",
	}).AddRow(nil).RowError(1, errors.New("An error has been ocurred"))

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT COUNT(*) as total_records FROM gain WHERE MONTH(pay_in) = ? AND YEAR(pay_in) = ? AND user_id = ?`).
		WithArgs(queryParams.month, queryParams.year, queryParams.userId).
		WillReturnRows(totalRecordsMock)

	_, err = _repository.GetTotalRecords(context.Background(), queryParams)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
