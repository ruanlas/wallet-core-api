package repository

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	Save(ctx context.Context, invoice Invoice) (*Invoice, error)
	GetById(ctx context.Context, id string, userId string) (*Invoice, error)
	Edit(ctx context.Context, invoice Invoice) (*Invoice, error)
	Remove(ctx context.Context, id string, userId string) error
	GetTotalRecords(ctx context.Context, params QueryParams) (*uint, error)
	GetAll(ctx context.Context, params QueryParams) (*[]Invoice, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{db: db}
}

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

func (r *repository) Remove(ctx context.Context, id string, userId string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	stmt, err := tx.PrepareContext(ctx, `DELETE FROM invoice WHERE id = ? AND user_id = ?`)
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
	query := `SELECT COUNT(*) as total_records FROM invoice WHERE MONTH(pay_at) = ? AND YEAR(pay_at) = ? AND user_id = ?`
	row := r.db.QueryRowContext(ctx, query, params.month, params.year, params.userId)
	err := row.Scan(&totalRecords)
	if err != nil {
		return nil, err
	}
	return &totalRecords, nil
}

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
