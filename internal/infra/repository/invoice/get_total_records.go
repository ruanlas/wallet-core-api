package invoice

import "context"

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
