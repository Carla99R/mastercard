# TransferGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The system-generated Transfer ID for the payment transfer. Alphanumeric and * , - . _ ~. | [optional] 
**ResourceType** | Pointer to **string** | The type of resource being returned: transfer. | [optional] 
**TransferReference** | Pointer to **string** | The partner-assigned unique reference identifier for the payment transfer. Alphanumeric and * , - . _ ~. Length 6-40. | [optional] 
**PaymentType** | Pointer to **string** | Payment type of the transfer: * A2A &#x3D; General Transfer to Own Account * ACO &#x3D; Agent Cash Out * C2C &#x3D; Cash to Card * CBP &#x3D; Payment of Own Credit Card Bill * P2P &#x3D; General Person-to-Person Transfer | [optional] 
**PaymentOriginationCountry** | Pointer to **string** | Country where the payment originated from as an ISO 3166-1 alpha-3 country code, in uppercase. Details- Alpha, Length: 3 | [optional] 
**SenderAccountUri** | Pointer to **string** | URI describing the sending account. It will include masked account information (e.g. \&quot;************9921\&quot; for a card account) but will not include security codes (e.g. CVC or expiration date for a card account). See [Account URIs](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/account-uris/). | [optional] 
**Sender** | Pointer to [**P2pSender**](P2pSender.md) |  | [optional] 
**FundingSource** | Pointer to **string** | Funding source must contain one of the following: CREDIT, DEBIT, PREPAID, DEPOSIT_ACCOUNT, MOBILE_MONEY_ACCOUNT, CASH or OTHER. | [optional] 
**TransactionPurpose** | Pointer to **string** | Purpose of the transaction. Valid values: &#39;00&#39; Family Support, &#39;01&#39; Regular Labor Transfers (expatriates), &#39;02&#39; Travel &amp; Tourism, &#39;03&#39; Education, &#39;04&#39; Hospitalization &amp; Medical Treatment, &#39;05&#39; Emergency Need, &#39;06&#39; Savings, &#39;07&#39; Gifts, &#39;08&#39; Other, &#39;09&#39; Salary, &#39;10&#39; Crowd lending, &#39;11&#39; Crypto currency, &#39;12&#39; Refund to original card, &#39;13&#39; Refund to new card. String, numeric, length 2. | [optional] 
**RecipientAccountUri** | Pointer to **string** | URI describing the receiving account. It will include masked account information (e.g. \&quot;************9913\&quot; for a card account) but will not include security codes (e.g. CVC or expiration date for a card account). See [Account URIs](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/account-uris/). | [optional] 
**Recipient** | Pointer to [**P2pRecipient**](P2pRecipient.md) |  | [optional] 
**SanctionScreeningOverride** | Pointer to **bool** | This field is not used at this time. The Sanctions Screening Service is suspended until further notice. | [optional] 
**TransferAmount** | Pointer to [**TransferAmount**](TransferAmount.md) |  | [optional] 
**Created** | Pointer to **string** | Date and time the original transfer was created, as an ISO 8601 timestamp in the format YYYY-MM-DDTHH:MM:SS±hh:mm. Refer to &#39;Date and Time Formats&#39;. | [optional] 
**TransactionHistory** | Pointer to [**TransactionHistoryGet**](TransactionHistoryGet.md) |  | [optional] 
**ReconciliationData** | Pointer to [**P2pReconciliationData**](P2pReconciliationData.md) |  | [optional] 
**StatementDescriptor** | Pointer to **string** | The statement descriptor is the value that will be displayed on the recipient&#39;s bank or card statement. It consists of two parts: the prefix and the content. The prefix is a short string typically used to identify the partner. The appended &lt;prefix&gt;*&lt;content&gt; will be displayed on the recipient&#39;s statement. Note: While most financial institutions display this information consistently, some may display it incorrectly or not at all. | [optional] 
**Channel** | Pointer to **string** | Initiation channel of the transfer request. Values: WEB, MOBILE, BANK, KIOSK. | [optional] 
**DeviceId** | Pointer to **string** | The serial number of a device that initiated the transfer. | [optional] 
**Location** | Pointer to **string** | Location where the transaction is initiated. | [optional] 
**OriginalStatus** | Pointer to **string** | Original status of the transfer. One of APPROVED, DECLINED, UNKNOWN, ERROR, PENDING, CANCELLED. CANCELLED is possible for non-card transactions (not applicable to card). | [optional] 
**Status** | Pointer to **string** | Status of the transfer. One of APPROVED, DECLINED, UNKNOWN, ERROR, PENDING, REVERSED, CANCELLED. PENDING indicates the transaction is still in progress. If you get PENDING for a card transaction, retry with another GET request. CANCELLED is possible for non-card transactions (not applicable to card). | [optional] 
**StatusTimestamp** | Pointer to **string** | Date and time the status was changed to its current value, as an ISO 8601 timestamp in the format YYYY-MM-DDTHH:MM:SS±hh:mm. Refer to &#39;Date and Time Formats&#39;. | [optional] 
**Participant** | Pointer to [**ResponseParticipant**](ResponseParticipant.md) |  | [optional] 
**TransferPriority** | Pointer to **string** | Type of payment instruction to send. Applicable for non-card only. Values: SIP - Standard immediate payment (default), FDP - Future dated payment | [optional] 
**FasterPaymentsId** | Pointer to **string** | Unique identifier of the payment within the scheme. 42 character string, in which the last five characters are almost always spaces. Applicable for non-card only. | [optional] 
**PaymentUid** | Pointer to **string** | Unique identifier of the payment the notification relates to. Applicable for non-card only. | [optional] 
**CanadaDomesticIndicator** | Pointer to **bool** | If a Canadian OI, TI or Reseller has elected to include Debit Mastercard in scope, the value should be true. Mastercard will decline a transaction if the value is not passed as true and it is a Canada DMC (Network Response: &#39;81&#39; Domestic Debit Transaction Not Allowed). | [optional] 

## Methods

### NewTransferGet

`func NewTransferGet() *TransferGet`

NewTransferGet instantiates a new TransferGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferGetWithDefaults

`func NewTransferGetWithDefaults() *TransferGet`

NewTransferGetWithDefaults instantiates a new TransferGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransferGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferGet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TransferGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceType

`func (o *TransferGet) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TransferGet) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TransferGet) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *TransferGet) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetTransferReference

`func (o *TransferGet) GetTransferReference() string`

GetTransferReference returns the TransferReference field if non-nil, zero value otherwise.

### GetTransferReferenceOk

`func (o *TransferGet) GetTransferReferenceOk() (*string, bool)`

GetTransferReferenceOk returns a tuple with the TransferReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferReference

`func (o *TransferGet) SetTransferReference(v string)`

SetTransferReference sets TransferReference field to given value.

### HasTransferReference

`func (o *TransferGet) HasTransferReference() bool`

HasTransferReference returns a boolean if a field has been set.

### GetPaymentType

`func (o *TransferGet) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *TransferGet) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *TransferGet) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *TransferGet) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentOriginationCountry

