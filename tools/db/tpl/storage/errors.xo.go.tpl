{{ define "errordef" -}}
// Error is an error.
type Error string

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// ErrInsertFailed is the insert failed error.
type ErrInsertFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrInsertFailed) Error() string {
	return fmt.Sprintf("insert failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrInsertFailed) Unwrap() error {
	return err.Err
}

// ErrUpdateFailed is the update failed error.
type ErrUpdateFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrUpdateFailed) Error() string {
	return fmt.Sprintf("update failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrUpdateFailed) Unwrap() error {
	return err.Err
}

// ErrUpsertFailed is the upsert failed error.
type ErrUpsertFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrUpsertFailed) Error() string {
	return fmt.Sprintf("upsert failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrUpsertFailed) Unwrap() error {
	return err.Err
}

// ErrDecodeFailed is the decode failed error.
type ErrDecodeFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrDecodeFailed) Error() string {
	return fmt.Sprintf("unable to decode: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrDecodeFailed) Unwrap() error {
	return err.Err
}

// ErrNilType is the nil type error.
type ErrNilType struct {
	typ string
}

func (err ErrNilType) Error() string {
	return fmt.Sprintf("unable to convert nil value for type %q", err.typ)
}
{{- end }}
