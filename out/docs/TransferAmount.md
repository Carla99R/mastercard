# TransferAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The transfer amount. Numeric integer, 1-999999999999. The decimal point is implied based on the relevant &#x60;currency&#x60; exponent. For example, a US Dollar $53 amount is a value of 5300. | [optional] 
**Currency** | Pointer to **string** | The currency of the transfer amount as an ISO 4217 uppercase alpha-3 currency code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, US Dollars is USD. | [optional] 

## Methods

### NewTransferAmount

`func NewTransferAmount() *TransferAmount`

NewTransferAmount instantiates a new TransferAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAmountWithDefaults

`func NewTransferAmountWithDefaults() *TransferAmount`

NewTransferAmountWithDefaults instantiates a new TransferAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *TransferAmount) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TransferAmount) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TransferAmount) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TransferAmount) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrency

`func (o *TransferAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransferAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransferAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransferAmount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


