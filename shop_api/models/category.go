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

// Category Document representing a category.
// swagger:model category
type Category struct {

	// Renders an alternative URL in main navigation. Uses Salesforce Commerce Cloud content url notation. For example: $url('Account-Show')$ or normal URL http://xchange.demandware.com
	CAlternativeURL string `json:"c_alternativeUrl,omitempty"`

	// Used to define the content asset used to populate a grid page banner for a category. This value is applied to all sub-category navigation (cascading) if no specific catBannerID has been defined for  a sub-category.
	CCatBannerID string `json:"c_catBannerID,omitempty"`

	// Use this attribute to apply custom styles for this category.
	CCustomCSSFile string `json:"c_customCSSFile,omitempty"`

	// Used to define if/when the Compare feature is to be visualized in the storefront based on navigation. If enableCompare = FALSE, no Compare checkboxes will be displayed in the grid view. If enableCompare = TRUE, the category (and its children) will support the Compare feature.
	CEnableCompare bool `json:"c_enableCompare,omitempty"`

	// c header menu banner
	CHeaderMenuBanner string `json:"c_headerMenuBanner,omitempty"`

	// Which way to orient the menu and optional header menu HTML. Vertical will list all in one line. Horizontal will list in columns.
	// Enum: [Horizontal Vertical]
	CHeaderMenuOrientation string `json:"c_headerMenuOrientation,omitempty"`

	// Used to indicate that a category (such as Mens -> Footwear -> Boots) will display in the roll-over navigation. A sub-category only shows if also the parent category is marked as showInMenu. Up to three category levels are shown in roll-over navigation.
	CShowInMenu bool `json:"c_showInMenu,omitempty"`

	// Used to define the content asset ID of the Size Chart that is appropriate for products whose PRIMARY category is the associated category (and its children). Whenever a product detail page (or quick view) is rendered, the Size Chart link is populated based on the value of this attribute for the products primary categorization. If not defined, NO size chart link is displayed.
	CSizeChartID string `json:"c_sizeChartID,omitempty"`

	// c slot banner Html
	CSlotBannerHTML string `json:"c_slotBannerHtml,omitempty"`

	// Image used on either the top or bottom slot on the category landing pages.
	CSlotBannerImage string `json:"c_slotBannerImage,omitempty"`

	// Array of subcategories. Can be empty.
	Categories []*Category `json:"categories"`

	// The localized description of the category.
	Description string `json:"description,omitempty"`

	// The id of the category.
	// Required: true
	ID *string `json:"id"`

	// The URL to the category image.
	Image string `json:"image,omitempty"`

	// The localized name of the category.
	Name string `json:"name,omitempty"`

	// The localized page description of the category.
	PageDescription string `json:"page_description,omitempty"`

	// The localized page keywords of the category.
	PageKeywords string `json:"page_keywords,omitempty"`

	// The localized page title of the category.
	PageTitle string `json:"page_title,omitempty"`

	// The id of the parent category.
	ParentCategoryID string `json:"parent_category_id,omitempty"`

	// The URL to the category thumbnail.
	Thumbnail string `json:"thumbnail,omitempty"`
}

// Validate validates this category
func (m *Category) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCHeaderMenuOrientation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var categoryTypeCHeaderMenuOrientationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Horizontal","Vertical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		categoryTypeCHeaderMenuOrientationPropEnum = append(categoryTypeCHeaderMenuOrientationPropEnum, v)
	}
}

const (

	// CategoryCHeaderMenuOrientationHorizontal captures enum value "Horizontal"
	CategoryCHeaderMenuOrientationHorizontal string = "Horizontal"

	// CategoryCHeaderMenuOrientationVertical captures enum value "Vertical"
	CategoryCHeaderMenuOrientationVertical string = "Vertical"
)

// prop value enum
func (m *Category) validateCHeaderMenuOrientationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, categoryTypeCHeaderMenuOrientationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Category) validateCHeaderMenuOrientation(formats strfmt.Registry) error {

	if swag.IsZero(m.CHeaderMenuOrientation) { // not required
		return nil
	}

	// value enum
	if err := m.validateCHeaderMenuOrientationEnum("c_headerMenuOrientation", "body", m.CHeaderMenuOrientation); err != nil {
		return err
	}

	return nil
}

func (m *Category) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {
		if swag.IsZero(m.Categories[i]) { // not required
			continue
		}

		if m.Categories[i] != nil {
			if err := m.Categories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("categories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Category) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Category) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Category) UnmarshalBinary(b []byte) error {
	var res Category
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
