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

// TemplateAssignmentResponse struct for TemplateAssignmentResponse
type TemplateAssignmentResponse struct {
	Id *string `json:"id,omitempty"`
	TemplateId *string `json:"templateId,omitempty"`
	TemplateName *string `json:"templateName,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
	AssignedTenantId *string `json:"assignedTenantId,omitempty"`
	Inheritable NullableBool `json:"inheritable,omitempty"`
	Status *TemplateStatus `json:"status,omitempty"`
	StatusDetails NullableString `json:"statusDetails,omitempty"`
	CreatedOn NullableTime `json:"createdOn,omitempty"`
	CreatedBy NullableString `json:"createdBy,omitempty"`
	ModifiedOn NullableTime `json:"modifiedOn,omitempty"`
	ModifiedBy NullableString `json:"modifiedBy,omitempty"`
	ResponseStatus *string `json:"responseStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TemplateAssignmentResponse TemplateAssignmentResponse

// NewTemplateAssignmentResponse instantiates a new TemplateAssignmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateAssignmentResponse() *TemplateAssignmentResponse {
	this := TemplateAssignmentResponse{}
	return &this
}

// NewTemplateAssignmentResponseWithDefaults instantiates a new TemplateAssignmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateAssignmentResponseWithDefaults() *TemplateAssignmentResponse {
	this := TemplateAssignmentResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateAssignmentResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssignmentResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TemplateAssignmentResponse) SetId(v string) {
	o.Id = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *TemplateAssignmentResponse) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssignmentResponse) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *TemplateAssignmentResponse) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *TemplateAssignmentResponse) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssignmentResponse) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *TemplateAssignmentResponse) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *TemplateAssignmentResponse) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssignmentResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *TemplateAssignmentResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetAssignedTenantId returns the AssignedTenantId field value if set, zero value otherwise.
func (o *TemplateAssignmentResponse) GetAssignedTenantId() string {
	if o == nil || o.AssignedTenantId == nil {
		var ret string
		return ret
	}
	return *o.AssignedTenantId
}

// GetAssignedTenantIdOk returns a tuple with the AssignedTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssignmentResponse) GetAssignedTenantIdOk() (*string, bool) {
	if o == nil || o.AssignedTenantId == nil {
		return nil, false
	}
	return o.AssignedTenantId, true
}

// HasAssignedTenantId returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasAssignedTenantId() bool {
	if o != nil && o.AssignedTenantId != nil {
		return true
	}

	return false
}

// SetAssignedTenantId gets a reference to the given string and assigns it to the AssignedTenantId field.
func (o *TemplateAssignmentResponse) SetAssignedTenantId(v string) {
	o.AssignedTenantId = &v
}

// GetInheritable returns the Inheritable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateAssignmentResponse) GetInheritable() bool {
	if o == nil || o.Inheritable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Inheritable.Get()
}

// GetInheritableOk returns a tuple with the Inheritable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateAssignmentResponse) GetInheritableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Inheritable.Get(), o.Inheritable.IsSet()
}

// HasInheritable returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasInheritable() bool {
	if o != nil && o.Inheritable.IsSet() {
		return true
	}

	return false
}

// SetInheritable gets a reference to the given NullableBool and assigns it to the Inheritable field.
func (o *TemplateAssignmentResponse) SetInheritable(v bool) {
	o.Inheritable.Set(&v)
}
// SetInheritableNil sets the value for Inheritable to be an explicit nil
func (o *TemplateAssignmentResponse) SetInheritableNil() {
	o.Inheritable.Set(nil)
}

// UnsetInheritable ensures that no value is present for Inheritable, not even an explicit nil
func (o *TemplateAssignmentResponse) UnsetInheritable() {
	o.Inheritable.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TemplateAssignmentResponse) GetStatus() TemplateStatus {
	if o == nil || o.Status == nil {
		var ret TemplateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssignmentResponse) GetStatusOk() (*TemplateStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TemplateStatus and assigns it to the Status field.
func (o *TemplateAssignmentResponse) SetStatus(v TemplateStatus) {
	o.Status = &v
}

// GetStatusDetails returns the StatusDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateAssignmentResponse) GetStatusDetails() string {
	if o == nil || o.StatusDetails.Get() == nil {
		var ret string
		return ret
	}
	return *o.StatusDetails.Get()
}

// GetStatusDetailsOk returns a tuple with the StatusDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateAssignmentResponse) GetStatusDetailsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatusDetails.Get(), o.StatusDetails.IsSet()
}

// HasStatusDetails returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasStatusDetails() bool {
	if o != nil && o.StatusDetails.IsSet() {
		return true
	}

	return false
}

// SetStatusDetails gets a reference to the given NullableString and assigns it to the StatusDetails field.
func (o *TemplateAssignmentResponse) SetStatusDetails(v string) {
	o.StatusDetails.Set(&v)
}
// SetStatusDetailsNil sets the value for StatusDetails to be an explicit nil
func (o *TemplateAssignmentResponse) SetStatusDetailsNil() {
	o.StatusDetails.Set(nil)
}

