package invoice

import "context"

func (r *repository) Save(ctx context.Context, invoice Invoice) (*Invoice, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO invoice (id, created_at, pay_at, buy_at, description, value, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		invoice.Id,
		invoice.CreatedAt.Unix(),
		invoice.PayAt,
		invoice.BuyAt,
		invoice.Description,
		invoice.Value,
		invoice.UserId,
		invoice.Category.Id,
		invoice.PaymentType.Id,
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
