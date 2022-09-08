# TransactionHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The type of resource being returned: list. | [optional] 
**ItemCount** | Pointer to **int64** | The number of Transactions in the list: 1. | [optional] 
**Data** | Pointer to [**Data**](Data.md) |  | [optional] 

## Methods

### NewTransactionHistory

`func NewTransactionHistory() *TransactionHistory`

NewTransactionHistory instantiates a new TransactionHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionHistoryWithDefaults

`func NewTransactionHistoryWithDefaults() *TransactionHistory`

NewTransactionHistoryWithDefaults instantiates a new TransactionHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *TransactionHistory) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TransactionHistory) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TransactionHistory) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *TransactionHistory) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetItemCount

`func (o *TransactionHistory) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *TransactionHistory) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *TransactionHistory) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *TransactionHistory) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetData

`func (o *TransactionHistory) GetData() Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionHistory) GetDataOk() (*Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionHistory) SetData(v Data)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionHistory) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


