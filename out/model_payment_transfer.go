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

// PaymentTransfer Contains the details of the request message.
type PaymentTransfer struct {
	// Required. A partner-assigned unique reference identifier for the payment transfer. Alphanumeric and * , - . _ ~. Length 6-40.
	TransferReference string `json:"transfer_reference"`
	// Conditional. Payment type of the transfer. The allowed payment types for the partner must be enabled by Mastercard during onboarding. If the partner is enabled for multiple payment types, the appropriate payment type must be provided. Otherwise the field is optional. Valid values: * A2A = General Transfer to Own Account * ACO = Agent Cash Out * C2C = Cash to Card * CBP = Payment of Own Credit Card Bill * P2P = General Person-to-Person Transfer
	PaymentType *string `json:"payment_type,omitempty"`
	// Required. The amount to be transferred. Numeric integer, 1-999999999999. The decimal point is implied based on the relevant `currency` exponent. For example, a US Dollar $53 amount is a value of 5300.
	Amount string `json:"amount"`
	// Required. The currency of the `amount` as an ISO 4217 uppercase alpha-3 currency code; see [Country and Currency Codes](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/country-currency-codes/). For example, US Dollars is USD.
	Currency string `json:"currency"`
	// Required. URI identifying the sending account. For the allowed schemes and format, see [Account URIs](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/account-uris/). See also [Funding Source](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/funding-source/).
	SenderAccountUri string `json:"sender_account_uri"`
	Sender *Sender `json:"sender,omitempty"`
	// Message a financial institution will associate to the transfer and may display. Details- 1-65
	AdditionalMessage *string `json:"additional_message,omitempty"`
	// Funding source must contain one of the following: CREDIT, DEBIT, PREPAID, DEPOSIT_ACCOUNT, MOBILE_MONEY_ACCOUNT, CASH or OTHER.
	FundingSource string `json:"funding_source"`
	// Purpose of the transaction. Valid values: '00' Family Support, '01'??Regular Labor Transfers (expatriates), '02' Travel & Tourism, '03' Education, '04' Hospitalization & Medical Treatment, '05' Emergency Need, '06' Savings, '07' Gifts, '08' Other, '09' Salary, '10' Crowd lending, '11' Crypto currency,??'12' Refund to original card, '13' Refund to new card. String, numeric, length 2.
	TransactionPurpose *string `json:"transaction_purpose,omitempty"`
	// Participation identifier of the sender. The receiving financial institution will associate the value to the transfer. Details- 1-30
	ParticipationId *string `json:"participation_id,omitempty"`
	// If a Canadian OI, TI or Reseller has elected to include Debit Mastercard in scope, the value should be true. Mastercard will decline a transaction if the value is not passed as true and it is a Canada DMC (Network Response: '81' Domestic Debit Transaction Not Allowed).
	CanadaDomesticIndicator *bool `json:"canada_domestic_indicator,omitempty"`
	TokenCryptogram *TokenCryptogram `json:"token_cryptogram,omitempty"`
	// The Directory Server Transaction ID is generated by the EMV 3DS Mastercard Directory Server during the authentication transaction and passed back to the merchant with the authentication results ds_transaction_id allows the OI to pass the Directory Server Transaction ID during authorization in order to link authentication and authorization data. Values- ans-36
	DsTransactionId *string `json:"ds_transaction_id,omitempty"`
	Authentication *Authentication `json:"authentication,omitempty"`
	// Required. URI identifying the receiving account. For the allowed schemes and format, see [Account URIs](https://developer.mastercard.com/mastercard-send-person-to-person/documentation/field-uris-codes/account-uris/).
	RecipientAccountUri string `json:"recipient_account_uri"`
	Recipient *Recipient `json:"recipient,omitempty"`
	// The country where the payment originates from as an ISO 3166-1 alpha-3 country code, in uppercase. Details - Conditional. If provided, this should match a valid country configured for the partner during onboarding. If the partner is configured for multiple origination countries this field is required and must be provided. Alpha, length: 3
	PaymentOriginationCountry *string `json:"payment_origination_country,omitempty"`
	// The Sanctions Screening Service is suspended until further notice. However, you must provide the value 'true'.
	SanctionScreeningOverride *bool `json:"sanction_screening_override,omitempty"`
	ReconciliationData *ReconciliationData `json:"reconciliation_data,omitempty"`
	// Initiation channel of the transfer request. This value can be defined in the onboarding process instead of passing in every call. Values: WEB, MOBILE, BANK, KIOSK. Details- Conditional
	Channel *string `json:"channel,omitempty"`
	// The statement descriptor is the value that will be displayed on the recipient's bank or card statement. It consists of two parts: the prefix and the content. The prefix is a short string typically used to identify the partner. It is defined during onboarding and the same value can be provided in the API call. If not provided in the API call, system will use the value defined in the onboarding process. The appended <prefix>*<content> will be displayed on the recipient's statement. The overall length may be at most 22 characters, including the prefix and the content. Note: While most financial institutions display this information consistently, some may display it per their discretion or not at all. Details- 1-22
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// The serial number of a device initiating the transfer. Details- 1-40
	DeviceId *string `json:"device_id,omitempty"`
	// Location where the transaction is initiated.
	Location *string `json:"location,omitempty"`
	FundingTransactionReference *FundingTransactionReference `json:"funding_transaction_reference,omitempty"`
	Participant *RequestParticipant `json:"participant,omitempty"`
	// Type of payment instruction to send. Applicable for non-card only. Values: SIP - Standard immediate payment (default), FDP - Future dated payment
	TransferPriority *string `json:"transfer_priority,omitempty"`
}

