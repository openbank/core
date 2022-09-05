package impl

import (
	"fmt"
	"strings"
)

// insert generates the SQL statement for an INSERT operation of the table
// specified in the typedef.
func (f *Funcs) insert(td Typedef) (string, error) {
	tbl := td.Tables[td.Table]
	s, _, err := f.insertOffset(td, tbl, 0)
	// Indent the statement for the template.
	s = strings.Join(strings.Split(s, "\n"), "\n\t")
	return s, err
}

// selectIndex generates the SQL statement for a SELECT operation on an index.
func (f *Funcs) selectIndex(td Typedef, i Index) (string, error) {
	tbl := td.Tables[td.Table]
	extraCols := []string{
		"\n\t" + tbl.SQLName + ".id",
	}
	whereClause := f.selectWhere(i)
	s, err := f.selectQuery(td, tbl, nil, extraCols, whereClause)
	// Indent the statement for the template.
	s = strings.Join(strings.Split(s, "\n"), "\n\t")
	return s, err
}

// list generates the SQL statement for a SELECT operation to list everything.
func (f *Funcs) list(td Typedef) (string, error) {
	tbl := td.Tables[td.Table]
	extraCols := []string{
		"\n\tquote_ident($1) AS ordered_idx, " + tbl.SQLName + ".id",
	}
	s, err := f.selectQuery(td, tbl, nil, extraCols, "%s")
	// TODO: Allow nested tables to be specified.

	// Indent the statement.
	s = strings.Join(strings.Split(s, "\n"), "\n\t\t")

	// Inject count details and filters.
	nulls := f.selectDataLiteral("NULL", td)
	s = fmt.Sprintf(
		`WITH all_entries AS (
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
		WHERE %%s
	), filtered_count AS (
		SELECT
			COUNT(*) AS count
		FROM
			filtered
	)
	SELECT
		all_count.count, filtered_count.count,
		NULL, NULL,
		%s
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
	)`,
		s, strings.Join(nulls, ",\n\t\t"),
	)
	// Indent the statement for the template.
	s = strings.Join(strings.Split(s, "\n"), "\n\t")
	return s, err
}

func (f *Funcs) update(td Typedef, by []Field) (string, error) {
	tbl := td.Tables[td.Table]
	s, _, err := f.updateOffset(td, tbl, 0, by, fromData{})
	s = strings.Join(strings.Split(s, "\n"), "\n\t")
	return s, err
}

func (f *Funcs) delete(td Typedef, by []Field) (string, error) {
	tbl := td.Tables[td.Table]
	s, err := f.recursiveDelete(td, tbl, by, fromData{})
	s = strings.Join(strings.Split(s, "\n"), "\n\t")
	return s, err
}

// insertOffset is a recursive variant of insert that accepts the current
// number of parameters used to offset its statement.
func (f *Funcs) insertOffset(td Typedef, t Table, offset int) (string, int, error) {
	// Generate the with query prefixing the insert statement.
	query, offset, err := f.insertWithQuery(td, t, offset)
	if err != nil {
		return "", 0, err
	}
	// Omit all sequences; they are automatically filled in by the database.
	seq, err := f.sequences(t)
	if err != nil {
		return "", 0, err
	}
	// Construct the parts of the insert statement.
	sqlName := f.sqlName(t)
	cols, err := f.columns("", t, seq)
	if err != nil {
		return "", 0, err
	}
	returning, err := f.columns("", seq)
	if err != nil {
		return "", 0, err
	}
	// If there are no prefixed query statements, we can use the VALUES syntax.
	if query.stmt == "" {
		vals, err := f.values(offset, t, seq)
		if err != nil {
			return "", 0, err
		}
		insertQuery := fmt.Sprintf(
			"INSERT INTO %s\n"+
				"\t(%s)\n"+
				"\tVALUES (%s)\n"+
				"\tRETURNING %s",
			sqlName, cols, vals, returning,
		)
		newOffset := offset + len(t.Fields) - len(seq)
		return insertQuery, newOffset, nil
	}
	// Create a query using a SELECT expression as the values.
	var vals []string
	for i, z := range t.Fields {
		if z.IsSequence {
			continue
		}
		if name := query.fields[i]; name != "" {
			vals = append(vals, name)
			continue
		}
		vals = append(vals, f.nth(offset))
		offset++
	}
	value := strings.Join(vals, ", ")
	withTbls := strings.Join(query.tableNames, ", ")
	insertQuery := fmt.Sprintf(
		"%s\n"+
			"INSERT INTO %s\n"+
			"\t(%s)\n"+
			"\tSELECT %s\n"+
			"\tFROM %s\n"+
			"\tRETURNING %s",
		query.stmt, sqlName, cols, value, withTbls, returning,
	)
	return insertQuery, offset, nil
}

