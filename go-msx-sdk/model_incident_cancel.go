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

// IncidentCancel struct for IncidentCancel
type IncidentCancel struct {
	Notes *string `json:"notes,omitempty"`
	Tenant string `json:"tenant"`
	AdditionalProperties map[string]interface{}
}

type _IncidentCancel IncidentCancel

// NewIncidentCancel instantiates a new IncidentCancel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncidentCancel(tenant string) *IncidentCancel {
	this := IncidentCancel{}
	this.Tenant = tenant
	return &this
}

// NewIncidentCancelWithDefaults instantiates a new IncidentCancel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentCancelWithDefaults() *IncidentCancel {
	this := IncidentCancel{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *IncidentCancel) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCancel) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *IncidentCancel) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *IncidentCancel) SetNotes(v string) {
	o.Notes = &v
}

// GetTenant returns the Tenant field value
func (o *IncidentCancel) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *IncidentCancel) GetTenantOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *IncidentCancel) SetTenant(v string) {
	o.Tenant = v
}

func (o IncidentCancel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncidentCancel) UnmarshalJSON(bytes []byte) (err error) {
	varIncidentCancel := _IncidentCancel{}

	if err = json.Unmarshal(bytes, &varIncidentCancel); err == nil {
		*o = IncidentCancel(varIncidentCancel)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "notes")
		delete(additionalProperties, "tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncidentCancel struct {
	value *IncidentCancel
	isSet bool
}

func (v NullableIncidentCancel) Get() *IncidentCancel {
	return v.value
}

func (v *NullableIncidentCancel) Set(val *IncidentCancel) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentCancel) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentCancel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentCancel(val *IncidentCancel) *NullableIncidentCancel {
	return &NullableIncidentCancel{value: val, isSet: true}
}

func (v NullableIncidentCancel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentCancel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

