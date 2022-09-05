{{ define "db" -}}
var (
	// logf is used by generated code to log SQL queries.
	logf = func(string, ...interface{}) {}
	// errf is used by generated code to log SQL errors.
	errf = func(string, ...interface{}) {}
)

// SetLogger sets the package logger. Valid logger types:
//
//     io.Writer
//     func(string, ...interface{}) (int, error) // fmt.Printf
//     func(string, ...interface{}) // log.Printf
//
func SetLogger(logger interface{}) {
	logf = newLogger(logger)
}

// SetErrorLogger sets the package error logger. Valid logger types:
//
//     io.Writer
//     func(string, ...interface{}) (int, error) // fmt.Printf
//     func(string, ...interface{}) // log.Printf
//
func SetErrorLogger(logger interface{}) {
	errf = newLogger(logger)
}

// Logf logs a message using the package logger.
func Logf(s string, v ...interface{}) {
	logf(s, v...)
}

// Errorf logs an error message using the package error logger.
func Errorf(s string, v ...interface{}) {
	errf(s, v...)
}

// logerror logs the error and returns it.
func logerror(err error) error {
	errf("Error: %v", err)
	return err
}

func newLogger(logger interface{}) func(string, ...interface{}) {
	switch z := logger.(type) {
	case io.Writer:
		return func(s string, v ...interface{}) {
			fmt.Fprintf(z, s, v...)
		}
	case func(string, ...interface{}) (int, error): // fmt.Printf
		return func(s string, v ...interface{}) {
			_, _ = z(s, v...)
		}
	case func(string, ...interface{}): // log.Printf
		return z
	}
	panic(fmt.Sprintf("unsupported logger type %T", logger))
}

// DB is the common interface for database operations.
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// convertEnum maps a type of enum to another type. It adds the provided offset
// to map the enums.
func convertEnum[T, U ~int32](t []T, offset U) []U {
	u := make([]U, len(t))
	for i := range t {
		u[i] = U(t[i]) + offset
	}
	return u
}

// protoMessage is an instance of proto.Message, asserted to be a pointer.
type protoMessage[V any] interface {
	proto.Message
	*V
}

// initProtoMessage initializes the proto.Message. It is used to avoid creating
// nil pointers.
func initProtoMessage[T protoMessage[V], V any]() T {
	return T(new(V))
}

// marshalArray marshals an array of proto.Messages into JSON.
func marshalArray[T proto.Message](t []T) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("[")
	for i, v := range t {
		msg, err := protojson.Marshal(v)
		if err != nil {
			return nil, err
		}
		b.Write(msg)
		if i != len(t)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString("]")
	return b.Bytes(), nil
}

// unmarshalArray unmarshals JSON into an array of proto.Messages.
func unmarshalArray[T protoMessage[V], V any](b []byte, t *[]T) error {
	var array []json.RawMessage
	if err := json.Unmarshal(b, &array); err != nil {
		return err
	}
	result := make([]T, len(array))
	for i := range array {
		if err := unmarshalMessage(array[i], &result[i]); err != nil {
			return err
		}
	}
	*t = result
	return nil
}

// unmarshalMessage unmarshals JSON into the provided proto.Message.
func unmarshalMessage[T protoMessage[V], V any](b []byte, t *T) error {
	res := initProtoMessage[T]()
	if err := protojson.Unmarshal(b, res); err != nil {
		return err
	}
	*t = res
	return nil
}

func mapError[F any, T any](pbs []F, f func(F) (T, error)) ([]T, error) {
	array  := make([]T, len(pbs))
	for i, pb := range pbs {
		v, err := f(pb)
		if err != nil {
			return nil, err
		}
		array[i] = v
	}
	return array, nil
}

func toStringPB(s sql.NullString) string {
	if !s.Valid {
		return ""
	}
	return s.String
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func toTimePB(t sql.NullTime) *timestamppb.Timestamp {
	if !t.Valid {
		return nil
	}
	return timestamppb.New(t.Time)
}

func toNullTime(t *timestamppb.Timestamp) sql.NullTime {
	if t != nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time: t.AsTime(),
		Valid: true,
	}
}

// ListStat is statistics related to a List endpoint.
type ListStat struct {
	// Total is the number of total entries in the list.
	Total int32
	// Remaining is the number of entries to be enumerated, including the
	// current page.
	Remaining int32
}

// ListPosition is the position of the last entry in the list query.
type ListPosition struct {
	// Data is the entry the query sorted by.
	Data interface{}
	// ID is the ID of the entry.
	ID int
}

// ErrInvalidStringSlice is the invalid StringSlice error.
const ErrInvalidStringSlice Error = "invalid StringSlice"

// StringSlice is a slice of strings.
type StringSlice []string

// Scan satisfies the sql.Scanner interface for StringSlice.
func (ss *StringSlice) Scan(v interface{}) error {
	buf, ok := v.([]byte)
	if !ok {
		return logerror(ErrInvalidStringSlice)
	}
	// change quote escapes for csv parser
	str := strings.Replace(quoteEscRE.ReplaceAllString(string(buf), `$1""`), `\\`, `\`, -1)
	str = str[1 : len(str)-1]
	// bail if only one
	if len(str) == 0 {
		return nil
	}
	// parse with csv reader
	r := csv.NewReader(strings.NewReader(str))
	line, err := r.Read()
	if err != nil {
		return logerror(&ErrDecodeFailed{err})
	}
	*ss = StringSlice(line)
	return nil
}

// quoteEscRE matches escaped characters in a string.
var quoteEscRE = regexp.MustCompile(`([^\\]([\\]{2})*)\\"`)

// Value satisfies the sql/driver.Valuer interface.
func (ss StringSlice) Value() (driver.Value, error) {
	v := make([]string, len(ss))
	for i, s := range ss {
		v[i] = `"` + strings.Replace(strings.Replace(s, `\`, `\\\`, -1), `"`, `\"`, -1) + `"`
	}
	return "{" + strings.Join(v, ",") + "}", nil
}
{{- end }}