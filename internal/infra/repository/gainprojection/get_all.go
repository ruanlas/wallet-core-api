package gainprojection

import (
	"context"
	"database/sql"
	"time"
)

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
			MONTH(gp.pay_in) = ? AND YEAR(gp.pay_in) = ? AND gp.user_id = ?
		LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query, params.month, params.year, params.userId, params.limit, params.offset)
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
