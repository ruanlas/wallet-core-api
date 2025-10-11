package invoice

import (
	"context"
	"database/sql"
	"time"
)

func (r *repository) GetById(ctx context.Context, id string, userId string) (*Invoice, error) {
	results, err := r.db.QueryContext(ctx, `
		SELECT
			i.id,
			i.created_at,
			i.pay_at,
			i.buy_at,
			i.description,
			i.value,
			i.user_id,
			i.invoice_projection_id,
			ic.id,
			ic.category,
			pt.id,
			pt.type_name
		FROM
			invoice i
		INNER JOIN invoice_category ic ON 
			ic.id = i.category_id
		INNER JOIN payment_type pt ON
			pt.id = i.payment_type_id
		WHERE i.id = ? AND i.user_id = ?`, id, userId)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	invoice := &Invoice{Category: InvoiceCategory{}, PaymentType: PaymentType{}}
	if results.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var paymentTypeId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		var invoiceProjectionId sql.NullString
		err := results.Scan(
			&invoice.Id,
			&createdAtTimestamp,
			&invoice.PayAt,
			&invoice.BuyAt,
			&invoice.Description,
			&value,
			&invoice.UserId,
			&invoiceProjectionId,
			&categoryId,
			&invoice.Category.Category,
			&paymentTypeId,
			&invoice.PaymentType.Type,
		)
		if err != nil {
			return nil, err
		}
		invoice.InvoiceProjectionId = invoiceProjectionId.String
		invoice.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		invoice.Category.Id = uint(categoryId.Int64)
		invoice.PaymentType.Id = uint(paymentTypeId.Int64)
		invoice.Value = value.Float64
	} else {
		return nil, nil
	}
	return invoice, nil
}
