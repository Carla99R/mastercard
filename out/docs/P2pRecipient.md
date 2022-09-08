# P2pRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The recipient&#39;s first name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**MiddleName** | Pointer to **string** | The recipient&#39;s middle name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**LastName** | Pointer to **string** | The recipient&#39;s last name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**Nationality** | Pointer to **string** | The recipient&#39;s nationality as an ISO 3166-1 uppercase alpha-3 country code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, the United States of America is USA. | [optional] 
**DateOfBirth** | Pointer to **string** | The recipient&#39;s date of birth in ISO 8601 full date format (YYYY-MM-DD). For example, 30 December 1999 is the value &#39;1999-12-30&#39;. Numeric and -, length 10. | [optional] 
**AccountType** | Pointer to **string** | Identifies the recipient&#39;s account number type. This field is optional. Valid values: &#39;00&#39; Other, &#39;01&#39; RTN + Bank Account, &#39;02&#39; IBAN, &#39;03&#39; Card Account, &#39;04&#39; Email, &#39;05&#39; Phone Number, &#39;06&#39; Bank account number (BAN) + Bank Identification Сode (BIC), &#39;07&#39; Wallet ID, &#39;08&#39; Social Network ID. Numeric, 2 characters. | [optional] 
**Address** | Pointer to [**P2pResponseRecipientAddress**](P2pResponseRecipientAddress.md) |  | [optional] 
**Phone** | Pointer to **string** | The recipient&#39;s phone number including country code. Numeric, length 1-15. | [optional] 
**Email** | Pointer to **string** | The recipient&#39;s email address. Alphanumeric Special [a-zA-Z0-9-@+.*$_], length 5-254. | [optional] 
**GovernmentIds** | Pointer to [**GovernmentIdsResponse**](GovernmentIdsResponse.md) |  | [optional] 
**NameOnAccount** | Pointer to **string** | Name on the account. Applicable for non-card only. Allowable characters are alphanumeric and - / ? : ( ) . + # &#x3D; ! % &amp; * &lt; &gt; ; { @ &#39; \&quot; characters. Details- 1-40 | [optional] 
**SanctionScore** | Pointer to **string** | This field is not used at this time. The Sanctions Screening Service is suspended until further notice. | [optional] 

## Methods

### NewP2pRecipient

`func NewP2pRecipient() *P2pRecipient`

NewP2pRecipient instantiates a new P2pRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP2pRecipientWithDefaults

`func NewP2pRecipientWithDefaults() *P2pRecipient`

NewP2pRecipientWithDefaults instantiates a new P2pRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *P2pRecipient) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *P2pRecipient) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *P2pRecipient) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *P2pRecipient) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *P2pRecipient) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *P2pRecipient) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *P2pRecipient) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *P2pRecipient) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *P2pRecipient) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *P2pRecipient) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *P2pRecipient) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *P2pRecipient) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetNationality

`func (o *P2pRecipient) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *P2pRecipient) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *P2pRecipient) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *P2pRecipient) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *P2pRecipient) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *P2pRecipient) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *P2pRecipient) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *P2pRecipient) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetAccountType

`func (o *P2pRecipient) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *P2pRecipient) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *P2pRecipient) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *P2pRecipient) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAddress

`func (o *P2pRecipient) GetAddress() P2pResponseRecipientAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *P2pRecipient) GetAddressOk() (*P2pResponseRecipientAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *P2pRecipient) SetAddress(v P2pResponseRecipientAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *P2pRecipient) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhone

`func (o *P2pRecipient) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *P2pRecipient) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *P2pRecipient) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *P2pRecipient) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *P2pRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *P2pRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *P2pRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *P2pRecipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGovernmentIds

`func (o *P2pRecipient) GetGovernmentIds() GovernmentIdsResponse`

GetGovernmentIds returns the GovernmentIds field if non-nil, zero value otherwise.

### GetGovernmentIdsOk

`func (o *P2pRecipient) GetGovernmentIdsOk() (*GovernmentIdsResponse, bool)`

GetGovernmentIdsOk returns a tuple with the GovernmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentIds

`func (o *P2pRecipient) SetGovernmentIds(v GovernmentIdsResponse)`

SetGovernmentIds sets GovernmentIds field to given value.

### HasGovernmentIds

`func (o *P2pRecipient) HasGovernmentIds() bool`

HasGovernmentIds returns a boolean if a field has been set.

### GetNameOnAccount

`func (o *P2pRecipient) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *P2pRecipient) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *P2pRecipient) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.

### HasNameOnAccount

`func (o *P2pRecipient) HasNameOnAccount() bool`

HasNameOnAccount returns a boolean if a field has been set.

### GetSanctionScore

`func (o *P2pRecipient) GetSanctionScore() string`

GetSanctionScore returns the SanctionScore field if non-nil, zero value otherwise.

### GetSanctionScoreOk

`func (o *P2pRecipient) GetSanctionScoreOk() (*string, bool)`

GetSanctionScoreOk returns a tuple with the SanctionScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionScore

`func (o *P2pRecipient) SetSanctionScore(v string)`

SetSanctionScore sets SanctionScore field to given value.

### HasSanctionScore

`func (o *P2pRecipient) HasSanctionScore() bool`

HasSanctionScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


