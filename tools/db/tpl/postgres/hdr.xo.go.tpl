{{ define "header" }}
// Package {{ pkg }} contains generated code from xo.
package {{ pkg }}

// Code generated by xo. DO NOT EDIT.

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"strings"
	"time"

	"bnk.to/core/tools/db"
)
{{ end }}
