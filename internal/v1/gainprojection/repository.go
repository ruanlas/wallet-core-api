package gainprojection

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	Save(ctx context.Context, gainProjections GainProjection) (*GainProjection, error)
	GetById(ctx context.Context, id string) (*GainProjection, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Save(ctx context.Context, gainProjection GainProjection) (*GainProjection, error) {
	// tx, err := r.db.Begin()
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO gain_projection (id, created_at, pay_in, description, value, is_passive, is_done, user_id, category_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		gainProjection.Id,
		gainProjection.CreatedAt.Unix(),
		gainProjection.PayIn,
		gainProjection.Description,
		gainProjection.Value,
		gainProjection.IsPassive,
		gainProjection.IsDone,
		gainProjection.UserId,
		gainProjection.Category.Id,
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &gainProjection, nil
}

func (r *repository) GetById(ctx context.Context, id string) (*GainProjection, error) {
	results, err := r.db.QueryContext(ctx, `
		SELECT
			gp.id,
			gp.created_at,
			gp.pay_in,
			gp.description,
			gp.value,
			gp.is_passive,
			gp.is_done,
			gp.user_id,
			gc.id,
			gc.category
		FROM
			gain_projection gp
		INNER JOIN gain_category gc ON 
			gc.id = gp.category_id
		WHERE gp.id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	gainProjection := &GainProjection{Category: GainCategory{}}
	if results.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		err := results.Scan(
			&gainProjection.Id,
			&createdAtTimestamp,
			&gainProjection.PayIn,
			&gainProjection.Description,
			&value,
			&gainProjection.IsPassive,
			&gainProjection.IsDone,
			&gainProjection.UserId,
			&categoryId,
			&gainProjection.Category.Category,
		)
		if err != nil {
			return nil, err
		}
		gainProjection.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		gainProjection.Category.Id = uint(categoryId.Int64)
		gainProjection.Value = value.Float64
	} else {
		return nil, nil
	}
	return gainProjection, nil
}
