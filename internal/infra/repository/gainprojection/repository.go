package gainprojection

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(ctx context.Context, gainProjection GainProjection) (*GainProjection, error)
	GetById(ctx context.Context, id string, userId string) (*GainProjection, error)
	Edit(ctx context.Context, gainProjection GainProjection) (*GainProjection, error)
	Remove(ctx context.Context, id string, userId string) error
	GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error)
	GetAll(ctx context.Context, params QueryParams) (*[]GainProjection, error)
	SaveGain(ctx context.Context, gain Gain) (*Gain, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}
