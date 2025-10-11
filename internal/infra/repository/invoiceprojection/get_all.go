package invoiceprojection

import (
	"context"
	"database/sql"
	"time"
)

func (r *repository) GetAll(ctx context.Context, params QueryParams) (*[]InvoiceProjection, error) {
	query := `
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
		WHERE 
			MONTH(ip.pay_in) = ? AND YEAR(ip.pay_in) = ? AND ip.user_id = ?
		LIMIT ? OFFSET ?`
	rows, err := r.db.QueryContext(ctx, query, params.month, params.year, params.userId, params.limit, params.offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoiceProjectionList []InvoiceProjection
	for rows.Next() {
		var value sql.NullFloat64
		var categoryId sql.NullInt64
		var paymentTypeId sql.NullInt64
		var createdAtTimestamp sql.NullInt64
		var ip InvoiceProjection
		var category InvoiceCategory
		var paymentType PaymentType

		err := rows.Scan(
			&ip.Id,
			&createdAtTimestamp,
			&ip.PayIn,
			&ip.BuyAt,
			&ip.Description,
			&value,
			&ip.IsAlreadyDone,
			&ip.UserId,
			&categoryId,
			&category.Category,
			&paymentTypeId,
			&paymentType.Type)
		if err != nil {
			return nil, err
		}
		ip.CreatedAt = time.Unix(createdAtTimestamp.Int64, 0)
		ip.Value = value.Float64
		category.Id = uint(categoryId.Int64)
		ip.Category = category
		paymentType.Id = uint(paymentTypeId.Int64)
		ip.PaymentType = paymentType

		invoiceProjectionList = append(invoiceProjectionList, ip)
	}

	return &invoiceProjectionList, nil
}
