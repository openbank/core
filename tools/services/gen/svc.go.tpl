package {{ .PackageName }}

import (
	"bnk.to/core/tools/services"

	pb "{{ .GunkPath }}"
)

{{ $serviceNames := .ServiceNames }}

{{ range $i, $serverName := .ServerNames -}}
// {{ $serverName }} is the implementation of {{ index $serviceNames $i }}.
type {{ $serverName }} struct {
	pb.Unsafe{{ index $serviceNames $i }}Server
	services.Common
}
{{ end }}

{{ range $serverName := .ServerNames -}}
// New{{ $serverName }} creates a new {{ $serverName }}.
func New{{ $serverName }}(common services.Common) *{{ $serverName }} {
	return &{{ $serverName }}{
		Common: common,
	}
}
{{ end }}
