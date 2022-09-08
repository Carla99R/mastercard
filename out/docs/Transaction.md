# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID for the Payment Transaction. Alphanumeric and * , - . _ ~. | [optional] 
**ResourceType** | Pointer to **string** | The type of resource being returned: transaction. | [optional] 
**AccountUri** | Pointer to **string** | URI describing the receiving account. It will include masked account information (e.g. \&quot;************9913\&quot; for a card account) but will not include security codes (e.g. CVC or expiration date for a card account). See [Account URIs](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/account-uris/). | [optional] 
**TransactionAmount** | Pointer to [**TransactionAmount**](TransactionAmount.md) |  | [optional] 
**Network** | Pointer to **string** | Proposed network for the transaction. | [optional] 
**FundsAvailability** | Pointer to **string** | An estimate of when the funds might be available. Actual Deposit Availability is determined by the financial institution. One of: IMMEDIATE, SAME_DAY, NEXT_CALENDAR_DAY, NEXT_BUSINESS_DAY, TWO_TO_FIVE_BUSINESS_DAYS, or UNKNOWN. | [optional] 
**Type** | Pointer to **string** | Type of the transaction. One of: FUNDING, PAYMENT, FUNDING_REVERSAL, or PAYMENT_REVERSAL. | [optional] 
**CreateTimestamp** | Pointer to **string** | Date and time the transaction was created, as an ISO 8601 timestamp in the format YYYY-MM-DDTHH:MM:SS±hh:mm. Refer to &#39;Date and Time Formats&#39;. | [optional] 
**Status** | Pointer to **string** | Status of the transfer. One of APPROVED, DECLINED, UNKNOWN, ERROR, PENDING. PENDING is possible for non-card transactions (not applicable to card). | [optional] 
**StatusReason** | Pointer to **string** | Reason for status. One of APPROVED, DECLINED, UNKNOWN, ERROR, PENDING, FRAUD, CARD_EXPIRED, LIMIT_EXCEEDED. PENDING is possible for non-card transactions (not applicable to card). | [optional] 
**StatusTimestamp** | Pointer to **string** | Date and time the status was changed to its current value, as an ISO 8601 timestamp in the format YYYY-MM-DDTHH:MM:SS±hh:mm. Refer to &#39;Date and Time Formats&#39;. | [optional] 
**RetrievalReference** | Pointer to **string** | Unique reference number that identifies the transaction at the network. Details- maximum length 24 | [optional] 
**SystemTraceAuditNumber** | Pointer to **string** | Unique system trace audit number for the transaction, the STAN (system trace audit number). Details- maximum length 6 | [optional] 
**UniqueReferenceNumber** | Pointer to **string** | Unique reference number for the transaction. Type: Alphanumeric [a-zA-Z 0-9], Maximum Length: 19 | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceType

`func (o *Transaction) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Transaction) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Transaction) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Transaction) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetAccountUri

`func (o *Transaction) GetAccountUri() string`

GetAccountUri returns the AccountUri field if non-nil, zero value otherwise.

### GetAccountUriOk

`func (o *Transaction) GetAccountUriOk() (*string, bool)`

GetAccountUriOk returns a tuple with the AccountUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUri

`func (o *Transaction) SetAccountUri(v string)`

SetAccountUri sets AccountUri field to given value.

### HasAccountUri

`func (o *Transaction) HasAccountUri() bool`

HasAccountUri returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *Transaction) GetTransactionAmount() TransactionAmount`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *Transaction) GetTransactionAmountOk() (*TransactionAmount, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *Transaction) SetTransactionAmount(v TransactionAmount)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *Transaction) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetNetwork

`func (o *Transaction) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Transaction) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Transaction) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Transaction) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetFundsAvailability

`func (o *Transaction) GetFundsAvailability() string`

GetFundsAvailability returns the FundsAvailability field if non-nil, zero value otherwise.

### GetFundsAvailabilityOk

`func (o *Transaction) GetFundsAvailabilityOk() (*string, bool)`

GetFundsAvailabilityOk returns a tuple with the FundsAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsAvailability

`func (o *Transaction) SetFundsAvailability(v string)`

SetFundsAvailability sets FundsAvailability field to given value.

### HasFundsAvailability

`func (o *Transaction) HasFundsAvailability() bool`

HasFundsAvailability returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *Transaction) GetCreateTimestamp() string`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *Transaction) GetCreateTimestampOk() (*string, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *Transaction) SetCreateTimestamp(v string)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *Transaction) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *Transaction) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *Transaction) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *Transaction) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *Transaction) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *Transaction) GetStatusTimestamp() string`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *Transaction) GetStatusTimestampOk() (*string, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *Transaction) SetStatusTimestamp(v string)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *Transaction) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.

### GetRetrievalReference

`func (o *Transaction) GetRetrievalReference() string`

GetRetrievalReference returns the RetrievalReference field if non-nil, zero value otherwise.

### GetRetrievalReferenceOk

`func (o *Transaction) GetRetrievalReferenceOk() (*string, bool)`

GetRetrievalReferenceOk returns a tuple with the RetrievalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievalReference

`func (o *Transaction) SetRetrievalReference(v string)`

SetRetrievalReference sets RetrievalReference field to given value.

### HasRetrievalReference

`func (o *Transaction) HasRetrievalReference() bool`

HasRetrievalReference returns a boolean if a field has been set.

### GetSystemTraceAuditNumber

`func (o *Transaction) GetSystemTraceAuditNumber() string`

GetSystemTraceAuditNumber returns the SystemTraceAuditNumber field if non-nil, zero value otherwise.

### GetSystemTraceAuditNumberOk

`func (o *Transaction) GetSystemTraceAuditNumberOk() (*string, bool)`

GetSystemTraceAuditNumberOk returns a tuple with the SystemTraceAuditNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTraceAuditNumber

`func (o *Transaction) SetSystemTraceAuditNumber(v string)`

SetSystemTraceAuditNumber sets SystemTraceAuditNumber field to given value.

### HasSystemTraceAuditNumber

`func (o *Transaction) HasSystemTraceAuditNumber() bool`

HasSystemTraceAuditNumber returns a boolean if a field has been set.

### GetUniqueReferenceNumber

`func (o *Transaction) GetUniqueReferenceNumber() string`

GetUniqueReferenceNumber returns the UniqueReferenceNumber field if non-nil, zero value otherwise.

### GetUniqueReferenceNumberOk

`func (o *Transaction) GetUniqueReferenceNumberOk() (*string, bool)`

GetUniqueReferenceNumberOk returns a tuple with the UniqueReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueReferenceNumber

`func (o *Transaction) SetUniqueReferenceNumber(v string)`

SetUniqueReferenceNumber sets UniqueReferenceNumber field to given value.

### HasUniqueReferenceNumber

`func (o *Transaction) HasUniqueReferenceNumber() bool`

HasUniqueReferenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


