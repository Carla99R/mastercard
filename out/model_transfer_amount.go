/*
Send Payment Transfer API

Mastercard Send Payment Transfer API

API version: 1.0.1
Contact: apisupport@mastercard.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TransferAmount Describes the amount paid to the recipient.
type TransferAmount struct {
	// The transfer amount. Numeric integer, 1-999999999999. The decimal point is implied based on the relevant `currency` exponent. For example, a US Dollar $53 amount is a value of 5300.
	Value *string `json:"value,omitempty"`
	// The currency of the transfer amount as an ISO 4217 uppercase alpha-3 currency code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, US Dollars is USD.
	Currency *string `json:"currency,omitempty"`
}

// NewTransferAmount instantiates a new TransferAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAmount() *TransferAmount {
	this := TransferAmount{}
	return &this
}

// NewTransferAmountWithDefaults instantiates a new TransferAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAmountWithDefaults() *TransferAmount {
	this := TransferAmount{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TransferAmount) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAmount) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TransferAmount) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TransferAmount) SetValue(v string) {
	o.Value = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TransferAmount) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferAmount) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TransferAmount) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TransferAmount) SetCurrency(v string) {
	o.Currency = &v
}

func (o TransferAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableTransferAmount struct {
	value *TransferAmount
	isSet bool
}

func (v NullableTransferAmount) Get() *TransferAmount {
	return v.value
}

func (v *NullableTransferAmount) Set(val *TransferAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAmount(val *TransferAmount) *NullableTransferAmount {
	return &NullableTransferAmount{value: val, isSet: true}
}

func (v NullableTransferAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

