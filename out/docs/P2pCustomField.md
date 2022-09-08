# P2pCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name to be displayed in the reconciliation report. This value will appear as a header in the reconciliation reports. Name values to be defined during the onboarding process. Details- 1-40 | [optional] 
**Value** | Pointer to **string** | Value to be displayed in the reconciliation report. Details- 1-200 | [optional] 

## Methods

### NewP2pCustomField

`func NewP2pCustomField() *P2pCustomField`

NewP2pCustomField instantiates a new P2pCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP2pCustomFieldWithDefaults

`func NewP2pCustomFieldWithDefaults() *P2pCustomField`

NewP2pCustomFieldWithDefaults instantiates a new P2pCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *P2pCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *P2pCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *P2pCustomField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *P2pCustomField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *P2pCustomField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *P2pCustomField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *P2pCustomField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *P2pCustomField) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


