package impl

import (
	"strings"
)

// insertData generates the data to insert. It generates one entry per table.
func (f *Funcs) insertData(varName string, t Typedef) []string {
	query := f.dataForTable(varName, t, t.Tables[t.Table])
	return query
}

// selectData generates the data to select. It generates one entry per table.
func (f *Funcs) selectData(varName string, t Typedef) []string {
	query := f.dataForTable("&"+varName, t, t.Tables[t.Table])
	return query
}

// selectDataLiteral generates the data to select, replacing every entry with
// the literal provided.
func (f *Funcs) selectDataLiteral(literal string, t Typedef) []string {
	query := f.dataForTable("a", t, t.Tables[t.Table])
	for k, v := range query {
		commaStr := strings.Repeat(literal+", ", strings.Count(v, ",")+1)
		query[k] = strings.TrimSuffix(commaStr, ", ")
	}
	return query
}

// updateData generates the data to update. It generates one entry per table.
//
// TODO(optimization): this could be changed to reference the same `varName.ID`
// every time instead, but this is realistically only going to reduce the total
// number of prepared statement arguments by 1-2.
func (f *Funcs) updateData(varName string, t Typedef, by []Field) []string {
	query := f.dataForTable(varName, t, t.Tables[t.Table], "create_time")
	var toAdd string
	for _, field := range by {
		toAdd += ", " + varName + "." + field.GoName
	}
	for i := range query {
		query[i] += toAdd
	}
	return query
}

// dataForTable generates the data for the table with the specified prefix.
// It recursively invokes itself for nested tables.
func (f *Funcs) dataForTable(prefix string, t Typedef, tbl Table, ignore ...string) []string {
	var result []string
	var nonMessageFields []string
	// Generate data of nested tables first.
	for _, field := range tbl.Fields[1:] {
		if contains(ignore, field.SQLName) {
			continue
		}
		refTable, ok := t.Tables[field.Type]
		if !ok {
			nonMessageFields = append(nonMessageFields, prefix+"."+field.GoName)
			continue
		}
		refData := f.dataForTable(prefix+"."+field.GoName, t, refTable, ignore...)
		result = append(result, refData...)
	}
	// Add all non-message fields in the correct order.
	if len(nonMessageFields) > 0 {
		result = append(result, strings.Join(nonMessageFields, ", "))
	}
	return result
}

func contains(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}
