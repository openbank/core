package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/kenshaw/inflector"
	"github.com/kenshaw/snaker"
	"github.com/xo/xo/loader"
	xo "github.com/xo/xo/types"
	"golang.org/x/tools/imports"
	"mvdan.cc/gofumpt/format"
)

// Init registers the template.
func Init(_ context.Context, f func(xo.TemplateType)) error {
	f(xo.TemplateType{
		Modes: []string{"schema"},
		Flags: []xo.Flag{
			{
				ContextKey: ProtoKey,
				Type:       "string",
				Desc:       "protobuf names for tables",
			},
			{
				ContextKey: ConflictKey,
				Type:       "string",
				Desc:       "name conflict suffix",
				Default:    "Val",
			},
			{
				ContextKey: InitialismKey,
				Type:       "[]string",
				Desc:       "add initialism (e.g. ID, API, URI, ...)",
			},
		},
		Funcs: func(ctx context.Context, _ string) (template.FuncMap, error) {
			return NewFuncs(ctx)
		},
		Order: func(ctx context.Context, mode string) []string {
			return []string{"header", "db", "typedef", "query"}
		},
		Pre: func(ctx context.Context, mode string, set *xo.Set, _ fs.FS, emit func(xo.Template)) error {
			if err := addInitialisms(ctx); err != nil {
				return err
			}
			proto, err := Proto(ctx)
			if err != nil {
				return err
			}
			files, err := fileNames(ctx, mode, set, proto)
			if err != nil {
				return err
			}
			emit(xo.Template{
				Partial: "db",
				Dest:    "db" + ext,
			})
			for filename := range files {
				emit(xo.Template{
					Partial: "header",
					Dest:    filename,
				})
			}
			return nil
		},
		Process: func(ctx context.Context, mode string, set *xo.Set, emit func(xo.Template)) error {
			if mode == "query" {
				for _, query := range set.Queries {
					if err := emitQuery(ctx, query, emit); err != nil {
						return err
					}
				}
			} else {
				proto, err := Proto(ctx)
				if err != nil {
					return err
				}
				for _, schema := range set.Schemas {
					if err := emitSchema(ctx, schema, proto, emit); err != nil {
						return err
					}
				}
			}
			return nil
		},
		Post: func(_ context.Context, _ string, files map[string][]byte, emit func(string, []byte)) error {
			for file, content := range files {
				// Run goimports.
				buf, err := imports.Process("", content, nil)
				if err != nil {
					return fmt.Errorf("%s:%w", file, err)
				}
				// Run gofumpt.
				formatted, err := format.Source(buf, format.Options{
					ExtraRules: true,
				})
				if err != nil {
					return fmt.Errorf("error formatting go source: %w", err)
				}
				lines := bytes.Split(formatted, []byte("\n"))
				for i := range lines {
					lines[i] = bytes.TrimRightFunc(lines[i], unicode.IsSpace)
				}
				emit(file, bytes.Join(lines, []byte("\n")))
			}
			return nil
		},
	})
	return nil
}

// fileNames returns a list of file names that will be generated by the
// template based on the parameters and schema.
func fileNames(_ context.Context, mode string, set *xo.Set, proto map[string]typeEntry) (map[string]bool, error) {
	// Infer filenames from set.
	files := make(map[string]bool)
	addFiles := func(filenames ...string) {
		// Filenames are always lowercase.
		for _, f := range filenames {
			f = strings.ToLower(f)
			files[f+ext] = true
		}
	}
	switch mode {
	case "schema":
		addFiles("db", "interface", "enums")
		for _, schema := range set.Schemas {
			for _, t := range schema.Tables {
				if _, ok := proto[t.Name]; !ok {
					continue
				}
				addFiles(camelExport(singularize(t.Name)))
			}
		}
	case "query":
		for _, query := range set.Queries {
			addFiles(query.Type)
			if query.Exec {
				return nil, fmt.Errorf("query exec mode is not supported")
			}
		}
	default:
		panic("unknown mode: " + mode)
	}
	return files, nil
}

