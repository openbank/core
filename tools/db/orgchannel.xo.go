// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"google.golang.org/protobuf/encoding/protojson"

	"bnk.to/core/api/v1/org"
)

// OrgChannel represents a row from 'org_channels'.
type OrgChannel struct {
	ID                 int32       `json:"id"`                  // id
	ChannelID          string      `json:"channel_id"`          // channel_id
	GLAccount          string      `json:"gl_account"`          // gl_account
	Name               string      `json:"name"`                // name
	LoansConstraints   []byte      `json:"loans_constraints"`   // loans_constraints
	DepositConstraints []byte      `json:"deposit_constraints"` // deposit_constraints
	State              State       `json:"state"`               // state
	Default            bool        `json:"default"`             // default
	AllowedUsers       StringSlice `json:"allowed_users"`       // allowed_users
	// xo fields
	Exists, Deleted bool
}

func NewOrgChannel(pb *org.Channel) (OrgChannel, error) {
	if pb == nil {
		return OrgChannel{}, ErrNilType{"OrgChannel"}
	}
	oc := OrgChannel{
		ChannelID:    pb.ChannelID,
		GLAccount:    pb.GLAccount,
		Name:         pb.Name,
		State:        NewState(pb.State),
		Default:      pb.Default,
		AllowedUsers: pb.AllowedUsers,
	}
	var err error
	oc.LoansConstraints, err = protojson.Marshal(pb.LoansConstraints)
	if err != nil {
		return OrgChannel{}, err
	}
	oc.DepositConstraints, err = protojson.Marshal(pb.DepositConstraints)
	if err != nil {
		return OrgChannel{}, err
	}
	return oc, nil
}

func (oc OrgChannel) PB() (*org.Channel, error) {
	pb := &org.Channel{
		ChannelID:    oc.ChannelID,
		GLAccount:    oc.GLAccount,
		Name:         oc.Name,
		State:        oc.State.PB(),
		Default:      oc.Default,
		AllowedUsers: oc.AllowedUsers,
	}
	var err error
	err = unmarshalMessage(oc.LoansConstraints, &pb.LoansConstraints)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(oc.DepositConstraints, &pb.DepositConstraints)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type OrgChannelRepository interface {
	InsertOrgChannel(context.Context, *OrgChannel) error
	ListOrgChannels(context.Context, string, int32, string, *ListPosition) (ListStat, []*OrgChannel, *ListPosition, error)

	// From org_channels_pkey
	OrgChannelByID(context.Context, int32) (*OrgChannel, error)

	UpdateOrgChannelByID(context.Context, *OrgChannel) error
	DeleteOrgChannelByID(context.Context, int32) error

	// From org_channels_channel_id_idx
	OrgChannelByChannelID(context.Context, string) (*OrgChannel, error)

	UpdateOrgChannelByChannelID(context.Context, *OrgChannel) error
	DeleteOrgChannelByChannelID(context.Context, string) error
}
