package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSaveInvoiceProjectionSuccess(t *testing.T) {
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
		INSERT INTO invoice_projection (id, created_at, pay_in, buy_at, description, value, is_already_done, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			invoicePMock.Id,
			invoicePMock.CreatedAt.Unix(),
			invoicePMock.PayIn,
			invoicePMock.BuyAt,
			invoicePMock.Description,
			invoicePMock.Value,
			invoicePMock.IsAlreadyDone,
			invoicePMock.UserId,
			invoicePMock.Category.Id,
			invoicePMock.PaymentType.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	invoicePSaved, err := _repository.Save(context.Background(), *invoicePMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", invoicePSaved.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoiceProjectionBeginFail(t *testing.T) {
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

	_, err = _repository.Save(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoiceProjectionPrepareFail(t *testing.T) {
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
		INSERT INTO invoice_projection (id, created_at, pay_in, buy_at, description, value, is_already_done, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoiceProjectionExecFail(t *testing.T) {
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
		INSERT INTO invoice_projection (id, created_at, pay_in, buy_at, description, value, is_already_done, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			invoicePMock.Id,
			invoicePMock.CreatedAt.Unix(),
			invoicePMock.PayIn,
			invoicePMock.BuyAt,
			invoicePMock.Description,
			invoicePMock.Value,
			invoicePMock.IsAlreadyDone,
			invoicePMock.UserId,
			invoicePMock.Category.Id,
			invoicePMock.PaymentType.Id).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoiceProjectionCommitFail(t *testing.T) {
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
		INSERT INTO invoice_projection (id, created_at, pay_in, buy_at, description, value, is_already_done, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			invoicePMock.Id,
			invoicePMock.CreatedAt.Unix(),
			invoicePMock.PayIn,
			invoicePMock.BuyAt,
			invoicePMock.Description,
			invoicePMock.Value,
			invoicePMock.IsAlreadyDone,
			invoicePMock.UserId,
			invoicePMock.Category.Id,
			invoicePMock.PaymentType.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *invoicePMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
