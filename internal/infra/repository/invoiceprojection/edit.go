package invoiceprojection

import "context"

func (r *repository) Edit(ctx context.Context, invoiceProjection InvoiceProjection) (*InvoiceProjection, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		UPDATE invoice_projection SET pay_in = ?, buy_at = ?, description = ?, value = ?, category_id = ?, is_already_done = ?, payment_type_id = ? 
		WHERE id = ? AND user_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		invoiceProjection.PayIn,
		invoiceProjection.BuyAt,
		invoiceProjection.Description,
		invoiceProjection.Value,
		invoiceProjection.Category.Id,
		invoiceProjection.IsAlreadyDone,
		invoiceProjection.PaymentType.Id,
		invoiceProjection.Id,
		invoiceProjection.UserId,
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
