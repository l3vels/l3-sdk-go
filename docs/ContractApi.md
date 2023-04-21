# client\ContractApi

All URIs are relative to *https://api-dev.l3vels.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionSize**](ContractApi.md#CollectionSize) | **Get** /v1/contract/collection-size | Collection size
[**ContractUri**](ContractApi.md#ContractUri) | **Get** /v1/contract/contract-uri | Get Contract URI
[**SetContractUri**](ContractApi.md#SetContractUri) | **Put** /v1/contract/contract-uri | Update Contract URI
[**SetSaleStatus**](ContractApi.md#SetSaleStatus) | **Put** /v1/contract/sale-status | Update Sale status



## CollectionSize

> CollectionSize(ctx).Authorization(authorization).CollectionId(collectionId).ProjectId(projectId).Execute()

Collection size



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
    collectionId := "collectionId_example" // string | 
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContractApi.CollectionSize(context.Background()).Authorization(authorization).CollectionId(collectionId).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.CollectionSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **collectionId** | **string** |  | 
 **projectId** | **string** |  | 

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


## ContractUri

> ContractUri(ctx).Authorization(authorization).CollectionId(collectionId).ProjectId(projectId).Execute()

Get Contract URI



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
    collectionId := "collectionId_example" // string | 
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContractApi.ContractUri(context.Background()).Authorization(authorization).CollectionId(collectionId).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.ContractUri``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **collectionId** | **string** |  | 
 **projectId** | **string** |  | 

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


## SetContractUri

> SetContractUri(ctx).Authorization(authorization).SetContractUriDto(setContractUriDto).Execute()

Update Contract URI



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
    setContractUriDto := *openapiclient.NewSetContractUriDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", "https://api.example.com/contract/{contractAddress}") // SetContractUriDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContractApi.SetContractUri(context.Background()).Authorization(authorization).SetContractUriDto(setContractUriDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.SetContractUri``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetContractUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **setContractUriDto** | [**SetContractUriDto**](SetContractUriDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSaleStatus

> SetSaleStatus(ctx).Authorization(authorization).SetSaleStatusDto(setSaleStatusDto).Execute()

Update Sale status



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
    setSaleStatusDto := *openapiclient.NewSetSaleStatusDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", map[string]interface{}(PUBLIC)) // SetSaleStatusDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ContractApi.SetSaleStatus(context.Background()).Authorization(authorization).SetSaleStatusDto(setSaleStatusDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractApi.SetSaleStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSaleStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **setSaleStatusDto** | [**SetSaleStatusDto**](SetSaleStatusDto.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

