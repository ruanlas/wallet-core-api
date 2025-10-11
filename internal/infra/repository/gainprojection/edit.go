package gainprojection

import "context"

func (r *repository) Edit(ctx context.Context, gainProjection GainProjection) (*GainProjection, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		UPDATE gain_projection SET pay_in = ?, description = ?, value = ?, is_passive = ?, category_id = ?, is_already_done = ? 
		WHERE id = ? AND user_id = ?`)
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
		gainProjection.UserId,
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
