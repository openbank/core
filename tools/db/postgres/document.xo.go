// Package postgres contains generated code from xo.
package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"bnk.to/core/tools/db"
)

func (repo *Repository) InsertDocument(ctx context.Context, d *db.Document) error {
	const stmt = `INSERT INTO documents
		(document_id, name, notes, owner_id, owner_type, template_id, file_name, size, create_time, expire_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		d.DocumentID, d.Name, d.Notes, d.OwnerID, d.OwnerType, d.TemplateID, d.FileName, d.Size, d.CreateTime, d.ExpireTime,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) ListDocuments(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.Document, *db.ListPosition, error) {
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
				quote_ident($1) AS ordered_idx, documents.id,

				documents.document_id,
				documents.name,
				documents.notes,
				documents.owner_id,
				documents.owner_type,
				documents.template_id,
				documents.file_name,
				documents.size,
				documents.create_time,
				documents.expire_time
			FROM
				documents
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
			NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL
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
	rows, err := repo.db.QueryContext(ctx, query, "documents."+orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.Document, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		&x, &x, &x, &x, &x, &x, &x, &x, &x, &x,
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.Document
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			&next.DocumentID, &next.Name, &next.Notes, &next.OwnerID, &next.OwnerType, &next.TemplateID, &next.FileName, &next.Size, &next.CreateTime, &next.ExpireTime,
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

func (repo *Repository) DocumentByID(ctx context.Context, id int32) (*db.Document, error) {
	const stmt = `SELECT
		documents.id,

		documents.document_id,
		documents.name,
		documents.notes,
		documents.owner_id,
		documents.owner_type,
		documents.template_id,
		documents.file_name,
		documents.size,
		documents.create_time,
		documents.expire_time
	FROM
		documents
	WHERE
		id = $1`

	var d db.Document
	row := repo.db.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(
		&d.ID,
		&d.DocumentID, &d.Name, &d.Notes, &d.OwnerID, &d.OwnerType, &d.TemplateID, &d.FileName, &d.Size, &d.CreateTime, &d.ExpireTime,
	); err != nil {
		return nil, err
	}

	return &d, nil
}

func (repo *Repository) UpdateDocumentByID(ctx context.Context, d *db.Document) error {
	const stmt = `UPDATE documents
	SET document_id=$1,
		name=$2,
		notes=$3,
		owner_id=$4,
		owner_type=$5,
		template_id=$6,
		file_name=$7,
		size=$8,
		expire_time=$9
	WHERE id = $10`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		d.DocumentID, d.Name, d.Notes, d.OwnerID, d.OwnerType, d.TemplateID, d.FileName, d.Size, d.ExpireTime, d.ID,
	)
	return err
}

func (repo *Repository) DeleteDocumentByID(ctx context.Context, id int32) error {
	const stmt = `DELETE FROM documents
	WHERE id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		id,
	)
	return err
}

func (repo *Repository) DocumentByDocumentID(ctx context.Context, documentID string) (*db.Document, error) {
	const stmt = `SELECT
		documents.id,

		documents.document_id,
		documents.name,
		documents.notes,
		documents.owner_id,
		documents.owner_type,
		documents.template_id,
		documents.file_name,
		documents.size,
		documents.create_time,
		documents.expire_time
	FROM
		documents
	WHERE
		document_id = $1`

	var d db.Document
	row := repo.db.QueryRowContext(ctx, stmt, documentID)
	if err := row.Scan(
		&d.ID,
		&d.DocumentID, &d.Name, &d.Notes, &d.OwnerID, &d.OwnerType, &d.TemplateID, &d.FileName, &d.Size, &d.CreateTime, &d.ExpireTime,
	); err != nil {
		return nil, err
	}

	return &d, nil
}

func (repo *Repository) UpdateDocumentByDocumentID(ctx context.Context, d *db.Document) error {
	const stmt = `UPDATE documents
	SET document_id=$1,
		name=$2,
		notes=$3,
		owner_id=$4,
		owner_type=$5,
		template_id=$6,
		file_name=$7,
		size=$8,
		expire_time=$9
	WHERE document_id = $10`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		d.DocumentID, d.Name, d.Notes, d.OwnerID, d.OwnerType, d.TemplateID, d.FileName, d.Size, d.ExpireTime, d.DocumentID,
	)
	return err
}

func (repo *Repository) DeleteDocumentByDocumentID(ctx context.Context, documentID string) error {
	const stmt = `DELETE FROM documents
	WHERE document_id = $1`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		documentID,
	)
	return err
}
