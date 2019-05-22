// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductInventoryRecord Document representing a product inventory record.
// swagger:model product_inventory_record
type ProductInventoryRecord struct {

	// The allocation quantity and reset date.
	Allocation *ProductInventoryRecordAllocation `json:"allocation,omitempty"`

	// The quantity of items available to sell (ATS). This is calculated as the
	//  allocation plus the preorderBackorderAllocation minus the turnover.
	Ats float64 `json:"ats,omitempty"`

	// Returns the value of attribute 'creationDate'.
	// Read Only: true
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creation_date,omitempty"`

	// The date that the item is expected to be in stock.
	// Format: date-time
	InStockDate strfmt.DateTime `json:"in_stock_date,omitempty"`

	// The user supplied ID of the inventory list.
	InventoryListID string `json:"inventory_list_id,omitempty"`

	// The sum of all inventory transactions (decrements and increments)
	//  that have been recorded subsequent to the allocation was reset date.
	//  The quantity value can be negative due to higher quantity of inventory decrements than increments.
	InventoryTurnover float64 `json:"inventory_turnover,omitempty"`

	// Returns the value of attribute 'lastModified'.
	// Read Only: true
	// Format: date-time
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// The URL that is used to get this instance.
	Link string `json:"link,omitempty"`

	// The flag that determines if the product is perpetually in stock.
	PerpetualFlag bool `json:"perpetual_flag,omitempty"`

	// The quantity of items that are allocated for sale, beyond the initial stock allocation.
	PreOrderBackOrderAllocation float64 `json:"pre_order_back_order_allocation,omitempty"`

	// The enum holding the records pre-backorder-handling configuration.
	//  Possible values are NONE, PREORDER and BACKORDER.
	//  Method returns NONE in case the record pre-backorder-handling-code is null or unknown.
	// Enum: [none preorder backorder]
	PreOrderBackOrderHandling string `json:"pre_order_back_order_handling,omitempty"`

	// The user supplied ID of the product.
	// Max Length: 256
	// Min Length: 1
	ProductID string `json:"product_id,omitempty"`

	// The name of the product.
	ProductName string `json:"product_name,omitempty"`

	// The on order quantity, the quantity of all transactions for
	//  this record since the allocation reset date.
	QuantityOnOrder float64 `json:"quantity_on_order,omitempty"`

	// The current stock level. This is calculated as the allocation minus the turnover.
	StockLevel float64 `json:"stock_level,omitempty"`
}

// Validate validates this product inventory record
func (m *ProductInventoryRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInStockDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreOrderBackOrderHandling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductInventoryRecord) validateAllocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Allocation) { // not required
		return nil
	}

	if m.Allocation != nil {
		if err := m.Allocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocation")
			}
			return err
		}
	}

	return nil
}

func (m *ProductInventoryRecord) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_date", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductInventoryRecord) validateInStockDate(formats strfmt.Registry) error {

	if swag.IsZero(m.InStockDate) { // not required
		return nil
	}

	if err := validate.FormatOf("in_stock_date", "body", "date-time", m.InStockDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductInventoryRecord) validateLastModified(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

var productInventoryRecordTypePreOrderBackOrderHandlingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","preorder","backorder"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		productInventoryRecordTypePreOrderBackOrderHandlingPropEnum = append(productInventoryRecordTypePreOrderBackOrderHandlingPropEnum, v)
	}
}

const (

	// ProductInventoryRecordPreOrderBackOrderHandlingNone captures enum value "none"
	ProductInventoryRecordPreOrderBackOrderHandlingNone string = "none"

	// ProductInventoryRecordPreOrderBackOrderHandlingPreorder captures enum value "preorder"
	ProductInventoryRecordPreOrderBackOrderHandlingPreorder string = "preorder"

	// ProductInventoryRecordPreOrderBackOrderHandlingBackorder captures enum value "backorder"
	ProductInventoryRecordPreOrderBackOrderHandlingBackorder string = "backorder"
)

// prop value enum
func (m *ProductInventoryRecord) validatePreOrderBackOrderHandlingEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, productInventoryRecordTypePreOrderBackOrderHandlingPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProductInventoryRecord) validatePreOrderBackOrderHandling(formats strfmt.Registry) error {

	if swag.IsZero(m.PreOrderBackOrderHandling) { // not required
		return nil
	}

	// value enum
	if err := m.validatePreOrderBackOrderHandlingEnum("pre_order_back_order_handling", "body", m.PreOrderBackOrderHandling); err != nil {
		return err
	}

	return nil
}

func (m *ProductInventoryRecord) validateProductID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductID) { // not required
		return nil
	}

	if err := validate.MinLength("product_id", "body", string(m.ProductID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("product_id", "body", string(m.ProductID), 256); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductInventoryRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductInventoryRecord) UnmarshalBinary(b []byte) error {
	var res ProductInventoryRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}