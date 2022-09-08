# TransactionGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID for the Payment Transaction. Alphanumeric and * , - . _ ~. | [optional] 
**ResourceType** | Pointer to **string** | The type of resource being returned: transaction. | [optional] 
**AccountUri** | Pointer to **string** | URI describing the receiving account. It will include masked account information (e.g. \&quot;************9913\&quot; for a card account) but will not include security codes (e.g. CVC or expiration date for a card account). See [Account URIs](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/account-uris/). | [optional] 
**TransactionAmount** | Pointer to [**TransactionAmount**](TransactionAmount.md) |  | [optional] 
**Network** | Pointer to **string** | Proposed network for the transaction. | [optional] 
**NetworkStatusCode** | Pointer to **string** | Network Status Code. | [optional] 
**NetworkStatusDescription** | Pointer to **string** | Network Status Description. | [optional] 
**FundsAvailability** | Pointer to **string** | An estimate of when the funds might be available. Actual Deposit Availability is determined by the financial institution. One of: IMMEDIATE, SAME_DAY, NEXT_CALENDAR_DAY, NEXT_BUSINESS_DAY, TWO_TO_FIVE_BUSINESS_DAYS, or UNKNOWN. | [optional] 
**Type** | Pointer to **string** | Type of the transaction. One of: FUNDING, PAYMENT, FUNDING_REVERSAL, or PAYMENT_REVERSAL. | [optional] 
**CreateTimestamp** | Pointer to **string** | Date and time the transaction was created, as an ISO 8601 timestamp in the format YYYY-MM-DDTHH:MM:SS±hh:mm. Refer to &#39;Date and Time Formats&#39;. | [optional] 
**Status** | Pointer to **string** | Status of the transfer. One of APPROVED, DECLINED, UNKNOWN, ERROR, PENDING. | [optional] 
**StatusReason** | Pointer to **string** | Reason for status. One of APPROVED, DECLINED, UNKNOWN, ERROR, PENDING, FRAUD, CARD_EXPIRED, LIMIT_EXCEEDED. | [optional] 
**StatusTimestamp** | Pointer to **string** | Date and time the status was changed to its current value, as an ISO 8601 timestamp in the format YYYY-MM-DDTHH:MM:SS±hh:mm. Refer to &#39;Date and Time Formats&#39;. | [optional] 
**RetrievalReference** | Pointer to **string** | Unique reference number that identifies the transaction at the network. Details- maxlength 24 | [optional] 
**SystemTraceAuditNumber** | Pointer to **string** | Unique system trace audit number for the transaction, the STAN (system trace audit number). Details- maxlength 6 | [optional] 
**UniqueReferenceNumber** | Pointer to **string** | Unique reference number for the transaction. Type: Alphanumeric [a-zA-Z 0-9], Maximum Length: 19 | [optional] 

## Methods

### NewTransactionGet

`func NewTransactionGet() *TransactionGet`

NewTransactionGet instantiates a new TransactionGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionGetWithDefaults

`func NewTransactionGetWithDefaults() *TransactionGet`

NewTransactionGetWithDefaults instantiates a new TransactionGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionGet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceType

`func (o *TransactionGet) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TransactionGet) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TransactionGet) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *TransactionGet) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetAccountUri

`func (o *TransactionGet) GetAccountUri() string`

GetAccountUri returns the AccountUri field if non-nil, zero value otherwise.

### GetAccountUriOk

`func (o *TransactionGet) GetAccountUriOk() (*string, bool)`

GetAccountUriOk returns a tuple with the AccountUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUri

`func (o *TransactionGet) SetAccountUri(v string)`

SetAccountUri sets AccountUri field to given value.

### HasAccountUri

`func (o *TransactionGet) HasAccountUri() bool`

HasAccountUri returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *TransactionGet) GetTransactionAmount() TransactionAmount`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *TransactionGet) GetTransactionAmountOk() (*TransactionAmount, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *TransactionGet) SetTransactionAmount(v TransactionAmount)`

SetTransactionAmount sets TransactionAmount field to given value.

### HasTransactionAmount

`func (o *TransactionGet) HasTransactionAmount() bool`

HasTransactionAmount returns a boolean if a field has been set.

### GetNetwork

`func (o *TransactionGet) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransactionGet) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransactionGet) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *TransactionGet) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkStatusCode

`func (o *TransactionGet) GetNetworkStatusCode() string`

GetNetworkStatusCode returns the NetworkStatusCode field if non-nil, zero value otherwise.

### GetNetworkStatusCodeOk

`func (o *TransactionGet) GetNetworkStatusCodeOk() (*string, bool)`

GetNetworkStatusCodeOk returns a tuple with the NetworkStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatusCode

