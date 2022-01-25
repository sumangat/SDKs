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
	"time"
)

// DeviceAllOf struct for DeviceAllOf
type DeviceAllOf struct {
	Id *string `json:"id,omitempty"`
	UserId *string `json:"userId,omitempty"`
	ProviderId *string `json:"providerId,omitempty"`
	VulnerabilityState *DeviceVulnerabilityState `json:"vulnerabilityState,omitempty"`
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	ModifiedOn NullableTime `json:"modifiedOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceAllOf DeviceAllOf

// NewDeviceAllOf instantiates a new DeviceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAllOf() *DeviceAllOf {
	this := DeviceAllOf{}
	return &this
}

// NewDeviceAllOfWithDefaults instantiates a new DeviceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAllOfWithDefaults() *DeviceAllOf {
	this := DeviceAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceAllOf) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DeviceAllOf) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DeviceAllOf) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *DeviceAllOf) SetUserId(v string) {
	o.UserId = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *DeviceAllOf) GetProviderId() string {
	if o == nil || o.ProviderId == nil {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetProviderIdOk() (*string, bool) {
	if o == nil || o.ProviderId == nil {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *DeviceAllOf) HasProviderId() bool {
	if o != nil && o.ProviderId != nil {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *DeviceAllOf) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetVulnerabilityState returns the VulnerabilityState field value if set, zero value otherwise.
func (o *DeviceAllOf) GetVulnerabilityState() DeviceVulnerabilityState {
	if o == nil || o.VulnerabilityState == nil {
		var ret DeviceVulnerabilityState
		return ret
	}
	return *o.VulnerabilityState
}

// GetVulnerabilityStateOk returns a tuple with the VulnerabilityState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetVulnerabilityStateOk() (*DeviceVulnerabilityState, bool) {
	if o == nil || o.VulnerabilityState == nil {
		return nil, false
	}
	return o.VulnerabilityState, true
}

// HasVulnerabilityState returns a boolean if a field has been set.
func (o *DeviceAllOf) HasVulnerabilityState() bool {
	if o != nil && o.VulnerabilityState != nil {
		return true
	}

	return false
}

// SetVulnerabilityState gets a reference to the given DeviceVulnerabilityState and assigns it to the VulnerabilityState field.
func (o *DeviceAllOf) SetVulnerabilityState(v DeviceVulnerabilityState) {
	o.VulnerabilityState = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *DeviceAllOf) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAllOf) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *DeviceAllOf) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *DeviceAllOf) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceAllOf) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceAllOf) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *DeviceAllOf) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn.IsSet() {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given NullableTime and assigns it to the ModifiedOn field.
func (o *DeviceAllOf) SetModifiedOn(v time.Time) {
	o.ModifiedOn.Set(&v)
}
// SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil
func (o *DeviceAllOf) SetModifiedOnNil() {
	o.ModifiedOn.Set(nil)
}

// UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
func (o *DeviceAllOf) UnsetModifiedOn() {
	o.ModifiedOn.Unset()
}

func (o DeviceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.ProviderId != nil {
		toSerialize["providerId"] = o.ProviderId
	}
	if o.VulnerabilityState != nil {
		toSerialize["vulnerabilityState"] = o.VulnerabilityState
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.ModifiedOn.IsSet() {
		toSerialize["modifiedOn"] = o.ModifiedOn.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceAllOf := _DeviceAllOf{}

	if err = json.Unmarshal(bytes, &varDeviceAllOf); err == nil {
		*o = DeviceAllOf(varDeviceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "providerId")
		delete(additionalProperties, "vulnerabilityState")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "modifiedOn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceAllOf struct {
	value *DeviceAllOf
	isSet bool
}

func (v NullableDeviceAllOf) Get() *DeviceAllOf {
	return v.value
}

func (v *NullableDeviceAllOf) Set(val *DeviceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAllOf(val *DeviceAllOf) *NullableDeviceAllOf {
	return &NullableDeviceAllOf{value: val, isSet: true}
}

func (v NullableDeviceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

