/*
 * MSX SDK
 *
 * MSX SDK client.
 *
 * API version: 1.0.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msxsdk

import (
	"encoding/json"
)

// TemplatesPageAllOf struct for TemplatesPageAllOf
type TemplatesPageAllOf struct {
	Contents *[]Template `json:"contents,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplatesPageAllOf TemplatesPageAllOf

// NewTemplatesPageAllOf instantiates a new TemplatesPageAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplatesPageAllOf() *TemplatesPageAllOf {
	this := TemplatesPageAllOf{}
	return &this
}

// NewTemplatesPageAllOfWithDefaults instantiates a new TemplatesPageAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplatesPageAllOfWithDefaults() *TemplatesPageAllOf {
	this := TemplatesPageAllOf{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *TemplatesPageAllOf) GetContents() []Template {
	if o == nil || o.Contents == nil {
		var ret []Template
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplatesPageAllOf) GetContentsOk() (*[]Template, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *TemplatesPageAllOf) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []Template and assigns it to the Contents field.
func (o *TemplatesPageAllOf) SetContents(v []Template) {
	o.Contents = &v
}

func (o TemplatesPageAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplatesPageAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTemplatesPageAllOf := _TemplatesPageAllOf{}

	if err = json.Unmarshal(bytes, &varTemplatesPageAllOf); err == nil {
		*o = TemplatesPageAllOf(varTemplatesPageAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplatesPageAllOf struct {
	value *TemplatesPageAllOf
	isSet bool
}

func (v NullableTemplatesPageAllOf) Get() *TemplatesPageAllOf {
	return v.value
}

func (v *NullableTemplatesPageAllOf) Set(val *TemplatesPageAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplatesPageAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplatesPageAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplatesPageAllOf(val *TemplatesPageAllOf) *NullableTemplatesPageAllOf {
	return &NullableTemplatesPageAllOf{value: val, isSet: true}
}

func (v NullableTemplatesPageAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplatesPageAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

