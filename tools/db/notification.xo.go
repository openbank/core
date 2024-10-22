// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/notifications"
)

// Notification represents a row from 'notifications'.
type Notification struct {
	ID               int32                      `json:"id"`                 // id
	NotificationID   string                     `json:"notification_id"`    // notification_id
	UserID           string                     `json:"user_id"`            // user_id
	ClientID         string                     `json:"client_id"`          // client_id
	GroupID          string                     `json:"group_id"`           // group_id
	DepositAccountID string                     `json:"deposit_account_id"` // deposit_account_id
	LoanAccountID    string                     `json:"loan_account_id"`    // loan_account_id
	RepaymentID      string                     `json:"repayment_id"`       // repayment_id
	Type             NotificationsType          `json:"type"`               // type
	Event            NotificationsEvent         `json:"event"`              // event
	Status           Status                     `json:"status"`             // status
	TemplateID       string                     `json:"template_id"`        // template_id
	Destination      string                     `json:"destination"`        // destination
	SenderID         string                     `json:"sender_id"`          // sender_id
	Subject          string                     `json:"subject"`            // subject
	Body             []byte                     `json:"body"`               // body
	NumRetries       int32                      `json:"num_retries"`        // num_retries
	CreateTime       time.Time                  `json:"create_time"`        // create_time
	SendTime         time.Time                  `json:"send_time"`          // send_time
	FailureReason    NotificationsFailureReason `json:"failure_reason"`     // failure_reason
	FailureCause     string                     `json:"failure_cause"`      // failure_cause
	// xo fields
	Exists, Deleted bool
}

func NewNotification(pb *notifications.Notification) (Notification, error) {
	if pb == nil {
		return Notification{}, ErrNilType{"Notification"}
	}
	n := Notification{
		NotificationID:   pb.NotificationID,
		UserID:           pb.UserID,
		ClientID:         pb.ClientID,
		GroupID:          pb.GroupID,
		DepositAccountID: pb.DepositAccountID,
		LoanAccountID:    pb.LoanAccountID,
		RepaymentID:      pb.RepaymentID,
		Type:             NewNotificationsType(pb.Type),
		Event:            NewNotificationsEvent(pb.Event),
		Status:           NewStatus(pb.Status),
		TemplateID:       pb.TemplateID,
		Destination:      pb.Destination,
		SenderID:         pb.SenderID,
		Subject:          pb.Subject,
		Body:             pb.Body,
		NumRetries:       pb.NumRetries,
		CreateTime:       pb.CreateTime.AsTime(),
		SendTime:         pb.SendTime.AsTime(),
		FailureReason:    NewNotificationsFailureReason(pb.FailureReason),
		FailureCause:     pb.FailureCause,
	}
	return n, nil
}

func (n Notification) PB() (*notifications.Notification, error) {
	pb := &notifications.Notification{
		NotificationID:   n.NotificationID,
		UserID:           n.UserID,
		ClientID:         n.ClientID,
		GroupID:          n.GroupID,
		DepositAccountID: n.DepositAccountID,
		LoanAccountID:    n.LoanAccountID,
		RepaymentID:      n.RepaymentID,
		Type:             n.Type.PB(),
		Event:            n.Event.PB(),
		Status:           n.Status.PB(),
		TemplateID:       n.TemplateID,
		Destination:      n.Destination,
		SenderID:         n.SenderID,
		Subject:          n.Subject,
		Body:             n.Body,
		NumRetries:       n.NumRetries,
		CreateTime:       timestamppb.New(n.CreateTime),
		SendTime:         timestamppb.New(n.SendTime),
		FailureReason:    n.FailureReason.PB(),
		FailureCause:     n.FailureCause,
	}
	return pb, nil
}

type NotificationRepository interface {
	InsertNotification(context.Context, *Notification) error
	ListNotifications(context.Context, string, int32, string, *ListPosition) (ListStat, []*Notification, *ListPosition, error)

	// From notifications_pkey
	NotificationByID(context.Context, int32) (*Notification, error)

	UpdateNotificationByID(context.Context, *Notification) error
	DeleteNotificationByID(context.Context, int32) error

	// From notifications_notification_id_idx
	NotificationByNotificationID(context.Context, string) (*Notification, error)

	UpdateNotificationByNotificationID(context.Context, *Notification) error
	DeleteNotificationByNotificationID(context.Context, string) error
}
