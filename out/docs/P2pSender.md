# P2pSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The sender&#39;s first name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**MiddleName** | Pointer to **string** | The sender&#39;s middle name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**LastName** | Pointer to **string** | The sender&#39;s last name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**Nationality** | Pointer to **string** | The sender&#39;s nationality as an ISO 3166-1 uppercase alpha-3 country code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, the United States of America is USA. | [optional] 
**DateOfBirth** | Pointer to **string** | The sender&#39;s date of birth in ISO 8601 full date format (YYYY-MM-DD). For example, 30 December 1999 is the value &#39;1999-12-30&#39;. Numeric and -, length 10. | [optional] 
**AccountType** | Pointer to **string** | Identifies the sender&#39;s account number type. This field is optional. Valid values: &#39;00&#39; Other, &#39;01&#39; RTN + Bank Account, &#39;02&#39; IBAN, &#39;03&#39; Card Account, &#39;04&#39; Email, &#39;05&#39; Phone Number, &#39;06&#39; Bank account number (BAN) + Bank Identification Сode (BIC), &#39;07&#39; Wallet ID, &#39;08&#39; Social Network ID. Numeric, 2 characters. | [optional] 
**Address** | Pointer to [**P2pResponseSenderAddress**](P2pResponseSenderAddress.md) |  | [optional] 
**Phone** | Pointer to **string** | The sender&#39;s phone number including country code. Numeric, length 1-15. | [optional] 
**Email** | Pointer to **string** | The sender&#39;s email address. Alphanumeric Special [a-zA-Z0-9-@+.*$_], length 5-254. | [optional] 
**GovernmentIds** | Pointer to [**GovernmentIdsResponse**](GovernmentIdsResponse.md) |  | [optional] 
**SanctionScore** | Pointer to **string** | This field is not used at this time. The Sanctions Screening Service is suspended until further notice. | [optional] 

## Methods

### NewP2pSender

`func NewP2pSender() *P2pSender`

NewP2pSender instantiates a new P2pSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP2pSenderWithDefaults

`func NewP2pSenderWithDefaults() *P2pSender`

NewP2pSenderWithDefaults instantiates a new P2pSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *P2pSender) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *P2pSender) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *P2pSender) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *P2pSender) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *P2pSender) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *P2pSender) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *P2pSender) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *P2pSender) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *P2pSender) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *P2pSender) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *P2pSender) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *P2pSender) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetNationality

`func (o *P2pSender) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *P2pSender) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *P2pSender) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *P2pSender) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *P2pSender) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *P2pSender) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *P2pSender) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *P2pSender) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetAccountType

`func (o *P2pSender) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *P2pSender) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *P2pSender) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *P2pSender) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAddress

`func (o *P2pSender) GetAddress() P2pResponseSenderAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *P2pSender) GetAddressOk() (*P2pResponseSenderAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *P2pSender) SetAddress(v P2pResponseSenderAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *P2pSender) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhone

`func (o *P2pSender) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *P2pSender) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *P2pSender) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *P2pSender) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *P2pSender) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *P2pSender) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *P2pSender) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *P2pSender) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGovernmentIds

`func (o *P2pSender) GetGovernmentIds() GovernmentIdsResponse`

GetGovernmentIds returns the GovernmentIds field if non-nil, zero value otherwise.

### GetGovernmentIdsOk

`func (o *P2pSender) GetGovernmentIdsOk() (*GovernmentIdsResponse, bool)`

GetGovernmentIdsOk returns a tuple with the GovernmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentIds

`func (o *P2pSender) SetGovernmentIds(v GovernmentIdsResponse)`

SetGovernmentIds sets GovernmentIds field to given value.

### HasGovernmentIds

`func (o *P2pSender) HasGovernmentIds() bool`

HasGovernmentIds returns a boolean if a field has been set.

### GetSanctionScore

`func (o *P2pSender) GetSanctionScore() string`

GetSanctionScore returns the SanctionScore field if non-nil, zero value otherwise.

### GetSanctionScoreOk

`func (o *P2pSender) GetSanctionScoreOk() (*string, bool)`

GetSanctionScoreOk returns a tuple with the SanctionScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionScore

`func (o *P2pSender) SetSanctionScore(v string)`

SetSanctionScore sets SanctionScore field to given value.

### HasSanctionScore

`func (o *P2pSender) HasSanctionScore() bool`

HasSanctionScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