// insertWithQuery generates the with query used in the INSERT statement of the
// specified table.
func (f *Funcs) insertWithQuery(td Typedef, t Table, offset int) (insertWithQuery, int, error) {
	query := insertWithQuery{
		fields: make([]string, len(t.Fields)),
	}

	withStmts := make([]string, 0, len(t.Fields))
	for i, field := range t.Fields {
		refTable, ok := td.Tables[field.Type]
		if !ok {
			continue
		}
		// Create the INSERT query of the nested table.
		stmt, newOffset, err := f.insertOffset(td, refTable, offset)
		if err != nil {
			return query, 0, err
		}
		offset = newOffset
		// Construct the WITH query.
		queryName := t.SQLName + "_" + field.SQLName
		stmt = strings.Join(strings.Split(stmt, "\n"), "\n\t")
		withStmt := fmt.Sprintf(
			"%s AS (\n"+
				"\t%s\n"+
				")",
			queryName, stmt,
		)
		withStmts = append(withStmts, withStmt)
		// Store extra metadata for INSERT statement generation.
		query.fields[i] = queryName + ".id"
		query.tableNames = append(query.tableNames, queryName)
	}
	if len(withStmts) == 0 {
		return query, offset, nil
	}
	// Generate the full WITH statement.
	query.stmt = "WITH " + strings.Join(withStmts, ", ")
	return query, offset, nil
}

type insertWithQuery struct {
	stmt       string
	fields     []string
	tableNames []string
}

func (f *Funcs) updateOffset(td Typedef, t Table, offset int, by []Field, from fromData) (string, int, error) {
	var stmts []string
	var updating []string

	for _, field := range t.Fields[1:] {
		if field.SQLName == "create_time" {
			continue
		}
		refTable, ok := td.Tables[field.Type]
		if !ok {
			updating = append(updating, field.SQLName)
			continue
		}

		// New scope for `from`.
		from := from
		from.tables = append(from.tables, t.SQLName)
		from.conds = append(from.conds, fmt.Sprintf("%s.%s = %s.id",
			t.SQLName, field.SQLName, refTable.SQLName))

		stmt, newOffset, err := f.updateOffset(td, refTable, offset, by, from)
		if err != nil {
			return "", 0, err
		}
		stmts = append(stmts, fmt.Sprintf(
			"%s AS (\n"+
				"\t%s\n"+
				")",
			t.SQLName+"_"+refTable.SQLName, stmt,
		))
		offset = newOffset
	}

	// Construct parts of the update statement
	sqlName := f.sqlName(t)
	for i := range updating {
		updating[i] += "=" + f.nth(offset+i)
	}
	offset += len(updating)

	var updateStmt string
	if len(from.tables) == 0 {
		cols := strings.Join(updating, ",\n\t")
		var byClauses []string
		for i, field := range by {
			byClauses = append(
				byClauses,
				fmt.Sprintf("%s = %s", field.SQLName, f.nth(offset+i)),
			)
		}
		byClause := strings.Join(byClauses, "AND\n\t")
		updateStmt = fmt.Sprintf(
			"UPDATE %s\n"+
				"SET %s\n"+
				"WHERE %s",
			sqlName, cols, byClause,
		)
	} else {
		// Build same query with `FROM ...` and extra `WHERE ...` clauses
		cols := strings.Join(updating, ",\n\t\t")
		tbls := strings.Join(from.tables, "\n\t\t")
		conds := strings.Join(from.conds, " AND\n\t\t")
		origTable := td.Tables[td.Table]
		var byClauses []string
		for i, field := range by {
			byClauses = append(
				byClauses,
				fmt.Sprintf("%s.%s = %s", origTable.SQLName, field.SQLName, f.nth(offset+i)),
			)
		}
		byClause := strings.Join(byClauses, "AND\n\t")
		updateStmt = fmt.Sprintf(
			"UPDATE %s\n"+
				"\tSET %s\n"+
				"\tFROM %s\n"+
				"\tWHERE %s AND\n"+
				"\t\t%s",
			sqlName, cols, tbls, conds, byClause,
		)
	}

	if len(stmts) == 0 {
		return updateStmt, offset + len(by), nil
	}

	stmt := fmt.Sprintf("WITH %s\n%s", strings.Join(stmts, ", "), updateStmt)

	return stmt, offset + 1, nil
}

