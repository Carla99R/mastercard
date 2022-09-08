# P2pReconciliationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomField** | Pointer to [**[]P2pCustomField**](P2pCustomField.md) | Contains custom field names and values to appear in the reconciliation report. This can be a list of multiple name/value pairs in response. The names provided must match the configured in the partner setup. | [optional] 

## Methods

### NewP2pReconciliationData

`func NewP2pReconciliationData() *P2pReconciliationData`

NewP2pReconciliationData instantiates a new P2pReconciliationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewP2pReconciliationDataWithDefaults

`func NewP2pReconciliationDataWithDefaults() *P2pReconciliationData`

NewP2pReconciliationDataWithDefaults instantiates a new P2pReconciliationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomField

`func (o *P2pReconciliationData) GetCustomField() []P2pCustomField`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *P2pReconciliationData) GetCustomFieldOk() (*[]P2pCustomField, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *P2pReconciliationData) SetCustomField(v []P2pCustomField)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *P2pReconciliationData) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


