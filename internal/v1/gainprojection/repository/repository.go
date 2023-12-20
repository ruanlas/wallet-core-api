package repository

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	Save(ctx context.Context, gainProjection GainProjection) (*GainProjection, error)
	GetById(ctx context.Context, id string) (*GainProjection, error)
	Edit(ctx context.Context, gainProjection GainProjection) (*GainProjection, error)
	Remove(ctx context.Context, id string) error
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

func (r *repository) Save(ctx context.Context, gainProjection GainProjection) (*GainProjection, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO gain_projection (id, created_at, pay_in, description, value, is_passive, is_already_done, user_id, category_id) 
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
		gainProjection.IsAlreadyDone,
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
			gp.is_already_done,
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
			&gainProjection.IsAlreadyDone,
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

func (r *repository) Edit(ctx context.Context, gainProjection GainProjection) (*GainProjection, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		UPDATE gain_projection SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ?, is_already_done = ? 
		WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		gainProjection.PayIn,
		gainProjection.Description,
		gainProjection.Value,
		gainProjection.IsPassive,
		gainProjection.Category.Id,
		gainProjection.IsAlreadyDone,
		gainProjection.Id,
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

func (r *repository) Remove(ctx context.Context, id string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, `DELETE FROM gain_projection WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error) {
	var totalRecords uint
	query := `SELECT COUNT(*) as total_records FROM gain_projection WHERE MONTH(pay_in) = ? AND YEAR(pay_in) = ?`
	row := r.db.QueryRowContext(ctx, query, params.month, params.year)
	err := row.Scan(&totalRecords)
	if err != nil {
		return nil, err
	}
	return &totalRecords, nil
}

func (r *repository) GetAll(ctx context.Context, params QueryParams) (*[]GainProjection, error) {
	query := `
		SELECT
			gp.id,
			gp.created_at,
			gp.pay_in,
			gp.description,
			gp.value,
			gp.is_passive,
			gp.is_already_done,
			gp.user_id,
			gc.id,
			gc.category
		FROM
			gain_projection gp
		INNER JOIN gain_category gc ON 
			gc.id = gp.category_id
		WHERE 
			MONTH(gp.pay_in) = ? AND YEAR(gp.pay_in) = ?
		LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query, params.month, params.year, params.limit, params.offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gainProjectionList []GainProjection
	for rows.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		var gp GainProjection
		var category GainCategory

		err := rows.Scan(
			&gp.Id,
			&createdAtTimestamp,
			&gp.PayIn,
			&gp.Description,
			&value,
			&gp.IsPassive,
			&gp.IsAlreadyDone,
			&gp.UserId,
			&categoryId,
			&category.Category)
		if err != nil {
			return nil, err
		}
		gp.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		gp.Value = value.Float64
		category.Id = uint(categoryId.Int64)
		gp.Category = category

		gainProjectionList = append(gainProjectionList, gp)
	}

	return &gainProjectionList, nil
}

func (r *repository) SaveGain(ctx context.Context, gain Gain) (*Gain, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO gain (id, created_at, pay_in, description, value, is_passive, user_id, category_id, gain_projection_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		gain.Id,
		gain.CreatedAt.Unix(),
		gain.PayIn,
		gain.Description,
		gain.Value,
		gain.IsPassive,
		gain.UserId,
		gain.Category.Id,
		gain.GainProjectionId,
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &gain, nil
}
