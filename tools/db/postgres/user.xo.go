// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertUser(ctx context.Context, u *db.User) error {
	const stmt = `INSERT INTO users
		(user_id, username, password, email, title, first_name, last_name, role_id, access, assigned_branch_id, mobile_phone, home_phone, language, notes, transaction_limits, two_factor_authentication, state, create_time, update_time, login_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		u.UserID, u.Username, u.Password, u.Email, u.Title, u.FirstName, u.LastName, u.RoleID, u.Access, u.AssignedBranchID, u.MobilePhone, u.HomePhone, u.Language, u.Notes, u.TransactionLimits, u.TwoFactorAuthentication, u.State, u.CreateTime, u.UpdateTime, u.LoginTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListUsers(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.User, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, users.id,

				users.user_id,
				users.username,
				users.password,
				users.email,
				users.title,
				users.first_name,
				users.last_name,
				users.role_id,
				users.access,
				users.assigned_branch_id,
				users.mobile_phone,
				users.home_phone,
				users.language,
				users.notes,
				users.transaction_limits,
				users.two_factor_authentication,
				users.state,
				users.create_time,
				users.update_time,
				users.login_time
			FROM
				users
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "users."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.User, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.User
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.UserID, &next.Username, &next.Password, &next.Email, &next.Title, &next.FirstName, &next.LastName, &next.RoleID, &next.Access, &next.AssignedBranchID, &next.MobilePhone, &next.HomePhone, &next.Language, &next.Notes, &next.TransactionLimits, &next.TwoFactorAuthentication, &next.State, &next.CreateTime, &next.UpdateTime, &next.LoginTime,
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

func (repo *Repository) UserByID(ctx context.Context, id int32) (*db.User, error) {
	const stmt = `SELECT
		users.id,

		users.user_id,
		users.username,
		users.password,
		users.email,
		users.title,
		users.first_name,
		users.last_name,
		users.role_id,
		users.access,
		users.assigned_branch_id,
		users.mobile_phone,
		users.home_phone,
		users.language,
		users.notes,
		users.transaction_limits,
		users.two_factor_authentication,
		users.state,
		users.create_time,
		users.update_time,
		users.login_time
	FROM
		users
	WHERE
		id = $1`

	var u db.User
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&u.ID,
		&u.UserID, &u.Username, &u.Password, &u.Email, &u.Title, &u.FirstName, &u.LastName, &u.RoleID, &u.Access, &u.AssignedBranchID, &u.MobilePhone, &u.HomePhone, &u.Language, &u.Notes, &u.TransactionLimits, &u.TwoFactorAuthentication, &u.State, &u.CreateTime, &u.UpdateTime, &u.LoginTime,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (repo *Repository) UpdateUserByID(ctx context.Context, u *db.User) error {
	const stmt = `UPDATE users
	SET user_id=$1,
		username=$2,
		password=$3,
		email=$4,
		title=$5,
		first_name=$6,
		last_name=$7,
		role_id=$8,
		access=$9,
		assigned_branch_id=$10,
		mobile_phone=$11,
		home_phone=$12,
		language=$13,
		notes=$14,
		transaction_limits=$15,
		two_factor_authentication=$16,
		state=$17,
		update_time=$18,
		login_time=$19
	WHERE id = $20`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		u.UserID, u.Username, u.Password, u.Email, u.Title, u.FirstName, u.LastName, u.RoleID, u.Access, u.AssignedBranchID, u.MobilePhone, u.HomePhone, u.Language, u.Notes, u.TransactionLimits, u.TwoFactorAuthentication, u.State, u.UpdateTime, u.LoginTime, u.ID,
	)
	return err
}

func (repo *Repository) DeleteUserByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM users
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) UserByUserID(ctx context.Context, userID string) (*db.User, error) {
	const stmt = `SELECT
		users.id,

		users.user_id,
		users.username,
		users.password,
		users.email,
		users.title,
		users.first_name,
		users.last_name,
		users.role_id,
		users.access,
		users.assigned_branch_id,
		users.mobile_phone,
		users.home_phone,
		users.language,
		users.notes,
		users.transaction_limits,
		users.two_factor_authentication,
		users.state,
		users.create_time,
		users.update_time,
		users.login_time
	FROM
		users
	WHERE
		user_id = $1`

	var u db.User
	row := repo.db.QueryRowContext(ctx, stmt, userID)
	if err := row.Scan(
		&u.ID,
		&u.UserID, &u.Username, &u.Password, &u.Email, &u.Title, &u.FirstName, &u.LastName, &u.RoleID, &u.Access, &u.AssignedBranchID, &u.MobilePhone, &u.HomePhone, &u.Language, &u.Notes, &u.TransactionLimits, &u.TwoFactorAuthentication, &u.State, &u.CreateTime, &u.UpdateTime, &u.LoginTime,
	); err != nil {
		return nil, err
	}

	return &u, nil
}

func (repo *Repository) UpdateUserByUserID(ctx context.Context, u *db.User) error {
	const stmt = `UPDATE users
	SET user_id=$1,
		username=$2,
		password=$3,
		email=$4,
		title=$5,
		first_name=$6,
		last_name=$7,
		role_id=$8,
		access=$9,
		assigned_branch_id=$10,
		mobile_phone=$11,
		home_phone=$12,
		language=$13,
		notes=$14,
		transaction_limits=$15,
		two_factor_authentication=$16,
		state=$17,
		update_time=$18,
		login_time=$19
	WHERE user_id = $20`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		u.UserID, u.Username, u.Password, u.Email, u.Title, u.FirstName, u.LastName, u.RoleID, u.Access, u.AssignedBranchID, u.MobilePhone, u.HomePhone, u.Language, u.Notes, u.TransactionLimits, u.TwoFactorAuthentication, u.State, u.UpdateTime, u.LoginTime, u.UserID,
	)
	return err
}

func (repo *Repository) DeleteUserByUserID(ctx context.Context, userID string) error {
	const stmt = `DELETE FROM users
	WHERE user_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		userID,
	)
	return err
}