// emitQuery emits the query.
func emitQuery(ctx context.Context, query xo.Query, emit func(xo.Template)) error {
	var table Table
	// build type if needed
	if !query.Exec {
		var err error
		if table, err = buildQueryType(ctx, query); err != nil {
			return err
		}
	}
	// emit type definition
	if !query.Exec && !query.Flat {
		emit(xo.Template{
			Partial:  "typedef",
			Dest:     strings.ToLower(table.GoName) + ext,
			SortType: query.Type,
			SortName: query.Name,
			Data:     table,
		})
	}
	// build query params
	params := make([]QueryParam, 0, len(query.Params))
	for _, param := range query.Params {
		params = append(params, QueryParam{
			Name:        param.Name,
			Type:        param.Type.Type,
			Interpolate: param.Interpolate,
			Join:        param.Join,
		})
	}
	// emit query
	emit(xo.Template{
		Partial:  "query",
		Dest:     strings.ToLower(table.GoName) + ext,
		SortType: query.Type,
		SortName: query.Name,
		Data: Query{
			Name:        buildQueryName(query),
			Query:       query.Query,
			Comments:    query.Comments,
			Params:      params,
			One:         query.Exec || query.Flat || query.One,
			Flat:        query.Flat,
			Exec:        query.Exec,
			Interpolate: query.Interpolate,
			Type:        table,
			Comment:     query.Comment,
		},
	})
	return nil
}

func buildQueryType(ctx context.Context, query xo.Query) (Table, error) {
	tf := camelExport
	if query.Flat {
		tf = camel
	}
	fields := make([]Field, 0, len(query.Fields))
	for _, z := range query.Fields {
		f, err := convertField(ctx, tf, z)
		if err != nil {
			return Table{}, err
		}
		// dont use convertField; the types are already provided by the user
		if query.ManualFields {
			f = Field{
				GoName:  z.Name,
				SQLName: snake(z.Name),
				Type:    z.Type.Type,
			}
		}
		fields = append(fields, f)
	}
	sqlName := snake(query.Type)
	return Table{
		GoName:  query.Type,
		SQLName: sqlName,
		Fields:  fields,
		Comment: query.TypeComment,
	}, nil
}

// buildQueryName builds a name for the query.
func buildQueryName(query xo.Query) string {
	if query.Name != "" {
		return query.Name
	}
	// generate name if not specified
	name := query.Type
	if !query.One {
		name = inflector.Pluralize(name)
	}
	// add params
	if len(query.Params) == 0 {
		name = "Get" + name
	} else {
		name += "By"
		for _, p := range query.Params {
			name += camelExport(p.Name)
		}
	}
	return name
}

// convertTable converts a xo.Table to a Table.
// emitSchema emits the xo schema for the template set.
func emitSchema(ctx context.Context, schema xo.Schema, proto map[string]typeEntry, emit func(xo.Template)) error {
	// emit tables
	tables := make(map[string]Table)
	for _, t := range schema.Tables {
		typeInfo, ok := proto[t.Name]
		if !ok {
			continue
		}
		table, err := convertTable(ctx, typeInfo, t)
		if err != nil {
			return err
		}
		// emit indexes
		for _, i := range t.Indexes {
			index, err := convertIndex(ctx, table, i)
			if err != nil {
				return err
			}
			table.Indexes = append(table.Indexes, index)
		}
		// emit fkeys
		for _, fk := range t.ForeignKeys {
			fkey, err := convertFKey(ctx, table, fk)
			if err != nil {
				return err
			}
			table.FKeys = append(table.FKeys, fkey)
		}
		tables[table.GoName] = table
	}
	for _, table := range tables {
		emit(xo.Template{
			Dest:     strings.ToLower(table.GoName) + ext,
			Partial:  "typedef",
			SortType: table.Type,
			SortName: table.GoName,
			Data: Typedef{
				Tables: tables,
				Table:  table.GoName,
			},
		})
	}
	return nil
}

