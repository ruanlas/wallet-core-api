package invoice

import "context"

func (r *repository) Edit(ctx context.Context, invoice Invoice) (*Invoice, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		UPDATE invoice SET pay_at = ?, buy_at = ?, description = ?, value = ?, category_id = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		invoice.PayAt,
		invoice.BuyAt,
		invoice.Description,
		invoice.Value,
		invoice.Category.Id,
		invoice.PaymentType.Id,
		invoice.Id,
		invoice.UserId,
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}
