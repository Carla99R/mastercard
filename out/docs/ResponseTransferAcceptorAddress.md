# ResponseTransferAcceptorAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | Pointer to **string** | First line of the transfer acceptor&#39;s address. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-50. | [optional] 
**Line2** | Pointer to **string** | Second line of the transfer acceptor&#39;s address. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-50. | [optional] 
**City** | Pointer to **string** | The transfer acceptor&#39;s city. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-25. | [optional] 
**CountrySubdivision** | Pointer to **string** | The transfer acceptor&#39;s province, state or territory, as an ISO 3166-2 uppercase alpha 2 or 3 character country subdivision code. For example, Missouri is MO. | [optional] 
**PostalCode** | Pointer to **string** | The transfer acceptor&#39;s postal code. For USA, this is 5 digits or 5 digits hyphen 4 digits, for example &#39;63368&#39;, &#39;63368-5555&#39;. For other regions, this is alphanumeric, length 1-10. | [optional] 
**Country** | Pointer to **string** | The transfer acceptor&#39;s country as an ISO 3166-1 uppercase alpha-3 country code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, the United States of America is USA. | [optional] 

## Methods

### NewResponseTransferAcceptorAddress

`func NewResponseTransferAcceptorAddress() *ResponseTransferAcceptorAddress`

NewResponseTransferAcceptorAddress instantiates a new ResponseTransferAcceptorAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTransferAcceptorAddressWithDefaults

`func NewResponseTransferAcceptorAddressWithDefaults() *ResponseTransferAcceptorAddress`

NewResponseTransferAcceptorAddressWithDefaults instantiates a new ResponseTransferAcceptorAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *ResponseTransferAcceptorAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *ResponseTransferAcceptorAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *ResponseTransferAcceptorAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *ResponseTransferAcceptorAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *ResponseTransferAcceptorAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *ResponseTransferAcceptorAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *ResponseTransferAcceptorAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *ResponseTransferAcceptorAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *ResponseTransferAcceptorAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ResponseTransferAcceptorAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ResponseTransferAcceptorAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ResponseTransferAcceptorAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountrySubdivision

`func (o *ResponseTransferAcceptorAddress) GetCountrySubdivision() string`

GetCountrySubdivision returns the CountrySubdivision field if non-nil, zero value otherwise.

### GetCountrySubdivisionOk

`func (o *ResponseTransferAcceptorAddress) GetCountrySubdivisionOk() (*string, bool)`

GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivision

`func (o *ResponseTransferAcceptorAddress) SetCountrySubdivision(v string)`

SetCountrySubdivision sets CountrySubdivision field to given value.

### HasCountrySubdivision

`func (o *ResponseTransferAcceptorAddress) HasCountrySubdivision() bool`

HasCountrySubdivision returns a boolean if a field has been set.

### GetPostalCode

`func (o *ResponseTransferAcceptorAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ResponseTransferAcceptorAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ResponseTransferAcceptorAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ResponseTransferAcceptorAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *ResponseTransferAcceptorAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponseTransferAcceptorAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponseTransferAcceptorAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResponseTransferAcceptorAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


