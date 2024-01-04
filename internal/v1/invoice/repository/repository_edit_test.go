package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestEditInvoiceSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	invoiceMock := NewInvoiceBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice SET pay_at = ?, buy_at = ?, description = ?, value = ?, category_id = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			invoiceMock.PayAt,
			invoiceMock.BuyAt,
			invoiceMock.Description,
			invoiceMock.Value,
			invoiceMock.Category.Id,
			invoiceMock.PaymentType.Id,
			invoiceMock.Id,
			invoiceMock.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	invoiceEdited, err := _repository.Edit(context.Background(), *invoiceMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", invoiceEdited.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoiceBeginFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	invoiceMock := NewInvoiceBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoicePrepareFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	invoiceMock := NewInvoiceBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice SET pay_at = ?, buy_at = ?, description = ?, value = ?, category_id = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoiceExecFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	invoiceMock := NewInvoiceBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice SET pay_at = ?, buy_at = ?, description = ?, value = ?, category_id = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			invoiceMock.PayAt,
			invoiceMock.BuyAt,
			invoiceMock.Description,
			invoiceMock.Value,
			invoiceMock.Category.Id,
			invoiceMock.PaymentType.Id,
			invoiceMock.Id,
			invoiceMock.UserId).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoiceCommitFail(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	invoiceMock := NewInvoiceBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddCreatedAt(now).
		AddPayAt(now).
		AddBuyAt(now).
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice SET pay_at = ?, buy_at = ?, description = ?, value = ?, category_id = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			invoiceMock.PayAt,
			invoiceMock.BuyAt,
			invoiceMock.Description,
			invoiceMock.Value,
			invoiceMock.Category.Id,
			invoiceMock.PaymentType.Id,
			invoiceMock.Id,
			invoiceMock.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
