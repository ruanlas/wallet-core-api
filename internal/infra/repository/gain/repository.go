package gain

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(ctx context.Context, gain Gain) (*Gain, error)
	GetById(ctx context.Context, id string, userId string) (*Gain, error)
	Edit(ctx context.Context, gain Gain) (*Gain, error)
	Remove(ctx context.Context, id string, userId string) error
	GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error)
	GetAll(ctx context.Context, params QueryParams) (*[]Gain, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}
