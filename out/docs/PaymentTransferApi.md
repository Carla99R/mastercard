# \PaymentTransferApi

All URIs are relative to *https://sandbox.api.mastercard.com/send*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayment**](PaymentTransferApi.md#CreatePayment) | **Post** /v1/partners/{partnerId}/transfers/payment | Create a payment transfer (Payment Transaction) to send funds to a recipient
[**GetTransferById**](PaymentTransferApi.md#GetTransferById) | **Get** /v1/partners/{partnerId}/transfers/{transferId} | Retrieve details of a payment transfer, including latest status, by using the transfer ID that was returned when you created the transfer
[**GetTransferByRef**](PaymentTransferApi.md#GetTransferByRef) | **Get** /v1/partners/{partnerId}/transfers | Retrieve details of a payment transfer, including latest status, by using the transfer_reference value you provided when you created the transfer



## CreatePayment

> TransferWrapper CreatePayment(ctx, partnerId).PaymentTransferWrapper(paymentTransferWrapper).Execute()

Create a payment transfer (Payment Transaction) to send funds to a recipient



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentTransferWrapper := *openapiclient.NewPaymentTransferWrapper() // PaymentTransferWrapper | Contains the details of the request message.
    partnerId := "ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo" // string | The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentTransferApi.CreatePayment(context.Background(), partnerId).PaymentTransferWrapper(paymentTransferWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentTransferApi.CreatePayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayment`: TransferWrapper
    fmt.Fprintf(os.Stdout, "Response from `PaymentTransferApi.CreatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentTransferWrapper** | [**PaymentTransferWrapper**](PaymentTransferWrapper.md) | Contains the details of the request message. | 


### Return type

[**TransferWrapper**](TransferWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferById

> TransferWrapperGet GetTransferById(ctx, partnerId, transferId).Execute()

Retrieve details of a payment transfer, including latest status, by using the transfer ID that was returned when you created the transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    partnerId := "ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo" // string | The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32.
    transferId := "trn_tN3S2cKGNt4hSmPwknYaiJd2BJrM" // string | Required. The system-generated Transfer ID (`id`) that was returned when you created the payment transfer. Alphanumeric and * , - . _ ~.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentTransferApi.GetTransferById(context.Background(), partnerId, transferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentTransferApi.GetTransferById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransferById`: TransferWrapperGet
    fmt.Fprintf(os.Stdout, "Response from `PaymentTransferApi.GetTransferById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32. | 
**transferId** | **string** | Required. The system-generated Transfer ID (&#x60;id&#x60;) that was returned when you created the payment transfer. Alphanumeric and * , - . _ ~. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransferWrapperGet**](TransferWrapperGet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransferByRef

> TransfersWrapper GetTransferByRef(ctx, partnerId).Ref(ref).Execute()

Retrieve details of a payment transfer, including latest status, by using the transfer_reference value you provided when you created the transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    partnerId := "ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo" // string | The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32.
    ref := "DEF123456" // string | Query parameter - Required. The unique `transfer_reference` that you provided when creating the payment transfer. Alphanumeric and * , - . _ ~. Length 6-40.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentTransferApi.GetTransferByRef(context.Background(), partnerId).Ref(ref).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentTransferApi.GetTransferByRef``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransferByRef`: TransfersWrapper
    fmt.Fprintf(os.Stdout, "Response from `PaymentTransferApi.GetTransferByRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** | The Mastercard-assigned unique ID for the partner registered to use Mastercard Send. String, length 32. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferByRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Query parameter - Required. The unique &#x60;transfer_reference&#x60; that you provided when creating the payment transfer. Alphanumeric and * , - . _ ~. Length 6-40. | 

### Return type

[**TransfersWrapper**](TransfersWrapper.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

