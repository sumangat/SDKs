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

// ServiceNowConfigurationCreate struct for ServiceNowConfigurationCreate
type ServiceNowConfigurationCreate struct {
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	CriticalEvent *bool `json:"criticalEvent,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Password *string `json:"password,omitempty"`
	UserName *string `json:"userName,omitempty"`
	TenantId NullableString `json:"tenantId,omitempty"`
	Proxy NullableString `json:"proxy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceNowConfigurationCreate ServiceNowConfigurationCreate

// NewServiceNowConfigurationCreate instantiates a new ServiceNowConfigurationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceNowConfigurationCreate() *ServiceNowConfigurationCreate {
	this := ServiceNowConfigurationCreate{}
	return &this
}

// NewServiceNowConfigurationCreateWithDefaults instantiates a new ServiceNowConfigurationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceNowConfigurationCreateWithDefaults() *ServiceNowConfigurationCreate {
	this := ServiceNowConfigurationCreate{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ServiceNowConfigurationCreate) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfigurationCreate) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ServiceNowConfigurationCreate) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ServiceNowConfigurationCreate) GetClientSecret() string {
	if o == nil || o.ClientSecret == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfigurationCreate) GetClientSecretOk() (*string, bool) {
	if o == nil || o.ClientSecret == nil {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasClientSecret() bool {
	if o != nil && o.ClientSecret != nil {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ServiceNowConfigurationCreate) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetCriticalEvent returns the CriticalEvent field value if set, zero value otherwise.
func (o *ServiceNowConfigurationCreate) GetCriticalEvent() bool {
	if o == nil || o.CriticalEvent == nil {
		var ret bool
		return ret
	}
	return *o.CriticalEvent
}

// GetCriticalEventOk returns a tuple with the CriticalEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfigurationCreate) GetCriticalEventOk() (*bool, bool) {
	if o == nil || o.CriticalEvent == nil {
		return nil, false
	}
	return o.CriticalEvent, true
}

// HasCriticalEvent returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasCriticalEvent() bool {
	if o != nil && o.CriticalEvent != nil {
		return true
	}

	return false
}

// SetCriticalEvent gets a reference to the given bool and assigns it to the CriticalEvent field.
func (o *ServiceNowConfigurationCreate) SetCriticalEvent(v bool) {
	o.CriticalEvent = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ServiceNowConfigurationCreate) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfigurationCreate) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ServiceNowConfigurationCreate) SetDomain(v string) {
	o.Domain = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ServiceNowConfigurationCreate) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfigurationCreate) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ServiceNowConfigurationCreate) SetPassword(v string) {
	o.Password = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ServiceNowConfigurationCreate) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceNowConfigurationCreate) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ServiceNowConfigurationCreate) SetUserName(v string) {
	o.UserName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceNowConfigurationCreate) GetTenantId() string {
	if o == nil || o.TenantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TenantId.Get()
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceNowConfigurationCreate) GetTenantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TenantId.Get(), o.TenantId.IsSet()
}

// HasTenantId returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasTenantId() bool {
	if o != nil && o.TenantId.IsSet() {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given NullableString and assigns it to the TenantId field.
func (o *ServiceNowConfigurationCreate) SetTenantId(v string) {
	o.TenantId.Set(&v)
}
// SetTenantIdNil sets the value for TenantId to be an explicit nil
func (o *ServiceNowConfigurationCreate) SetTenantIdNil() {
	o.TenantId.Set(nil)
}

// UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil
func (o *ServiceNowConfigurationCreate) UnsetTenantId() {
	o.TenantId.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceNowConfigurationCreate) GetProxy() string {
	if o == nil || o.Proxy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceNowConfigurationCreate) GetProxyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *ServiceNowConfigurationCreate) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *ServiceNowConfigurationCreate) SetProxy(v string) {
	o.Proxy.Set(&v)
}
// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *ServiceNowConfigurationCreate) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *ServiceNowConfigurationCreate) UnsetProxy() {
	o.Proxy.Unset()
}

func (o ServiceNowConfigurationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientSecret != nil {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if o.CriticalEvent != nil {
		toSerialize["criticalEvent"] = o.CriticalEvent
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.TenantId.IsSet() {
		toSerialize["tenantId"] = o.TenantId.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceNowConfigurationCreate) UnmarshalJSON(bytes []byte) (err error) {
	varServiceNowConfigurationCreate := _ServiceNowConfigurationCreate{}

	if err = json.Unmarshal(bytes, &varServiceNowConfigurationCreate); err == nil {
		*o = ServiceNowConfigurationCreate(varServiceNowConfigurationCreate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientSecret")
		delete(additionalProperties, "criticalEvent")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "password")
		delete(additionalProperties, "userName")
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "proxy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceNowConfigurationCreate struct {
	value *ServiceNowConfigurationCreate
	isSet bool
}

func (v NullableServiceNowConfigurationCreate) Get() *ServiceNowConfigurationCreate {
	return v.value
}

func (v *NullableServiceNowConfigurationCreate) Set(val *ServiceNowConfigurationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNowConfigurationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNowConfigurationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNowConfigurationCreate(val *ServiceNowConfigurationCreate) *NullableServiceNowConfigurationCreate {
	return &NullableServiceNowConfigurationCreate{value: val, isSet: true}
}

func (v NullableServiceNowConfigurationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNowConfigurationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