func convertTable(ctx context.Context, entry typeEntry, t xo.Table) (Table, error) {
	var cols, pkCols []Field
	for _, z := range t.Columns {
		f, err := convertField(ctx, camelExport, z)
		if ref, ok := getReference(ctx, entry, t, z); ok {
			f.Type = camelExport(singularize(ref))
		}
		if err != nil {
			return Table{}, err
		}
		cols = append(cols, f)
		if z.IsPrimary {
			pkCols = append(pkCols, f)
		}
	}
	return Table{
		GoName:      camelExport(singularize(t.Name)),
		SQLName:     t.Name,
		Fields:      cols,
		PrimaryKeys: pkCols,
		Manual:      t.Manual,
	}, nil
}

func convertIndex(ctx context.Context, _ Table, i xo.Index) (Index, error) {
	var fields []Field
	for _, z := range i.Fields {
		f, err := convertField(ctx, camelExport, z)
		if err != nil {
			return Index{}, err
		}
		fields = append(fields, f)
	}
	return Index{
		SQLName:   i.Name,
		Func:      camelExport(i.Func),
		Fields:    fields,
		IsUnique:  i.IsUnique,
		IsPrimary: i.IsPrimary,
	}, nil
}

func convertFKey(ctx context.Context, _ Table, fk xo.ForeignKey) (ForeignKey, error) {
	var fields, refFields []Field
	// convert fields
	for _, f := range fk.Fields {
		field, err := convertField(ctx, camelExport, f)
		if err != nil {
			return ForeignKey{}, err
		}
		fields = append(fields, field)
	}
	// convert ref fields
	for _, f := range fk.RefFields {
		refField, err := convertField(ctx, camelExport, f)
		if err != nil {
			return ForeignKey{}, err
		}
		refFields = append(refFields, refField)
	}
	return ForeignKey{
		GoName:    camelExport(fk.Func),
		SQLName:   fk.Name,
		Fields:    fields,
		RefTable:  camelExport(singularize(fk.RefTable)),
		RefFields: refFields,
		RefFunc:   camelExport(fk.RefFunc),
	}, nil
}

func convertField(ctx context.Context, tf transformFunc, f xo.Field) (Field, error) {
	typ, zero, err := goType(ctx, f.Type)
	if err != nil {
		return Field{}, err
	}
	return Field{
		Type:       typ,
		GoName:     tf(f.Name),
		SQLName:    f.Name,
		Zero:       zero,
		IsPrimary:  f.IsPrimary,
		IsSequence: f.IsSequence,
	}, nil
}

func getReference(ctx context.Context, entry typeEntry, t xo.Table, f xo.Field) (string, bool) {
	for _, fk := range t.ForeignKeys {
		if len(fk.Fields) != 1 || fk.Fields[0] != f {
			continue
		}
		field, _ := convertField(ctx, camelExport, f)
		if entry.Fields[field.GoName].Kind == "basic" {
			return "", false
		}
		return fk.RefTable, true
	}
	return "", false
}

func goType(ctx context.Context, typ xo.Type) (string, string, error) {
	driver, _, schema := xo.DriverDbSchema(ctx)
	var f func(xo.Type, string, string, string) (string, string, error)
	switch driver {
	case "mysql":
		f = loader.MysqlGoType
	case "oracle":
		f = loader.OracleGoType
	case "postgres":
		f = loader.PostgresGoType
	case "sqlite3":
		f = loader.Sqlite3GoType
	case "sqlserver":
		f = loader.SqlserverGoType
	default:
		return "", "", fmt.Errorf("unknown driver %q", driver)
	}
	return f(typ, schema, "int32", "uint32")
}

type transformFunc func(...string) string

func snake(names ...string) string {
	return snaker.CamelToSnake(strings.Join(names, "_"))
}

func camel(names ...string) string {
	return snaker.ForceLowerCamelIdentifier(strings.Join(names, "_"))
}

func camelExport(names ...string) string {
	return snaker.ForceCamelIdentifier(strings.Join(names, "_"))
}

const ext = ".xo.go"

// Funcs is a set of template funcs.
type Funcs struct {
	driver   string
	schema   string
	nth      func(int) string
	pkg      string
	conflict string
	// knownTypes is the collection of known Go types.
	knownTypes map[string]bool
	// shorts is the collection of Go style short names for types, mainly
	// used for use with declaring a func receiver on a type.
	shorts map[string]string
}

