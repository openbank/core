{{ define "typedef" }}
{{- $typedef := .Data }}
{{- $t := index .Data.Tables .Data.Table -}}

func (repo *Repository) Insert{{ $t.GoName }}(ctx context.Context, {{ short $t }} *db.{{ $t.GoName }}) error {
	{{- $var := short $t }}
	const stmt = `{{ insert $typedef }}`
	var id int
	row := repo.db.QueryRowContext(
		ctx, stmt,
		{{- range insert_data $var $typedef }}
		{{ . }},
		{{- end }}
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) List{{ plural $t.GoName }}(ctx context.Context, filter string, pageSize int32, orderBy string, after *db.ListPosition) (db.ListStat, []*db.{{ $t.GoName }}, *db.ListPosition, error) {
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
			ID: 0,
		}
	}
	const stmt = `{{ list $typedef }}`

	filterSQL := "TRUE" // TODO
	query := fmt.Sprintf(stmt, filterSQL, whereClause)
	rows, err := repo.db.QueryContext(ctx, query, "{{ $t.SQLName }}." + orderBy, after.Data, after.ID, pageSize)
	if err != nil {
		return db.ListStat{}, nil, nil, err
	}

	defer rows.Close()
	result := make([]*db.{{ $t.GoName }}, 0, pageSize)
	var nextPos db.ListPosition
	var listStat db.ListStat
	if !rows.Next() {
		return db.ListStat{}, nil, nil, rows.Err()
	}
	var x any
	if err := rows.Scan(
		&listStat.Total, &listStat.Remaining,
		&x, &x,
		{{- range select_data_literal "&x" $typedef }}
		{{ . }},
		{{- end }}
	); err != nil {
		return db.ListStat{}, nil, nil, err
	}
	for rows.Next() {
		var next db.{{ $t.GoName }}
		if err := rows.Scan(
			&listStat.Total, &listStat.Remaining,
			&nextPos.Data, &nextPos.ID,
			{{- range select_data "next" $typedef }}
			{{ . }},
			{{- end }}
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

{{ range $i := $t.Indexes }}
func (repo *Repository) {{ $i.Func }}(ctx context.Context, {{ typed_params $i.Fields }}) ({{ if not $i.IsUnique }}[]{{ end }}*db.{{ $t.GoName }}, error) {
	const stmt = `{{ select_idx $typedef $i }}`
	{{ if $i.IsUnique }}
	var {{ short $t }} db.{{ $t.GoName }}
	row := repo.db.QueryRowContext(ctx, stmt, {{ untyped_params $i.Fields }})
	if err := row.Scan(
		&{{ short $t }}.ID,
		{{- range select_data (short $t) $typedef }}
		{{ . }},
		{{- end }}
	); err != nil {
		return nil, err
	}

	return &{{ short $t }}, nil
	{{- else }}
	rows, err := repo.db.QueryContext(ctx, stmt, {{ untyped_params $i.Fields }})
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var result []*db.{{ $t.GoName }}
	for rows.Next() {
		var next db.{{ $t.GoName }}
		if err := rows.Scan(
			&next.ID,
			{{- range select_data "next" $typedef }}
			{{ . }},
			{{- end }}
		); err != nil {
			return nil, err
		}
		result = append(result, &next)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
	{{- end }}
}

{{- if $i.IsUnique }}
{{- $fields := $i.Fields }}
{{- $by := join $fields "" }}

func (repo *Repository) Update{{ $t.GoName }}By{{ $by }}(ctx context.Context, {{ short $t }} *db.{{ $t.GoName }}) error {
	{{- $var := short $t }}
	const stmt = `{{ update $typedef $fields }}`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		{{- range update_data $var $typedef $fields }}
		{{ . }},
		{{- end }}
	)
	return err
}

func (repo *Repository) Delete{{ $t.GoName }}By{{ $by }}(ctx context.Context, {{ typed_params $fields }}) error {
	{{- $var := short $t }}
	const stmt = `{{ delete $typedef $fields }}`
	_, err := repo.db.ExecContext(
		ctx, stmt,
		{{ range $f := $fields }}{{ param_name $f }}, {{ end }}
	)
	return err
}
{{- end }}
{{- end }}
{{- end }}
