// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddressPair Addressing information of an endpoint
//
// swagger:model AddressPair
type AddressPair struct {

	// IPv4 address
	IPV4 string `json:"ipv4,omitempty"`

	// UUID of IPv4 expiration timer
	IPV4ExpirationUUID string `json:"ipv4-expiration-uuid,omitempty"`

	// IPAM pool from which this IPv4 address was allocated
	IPV4PoolName string `json:"ipv4-pool-name,omitempty"`

	// IPv6 address
	IPV6 string `json:"ipv6,omitempty"`

	// UUID of IPv6 expiration timer
	IPV6ExpirationUUID string `json:"ipv6-expiration-uuid,omitempty"`

	// IPAM pool from which this IPv6 address was allocated
	IPV6PoolName string `json:"ipv6-pool-name,omitempty"`
}

// Validate validates this address pair
func (m *AddressPair) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this address pair based on context it is used
func (m *AddressPair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddressPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressPair) UnmarshalBinary(b []byte) error {
	var res AddressPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