type typeEntry struct {
	PkgName  string                `json:"pkg_name"`
	PkgPath  string                `json:"pkg_path"`
	TypeName string                `json:"type"`
	Fields   map[string]fieldEntry `json:"fields"`
}

type fieldEntry struct {
	PkgName string `json:"pkg_name"`
	PkgPath string `json:"pkg_path"`
	Type    string `json:"type"`
	Kind    string `json:"kind"`
	Array   bool   `json:"repeated"`
	Map     bool   `json:"map"`
}

// NewFuncs creates custom template funcs for the context.
func NewFuncs(ctx context.Context) (template.FuncMap, error) {
	driver, _, schema := xo.DriverDbSchema(ctx)
	nth, err := loader.NthParam(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get nth param func: %w", err)
	}
	knownTypes := map[string]bool{
		"bool":        true,
		"string":      true,
		"byte":        true,
		"rune":        true,
		"int":         true,
		"int16":       true,
		"int32":       true,
		"int64":       true,
		"uint":        true,
		"uint8":       true,
		"uint16":      true,
		"uint32":      true,
		"uint64":      true,
		"float32":     true,
		"float64":     true,
		"Slice":       true,
		"StringSlice": true,
	}
	funcs := &Funcs{
		driver:     driver,
		schema:     schema,
		nth:        nth,
		pkg:        filepath.Base(xo.Out(ctx)),
		conflict:   Conflict(ctx),
		knownTypes: knownTypes,
		shorts:     make(map[string]string),
	}
	return funcs.FuncMap(), nil
}

// FuncMap returns the func map.
func (f *Funcs) FuncMap() template.FuncMap {
	return template.FuncMap{
		// general
		"pkg":    f.pkgfn,
		"driver": f.driverfn,
		"series": series,
		"add":    add,
		// type
		"names":          f.names,
		"typed_params":   f.typedParams,
		"untyped_params": f.untypedParams,
		"param_name":     f.paramName,
		"type":           f.typefn,
		"short":          f.short,
		"plural":         plural,
		// sql funcs
		"insert":              f.insert,
		"insert_data":         f.insertData,
		"select_idx":          f.selectIndex,
		"select_data":         f.selectData,
		"select_data_literal": f.selectDataLiteral,
		"update":              f.update,
		"update_data":         f.updateData,
		"delete":              f.delete,
		"list":                f.list,
		// helpers
		"check_name": checkName,
		"join":       join,
		"eval":       eval,
	}
}

// pkgfn returns the package name.
func (f *Funcs) pkgfn() string {
	return f.pkg
}

// driverfn returns the driver name.
func (f *Funcs) driverfn() string {
	return f.driver
}

func series(start int, end int) []int {
	ret := make([]int, 0, end-start)
	for i := start; i < end; i++ {
		ret = append(ret, i)
	}
	return ret
}

func add(a, b int) int {
	return a + b
}

// names generates a list of names (excluding certain ones such as interpolated
// names).
func (f *Funcs) names(prefix string, z ...interface{}) (string, error) {
	var names []string
	for i, v := range z {
		switch x := v.(type) {
		case string:
			names = append(names, x)
		case Table:
			for _, p := range x.Fields {
				names = append(names, prefix+checkName(p.GoName))
			}
		case []Field:
			for _, p := range x {
				names = append(names, prefix+checkName(p.GoName))
			}
		case Index:
			param, err := f.untypedParams(x.Fields)
			if err != nil {
				return "", err
			}
			names = append(names, param)
		default:
			return "", fmt.Errorf("unsupported type for namesfn(%d): %T", i, v)
		}
	}
	return strings.Join(names, ", "), nil
}

// typedParams converts a list of fields into their named Go parameters.
//
// Any field name encountered will be checked against goReservedNames, and will
// have its name substituted by its corresponding looked up value.
func (f *Funcs) typedParams(fields []Field) (string, error) {
	vals := make([]string, 0, len(fields))
	for _, field := range fields {
		vals = append(vals, f.paramName(field)+" "+f.typefn(field.Type))
	}
	return strings.Join(vals, ", "), nil
}

