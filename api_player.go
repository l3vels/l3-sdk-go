/*
L3vels Api

L3vels API for Game developers

API version: 0.3
Contact: support@l3vels.xyz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package l3vels-sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PlayerApiService PlayerApi service
type PlayerApiService service

type ApiCreatePlayerRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
	createPlayerDto *CreatePlayerDto
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiCreatePlayerRequest) Authorization(authorization string) ApiCreatePlayerRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreatePlayerRequest) CreatePlayerDto(createPlayerDto CreatePlayerDto) ApiCreatePlayerRequest {
	r.createPlayerDto = &createPlayerDto
	return r
}

func (r ApiCreatePlayerRequest) Execute() (*Player, *http.Response, error) {
	return r.ApiService.CreatePlayerExecute(r)
}

/*
CreatePlayer Create new player

Create new player for game/project. Example: Create new player Jack in game Call of Duty.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePlayerRequest
*/
func (a *PlayerApiService) CreatePlayer(ctx context.Context) ApiCreatePlayerRequest {
	return ApiCreatePlayerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Player
func (a *PlayerApiService) CreatePlayerExecute(r ApiCreatePlayerRequest) (*Player, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Player
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.CreatePlayer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.createPlayerDto == nil {
		return localVarReturnValue, nil, reportError("createPlayerDto is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	// body params
	localVarPostBody = r.createPlayerDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeletePlayerRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiDeletePlayerRequest) Authorization(authorization string) ApiDeletePlayerRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeletePlayerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePlayerExecute(r)
}

/*
DeletePlayer Delete a Player

This API method allows developers to delete a Player by providing the ID of the Player. Once deleted, the Player and all associated assets will be removed from the system.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeletePlayerRequest
*/
func (a *PlayerApiService) DeletePlayer(ctx context.Context) ApiDeletePlayerRequest {
	return ApiDeletePlayerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PlayerApiService) DeletePlayerExecute(r ApiDeletePlayerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.DeletePlayer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPlayersRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
	projectId *string
	sort *string
	order *string
	searchText *string
	limit *float32
	page *float32
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiGetPlayersRequest) Authorization(authorization string) ApiGetPlayersRequest {
	r.authorization = &authorization
	return r
}

// Game/project ID to find player in your game. Example: Fortnite, Minecraft, etc.
func (r ApiGetPlayersRequest) ProjectId(projectId string) ApiGetPlayersRequest {
	r.projectId = &projectId
	return r
}

// Player field to sort by. You can sort by name, created_on and etc.
func (r ApiGetPlayersRequest) Sort(sort string) ApiGetPlayersRequest {
	r.sort = &sort
	return r
}

// Sort order (ASC for ascending or DESC for descending)
func (r ApiGetPlayersRequest) Order(order string) ApiGetPlayersRequest {
	r.order = &order
	return r
}

// Search player by name
func (r ApiGetPlayersRequest) SearchText(searchText string) ApiGetPlayersRequest {
	r.searchText = &searchText
	return r
}

// Number of players to return per page
func (r ApiGetPlayersRequest) Limit(limit float32) ApiGetPlayersRequest {
	r.limit = &limit
	return r
}

// Page number
func (r ApiGetPlayersRequest) Page(page float32) ApiGetPlayersRequest {
	r.page = &page
	return r
}

func (r ApiGetPlayersRequest) Execute() ([]Player, *http.Response, error) {
	return r.ApiService.GetPlayersExecute(r)
}

/*
GetPlayers Retrieve players

Retrieve a list of players that match the specified filter criteria. Developers can use this method to retrieve players by name, category, status, or other properties. Example: Retrieve players from game Call of Duty.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPlayersRequest
*/
func (a *PlayerApiService) GetPlayers(ctx context.Context) ApiGetPlayersRequest {
	return ApiGetPlayersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Player
func (a *PlayerApiService) GetPlayersExecute(r ApiGetPlayersRequest) ([]Player, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Player
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.GetPlayers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.projectId == nil {
		return localVarReturnValue, nil, reportError("projectId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId, "")
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.searchText != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_text", r.searchText, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPlayerAssetByIdRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
	id string
	projectId string
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiPlayerAssetByIdRequest) Authorization(authorization string) ApiPlayerAssetByIdRequest {
	r.authorization = &authorization
	return r
}

func (r ApiPlayerAssetByIdRequest) Execute() (*PlayerAsset, *http.Response, error) {
	return r.ApiService.PlayerAssetByIdExecute(r)
}

/*
PlayerAssetById Retrieve player asset by ID

Retrieve player asset by ID. Player asset represents a single asset that a player owns. It has amount field that represents how many of this asset player owns.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param projectId
 @return ApiPlayerAssetByIdRequest
*/
func (a *PlayerApiService) PlayerAssetById(ctx context.Context, id string, projectId string) ApiPlayerAssetByIdRequest {
	return ApiPlayerAssetByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return PlayerAsset
func (a *PlayerApiService) PlayerAssetByIdExecute(r ApiPlayerAssetByIdRequest) (*PlayerAsset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PlayerAsset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.PlayerAssetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player-asset/{project_id}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPlayerAssetsRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
	projectId *string
	assetId *string
	playerId *string
	sort *string
	order *string
	limit *float32
	page *float32
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiPlayerAssetsRequest) Authorization(authorization string) ApiPlayerAssetsRequest {
	r.authorization = &authorization
	return r
}

// Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc.
func (r ApiPlayerAssetsRequest) ProjectId(projectId string) ApiPlayerAssetsRequest {
	r.projectId = &projectId
	return r
}

// Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc.
func (r ApiPlayerAssetsRequest) AssetId(assetId string) ApiPlayerAssetsRequest {
	r.assetId = &assetId
	return r
}

// Game/project ID to find player assets in your game. Example: Fortnite, Minecraft, etc.
func (r ApiPlayerAssetsRequest) PlayerId(playerId string) ApiPlayerAssetsRequest {
	r.playerId = &playerId
	return r
}

// Player asset field to sort by. You can sort by name, created_on and etc.
func (r ApiPlayerAssetsRequest) Sort(sort string) ApiPlayerAssetsRequest {
	r.sort = &sort
	return r
}

// Sort order (ASC for ascending or DESC for descending)
func (r ApiPlayerAssetsRequest) Order(order string) ApiPlayerAssetsRequest {
	r.order = &order
	return r
}

// Number of player assets to return per page
func (r ApiPlayerAssetsRequest) Limit(limit float32) ApiPlayerAssetsRequest {
	r.limit = &limit
	return r
}

// Page number
func (r ApiPlayerAssetsRequest) Page(page float32) ApiPlayerAssetsRequest {
	r.page = &page
	return r
}

func (r ApiPlayerAssetsRequest) Execute() ([]PlayerAsset, *http.Response, error) {
	return r.ApiService.PlayerAssetsExecute(r)
}

/*
PlayerAssets Retrieve player assets

This API method retrieves a list of Player assets that match the specified filter criteria. Developers can use this method to retrieve Player assets by player, game/project or other properties.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPlayerAssetsRequest
*/
func (a *PlayerApiService) PlayerAssets(ctx context.Context) ApiPlayerAssetsRequest {
	return ApiPlayerAssetsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PlayerAsset
func (a *PlayerApiService) PlayerAssetsExecute(r ApiPlayerAssetsRequest) ([]PlayerAsset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PlayerAsset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.PlayerAssets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player-asset"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.projectId == nil {
		return localVarReturnValue, nil, reportError("projectId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "project_id", r.projectId, "")
	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId, "")
	}
	if r.playerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "player_id", r.playerId, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPlayerByIdRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
	id string
	projectId string
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiPlayerByIdRequest) Authorization(authorization string) ApiPlayerByIdRequest {
	r.authorization = &authorization
	return r
}

func (r ApiPlayerByIdRequest) Execute() (*Player, *http.Response, error) {
	return r.ApiService.PlayerByIdExecute(r)
}

/*
PlayerById Retrieve player by ID

Retrieves a specific player by ID associated with game/project. Example: retrieve player Jack from game Call of Duty.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Player ID that you created in your game/project. Example: Jack, George, etc.
 @param projectId Game/project ID to find asset in. Example: Call of Duty, Fortnite, etc.
 @return ApiPlayerByIdRequest
*/
func (a *PlayerApiService) PlayerById(ctx context.Context, id string, projectId string) ApiPlayerByIdRequest {
	return ApiPlayerByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return Player
func (a *PlayerApiService) PlayerByIdExecute(r ApiPlayerByIdRequest) (*Player, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Player
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.PlayerById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player/{project_id}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiPlayersCountByGameIdRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
	projectId string
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiPlayersCountByGameIdRequest) Authorization(authorization string) ApiPlayersCountByGameIdRequest {
	r.authorization = &authorization
	return r
}

func (r ApiPlayersCountByGameIdRequest) Execute() (float32, *http.Response, error) {
	return r.ApiService.PlayersCountByGameIdExecute(r)
}

/*
PlayersCountByGameId Count players

Count players in game. Example: count players in game Call of Duty.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Game Id
 @return ApiPlayersCountByGameIdRequest
*/
func (a *PlayerApiService) PlayersCountByGameId(ctx context.Context, projectId string) ApiPlayersCountByGameIdRequest {
	return ApiPlayersCountByGameIdRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return float32
func (a *PlayerApiService) PlayersCountByGameIdExecute(r ApiPlayersCountByGameIdRequest) (float32, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  float32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.PlayersCountByGameId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player/count/{project_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdatePlayerRequest struct {
	ctx context.Context
	ApiService *PlayerApiService
	authorization *string
}

// API key is associated with multiple projects. Please include it in to use developers API.
func (r ApiUpdatePlayerRequest) Authorization(authorization string) ApiUpdatePlayerRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdatePlayerRequest) Execute() (*Player, *http.Response, error) {
	return r.ApiService.UpdatePlayerExecute(r)
}

/*
UpdatePlayer Update an existing Player

This API method allows developers to update an existing Player by providing the ID of the Player and the updated properties and associated assets.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdatePlayerRequest
*/
func (a *PlayerApiService) UpdatePlayer(ctx context.Context) ApiUpdatePlayerRequest {
	return ApiUpdatePlayerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Player
func (a *PlayerApiService) UpdatePlayerExecute(r ApiUpdatePlayerRequest) (*Player, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Player
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PlayerApiService.UpdatePlayer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/player"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
