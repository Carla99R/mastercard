/*
Send Payment Transfer API

Mastercard Send Payment Transfer API

API version: 1.0.1
Contact: apisupport@mastercard.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// PaymentTransferApiService PaymentTransferApi service
type PaymentTransferApiService service

type ApiCreatePaymentRequest struct {
	ctx context.Context
	ApiService *PaymentTransferApiService
	paymentTransferWrapper *PaymentTransferWrapper
	partnerId string
}

// Contains the details of the request message.
func (r ApiCreatePaymentRequest) PaymentTransferWrapper(paymentTransferWrapper PaymentTransferWrapper) ApiCreatePaymentRequest {
	r.paymentTransferWrapper = &paymentTransferWrapper
	return r
}

func (r ApiCreatePaymentRequest) Execute() (*TransferWrapper, *http.Response, error) {
	return r.ApiService.CreatePaymentExecute(r)
}

/*
CreatePayment Create a payment transfer (Payment Transaction) to send funds to a recipient

Each payment transfer has a unique `transfer_reference` ID, which you must provide in the request, and it will be assigned a unique system-generated ID (`id`), which is returned in the response. This API service supports the optional extra Idempotency feature. If you have that feature enabled (contact your Mastercard representative), you can make the same payment transfer request multiple times without worrying about the transaction being processed multiple times; see the Mastercard Send Person-to-Person documentation.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32.
 @return ApiCreatePaymentRequest
*/
func (a *PaymentTransferApiService) CreatePayment(ctx context.Context, partnerId string) ApiCreatePaymentRequest {
	return ApiCreatePaymentRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return TransferWrapper
func (a *PaymentTransferApiService) CreatePaymentExecute(r ApiCreatePaymentRequest) (*TransferWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransferWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentTransferApiService.CreatePayment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/partners/{partnerId}/transfers/payment"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.paymentTransferWrapper == nil {
		return localVarReturnValue, nil, reportError("paymentTransferWrapper is required and must be specified")
	}
	if strlen(r.partnerId) < 32 {
		return localVarReturnValue, nil, reportError("partnerId must have at least 32 elements")
	}
	if strlen(r.partnerId) > 32 {
		return localVarReturnValue, nil, reportError("partnerId must have less than 32 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.paymentTransferWrapper
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTransferByIdRequest struct {
	ctx context.Context
	ApiService *PaymentTransferApiService
	partnerId string
	transferId string
}

func (r ApiGetTransferByIdRequest) Execute() (*TransferWrapperGet, *http.Response, error) {
	return r.ApiService.GetTransferByIdExecute(r)
}

/*
GetTransferById Retrieve details of a payment transfer, including latest status, by using the transfer ID that was returned when you created the transfer

Example uses for this API call:

 - Get the network decline code and description when a transaction is declined

 - Get the latest status when a transaction response was 'UNKNOWN'

 - Transaction history and research purposes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32.
 @param transferId Required. The system-generated Transfer ID (`id`) that was returned when you created the payment transfer. Alphanumeric and * , - . _ ~.
 @return ApiGetTransferByIdRequest
*/
func (a *PaymentTransferApiService) GetTransferById(ctx context.Context, partnerId string, transferId string) ApiGetTransferByIdRequest {
	return ApiGetTransferByIdRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
		transferId: transferId,
	}
}

// Execute executes the request
//  @return TransferWrapperGet
func (a *PaymentTransferApiService) GetTransferByIdExecute(r ApiGetTransferByIdRequest) (*TransferWrapperGet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransferWrapperGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentTransferApiService.GetTransferById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/partners/{partnerId}/transfers/{transferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"transferId"+"}", url.PathEscape(parameterToString(r.transferId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.partnerId) < 32 {
		return localVarReturnValue, nil, reportError("partnerId must have at least 32 elements")
	}
	if strlen(r.partnerId) > 32 {
		return localVarReturnValue, nil, reportError("partnerId must have less than 32 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTransferByRefRequest struct {
	ctx context.Context
	ApiService *PaymentTransferApiService
	partnerId string
	ref *string
}

// Query parameter - Required. The unique &#x60;transfer_reference&#x60; that you provided when creating the payment transfer. Alphanumeric and * , - . _ ~. Length 6-40.
func (r ApiGetTransferByRefRequest) Ref(ref string) ApiGetTransferByRefRequest {
	r.ref = &ref
	return r
}

func (r ApiGetTransferByRefRequest) Execute() (*TransfersWrapper, *http.Response, error) {
	return r.ApiService.GetTransferByRefExecute(r)
}

/*
GetTransferByRef Retrieve details of a payment transfer, including latest status, by using the transfer_reference value you provided when you created the transfer

Example uses for this API call:

 - Get the network decline code and description when a transaction is declined

 - Get the latest status when a transaction response was 'UNKNOWN'

 - Transaction history and research purposes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param partnerId The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32.
 @return ApiGetTransferByRefRequest
*/
func (a *PaymentTransferApiService) GetTransferByRef(ctx context.Context, partnerId string) ApiGetTransferByRefRequest {
	return ApiGetTransferByRefRequest{
		ApiService: a,
		ctx: ctx,
		partnerId: partnerId,
	}
}

// Execute executes the request
//  @return TransfersWrapper
func (a *PaymentTransferApiService) GetTransferByRefExecute(r ApiGetTransferByRefRequest) (*TransfersWrapper, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransfersWrapper
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentTransferApiService.GetTransferByRef")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/partners/{partnerId}/transfers"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerId"+"}", url.PathEscape(parameterToString(r.partnerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.partnerId) < 32 {
		return localVarReturnValue, nil, reportError("partnerId must have at least 32 elements")
	}
	if strlen(r.partnerId) > 32 {
		return localVarReturnValue, nil, reportError("partnerId must have less than 32 elements")
	}
	if r.ref == nil {
		return localVarReturnValue, nil, reportError("ref is required and must be specified")
	}
	if strlen(*r.ref) < 6 {
		return localVarReturnValue, nil, reportError("ref must have at least 6 elements")
	}
	if strlen(*r.ref) > 40 {
		return localVarReturnValue, nil, reportError("ref must have less than 40 elements")
	}

	localVarQueryParams.Add("ref", parameterToString(*r.ref, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
