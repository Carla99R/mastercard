# RequestParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantCategoryCode** | Pointer to **string** | Classification of the merchant&#39;s type of business or service. Please contact your Mastercard representative to enable the usage of fields in this section. Details- numeric, 4. | [optional] 
**InterchangeRateDesignator** | Pointer to **string** | The IRD used for clearing the transaction on the Mastercard network. Details - Alphanumeric, length 2 characters. | [optional] 
**CardAcceptorId** | Pointer to **string** | Code identifying the card acceptor. Please contact your Mastercard representative to enable the usage of fields in this section. Details- alphanumeric, 10-15. | [optional] 
**CustomerServiceContactInfo** | Pointer to **string** | Customer service phone number that may be displayed on the recipient&#39;s bank or card statement. For U.S. or Canada regions, this must be a 10-digit phone number (e.g. 8005551234). For other regions, any length up to 13 digits is permitted. The phone number must be numeric without special characters such as spaces, - or ( ). To enable usage of this field, please contact your Mastercard representative. Details - numeric, 1-13. | [optional] 
**TransferAcceptorAddress** | Pointer to [**RequestTransferAcceptorAddress**](RequestTransferAcceptorAddress.md) |  | [optional] 

## Methods

### NewRequestParticipant

`func NewRequestParticipant() *RequestParticipant`

NewRequestParticipant instantiates a new RequestParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestParticipantWithDefaults

`func NewRequestParticipantWithDefaults() *RequestParticipant`

NewRequestParticipantWithDefaults instantiates a new RequestParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantCategoryCode

`func (o *RequestParticipant) GetMerchantCategoryCode() string`

GetMerchantCategoryCode returns the MerchantCategoryCode field if non-nil, zero value otherwise.

### GetMerchantCategoryCodeOk

`func (o *RequestParticipant) GetMerchantCategoryCodeOk() (*string, bool)`

GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantCategoryCode

`func (o *RequestParticipant) SetMerchantCategoryCode(v string)`

SetMerchantCategoryCode sets MerchantCategoryCode field to given value.

### HasMerchantCategoryCode

`func (o *RequestParticipant) HasMerchantCategoryCode() bool`

HasMerchantCategoryCode returns a boolean if a field has been set.

### GetInterchangeRateDesignator

`func (o *RequestParticipant) GetInterchangeRateDesignator() string`

GetInterchangeRateDesignator returns the InterchangeRateDesignator field if non-nil, zero value otherwise.

### GetInterchangeRateDesignatorOk

`func (o *RequestParticipant) GetInterchangeRateDesignatorOk() (*string, bool)`

GetInterchangeRateDesignatorOk returns a tuple with the InterchangeRateDesignator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterchangeRateDesignator

`func (o *RequestParticipant) SetInterchangeRateDesignator(v string)`

SetInterchangeRateDesignator sets InterchangeRateDesignator field to given value.

### HasInterchangeRateDesignator

`func (o *RequestParticipant) HasInterchangeRateDesignator() bool`

HasInterchangeRateDesignator returns a boolean if a field has been set.

### GetCardAcceptorId

`func (o *RequestParticipant) GetCardAcceptorId() string`

GetCardAcceptorId returns the CardAcceptorId field if non-nil, zero value otherwise.

### GetCardAcceptorIdOk

`func (o *RequestParticipant) GetCardAcceptorIdOk() (*string, bool)`

GetCardAcceptorIdOk returns a tuple with the CardAcceptorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardAcceptorId

`func (o *RequestParticipant) SetCardAcceptorId(v string)`

SetCardAcceptorId sets CardAcceptorId field to given value.

### HasCardAcceptorId

`func (o *RequestParticipant) HasCardAcceptorId() bool`

HasCardAcceptorId returns a boolean if a field has been set.

### GetCustomerServiceContactInfo

`func (o *RequestParticipant) GetCustomerServiceContactInfo() string`

GetCustomerServiceContactInfo returns the CustomerServiceContactInfo field if non-nil, zero value otherwise.

### GetCustomerServiceContactInfoOk

`func (o *RequestParticipant) GetCustomerServiceContactInfoOk() (*string, bool)`

GetCustomerServiceContactInfoOk returns a tuple with the CustomerServiceContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerServiceContactInfo

`func (o *RequestParticipant) SetCustomerServiceContactInfo(v string)`

SetCustomerServiceContactInfo sets CustomerServiceContactInfo field to given value.

### HasCustomerServiceContactInfo

`func (o *RequestParticipant) HasCustomerServiceContactInfo() bool`

HasCustomerServiceContactInfo returns a boolean if a field has been set.

### GetTransferAcceptorAddress

`func (o *RequestParticipant) GetTransferAcceptorAddress() RequestTransferAcceptorAddress`

GetTransferAcceptorAddress returns the TransferAcceptorAddress field if non-nil, zero value otherwise.

### GetTransferAcceptorAddressOk

`func (o *RequestParticipant) GetTransferAcceptorAddressOk() (*RequestTransferAcceptorAddress, bool)`

GetTransferAcceptorAddressOk returns a tuple with the TransferAcceptorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcceptorAddress

`func (o *RequestParticipant) SetTransferAcceptorAddress(v RequestTransferAcceptorAddress)`

SetTransferAcceptorAddress sets TransferAcceptorAddress field to given value.

### HasTransferAcceptorAddress

`func (o *RequestParticipant) HasTransferAcceptorAddress() bool`

HasTransferAcceptorAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


