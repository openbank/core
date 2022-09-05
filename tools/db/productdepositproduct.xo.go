// Package db contains generated code from xo.
package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"bnk.to/core/api/v1/products"
)

// ProductDepositProduct represents a row from 'product_deposit_products'.
type ProductDepositProduct struct {
	ID          int32                      `json:"id"`           // id
	ProductID   string                     `json:"product_id"`   // product_id
	Name        string                     `json:"name"`         // name
	Notes       string                     `json:"notes"`        // notes
	Category    ProductsDepositCategory    `json:"category"`     // category
	Currencies  StringSlice                `json:"currencies"`   // currencies
	Controls    []byte                     `json:"controls"`     // controls
	Settings    []byte                     `json:"settings"`     // settings
	Active      bool                       `json:"active"`       // active
	TemplateIDs StringSlice                `json:"template_ids"` // template_ids
	Type        ProductsDepositProductType `json:"type"`         // type
	CreateTime  time.Time                  `json:"create_time"`  // create_time
	UpdateTime  time.Time                  `json:"update_time"`  // update_time
	// xo fields
	Exists, Deleted bool
}

func NewProductDepositProduct(pb *products.DepositProduct) (ProductDepositProduct, error) {
	if pb == nil {
		return ProductDepositProduct{}, ErrNilType{"ProductDepositProduct"}
	}
	pdp := ProductDepositProduct{
		ProductID:   pb.ProductID,
		Name:        pb.Name,
		Notes:       pb.Notes,
		Category:    NewProductsDepositCategory(pb.Category),
		Currencies:  pb.Currencies,
		Active:      pb.Active,
		TemplateIDs: pb.TemplateIDs,
		Type:        NewProductsDepositProductType(pb.Type),
		CreateTime:  pb.CreateTime.AsTime(),
		UpdateTime:  pb.UpdateTime.AsTime(),
	}
	var err error
	pdp.Controls, err = protojson.Marshal(pb.Controls)
	if err != nil {
		return ProductDepositProduct{}, err
	}
	pdp.Settings, err = protojson.Marshal(pb.Settings)
	if err != nil {
		return ProductDepositProduct{}, err
	}
	return pdp, nil
}

func (pdp ProductDepositProduct) PB() (*products.DepositProduct, error) {
	pb := &products.DepositProduct{
		ProductID:   pdp.ProductID,
		Name:        pdp.Name,
		Notes:       pdp.Notes,
		Category:    pdp.Category.PB(),
		Currencies:  pdp.Currencies,
		Active:      pdp.Active,
		TemplateIDs: pdp.TemplateIDs,
		Type:        pdp.Type.PB(),
		CreateTime:  timestamppb.New(pdp.CreateTime),
		UpdateTime:  timestamppb.New(pdp.UpdateTime),
	}
	var err error
	err = unmarshalMessage(pdp.Controls, &pb.Controls)
	if err != nil {
		return nil, err
	}
	err = unmarshalMessage(pdp.Settings, &pb.Settings)
	if err != nil {
		return nil, err
	}
	return pb, nil
}

type ProductDepositProductRepository interface {
	InsertProductDepositProduct(context.Context, *ProductDepositProduct) error
	ListProductDepositProducts(context.Context, string, int32, string, *ListPosition) (ListStat, []*ProductDepositProduct, *ListPosition, error)

	// From product_deposit_products_pkey
	ProductDepositProductByID(context.Context, int32) (*ProductDepositProduct, error)

	UpdateProductDepositProductByID(context.Context, *ProductDepositProduct) error
	DeleteProductDepositProductByID(context.Context, int32) error

	// From product_deposit_products_product_id_idx
	ProductDepositProductByProductID(context.Context, string) (*ProductDepositProduct, error)

	UpdateProductDepositProductByProductID(context.Context, *ProductDepositProduct) error
	DeleteProductDepositProductByProductID(context.Context, string) error
}