// NewPaymentTransfer instantiates a new PaymentTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentTransfer(transferReference string, amount string, currency string, senderAccountUri string, fundingSource string, recipientAccountUri string) *PaymentTransfer {
	this := PaymentTransfer{}
	this.TransferReference = transferReference
	this.Amount = amount
	this.Currency = currency
	this.SenderAccountUri = senderAccountUri
	this.FundingSource = fundingSource
	this.RecipientAccountUri = recipientAccountUri
	return &this
}

// NewPaymentTransferWithDefaults instantiates a new PaymentTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentTransferWithDefaults() *PaymentTransfer {
	this := PaymentTransfer{}
	return &this
}

// GetTransferReference returns the TransferReference field value
func (o *PaymentTransfer) GetTransferReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferReference
}

// GetTransferReferenceOk returns a tuple with the TransferReference field value
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetTransferReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferReference, true
}

// SetTransferReference sets field value
func (o *PaymentTransfer) SetTransferReference(v string) {
	o.TransferReference = v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *PaymentTransfer) GetPaymentType() string {
	if o == nil || o.PaymentType == nil {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetPaymentTypeOk() (*string, bool) {
	if o == nil || o.PaymentType == nil {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *PaymentTransfer) HasPaymentType() bool {
	if o != nil && o.PaymentType != nil {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *PaymentTransfer) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetAmount returns the Amount field value
func (o *PaymentTransfer) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentTransfer) SetAmount(v string) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *PaymentTransfer) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentTransfer) SetCurrency(v string) {
	o.Currency = v
}

// GetSenderAccountUri returns the SenderAccountUri field value
func (o *PaymentTransfer) GetSenderAccountUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderAccountUri
}

// GetSenderAccountUriOk returns a tuple with the SenderAccountUri field value
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetSenderAccountUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderAccountUri, true
}

