package invoiceprojection

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(ctx context.Context, invoiceProjection InvoiceProjection) (*InvoiceProjection, error)
	GetById(ctx context.Context, id string, userId string) (*InvoiceProjection, error)
	Edit(ctx context.Context, invoiceProjection InvoiceProjection) (*InvoiceProjection, error)
	Remove(ctx context.Context, id string, userId string) error
	GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error)
	GetAll(ctx context.Context, params QueryParams) (*[]InvoiceProjection, error)
	SaveInvoice(ctx context.Context, invoice Invoice) (*Invoice, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}
