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

// TemplateSummaryCore struct for TemplateSummaryCore
type TemplateSummaryCore struct {
	Kind             string                 `json:"kind" yaml:"kind"`
	TemplateMetaName *string                `json:"templateMetaName,omitempty" yaml:"templateMetaName,omitempty"`
	EnvReferences    []TemplateEnvReference `json:"envReferences" yaml:"envReferences"`
}

// NewTemplateSummaryCore instantiates a new TemplateSummaryCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateSummaryCore(kind string, envReferences []TemplateEnvReference) *TemplateSummaryCore {
	this := TemplateSummaryCore{}
	this.Kind = kind
	this.EnvReferences = envReferences
	return &this
}

// NewTemplateSummaryCoreWithDefaults instantiates a new TemplateSummaryCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateSummaryCoreWithDefaults() *TemplateSummaryCore {
	this := TemplateSummaryCore{}
	return &this
}

// GetKind returns the Kind field value
func (o *TemplateSummaryCore) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCore) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TemplateSummaryCore) SetKind(v string) {
	o.Kind = v
}

// GetTemplateMetaName returns the TemplateMetaName field value if set, zero value otherwise.
func (o *TemplateSummaryCore) GetTemplateMetaName() string {
	if o == nil || o.TemplateMetaName == nil {
		var ret string
		return ret
	}
	return *o.TemplateMetaName
}

// GetTemplateMetaNameOk returns a tuple with the TemplateMetaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCore) GetTemplateMetaNameOk() (*string, bool) {
	if o == nil || o.TemplateMetaName == nil {
		return nil, false
	}
	return o.TemplateMetaName, true
}

// HasTemplateMetaName returns a boolean if a field has been set.
func (o *TemplateSummaryCore) HasTemplateMetaName() bool {
	if o != nil && o.TemplateMetaName != nil {
		return true
	}

	return false
}

// SetTemplateMetaName gets a reference to the given string and assigns it to the TemplateMetaName field.
func (o *TemplateSummaryCore) SetTemplateMetaName(v string) {
	o.TemplateMetaName = &v
}

// GetEnvReferences returns the EnvReferences field value
func (o *TemplateSummaryCore) GetEnvReferences() []TemplateEnvReference {
	if o == nil {
		var ret []TemplateEnvReference
		return ret
	}

	return o.EnvReferences
}

// GetEnvReferencesOk returns a tuple with the EnvReferences field value
// and a boolean to check if the value has been set.
func (o *TemplateSummaryCore) GetEnvReferencesOk() (*[]TemplateEnvReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvReferences, true
}

// SetEnvReferences sets field value
func (o *TemplateSummaryCore) SetEnvReferences(v []TemplateEnvReference) {
	o.EnvReferences = v
}

func (o TemplateSummaryCore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.TemplateMetaName != nil {
		toSerialize["templateMetaName"] = o.TemplateMetaName
	}
	if true {
		toSerialize["envReferences"] = o.EnvReferences
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateSummaryCore struct {
	value *TemplateSummaryCore
	isSet bool
}

func (v NullableTemplateSummaryCore) Get() *TemplateSummaryCore {
	return v.value
}

func (v *NullableTemplateSummaryCore) Set(val *TemplateSummaryCore) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateSummaryCore) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateSummaryCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateSummaryCore(val *TemplateSummaryCore) *NullableTemplateSummaryCore {
	return &NullableTemplateSummaryCore{value: val, isSet: true}
}

func (v NullableTemplateSummaryCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateSummaryCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}