`func (o *TransferGet) GetPaymentOriginationCountry() string`

GetPaymentOriginationCountry returns the PaymentOriginationCountry field if non-nil, zero value otherwise.

### GetPaymentOriginationCountryOk

`func (o *TransferGet) GetPaymentOriginationCountryOk() (*string, bool)`

GetPaymentOriginationCountryOk returns a tuple with the PaymentOriginationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOriginationCountry

`func (o *TransferGet) SetPaymentOriginationCountry(v string)`

SetPaymentOriginationCountry sets PaymentOriginationCountry field to given value.

### HasPaymentOriginationCountry

`func (o *TransferGet) HasPaymentOriginationCountry() bool`

HasPaymentOriginationCountry returns a boolean if a field has been set.

### GetSenderAccountUri

`func (o *TransferGet) GetSenderAccountUri() string`

GetSenderAccountUri returns the SenderAccountUri field if non-nil, zero value otherwise.

### GetSenderAccountUriOk

`func (o *TransferGet) GetSenderAccountUriOk() (*string, bool)`

GetSenderAccountUriOk returns a tuple with the SenderAccountUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccountUri

`func (o *TransferGet) SetSenderAccountUri(v string)`

SetSenderAccountUri sets SenderAccountUri field to given value.

### HasSenderAccountUri

`func (o *TransferGet) HasSenderAccountUri() bool`

HasSenderAccountUri returns a boolean if a field has been set.

### GetSender

`func (o *TransferGet) GetSender() P2pSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TransferGet) GetSenderOk() (*P2pSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TransferGet) SetSender(v P2pSender)`

SetSender sets Sender field to given value.

### HasSender

