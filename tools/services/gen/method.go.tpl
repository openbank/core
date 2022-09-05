package {{ .PackageName }}

import (
	"context"
	"fmt"

	pb "{{ .GunkPath }}"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"bnk.to/core/tools/db"
	"bnk.to/core/tools/db/mux"
	"bnk.to/core/tools/services"

	{{ range .Imports }}
	{{ .Name }}pb "{{ .Path }}"
	{{ end }}
)

{{ $methodData := . }}

{{/* Generate function header, skipping RequestType and ResponseType as applicable. */ -}}

func (s *{{ .ServerName }}) {{ .MethodName }}(ctx context.Context,

{{- /* Add request type if applicable. */ -}}
{{if eq .RequestType "" -}}
	_ *emptypb.Empty
{{- else -}}
	req *{{ .RequestPkg }}pb.{{ .RequestType }}
{{- end -}}

{{- if eq .ResponseType "" -}}
) (*emptypb.Empty, error) {
{{ else -}}
) (*{{ .ResponsePkg }}pb.{{ .ResponseType}}, error) {
{{ end }}

{{- /* Check for permissions. */ -}}
	if err := s.Auth.CheckPerm(ctx, req, "{{ .PermissionName }}"); err != nil {
		return nil, err
	}

{{if hasPrefix .MethodName "Create" -}}

	storage := mux.Storage(ctx)
	{{- if hasRequestBodyField . "CreateTime" }}
	req.Body.CreateTime = timestamppb.New(time.Now())
	{{- end }}
	{{- if hasRequestBodyField . "UpdateTime" }}
	req.Body.UpdateTime = timestamppb.New(time.Now())
	{{- end }}
	v, err := db.New{{ .ModelName }}(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.Insert{{ .ModelName }}(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil

{{- else if hasPrefix .MethodName "List" -}}

	storage := mux.Storage(ctx)
	var after *db.ListPosition
	if req.PageToken != "" {
		if req.Filter != "" {
			return nil, services.ErrFilterOnPagination
		}
		if req.PageSize != 0 {
			return nil, services.ErrSizeOnPagination
		}
		if req.OrderBy != "" {
			return nil, services.ErrOrderOnPagination
		}
		var err error
		req.Filter, req.PageSize, req.OrderBy, after, err = services.DecodePageToken(req.PageToken)
		if err != nil {
			return nil, err
		}
	}
	if req.PageSize == 0 {
		req.PageSize = services.DefaultPageSize
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		return nil, services.ErrInvalidPageSize
	}
	stat, v, newAfter, err := storage.List{{ .PluralName }}(ctx, req.Filter, req.PageSize, req.OrderBy, after)
	if err != nil {
		return nil, err
	}
	results := make([]*{{ .ResponsePkg }}pb.{{ pbName . }}, 0, len(v))
	for _, w := range v {
		pb, err := w.PB()
		if err != nil {
			return nil, err
		}
		results = append(results, pb)
	}
	return &{{ .ResponsePkg }}pb.{{ .ResponseType }}{
		Total:     stat.Total,
		Remaining: stat.Remaining,
		{{ shortName .PluralName }}: results,
		NextPageToken: services.EncodePageToken(
			req.Filter, req.PageSize, req.OrderBy, *newAfter,
		),
	}, nil

{{- else if hasPrefix .MethodName "Get" -}}

	storage := mux.Storage(ctx)
	v, err := storage.{{ .ModelName }}By{{ uniqueField . }}(ctx, req.{{ uniqueField . }})
	if err != nil {
		return nil, err
	}
	return v.PB()

{{- else if hasPrefix .MethodName "Update" -}}

	storage := mux.Storage(ctx)
	{{- if hasRequestBodyField . "UpdateTime" }}
	// Override UpdateTime to the current time.
	req.Body.UpdateTime = timestamppb.New(time.Now())
	{{- end }}
	v, err := db.New{{ .ModelName }}(req.Body)
	if err != nil {
		return nil, err
	}
	if err := storage.Update{{ .ModelName }}By{{ uniqueField . }}(ctx, &v); err != nil {
		return nil, err
	}
	return req.Body, nil

{{- else if hasPrefix .MethodName "Delete" -}}

	storage := mux.Storage(ctx)
	if err := storage.Delete{{ .ModelName }}By{{ uniqueField . }}(ctx, req.{{ uniqueField . }}); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil

{{- else -}}

	return nil, fmt.Errorf("unimplemented")

{{- end }}
}

