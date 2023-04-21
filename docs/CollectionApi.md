# client\CollectionApi

All URIs are relative to *https://api-dev.l3vels.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountCollectionsByGameId**](CollectionApi.md#CountCollectionsByGameId) | **Get** /v1/collection/count/{project_id} | Count collections
[**GetCollectionById**](CollectionApi.md#GetCollectionById) | **Get** /v1/collection/{project_id}/{id} | Retrieve collection by ID
[**GetCollections**](CollectionApi.md#GetCollections) | **Get** /v1/collection | Retrieve collections



## CountCollectionsByGameId

> float32 CountCollectionsByGameId(ctx, projectId).Authorization(authorization).Execute()

Count collections



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
    projectId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to count collections in

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionApi.CountCollectionsByGameId(context.Background(), projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.CountCollectionsByGameId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCollectionsByGameId`: float32
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.CountCollectionsByGameId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Game/project ID to count collections in | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountCollectionsByGameIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 


### Return type

**float32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionById

> Collection GetCollectionById(ctx, id, projectId).Authorization(authorization).Execute()

Retrieve collection by ID



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
    id := "229fd9e0-b51f-4b20-9203-9db60995e6b1" // string | Collection ID to find
    projectId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find collection in

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionApi.GetCollectionById(context.Background(), id, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetCollectionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollectionById`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetCollectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Collection ID to find | 
**projectId** | **string** | Game/project ID to find collection in | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 



### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollections

> []Collection GetCollections(ctx).Authorization(authorization).ProjectId(projectId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()

Retrieve collections



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
    projectId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find collections in your game. Example: Fortnite, Minecraft, etc.
    sort := "name" // string | Collection field to sort by. You can sort by name, created_on and etc. (optional)
    order := "ASC" // string | Sort order (ASC for ascending or DESC for descending) (optional)
    searchText := "Weapons" // string | Search collections by name (optional)
    limit := float32(10) // float32 | Number of collections to return per page (optional)
    page := float32(1) // float32 | Page number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionApi.GetCollections(context.Background()).Authorization(authorization).ProjectId(projectId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionApi.GetCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollections`: []Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionApi.GetCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **projectId** | **string** | Game/project ID to find collections in your game. Example: Fortnite, Minecraft, etc. | 
 **sort** | **string** | Collection field to sort by. You can sort by name, created_on and etc. | 
 **order** | **string** | Sort order (ASC for ascending or DESC for descending) | 
 **searchText** | **string** | Search collections by name | 
 **limit** | **float32** | Number of collections to return per page | 
 **page** | **float32** | Page number | 

### Return type

[**[]Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

