package gain

import "context"

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
