// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/ledgers"
)

// LedgerAccount represents a row from 'ledger_accounts'.
type LedgerAccount struct {
	ID                 int32              `json:"id"`                   // id
	AccountID          string             `json:"account_id"`           // account_id
	Name               string             `json:"name"`                 // name
	Description        string             `json:"description"`          // description
	Amount             []byte             `json:"amount"`               // amount
	Activated          bool               `json:"activated"`            // activated
	AllowManualEntries bool               `json:"allow_manual_entries"` // allow_manual_entries
	GLCode             string             `json:"gl_code"`              // gl_code
	MigrationEventID   string             `json:"migration_event_id"`   // migration_event_id
	StripTrailingZeros bool               `json:"strip_trailing_zeros"` // strip_trailing_zeros
	Type               LedgersAccountType `json:"type"`                 // type
	Usage              LedgersUsage       `json:"usage"`                // usage
	CreateTime         time.Time          `json:"create_time"`          // create_time
	UpdateTime         time.Time          `json:"update_time"`          // update_time
	// xo fields
	Exists, Deleted bool
}

func NewLedgerAccount(pb *ledgers.Account) (LedgerAccount, error) {
	if pb == nil {
		return LedgerAccount{}, ErrNilType{"LedgerAccount"}
	}
	la := LedgerAccount{
		AccountID:          pb.AccountID,
		Name:               pb.Name,
		Description:        pb.Description,
		Activated:          pb.Activated,
		AllowManualEntries: pb.AllowManualEntries,
		GLCode:             pb.GLCode,
		MigrationEventID:   pb.MigrationEventID,
		StripTrailingZeros: pb.StripTrailingZeros,
		Type:               NewLedgersAccountType(pb.Type),
		Usage:              NewLedgersUsage(pb.Usage),
		CreateTime:         pb.CreateTime.AsTime(),
		UpdateTime:         pb.UpdateTime.AsTime(),
	}
	var err error
	la.Amount, err = protojson.Marshal(pb.Amount)
	if err != nil {
		return LedgerAccount{}, err
	}
	return la, nil
}

func (la LedgerAccount) PB() (*ledgers.Account, error) {
	pb := &ledgers.Account{
		AccountID:          la.AccountID,
		Name:               la.Name,
		Description:        la.Description,
		Activated:          la.Activated,
		AllowManualEntries: la.AllowManualEntries,
		GLCode:             la.GLCode,
		MigrationEventID:   la.MigrationEventID,
		StripTrailingZeros: la.StripTrailingZeros,
		Type:               la.Type.PB(),
		Usage:              la.Usage.PB(),
		CreateTime:         timestamppb.New(la.CreateTime),
		UpdateTime:         timestamppb.New(la.UpdateTime),
	}
	var err error
	err = unmarshalMessage(la.Amount, &pb.Amount)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type LedgerAccountRepository interface {
	InsertLedgerAccount(context.Context, *LedgerAccount) error
	ListLedgerAccounts(context.Context, string, int32, string, *ListPosition) (ListStat, []*LedgerAccount, *ListPosition, error)

	// From ledger_accounts_pkey
	LedgerAccountByID(context.Context, int32) (*LedgerAccount, error)

	UpdateLedgerAccountByID(context.Context, *LedgerAccount) error
	DeleteLedgerAccountByID(context.Context, int32) error

	// From ledger_accounts_account_id_idx
	LedgerAccountByAccountID(context.Context, string) (*LedgerAccount, error)

	UpdateLedgerAccountByAccountID(context.Context, *LedgerAccount) error
	DeleteLedgerAccountByAccountID(context.Context, string) error
}