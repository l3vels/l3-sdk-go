# MintDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | Pointer to **string** | Contract ID | [optional] 
**ProjectId** | **string** | Game/project ID. Example: Call of Duty | 
**CollectionId** | **string** | Collection ID to use. Example: Characters, Weapons, etc. | 
**PlayerAddress** | Pointer to **string** | Player address to mint token to. You can provide player ID or player address | [optional] 
**PlayerId** | Pointer to **string** | Player ID to mint to. You can provide player ID or player address | [optional] 
**Asset** | [**MintDtoAsset**](MintDtoAsset.md) |  | 

## Methods

### NewMintDto

`func NewMintDto(projectId string, collectionId string, asset MintDtoAsset, ) *MintDto`

NewMintDto instantiates a new MintDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintDtoWithDefaults

`func NewMintDtoWithDefaults() *MintDto`

NewMintDtoWithDefaults instantiates a new MintDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *MintDto) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *MintDto) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *MintDto) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *MintDto) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetProjectId

`func (o *MintDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MintDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MintDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCollectionId

`func (o *MintDto) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *MintDto) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *MintDto) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetPlayerAddress

`func (o *MintDto) GetPlayerAddress() string`

GetPlayerAddress returns the PlayerAddress field if non-nil, zero value otherwise.

### GetPlayerAddressOk

`func (o *MintDto) GetPlayerAddressOk() (*string, bool)`

GetPlayerAddressOk returns a tuple with the PlayerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerAddress

`func (o *MintDto) SetPlayerAddress(v string)`

SetPlayerAddress sets PlayerAddress field to given value.

### HasPlayerAddress

`func (o *MintDto) HasPlayerAddress() bool`

HasPlayerAddress returns a boolean if a field has been set.

### GetPlayerId

`func (o *MintDto) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *MintDto) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *MintDto) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *MintDto) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetAsset

`func (o *MintDto) GetAsset() MintDtoAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *MintDto) GetAssetOk() (*MintDtoAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *MintDto) SetAsset(v MintDtoAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


