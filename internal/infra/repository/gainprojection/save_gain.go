package gainprojection

import "context"

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
