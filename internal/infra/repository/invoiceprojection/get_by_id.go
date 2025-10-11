package invoiceprojection

import (
	"context"
	"database/sql"
	"time"
)

func (r *repository) GetById(ctx context.Context, id string, userId string) (*InvoiceProjection, error) {
	results, err := r.db.QueryContext(ctx, `
		SELECT
			ip.id,
			ip.created_at,
			ip.pay_in,
			ip.buy_at,
			ip.description,
			ip.value,
			ip.is_already_done,
			ip.user_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice_projection ip
		INNER JOIN invoice_category ic ON 
			ic.id = ip.category_id
		INNER JOIN payment_type pt ON
			pt.id = ip.payment_type_id
		WHERE ip.id = ? AND ip.user_id = ?`, id, userId)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	invoiceProjection := &InvoiceProjection{Category: InvoiceCategory{}, PaymentType: PaymentType{}}
	if results.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var paymentTypeId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		err := results.Scan(
			&invoiceProjection.Id,
			&createdAtTimestamp,
			&invoiceProjection.PayIn,
			&invoiceProjection.BuyAt,
			&invoiceProjection.Description,
			&value,
			&invoiceProjection.IsAlreadyDone,
			&invoiceProjection.UserId,
			&categoryId,
			&invoiceProjection.Category.Category,
			&paymentTypeId,
			&invoiceProjection.PaymentType.Type,
		)
		if err != nil {
			return nil, err
		}
		invoiceProjection.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		invoiceProjection.Category.Id = uint(categoryId.Int64)
		invoiceProjection.PaymentType.Id = uint(paymentTypeId.Int64)
		invoiceProjection.Value = value.Float64
	} else {
		return nil, nil
	}
	return invoiceProjection, nil
}
