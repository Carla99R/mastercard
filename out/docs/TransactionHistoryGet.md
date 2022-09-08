# TransactionHistoryGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The type of resource being returned: list. | [optional] 
**ItemCount** | Pointer to **int64** | Number of Transactions in the list. Details- Numeric Example: 1 | [optional] 
**Data** | Pointer to [**DataGet**](DataGet.md) |  | [optional] 

## Methods

### NewTransactionHistoryGet

`func NewTransactionHistoryGet() *TransactionHistoryGet`

NewTransactionHistoryGet instantiates a new TransactionHistoryGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionHistoryGetWithDefaults

`func NewTransactionHistoryGetWithDefaults() *TransactionHistoryGet`

NewTransactionHistoryGetWithDefaults instantiates a new TransactionHistoryGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *TransactionHistoryGet) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TransactionHistoryGet) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TransactionHistoryGet) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *TransactionHistoryGet) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetItemCount

`func (o *TransactionHistoryGet) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *TransactionHistoryGet) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *TransactionHistoryGet) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *TransactionHistoryGet) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetData

`func (o *TransactionHistoryGet) GetData() DataGet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionHistoryGet) GetDataOk() (*DataGet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionHistoryGet) SetData(v DataGet)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionHistoryGet) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


