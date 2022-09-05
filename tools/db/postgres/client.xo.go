// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertClient(ctx context.Context, c *db.Client) error {
	const stmt = `WITH clients_info AS (
		INSERT INTO contact_infos
			(full_name, addresses, telephones, emails, language)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
	)
	INSERT INTO clients
		(client_id, info, assigned_branch_id, assigned_centre_id, assigned_user_id, birthday, role_id, gender, group_loan_cycle, loan_cycle, portal_settings, migration_event_id, profile_picture_id, profile_signature_id, state, notes, create_time, update_time, approve_time, activate_time, close_time)
		SELECT $6, clients_info.id, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25
		FROM clients_info
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		c.Info.FullName, c.Info.Addresses, c.Info.Telephones, c.Info.Emails, c.Info.Language,
		c.ClientID, c.AssignedBranchID, c.AssignedCentreID, c.AssignedUserID, c.Birthday, c.RoleID, c.Gender, c.GroupLoanCycle, c.LoanCycle, c.PortalSettings, c.MigrationEventID, c.ProfilePictureID, c.ProfileSignatureID, c.State, c.Notes, c.CreateTime, c.UpdateTime, c.ApproveTime, c.ActivateTime, c.CloseTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListClients(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.Client, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, clients.id,

				clients_info_tbl.full_name,
				clients_info_tbl.addresses,
				clients_info_tbl.telephones,
				clients_info_tbl.emails,
				clients_info_tbl.language,

				clients.client_id,
				clients.assigned_branch_id,
				clients.assigned_centre_id,
				clients.assigned_user_id,
				clients.birthday,
				clients.role_id,
				clients.gender,
				clients.group_loan_cycle,
				clients.loan_cycle,
				clients.portal_settings,
				clients.migration_event_id,
				clients.profile_picture_id,
				clients.profile_signature_id,
				clients.state,
				clients.notes,
				clients.create_time,
				clients.update_time,
				clients.approve_time,
				clients.activate_time,
				clients.close_time
			FROM
				clients
				JOIN contact_infos AS clients_info_tbl ON clients_info_tbl.id = clients.info
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
			NULL, NULL, NULL, NULL, NULL,
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
	rows, err := repo.db.QueryContext(ctx, query, "clients."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.Client, 0, pageSize)
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
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.Client
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.Info.FullName, &next.Info.Addresses, &next.Info.Telephones, &next.Info.Emails, &next.Info.Language,
			&next.ClientID, &next.AssignedBranchID, &next.AssignedCentreID, &next.AssignedUserID, &next.Birthday, &next.RoleID, &next.Gender, &next.GroupLoanCycle, &next.LoanCycle, &next.PortalSettings, &next.MigrationEventID, &next.ProfilePictureID, &next.ProfileSignatureID, &next.State, &next.Notes, &next.CreateTime, &next.UpdateTime, &next.ApproveTime, &next.ActivateTime, &next.CloseTime,
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

func (repo *Repository) ClientByID(ctx context.Context, id int32) (*db.Client, error) {
	const stmt = `SELECT
		clients.id,

		clients_info_tbl.full_name,
		clients_info_tbl.addresses,
		clients_info_tbl.telephones,
		clients_info_tbl.emails,
		clients_info_tbl.language,

		clients.client_id,
		clients.assigned_branch_id,
		clients.assigned_centre_id,
		clients.assigned_user_id,
		clients.birthday,
		clients.role_id,
		clients.gender,
		clients.group_loan_cycle,
		clients.loan_cycle,
		clients.portal_settings,
		clients.migration_event_id,
		clients.profile_picture_id,
		clients.profile_signature_id,
		clients.state,
		clients.notes,
		clients.create_time,
		clients.update_time,
		clients.approve_time,
		clients.activate_time,
		clients.close_time
	FROM
		clients
		JOIN contact_infos AS clients_info_tbl ON clients_info_tbl.id = clients.info
	WHERE
		id = $1`

	var c db.Client
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&c.ID,
		&c.Info.FullName, &c.Info.Addresses, &c.Info.Telephones, &c.Info.Emails, &c.Info.Language,
		&c.ClientID, &c.AssignedBranchID, &c.AssignedCentreID, &c.AssignedUserID, &c.Birthday, &c.RoleID, &c.Gender, &c.GroupLoanCycle, &c.LoanCycle, &c.PortalSettings, &c.MigrationEventID, &c.ProfilePictureID, &c.ProfileSignatureID, &c.State, &c.Notes, &c.CreateTime, &c.UpdateTime, &c.ApproveTime, &c.ActivateTime, &c.CloseTime,
	); err != nil {
		return nil, err
	}

	return &c, nil
}

func (repo *Repository) UpdateClientByID(ctx context.Context, c *db.Client) error {
	const stmt = `WITH clients_contact_infos AS (
		UPDATE contact_infos
		SET full_name=$1,
			addresses=$2,
			telephones=$3,
			emails=$4,
			language=$5
		FROM clients
		WHERE clients.info = contact_infos.id AND
			clients.id = $6
	)
	UPDATE clients
	SET client_id=$7,
		assigned_branch_id=$8,
		assigned_centre_id=$9,
		assigned_user_id=$10,
		birthday=$11,
		role_id=$12,
		gender=$13,
		group_loan_cycle=$14,
		loan_cycle=$15,
		portal_settings=$16,
		migration_event_id=$17,
		profile_picture_id=$18,
		profile_signature_id=$19,
		state=$20,
		notes=$21,
		update_time=$22,
		approve_time=$23,
		activate_time=$24,
		close_time=$25
	WHERE id = $26`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		c.Info.FullName, c.Info.Addresses, c.Info.Telephones, c.Info.Emails, c.Info.Language, c.ID,
		c.ClientID, c.AssignedBranchID, c.AssignedCentreID, c.AssignedUserID, c.Birthday, c.RoleID, c.Gender, c.GroupLoanCycle, c.LoanCycle, c.PortalSettings, c.MigrationEventID, c.ProfilePictureID, c.ProfileSignatureID, c.State, c.Notes, c.UpdateTime, c.ApproveTime, c.ActivateTime, c.CloseTime, c.ID,
	)
	return err
}

func (repo *Repository) DeleteClientByID(ctx context.Context, id int32) error {
	const stmt = `WITH clients_contact_infos AS (
		DELETE FROM contact_infos
		USING clients
		WHERE clients.info = contact_infos.id AND
			clients.id = $1
	)
	DELETE FROM clients
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) ClientByClientID(ctx context.Context, clientID string) (*db.Client, error) {
	const stmt = `SELECT
		clients.id,

		clients_info_tbl.full_name,
		clients_info_tbl.addresses,
		clients_info_tbl.telephones,
		clients_info_tbl.emails,
		clients_info_tbl.language,

		clients.client_id,
		clients.assigned_branch_id,
		clients.assigned_centre_id,
		clients.assigned_user_id,
		clients.birthday,
		clients.role_id,
		clients.gender,
		clients.group_loan_cycle,
		clients.loan_cycle,
		clients.portal_settings,
		clients.migration_event_id,
		clients.profile_picture_id,
		clients.profile_signature_id,
		clients.state,
		clients.notes,
		clients.create_time,
		clients.update_time,
		clients.approve_time,
		clients.activate_time,
		clients.close_time
	FROM
		clients
		JOIN contact_infos AS clients_info_tbl ON clients_info_tbl.id = clients.info
	WHERE
		client_id = $1`

	var c db.Client
	row := repo.db.QueryRowContext(ctx, stmt, clientID)
	if err := row.Scan(
		&c.ID,
		&c.Info.FullName, &c.Info.Addresses, &c.Info.Telephones, &c.Info.Emails, &c.Info.Language,
		&c.ClientID, &c.AssignedBranchID, &c.AssignedCentreID, &c.AssignedUserID, &c.Birthday, &c.RoleID, &c.Gender, &c.GroupLoanCycle, &c.LoanCycle, &c.PortalSettings, &c.MigrationEventID, &c.ProfilePictureID, &c.ProfileSignatureID, &c.State, &c.Notes, &c.CreateTime, &c.UpdateTime, &c.ApproveTime, &c.ActivateTime, &c.CloseTime,
	); err != nil {
		return nil, err
	}

	return &c, nil
}

