// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"

	"bnk.to/core/api/v1/fields"
)

// Field represents a row from 'fields'.
type Field struct {
	ID             int32       `json:"id"`              // id
	FieldID        string      `json:"field_id"`        // field_id
	Type           FieldsType  `json:"type"`            // type
	Name           string      `json:"name"`            // name
	Description    string      `json:"description"`     // description
	Options        StringSlice `json:"options"`         // options
	Usage          []byte      `json:"usage"`           // usage
	ValidationRule []byte      `json:"validation_rule"` // validation_rule
	ViewRights     StringSlice `json:"view_rights"`     // view_rights
	EditRights     StringSlice `json:"edit_rights"`     // edit_rights
	// xo fields
	Exists, Deleted bool
}

func NewField(pb *fields.Field) (Field, error) {
	if pb == nil {
		return Field{}, ErrNilType{"Field"}
	}
	f := Field{
		FieldID:     pb.FieldID,
		Type:        NewFieldsType(pb.Type),
		Name:        pb.Name,
		Description: pb.Description,
		Options:     pb.Options,
		ViewRights:  pb.ViewRights,
		EditRights:  pb.EditRights,
	}
	var err error
	f.Usage, err = protojson.Marshal(pb.Usage)
	if err != nil {
		return Field{}, err
	}
	f.ValidationRule, err = protojson.Marshal(pb.ValidationRule)
	if err != nil {
		return Field{}, err
	}
	return f, nil
}

func (f Field) PB() (*fields.Field, error) {
	pb := &fields.Field{
		FieldID:     f.FieldID,
		Type:        f.Type.PB(),
		Name:        f.Name,
		Description: f.Description,
		Options:     f.Options,
		ViewRights:  f.ViewRights,
		EditRights:  f.EditRights,
	}
	var err error
	err = unmarshalMessage(f.Usage, &pb.Usage)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(f.ValidationRule, &pb.ValidationRule)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type FieldRepository interface {
	InsertField(context.Context, *Field) error
	ListFields(context.Context, string, int32, string, *ListPosition) (ListStat, []*Field, *ListPosition, error)

	// From fields_pkey
	FieldByID(context.Context, int32) (*Field, error)

	UpdateFieldByID(context.Context, *Field) error
	DeleteFieldByID(context.Context, int32) error
}
