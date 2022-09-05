{{ define "interface" }}
type Storage interface {
	// Create and close transactions and database connections
	WithTx(tx *sql.Tx) Storage
	// Available table repositories
{{- range .Data }}
	{{ .GoName }}Repository
{{- end }}
}
{{ end }}

{{ define "enum" }}
{{- $e := .Data -}}
// {{ $e.GoName }} is the '{{ $e.SQLName }}' enum.
type {{ $e.GoName }} int32

// {{ $e.GoName }} values.
const (
{{ range $e.Values -}}
	// {{ $e.GoName }}{{ .GoName }} is the '{{ .SQLName }}' {{ $e.SQLName }}.
	{{ $e.GoName }}{{ .GoName }} {{ $e.GoName }} = {{ .ConstValue }}
{{ end -}}
)

func New{{ $e.GoName }} ({{ short $e }} {{ protoname $e.SQLName }}) {{$e.GoName }} {
	return {{ $e.GoName }}({{ short $e }} + 1)
}

func ({{ short $e }} {{ $e.GoName }}) PB() {{ protoname $e.SQLName }} {
	return {{ protoname $e.SQLName }}({{ short $e }}-1)
}

// UnmarshalText unmarshals {{ $e.GoName }} from text.
func ({{ short $e.GoName }} *{{ $e.GoName }}) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
{{ range $e.Values -}}
	case "{{ .SQLName }}":
		*{{ short $e.GoName }} = {{ $e.GoName }}{{ .GoName }}
{{ end -}}
	default:
		return ErrInvalid{{ $e.GoName }}(str)
	}
	return nil
}

// String satisfies the fmt.Stringer interface.
func ({{ short $e }} {{ $e.GoName }}) String() string {
	switch {{ short $e }} {
{{ range $e.Values -}}
	case {{ $e.GoName }}{{ .GoName }}:
		return "{{ .SQLName }}"
{{ end -}}
	}
	return fmt.Sprintf("{{ $e.GoName }}(%d)", {{ short $e }})
}


// Value satisfies the driver.Valuer interface.
func ({{ short $e }} {{ $e.GoName }}) Value() (driver.Value, error) {
	return {{ short $e }}.String(), nil
}

// Scan satisfies the sql.Scanner interface.
func ({{ short $e }} *{{ $e.GoName }}) Scan(v interface{}) error {
	if buf, ok := v.([]byte); ok {
		return {{ short $e }}.UnmarshalText(buf)
	}
	return ErrInvalid{{ $e.GoName }}(fmt.Sprintf("%T", v))
}

// ErrInvalid{{ $e.GoName }} is the invalid {{ $e.GoName }} error.
type ErrInvalid{{ $e.GoName }} string

// Error satisfies the error interface.
func (err ErrInvalid{{ $e.GoName }}) Error() string {
	return fmt.Sprintf("invalid {{ $e.GoName }}(%s)", string(err))
}
{{ end }}

{{ define "typedef" }}
{{- $t := .Data -}}
{{- if $t.Comment -}}
// {{ $t.Comment | eval $t.GoName }}
{{- else -}}
// {{ $t.GoName }} represents a row from '{{ $t.SQLName }}'.
{{- end }}
type {{ $t.GoName }} struct {
{{ range $t.Fields -}}
	{{ field . }}
{{ end }}
{{- if $t.PrimaryKeys -}}
	// xo fields
	Exists, Deleted bool
{{ end -}}
}

func New{{ $t.GoName }}(pb *{{ protoname $t.SQLName }}) ({{ $t.GoName }}, error) {
	if pb == nil {
		return {{ $t.GoName }}{}, ErrNilType{"{{ $t.GoName }}"}
	}
	{{ short $t }} := {{ $t.GoName }}{
{{- range fields $t false }}
		{{ .GoName }}: {{ from_pb $t . }},
{{- end }}
	}

{{- with fields $t true }}
	var err error
{{- range . }}
{{- if eq .Type "[]byte" }}
	{{ short $t }}.{{ .GoName }}, err = {{ marshaler $t . }}
	if err != nil {
		return {{ $t.GoName }}{}, err
	}
{{- else if has_prefix .Type "[]" }}
	{{ short $t }}.{{ .GoName }}, err = mapError(pb.{{ .GoName }}, New{{ slice .Type 2 }})
	if err != nil {
		return {{ $t.GoName }}{}, err
	}
{{- else }}
	{{ short $t }}.{{ .GoName }}, err = New{{ .Type }}(pb.{{ .GoName }})
	if err != nil {
		return {{ $t.GoName }}{}, err
	}
{{- end }}
{{- end }}
{{- end }}
	return {{ short $t }}, nil
}

func ({{ short $t }} {{ $t.GoName }}) PB() (*{{ protoname $t.SQLName }}, error) {
	pb := &{{ protoname $t.SQLName }}{
{{- range fields $t false }}
		{{ .GoName }}: {{ pb $t . }},
{{- end }}
	}

{{- with fields $t true }}
	var err error
{{- range . }}
{{- if eq .Type "[]byte" }}
	err = {{ unmarshaler $t . }}
	if err != nil {
		return nil, err
	}
{{- else if has_prefix .Type "[]" }}
	pb.{{ .GoName }}, err = mapError({{ short $t }}.{{ .GoName }}, {{ slice .Type 2 }}.PB)
	if err != nil {
		return nil, err
	}
{{- else }}
	pb.{{ .GoName }}, err = {{ short $t }}.{{ .GoName }}.PB()
	if err != nil {
		return nil, err
	}
{{- end }}
{{- end }}
{{- end }}
	return pb, nil
}

{{ if $t.PrimaryKeys -}}
type {{ $t.GoName}}Repository interface {
	{{ method $t "Insert" $t.GoName }}
	List{{ plural $t.GoName }}(context.Context, string, int32, string, *ListPosition) (ListStat, []*{{ $t.GoName }}, *ListPosition, error)
{{- range $i := $t.Indexes }}

	// From {{ $i.SQLName }}
	{{ func $t $i }}
{{- if $i.IsUnique }}
{{ $fields := join $i.Fields "" }}
	{{ method $t (printf "Update%sBy" ($t.GoName)) ($fields) }}
	Delete{{ $t.GoName }}By{{ $fields }}(context.Context, {{ params $i.Fields }}) error
{{- end }}
{{- end }}
}
{{- end }}
{{ end }}
