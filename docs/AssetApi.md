# client\AssetApi

All URIs are relative to *https://api-dev.l3vels.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountByGame**](AssetApi.md#CountByGame) | **Get** /v1/asset/count/{project_id} | Count assets
[**GetAssetById**](AssetApi.md#GetAssetById) | **Get** /v1/asset/{project_id}/{id} | Retrieve asset by ID
[**GetAssets**](AssetApi.md#GetAssets) | **Get** /v1/asset | Retrieve assets
[**UpdateAsset**](AssetApi.md#UpdateAsset) | **Patch** /v1/asset/{id} | Update asset



## CountByGame

> float32 CountByGame(ctx, projectId).Authorization(authorization).Execute()

Count assets



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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetApi.CountByGame(context.Background(), projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.CountByGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountByGame`: float32
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.CountByGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountByGameRequest struct via the builder pattern


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


## GetAssetById

> Asset GetAssetById(ctx, id, projectId).Authorization(authorization).Execute()

Retrieve asset by ID



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
    id := "9abd57ce-b11c-4828-ab6a-19f568a8081a" // string | Asset ID to find
    projectId := "556a2843-b302-4b9d-916c-cefcb5d66053" // string | Game/project ID to find asset in

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetApi.GetAssetById(context.Background(), id, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetById`: Asset
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Asset ID to find | 
**projectId** | **string** | Game/project ID to find asset in | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 



### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets

> AssetsResponseDto GetAssets(ctx).Authorization(authorization).ProjectId(projectId).CollectionId(collectionId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()

Retrieve assets



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
    projectId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find assets in your game. Example: Fortnite, Minecraft, etc.
    collectionId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Filter assets by collection. Example: Get assets only from Weapons collection. (optional)
    sort := "name" // string | Asset field to sort by. You can sort by name, created_on and etc. (optional)
    order := "ASC" // string | Sort order (ASC for ascending or DESC for descending) (optional)
    searchText := "Hammer" // string | Search assets by name (optional)
    limit := float32(10) // float32 | Number of assets to return per page (optional)
    page := float32(1) // float32 | Page number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetApi.GetAssets(context.Background()).Authorization(authorization).ProjectId(projectId).CollectionId(collectionId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.GetAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssets`: AssetsResponseDto
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.GetAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **projectId** | **string** | Game/project ID to find assets in your game. Example: Fortnite, Minecraft, etc. | 
 **collectionId** | **string** | Filter assets by collection. Example: Get assets only from Weapons collection. | 
 **sort** | **string** | Asset field to sort by. You can sort by name, created_on and etc. | 
 **order** | **string** | Sort order (ASC for ascending or DESC for descending) | 
 **searchText** | **string** | Search assets by name | 
 **limit** | **float32** | Number of assets to return per page | 
 **page** | **float32** | Page number | 

### Return type

[**AssetsResponseDto**](AssetsResponseDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAsset

> Asset UpdateAsset(ctx, id).Authorization(authorization).UpdateAssetDto(updateAssetDto).Execute()

Update asset



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
    id := "353c69f6-76a6-4baa-b68b-852c1c531953" // string | Asset ID to update
    updateAssetDto := *openapiclient.NewUpdateAssetDto("353c69f6-76a6-4baa-b68b-852c1c531953", "353c69f6-76a6-4baa-b68b-852c1c531953") // UpdateAssetDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetApi.UpdateAsset(context.Background(), id).Authorization(authorization).UpdateAssetDto(updateAssetDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetApi.UpdateAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAsset`: Asset
    fmt.Fprintf(os.Stdout, "Response from `AssetApi.UpdateAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Asset ID to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 

 **updateAssetDto** | [**UpdateAssetDto**](UpdateAssetDto.md) |  | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