// SetSenderAccountUri sets field value
func (o *PaymentTransfer) SetSenderAccountUri(v string) {
	o.SenderAccountUri = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *PaymentTransfer) GetSender() Sender {
	if o == nil || o.Sender == nil {
		var ret Sender
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetSenderOk() (*Sender, bool) {
	if o == nil || o.Sender == nil {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *PaymentTransfer) HasSender() bool {
	if o != nil && o.Sender != nil {
		return true
	}

	return false
}

// SetSender gets a reference to the given Sender and assigns it to the Sender field.
func (o *PaymentTransfer) SetSender(v Sender) {
	o.Sender = &v
}

// GetAdditionalMessage returns the AdditionalMessage field value if set, zero value otherwise.
func (o *PaymentTransfer) GetAdditionalMessage() string {
	if o == nil || o.AdditionalMessage == nil {
		var ret string
		return ret
	}
	return *o.AdditionalMessage
}

// GetAdditionalMessageOk returns a tuple with the AdditionalMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetAdditionalMessageOk() (*string, bool) {
	if o == nil || o.AdditionalMessage == nil {
		return nil, false
	}
	return o.AdditionalMessage, true
}

// HasAdditionalMessage returns a boolean if a field has been set.
func (o *PaymentTransfer) HasAdditionalMessage() bool {
	if o != nil && o.AdditionalMessage != nil {
		return true
	}

	return false
}

// SetAdditionalMessage gets a reference to the given string and assigns it to the AdditionalMessage field.
func (o *PaymentTransfer) SetAdditionalMessage(v string) {
	o.AdditionalMessage = &v
}

// GetFundingSource returns the FundingSource field value
func (o *PaymentTransfer) GetFundingSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetFundingSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FundingSource, true
}

// SetFundingSource sets field value
func (o *PaymentTransfer) SetFundingSource(v string) {
	o.FundingSource = v
}

// GetTransactionPurpose returns the TransactionPurpose field value if set, zero value otherwise.
func (o *PaymentTransfer) GetTransactionPurpose() string {
	if o == nil || o.TransactionPurpose == nil {
		var ret string
		return ret
	}
	return *o.TransactionPurpose
}

// GetTransactionPurposeOk returns a tuple with the TransactionPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetTransactionPurposeOk() (*string, bool) {
	if o == nil || o.TransactionPurpose == nil {
		return nil, false
	}
	return o.TransactionPurpose, true
}

// HasTransactionPurpose returns a boolean if a field has been set.
func (o *PaymentTransfer) HasTransactionPurpose() bool {
	if o != nil && o.TransactionPurpose != nil {
		return true
	}

	return false
}

// SetTransactionPurpose gets a reference to the given string and assigns it to the TransactionPurpose field.
func (o *PaymentTransfer) SetTransactionPurpose(v string) {
	o.TransactionPurpose = &v
}

// GetParticipationId returns the ParticipationId field value if set, zero value otherwise.
func (o *PaymentTransfer) GetParticipationId() string {
	if o == nil || o.ParticipationId == nil {
		var ret string
		return ret
	}
	return *o.ParticipationId
}

// GetParticipationIdOk returns a tuple with the ParticipationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetParticipationIdOk() (*string, bool) {
	if o == nil || o.ParticipationId == nil {
		return nil, false
	}
	return o.ParticipationId, true
}

// HasParticipationId returns a boolean if a field has been set.
func (o *PaymentTransfer) HasParticipationId() bool {
	if o != nil && o.ParticipationId != nil {
		return true
	}

	return false
}

// SetParticipationId gets a reference to the given string and assigns it to the ParticipationId field.
func (o *PaymentTransfer) SetParticipationId(v string) {
	o.ParticipationId = &v
}

// GetCanadaDomesticIndicator returns the CanadaDomesticIndicator field value if set, zero value otherwise.
func (o *PaymentTransfer) GetCanadaDomesticIndicator() bool {
	if o == nil || o.CanadaDomesticIndicator == nil {
		var ret bool
		return ret
	}
	return *o.CanadaDomesticIndicator
}

// GetCanadaDomesticIndicatorOk returns a tuple with the CanadaDomesticIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetCanadaDomesticIndicatorOk() (*bool, bool) {
	if o == nil || o.CanadaDomesticIndicator == nil {
		return nil, false
	}
	return o.CanadaDomesticIndicator, true
}

// HasCanadaDomesticIndicator returns a boolean if a field has been set.
func (o *PaymentTransfer) HasCanadaDomesticIndicator() bool {
	if o != nil && o.CanadaDomesticIndicator != nil {
		return true
	}

	return false
}

// SetCanadaDomesticIndicator gets a reference to the given bool and assigns it to the CanadaDomesticIndicator field.
func (o *PaymentTransfer) SetCanadaDomesticIndicator(v bool) {
	o.CanadaDomesticIndicator = &v
}

// GetTokenCryptogram returns the TokenCryptogram field value if set, zero value otherwise.
func (o *PaymentTransfer) GetTokenCryptogram() TokenCryptogram {
	if o == nil || o.TokenCryptogram == nil {
		var ret TokenCryptogram
		return ret
	}
	return *o.TokenCryptogram
}

