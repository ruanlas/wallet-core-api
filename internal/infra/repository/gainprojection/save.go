package gainprojection

import "context"

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
