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

// WorkflowTargetUpdate struct for WorkflowTargetUpdate
type WorkflowTargetUpdate struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Title *string `json:"title,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Properties map[string]interface{} `json:"properties"`
	UniqueName *string `json:"unique_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetUpdate WorkflowTargetUpdate

// NewWorkflowTargetUpdate instantiates a new WorkflowTargetUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetUpdate(name string, properties map[string]interface{}) *WorkflowTargetUpdate {
	this := WorkflowTargetUpdate{}
	this.Name = name
	this.Properties = properties
	return &this
}

// NewWorkflowTargetUpdateWithDefaults instantiates a new WorkflowTargetUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetUpdateWithDefaults() *WorkflowTargetUpdate {
	this := WorkflowTargetUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowTargetUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetUpdate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowTargetUpdate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowTargetUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowTargetUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowTargetUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkflowTargetUpdate) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetUpdate) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkflowTargetUpdate) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkflowTargetUpdate) SetTitle(v string) {
	o.Title = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowTargetUpdate) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetUpdate) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowTargetUpdate) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowTargetUpdate) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowTargetUpdate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetUpdate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowTargetUpdate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowTargetUpdate) SetType(v string) {
	o.Type = &v
}

// GetProperties returns the Properties field value
func (o *WorkflowTargetUpdate) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetUpdate) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *WorkflowTargetUpdate) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise.
func (o *WorkflowTargetUpdate) GetUniqueName() string {
	if o == nil || o.UniqueName == nil {
		var ret string
		return ret
	}
	return *o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetUpdate) GetUniqueNameOk() (*string, bool) {
	if o == nil || o.UniqueName == nil {
		return nil, false
	}
	return o.UniqueName, true
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowTargetUpdate) HasUniqueName() bool {
	if o != nil && o.UniqueName != nil {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given string and assigns it to the UniqueName field.
func (o *WorkflowTargetUpdate) SetUniqueName(v string) {
	o.UniqueName = &v
}

func (o WorkflowTargetUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.SchemaId != nil {
		toSerialize["schema_id"] = o.SchemaId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if o.UniqueName != nil {
		toSerialize["unique_name"] = o.UniqueName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowTargetUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTargetUpdate := _WorkflowTargetUpdate{}

	if err = json.Unmarshal(bytes, &varWorkflowTargetUpdate); err == nil {
		*o = WorkflowTargetUpdate(varWorkflowTargetUpdate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "title")
		delete(additionalProperties, "schema_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "unique_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowTargetUpdate struct {
	value *WorkflowTargetUpdate
	isSet bool
}

func (v NullableWorkflowTargetUpdate) Get() *WorkflowTargetUpdate {
	return v.value
}

func (v *NullableWorkflowTargetUpdate) Set(val *WorkflowTargetUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetUpdate(val *WorkflowTargetUpdate) *NullableWorkflowTargetUpdate {
	return &NullableWorkflowTargetUpdate{value: val, isSet: true}
}

func (v NullableWorkflowTargetUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