// untypedParams converts a list of fields into their Go names, separated by
// commas. To include types, use typedParams.
func (f *Funcs) untypedParams(fields []Field) (string, error) {
	vals := make([]string, 0, len(fields))
	for _, field := range fields {
		vals = append(vals, f.paramName(field))
	}
	return strings.Join(vals, ", "), nil
}

// paramName returns the name of the field as a parameter.
func (f *Funcs) paramName(field Field) string {
	name := snaker.ForceLowerCamelIdentifier(field.GoName)
	if goName, ok := goReservedNames[name]; ok {
		return goName
	}
	return name
}

// typefn generates the Go type, prefixing the custom package name if applicable.
func (f *Funcs) typefn(typ string) string {
	if strings.Contains(typ, ".") {
		return typ
	}
	var prefix string
	for strings.HasPrefix(typ, "[]") {
		typ = typ[2:]
		prefix += "[]"
	}
	if _, ok := f.knownTypes[typ]; !ok {
		return prefix + "db." + typ
	}
	return prefix + typ
}

// short generates a safe Go identifier for typ. typ is first checked
// against shorts, and if not found, then the value is calculated and
// stored in the shorts for future use.
//
// A short is the concatenation of the lowercase of the first character in
// the words comprising the name. For example, "MyCustomName" will have have
// the short of "mcn".
//
// If a generated short conflicts with a Go reserved name or a name used in
// the templates, then the corresponding value in goReservedNames map will be
// used.
//
// Generated shorts that have conflicts with any scopeConflicts member will
// have nameConflictSuffix appended.
func (f *Funcs) short(v interface{}) (string, error) {
	var n string
	switch x := v.(type) {
	case string:
		n = x
	case Table:
		n = x.GoName
	case Enum:
		n = x.GoName
	default:
		return "", fmt.Errorf("unsupported type for short: %T", v)
	}
	// check short name map
	name, ok := f.shorts[n]
	if !ok {
		// calc the short name
		var u []string
		for _, s := range strings.Split(strings.ToLower(snaker.CamelToSnake(n)), "_") {
			if len(s) > 0 && s != "id" {
				u = append(u, s[:1])
			}
		}
		// ensure no name conflict
		name = checkName(strings.Join(u, ""))
		// store back to short name map
		f.shorts[n] = name
	}
	// append suffix if conflict exists
	if _, ok := templateReservedNames[name]; ok {
		name += f.conflict
	}
	return name, nil
}

func plural(s string) string {
	return inflector.Pluralize(s)
}

func checkName(name string) string {
	if n, ok := goReservedNames[name]; ok {
		return n
	}
	return name
}

func join(v interface{}, sep string) (string, error) {
	switch x := v.(type) {
	case []string:
		return strings.Join(x, sep), nil
	case []Field:
		var fields []string
		for _, f := range x {
			fields = append(fields, f.GoName)
		}
		return strings.Join(fields, sep), nil
	}
	return "", fmt.Errorf("unsupported type for join: %T", v)
}

// eval evalutates a template s against v.
func eval(v interface{}, s string) (string, error) {
	tpl, err := template.New(fmt.Sprintf("[EVAL %q]", s)).Parse(s)
	if err != nil {
		return "", fmt.Errorf("error parsing eval template %q: %w", s, err)
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, v); err != nil {
		return "", fmt.Errorf("error executing eval template: %w", err)
	}
	return buf.String(), nil
}

// templateReservedNames are the template reserved names.
var templateReservedNames = map[string]bool{
	// variables
	"ctx":  true,
	"db":   true,
	"err":  true,
	"log":  true,
	"logf": true,
	"res":  true,
	"rows": true,

	// packages
	"context": true,
	"csv":     true,
	"driver":  true,
	"errors":  true,
	"fmt":     true,
	"hstore":  true,
	"regexp":  true,
	"sql":     true,
	"strings": true,
	"time":    true,
	"uuid":    true,
}

