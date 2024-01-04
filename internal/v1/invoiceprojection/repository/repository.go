package repository

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	Save(ctx context.Context, invoiceProjection InvoiceProjection) (*InvoiceProjection, error)
	GetById(ctx context.Context, id string, userId string) (*InvoiceProjection, error)
	Edit(ctx context.Context, invoiceProjection InvoiceProjection) (*InvoiceProjection, error)
	Remove(ctx context.Context, id string, userId string) error
	GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error)
	GetAll(ctx context.Context, params QueryParams) (*[]InvoiceProjection, error)
	SaveInvoice(ctx context.Context, invoice Invoice) (*Invoice, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

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

func (r *repository) Remove(ctx context.Context, id string, userId string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, `DELETE FROM invoice_projection WHERE id = ? AND user_id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, userId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error) {
	var totalRecords uint
	query := `SELECT COUNT(*) as total_records FROM invoice_projection WHERE MONTH(pay_in) = ? AND YEAR(pay_in) = ? AND user_id = ?`
	row := r.db.QueryRowContext(ctx, query, params.month, params.year, params.userId)
	err := row.Scan(&totalRecords)
	if err != nil {
		return nil, err
	}
	return &totalRecords, nil
}

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

func (r *repository) SaveInvoice(ctx context.Context, invoice Invoice) (*Invoice, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO invoice (id, created_at, pay_at, buy_at, description, value, user_id, category_id, invoice_projection_id, payment_type_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
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
		invoice.InvoiceProjectionId,
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