`func (o *TransferGet) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetFundingSource

`func (o *TransferGet) GetFundingSource() string`

GetFundingSource returns the FundingSource field if non-nil, zero value otherwise.

### GetFundingSourceOk

`func (o *TransferGet) GetFundingSourceOk() (*string, bool)`

GetFundingSourceOk returns a tuple with the FundingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundingSource

`func (o *TransferGet) SetFundingSource(v string)`

SetFundingSource sets FundingSource field to given value.

### HasFundingSource

`func (o *TransferGet) HasFundingSource() bool`

HasFundingSource returns a boolean if a field has been set.

### GetTransactionPurpose

`func (o *TransferGet) GetTransactionPurpose() string`

GetTransactionPurpose returns the TransactionPurpose field if non-nil, zero value otherwise.

### GetTransactionPurposeOk

`func (o *TransferGet) GetTransactionPurposeOk() (*string, bool)`

GetTransactionPurposeOk returns a tuple with the TransactionPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionPurpose

`func (o *TransferGet) SetTransactionPurpose(v string)`

SetTransactionPurpose sets TransactionPurpose field to given value.

### HasTransactionPurpose

`func (o *TransferGet) HasTransactionPurpose() bool`

HasTransactionPurpose returns a boolean if a field has been set.

### GetRecipientAccountUri

`func (o *TransferGet) GetRecipientAccountUri() string`

GetRecipientAccountUri returns the RecipientAccountUri field if non-nil, zero value otherwise.

### GetRecipientAccountUriOk

`func (o *TransferGet) GetRecipientAccountUriOk() (*string, bool)`

GetRecipientAccountUriOk returns a tuple with the RecipientAccountUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAccountUri

`func (o *TransferGet) SetRecipientAccountUri(v string)`

SetRecipientAccountUri sets RecipientAccountUri field to given value.

### HasRecipientAccountUri

`func (o *TransferGet) HasRecipientAccountUri() bool`

HasRecipientAccountUri returns a boolean if a field has been set.

### GetRecipient

`func (o *TransferGet) GetRecipient() P2pRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TransferGet) GetRecipientOk() (*P2pRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TransferGet) SetRecipient(v P2pRecipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *TransferGet) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetSanctionScreeningOverride

`func (o *TransferGet) GetSanctionScreeningOverride() bool`

GetSanctionScreeningOverride returns the SanctionScreeningOverride field if non-nil, zero value otherwise.

### GetSanctionScreeningOverrideOk

`func (o *TransferGet) GetSanctionScreeningOverrideOk() (*bool, bool)`

GetSanctionScreeningOverrideOk returns a tuple with the SanctionScreeningOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionScreeningOverride

`func (o *TransferGet) SetSanctionScreeningOverride(v bool)`

SetSanctionScreeningOverride sets SanctionScreeningOverride field to given value.

### HasSanctionScreeningOverride

`func (o *TransferGet) HasSanctionScreeningOverride() bool`

HasSanctionScreeningOverride returns a boolean if a field has been set.

### GetTransferAmount

`func (o *TransferGet) GetTransferAmount() TransferAmount`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *TransferGet) GetTransferAmountOk() (*TransferAmount, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *TransferGet) SetTransferAmount(v TransferAmount)`

SetTransferAmount sets TransferAmount field to given value.

### HasTransferAmount

`func (o *TransferGet) HasTransferAmount() bool`

HasTransferAmount returns a boolean if a field has been set.

### GetCreated

`func (o *TransferGet) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransferGet) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransferGet) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TransferGet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetTransactionHistory

`func (o *TransferGet) GetTransactionHistory() TransactionHistoryGet`

GetTransactionHistory returns the TransactionHistory field if non-nil, zero value otherwise.

### GetTransactionHistoryOk

`func (o *TransferGet) GetTransactionHistoryOk() (*TransactionHistoryGet, bool)`

GetTransactionHistoryOk returns a tuple with the TransactionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHistory

`func (o *TransferGet) SetTransactionHistory(v TransactionHistoryGet)`

SetTransactionHistory sets TransactionHistory field to given value.

### HasTransactionHistory

`func (o *TransferGet) HasTransactionHistory() bool`

HasTransactionHistory returns a boolean if a field has been set.

### GetReconciliationData

`func (o *TransferGet) GetReconciliationData() P2pReconciliationData`

GetReconciliationData returns the ReconciliationData field if non-nil, zero value otherwise.

### GetReconciliationDataOk

`func (o *TransferGet) GetReconciliationDataOk() (*P2pReconciliationData, bool)`

GetReconciliationDataOk returns a tuple with the ReconciliationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationData

`func (o *TransferGet) SetReconciliationData(v P2pReconciliationData)`

SetReconciliationData sets ReconciliationData field to given value.

### HasReconciliationData

`func (o *TransferGet) HasReconciliationData() bool`

HasReconciliationData returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *TransferGet) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *TransferGet) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *TransferGet) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *TransferGet) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetChannel

