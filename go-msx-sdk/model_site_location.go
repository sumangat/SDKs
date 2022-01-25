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

// SiteLocation struct for SiteLocation
type SiteLocation struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	AdditionalProperties map[string]interface{}
}

type _SiteLocation SiteLocation

// NewSiteLocation instantiates a new SiteLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteLocation(latitude float64, longitude float64) *SiteLocation {
	this := SiteLocation{}
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewSiteLocationWithDefaults instantiates a new SiteLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteLocationWithDefaults() *SiteLocation {
	this := SiteLocation{}
	return &this
}

// GetLatitude returns the Latitude field value
func (o *SiteLocation) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *SiteLocation) GetLatitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *SiteLocation) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *SiteLocation) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *SiteLocation) GetLongitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *SiteLocation) SetLongitude(v float64) {
	o.Longitude = v
}

func (o SiteLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["latitude"] = o.Latitude
	}
	if true {
		toSerialize["longitude"] = o.Longitude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SiteLocation) UnmarshalJSON(bytes []byte) (err error) {
	varSiteLocation := _SiteLocation{}

	if err = json.Unmarshal(bytes, &varSiteLocation); err == nil {
		*o = SiteLocation(varSiteLocation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSiteLocation struct {
	value *SiteLocation
	isSet bool
}

func (v NullableSiteLocation) Get() *SiteLocation {
	return v.value
}

func (v *NullableSiteLocation) Set(val *SiteLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteLocation(val *SiteLocation) *NullableSiteLocation {
	return &NullableSiteLocation{value: val, isSet: true}
}

func (v NullableSiteLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

