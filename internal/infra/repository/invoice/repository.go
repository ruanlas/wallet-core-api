package invoice

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(ctx context.Context, invoice Invoice) (*Invoice, error)
	GetById(ctx context.Context, id string, userId string) (*Invoice, error)
	Edit(ctx context.Context, invoice Invoice) (*Invoice, error)
	Remove(ctx context.Context, id string, userId string) error
	GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error)
	GetAll(ctx context.Context, params QueryParams) (*[]Invoice, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}
