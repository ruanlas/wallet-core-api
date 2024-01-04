package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestEditInvoiceProjectionSuccess(t *testing.T) {
	dbMock, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	now := time.Now()
	invoicePMock := NewInvoiceProjectionBuilder().
		AddId("519fd73e-45e6-4471-8a66-5057486f5cc8").
		AddPayIn(now).
		AddBuyAt(now).
		AddIsAlreadyDone(false).
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice_projection SET pay_in = ?, buy_at = ?, description = ?, value = ?, category_id = ?, is_already_done = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			invoicePMock.PayIn,
			invoicePMock.BuyAt,
			invoicePMock.Description,
			invoicePMock.Value,
			invoicePMock.Category.Id,
			invoicePMock.IsAlreadyDone,
			invoicePMock.PaymentType.Id,
			invoicePMock.Id,
			invoicePMock.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	invoicePEdited, err := _repository.Edit(context.Background(), *invoicePMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", invoicePEdited.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoiceProjectionBeginFail(t *testing.T) {
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
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoiceProjectionPrepareFail(t *testing.T) {
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
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice_projection SET pay_in = ?, buy_at = ?, description = ?, value = ?, category_id = ?, is_already_done = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoiceProjectionExecFail(t *testing.T) {
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
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice_projection SET pay_in = ?, buy_at = ?, description = ?, value = ?, category_id = ?, is_already_done = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			invoicePMock.PayIn,
			invoicePMock.BuyAt,
			invoicePMock.Description,
			invoicePMock.Value,
			invoicePMock.Category.Id,
			invoicePMock.IsAlreadyDone,
			invoicePMock.PaymentType.Id,
			invoicePMock.Id,
			invoicePMock.UserId).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditInvoiceProjectionCommitFail(t *testing.T) {
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
		AddCategory(InvoiceCategory{Id: 1}).
		AddDescription("Description de teste").
		AddPaymentType(PaymentType{Id: 2}).
		AddValue(500.50).
		AddUserId("User1").
		Build()

	_repository := New(dbMock)

	sqlMock.ExpectBegin()
	sqlMock.ExpectPrepare(`
		UPDATE invoice_projection SET pay_in = ?, buy_at = ?, description = ?, value = ?, category_id = ?, is_already_done = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			invoicePMock.PayIn,
			invoicePMock.BuyAt,
			invoicePMock.Description,
			invoicePMock.Value,
			invoicePMock.Category.Id,
			invoicePMock.IsAlreadyDone,
			invoicePMock.PaymentType.Id,
			invoicePMock.Id,
			invoicePMock.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Edit(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
