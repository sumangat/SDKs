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

// WorkflowAccessMeta struct for WorkflowAccessMeta
type WorkflowAccessMeta struct {
	Adapter *WorkflowAccessMetaType `json:"adapter,omitempty"`
	RuntimeUsers *[]WorkflowAccessMetaType `json:"runtime_users,omitempty"`
	Targets *[]WorkflowAccessMetaType `json:"targets,omitempty"`
	IsIntegration *bool `json:"is_integration,omitempty"`
	IsInternal *bool `json:"is_internal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAccessMeta WorkflowAccessMeta

// NewWorkflowAccessMeta instantiates a new WorkflowAccessMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAccessMeta() *WorkflowAccessMeta {
	this := WorkflowAccessMeta{}
	return &this
}

// NewWorkflowAccessMetaWithDefaults instantiates a new WorkflowAccessMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAccessMetaWithDefaults() *WorkflowAccessMeta {
	this := WorkflowAccessMeta{}
	return &this
}

// GetAdapter returns the Adapter field value if set, zero value otherwise.
func (o *WorkflowAccessMeta) GetAdapter() WorkflowAccessMetaType {
	if o == nil || o.Adapter == nil {
		var ret WorkflowAccessMetaType
		return ret
	}
	return *o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAccessMeta) GetAdapterOk() (*WorkflowAccessMetaType, bool) {
	if o == nil || o.Adapter == nil {
		return nil, false
	}
	return o.Adapter, true
}

// HasAdapter returns a boolean if a field has been set.
func (o *WorkflowAccessMeta) HasAdapter() bool {
	if o != nil && o.Adapter != nil {
		return true
	}

	return false
}

// SetAdapter gets a reference to the given WorkflowAccessMetaType and assigns it to the Adapter field.
func (o *WorkflowAccessMeta) SetAdapter(v WorkflowAccessMetaType) {
	o.Adapter = &v
}

// GetRuntimeUsers returns the RuntimeUsers field value if set, zero value otherwise.
func (o *WorkflowAccessMeta) GetRuntimeUsers() []WorkflowAccessMetaType {
	if o == nil || o.RuntimeUsers == nil {
		var ret []WorkflowAccessMetaType
		return ret
	}
	return *o.RuntimeUsers
}

// GetRuntimeUsersOk returns a tuple with the RuntimeUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAccessMeta) GetRuntimeUsersOk() (*[]WorkflowAccessMetaType, bool) {
	if o == nil || o.RuntimeUsers == nil {
		return nil, false
	}
	return o.RuntimeUsers, true
}

// HasRuntimeUsers returns a boolean if a field has been set.
func (o *WorkflowAccessMeta) HasRuntimeUsers() bool {
	if o != nil && o.RuntimeUsers != nil {
		return true
	}

	return false
}

// SetRuntimeUsers gets a reference to the given []WorkflowAccessMetaType and assigns it to the RuntimeUsers field.
func (o *WorkflowAccessMeta) SetRuntimeUsers(v []WorkflowAccessMetaType) {
	o.RuntimeUsers = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *WorkflowAccessMeta) GetTargets() []WorkflowAccessMetaType {
	if o == nil || o.Targets == nil {
		var ret []WorkflowAccessMetaType
		return ret
	}
	return *o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAccessMeta) GetTargetsOk() (*[]WorkflowAccessMetaType, bool) {
	if o == nil || o.Targets == nil {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *WorkflowAccessMeta) HasTargets() bool {
	if o != nil && o.Targets != nil {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []WorkflowAccessMetaType and assigns it to the Targets field.
func (o *WorkflowAccessMeta) SetTargets(v []WorkflowAccessMetaType) {
	o.Targets = &v
}

// GetIsIntegration returns the IsIntegration field value if set, zero value otherwise.
func (o *WorkflowAccessMeta) GetIsIntegration() bool {
	if o == nil || o.IsIntegration == nil {
		var ret bool
		return ret
	}
	return *o.IsIntegration
}

// GetIsIntegrationOk returns a tuple with the IsIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAccessMeta) GetIsIntegrationOk() (*bool, bool) {
	if o == nil || o.IsIntegration == nil {
		return nil, false
	}
	return o.IsIntegration, true
}

// HasIsIntegration returns a boolean if a field has been set.
func (o *WorkflowAccessMeta) HasIsIntegration() bool {
	if o != nil && o.IsIntegration != nil {
		return true
	}

	return false
}

// SetIsIntegration gets a reference to the given bool and assigns it to the IsIntegration field.
func (o *WorkflowAccessMeta) SetIsIntegration(v bool) {
	o.IsIntegration = &v
}

// GetIsInternal returns the IsInternal field value if set, zero value otherwise.
func (o *WorkflowAccessMeta) GetIsInternal() bool {
	if o == nil || o.IsInternal == nil {
		var ret bool
		return ret
	}
	return *o.IsInternal
}

// GetIsInternalOk returns a tuple with the IsInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowAccessMeta) GetIsInternalOk() (*bool, bool) {
	if o == nil || o.IsInternal == nil {
		return nil, false
	}
	return o.IsInternal, true
}

// HasIsInternal returns a boolean if a field has been set.
func (o *WorkflowAccessMeta) HasIsInternal() bool {
	if o != nil && o.IsInternal != nil {
		return true
	}

	return false
}

// SetIsInternal gets a reference to the given bool and assigns it to the IsInternal field.
func (o *WorkflowAccessMeta) SetIsInternal(v bool) {
	o.IsInternal = &v
}

func (o WorkflowAccessMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Adapter != nil {
		toSerialize["adapter"] = o.Adapter
	}
	if o.RuntimeUsers != nil {
		toSerialize["runtime_users"] = o.RuntimeUsers
	}
	if o.Targets != nil {
		toSerialize["targets"] = o.Targets
	}
	if o.IsIntegration != nil {
		toSerialize["is_integration"] = o.IsIntegration
	}
	if o.IsInternal != nil {
		toSerialize["is_internal"] = o.IsInternal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowAccessMeta) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowAccessMeta := _WorkflowAccessMeta{}

	if err = json.Unmarshal(bytes, &varWorkflowAccessMeta); err == nil {
		*o = WorkflowAccessMeta(varWorkflowAccessMeta)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "adapter")
		delete(additionalProperties, "runtime_users")
		delete(additionalProperties, "targets")
		delete(additionalProperties, "is_integration")
		delete(additionalProperties, "is_internal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowAccessMeta struct {
	value *WorkflowAccessMeta
	isSet bool
}

func (v NullableWorkflowAccessMeta) Get() *WorkflowAccessMeta {
	return v.value
}

func (v *NullableWorkflowAccessMeta) Set(val *WorkflowAccessMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAccessMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAccessMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAccessMeta(val *WorkflowAccessMeta) *NullableWorkflowAccessMeta {
	return &NullableWorkflowAccessMeta{value: val, isSet: true}
}

func (v NullableWorkflowAccessMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAccessMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

