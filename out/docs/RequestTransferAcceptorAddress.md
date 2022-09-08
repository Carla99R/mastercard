# RequestTransferAcceptorAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line1** | **string** | Required. First line of the transfer acceptor&#39;s address. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-50. | 
**Line2** | Pointer to **string** | Second line of the transfer acceptor&#39;s address. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-50. | [optional] 
**City** | **string** | Required. The transfer acceptor&#39;s city. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length: 1-25. | 
**CountrySubdivision** | Pointer to **string** | The transfer acceptor&#39;s province, state or territory. Conditional, required if transfer acceptor&#39;s &#x60;country&#x60; is USA or CAN. Must be an ISO 3166-2 uppercase alpha 2 or 3 character country subdivision code. For example, Missouri is MO. | [optional] 
**PostalCode** | Pointer to **string** | The transfer acceptor&#39;s postal code. For USA, this must be a valid value of 5 digits or 5 digits hyphen 4 digits, for example &#39;63368&#39;, &#39;63368-5555&#39;. For other regions, this can be alphanumeric, length 1-10. | [optional] 
**Country** | **string** | Required. The transfer acceptor&#39;s country as an ISO 3166-1 uppercase alpha-3 country code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, the United States of America is USA. | 

## Methods

### NewRequestTransferAcceptorAddress

`func NewRequestTransferAcceptorAddress(line1 string, city string, country string, ) *RequestTransferAcceptorAddress`

NewRequestTransferAcceptorAddress instantiates a new RequestTransferAcceptorAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTransferAcceptorAddressWithDefaults

`func NewRequestTransferAcceptorAddressWithDefaults() *RequestTransferAcceptorAddress`

NewRequestTransferAcceptorAddressWithDefaults instantiates a new RequestTransferAcceptorAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine1

`func (o *RequestTransferAcceptorAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *RequestTransferAcceptorAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *RequestTransferAcceptorAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.


### GetLine2

`func (o *RequestTransferAcceptorAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *RequestTransferAcceptorAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *RequestTransferAcceptorAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *RequestTransferAcceptorAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetCity

`func (o *RequestTransferAcceptorAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RequestTransferAcceptorAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RequestTransferAcceptorAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetCountrySubdivision

`func (o *RequestTransferAcceptorAddress) GetCountrySubdivision() string`

GetCountrySubdivision returns the CountrySubdivision field if non-nil, zero value otherwise.

### GetCountrySubdivisionOk

`func (o *RequestTransferAcceptorAddress) GetCountrySubdivisionOk() (*string, bool)`

GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivision

`func (o *RequestTransferAcceptorAddress) SetCountrySubdivision(v string)`

SetCountrySubdivision sets CountrySubdivision field to given value.

### HasCountrySubdivision

`func (o *RequestTransferAcceptorAddress) HasCountrySubdivision() bool`

HasCountrySubdivision returns a boolean if a field has been set.

### GetPostalCode

`func (o *RequestTransferAcceptorAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *RequestTransferAcceptorAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *RequestTransferAcceptorAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *RequestTransferAcceptorAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *RequestTransferAcceptorAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RequestTransferAcceptorAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RequestTransferAcceptorAddress) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


