# ReconciliationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomField** | Pointer to [**[]CustomField**](CustomField.md) | Contains custom field names and values to appear in the reconciliation report. This can be a list of multiple name/value pairs in request. The names provided must match the configured in the partner setup. | [optional] 

## Methods

### NewReconciliationData

`func NewReconciliationData() *ReconciliationData`

NewReconciliationData instantiates a new ReconciliationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationDataWithDefaults

`func NewReconciliationDataWithDefaults() *ReconciliationData`

NewReconciliationDataWithDefaults instantiates a new ReconciliationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomField

`func (o *ReconciliationData) GetCustomField() []CustomField`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *ReconciliationData) GetCustomFieldOk() (*[]CustomField, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *ReconciliationData) SetCustomField(v []CustomField)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *ReconciliationData) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


