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

// TokenCryptogram Contains the Token Cryptogram fields that relate to a Recipient token. If you provide a tokenized PAN (token) for the Recipient Account URI, you may optionally provide the associated cryptogram information using these fields.
type TokenCryptogram struct {
	// May be provided when the Recipient Account URI is a tokenized PAN, refer to 'Token Cryptogram' in the documentation. Alphanumeric special, length 22. Valid values: CONTACTLESS_CHIP, CONTACTLESS_MAGSTRIPE, DSRP_UCAF, CONTACT_CHIP, DSRP_DPD, DSRP_CHIP. - CONTACTLESS_CHIP: The cryptogram in token_cryptogram.value is the result of a contactless tap and chip information is read by the terminal. - CONTACTLESS_MAGSTRIPE: The cryptogram in token_cryptogram.value is the result of a contactless tap and the magstripe information is read by the terminal and the full track data is included without alteration or truncation. - DSRP_UCAF: The cryptogram in token_cryptogram.value is the result of an in-app purchase and the chip information is included as a Universal Cardholder Authentication Field (UCAF) value. - CONTACT_CHIP: The cryptogram in token_cryptogram.value is the result of PAN auto-entry via chip. - DSRP_DPD: The cryptogram in token_cryptogram.value is used to transport the digital payment data for electronic commerce (ecommerce) transactions. - DSRP_CHIP: The cryptogram in token_cryptogram.value is the result of PAN/token entry via ecommerce containing a Digital Secure Remote Payments (DSRP) cryptogram.
	Type string `json:"type"`
	// Contains formatted chip/cryptogram data relating to the token cryptogram. Format and length depends on the cryptogram type (token_cryptogram.type): - CONTACTLESS_CHIP: Hexadecimal [A-F0-9], length 2-510 - CONTACTLESS_MAGSTRIPE: Hexadecimal [A-F0-9], length 2-510 - DSRP_UCAF: Base64 [A-Za-z0-9+/=?:], length 1-510 - CONTACT_CHIP: Hexadecimal [A-F0-9], length 2-510 - DSRP_DPD: Base64 [A-Za-z0-9+/=?:], length 1-510 - DSRP_CHIP: Hexadecimal [A-F0-9], length 2-510
	Value string `json:"value"`
	// The PAN Sequence Number identifies and differentiates cards with the same PAN. Processors with chip-reading capability may pass this information for Contactless Chip and Contactless Magstripe transactions. This field may be present when token_cryptogram.type contains CONTACTLESS_CHIP, CONTACTLESS_MAGSTRIPE, CONTACT_CHIP, DSRP_DPD, or DSRP_CHIP. Numeric [0-9], length 3.
	PanSequenceNumber *string `json:"pan_sequence_number,omitempty"`
}

// NewTokenCryptogram instantiates a new TokenCryptogram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCryptogram(type_ string, value string) *TokenCryptogram {
	this := TokenCryptogram{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewTokenCryptogramWithDefaults instantiates a new TokenCryptogram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCryptogramWithDefaults() *TokenCryptogram {
	this := TokenCryptogram{}
	return &this
}

// GetType returns the Type field value
func (o *TokenCryptogram) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TokenCryptogram) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TokenCryptogram) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *TokenCryptogram) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TokenCryptogram) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TokenCryptogram) SetValue(v string) {
	o.Value = v
}

// GetPanSequenceNumber returns the PanSequenceNumber field value if set, zero value otherwise.
func (o *TokenCryptogram) GetPanSequenceNumber() string {
	if o == nil || o.PanSequenceNumber == nil {
		var ret string
		return ret
	}
	return *o.PanSequenceNumber
}

// GetPanSequenceNumberOk returns a tuple with the PanSequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCryptogram) GetPanSequenceNumberOk() (*string, bool) {
	if o == nil || o.PanSequenceNumber == nil {
		return nil, false
	}
	return o.PanSequenceNumber, true
}

// HasPanSequenceNumber returns a boolean if a field has been set.
func (o *TokenCryptogram) HasPanSequenceNumber() bool {
	if o != nil && o.PanSequenceNumber != nil {
		return true
	}

	return false
}

// SetPanSequenceNumber gets a reference to the given string and assigns it to the PanSequenceNumber field.
func (o *TokenCryptogram) SetPanSequenceNumber(v string) {
	o.PanSequenceNumber = &v
}

func (o TokenCryptogram) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.PanSequenceNumber != nil {
		toSerialize["pan_sequence_number"] = o.PanSequenceNumber
	}
	return json.Marshal(toSerialize)
}

type NullableTokenCryptogram struct {
	value *TokenCryptogram
	isSet bool
}

func (v NullableTokenCryptogram) Get() *TokenCryptogram {
	return v.value
}

func (v *NullableTokenCryptogram) Set(val *TokenCryptogram) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCryptogram) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCryptogram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCryptogram(val *TokenCryptogram) *NullableTokenCryptogram {
	return &NullableTokenCryptogram{value: val, isSet: true}
}

func (v NullableTokenCryptogram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCryptogram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


