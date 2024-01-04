package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetInvoiceProjectionByIdSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	invoicePMock := NewInvoiceProjectionBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayIn(now).
		AddBuyAt(now).
		AddIsAlreadyDone(false).
		AddCategory(InvoiceCategory{Id: 1, Category: "Moradia"}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2, Type: "Transferência"}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	rowsInvoiceProjectionMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"buy_at",
		"description",
		"value",
		"is_already_done",
		"user_id",
		"category_id",
		"category",
		"payment_type_id",
		"payment_type",
	}).AddRow(
		invoicePMock.Id,
		invoicePMock.CreatedAt.Unix(),
		invoicePMock.PayIn,
		invoicePMock.BuyAt,
		invoicePMock.Description,
		invoicePMock.Value,
		invoicePMock.IsAlreadyDone,
		invoicePMock.UserId,
		invoicePMock.Category.Id,
		invoicePMock.Category.Category,
		invoicePMock.PaymentType.Id,
		invoicePMock.PaymentType.Type,
	)

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			ip.id,
			ip.created_at,
			ip.pay_in,
			ip.buy_at,
			ip.description,
			ip.value,
			ip.is_already_done,
			ip.user_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice_projection ip
		INNER JOIN invoice_category ic ON 
			ic.id = ip.category_id
		INNER JOIN payment_type pt ON
			pt.id = ip.payment_type_id
		WHERE ip.id = ? AND ip.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnRows(rowsInvoiceProjectionMock)

	invoicePReturn, err := _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", invoicePReturn.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetInvoiceProjectionByIdQueryFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			ip.id,
			ip.created_at,
			ip.pay_in,
			ip.buy_at,
			ip.description,
			ip.value,
			ip.is_already_done,
			ip.user_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice_projection ip
		INNER JOIN invoice_category ic ON 
			ic.id = ip.category_id
		INNER JOIN payment_type pt ON
			pt.id = ip.payment_type_id
		WHERE ip.id = ? AND ip.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetInvoiceProjectionByIdRowEmpty(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	rowsInvoiceProjectionMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"buy_at",
		"description",
		"value",
		"is_already_done",
		"user_id",
		"category_id",
		"category",
		"payment_type",
		"payment_type_id",
	})

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			ip.id,
			ip.created_at,
			ip.pay_in,
			ip.buy_at,
			ip.description,
			ip.value,
			ip.is_already_done,
			ip.user_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice_projection ip
		INNER JOIN invoice_category ic ON 
			ic.id = ip.category_id
		INNER JOIN payment_type pt ON
			pt.id = ip.payment_type_id
		WHERE ip.id = ? AND ip.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnRows(rowsInvoiceProjectionMock)

	invoicePReturn, err := _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.NoError(t, err)
	assert.Empty(t, invoicePReturn)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetInvoiceProjectionByIdRowScanFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	rowsInvoiceProjectionMock := sqlMock.NewRows([]string{
		"id",
		"created_at",
		"pay_in",
		"buy_at",
		"description",
		"value",
		"is_already_done",
		"user_id",
		"category_id",
		"category",
		"payment_type_id",
		"payment_type",
	}).AddRow(
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil).
		RowError(1, errors.New("An error has been ocurred"))

	_repository := New(dbMock)

	sqlMock.ExpectQuery(`
		SELECT
			ip.id,
			ip.created_at,
			ip.pay_in,
			ip.buy_at,
			ip.description,
			ip.value,
			ip.is_already_done,
			ip.user_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice_projection ip
		INNER JOIN invoice_category ic ON 
			ic.id = ip.category_id
		INNER JOIN payment_type pt ON
			pt.id = ip.payment_type_id
		WHERE ip.id = ? AND ip.user_id = ?`).
		WithArgs("519fd73e-45e6-4471-8a66-5057486f5cc8", "User1").
		WillReturnRows(rowsInvoiceProjectionMock)

	_, err = _repository.GetById(context.Background(), "519fd73e-45e6-4471-8a66-5057486f5cc8", "User1")
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
