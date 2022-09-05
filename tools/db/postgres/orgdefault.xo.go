// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertOrgDefault(ctx context.Context, od *db.OrgDefault) error {
	const stmt = `INSERT INTO org_defaults
		(client_role_id, client_state, group_role_id, line_of_credit_state, transaction_channel_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		od.ClientRoleID, od.ClientState, od.GroupRoleID, od.LineOfCreditState, od.TransactionChannelID,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListOrgDefaults(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.OrgDefault, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, org_defaults.id,

				org_defaults.client_role_id,
				org_defaults.client_state,
				org_defaults.group_role_id,
				org_defaults.line_of_credit_state,
				org_defaults.transaction_channel_id
			FROM
				org_defaults
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
	rows, err := repo.db.QueryContext(ctx, query, "org_defaults."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.OrgDefault, 0, pageSize)
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
		var next db.OrgDefault
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.ClientRoleID, &next.ClientState, &next.GroupRoleID, &next.LineOfCreditState, &next.TransactionChannelID,
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

func (repo *Repository) OrgDefaultByID(ctx context.Context, id int32) (*db.OrgDefault, error) {
	const stmt = `SELECT
		org_defaults.id,

		org_defaults.client_role_id,
		org_defaults.client_state,
		org_defaults.group_role_id,
		org_defaults.line_of_credit_state,
		org_defaults.transaction_channel_id
	FROM
		org_defaults
	WHERE
		id = $1`

	var od db.OrgDefault
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&od.ID,
		&od.ClientRoleID, &od.ClientState, &od.GroupRoleID, &od.LineOfCreditState, &od.TransactionChannelID,
	); err != nil {
		return nil, err
	}

	return &od, nil
}

func (repo *Repository) UpdateOrgDefaultByID(ctx context.Context, od *db.OrgDefault) error {
	const stmt = `UPDATE org_defaults
	SET client_role_id=$1,
		client_state=$2,
		group_role_id=$3,
		line_of_credit_state=$4,
		transaction_channel_id=$5
	WHERE id = $6`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		od.ClientRoleID, od.ClientState, od.GroupRoleID, od.LineOfCreditState, od.TransactionChannelID, od.ID,
	)
	return err
}

func (repo *Repository) DeleteOrgDefaultByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM org_defaults
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}