func (repo *Repository) UpdateClientByClientID(ctx context.Context, c *db.Client) error {
	const stmt = `WITH clients_contact_infos AS (
		UPDATE contact_infos
		SET full_name=$1,
			addresses=$2,
			telephones=$3,
			emails=$4,
			language=$5
		FROM clients
		WHERE clients.info = contact_infos.id AND
			clients.client_id = $6
	)
	UPDATE clients
	SET client_id=$7,
		assigned_branch_id=$8,
		assigned_centre_id=$9,
		assigned_user_id=$10,
		birthday=$11,
		role_id=$12,
		gender=$13,
		group_loan_cycle=$14,
		loan_cycle=$15,
		portal_settings=$16,
		migration_event_id=$17,
		profile_picture_id=$18,
		profile_signature_id=$19,
		state=$20,
		notes=$21,
		update_time=$22,
		approve_time=$23,
		activate_time=$24,
		close_time=$25
	WHERE client_id = $26`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		c.Info.FullName, c.Info.Addresses, c.Info.Telephones, c.Info.Emails, c.Info.Language, c.ClientID,
		c.ClientID, c.AssignedBranchID, c.AssignedCentreID, c.AssignedUserID, c.Birthday, c.RoleID, c.Gender, c.GroupLoanCycle, c.LoanCycle, c.PortalSettings, c.MigrationEventID, c.ProfilePictureID, c.ProfileSignatureID, c.State, c.Notes, c.UpdateTime, c.ApproveTime, c.ActivateTime, c.CloseTime, c.ClientID,
	)
	return err
}

func (repo *Repository) DeleteClientByClientID(ctx context.Context, clientID string) error {
	const stmt = `WITH clients_contact_infos AS (
		DELETE FROM contact_infos
		USING clients
		WHERE clients.info = contact_infos.id AND
			clients.client_id = $1
	)
	DELETE FROM clients
	WHERE client_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		clientID,
	)
	return err
}