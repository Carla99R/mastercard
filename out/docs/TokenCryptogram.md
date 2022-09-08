# TokenCryptogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | May be provided when the Recipient Account URI is a tokenized PAN, refer to &#39;Token Cryptogram&#39; in the documentation. Alphanumeric special, length 22. Valid values: CONTACTLESS_CHIP, CONTACTLESS_MAGSTRIPE, DSRP_UCAF, CONTACT_CHIP, DSRP_DPD, DSRP_CHIP. - CONTACTLESS_CHIP: The cryptogram in token_cryptogram.value is the result of a contactless tap and chip information is read by the terminal. - CONTACTLESS_MAGSTRIPE: The cryptogram in token_cryptogram.value is the result of a contactless tap and the magstripe information is read by the terminal and the full track data is included without alteration or truncation. - DSRP_UCAF: The cryptogram in token_cryptogram.value is the result of an in-app purchase and the chip information is included as a Universal Cardholder Authentication Field (UCAF) value. - CONTACT_CHIP: The cryptogram in token_cryptogram.value is the result of PAN auto-entry via chip. - DSRP_DPD: The cryptogram in token_cryptogram.value is used to transport the digital payment data for electronic commerce (ecommerce) transactions. - DSRP_CHIP: The cryptogram in token_cryptogram.value is the result of PAN/token entry via ecommerce containing a Digital Secure Remote Payments (DSRP) cryptogram. | 
**Value** | **string** | Contains formatted chip/cryptogram data relating to the token cryptogram. Format and length depends on the cryptogram type (token_cryptogram.type): - CONTACTLESS_CHIP: Hexadecimal [A-F0-9], length 2-510 - CONTACTLESS_MAGSTRIPE: Hexadecimal [A-F0-9], length 2-510 - DSRP_UCAF: Base64 [A-Za-z0-9+/&#x3D;?:], length 1-510 - CONTACT_CHIP: Hexadecimal [A-F0-9], length 2-510 - DSRP_DPD: Base64 [A-Za-z0-9+/&#x3D;?:], length 1-510 - DSRP_CHIP: Hexadecimal [A-F0-9], length 2-510 | 
**PanSequenceNumber** | Pointer to **string** | The PAN Sequence Number identifies and differentiates cards with the same PAN. Processors with chip-reading capability may pass this information for Contactless Chip and Contactless Magstripe transactions. This field may be present when token_cryptogram.type contains CONTACTLESS_CHIP, CONTACTLESS_MAGSTRIPE, CONTACT_CHIP, DSRP_DPD, or DSRP_CHIP. Numeric [0-9], length 3. | [optional] 

## Methods

### NewTokenCryptogram

`func NewTokenCryptogram(type_ string, value string, ) *TokenCryptogram`

NewTokenCryptogram instantiates a new TokenCryptogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCryptogramWithDefaults

`func NewTokenCryptogramWithDefaults() *TokenCryptogram`

NewTokenCryptogramWithDefaults instantiates a new TokenCryptogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TokenCryptogram) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenCryptogram) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenCryptogram) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *TokenCryptogram) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TokenCryptogram) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TokenCryptogram) SetValue(v string)`

SetValue sets Value field to given value.


### GetPanSequenceNumber

`func (o *TokenCryptogram) GetPanSequenceNumber() string`

GetPanSequenceNumber returns the PanSequenceNumber field if non-nil, zero value otherwise.

### GetPanSequenceNumberOk

`func (o *TokenCryptogram) GetPanSequenceNumberOk() (*string, bool)`

GetPanSequenceNumberOk returns a tuple with the PanSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanSequenceNumber

`func (o *TokenCryptogram) SetPanSequenceNumber(v string)`

SetPanSequenceNumber sets PanSequenceNumber field to given value.

### HasPanSequenceNumber

`func (o *TokenCryptogram) HasPanSequenceNumber() bool`

HasPanSequenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


