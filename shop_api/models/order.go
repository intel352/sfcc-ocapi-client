// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Order Document representing an order.
// swagger:model order
type Order struct {

	// The products tax after discounts applying in purchase currency.
	//
	//  Adjusted merchandize prices represent the sum of product prices before
	//  services such as shipping have been added, but after adjustment from
	//  promotions have been added.
	AdjustedMerchandizeTotalTax float64 `json:"adjusted_merchandize_total_tax,omitempty"`

	// The tax of all shipping line items of the line item container after
	//  shipping adjustments have been applied.
	AdjustedShippingTotalTax float64 `json:"adjusted_shipping_total_tax,omitempty"`

	// The billing address. This property is part of basket checkout information only.
	BillingAddress *OrderAddress `json:"billing_address,omitempty"`

	// The bonus discount line items of the line item container.
	BonusDiscountLineItems []*BonusDiscountLineItem `json:"bonus_discount_line_items"`

	// The sales channel for the order.
	// Read Only: true
	// Enum: [storefront callcenter marketplace dss store pinterest twitter facebookads subscriptions onlinereservation customerservicecenter instagramcommerce]
	ChannelType string `json:"channel_type,omitempty"`

	// The confirmation status of the order.
	// Required: true
	// Enum: [not_confirmed confirmed]
	ConfirmationStatus *string `json:"confirmation_status"`

	// The sorted array of coupon items. This array can be empty.
	CouponItems []*CouponItem `json:"coupon_items"`

	// The name of the user who created the order.
	CreatedBy string `json:"created_by,omitempty"`

	// Returns the value of attribute 'creationDate'.
	// Read Only: true
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creation_date,omitempty"`

	// The ISO 4217 mnemonic code of the currency.
	Currency string `json:"currency,omitempty"`

	// The customer information for logged in customers. This property is part of basket checkout information only.
	CustomerInfo *CustomerInfo `json:"customer_info,omitempty"`

	// The name of the customer associated with this order.
	CustomerName string `json:"customer_name,omitempty"`

	// The export status of the order.
	// Required: true
	// Enum: [not_exported exported ready failed]
	ExportStatus *string `json:"export_status"`

	// The external status of the order.
	ExternalOrderStatus string `json:"external_order_status,omitempty"`

	// The sorted array of gift certificate line items. This array can be empty.
	GiftCertificateItems []*GiftCertificateItem `json:"gift_certificate_items"`

	// globalPartyId is managed by Customer 360. Its value can be changed.
	GlobalPartyID string `json:"global_party_id,omitempty"`

	// Returns the value of attribute 'lastModified'.
	// Read Only: true
	// Format: date-time
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// The products total tax in purchase currency.
	//
	//  Merchandize total prices represent the sum of product prices before
	//  services such as shipping or adjustment from promotions have
	//  been added.
	MerchandizeTotalTax float64 `json:"merchandize_total_tax,omitempty"`

	// The notes for the line item container.
	Notes *SimpleLink `json:"notes,omitempty"`

	// The order number of the order.
	// Read Only: true
	OrderNo string `json:"order_no,omitempty"`

	// The array of order level price adjustments. This array can be empty.
	OrderPriceAdjustments []*PriceAdjustment `json:"order_price_adjustments"`

	// The order token used to secure the lookup of an order on base of the
	//  plain order number. The order token contains only URL safe characters.
	OrderToken string `json:"order_token,omitempty"`

	// The total price of the order, including products, shipping and tax. This property is part of basket checkout
	//  information only.
	OrderTotal float64 `json:"order_total,omitempty"`

	// The payment instruments list for the order.
	PaymentInstruments []*OrderPaymentInstrument `json:"payment_instruments"`

	// The payment status of the order.
	// Required: true
	// Enum: [not_paid part_paid paid]
	PaymentStatus *string `json:"payment_status"`

	// The sorted array of product items (up to a maximum of 50 items). This array can be empty.
	ProductItems []*ProductItem `json:"product_items"`

	// The total price of all product items after all product discounts.
	//  Depending on taxation policy the returned price is net or gross.
	ProductSubTotal float64 `json:"product_sub_total,omitempty"`

	// The total price of all product items after all product and order discounts.
	//  Depending on taxation policy the returned price is net or gross.
	ProductTotal float64 `json:"product_total,omitempty"`

	// The array of shipments. This property is part of basket checkout information only.
	Shipments []*Shipment `json:"shipments"`

	// The sorted array of shipping items. This array can be empty.
	ShippingItems []*ShippingItem `json:"shipping_items"`

	// The shipping status of the order.
	// Required: true
	// Enum: [not_shipped part_shipped shipped]
	ShippingStatus *string `json:"shipping_status"`

	// The total shipping price of the order after all shipping discounts. Excludes tax if taxation policy is net. Includes
	//  tax if taxation policy is gross. This property is part of basket checkout information only.
	ShippingTotal float64 `json:"shipping_total,omitempty"`

	// The tax of all shipping line items of the line item container before
	//  shipping adjustments have been applied.
	ShippingTotalTax float64 `json:"shipping_total_tax,omitempty"`

	// The site where the order resides.
	SiteID string `json:"site_id,omitempty"`

	// Gets the source code assigned to this basket.
	SourceCode string `json:"source_code,omitempty"`

	// The status of the order.
	// Required: true
	// Enum: [created new open completed cancelled replaced failed]
	Status *string `json:"status"`

	// The total tax amount of the order. This property is part of basket checkout information only.
	TaxTotal float64 `json:"tax_total,omitempty"`

	// The taxation the line item container is based on.
	// Enum: [gross net]
	Taxation string `json:"taxation,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBonusDiscountLineItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfirmationStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCouponItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGiftCertificateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderPriceAdjustments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentInstruments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Order) validateBillingAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingAddress) { // not required
		return nil
	}

	if m.BillingAddress != nil {
		if err := m.BillingAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("billing_address")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateBonusDiscountLineItems(formats strfmt.Registry) error {

	if swag.IsZero(m.BonusDiscountLineItems) { // not required
		return nil
	}

	for i := 0; i < len(m.BonusDiscountLineItems); i++ {
		if swag.IsZero(m.BonusDiscountLineItems[i]) { // not required
			continue
		}

		if m.BonusDiscountLineItems[i] != nil {
			if err := m.BonusDiscountLineItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bonus_discount_line_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var orderTypeChannelTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["storefront","callcenter","marketplace","dss","store","pinterest","twitter","facebookads","subscriptions","onlinereservation","customerservicecenter","instagramcommerce"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeChannelTypePropEnum = append(orderTypeChannelTypePropEnum, v)
	}
}

const (

	// OrderChannelTypeStorefront captures enum value "storefront"
	OrderChannelTypeStorefront string = "storefront"

	// OrderChannelTypeCallcenter captures enum value "callcenter"
	OrderChannelTypeCallcenter string = "callcenter"

	// OrderChannelTypeMarketplace captures enum value "marketplace"
	OrderChannelTypeMarketplace string = "marketplace"

	// OrderChannelTypeDss captures enum value "dss"
	OrderChannelTypeDss string = "dss"

	// OrderChannelTypeStore captures enum value "store"
	OrderChannelTypeStore string = "store"

	// OrderChannelTypePinterest captures enum value "pinterest"
	OrderChannelTypePinterest string = "pinterest"

	// OrderChannelTypeTwitter captures enum value "twitter"
	OrderChannelTypeTwitter string = "twitter"

	// OrderChannelTypeFacebookads captures enum value "facebookads"
	OrderChannelTypeFacebookads string = "facebookads"

	// OrderChannelTypeSubscriptions captures enum value "subscriptions"
	OrderChannelTypeSubscriptions string = "subscriptions"

	// OrderChannelTypeOnlinereservation captures enum value "onlinereservation"
	OrderChannelTypeOnlinereservation string = "onlinereservation"

	// OrderChannelTypeCustomerservicecenter captures enum value "customerservicecenter"
	OrderChannelTypeCustomerservicecenter string = "customerservicecenter"

	// OrderChannelTypeInstagramcommerce captures enum value "instagramcommerce"
	OrderChannelTypeInstagramcommerce string = "instagramcommerce"
)

// prop value enum
func (m *Order) validateChannelTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeChannelTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateChannelType(formats strfmt.Registry) error {

	if swag.IsZero(m.ChannelType) { // not required
		return nil
	}

	// value enum
	if err := m.validateChannelTypeEnum("channel_type", "body", m.ChannelType); err != nil {
		return err
	}

	return nil
}

var orderTypeConfirmationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_confirmed","confirmed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeConfirmationStatusPropEnum = append(orderTypeConfirmationStatusPropEnum, v)
	}
}

const (

	// OrderConfirmationStatusNotConfirmed captures enum value "not_confirmed"
	OrderConfirmationStatusNotConfirmed string = "not_confirmed"

	// OrderConfirmationStatusConfirmed captures enum value "confirmed"
	OrderConfirmationStatusConfirmed string = "confirmed"
)

// prop value enum
func (m *Order) validateConfirmationStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeConfirmationStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateConfirmationStatus(formats strfmt.Registry) error {

	if err := validate.Required("confirmation_status", "body", m.ConfirmationStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateConfirmationStatusEnum("confirmation_status", "body", *m.ConfirmationStatus); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCouponItems(formats strfmt.Registry) error {

	if swag.IsZero(m.CouponItems) { // not required
		return nil
	}

	for i := 0; i < len(m.CouponItems); i++ {
		if swag.IsZero(m.CouponItems[i]) { // not required
			continue
		}

		if m.CouponItems[i] != nil {
			if err := m.CouponItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("coupon_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Order) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_date", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateCustomerInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerInfo) { // not required
		return nil
	}

	if m.CustomerInfo != nil {
		if err := m.CustomerInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer_info")
			}
			return err
		}
	}

	return nil
}

var orderTypeExportStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_exported","exported","ready","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeExportStatusPropEnum = append(orderTypeExportStatusPropEnum, v)
	}
}

const (

	// OrderExportStatusNotExported captures enum value "not_exported"
	OrderExportStatusNotExported string = "not_exported"

	// OrderExportStatusExported captures enum value "exported"
	OrderExportStatusExported string = "exported"

	// OrderExportStatusReady captures enum value "ready"
	OrderExportStatusReady string = "ready"

	// OrderExportStatusFailed captures enum value "failed"
	OrderExportStatusFailed string = "failed"
)

// prop value enum
func (m *Order) validateExportStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeExportStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateExportStatus(formats strfmt.Registry) error {

	if err := validate.Required("export_status", "body", m.ExportStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateExportStatusEnum("export_status", "body", *m.ExportStatus); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateGiftCertificateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.GiftCertificateItems) { // not required
		return nil
	}

	for i := 0; i < len(m.GiftCertificateItems); i++ {
		if swag.IsZero(m.GiftCertificateItems[i]) { // not required
			continue
		}

		if m.GiftCertificateItems[i] != nil {
			if err := m.GiftCertificateItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gift_certificate_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Order) validateLastModified(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateNotes(formats strfmt.Registry) error {

	if swag.IsZero(m.Notes) { // not required
		return nil
	}

	if m.Notes != nil {
		if err := m.Notes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notes")
			}
			return err
		}
	}

	return nil
}

func (m *Order) validateOrderPriceAdjustments(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderPriceAdjustments) { // not required
		return nil
	}

	for i := 0; i < len(m.OrderPriceAdjustments); i++ {
		if swag.IsZero(m.OrderPriceAdjustments[i]) { // not required
			continue
		}

		if m.OrderPriceAdjustments[i] != nil {
			if err := m.OrderPriceAdjustments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("order_price_adjustments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Order) validatePaymentInstruments(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentInstruments) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentInstruments); i++ {
		if swag.IsZero(m.PaymentInstruments[i]) { // not required
			continue
		}

		if m.PaymentInstruments[i] != nil {
			if err := m.PaymentInstruments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_instruments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var orderTypePaymentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_paid","part_paid","paid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypePaymentStatusPropEnum = append(orderTypePaymentStatusPropEnum, v)
	}
}

const (

	// OrderPaymentStatusNotPaid captures enum value "not_paid"
	OrderPaymentStatusNotPaid string = "not_paid"

	// OrderPaymentStatusPartPaid captures enum value "part_paid"
	OrderPaymentStatusPartPaid string = "part_paid"

	// OrderPaymentStatusPaid captures enum value "paid"
	OrderPaymentStatusPaid string = "paid"
)

// prop value enum
func (m *Order) validatePaymentStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypePaymentStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validatePaymentStatus(formats strfmt.Registry) error {

	if err := validate.Required("payment_status", "body", m.PaymentStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentStatusEnum("payment_status", "body", *m.PaymentStatus); err != nil {
		return err
	}

	return nil
}

func (m *Order) validateProductItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ProductItems); i++ {
		if swag.IsZero(m.ProductItems[i]) { // not required
			continue
		}

		if m.ProductItems[i] != nil {
			if err := m.ProductItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("product_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Order) validateShipments(formats strfmt.Registry) error {

	if swag.IsZero(m.Shipments) { // not required
		return nil
	}

	for i := 0; i < len(m.Shipments); i++ {
		if swag.IsZero(m.Shipments[i]) { // not required
			continue
		}

		if m.Shipments[i] != nil {
			if err := m.Shipments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Order) validateShippingItems(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ShippingItems); i++ {
		if swag.IsZero(m.ShippingItems[i]) { // not required
			continue
		}

		if m.ShippingItems[i] != nil {
			if err := m.ShippingItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shipping_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var orderTypeShippingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["not_shipped","part_shipped","shipped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeShippingStatusPropEnum = append(orderTypeShippingStatusPropEnum, v)
	}
}

const (

	// OrderShippingStatusNotShipped captures enum value "not_shipped"
	OrderShippingStatusNotShipped string = "not_shipped"

	// OrderShippingStatusPartShipped captures enum value "part_shipped"
	OrderShippingStatusPartShipped string = "part_shipped"

	// OrderShippingStatusShipped captures enum value "shipped"
	OrderShippingStatusShipped string = "shipped"
)

// prop value enum
func (m *Order) validateShippingStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeShippingStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateShippingStatus(formats strfmt.Registry) error {

	if err := validate.Required("shipping_status", "body", m.ShippingStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateShippingStatusEnum("shipping_status", "body", *m.ShippingStatus); err != nil {
		return err
	}

	return nil
}

var orderTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["created","new","open","completed","cancelled","replaced","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeStatusPropEnum = append(orderTypeStatusPropEnum, v)
	}
}

const (

	// OrderStatusCreated captures enum value "created"
	OrderStatusCreated string = "created"

	// OrderStatusNew captures enum value "new"
	OrderStatusNew string = "new"

	// OrderStatusOpen captures enum value "open"
	OrderStatusOpen string = "open"

	// OrderStatusCompleted captures enum value "completed"
	OrderStatusCompleted string = "completed"

	// OrderStatusCancelled captures enum value "cancelled"
	OrderStatusCancelled string = "cancelled"

	// OrderStatusReplaced captures enum value "replaced"
	OrderStatusReplaced string = "replaced"

	// OrderStatusFailed captures enum value "failed"
	OrderStatusFailed string = "failed"
)

// prop value enum
func (m *Order) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var orderTypeTaxationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["gross","net"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderTypeTaxationPropEnum = append(orderTypeTaxationPropEnum, v)
	}
}

const (

	// OrderTaxationGross captures enum value "gross"
	OrderTaxationGross string = "gross"

	// OrderTaxationNet captures enum value "net"
	OrderTaxationNet string = "net"
)

// prop value enum
func (m *Order) validateTaxationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, orderTypeTaxationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Order) validateTaxation(formats strfmt.Registry) error {

	if swag.IsZero(m.Taxation) { // not required
		return nil
	}

	// value enum
	if err := m.validateTaxationEnum("taxation", "body", m.Taxation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
