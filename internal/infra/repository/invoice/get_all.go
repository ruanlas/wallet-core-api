package invoice

import (
	"context"
	"database/sql"
	"time"
)

func (r *repository) GetAll(ctx context.Context, params QueryParams) (*[]Invoice, error) {
	query := `
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
		WHERE 
			MONTH(i.pay_at) = ? AND YEAR(i.pay_at) = ? AND i.user_id = ?
		LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query, params.month, params.year, params.userId, params.limit, params.offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoiceList []Invoice
	for rows.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var paymentTypeId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		var invoiceProjectionId sql.NullString

		var invoice Invoice
		var category InvoiceCategory
		var paymentType PaymentType

		err := rows.Scan(
			&invoice.Id,
			&createdAtTimestamp,
			&invoice.PayAt,
			&invoice.BuyAt,
			&invoice.Description,
			&value,
			&invoice.UserId,
			&invoiceProjectionId,
			&categoryId,
			&category.Category,
			&paymentTypeId,
			&paymentType.Type)
		if err != nil {
			return nil, err
		}
		invoice.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		invoice.Value = value.Float64
		category.Id = uint(categoryId.Int64)
		invoice.Category = category
		paymentType.Id = uint(paymentTypeId.Int64)
		invoice.PaymentType = paymentType
		invoice.InvoiceProjectionId = invoiceProjectionId.String

		invoiceList = append(invoiceList, invoice)
	}

	return &invoiceList, nil
}
