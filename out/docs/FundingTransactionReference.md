# FundingTransactionReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | For future use. Network on which funding transaction is processed. Valid values will be provided during partner onboarding. | 
**ReferenceNumber** | **string** | For future use. Provide the ID from the funding transfer resource. Details- 6-40 | 

## Methods

### NewFundingTransactionReference

`func NewFundingTransactionReference(network string, referenceNumber string, ) *FundingTransactionReference`

NewFundingTransactionReference instantiates a new FundingTransactionReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundingTransactionReferenceWithDefaults

`func NewFundingTransactionReferenceWithDefaults() *FundingTransactionReference`

NewFundingTransactionReferenceWithDefaults instantiates a new FundingTransactionReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *FundingTransactionReference) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *FundingTransactionReference) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *FundingTransactionReference) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetReferenceNumber

`func (o *FundingTransactionReference) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *FundingTransactionReference) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *FundingTransactionReference) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