// GetTokenCryptogramOk returns a tuple with the TokenCryptogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetTokenCryptogramOk() (*TokenCryptogram, bool) {
	if o == nil || o.TokenCryptogram == nil {
		return nil, false
	}
	return o.TokenCryptogram, true
}

// HasTokenCryptogram returns a boolean if a field has been set.
func (o *PaymentTransfer) HasTokenCryptogram() bool {
	if o != nil && o.TokenCryptogram != nil {
		return true
	}

	return false
}

// SetTokenCryptogram gets a reference to the given TokenCryptogram and assigns it to the TokenCryptogram field.
func (o *PaymentTransfer) SetTokenCryptogram(v TokenCryptogram) {
	o.TokenCryptogram = &v
}

// GetDsTransactionId returns the DsTransactionId field value if set, zero value otherwise.
func (o *PaymentTransfer) GetDsTransactionId() string {
	if o == nil || o.DsTransactionId == nil {
		var ret string
		return ret
	}
	return *o.DsTransactionId
}

// GetDsTransactionIdOk returns a tuple with the DsTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetDsTransactionIdOk() (*string, bool) {
	if o == nil || o.DsTransactionId == nil {
		return nil, false
	}
	return o.DsTransactionId, true
}

// HasDsTransactionId returns a boolean if a field has been set.
func (o *PaymentTransfer) HasDsTransactionId() bool {
	if o != nil && o.DsTransactionId != nil {
		return true
	}

	return false
}

