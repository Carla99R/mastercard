# P2pData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transfer** | Pointer to [**[]TransferGet**](TransferGet.md) | Transfer array | [optional] 

## Methods

### NewP2pData

`func NewP2pData() *P2pData`

NewP2pData instantiates a new P2pData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP2pDataWithDefaults

`func NewP2pDataWithDefaults() *P2pData`

NewP2pDataWithDefaults instantiates a new P2pData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransfer

`func (o *P2pData) GetTransfer() []TransferGet`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *P2pData) GetTransferOk() (*[]TransferGet, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *P2pData) SetTransfer(v []TransferGet)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *P2pData) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


