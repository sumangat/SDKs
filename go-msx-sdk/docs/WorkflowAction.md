# WorkflowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **NullableString** |  | [optional] 
**DefinitionId** | Pointer to **NullableString** |  | [optional] 
**SchemaId** | Pointer to **NullableString** |  | [optional] 
**AdapterId** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**BaseType** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**ObjectType** | Pointer to **NullableString** |  | [optional] 
**StartedBy** | Pointer to **NullableString** |  | [optional] 
**Blocks** | Pointer to [**[]WorkflowActionBlock**](WorkflowActionBlock.md) |  | [optional] 
**Status** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Output** | Pointer to **map[string]interface{}** |  | [optional] 
**StartedOn** | Pointer to **NullableString** |  | [optional] 
**EndedOn** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWorkflowAction

`func NewWorkflowAction() *WorkflowAction`

NewWorkflowAction instantiates a new WorkflowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionWithDefaults

`func NewWorkflowActionWithDefaults() *WorkflowAction`

NewWorkflowActionWithDefaults instantiates a new WorkflowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowAction) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowAction) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowAction) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowAction) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *WorkflowAction) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *WorkflowAction) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetDefinitionId

`func (o *WorkflowAction) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *WorkflowAction) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *WorkflowAction) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *WorkflowAction) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### SetDefinitionIdNil

`func (o *WorkflowAction) SetDefinitionIdNil(b bool)`

 SetDefinitionIdNil sets the value for DefinitionId to be an explicit nil

### UnsetDefinitionId
`func (o *WorkflowAction) UnsetDefinitionId()`

UnsetDefinitionId ensures that no value is present for DefinitionId, not even an explicit nil
### GetSchemaId

`func (o *WorkflowAction) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *WorkflowAction) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *WorkflowAction) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *WorkflowAction) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### SetSchemaIdNil

`func (o *WorkflowAction) SetSchemaIdNil(b bool)`

 SetSchemaIdNil sets the value for SchemaId to be an explicit nil

### UnsetSchemaId
`func (o *WorkflowAction) UnsetSchemaId()`

UnsetSchemaId ensures that no value is present for SchemaId, not even an explicit nil
### GetAdapterId

`func (o *WorkflowAction) GetAdapterId() string`

GetAdapterId returns the AdapterId field if non-nil, zero value otherwise.

### GetAdapterIdOk

`func (o *WorkflowAction) GetAdapterIdOk() (*string, bool)`

GetAdapterIdOk returns a tuple with the AdapterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterId

`func (o *WorkflowAction) SetAdapterId(v string)`

SetAdapterId sets AdapterId field to given value.

### HasAdapterId

`func (o *WorkflowAction) HasAdapterId() bool`

HasAdapterId returns a boolean if a field has been set.

### SetAdapterIdNil

`func (o *WorkflowAction) SetAdapterIdNil(b bool)`

 SetAdapterIdNil sets the value for AdapterId to be an explicit nil

### UnsetAdapterId
`func (o *WorkflowAction) UnsetAdapterId()`

UnsetAdapterId ensures that no value is present for AdapterId, not even an explicit nil
### GetUniqueName

`func (o *WorkflowAction) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WorkflowAction) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WorkflowAction) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WorkflowAction) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WorkflowAction) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WorkflowAction) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetName

`func (o *WorkflowAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *WorkflowAction) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkflowAction) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkflowAction) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkflowAction) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *WorkflowAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseType

`func (o *WorkflowAction) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *WorkflowAction) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *WorkflowAction) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *WorkflowAction) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowAction) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowAction) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowAction) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowAction) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowAction) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowAction) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetObjectType

`func (o *WorkflowAction) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowAction) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowAction) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *WorkflowAction) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### SetObjectTypeNil

`func (o *WorkflowAction) SetObjectTypeNil(b bool)`

 SetObjectTypeNil sets the value for ObjectType to be an explicit nil

### UnsetObjectType
`func (o *WorkflowAction) UnsetObjectType()`

UnsetObjectType ensures that no value is present for ObjectType, not even an explicit nil
### GetStartedBy

`func (o *WorkflowAction) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *WorkflowAction) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *WorkflowAction) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.

### HasStartedBy

`func (o *WorkflowAction) HasStartedBy() bool`

HasStartedBy returns a boolean if a field has been set.

### SetStartedByNil

`func (o *WorkflowAction) SetStartedByNil(b bool)`

 SetStartedByNil sets the value for StartedBy to be an explicit nil

### UnsetStartedBy
`func (o *WorkflowAction) UnsetStartedBy()`

UnsetStartedBy ensures that no value is present for StartedBy, not even an explicit nil
### GetBlocks

`func (o *WorkflowAction) GetBlocks() []WorkflowActionBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *WorkflowAction) GetBlocksOk() (*[]WorkflowActionBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *WorkflowAction) SetBlocks(v []WorkflowActionBlock)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *WorkflowAction) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### SetBlocksNil

`func (o *WorkflowAction) SetBlocksNil(b bool)`

 SetBlocksNil sets the value for Blocks to be an explicit nil

### UnsetBlocks
`func (o *WorkflowAction) UnsetBlocks()`

UnsetBlocks ensures that no value is present for Blocks, not even an explicit nil
### GetStatus

`func (o *WorkflowAction) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowAction) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowAction) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *WorkflowAction) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkflowAction) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *WorkflowAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowAction) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowAction) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOutput

`func (o *WorkflowAction) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowAction) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowAction) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowAction) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowAction) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowAction) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetStartedOn

`func (o *WorkflowAction) GetStartedOn() string`

GetStartedOn returns the StartedOn field if non-nil, zero value otherwise.

### GetStartedOnOk

`func (o *WorkflowAction) GetStartedOnOk() (*string, bool)`

GetStartedOnOk returns a tuple with the StartedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedOn

`func (o *WorkflowAction) SetStartedOn(v string)`

SetStartedOn sets StartedOn field to given value.

### HasStartedOn

`func (o *WorkflowAction) HasStartedOn() bool`

HasStartedOn returns a boolean if a field has been set.

### SetStartedOnNil

`func (o *WorkflowAction) SetStartedOnNil(b bool)`

 SetStartedOnNil sets the value for StartedOn to be an explicit nil

### UnsetStartedOn
`func (o *WorkflowAction) UnsetStartedOn()`

UnsetStartedOn ensures that no value is present for StartedOn, not even an explicit nil
### GetEndedOn

`func (o *WorkflowAction) GetEndedOn() string`

GetEndedOn returns the EndedOn field if non-nil, zero value otherwise.

### GetEndedOnOk

`func (o *WorkflowAction) GetEndedOnOk() (*string, bool)`

GetEndedOnOk returns a tuple with the EndedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedOn

`func (o *WorkflowAction) SetEndedOn(v string)`

SetEndedOn sets EndedOn field to given value.

### HasEndedOn

`func (o *WorkflowAction) HasEndedOn() bool`

HasEndedOn returns a boolean if a field has been set.

### SetEndedOnNil

`func (o *WorkflowAction) SetEndedOnNil(b bool)`

 SetEndedOnNil sets the value for EndedOn to be an explicit nil

### UnsetEndedOn
`func (o *WorkflowAction) UnsetEndedOn()`

UnsetEndedOn ensures that no value is present for EndedOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


