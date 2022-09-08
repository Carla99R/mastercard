# CustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name to be displayed in the reconciliation report. This value will appear as a header in the reconciliation reports. Name values to be defined during the onboarding process. Details- 1-40 | 
**Value** | **string** | Value to be displayed in the reconciliation report. Details- 1-40 | 

## Methods

### NewCustomField

`func NewCustomField(name string, value string, ) *CustomField`

NewCustomField instantiates a new CustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldWithDefaults

`func NewCustomFieldWithDefaults() *CustomField`

NewCustomFieldWithDefaults instantiates a new CustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomField) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CustomField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomField) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


