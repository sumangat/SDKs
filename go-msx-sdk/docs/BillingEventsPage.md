# BillingEventsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] [readonly] 
**PageSize** | Pointer to **int32** |  | [optional] [readonly] 
**TotalItems** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**HasNext** | Pointer to **NullableBool** |  | [optional] [readonly] 
**HasPrevious** | Pointer to **NullableBool** |  | [optional] [readonly] 
**SortBy** | Pointer to **NullableString** |  | [optional] [readonly] 
**SortOrder** | Pointer to **NullableString** |  | [optional] [readonly] 
**Contents** | Pointer to [**[]BillingEvent**](BillingEvent.md) |  | [optional] 

## Methods

### NewBillingEventsPage

`func NewBillingEventsPage() *BillingEventsPage`

NewBillingEventsPage instantiates a new BillingEventsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingEventsPageWithDefaults

`func NewBillingEventsPageWithDefaults() *BillingEventsPage`

NewBillingEventsPageWithDefaults instantiates a new BillingEventsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *BillingEventsPage) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BillingEventsPage) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BillingEventsPage) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *BillingEventsPage) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *BillingEventsPage) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *BillingEventsPage) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *BillingEventsPage) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *BillingEventsPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalItems

`func (o *BillingEventsPage) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *BillingEventsPage) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *BillingEventsPage) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *BillingEventsPage) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### SetTotalItemsNil

`func (o *BillingEventsPage) SetTotalItemsNil(b bool)`

 SetTotalItemsNil sets the value for TotalItems to be an explicit nil

### UnsetTotalItems
`func (o *BillingEventsPage) UnsetTotalItems()`

UnsetTotalItems ensures that no value is present for TotalItems, not even an explicit nil
### GetHasNext

`func (o *BillingEventsPage) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *BillingEventsPage) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *BillingEventsPage) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *BillingEventsPage) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### SetHasNextNil

`func (o *BillingEventsPage) SetHasNextNil(b bool)`

 SetHasNextNil sets the value for HasNext to be an explicit nil

### UnsetHasNext
`func (o *BillingEventsPage) UnsetHasNext()`

UnsetHasNext ensures that no value is present for HasNext, not even an explicit nil
### GetHasPrevious

`func (o *BillingEventsPage) GetHasPrevious() bool`

GetHasPrevious returns the HasPrevious field if non-nil, zero value otherwise.

### GetHasPreviousOk

`func (o *BillingEventsPage) GetHasPreviousOk() (*bool, bool)`

GetHasPreviousOk returns a tuple with the HasPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrevious

`func (o *BillingEventsPage) SetHasPrevious(v bool)`

SetHasPrevious sets HasPrevious field to given value.

### HasHasPrevious

`func (o *BillingEventsPage) HasHasPrevious() bool`

HasHasPrevious returns a boolean if a field has been set.

### SetHasPreviousNil

`func (o *BillingEventsPage) SetHasPreviousNil(b bool)`

 SetHasPreviousNil sets the value for HasPrevious to be an explicit nil

### UnsetHasPrevious
`func (o *BillingEventsPage) UnsetHasPrevious()`

UnsetHasPrevious ensures that no value is present for HasPrevious, not even an explicit nil
### GetSortBy

`func (o *BillingEventsPage) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *BillingEventsPage) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *BillingEventsPage) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *BillingEventsPage) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *BillingEventsPage) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *BillingEventsPage) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *BillingEventsPage) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *BillingEventsPage) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *BillingEventsPage) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *BillingEventsPage) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *BillingEventsPage) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *BillingEventsPage) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetContents

`func (o *BillingEventsPage) GetContents() []BillingEvent`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BillingEventsPage) GetContentsOk() (*[]BillingEvent, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BillingEventsPage) SetContents(v []BillingEvent)`

SetContents sets Contents field to given value.

### HasContents

`func (o *BillingEventsPage) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