func (f *Funcs) recursiveDelete(td Typedef, t Table, by []Field, from fromData) (string, error) {
	var stmts []string

	for _, field := range t.Fields[1:] {
		refTable, ok := td.Tables[field.Type]
		if !ok {
			continue
		}

		// New scope for `from`.
		from := from
		from.tables = append(from.tables, t.SQLName)
		from.conds = append(from.conds, fmt.Sprintf("%s.%s = %s.id",
			t.SQLName, field.SQLName, refTable.SQLName))

		stmt, err := f.recursiveDelete(td, refTable, by, from)
		if err != nil {
			return "", err
		}
		stmts = append(stmts, fmt.Sprintf(
			"%s AS (\n"+
				"\t%s\n"+
				")",
			t.SQLName+"_"+refTable.SQLName, stmt,
		))
	}

	sqlName := f.sqlName(t)
	var byClauses []string
	var deleteStmt string
	if len(from.tables) == 0 {
		for i, field := range by {
			byClauses = append(
				byClauses,
				fmt.Sprintf("%s = %s", field.SQLName, f.nth(i)),
			)
		}
		byClause := strings.Join(byClauses, " AND\n\t")

		deleteStmt = fmt.Sprintf(
			"DELETE FROM %s\n"+
				"WHERE %s",
			sqlName, byClause,
		)
	} else {
		// Build same query with `FROM ...` and extra `WHERE ...` clauses
		tbls := strings.Join(from.tables, "\n\t\t")
		conds := strings.Join(from.conds, " AND \n\t\t")
		origTable := td.Tables[td.Table]
		for i, field := range by {
			byClauses = append(
				byClauses,
				fmt.Sprintf("%s.%s = %s", origTable.SQLName, field.SQLName, f.nth(i)),
			)
		}
		byClause := strings.Join(byClauses, " AND\n\t\t")

		deleteStmt = fmt.Sprintf(
			"DELETE FROM %s\n"+
				"\tUSING %s\n"+
				"\tWHERE %s AND\n"+
				"\t\t%s",
			sqlName, tbls, conds, byClause,
		)
	}
	if len(stmts) == 0 {
		return deleteStmt, nil
	}

	stmt := fmt.Sprintf("WITH %s\n%s", strings.Join(stmts, ", "), deleteStmt)
	return stmt, nil
}

type fromData struct {
	tables []string
	conds  []string
}

// selectQuery generates the SELECT statement based on the provided typedef. It
// injects whereClause into WHERE without checks.
func (f *Funcs) selectQuery(td Typedef, t Table, joins, cols []string, whereClause string) (string, error) {
	tblJoins, tblCols := f.retrieveJoinCols(td, t, t.SQLName)
	joins = append(joins, tblJoins...)
	cols = append(cols, tblCols...)
	completeCols := strings.Join(cols, ",\n\t")
	completeJoin := strings.Join(joins, "\n\t")
	if len(joins) != 0 {
		completeJoin = "\n\t" + completeJoin
	}
	stmt := fmt.Sprintf(
		"SELECT"+
			"%s\n"+
			"FROM\n"+
			"\t%s%s\n"+
			"WHERE\n"+
			"\t%s",
		completeCols, t.SQLName, completeJoin, whereClause,
	)
	return stmt, nil
}

