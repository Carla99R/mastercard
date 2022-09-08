# Sender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Required. The sender&#39;s first name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. If you cannot provide the actual value, you must provide the alternative value &#39;#NOTINCLUDED&#39;. | 
**MiddleName** | Pointer to **string** | The sender&#39;s middle name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. | [optional] 
**LastName** | **string** | Required. The sender&#39;s last name. Alphanumeric Special [a-zA-Z0-9 !\&quot;#$%&amp;&#39;()*+,-./&amp;#92;:;&lt;&#x3D;&gt;?@[]_&#x60;{|}~ÀÁÂÃÄÅÇÈÉÊËÌÍÎÏÑÒÓÔÕÖÙÚÛÜÝàáâãäåçèéêëìíîïñòóôõöùúûüýÿ], length 1-40. If you cannot provide the actual value, you must provide the alternative value &#39;#NOTINCLUDED&#39;. | 
**Nationality** | Pointer to **string** | The sender&#39;s nationality as an ISO 3166-1 uppercase alpha-3 country code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, the United States of America is USA. | [optional] 
**DateOfBirth** | Pointer to **string** | The sender&#39;s date of birth in ISO 8601 full date format (YYYY-MM-DD). For example, 30 December 1999 is the value &#39;1999-12-30&#39;. Numeric and -, length 10. | [optional] 
**AccountType** | Pointer to **string** | Identifies the sender&#39;s account number type. This field is optional. Valid values: &#39;00&#39; Other, &#39;01&#39; RTN + Bank Account, &#39;02&#39; IBAN, &#39;03&#39; Card Account, &#39;04&#39; Email, &#39;05&#39; Phone Number, &#39;06&#39; Bank account number (BAN) + Bank Identification Сode (BIC), &#39;07&#39; Wallet ID, &#39;08&#39; Social Network ID. Numeric, 2 characters. | [optional] 
**Address** | [**RequestSenderAddress**](RequestSenderAddress.md) |  | 
**Phone** | Pointer to **string** | The sender&#39;s phone number including country code. Numeric, length 1-15. | [optional] 
**Email** | Pointer to **string** | The sender&#39;s email address. Alphanumeric Special [a-zA-Z0-9-@+.*$_], length 5-254. | [optional] 
**GovernmentIds** | Pointer to [**GovernmentIds**](GovernmentIds.md) |  | [optional] 

## Methods

### NewSender

`func NewSender(firstName string, lastName string, address RequestSenderAddress, ) *Sender`

NewSender instantiates a new Sender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSenderWithDefaults

`func NewSenderWithDefaults() *Sender`

NewSenderWithDefaults instantiates a new Sender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *Sender) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Sender) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Sender) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetMiddleName

`func (o *Sender) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *Sender) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *Sender) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *Sender) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *Sender) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Sender) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Sender) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetNationality

`func (o *Sender) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *Sender) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *Sender) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *Sender) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *Sender) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *Sender) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *Sender) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *Sender) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetAccountType

`func (o *Sender) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Sender) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Sender) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *Sender) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAddress

`func (o *Sender) GetAddress() RequestSenderAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Sender) GetAddressOk() (*RequestSenderAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Sender) SetAddress(v RequestSenderAddress)`

SetAddress sets Address field to given value.


### GetPhone

`func (o *Sender) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Sender) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Sender) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Sender) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *Sender) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Sender) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Sender) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Sender) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGovernmentIds

`func (o *Sender) GetGovernmentIds() GovernmentIds`

GetGovernmentIds returns the GovernmentIds field if non-nil, zero value otherwise.

### GetGovernmentIdsOk

`func (o *Sender) GetGovernmentIdsOk() (*GovernmentIds, bool)`

GetGovernmentIdsOk returns a tuple with the GovernmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentIds

`func (o *Sender) SetGovernmentIds(v GovernmentIds)`

SetGovernmentIds sets GovernmentIds field to given value.

### HasGovernmentIds

`func (o *Sender) HasGovernmentIds() bool`

HasGovernmentIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


