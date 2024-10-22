// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertTask(ctx context.Context, t *db.Task) error {
	const stmt = `INSERT INTO tasks
		(task_id, title, status, description, assignee_id, assigner_full_name, assigner_id, linked_id, linked_type, template_id, due_time, create_time, update_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		t.TaskID, t.Title, t.Status, t.Description, t.AssigneeID, t.AssignerFullName, t.AssignerID, t.LinkedID, t.LinkedType, t.TemplateID, t.DueTime, t.CreateTime, t.UpdateTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListTasks(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.Task, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, tasks.id,

				tasks.task_id,
				tasks.title,
				tasks.status,
				tasks.description,
				tasks.assignee_id,
				tasks.assigner_full_name,
				tasks.assigner_id,
				tasks.linked_id,
				tasks.linked_type,
				tasks.template_id,
				tasks.due_time,
				tasks.create_time,
				tasks.update_time
			FROM
				tasks
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "tasks."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.Task, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.Task
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.TaskID, &next.Title, &next.Status, &next.Description, &next.AssigneeID, &next.AssignerFullName, &next.AssignerID, &next.LinkedID, &next.LinkedType, &next.TemplateID, &next.DueTime, &next.CreateTime, &next.UpdateTime,
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

func (repo *Repository) TaskByID(ctx context.Context, id int32) (*db.Task, error) {
	const stmt = `SELECT
		tasks.id,

		tasks.task_id,
		tasks.title,
		tasks.status,
		tasks.description,
		tasks.assignee_id,
		tasks.assigner_full_name,
		tasks.assigner_id,
		tasks.linked_id,
		tasks.linked_type,
		tasks.template_id,
		tasks.due_time,
		tasks.create_time,
		tasks.update_time
	FROM
		tasks
	WHERE
		id = $1`

	var t db.Task
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&t.ID,
		&t.TaskID, &t.Title, &t.Status, &t.Description, &t.AssigneeID, &t.AssignerFullName, &t.AssignerID, &t.LinkedID, &t.LinkedType, &t.TemplateID, &t.DueTime, &t.CreateTime, &t.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &t, nil
}

func (repo *Repository) UpdateTaskByID(ctx context.Context, t *db.Task) error {
	const stmt = `UPDATE tasks
	SET task_id=$1,
		title=$2,
		status=$3,
		description=$4,
		assignee_id=$5,
		assigner_full_name=$6,
		assigner_id=$7,
		linked_id=$8,
		linked_type=$9,
		template_id=$10,
		due_time=$11,
		update_time=$12
	WHERE id = $13`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		t.TaskID, t.Title, t.Status, t.Description, t.AssigneeID, t.AssignerFullName, t.AssignerID, t.LinkedID, t.LinkedType, t.TemplateID, t.DueTime, t.UpdateTime, t.ID,
	)
	return err
}

func (repo *Repository) DeleteTaskByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM tasks
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) TaskByTaskID(ctx context.Context, taskID string) (*db.Task, error) {
	const stmt = `SELECT
		tasks.id,

		tasks.task_id,
		tasks.title,
		tasks.status,
		tasks.description,
		tasks.assignee_id,
		tasks.assigner_full_name,
		tasks.assigner_id,
		tasks.linked_id,
		tasks.linked_type,
		tasks.template_id,
		tasks.due_time,
		tasks.create_time,
		tasks.update_time
	FROM
		tasks
	WHERE
		task_id = $1`

	var t db.Task
	row := repo.db.QueryRowContext(ctx, stmt, taskID)
	if err := row.Scan(
		&t.ID,
		&t.TaskID, &t.Title, &t.Status, &t.Description, &t.AssigneeID, &t.AssignerFullName, &t.AssignerID, &t.LinkedID, &t.LinkedType, &t.TemplateID, &t.DueTime, &t.CreateTime, &t.UpdateTime,
	); err != nil {
		return nil, err
	}

	return &t, nil
}

func (repo *Repository) UpdateTaskByTaskID(ctx context.Context, t *db.Task) error {
	const stmt = `UPDATE tasks
	SET task_id=$1,
		title=$2,
		status=$3,
		description=$4,
		assignee_id=$5,
		assigner_full_name=$6,
		assigner_id=$7,
		linked_id=$8,
		linked_type=$9,
		template_id=$10,
		due_time=$11,
		update_time=$12
	WHERE task_id = $13`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		t.TaskID, t.Title, t.Status, t.Description, t.AssigneeID, t.AssignerFullName, t.AssignerID, t.LinkedID, t.LinkedType, t.TemplateID, t.DueTime, t.UpdateTime, t.TaskID,
	)
	return err
}

func (repo *Repository) DeleteTaskByTaskID(ctx context.Context, taskID string) error {
	const stmt = `DELETE FROM tasks
	WHERE task_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		taskID,
	)
	return err
}