// retrieveJoinCols recursively JOINs and lists all simple columns in fields of
// the table, returning joins required to fetch all simple columns and the list
// of columns.
func (f *Funcs) retrieveJoinCols(td Typedef, t Table, name string) ([]string, []string) {
	var joins []string
	var cols []string
	var simpleCols []string
	for _, field := range t.Fields[1:] {
		refTable, ok := td.Tables[field.Type]
		if !ok {
			// Prepend \n\t so each column is on its own line.
			simpleCols = append(simpleCols, "\n\t"+name+"."+field.SQLName)
			continue
		}
		joinAlias := t.SQLName + "_" + field.SQLName + "_tbl"
		joinStr := fmt.Sprintf(
			"JOIN %s AS %s ON %s.id = %s.%s",
			refTable.SQLName, joinAlias,
			joinAlias, t.SQLName, field.SQLName,
		)
		joins = append(joins, joinStr)
		newJoins, newCols := f.retrieveJoinCols(td, refTable, joinAlias)
		joins = append(joins, newJoins...)
		cols = append(cols, newCols...)
	}
	if len(simpleCols) > 0 {
		cols = append(cols, strings.Join(simpleCols, ", "))
	}
	return joins, cols
}

// selectWhere returns the WHERE clause that queries based on the provided
// index without the WHERE keyword.
func (f *Funcs) selectWhere(i Index) string {
	var cond []string
	for id, field := range i.Fields {
		cond = append(cond, field.SQLName+" = "+f.nth(id))
	}
	return strings.Join(cond, "\n\tAND ")
}

// sqlName returns the SQL name of the table.
func (f *Funcs) sqlName(x Table) string {
	return x.SQLName
}

// ignoreMap normalizes the list of values to a map containing all the SQL name
// to be ignored.
func ignoreMap(vals []interface{}) (map[string]bool, error) {
	ignore := map[string]bool{}
	for i, z := range vals {
		switch x := z.(type) {
		case string:
			ignore[x] = true
		case Field:
			ignore[x.SQLName] = true
		case []Field:
			for _, field := range x {
				ignore[field.SQLName] = true
			}
		default:
			return nil, fmt.Errorf("unknown ignore(%d) type: %T", i, z)
		}
	}
	return ignore, nil
}

// columns returns the SQL names of all columns in v excluding ones specified
// in ignoreVals, joined together with commas in between.
func (f *Funcs) columns(prefix string, v interface{}, ignoreVals ...interface{}) (string, error) {
	ignore, err := ignoreMap(ignoreVals)
	if err != nil {
		return "", err
	}

	var vals []string
	switch x := v.(type) {
	case Table:
		return f.columns(prefix, x.Fields, ignoreVals)
	case []Field:
		for _, z := range x {
			if ignore[z.SQLName] {
				continue
			}
			vals = append(vals, prefix+z.SQLName)
		}
	default:
		return "", fmt.Errorf("invalid provided type for columns: %T", v)
	}
	return strings.Join(vals, ", "), nil
}

// sequences returns a list of fields that are sequences in v.
func (f *Funcs) sequences(v interface{}) ([]Field, error) {
	switch x := v.(type) {
	case Table:
		return f.sequences(x.Fields)
	case []Field:
		var fields []Field
		for _, z := range x {
			if z.IsSequence {
				fields = append(fields, z)
			}
		}
		return fields, nil
	}
	return nil, fmt.Errorf("invalid provided type for sequences: %T", v)
}

// values returns value placeholders for all columns in v that are not present
// in ignoreVals. It starts numbering the placeholders from offset if
// applicable.
func (f *Funcs) values(offset int, v interface{}, ignoreVals ...interface{}) (string, error) {
	ignore, err := ignoreMap(ignoreVals)
	if err != nil {
		return "", err
	}

	var vals []string
	var i int
	switch x := v.(type) {
	case Table:
		return f.values(offset, x.Fields, ignoreVals)
	case []Field:
		for _, z := range x {
			if ignore[z.SQLName] {
				continue
			}
			vals = append(vals, f.nth(i+offset))
			i++
		}
	default:
		return "", fmt.Errorf("unsupported type for values: %T", v)
	}
	return strings.Join(vals, ", "), nil
}
