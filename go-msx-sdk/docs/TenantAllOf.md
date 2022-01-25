# TenantAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**ModifiedOn** | Pointer to **time.Time** |  | [optional] [readonly] 
**Suspended** | Pointer to **NullableBool** |  | [optional] [readonly] 
**NumberOfChildren** | Pointer to **NullableInt64** |  | [optional] [readonly] 

## Methods

### NewTenantAllOf

`func NewTenantAllOf() *TenantAllOf`

NewTenantAllOf instantiates a new TenantAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAllOfWithDefaults

`func NewTenantAllOfWithDefaults() *TenantAllOf`

NewTenantAllOfWithDefaults instantiates a new TenantAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *TenantAllOf) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *TenantAllOf) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *TenantAllOf) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *TenantAllOf) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *TenantAllOf) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *TenantAllOf) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *TenantAllOf) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *TenantAllOf) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetSuspended

`func (o *TenantAllOf) GetSuspended() bool`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *TenantAllOf) GetSuspendedOk() (*bool, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *TenantAllOf) SetSuspended(v bool)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *TenantAllOf) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### SetSuspendedNil

`func (o *TenantAllOf) SetSuspendedNil(b bool)`

 SetSuspendedNil sets the value for Suspended to be an explicit nil

### UnsetSuspended
`func (o *TenantAllOf) UnsetSuspended()`

UnsetSuspended ensures that no value is present for Suspended, not even an explicit nil
### GetNumberOfChildren

`func (o *TenantAllOf) GetNumberOfChildren() int64`

GetNumberOfChildren returns the NumberOfChildren field if non-nil, zero value otherwise.

### GetNumberOfChildrenOk

`func (o *TenantAllOf) GetNumberOfChildrenOk() (*int64, bool)`

GetNumberOfChildrenOk returns a tuple with the NumberOfChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfChildren

`func (o *TenantAllOf) SetNumberOfChildren(v int64)`

SetNumberOfChildren sets NumberOfChildren field to given value.

### HasNumberOfChildren

`func (o *TenantAllOf) HasNumberOfChildren() bool`

HasNumberOfChildren returns a boolean if a field has been set.

### SetNumberOfChildrenNil

`func (o *TenantAllOf) SetNumberOfChildrenNil(b bool)`

 SetNumberOfChildrenNil sets the value for NumberOfChildren to be an explicit nil

### UnsetNumberOfChildren
`func (o *TenantAllOf) UnsetNumberOfChildren()`

UnsetNumberOfChildren ensures that no value is present for NumberOfChildren, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


