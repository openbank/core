// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertOrgHoliday(ctx context.Context, oh *db.OrgHoliday) error {
	const stmt = `INSERT INTO org_holidays
		(holiday_id, name, date, is_annually_recurring, create_time)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		oh.HolidayID, oh.Name, oh.Date, oh.IsAnnuallyRecurring, oh.CreateTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListOrgHolidays(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.OrgHoliday, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, org_holidays.id,

				org_holidays.holiday_id,
				org_holidays.name,
				org_holidays.date,
				org_holidays.is_annually_recurring,
				org_holidays.create_time
			FROM
				org_holidays
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
			NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "org_holidays."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.OrgHoliday, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.OrgHoliday
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.HolidayID, &next.Name, &next.Date, &next.IsAnnuallyRecurring, &next.CreateTime,
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

func (repo *Repository) OrgHolidayByID(ctx context.Context, id int32) (*db.OrgHoliday, error) {
	const stmt = `SELECT
		org_holidays.id,

		org_holidays.holiday_id,
		org_holidays.name,
		org_holidays.date,
		org_holidays.is_annually_recurring,
		org_holidays.create_time
	FROM
		org_holidays
	WHERE
		id = $1`

	var oh db.OrgHoliday
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&oh.ID,
		&oh.HolidayID, &oh.Name, &oh.Date, &oh.IsAnnuallyRecurring, &oh.CreateTime,
	); err != nil {
		return nil, err
	}

	return &oh, nil
}

func (repo *Repository) UpdateOrgHolidayByID(ctx context.Context, oh *db.OrgHoliday) error {
	const stmt = `UPDATE org_holidays
	SET holiday_id=$1,
		name=$2,
		date=$3,
		is_annually_recurring=$4
	WHERE id = $5`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		oh.HolidayID, oh.Name, oh.Date, oh.IsAnnuallyRecurring, oh.ID,
	)
	return err
}

func (repo *Repository) DeleteOrgHolidayByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM org_holidays
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) OrgHolidayByHolidayID(ctx context.Context, holidayID string) (*db.OrgHoliday, error) {
	const stmt = `SELECT
		org_holidays.id,

		org_holidays.holiday_id,
		org_holidays.name,
		org_holidays.date,
		org_holidays.is_annually_recurring,
		org_holidays.create_time
	FROM
		org_holidays
	WHERE
		holiday_id = $1`

	var oh db.OrgHoliday
	row := repo.db.QueryRowContext(ctx, stmt, holidayID)
	if err := row.Scan(
		&oh.ID,
		&oh.HolidayID, &oh.Name, &oh.Date, &oh.IsAnnuallyRecurring, &oh.CreateTime,
	); err != nil {
		return nil, err
	}

	return &oh, nil
}

func (repo *Repository) UpdateOrgHolidayByHolidayID(ctx context.Context, oh *db.OrgHoliday) error {
	const stmt = `UPDATE org_holidays
	SET holiday_id=$1,
		name=$2,
		date=$3,
		is_annually_recurring=$4
	WHERE holiday_id = $5`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		oh.HolidayID, oh.Name, oh.Date, oh.IsAnnuallyRecurring, oh.HolidayID,
	)
	return err
}

func (repo *Repository) DeleteOrgHolidayByHolidayID(ctx context.Context, holidayID string) error {
	const stmt = `DELETE FROM org_holidays
	WHERE holiday_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		holidayID,
	)
	return err
}
