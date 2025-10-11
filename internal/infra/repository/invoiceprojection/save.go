package invoiceprojection

import "context"

func (r *repository) Save(ctx context.Context, invoiceProjection InvoiceProjection) (*InvoiceProjection, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO invoice_projection (id, created_at, pay_in, buy_at, description, value, is_already_done, user_id, category_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		invoiceProjection.Id,
		invoiceProjection.CreatedAt.Unix(),
		invoiceProjection.PayIn,
		invoiceProjection.BuyAt,
		invoiceProjection.Description,
		invoiceProjection.Value,
		invoiceProjection.IsAlreadyDone,
		invoiceProjection.UserId,
		invoiceProjection.Category.Id,
		invoiceProjection.PaymentType.Id,
	)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &invoiceProjection, nil
}
