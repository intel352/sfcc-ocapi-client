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

// Site Document representing a site.
// swagger:model site
type Site struct {

	// The list of allowed currencies.
	AllowedCurrencies []string `json:"allowed_currencies"`

	// A list of all allowed site locales.
	AllowedLocales []*Locale `json:"allowed_locales"`

	// The currency mnemonic of the site.
	DefaultCurrency string `json:"default_currency,omitempty"`

	// The default locale of the site.
	DefaultLocale string `json:"default_locale,omitempty"`

	// The HTTP DIS base URL.
	HTTPDisBaseURL string `json:"http_dis_base_url,omitempty"`

	// The configured HTTP host name. If no host name is configured the instance host name is returned.
	HTTPHostname string `json:"http_hostname,omitempty"`

	// The HTTP URL to the library content location of the site.
	HTTPLibraryContentURL string `json:"http_library_content_url,omitempty"`

	// The HTTP URL to the site content location.
	HTTPSiteContentURL string `json:"http_site_content_url,omitempty"`

	// The HTTPS DIS base URL.
	HTTPSDisBaseURL string `json:"https_dis_base_url,omitempty"`

	// The configured HTTPS host name. If no host name is configured the instance host name is returned.
	HTTPSHostname string `json:"https_hostname,omitempty"`

	// The HTTPS URL to the library content location of the site.
	HTTPSLibraryContentURL string `json:"https_library_content_url,omitempty"`

	// The HTTPS URL to the site content location.
	HTTPSSiteContentURL string `json:"https_site_content_url,omitempty"`

	// The id of the site.
	ID string `json:"id,omitempty"`

	// The descriptive name for the site.
	Name string `json:"name,omitempty"`

	// The site status online/offline.
	// Enum: [online offline]
	Status string `json:"status,omitempty"`

	// The time zone of the site (for example, USA/Eastern).
	Timezone string `json:"timezone,omitempty"`

	// The time zone offset from UTC for the current time in milliseconds (for example, -14400000).
	TimezoneOffset int32 `json:"timezone_offset,omitempty"`
}

// Validate validates this site
func (m *Site) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedLocales(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Site) validateAllowedLocales(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedLocales) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedLocales); i++ {
		if swag.IsZero(m.AllowedLocales[i]) { // not required
			continue
		}

		if m.AllowedLocales[i] != nil {
			if err := m.AllowedLocales[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowed_locales" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var siteTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		siteTypeStatusPropEnum = append(siteTypeStatusPropEnum, v)
	}
}

const (

	// SiteStatusOnline captures enum value "online"
	SiteStatusOnline string = "online"

	// SiteStatusOffline captures enum value "offline"
	SiteStatusOffline string = "offline"
)

// prop value enum
func (m *Site) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, siteTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Site) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Site) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Site) UnmarshalBinary(b []byte) error {
	var res Site
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}