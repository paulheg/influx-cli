/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UnsignedIntegerLiteral Represents integer numbers
type UnsignedIntegerLiteral struct {
	// Type of AST node
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewUnsignedIntegerLiteral instantiates a new UnsignedIntegerLiteral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsignedIntegerLiteral() *UnsignedIntegerLiteral {
	this := UnsignedIntegerLiteral{}
	return &this
}

// NewUnsignedIntegerLiteralWithDefaults instantiates a new UnsignedIntegerLiteral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsignedIntegerLiteralWithDefaults() *UnsignedIntegerLiteral {
	this := UnsignedIntegerLiteral{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnsignedIntegerLiteral) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsignedIntegerLiteral) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnsignedIntegerLiteral) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UnsignedIntegerLiteral) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UnsignedIntegerLiteral) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsignedIntegerLiteral) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UnsignedIntegerLiteral) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UnsignedIntegerLiteral) SetValue(v string) {
	o.Value = &v
}

func (o UnsignedIntegerLiteral) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableUnsignedIntegerLiteral struct {
	value *UnsignedIntegerLiteral
	isSet bool
}

func (v NullableUnsignedIntegerLiteral) Get() *UnsignedIntegerLiteral {
	return v.value
}

func (v *NullableUnsignedIntegerLiteral) Set(val *UnsignedIntegerLiteral) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsignedIntegerLiteral) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsignedIntegerLiteral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsignedIntegerLiteral(val *UnsignedIntegerLiteral) *NullableUnsignedIntegerLiteral {
	return &NullableUnsignedIntegerLiteral{value: val, isSet: true}
}

func (v NullableUnsignedIntegerLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsignedIntegerLiteral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}