// goReservedNames is a map of of go reserved names to "safe" names.
var goReservedNames = map[string]string{
	"break":       "brk",
	"case":        "cs",
	"chan":        "chn",
	"const":       "cnst",
	"continue":    "cnt",
	"default":     "def",
	"defer":       "dfr",
	"else":        "els",
	"fallthrough": "flthrough",
	"for":         "fr",
	"func":        "fn",
	"go":          "goVal",
	"goto":        "gt",
	"if":          "ifVal",
	"import":      "imp",
	"interface":   "iface",
	"map":         "mp",
	"package":     "pkg",
	"range":       "rnge",
	"return":      "ret",
	"select":      "slct",
	"struct":      "strct",
	"switch":      "swtch",
	"type":        "typ",
	"var":         "vr",
	// go types
	"error":      "e",
	"bool":       "b",
	"string":     "str",
	"byte":       "byt",
	"rune":       "r",
	"uintptr":    "uptr",
	"int":        "i",
	"int8":       "i8",
	"int16":      "i16",
	"int32":      "i32",
	"int64":      "i64",
	"uint":       "u",
	"uint8":      "u8",
	"uint16":     "u16",
	"uint32":     "u32",
	"uint64":     "u64",
	"float32":    "z",
	"float64":    "f",
	"complex64":  "c",
	"complex128": "c128",
}

// Context keys.
var (
	ProtoKey      xo.ContextKey = "protobuf"
	ConflictKey   xo.ContextKey = "conflict"
	InitialismKey xo.ContextKey = "initialism"
)

// Proto returns the proto names json from the context.
func Proto(ctx context.Context) (map[string]typeEntry, error) {
	s, _ := ctx.Value(ProtoKey).(string)
	proto := make(map[string]typeEntry)
	if err := json.Unmarshal([]byte(s), &proto); err != nil {
		return nil, fmt.Errorf("could not unmarshal proto: %w", err)
	}
	return proto, nil
}

// Conflict returns conflict from the context.
func Conflict(ctx context.Context) string {
	s, _ := ctx.Value(ConflictKey).(string)
	return s
}

// addInitialisms adds snaker initialisms from the context.
func addInitialisms(ctx context.Context) error {
	z := ctx.Value(InitialismKey)
	y, _ := z.([]string)
	var v []string
	for _, s := range y {
		if s != "" {
			v = append(v, s)
		}
	}
	return snaker.DefaultInitialisms.Add(v...)
}

// singularize singularizes s.
func singularize(s string) string {
	if i := strings.LastIndex(s, "_"); i != -1 {
		return s[:i+1] + inflector.Singularize(s[i+1:])
	}
	return inflector.Singularize(s)
}

// EnumValue is a enum value template.
type EnumValue struct {
	GoName     string
	SQLName    string
	ConstValue int
}

// Enum is a enum type template.
type Enum struct {
	GoName  string
	SQLName string
	Values  []EnumValue
	Comment string
}

type Typedef struct {
	Tables map[string]Table
	Table  string
}

// Table is a type (ie, table/view/custom query) template.
type Table struct {
	Type        string
	GoName      string
	SQLName     string
	PrimaryKeys []Field
	Fields      []Field
	Manual      bool
	Comment     string
	FKeys       []ForeignKey
	Indexes     []Index
}

// ForeignKey is a foreign key template.
type ForeignKey struct {
	GoName    string
	SQLName   string
	Fields    []Field
	RefTable  string
	RefFields []Field
	RefFunc   string
	Comment   string
}

// Index is an index template.
type Index struct {
	SQLName   string
	Func      string
	Fields    []Field
	IsUnique  bool
	IsPrimary bool
	Comment   string
}

// Field is a field template.
type Field struct {
	GoName     string
	SQLName    string
	Type       string
	Zero       string
	IsPrimary  bool
	IsSequence bool
	Comment    string
}

// QueryParam is a custom query parameter template.
type QueryParam struct {
	Name        string
	Type        string
	Interpolate bool
	Join        bool
}

// Query is a custom query template.
type Query struct {
	Name        string
	Query       []string
	Comments    []string
	Params      []QueryParam
	One         bool
	Flat        bool
	Exec        bool
	Interpolate bool
	Type        Table
	Comment     string
}