// UnsetStatusDetails ensures that no value is present for StatusDetails, not even an explicit nil
func (o *TemplateAssignmentResponse) UnsetStatusDetails() {
	o.StatusDetails.Unset()
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateAssignmentResponse) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn.Get()
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateAssignmentResponse) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedOn.Get(), o.CreatedOn.IsSet()
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasCreatedOn() bool {
	if o != nil && o.CreatedOn.IsSet() {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given NullableTime and assigns it to the CreatedOn field.
func (o *TemplateAssignmentResponse) SetCreatedOn(v time.Time) {
	o.CreatedOn.Set(&v)
}
// SetCreatedOnNil sets the value for CreatedOn to be an explicit nil
func (o *TemplateAssignmentResponse) SetCreatedOnNil() {
	o.CreatedOn.Set(nil)
}

// UnsetCreatedOn ensures that no value is present for CreatedOn, not even an explicit nil
func (o *TemplateAssignmentResponse) UnsetCreatedOn() {
	o.CreatedOn.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateAssignmentResponse) GetCreatedBy() string {
	if o == nil || o.CreatedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateAssignmentResponse) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *TemplateAssignmentResponse) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *TemplateAssignmentResponse) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *TemplateAssignmentResponse) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetModifiedOn returns the ModifiedOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateAssignmentResponse) GetModifiedOn() time.Time {
	if o == nil || o.ModifiedOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedOn.Get()
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateAssignmentResponse) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedOn.Get(), o.ModifiedOn.IsSet()
}

// HasModifiedOn returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasModifiedOn() bool {
	if o != nil && o.ModifiedOn.IsSet() {
		return true
	}

	return false
}

// SetModifiedOn gets a reference to the given NullableTime and assigns it to the ModifiedOn field.
func (o *TemplateAssignmentResponse) SetModifiedOn(v time.Time) {
	o.ModifiedOn.Set(&v)
}
// SetModifiedOnNil sets the value for ModifiedOn to be an explicit nil
func (o *TemplateAssignmentResponse) SetModifiedOnNil() {
	o.ModifiedOn.Set(nil)
}

// UnsetModifiedOn ensures that no value is present for ModifiedOn, not even an explicit nil
func (o *TemplateAssignmentResponse) UnsetModifiedOn() {
	o.ModifiedOn.Unset()
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateAssignmentResponse) GetModifiedBy() string {
	if o == nil || o.ModifiedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.ModifiedBy.Get()
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateAssignmentResponse) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedBy.Get(), o.ModifiedBy.IsSet()
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given NullableString and assigns it to the ModifiedBy field.
func (o *TemplateAssignmentResponse) SetModifiedBy(v string) {
	o.ModifiedBy.Set(&v)
}
// SetModifiedByNil sets the value for ModifiedBy to be an explicit nil
func (o *TemplateAssignmentResponse) SetModifiedByNil() {
	o.ModifiedBy.Set(nil)
}

// UnsetModifiedBy ensures that no value is present for ModifiedBy, not even an explicit nil
func (o *TemplateAssignmentResponse) UnsetModifiedBy() {
	o.ModifiedBy.Unset()
}

// GetResponseStatus returns the ResponseStatus field value if set, zero value otherwise.
func (o *TemplateAssignmentResponse) GetResponseStatus() string {
	if o == nil || o.ResponseStatus == nil {
		var ret string
		return ret
	}
	return *o.ResponseStatus
}

// GetResponseStatusOk returns a tuple with the ResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateAssignmentResponse) GetResponseStatusOk() (*string, bool) {
	if o == nil || o.ResponseStatus == nil {
		return nil, false
	}
	return o.ResponseStatus, true
}

// HasResponseStatus returns a boolean if a field has been set.
func (o *TemplateAssignmentResponse) HasResponseStatus() bool {
	if o != nil && o.ResponseStatus != nil {
		return true
	}

	return false
}

// SetResponseStatus gets a reference to the given string and assigns it to the ResponseStatus field.
func (o *TemplateAssignmentResponse) SetResponseStatus(v string) {
	o.ResponseStatus = &v
}

func (o TemplateAssignmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.TemplateName != nil {
		toSerialize["templateName"] = o.TemplateName
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.AssignedTenantId != nil {
		toSerialize["assignedTenantId"] = o.AssignedTenantId
	}
	if o.Inheritable.IsSet() {
		toSerialize["inheritable"] = o.Inheritable.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetails.IsSet() {
		toSerialize["statusDetails"] = o.StatusDetails.Get()
	}
	if o.CreatedOn.IsSet() {
		toSerialize["createdOn"] = o.CreatedOn.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.ModifiedOn.IsSet() {
		toSerialize["modifiedOn"] = o.ModifiedOn.Get()
	}
	if o.ModifiedBy.IsSet() {
		toSerialize["modifiedBy"] = o.ModifiedBy.Get()
	}
	if o.ResponseStatus != nil {
		toSerialize["responseStatus"] = o.ResponseStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TemplateAssignmentResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTemplateAssignmentResponse := _TemplateAssignmentResponse{}

	if err = json.Unmarshal(bytes, &varTemplateAssignmentResponse); err == nil {
		*o = TemplateAssignmentResponse(varTemplateAssignmentResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "templateName")
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "assignedTenantId")
		delete(additionalProperties, "inheritable")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetails")
		delete(additionalProperties, "createdOn")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "modifiedOn")
		delete(additionalProperties, "modifiedBy")
		delete(additionalProperties, "responseStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTemplateAssignmentResponse struct {
	value *TemplateAssignmentResponse
	isSet bool
}

func (v NullableTemplateAssignmentResponse) Get() *TemplateAssignmentResponse {
	return v.value
}

func (v *NullableTemplateAssignmentResponse) Set(val *TemplateAssignmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateAssignmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateAssignmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateAssignmentResponse(val *TemplateAssignmentResponse) *NullableTemplateAssignmentResponse {
	return &NullableTemplateAssignmentResponse{value: val, isSet: true}
}

func (v NullableTemplateAssignmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateAssignmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

