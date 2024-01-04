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
	invoiceMock := NewInvoiceBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(InvoiceCategory{Id: 1, Category: "Moradia"}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2, Type: "Transferência"}).
		AddValue(500.50).
		AddInvoiceProjectionId("4c3939f7-2b39-4bb1-8367-54fc56abea3a").
		AddUserId("User1").
		Build()

	rowsInvoiceMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_at",
		"buy_at",
		"description",
		"value",
		"user_id",
		"invoice_projection_id",
		"category_id",
		"category",
		"payment_type_id",
		"payment_type",
	}).AddRow(
		invoiceMock.Id,
		invoiceMock.CreatedAt.Unix(),
		invoiceMock.PayAt,
		invoiceMock.BuyAt,
		invoiceMock.Description,
		invoiceMock.Value,
		invoiceMock.UserId,
		invoiceMock.InvoiceProjectionId,
		invoiceMock.Category.Id,
		invoiceMock.Category.Category,
		invoiceMock.PaymentType.Id,
		invoiceMock.PaymentType.Type,
	)

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			i.id,
			i.created_at,
			i.pay_at,
			i.buy_at,
			i.description,
			i.value,
			i.user_id,
			i.invoice_projection_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice i
		INNER JOIN invoice_category ic ON 
			ic.id = i.category_id
		INNER JOIN payment_type pt ON
			pt.id = i.payment_type_id
		WHERE 
			MONTH(i.pay_at) = ? AND YEAR(i.pay_at) = ? AND i.user_id = ?
		LIMIT ? OFFSET ?`).
		WithArgs(queryParams.month, queryParams.year, queryParams.userId, queryParams.limit, queryParams.offset).
		WillReturnRows(rowsInvoiceMock)

	listInvoice, err := _repository.GetAll(context.Background(), queryParams)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", (*listInvoice)[0].Id)

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
			i.id,
			i.created_at,
			i.pay_at,
			i.buy_at,
			i.description,
			i.value,
			i.user_id,
			i.invoice_projection_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice i
		INNER JOIN invoice_category ic ON 
			ic.id = i.category_id
		INNER JOIN payment_type pt ON
			pt.id = i.payment_type_id
		WHERE 
			MONTH(i.pay_at) = ? AND YEAR(i.pay_at) = ? AND i.user_id = ?
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

	rowsInvoiceMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_at",
		"buy_at",
		"description",
		"value",
		"user_id",
		"invoice_projection_id",
		"category_id",
		"category",
		"payment_type_id",
		"payment_type",
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
		nil,
		nil,
		nil,
	).RowError(1, errors.New("An error has been ocurred"))

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			i.id,
			i.created_at,
			i.pay_at,
			i.buy_at,
			i.description,
			i.value,
			i.user_id,
			i.invoice_projection_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice i
		INNER JOIN invoice_category ic ON 
			ic.id = i.category_id
		INNER JOIN payment_type pt ON
			pt.id = i.payment_type_id
		WHERE 
			MONTH(i.pay_at) = ? AND YEAR(i.pay_at) = ? AND i.user_id = ?
		LIMIT ? OFFSET ?`).
		WithArgs(queryParams.month, queryParams.year, queryParams.userId, queryParams.limit, queryParams.offset).
		WillReturnRows(rowsInvoiceMock)

	_, err = _repository.GetAll(context.Background(), queryParams)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
