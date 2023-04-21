# client\TransactionApi

All URIs are relative to *https://api-dev.l3vels.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionById**](TransactionApi.md#TransactionById) | **Get** /v1/transaction/{project_id}/{id} | Retrieve Transaction by ID
[**Transactions**](TransactionApi.md#Transactions) | **Get** /v1/transaction | Retrieve transactions
[**Webhook**](TransactionApi.md#Webhook) | **Post** /v1/transaction/webhook | 



## TransactionById

> Transaction TransactionById(ctx, id, projectId).Authorization(authorization).Execute()

Retrieve Transaction by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chkhikvadze/l3vels-sdk"
)

func main() {
    authorization := "authorization_example" // string | API key is associated with multiple projects. Please include it in to use developers API.
    id := "id_example" // string | 
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionApi.TransactionById(context.Background(), id, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.TransactionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionById`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.TransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 



### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Transactions

> Transaction Transactions(ctx).Authorization(authorization).ProjectId(projectId).CollectionId(collectionId).PlayerId(playerId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()

Retrieve transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chkhikvadze/l3vels-sdk"
)

func main() {
    authorization := "authorization_example" // string | API key is associated with multiple projects. Please include it in to use developers API.
    projectId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find transactions in your game. Example: Fortnite, Minecraft, etc.
    collectionId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Filter transactions by collection. Example: Get transactions only from Weapons collection. (optional)
    playerId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Player ID to mint to. You can provide player ID or player address (optional)
    sort := "name" // string | Asset field to sort by. You can sort by name, created_on and etc. (optional)
    order := "ASC" // string | Sort order (ASC for ascending or DESC for descending) (optional)
    searchText := "Hammer" // string | Search transactions by name (optional)
    limit := float32(10) // float32 | Number of transactions to return per page (optional)
    page := float32(1) // float32 | Page number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionApi.Transactions(context.Background()).Authorization(authorization).ProjectId(projectId).CollectionId(collectionId).PlayerId(playerId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.Transactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Transactions`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.Transactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **projectId** | **string** | Game/project ID to find transactions in your game. Example: Fortnite, Minecraft, etc. | 
 **collectionId** | **string** | Filter transactions by collection. Example: Get transactions only from Weapons collection. | 
 **playerId** | **string** | Player ID to mint to. You can provide player ID or player address | 
 **sort** | **string** | Asset field to sort by. You can sort by name, created_on and etc. | 
 **order** | **string** | Sort order (ASC for ascending or DESC for descending) | 
 **searchText** | **string** | Search transactions by name | 
 **limit** | **float32** | Number of transactions to return per page | 
 **page** | **float32** | Page number | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Webhook

> Webhook(ctx).Authorization(authorization).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chkhikvadze/l3vels-sdk"
)

func main() {
    authorization := "authorization_example" // string | API key is associated with multiple projects. Please include it in to use developers API.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransactionApi.Webhook(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.Webhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