`func (o *TransactionGet) SetNetworkStatusCode(v string)`

SetNetworkStatusCode sets NetworkStatusCode field to given value.

### HasNetworkStatusCode

`func (o *TransactionGet) HasNetworkStatusCode() bool`

HasNetworkStatusCode returns a boolean if a field has been set.

### GetNetworkStatusDescription

`func (o *TransactionGet) GetNetworkStatusDescription() string`

GetNetworkStatusDescription returns the NetworkStatusDescription field if non-nil, zero value otherwise.

### GetNetworkStatusDescriptionOk

`func (o *TransactionGet) GetNetworkStatusDescriptionOk() (*string, bool)`

GetNetworkStatusDescriptionOk returns a tuple with the NetworkStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStatusDescription

`func (o *TransactionGet) SetNetworkStatusDescription(v string)`

SetNetworkStatusDescription sets NetworkStatusDescription field to given value.

### HasNetworkStatusDescription

`func (o *TransactionGet) HasNetworkStatusDescription() bool`

HasNetworkStatusDescription returns a boolean if a field has been set.

### GetFundsAvailability

`func (o *TransactionGet) GetFundsAvailability() string`

GetFundsAvailability returns the FundsAvailability field if non-nil, zero value otherwise.

### GetFundsAvailabilityOk

`func (o *TransactionGet) GetFundsAvailabilityOk() (*string, bool)`

GetFundsAvailabilityOk returns a tuple with the FundsAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsAvailability

`func (o *TransactionGet) SetFundsAvailability(v string)`

SetFundsAvailability sets FundsAvailability field to given value.

### HasFundsAvailability

`func (o *TransactionGet) HasFundsAvailability() bool`

HasFundsAvailability returns a boolean if a field has been set.

### GetType

`func (o *TransactionGet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionGet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionGet) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionGet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *TransactionGet) GetCreateTimestamp() string`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *TransactionGet) GetCreateTimestampOk() (*string, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *TransactionGet) SetCreateTimestamp(v string)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *TransactionGet) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *TransactionGet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionGet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionGet) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionGet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *TransactionGet) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *TransactionGet) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *TransactionGet) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *TransactionGet) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *TransactionGet) GetStatusTimestamp() string`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *TransactionGet) GetStatusTimestampOk() (*string, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *TransactionGet) SetStatusTimestamp(v string)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *TransactionGet) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.

### GetRetrievalReference

`func (o *TransactionGet) GetRetrievalReference() string`

GetRetrievalReference returns the RetrievalReference field if non-nil, zero value otherwise.

### GetRetrievalReferenceOk

`func (o *TransactionGet) GetRetrievalReferenceOk() (*string, bool)`

GetRetrievalReferenceOk returns a tuple with the RetrievalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievalReference

`func (o *TransactionGet) SetRetrievalReference(v string)`

SetRetrievalReference sets RetrievalReference field to given value.

### HasRetrievalReference

`func (o *TransactionGet) HasRetrievalReference() bool`

HasRetrievalReference returns a boolean if a field has been set.

### GetSystemTraceAuditNumber

`func (o *TransactionGet) GetSystemTraceAuditNumber() string`

GetSystemTraceAuditNumber returns the SystemTraceAuditNumber field if non-nil, zero value otherwise.

### GetSystemTraceAuditNumberOk

`func (o *TransactionGet) GetSystemTraceAuditNumberOk() (*string, bool)`

GetSystemTraceAuditNumberOk returns a tuple with the SystemTraceAuditNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemTraceAuditNumber

`func (o *TransactionGet) SetSystemTraceAuditNumber(v string)`

SetSystemTraceAuditNumber sets SystemTraceAuditNumber field to given value.

### HasSystemTraceAuditNumber

`func (o *TransactionGet) HasSystemTraceAuditNumber() bool`

HasSystemTraceAuditNumber returns a boolean if a field has been set.

### GetUniqueReferenceNumber

`func (o *TransactionGet) GetUniqueReferenceNumber() string`

GetUniqueReferenceNumber returns the UniqueReferenceNumber field if non-nil, zero value otherwise.

### GetUniqueReferenceNumberOk

`func (o *TransactionGet) GetUniqueReferenceNumberOk() (*string, bool)`

GetUniqueReferenceNumberOk returns a tuple with the UniqueReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueReferenceNumber

`func (o *TransactionGet) SetUniqueReferenceNumber(v string)`

SetUniqueReferenceNumber sets UniqueReferenceNumber field to given value.

### HasUniqueReferenceNumber

`func (o *TransactionGet) HasUniqueReferenceNumber() bool`

HasUniqueReferenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


