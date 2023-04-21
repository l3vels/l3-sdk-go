# MintBatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Project Id | 
**CollectionId** | **string** | Collection Id | 
**PlayerAddress** | **string** | Player address to mint token to | 
**PlayerId** | **string** | Player ID to mint to. You can provide this or player_address | 
**Assets** | **[]string** | Array of tokens to mint | 

## Methods

### NewMintBatchDto

`func NewMintBatchDto(projectId string, collectionId string, playerAddress string, playerId string, assets []string, ) *MintBatchDto`

NewMintBatchDto instantiates a new MintBatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintBatchDtoWithDefaults

`func NewMintBatchDtoWithDefaults() *MintBatchDto`

NewMintBatchDtoWithDefaults instantiates a new MintBatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *MintBatchDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MintBatchDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MintBatchDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCollectionId

`func (o *MintBatchDto) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *MintBatchDto) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *MintBatchDto) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetPlayerAddress

`func (o *MintBatchDto) GetPlayerAddress() string`

GetPlayerAddress returns the PlayerAddress field if non-nil, zero value otherwise.

### GetPlayerAddressOk

`func (o *MintBatchDto) GetPlayerAddressOk() (*string, bool)`

GetPlayerAddressOk returns a tuple with the PlayerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerAddress

`func (o *MintBatchDto) SetPlayerAddress(v string)`

SetPlayerAddress sets PlayerAddress field to given value.


### GetPlayerId

`func (o *MintBatchDto) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *MintBatchDto) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *MintBatchDto) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.


### GetAssets

`func (o *MintBatchDto) GetAssets() []string`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *MintBatchDto) GetAssetsOk() (*[]string, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *MintBatchDto) SetAssets(v []string)`

SetAssets sets Assets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


