package repository

import (
	"context"
	"database/sql"
	"time"
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

func (r *repository) Save(ctx context.Context, gain Gain) (*Gain, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO gain (id, created_at, pay_in, description, value, is_passive, user_id, category_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
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

func (r *repository) GetById(ctx context.Context, id string, userId string) (*Gain, error) {
	results, err := r.db.QueryContext(ctx, `
		SELECT
			g.id,
			g.created_at,
			g.pay_in,
			g.description,
			g.value,
			g.is_passive,
			g.user_id,
			gc.id,
			gc.category,
			g.gain_projection_id
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE g.id = ? AND g.user_id = ?`, id, userId)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	gain := &Gain{Category: GainCategory{}}
	if results.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		var gainProjectionId sql.NullString
		err := results.Scan(
			&gain.Id,
			&createdAtTimestamp,
			&gain.PayIn,
			&gain.Description,
			&value,
			&gain.IsPassive,
			&gain.UserId,
			&categoryId,
			&gain.Category.Category,
			&gainProjectionId,
		)
		if err != nil {
			return nil, err
		}
		gain.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		gain.Category.Id = uint(categoryId.Int64)
		gain.Value = value.Float64
		gain.GainProjectionId = gainProjectionId.String
	} else {
		return nil, nil
	}
	return gain, nil
}

func (r *repository) Edit(ctx context.Context, gain Gain) (*Gain, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		UPDATE gain SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ? 
		WHERE id = ? AND user_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		gain.PayIn,
		gain.Description,
		gain.Value,
		gain.IsPassive,
		gain.Category.Id,
		gain.Id,
		gain.UserId,
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

func (r *repository) Remove(ctx context.Context, id string, userId string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, `DELETE FROM gain WHERE id = ? AND user_id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, userId)
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
	query := `SELECT COUNT(*) as total_records FROM gain WHERE MONTH(pay_in) = ? AND YEAR(pay_in) = ? AND user_id = ?`
	row := r.db.QueryRowContext(ctx, query, params.month, params.year, params.userId)
	err := row.Scan(&totalRecords)
	if err != nil {
		return nil, err
	}
	return &totalRecords, nil
}

func (r *repository) GetAll(ctx context.Context, params QueryParams) (*[]Gain, error) {
	query := `
		SELECT
			g.id,
			g.created_at,
			g.pay_in,
			g.description,
			g.value,
			g.is_passive,
			g.user_id,
			gc.id,
			gc.category
		FROM
			gain g
		INNER JOIN gain_category gc ON 
			gc.id = g.category_id
		WHERE 
			MONTH(g.pay_in) = ? AND YEAR(g.pay_in) = ? AND g.user_id = ?
		LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query, params.month, params.year, params.userId, params.limit, params.offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gainList []Gain
	for rows.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		var g Gain
		var category GainCategory

		err := rows.Scan(
			&g.Id,
			&createdAtTimestamp,
			&g.PayIn,
			&g.Description,
			&value,
			&g.IsPassive,
			&g.UserId,
			&categoryId,
			&category.Category)
		if err != nil {
			return nil, err
		}
		g.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		g.Value = value.Float64
		category.Id = uint(categoryId.Int64)
		g.Category = category

		gainList = append(gainList, g)
	}

	return &gainList, nil
}
