package gain

import (
	"context"
	"database/sql"
	"time"
)

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
