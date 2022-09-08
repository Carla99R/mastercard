# Authentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationValue** | Pointer to **string** | List of name/value pairs containing authentication values. Refer &#39;Authentication Value URIs&#39; | [optional] 
**AuthenticationCode** | Pointer to **string** | Authentication Code indicates the authentication mechanism. This can contain one of the following values - AUTHENTICATED_CAVV, NON_AUTHENTICATED_3D_SECURE, NON_AUTHENTICATED and NON_SECURE. | [optional] 

## Methods

### NewAuthentication

`func NewAuthentication() *Authentication`

NewAuthentication instantiates a new Authentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationWithDefaults

`func NewAuthenticationWithDefaults() *Authentication`

NewAuthenticationWithDefaults instantiates a new Authentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationValue

`func (o *Authentication) GetAuthenticationValue() string`

GetAuthenticationValue returns the AuthenticationValue field if non-nil, zero value otherwise.

### GetAuthenticationValueOk

`func (o *Authentication) GetAuthenticationValueOk() (*string, bool)`

GetAuthenticationValueOk returns a tuple with the AuthenticationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationValue

`func (o *Authentication) SetAuthenticationValue(v string)`

SetAuthenticationValue sets AuthenticationValue field to given value.

### HasAuthenticationValue

`func (o *Authentication) HasAuthenticationValue() bool`

HasAuthenticationValue returns a boolean if a field has been set.

### GetAuthenticationCode

`func (o *Authentication) GetAuthenticationCode() string`

GetAuthenticationCode returns the AuthenticationCode field if non-nil, zero value otherwise.

### GetAuthenticationCodeOk

`func (o *Authentication) GetAuthenticationCodeOk() (*string, bool)`

GetAuthenticationCodeOk returns a tuple with the AuthenticationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCode

`func (o *Authentication) SetAuthenticationCode(v string)`

SetAuthenticationCode sets AuthenticationCode field to given value.

### HasAuthenticationCode

`func (o *Authentication) HasAuthenticationCode() bool`

HasAuthenticationCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


