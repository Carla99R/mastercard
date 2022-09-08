# ResponseParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantCategoryCode** | Pointer to **string** | Classification of the merchant&#39;s type of business or service. The value in the response is the code passed to the network and it may be different from the value provided in the request. Details- numeric, 4. | [optional] 
**InterchangeRateDesignator** | Pointer to **string** | The IRD used for clearing the transaction on the Mastercard network. Details - Alphanumeric, length 2 characters. | [optional] 
**CardAcceptorId** | Pointer to **string** | Code identifying the card acceptor. Please contact your Mastercard representative to enable the usage of fields in this section. Details- alphanumeric, 10-15. | [optional] 
**CustomerServiceContactInfo** | Pointer to **string** | Customer service phone number that may be displayed on the recipient&#39;s bank or card statement. For U.S. or Canada regions, this must be a 10-digit phone number (e.g. 8005551234). For other regions, any length up to 13 digits is permitted. The phone number must be numeric without special characters such as spaces, - or ( ). To enable usage of this field, please contact your Mastercard representative. Details - numeric, 1-13. | [optional] 
**TransferAcceptorAddress** | Pointer to [**ResponseTransferAcceptorAddress**](ResponseTransferAcceptorAddress.md) |  | [optional] 

## Methods

### NewResponseParticipant

`func NewResponseParticipant() *ResponseParticipant`

NewResponseParticipant instantiates a new ResponseParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseParticipantWithDefaults

`func NewResponseParticipantWithDefaults() *ResponseParticipant`

NewResponseParticipantWithDefaults instantiates a new ResponseParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantCategoryCode

`func (o *ResponseParticipant) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *ResponseParticipant) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *ResponseParticipant) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *ResponseParticipant) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetInterchangeRateDesignator

`func (o *ResponseParticipant) GetInterchangeRateDesignator() string`

GetInterchangeRateDesignator returns the InterchangeRateDesignator field if non-nil, zero value otherwise.

### GetInterchangeRateDesignatorOk

`func (o *ResponseParticipant) GetInterchangeRateDesignatorOk() (*string, bool)`

GetInterchangeRateDesignatorOk returns a tuple with the InterchangeRateDesignator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterchangeRateDesignator

`func (o *ResponseParticipant) SetInterchangeRateDesignator(v string)`

SetInterchangeRateDesignator sets InterchangeRateDesignator field to given value.

### HasInterchangeRateDesignator

`func (o *ResponseParticipant) HasInterchangeRateDesignator() bool`

HasInterchangeRateDesignator returns a boolean if a field has been set.

### GetCardAcceptorId

`func (o *ResponseParticipant) GetCardAcceptorId() string`

GetCardAcceptorId returns the CardAcceptorId field if non-nil, zero value otherwise.

### GetCardAcceptorIdOk

`func (o *ResponseParticipant) GetCardAcceptorIdOk() (*string, bool)`

GetCardAcceptorIdOk returns a tuple with the CardAcceptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptorId

`func (o *ResponseParticipant) SetCardAcceptorId(v string)`

SetCardAcceptorId sets CardAcceptorId field to given value.

### HasCardAcceptorId

`func (o *ResponseParticipant) HasCardAcceptorId() bool`

HasCardAcceptorId returns a boolean if a field has been set.

### GetCustomerServiceContactInfo

`func (o *ResponseParticipant) GetCustomerServiceContactInfo() string`

GetCustomerServiceContactInfo returns the CustomerServiceContactInfo field if non-nil, zero value otherwise.

### GetCustomerServiceContactInfoOk

`func (o *ResponseParticipant) GetCustomerServiceContactInfoOk() (*string, bool)`

GetCustomerServiceContactInfoOk returns a tuple with the CustomerServiceContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerServiceContactInfo

`func (o *ResponseParticipant) SetCustomerServiceContactInfo(v string)`

SetCustomerServiceContactInfo sets CustomerServiceContactInfo field to given value.

### HasCustomerServiceContactInfo

`func (o *ResponseParticipant) HasCustomerServiceContactInfo() bool`

HasCustomerServiceContactInfo returns a boolean if a field has been set.

### GetTransferAcceptorAddress

`func (o *ResponseParticipant) GetTransferAcceptorAddress() ResponseTransferAcceptorAddress`

GetTransferAcceptorAddress returns the TransferAcceptorAddress field if non-nil, zero value otherwise.

### GetTransferAcceptorAddressOk

`func (o *ResponseParticipant) GetTransferAcceptorAddressOk() (*ResponseTransferAcceptorAddress, bool)`

GetTransferAcceptorAddressOk returns a tuple with the TransferAcceptorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcceptorAddress

`func (o *ResponseParticipant) SetTransferAcceptorAddress(v ResponseTransferAcceptorAddress)`

SetTransferAcceptorAddress sets TransferAcceptorAddress field to given value.

### HasTransferAcceptorAddress

`func (o *ResponseParticipant) HasTransferAcceptorAddress() bool`

HasTransferAcceptorAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


