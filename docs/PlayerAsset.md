# PlayerAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the Player asset entity. | 
**Amount** | **float32** | The amount of the specific asset that the player has. | 
**PlayerId** | **string** | The unique identifier of the player that the asset is associated with. | 
**AssetId** | **string** | The unique identifier of the asset that the asset is associated with. | 
**CollectionId** | **string** | The unique identifier of the collection that the Player asset is associated with. | 
**AccountId** | **float32** | The unique identifier of the account that the Player belongs to. | 
**ProjectId** | **string** | The unique identifier of the project that the Player is associated with. This allows developers to organize their players by project and helps with tracking and reporting. | 
**CreatedOn** | **time.Time** | The date when the player was created. | 
**ModifiedOn** | **time.Time** | The date when the player was last modified. | 
**CreatedBy** | **float32** | The Id of the user who created the player. | 
**ModifiedBy** | **float32** | The Id of the user who last modified the player. | 

## Methods

### NewPlayerAsset

`func NewPlayerAsset(id string, amount float32, playerId string, assetId string, collectionId string, accountId float32, projectId string, createdOn time.Time, modifiedOn time.Time, createdBy float32, modifiedBy float32, ) *PlayerAsset`

NewPlayerAsset instantiates a new PlayerAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerAssetWithDefaults

`func NewPlayerAssetWithDefaults() *PlayerAsset`

NewPlayerAssetWithDefaults instantiates a new PlayerAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlayerAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayerAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayerAsset) SetId(v string)`

SetId sets Id field to given value.


### GetAmount

`func (o *PlayerAsset) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PlayerAsset) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PlayerAsset) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetPlayerId

`func (o *PlayerAsset) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *PlayerAsset) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *PlayerAsset) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.


### GetAssetId

`func (o *PlayerAsset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *PlayerAsset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *PlayerAsset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetCollectionId

`func (o *PlayerAsset) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *PlayerAsset) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *PlayerAsset) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetAccountId

`func (o *PlayerAsset) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlayerAsset) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlayerAsset) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.


### GetProjectId

`func (o *PlayerAsset) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PlayerAsset) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PlayerAsset) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedOn

`func (o *PlayerAsset) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *PlayerAsset) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *PlayerAsset) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *PlayerAsset) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *PlayerAsset) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *PlayerAsset) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.


### GetCreatedBy

`func (o *PlayerAsset) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PlayerAsset) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PlayerAsset) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *PlayerAsset) GetModifiedBy() float32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PlayerAsset) GetModifiedByOk() (*float32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PlayerAsset) SetModifiedBy(v float32)`

SetModifiedBy sets ModifiedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