`func (o *TransferGet) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TransferGet) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TransferGet) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TransferGet) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDeviceId

`func (o *TransferGet) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *TransferGet) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *TransferGet) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *TransferGet) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetLocation

`func (o *TransferGet) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TransferGet) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TransferGet) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TransferGet) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOriginalStatus

`func (o *TransferGet) GetOriginalStatus() string`

GetOriginalStatus returns the OriginalStatus field if non-nil, zero value otherwise.

### GetOriginalStatusOk

`func (o *TransferGet) GetOriginalStatusOk() (*string, bool)`

GetOriginalStatusOk returns a tuple with the OriginalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStatus

`func (o *TransferGet) SetOriginalStatus(v string)`

SetOriginalStatus sets OriginalStatus field to given value.

### HasOriginalStatus

`func (o *TransferGet) HasOriginalStatus() bool`

HasOriginalStatus returns a boolean if a field has been set.

### GetStatus

`func (o *TransferGet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferGet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferGet) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransferGet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *TransferGet) GetStatusTimestamp() string`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *TransferGet) GetStatusTimestampOk() (*string, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *TransferGet) SetStatusTimestamp(v string)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *TransferGet) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.

### GetParticipant

`func (o *TransferGet) GetParticipant() ResponseParticipant`

GetParticipant returns the Participant field if non-nil, zero value otherwise.

### GetParticipantOk

`func (o *TransferGet) GetParticipantOk() (*ResponseParticipant, bool)`

GetParticipantOk returns a tuple with the Participant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipant

`func (o *TransferGet) SetParticipant(v ResponseParticipant)`

SetParticipant sets Participant field to given value.

### HasParticipant

`func (o *TransferGet) HasParticipant() bool`

HasParticipant returns a boolean if a field has been set.

### GetTransferPriority

`func (o *TransferGet) GetTransferPriority() string`

GetTransferPriority returns the TransferPriority field if non-nil, zero value otherwise.

### GetTransferPriorityOk

`func (o *TransferGet) GetTransferPriorityOk() (*string, bool)`

GetTransferPriorityOk returns a tuple with the TransferPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferPriority

`func (o *TransferGet) SetTransferPriority(v string)`

SetTransferPriority sets TransferPriority field to given value.

### HasTransferPriority

`func (o *TransferGet) HasTransferPriority() bool`

HasTransferPriority returns a boolean if a field has been set.

### GetFasterPaymentsId

`func (o *TransferGet) GetFasterPaymentsId() string`

GetFasterPaymentsId returns the FasterPaymentsId field if non-nil, zero value otherwise.

### GetFasterPaymentsIdOk

`func (o *TransferGet) GetFasterPaymentsIdOk() (*string, bool)`

GetFasterPaymentsIdOk returns a tuple with the FasterPaymentsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFasterPaymentsId

`func (o *TransferGet) SetFasterPaymentsId(v string)`

SetFasterPaymentsId sets FasterPaymentsId field to given value.

### HasFasterPaymentsId

`func (o *TransferGet) HasFasterPaymentsId() bool`

HasFasterPaymentsId returns a boolean if a field has been set.

### GetPaymentUid

`func (o *TransferGet) GetPaymentUid() string`

GetPaymentUid returns the PaymentUid field if non-nil, zero value otherwise.

### GetPaymentUidOk

`func (o *TransferGet) GetPaymentUidOk() (*string, bool)`

GetPaymentUidOk returns a tuple with the PaymentUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUid

`func (o *TransferGet) SetPaymentUid(v string)`

SetPaymentUid sets PaymentUid field to given value.

### HasPaymentUid

`func (o *TransferGet) HasPaymentUid() bool`

HasPaymentUid returns a boolean if a field has been set.

### GetCanadaDomesticIndicator

`func (o *TransferGet) GetCanadaDomesticIndicator() bool`

GetCanadaDomesticIndicator returns the CanadaDomesticIndicator field if non-nil, zero value otherwise.

### GetCanadaDomesticIndicatorOk

`func (o *TransferGet) GetCanadaDomesticIndicatorOk() (*bool, bool)`

GetCanadaDomesticIndicatorOk returns a tuple with the CanadaDomesticIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanadaDomesticIndicator

`func (o *TransferGet) SetCanadaDomesticIndicator(v bool)`

SetCanadaDomesticIndicator sets CanadaDomesticIndicator field to given value.

### HasCanadaDomesticIndicator

`func (o *TransferGet) HasCanadaDomesticIndicator() bool`

HasCanadaDomesticIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


