# Transfers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **string** | The type of resource being returned: list. | [optional] 
**ItemCount** | Pointer to **int64** | Number of Transactions in the list | [optional] 
**Data** | Pointer to [**P2pData**](P2pData.md) |  | [optional] 

## Methods

### NewTransfers

`func NewTransfers() *Transfers`

NewTransfers instantiates a new Transfers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransfersWithDefaults

`func NewTransfersWithDefaults() *Transfers`

NewTransfersWithDefaults instantiates a new Transfers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *Transfers) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Transfers) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Transfers) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Transfers) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetItemCount

`func (o *Transfers) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *Transfers) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *Transfers) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *Transfers) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetData

`func (o *Transfers) GetData() P2pData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Transfers) GetDataOk() (*P2pData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Transfers) SetData(v P2pData)`

SetData sets Data field to given value.

### HasData

`func (o *Transfers) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


