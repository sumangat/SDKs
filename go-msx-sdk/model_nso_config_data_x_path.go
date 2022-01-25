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

// NSOConfigDataXPath struct for NSOConfigDataXPath
type NSOConfigDataXPath struct {
	ServiceInstanceXPath string `json:"serviceInstanceXPath"`
	ServiceType *string `json:"serviceType,omitempty"`
	PossibleXPathLocations *[]string `json:"possibleXPathLocations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NSOConfigDataXPath NSOConfigDataXPath

// NewNSOConfigDataXPath instantiates a new NSOConfigDataXPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNSOConfigDataXPath(serviceInstanceXPath string) *NSOConfigDataXPath {
	this := NSOConfigDataXPath{}
	this.ServiceInstanceXPath = serviceInstanceXPath
	return &this
}

// NewNSOConfigDataXPathWithDefaults instantiates a new NSOConfigDataXPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNSOConfigDataXPathWithDefaults() *NSOConfigDataXPath {
	this := NSOConfigDataXPath{}
	return &this
}

// GetServiceInstanceXPath returns the ServiceInstanceXPath field value
func (o *NSOConfigDataXPath) GetServiceInstanceXPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceInstanceXPath
}

// GetServiceInstanceXPathOk returns a tuple with the ServiceInstanceXPath field value
// and a boolean to check if the value has been set.
func (o *NSOConfigDataXPath) GetServiceInstanceXPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceInstanceXPath, true
}

// SetServiceInstanceXPath sets field value
func (o *NSOConfigDataXPath) SetServiceInstanceXPath(v string) {
	o.ServiceInstanceXPath = v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *NSOConfigDataXPath) GetServiceType() string {
	if o == nil || o.ServiceType == nil {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSOConfigDataXPath) GetServiceTypeOk() (*string, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *NSOConfigDataXPath) HasServiceType() bool {
	if o != nil && o.ServiceType != nil {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *NSOConfigDataXPath) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetPossibleXPathLocations returns the PossibleXPathLocations field value if set, zero value otherwise.
func (o *NSOConfigDataXPath) GetPossibleXPathLocations() []string {
	if o == nil || o.PossibleXPathLocations == nil {
		var ret []string
		return ret
	}
	return *o.PossibleXPathLocations
}

// GetPossibleXPathLocationsOk returns a tuple with the PossibleXPathLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSOConfigDataXPath) GetPossibleXPathLocationsOk() (*[]string, bool) {
	if o == nil || o.PossibleXPathLocations == nil {
		return nil, false
	}
	return o.PossibleXPathLocations, true
}

// HasPossibleXPathLocations returns a boolean if a field has been set.
func (o *NSOConfigDataXPath) HasPossibleXPathLocations() bool {
	if o != nil && o.PossibleXPathLocations != nil {
		return true
	}

	return false
}

// SetPossibleXPathLocations gets a reference to the given []string and assigns it to the PossibleXPathLocations field.
func (o *NSOConfigDataXPath) SetPossibleXPathLocations(v []string) {
	o.PossibleXPathLocations = &v
}

func (o NSOConfigDataXPath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceInstanceXPath"] = o.ServiceInstanceXPath
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}
	if o.PossibleXPathLocations != nil {
		toSerialize["possibleXPathLocations"] = o.PossibleXPathLocations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NSOConfigDataXPath) UnmarshalJSON(bytes []byte) (err error) {
	varNSOConfigDataXPath := _NSOConfigDataXPath{}

	if err = json.Unmarshal(bytes, &varNSOConfigDataXPath); err == nil {
		*o = NSOConfigDataXPath(varNSOConfigDataXPath)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "serviceInstanceXPath")
		delete(additionalProperties, "serviceType")
		delete(additionalProperties, "possibleXPathLocations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNSOConfigDataXPath struct {
	value *NSOConfigDataXPath
	isSet bool
}

func (v NullableNSOConfigDataXPath) Get() *NSOConfigDataXPath {
	return v.value
}

func (v *NullableNSOConfigDataXPath) Set(val *NSOConfigDataXPath) {
	v.value = val
	v.isSet = true
}

func (v NullableNSOConfigDataXPath) IsSet() bool {
	return v.isSet
}

func (v *NullableNSOConfigDataXPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNSOConfigDataXPath(val *NSOConfigDataXPath) *NullableNSOConfigDataXPath {
	return &NullableNSOConfigDataXPath{value: val, isSet: true}
}

func (v NullableNSOConfigDataXPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNSOConfigDataXPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

