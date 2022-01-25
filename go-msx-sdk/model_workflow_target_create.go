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

// WorkflowTargetCreate struct for WorkflowTargetCreate
type WorkflowTargetCreate struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Title *string `json:"title,omitempty"`
	SchemaId *string `json:"schema_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Properties map[string]interface{} `json:"properties"`
	UniqueName *string `json:"unique_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowTargetCreate WorkflowTargetCreate

// NewWorkflowTargetCreate instantiates a new WorkflowTargetCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTargetCreate(name string, properties map[string]interface{}) *WorkflowTargetCreate {
	this := WorkflowTargetCreate{}
	this.Name = name
	this.Properties = properties
	return &this
}

// NewWorkflowTargetCreateWithDefaults instantiates a new WorkflowTargetCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTargetCreateWithDefaults() *WorkflowTargetCreate {
	this := WorkflowTargetCreate{}
	return &this
}

// GetName returns the Name field value
func (o *WorkflowTargetCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowTargetCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowTargetCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowTargetCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowTargetCreate) SetDescription(v string) {
	o.Description = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkflowTargetCreate) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetCreate) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkflowTargetCreate) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkflowTargetCreate) SetTitle(v string) {
	o.Title = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *WorkflowTargetCreate) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetCreate) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *WorkflowTargetCreate) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *WorkflowTargetCreate) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowTargetCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowTargetCreate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkflowTargetCreate) SetType(v string) {
	o.Type = &v
}

// GetProperties returns the Properties field value
func (o *WorkflowTargetCreate) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *WorkflowTargetCreate) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *WorkflowTargetCreate) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetUniqueName returns the UniqueName field value if set, zero value otherwise.
func (o *WorkflowTargetCreate) GetUniqueName() string {
	if o == nil || o.UniqueName == nil {
		var ret string
		return ret
	}
	return *o.UniqueName
}

// GetUniqueNameOk returns a tuple with the UniqueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTargetCreate) GetUniqueNameOk() (*string, bool) {
	if o == nil || o.UniqueName == nil {
		return nil, false
	}
	return o.UniqueName, true
}

// HasUniqueName returns a boolean if a field has been set.
func (o *WorkflowTargetCreate) HasUniqueName() bool {
	if o != nil && o.UniqueName != nil {
		return true
	}

	return false
}

// SetUniqueName gets a reference to the given string and assigns it to the UniqueName field.
func (o *WorkflowTargetCreate) SetUniqueName(v string) {
	o.UniqueName = &v
}

func (o WorkflowTargetCreate) MarshalJSON() ([]byte, error) {
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

func (o *WorkflowTargetCreate) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowTargetCreate := _WorkflowTargetCreate{}

	if err = json.Unmarshal(bytes, &varWorkflowTargetCreate); err == nil {
		*o = WorkflowTargetCreate(varWorkflowTargetCreate)
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

type NullableWorkflowTargetCreate struct {
	value *WorkflowTargetCreate
	isSet bool
}

func (v NullableWorkflowTargetCreate) Get() *WorkflowTargetCreate {
	return v.value
}

func (v *NullableWorkflowTargetCreate) Set(val *WorkflowTargetCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTargetCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTargetCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTargetCreate(val *WorkflowTargetCreate) *NullableWorkflowTargetCreate {
	return &NullableWorkflowTargetCreate{value: val, isSet: true}
}

func (v NullableWorkflowTargetCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTargetCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

