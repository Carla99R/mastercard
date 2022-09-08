# Recipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Required. The recipient&#39;s first name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. If you cannot provide the actual value, you must provide the alternative value &#39;#NOTINCLUDED&#39;. | 
**MiddleName** | Pointer to **string** | The recipient&#39;s middle name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**LastName** | **string** | Required. The recipient&#39;s last name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. If you cannot provide the actual value, you must provide the alternative value &#39;#NOTINCLUDED&#39;. | 
**Nationality** | Pointer to **string** | The recipient&#39;s nationality as an ISO 3166-1 uppercase alpha-3 country code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, the United States of America is USA. | [optional] 
**DateOfBirth** | Pointer to **string** | The recipient&#39;s date of birth in ISO 8601 full date format (YYYY-MM-DD). For example, 30 December 1999 is the value &#39;1999-12-30&#39;. Numeric and -, length 10. | [optional] 
**AccountType** | Pointer to **string** | Identifies the recipient&#39;s account number type. This field is optional. Valid values: &#39;00&#39; Other, &#39;01&#39; RTN + Bank Account, &#39;02&#39; IBAN, &#39;03&#39; Card Account, &#39;04&#39; Email, &#39;05&#39; Phone Number, &#39;06&#39; Bank account number (BAN) + Bank Identification Сode (BIC), &#39;07&#39; Wallet ID, &#39;08&#39; Social Network ID. Numeric, 2 characters. | [optional] 
**Address** | Pointer to [**RequestRecipientAddress**](RequestRecipientAddress.md) |  | [optional] 
**Phone** | Pointer to **string** | The recipient&#39;s phone number including country code. Numeric, length 1-15. | [optional] 
**Email** | Pointer to **string** | The recipient&#39;s email address. Alphanumeric Special [a-zA-Z0-9-@+.*$_], length 5-254. | [optional] 
**GovernmentIds** | Pointer to [**GovernmentIds**](GovernmentIds.md) |  | [optional] 
**NameOnAccount** | Pointer to **string** | Name on the account. Applicable for non-card only. Allowable characters are alphanumeric and - / ? : ( ) . + # &#x3D; ! % &amp; * &lt; &gt; ; { @ &#39; \&quot; characters. Details- Conditional, Required for non-card only, 1-40 | [optional] 

## Methods

### NewRecipient

`func NewRecipient(firstName string, lastName string, ) *Recipient`

NewRecipient instantiates a new Recipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientWithDefaults

`func NewRecipientWithDefaults() *Recipient`

NewRecipientWithDefaults instantiates a new Recipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *Recipient) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Recipient) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Recipient) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *Recipient) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Recipient) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *Recipient) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *Recipient) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *Recipient) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Recipient) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Recipient) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetNationality

`func (o *Recipient) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *Recipient) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *Recipient) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *Recipient) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *Recipient) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Recipient) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Recipient) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Recipient) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetAccountType

`func (o *Recipient) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Recipient) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Recipient) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *Recipient) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAddress

`func (o *Recipient) GetAddress() RequestRecipientAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Recipient) GetAddressOk() (*RequestRecipientAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Recipient) SetAddress(v RequestRecipientAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Recipient) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhone

`func (o *Recipient) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Recipient) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Recipient) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Recipient) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *Recipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Recipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Recipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Recipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGovernmentIds

`func (o *Recipient) GetGovernmentIds() GovernmentIds`

GetGovernmentIds returns the GovernmentIds field if non-nil, zero value otherwise.

### GetGovernmentIdsOk

`func (o *Recipient) GetGovernmentIdsOk() (*GovernmentIds, bool)`

GetGovernmentIdsOk returns a tuple with the GovernmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentIds

`func (o *Recipient) SetGovernmentIds(v GovernmentIds)`

SetGovernmentIds sets GovernmentIds field to given value.

### HasGovernmentIds

`func (o *Recipient) HasGovernmentIds() bool`

HasGovernmentIds returns a boolean if a field has been set.

### GetNameOnAccount

`func (o *Recipient) GetNameOnAccount() string`

GetNameOnAccount returns the NameOnAccount field if non-nil, zero value otherwise.

### GetNameOnAccountOk

`func (o *Recipient) GetNameOnAccountOk() (*string, bool)`

GetNameOnAccountOk returns a tuple with the NameOnAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnAccount

`func (o *Recipient) SetNameOnAccount(v string)`

SetNameOnAccount sets NameOnAccount field to given value.

### HasNameOnAccount

`func (o *Recipient) HasNameOnAccount() bool`

HasNameOnAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


