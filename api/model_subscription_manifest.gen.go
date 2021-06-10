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

// SubscriptionManifest struct for SubscriptionManifest
type SubscriptionManifest struct {
	Name         string   `json:"name"`
	Mode         string   `json:"mode"`
	Destinations []string `json:"destinations"`
}

// NewSubscriptionManifest instantiates a new SubscriptionManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionManifest(name string, mode string, destinations []string) *SubscriptionManifest {
	this := SubscriptionManifest{}
	this.Name = name
	this.Mode = mode
	this.Destinations = destinations
	return &this
}

// NewSubscriptionManifestWithDefaults instantiates a new SubscriptionManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionManifestWithDefaults() *SubscriptionManifest {
	this := SubscriptionManifest{}
	return &this
}

// GetName returns the Name field value
func (o *SubscriptionManifest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionManifest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionManifest) SetName(v string) {
	o.Name = v
}

// GetMode returns the Mode field value
func (o *SubscriptionManifest) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionManifest) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *SubscriptionManifest) SetMode(v string) {
	o.Mode = v
}

// GetDestinations returns the Destinations field value
func (o *SubscriptionManifest) GetDestinations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
func (o *SubscriptionManifest) GetDestinationsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destinations, true
}

// SetDestinations sets field value
func (o *SubscriptionManifest) SetDestinations(v []string) {
	o.Destinations = v
}

func (o SubscriptionManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["destinations"] = o.Destinations
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionManifest struct {
	value *SubscriptionManifest
	isSet bool
}

func (v NullableSubscriptionManifest) Get() *SubscriptionManifest {
	return v.value
}

func (v *NullableSubscriptionManifest) Set(val *SubscriptionManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionManifest(val *SubscriptionManifest) *NullableSubscriptionManifest {
	return &NullableSubscriptionManifest{value: val, isSet: true}
}

func (v NullableSubscriptionManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}