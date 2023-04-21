# client\PlayerApi

All URIs are relative to *https://api-dev.l3vels.xyz*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlayer**](PlayerApi.md#CreatePlayer) | **Post** /v1/player | Create new player
[**DeletePlayer**](PlayerApi.md#DeletePlayer) | **Delete** /v1/player | Delete a Player
[**GetPlayers**](PlayerApi.md#GetPlayers) | **Get** /v1/player | Retrieve players
[**PlayerAssetById**](PlayerApi.md#PlayerAssetById) | **Get** /v1/player-asset/{project_id}/{id} | Retrieve player asset by ID
[**PlayerAssets**](PlayerApi.md#PlayerAssets) | **Get** /v1/player-asset | Retrieve player assets
[**PlayerById**](PlayerApi.md#PlayerById) | **Get** /v1/player/{project_id}/{id} | Retrieve player by ID
[**PlayersCountByGameId**](PlayerApi.md#PlayersCountByGameId) | **Get** /v1/player/count/{project_id} | Count players
[**UpdatePlayer**](PlayerApi.md#UpdatePlayer) | **Put** /v1/player | Update an existing Player



## CreatePlayer

> Player CreatePlayer(ctx).Authorization(authorization).CreatePlayerDto(createPlayerDto).Execute()

Create new player



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
    createPlayerDto := *openapiclient.NewCreatePlayerDto("Jack", "353c69f6-76a6-4baa-b68b-852c1c531953") // CreatePlayerDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.CreatePlayer(context.Background()).Authorization(authorization).CreatePlayerDto(createPlayerDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.CreatePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlayer`: Player
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.CreatePlayer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **createPlayerDto** | [**CreatePlayerDto**](CreatePlayerDto.md) |  | 

### Return type

[**Player**](Player.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlayer

> DeletePlayer(ctx).Authorization(authorization).Execute()

Delete a Player



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
    r, err := apiClient.PlayerApi.DeletePlayer(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.DeletePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlayerRequest struct via the builder pattern


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


## GetPlayers

> []Player GetPlayers(ctx).Authorization(authorization).ProjectId(projectId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()

Retrieve players



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
    projectId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find player in your game. Example: Fortnite, Minecraft, etc.
    sort := "name" // string | Player field to sort by. You can sort by name, created_on and etc. (optional)
    order := "ASC" // string | Sort order (ASC for ascending or DESC for descending) (optional)
    searchText := "Jack" // string | Search player by name (optional)
    limit := float32(10) // float32 | Number of players to return per page (optional)
    page := float32(1) // float32 | Page number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.GetPlayers(context.Background()).Authorization(authorization).ProjectId(projectId).Sort(sort).Order(order).SearchText(searchText).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.GetPlayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayers`: []Player
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.GetPlayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **projectId** | **string** | Game/project ID to find player in your game. Example: Fortnite, Minecraft, etc. | 
 **sort** | **string** | Player field to sort by. You can sort by name, created_on and etc. | 
 **order** | **string** | Sort order (ASC for ascending or DESC for descending) | 
 **searchText** | **string** | Search player by name | 
 **limit** | **float32** | Number of players to return per page | 
 **page** | **float32** | Page number | 

### Return type

[**[]Player**](Player.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayerAssetById

> PlayerAsset PlayerAssetById(ctx, id, projectId).Authorization(authorization).Execute()

Retrieve player asset by ID



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
    resp, r, err := apiClient.PlayerApi.PlayerAssetById(context.Background(), id, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.PlayerAssetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayerAssetById`: PlayerAsset
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.PlayerAssetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerAssetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 



### Return type

[**PlayerAsset**](PlayerAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayerAssets

> []PlayerAsset PlayerAssets(ctx).Authorization(authorization).ProjectId(projectId).AssetId(assetId).PlayerId(playerId).Sort(sort).Order(order).Limit(limit).Page(page).Execute()

Retrieve player assets



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
    projectId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc.
    assetId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc. (optional)
    playerId := "a44b646a-ae14-4e05-ae09-b12d5e7269bf" // string | Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc. (optional)
    sort := "name" // string | Player asset field to sort by. You can sort by name, created_on and etc. (optional)
    order := "ASC" // string | Sort order (ASC for ascending or DESC for descending) (optional)
    limit := float32(10) // float32 | Number of player assets to return per page (optional)
    page := float32(1) // float32 | Page number (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.PlayerAssets(context.Background()).Authorization(authorization).ProjectId(projectId).AssetId(assetId).PlayerId(playerId).Sort(sort).Order(order).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.PlayerAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayerAssets`: []PlayerAsset
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.PlayerAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayerAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 
 **projectId** | **string** | Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc. | 
 **assetId** | **string** | Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc. | 
 **playerId** | **string** | Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc. | 
 **sort** | **string** | Player asset field to sort by. You can sort by name, created_on and etc. | 
 **order** | **string** | Sort order (ASC for ascending or DESC for descending) | 
 **limit** | **float32** | Number of player assets to return per page | 
 **page** | **float32** | Page number | 

### Return type

[**[]PlayerAsset**](PlayerAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayerById

> Player PlayerById(ctx, id, projectId).Authorization(authorization).Execute()

Retrieve player by ID



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
    id := "9abd57ce-b11c-4828-ab6a-19f568a8081a" // string | Player ID that you created in your game/project. Example: Jack, George, etc.
    projectId := "556a2843-b302-4b9d-916c-cefcb5d66053" // string | Game/project ID to find asset in. Example: Call of Duty, Fortnite, etc.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.PlayerById(context.Background(), id, projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.PlayerById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayerById`: Player
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.PlayerById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Player ID that you created in your game/project. Example: Jack, George, etc. | 
**projectId** | **string** | Game/project ID to find asset in. Example: Call of Duty, Fortnite, etc. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 



### Return type

[**Player**](Player.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayersCountByGameId

> float32 PlayersCountByGameId(ctx, projectId).Authorization(authorization).Execute()

Count players



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
    projectId := "1" // string | Game Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.PlayersCountByGameId(context.Background(), projectId).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.PlayersCountByGameId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayersCountByGameId`: float32
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.PlayersCountByGameId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | Game Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayersCountByGameIdRequest struct via the builder pattern


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


## UpdatePlayer

> Player UpdatePlayer(ctx).Authorization(authorization).Execute()

Update an existing Player



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
    resp, r, err := apiClient.PlayerApi.UpdatePlayer(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.UpdatePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlayer`: Player
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.UpdatePlayer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | API key is associated with multiple projects. Please include it in to use developers API. | 

### Return type

[**Player**](Player.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

