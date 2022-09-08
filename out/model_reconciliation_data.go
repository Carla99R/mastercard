/*
Send Payment Transfer API

Mastercard Send Payment Transfer API

API version: 1.0.1
Contact: apisupport@mastercard.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ReconciliationData Reconciliation information in request that will be associated to the transfer and will appear in the reconciliation reporting for the partner.
type ReconciliationData struct {
	// Contains custom field names and values to appear in the reconciliation report. This can be a list of multiple name/value pairs in request. The names provided must match the configured in the partner setup.
	CustomField []CustomField `json:"custom_field,omitempty"`
}

// NewReconciliationData instantiates a new ReconciliationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReconciliationData() *ReconciliationData {
	this := ReconciliationData{}
	return &this
}

// NewReconciliationDataWithDefaults instantiates a new ReconciliationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReconciliationDataWithDefaults() *ReconciliationData {
	this := ReconciliationData{}
	return &this
}

// GetCustomField returns the CustomField field value if set, zero value otherwise.
func (o *ReconciliationData) GetCustomField() []CustomField {
	if o == nil || o.CustomField == nil {
		var ret []CustomField
		return ret
	}
	return o.CustomField
}

// GetCustomFieldOk returns a tuple with the CustomField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReconciliationData) GetCustomFieldOk() ([]CustomField, bool) {
	if o == nil || o.CustomField == nil {
		return nil, false
	}
	return o.CustomField, true
}

// HasCustomField returns a boolean if a field has been set.
func (o *ReconciliationData) HasCustomField() bool {
	if o != nil && o.CustomField != nil {
		return true
	}

	return false
}

// SetCustomField gets a reference to the given []CustomField and assigns it to the CustomField field.
func (o *ReconciliationData) SetCustomField(v []CustomField) {
	o.CustomField = v
}

func (o ReconciliationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomField != nil {
		toSerialize["custom_field"] = o.CustomField
	}
	return json.Marshal(toSerialize)
}

type NullableReconciliationData struct {
	value *ReconciliationData
	isSet bool
}

func (v NullableReconciliationData) Get() *ReconciliationData {
	return v.value
}

func (v *NullableReconciliationData) Set(val *ReconciliationData) {
	v.value = val
	v.isSet = true
}

func (v NullableReconciliationData) IsSet() bool {
	return v.isSet
}

func (v *NullableReconciliationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReconciliationData(val *ReconciliationData) *NullableReconciliationData {
	return &NullableReconciliationData{value: val, isSet: true}
}

func (v NullableReconciliationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReconciliationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

