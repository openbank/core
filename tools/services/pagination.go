package services

import (
	"errors"
	"fmt"

	"bnk.to/core/tools/db"
)

// DefaultPageSize is the default page size if the page size was not provided
// (set to 0).
const DefaultPageSize = 25

var (
	ErrFilterOnPagination = fmt.Errorf("filter must not be provided when using a page token")
	ErrSizeOnPagination   = fmt.Errorf("page_size must not be provided when using a page token")
	ErrOrderOnPagination  = fmt.Errorf("order_by must not be provided when using a page token")
	ErrInvalidPageSize    = fmt.Errorf("page size must be between 1-100")
)

// EncodePageToken encodes the provided values into a single page token value
// which can be used to retrieve the values in the future.
func EncodePageToken(query string, pageSize int32, orderBy string, after db.ListPosition) string {
	// TODO: Implement page tokens.
	return "unimplemented"
}

// DecodePageToken decodes the provided token into the values passed to EncodePageToken.
func DecodePageToken(token string) (query string, pageSize int32, orderBy string, after *db.ListPosition, err error) {
	return "", 0, "", nil, errors.New("page token is unimplemented")
}
