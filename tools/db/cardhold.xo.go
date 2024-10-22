// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/cards"
)

// CardHold represents a row from 'card_holds'.
type CardHold struct {
	ID                  int32           `json:"id"`                    // id
	HoldID              string          `json:"hold_id"`               // hold_id
	AccountID           string          `json:"account_id"`            // account_id
	Advice              bool            `json:"advice"`                // advice
	Amount              []byte          `json:"amount"`                // amount
	OriginalAmount      []byte          `json:"original_amount"`       // original_amount
	CardAcceptor        []byte          `json:"card_acceptor"`         // card_acceptor
	ExchangeRate        float64         `json:"exchange_rate"`         // exchange_rate
	CardID              string          `json:"card_id"`               // card_id
	Type                CardsType       `json:"type"`                  // type
	Source              CardsHoldSource `json:"source"`                // source
	Status              CardsHoldStatus `json:"status"`                // status
	UserTransactionTime time.Time       `json:"user_transaction_time"` // user_transaction_time
	// xo fields
	Exists, Deleted bool
}

func NewCardHold(pb *cards.Hold) (CardHold, error) {
	if pb == nil {
		return CardHold{}, ErrNilType{"CardHold"}
	}
	ch := CardHold{
		HoldID:              pb.HoldID,
		AccountID:           pb.AccountID,
		Advice:              pb.Advice,
		ExchangeRate:        pb.ExchangeRate,
		CardID:              pb.CardID,
		Type:                NewCardsType(pb.Type),
		Source:              NewCardsHoldSource(pb.Source),
		Status:              NewCardsHoldStatus(pb.Status),
		UserTransactionTime: pb.UserTransactionTime.AsTime(),
	}
	var err error
	ch.Amount, err = protojson.Marshal(pb.Amount)
	if err != nil {
		return CardHold{}, err
	}
	ch.OriginalAmount, err = protojson.Marshal(pb.OriginalAmount)
	if err != nil {
		return CardHold{}, err
	}
	ch.CardAcceptor, err = protojson.Marshal(pb.CardAcceptor)
	if err != nil {
		return CardHold{}, err
	}
	return ch, nil
}

func (ch CardHold) PB() (*cards.Hold, error) {
	pb := &cards.Hold{
		HoldID:              ch.HoldID,
		AccountID:           ch.AccountID,
		Advice:              ch.Advice,
		ExchangeRate:        ch.ExchangeRate,
		CardID:              ch.CardID,
		Type:                ch.Type.PB(),
		Source:              ch.Source.PB(),
		Status:              ch.Status.PB(),
		UserTransactionTime: timestamppb.New(ch.UserTransactionTime),
	}
	var err error
	err = unmarshalMessage(ch.Amount, &pb.Amount)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(ch.OriginalAmount, &pb.OriginalAmount)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(ch.CardAcceptor, &pb.CardAcceptor)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type CardHoldRepository interface {
	InsertCardHold(context.Context, *CardHold) error
	ListCardHolds(context.Context, string, int32, string, *ListPosition) (ListStat, []*CardHold, *ListPosition, error)

	// From card_holds_pkey
	CardHoldByID(context.Context, int32) (*CardHold, error)

	UpdateCardHoldByID(context.Context, *CardHold) error
	DeleteCardHoldByID(context.Context, int32) error

	// From card_holds_hold_id_idx
	CardHoldByHoldID(context.Context, string) (*CardHold, error)

	UpdateCardHoldByHoldID(context.Context, *CardHold) error
	DeleteCardHoldByHoldID(context.Context, string) error
}
