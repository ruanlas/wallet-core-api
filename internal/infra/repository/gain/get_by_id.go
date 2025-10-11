package gain

import (
	"context"
	"database/sql"
	"time"
)

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
