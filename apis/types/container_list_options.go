// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerListOptions options of list container, filters (a `map[string][]string`) to process on the container list. Available filters:
//
// - `id=container-id`
// - `name=container-name`
// - `status=running`
// - `label=key` or `label=key=value`
// - `network=container-network`
// - `volume=volume-id`
//
// swagger:model ContainerListOptions

type ContainerListOptions struct {

	// all
	All bool `json:"All,omitempty"`

	// before
	Before string `json:"Before,omitempty"`

	// filter
	Filter map[string][]string `json:"Filter,omitempty"`

	// limit
	Limit int64 `json:"Limit,omitempty"`

	// since
	Since string `json:"Since,omitempty"`
}

/* polymorph ContainerListOptions All false */

/* polymorph ContainerListOptions Before false */

/* polymorph ContainerListOptions Filter false */

/* polymorph ContainerListOptions Limit false */

/* polymorph ContainerListOptions Since false */

// Validate validates this container list options
func (m *ContainerListOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerListOptions) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerListOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerListOptions) UnmarshalBinary(b []byte) error {
	var res ContainerListOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
