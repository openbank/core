// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertProductFee(ctx context.Context, pf *db.ProductFee) error {
	const stmt = `INSERT INTO product_fees
		(fee_id, name, calculation_method, required, amount, percent, apply_date_method, trigger, accounting_rules, amortization_settings, is_active, is_taxable, create_time, update_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		pf.FeeID, pf.Name, pf.CalculationMethod, pf.Required, pf.Amount, pf.Percent, pf.ApplyDateMethod, pf.Trigger, pf.AccountingRules, pf.AmortizationSettings, pf.IsActive, pf.IsTaxable, pf.CreateTime, pf.UpdateTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListProductFees(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.ProductFee, *db.ListPosition, error) {
	if filter != "" {
		return db.ListStat{}, nil, nil, fmt.Errorf("filter is unimplemented")
	}
	if orderBy == "" {
		orderBy = "id"
	}
	whereClause := `
		(
			(all_entries.ordered_idx > $2) OR
			(all_entries.ordered_idx = $2 AND all_entries.id > $3)
		)
	`
	if after == nil {
		// Use placeholder values but always evaluate to true.
		whereClause = "$2::INTEGER = $3"
		after = &db.ListPosition{
			Data: 0,
			ID:   0,
		}
	}
	const stmt = `WITH all_entries AS (
			SELECT
				quote_ident($1) AS ordered_idx, product_fees.id,

				product_fees.fee_id,
				product_fees.name,
				product_fees.calculation_method,
				product_fees.required,
				product_fees.amount,
				product_fees.percent,
				product_fees.apply_date_method,
				product_fees.trigger,
				product_fees.accounting_rules,
				product_fees.amortization_settings,
				product_fees.is_active,
				product_fees.is_taxable,
				product_fees.create_time,
				product_fees.update_time
			FROM
				product_fees
			WHERE
				%s
		), all_count AS (
			SELECT
				COUNT(*) AS count
			FROM
				all_entries
		), filtered AS (
			SELECT
				*
			FROM
				all_entries
			WHERE %s
		), filtered_count AS (
			SELECT
				COUNT(*) AS count
			FROM
				filtered
		)
		SELECT
			all_count.count, filtered_count.count,
			NULL, NULL,
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
		FROM
			all_count
			CROSS JOIN filtered_count
		UNION ALL
		(
			SELECT
				*
			FROM
				all_count
				CROSS JOIN filtered_count
				CROSS JOIN filtered
			ORDER BY
				quote_ident($1), filtered.id
			LIMIT
				$4
		)`

	filterSQL := "TRUE" // TODO
	query := fmt.Sprintf(stmt, filterSQL, whereClause)
	rows, err := repo.db.QueryContext(ctx, query, "product_fees."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.ProductFee, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.ProductFee
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.FeeID, &next.Name, &next.CalculationMethod, &next.Required, &next.Amount, &next.Percent, &next.ApplyDateMethod, &next.Trigger, &next.AccountingRules, &next.AmortizationSettings, &next.IsActive, &next.IsTaxable, &next.CreateTime, &next.UpdateTime,
		); err != nil {
			return db.ListStat{}, nil, nil, err
		}
		result = append(result, &next)
	}
	if rows.Err() != nil {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	return listStat, result, &nextPos, nil
}

func (repo *Repository) ProductFeeByID(ctx context.Context, id int32) (*db.ProductFee, error) {
	const stmt = `SELECT
		product_fees.id,

		product_fees.fee_id,
		product_fees.name,
		product_fees.calculation_method,
		product_fees.required,
		product_fees.amount,
		product_fees.percent,
		product_fees.apply_date_method,
		product_fees.trigger,
		product_fees.accounting_rules,
		product_fees.amortization_settings,
		product_fees.is_active,
		product_fees.is_taxable,
		product_fees.create_time,
		product_fees.update_time
	FROM
		product_fees
	WHERE
		id = $1`

	var pf db.ProductFee
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&pf.ID,
		&pf.FeeID, &pf.Name, &pf.CalculationMethod, &pf.Required, &pf.Amount, &pf.Percent, &pf.ApplyDateMethod, &pf.Trigger, &pf.AccountingRules, &pf.AmortizationSettings, &pf.IsActive, &pf.IsTaxable, &pf.CreateTime, &pf.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &pf, nil
}

func (repo *Repository) UpdateProductFeeByID(ctx context.Context, pf *db.ProductFee) error {
	const stmt = `UPDATE product_fees
	SET fee_id=$1,
		name=$2,
		calculation_method=$3,
		required=$4,
		amount=$5,
		percent=$6,
		apply_date_method=$7,
		trigger=$8,
		accounting_rules=$9,
		amortization_settings=$10,
		is_active=$11,
		is_taxable=$12,
		update_time=$13
	WHERE id = $14`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		pf.FeeID, pf.Name, pf.CalculationMethod, pf.Required, pf.Amount, pf.Percent, pf.ApplyDateMethod, pf.Trigger, pf.AccountingRules, pf.AmortizationSettings, pf.IsActive, pf.IsTaxable, pf.UpdateTime, pf.ID,
	)
	return err
}

func (repo *Repository) DeleteProductFeeByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM product_fees
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) ProductFeeByFeeID(ctx context.Context, feeID string) (*db.ProductFee, error) {
	const stmt = `SELECT
		product_fees.id,

		product_fees.fee_id,
		product_fees.name,
		product_fees.calculation_method,
		product_fees.required,
		product_fees.amount,
		product_fees.percent,
		product_fees.apply_date_method,
		product_fees.trigger,
		product_fees.accounting_rules,
		product_fees.amortization_settings,
		product_fees.is_active,
		product_fees.is_taxable,
		product_fees.create_time,
		product_fees.update_time
	FROM
		product_fees
	WHERE
		fee_id = $1`

	var pf db.ProductFee
	row := repo.db.QueryRowContext(ctx, stmt, feeID)
	if err := row.Scan(
		&pf.ID,
		&pf.FeeID, &pf.Name, &pf.CalculationMethod, &pf.Required, &pf.Amount, &pf.Percent, &pf.ApplyDateMethod, &pf.Trigger, &pf.AccountingRules, &pf.AmortizationSettings, &pf.IsActive, &pf.IsTaxable, &pf.CreateTime, &pf.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &pf, nil
}

func (repo *Repository) UpdateProductFeeByFeeID(ctx context.Context, pf *db.ProductFee) error {
	const stmt = `UPDATE product_fees
	SET fee_id=$1,
		name=$2,
		calculation_method=$3,
		required=$4,
		amount=$5,
		percent=$6,
		apply_date_method=$7,
		trigger=$8,
		accounting_rules=$9,
		amortization_settings=$10,
		is_active=$11,
		is_taxable=$12,
		update_time=$13
	WHERE fee_id = $14`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		pf.FeeID, pf.Name, pf.CalculationMethod, pf.Required, pf.Amount, pf.Percent, pf.ApplyDateMethod, pf.Trigger, pf.AccountingRules, pf.AmortizationSettings, pf.IsActive, pf.IsTaxable, pf.UpdateTime, pf.FeeID,
	)
	return err
}

func (repo *Repository) DeleteProductFeeByFeeID(ctx context.Context, feeID string) error {
	const stmt = `DELETE FROM product_fees
	WHERE fee_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		feeID,
	)
	return err
}
