package gainprojection

import (
	"context"
	"database/sql"
	"time"
)

func (r *repository) GetById(ctx context.Context, id string, userId string) (*GainProjection, error) {
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
		WHERE gp.id = ? AND gp.user_id = ?`, id, userId)
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