// SetDsTransactionId gets a reference to the given string and assigns it to the DsTransactionId field.
func (o *PaymentTransfer) SetDsTransactionId(v string) {
	o.DsTransactionId = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *PaymentTransfer) GetAuthentication() Authentication {
	if o == nil || o.Authentication == nil {
		var ret Authentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetAuthenticationOk() (*Authentication, bool) {
	if o == nil || o.Authentication == nil {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *PaymentTransfer) HasAuthentication() bool {
	if o != nil && o.Authentication != nil {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given Authentication and assigns it to the Authentication field.
func (o *PaymentTransfer) SetAuthentication(v Authentication) {
	o.Authentication = &v
}

// GetRecipientAccountUri returns the RecipientAccountUri field value
func (o *PaymentTransfer) GetRecipientAccountUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAccountUri
}

// GetRecipientAccountUriOk returns a tuple with the RecipientAccountUri field value
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetRecipientAccountUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientAccountUri, true
}

// SetRecipientAccountUri sets field value
func (o *PaymentTransfer) SetRecipientAccountUri(v string) {
	o.RecipientAccountUri = v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *PaymentTransfer) GetRecipient() Recipient {
	if o == nil || o.Recipient == nil {
		var ret Recipient
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetRecipientOk() (*Recipient, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *PaymentTransfer) HasRecipient() bool {
	if o != nil && o.Recipient != nil {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given Recipient and assigns it to the Recipient field.
func (o *PaymentTransfer) SetRecipient(v Recipient) {
	o.Recipient = &v
}

// GetPaymentOriginationCountry returns the PaymentOriginationCountry field value if set, zero value otherwise.
func (o *PaymentTransfer) GetPaymentOriginationCountry() string {
	if o == nil || o.PaymentOriginationCountry == nil {
		var ret string
		return ret
	}
	return *o.PaymentOriginationCountry
}

// GetPaymentOriginationCountryOk returns a tuple with the PaymentOriginationCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetPaymentOriginationCountryOk() (*string, bool) {
	if o == nil || o.PaymentOriginationCountry == nil {
		return nil, false
	}
	return o.PaymentOriginationCountry, true
}

// HasPaymentOriginationCountry returns a boolean if a field has been set.
func (o *PaymentTransfer) HasPaymentOriginationCountry() bool {
	if o != nil && o.PaymentOriginationCountry != nil {
		return true
	}

	return false
}

// SetPaymentOriginationCountry gets a reference to the given string and assigns it to the PaymentOriginationCountry field.
func (o *PaymentTransfer) SetPaymentOriginationCountry(v string) {
	o.PaymentOriginationCountry = &v
}

// GetSanctionScreeningOverride returns the SanctionScreeningOverride field value if set, zero value otherwise.
func (o *PaymentTransfer) GetSanctionScreeningOverride() bool {
	if o == nil || o.SanctionScreeningOverride == nil {
		var ret bool
		return ret
	}
	return *o.SanctionScreeningOverride
}

// GetSanctionScreeningOverrideOk returns a tuple with the SanctionScreeningOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetSanctionScreeningOverrideOk() (*bool, bool) {
	if o == nil || o.SanctionScreeningOverride == nil {
		return nil, false
	}
	return o.SanctionScreeningOverride, true
}

// HasSanctionScreeningOverride returns a boolean if a field has been set.
func (o *PaymentTransfer) HasSanctionScreeningOverride() bool {
	if o != nil && o.SanctionScreeningOverride != nil {
		return true
	}

	return false
}

// SetSanctionScreeningOverride gets a reference to the given bool and assigns it to the SanctionScreeningOverride field.
func (o *PaymentTransfer) SetSanctionScreeningOverride(v bool) {
	o.SanctionScreeningOverride = &v
}

// GetReconciliationData returns the ReconciliationData field value if set, zero value otherwise.
func (o *PaymentTransfer) GetReconciliationData() ReconciliationData {
	if o == nil || o.ReconciliationData == nil {
		var ret ReconciliationData
		return ret
	}
	return *o.ReconciliationData
}

// GetReconciliationDataOk returns a tuple with the ReconciliationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetReconciliationDataOk() (*ReconciliationData, bool) {
	if o == nil || o.ReconciliationData == nil {
		return nil, false
	}
	return o.ReconciliationData, true
}

// HasReconciliationData returns a boolean if a field has been set.
func (o *PaymentTransfer) HasReconciliationData() bool {
	if o != nil && o.ReconciliationData != nil {
		return true
	}

	return false
}

// SetReconciliationData gets a reference to the given ReconciliationData and assigns it to the ReconciliationData field.
func (o *PaymentTransfer) SetReconciliationData(v ReconciliationData) {
	o.ReconciliationData = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *PaymentTransfer) GetChannel() string {
	if o == nil || o.Channel == nil {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetChannelOk() (*string, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *PaymentTransfer) HasChannel() bool {
	if o != nil && o.Channel != nil {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *PaymentTransfer) SetChannel(v string) {
	o.Channel = &v
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise.
func (o *PaymentTransfer) GetStatementDescriptor() string {
	if o == nil || o.StatementDescriptor == nil {
		var ret string
		return ret
	}
	return *o.StatementDescriptor
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetStatementDescriptorOk() (*string, bool) {
	if o == nil || o.StatementDescriptor == nil {
		return nil, false
	}
	return o.StatementDescriptor, true
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *PaymentTransfer) HasStatementDescriptor() bool {
	if o != nil && o.StatementDescriptor != nil {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given string and assigns it to the StatementDescriptor field.
func (o *PaymentTransfer) SetStatementDescriptor(v string) {
	o.StatementDescriptor = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *PaymentTransfer) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *PaymentTransfer) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *PaymentTransfer) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PaymentTransfer) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PaymentTransfer) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PaymentTransfer) SetLocation(v string) {
	o.Location = &v
}

// GetFundingTransactionReference returns the FundingTransactionReference field value if set, zero value otherwise.
func (o *PaymentTransfer) GetFundingTransactionReference() FundingTransactionReference {
	if o == nil || o.FundingTransactionReference == nil {
		var ret FundingTransactionReference
		return ret
	}
	return *o.FundingTransactionReference
}

// GetFundingTransactionReferenceOk returns a tuple with the FundingTransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetFundingTransactionReferenceOk() (*FundingTransactionReference, bool) {
	if o == nil || o.FundingTransactionReference == nil {
		return nil, false
	}
	return o.FundingTransactionReference, true
}

// HasFundingTransactionReference returns a boolean if a field has been set.
func (o *PaymentTransfer) HasFundingTransactionReference() bool {
	if o != nil && o.FundingTransactionReference != nil {
		return true
	}

	return false
}

// SetFundingTransactionReference gets a reference to the given FundingTransactionReference and assigns it to the FundingTransactionReference field.
func (o *PaymentTransfer) SetFundingTransactionReference(v FundingTransactionReference) {
	o.FundingTransactionReference = &v
}

// GetParticipant returns the Participant field value if set, zero value otherwise.
func (o *PaymentTransfer) GetParticipant() RequestParticipant {
	if o == nil || o.Participant == nil {
		var ret RequestParticipant
		return ret
	}
	return *o.Participant
}

// GetParticipantOk returns a tuple with the Participant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetParticipantOk() (*RequestParticipant, bool) {
	if o == nil || o.Participant == nil {
		return nil, false
	}
	return o.Participant, true
}

// HasParticipant returns a boolean if a field has been set.
func (o *PaymentTransfer) HasParticipant() bool {
	if o != nil && o.Participant != nil {
		return true
	}

	return false
}

// SetParticipant gets a reference to the given RequestParticipant and assigns it to the Participant field.
func (o *PaymentTransfer) SetParticipant(v RequestParticipant) {
	o.Participant = &v
}

// GetTransferPriority returns the TransferPriority field value if set, zero value otherwise.
func (o *PaymentTransfer) GetTransferPriority() string {
	if o == nil || o.TransferPriority == nil {
		var ret string
		return ret
	}
	return *o.TransferPriority
}

// GetTransferPriorityOk returns a tuple with the TransferPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTransfer) GetTransferPriorityOk() (*string, bool) {
	if o == nil || o.TransferPriority == nil {
		return nil, false
	}
	return o.TransferPriority, true
}

// HasTransferPriority returns a boolean if a field has been set.
func (o *PaymentTransfer) HasTransferPriority() bool {
	if o != nil && o.TransferPriority != nil {
		return true
	}

	return false
}

// SetTransferPriority gets a reference to the given string and assigns it to the TransferPriority field.
func (o *PaymentTransfer) SetTransferPriority(v string) {
	o.TransferPriority = &v
}

func (o PaymentTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer_reference"] = o.TransferReference
	}
	if o.PaymentType != nil {
		toSerialize["payment_type"] = o.PaymentType
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["sender_account_uri"] = o.SenderAccountUri
	}
	if o.Sender != nil {
		toSerialize["sender"] = o.Sender
	}
	if o.AdditionalMessage != nil {
		toSerialize["additional_message"] = o.AdditionalMessage
	}
	if true {
		toSerialize["funding_source"] = o.FundingSource
	}
	if o.TransactionPurpose != nil {
		toSerialize["transaction_purpose"] = o.TransactionPurpose
	}
	if o.ParticipationId != nil {
		toSerialize["participation_id"] = o.ParticipationId
	}
	if o.CanadaDomesticIndicator != nil {
		toSerialize["canada_domestic_indicator"] = o.CanadaDomesticIndicator
	}
	if o.TokenCryptogram != nil {
		toSerialize["token_cryptogram"] = o.TokenCryptogram
	}
	if o.DsTransactionId != nil {
		toSerialize["ds_transaction_id"] = o.DsTransactionId
	}
	if o.Authentication != nil {
		toSerialize["authentication"] = o.Authentication
	}
	if true {
		toSerialize["recipient_account_uri"] = o.RecipientAccountUri
	}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	if o.PaymentOriginationCountry != nil {
		toSerialize["payment_origination_country"] = o.PaymentOriginationCountry
	}
	if o.SanctionScreeningOverride != nil {
		toSerialize["sanction_screening_override"] = o.SanctionScreeningOverride
	}
	if o.ReconciliationData != nil {
		toSerialize["reconciliation_data"] = o.ReconciliationData
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.StatementDescriptor != nil {
		toSerialize["statement_descriptor"] = o.StatementDescriptor
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.FundingTransactionReference != nil {
		toSerialize["funding_transaction_reference"] = o.FundingTransactionReference
	}
	if o.Participant != nil {
		toSerialize["participant"] = o.Participant
	}
	if o.TransferPriority != nil {
		toSerialize["transfer_priority"] = o.TransferPriority
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentTransfer struct {
	value *PaymentTransfer
	isSet bool
}

func (v NullablePaymentTransfer) Get() *PaymentTransfer {
	return v.value
}

func (v *NullablePaymentTransfer) Set(val *PaymentTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTransfer(val *PaymentTransfer) *NullablePaymentTransfer {
	return &NullablePaymentTransfer{value: val, isSet: true}
}

func (v NullablePaymentTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


