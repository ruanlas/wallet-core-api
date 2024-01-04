package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSaveInvoiceSuccess(t *testing.T) {
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
		INSERT INTO invoice (id, created_at, pay_at, buy_at, description, value, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			invoiceMock.Id,
			invoiceMock.CreatedAt.Unix(),
			invoiceMock.PayAt,
			invoiceMock.BuyAt,
			invoiceMock.Description,
			invoiceMock.Value,
			invoiceMock.UserId,
			invoiceMock.Category.Id,
			invoiceMock.PaymentType.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	invoiceSaved, err := _repository.Save(context.Background(), *invoiceMock)
	assert.NoError(t, err)
	assert.Equal(t, "519fd73e-45e6-4471-8a66-5057486f5cc8", invoiceSaved.Id)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoiceBeginFail(t *testing.T) {
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

	_, err = _repository.Save(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoicePrepareFail(t *testing.T) {
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
		INSERT INTO invoice (id, created_at, pay_at, buy_at, description, value, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoiceExecFail(t *testing.T) {
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
		INSERT INTO invoice (id, created_at, pay_at, buy_at, description, value, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			invoiceMock.Id,
			invoiceMock.CreatedAt.Unix(),
			invoiceMock.PayAt,
			invoiceMock.BuyAt,
			invoiceMock.Description,
			invoiceMock.Value,
			invoiceMock.UserId,
			invoiceMock.Category.Id,
			invoiceMock.PaymentType.Id).
		WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSaveInvoiceCommitFail(t *testing.T) {
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
		INSERT INTO invoice (id, created_at, pay_at, buy_at, description, value, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`).
		ExpectExec().
		WithArgs(
			invoiceMock.Id,
			invoiceMock.CreatedAt.Unix(),
			invoiceMock.PayAt,
			invoiceMock.BuyAt,
			invoiceMock.Description,
			invoiceMock.Value,
			invoiceMock.UserId,
			invoiceMock.Category.Id,
			invoiceMock.PaymentType.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit().WillReturnError(errors.New("An error has been ocurred"))

	_, err = _repository.Save(context.Background(), *invoiceMock)
	assert.Error(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
