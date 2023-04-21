# client\MintApi

All URIs are relative to *https://api-dev.l3vels.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Airdrop**](MintApi.md#Airdrop) | **Post** /v1/mint/airdrop | Airdrop asset to player
[**Award**](MintApi.md#Award) | **Post** /v1/mint/award | Award asset to player
[**Mint**](MintApi.md#Mint) | **Post** /v1/mint | Mint asset
[**MintBatch**](MintApi.md#MintBatch) | **Post** /v1/mint/batch | Batch mint assets
[**PlayerMint**](MintApi.md#PlayerMint) | **Post** /v1/mint/player | Mint asset by player
[**PlayerMintBatch**](MintApi.md#PlayerMintBatch) | **Post** /v1/mint/batch-player | Batch mint assets by player



## Airdrop

> Airdrop(ctx).Authorization(authorization).MintDto(mintDto).Execute()

Airdrop asset to player



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
    mintDto := *openapiclient.NewMintDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", *openapiclient.NewMintDtoAsset(float32(1))) // MintDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MintApi.Airdrop(context.Background()).Authorization(authorization).MintDto(mintDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintApi.Airdrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAirdropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **mintDto** | [**MintDto**](MintDto.md) |  | 

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


## Award

> Award(ctx).Authorization(authorization).MintDto(mintDto).Execute()

Award asset to player



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
    mintDto := *openapiclient.NewMintDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", *openapiclient.NewMintDtoAsset(float32(1))) // MintDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MintApi.Award(context.Background()).Authorization(authorization).MintDto(mintDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintApi.Award``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **mintDto** | [**MintDto**](MintDto.md) |  | 

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


## Mint

> Mint(ctx).Authorization(authorization).MintDto(mintDto).Execute()

Mint asset



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
    mintDto := *openapiclient.NewMintDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", *openapiclient.NewMintDtoAsset(float32(1))) // MintDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MintApi.Mint(context.Background()).Authorization(authorization).MintDto(mintDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintApi.Mint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **mintDto** | [**MintDto**](MintDto.md) |  | 

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


## MintBatch

> MintBatch(ctx).Authorization(authorization).MintBatchDto(mintBatchDto).Execute()

Batch mint assets



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
    mintBatchDto := *openapiclient.NewMintBatchDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", "0x0000000000000000000000000000000000000000", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", []string{"Assets_example"}) // MintBatchDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MintApi.MintBatch(context.Background()).Authorization(authorization).MintBatchDto(mintBatchDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintApi.MintBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMintBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **mintBatchDto** | [**MintBatchDto**](MintBatchDto.md) |  | 

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


## PlayerMint

> PlayerMint(ctx).Authorization(authorization).MintDto(mintDto).Execute()

Mint asset by player



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
    mintDto := *openapiclient.NewMintDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", *openapiclient.NewMintDtoAsset(float32(1))) // MintDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MintApi.PlayerMint(context.Background()).Authorization(authorization).MintDto(mintDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintApi.PlayerMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayerMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **mintDto** | [**MintDto**](MintDto.md) |  | 

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


## PlayerMintBatch

> PlayerMintBatch(ctx).Authorization(authorization).MintBatchDto(mintBatchDto).Execute()

Batch mint assets by player



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
    mintBatchDto := *openapiclient.NewMintBatchDto("a44b646a-ae14-4e05-ae09-b12d5e7269bf", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", "0x0000000000000000000000000000000000000000", "a44b646a-ae14-4e05-ae09-b12d5e7269bf", []string{"Assets_example"}) // MintBatchDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MintApi.PlayerMintBatch(context.Background()).Authorization(authorization).MintBatchDto(mintBatchDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintApi.PlayerMintBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayerMintBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **mintBatchDto** | [**MintBatchDto**](MintBatchDto.md) |  | 

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

