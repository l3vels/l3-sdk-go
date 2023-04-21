# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the asset entity. | 
**TokenId** | **float32** | The token id of the asset. | 
**Name** | **string** | The name of the asset. | 
**ParentId** | **string** | ID of the parent asset. | 
**Properties** | **map[string]interface{}** | Custom properties of the asset. | 
**Attributes** | **map[string]interface{}** | Custom attributes of the asset. | 
**Description** | **string** | The description of the asset. | 
**Status** | **string** | The status of the asset. | 
**Price** | **float32** | The price of the asset. | 
**Supply** | **float32** | The supply of the asset. | 
**MintedAmount** | **float32** | The minted amount of the asset. | 
**AssetType** | **string** | The asset type. Can be main or nested. | 
**AssetUrl** | **string** | The asset URL. | 
**Medias** | **[]string** | Additional images associated with the asset, such as screenshots or promotional images. | 
**MainMedia** | **string** | The main or featured image associated with the asset. You can set it from the Dashboard as main image. | 
**AccountId** | **float32** | The unique identifier of the account that the Collection belongs to. | 
**ProjectId** | **string** | The unique identifier of the project that the asset is associated with. This allows developers to organize their assets by project and helps with tracking and reporting. | 
**CollectionId** | **string** | The unique identifier of the collection that the asset is associated with. This allows developers to organize their collections by project and helps with tracking and reporting. | 
**CreatedOn** | **time.Time** | The date when the collection was created. | 
**ModifiedOn** | **time.Time** | The date when the collection was last modified. | 
**CreatedBy** | **float32** | The Id of the user who created the collection. | 
**ModifiedBy** | **float32** | The Id of the user who last modified the collection. | 

## Methods

### NewAsset

`func NewAsset(id string, tokenId float32, name string, parentId string, properties map[string]interface{}, attributes map[string]interface{}, description string, status string, price float32, supply float32, mintedAmount float32, assetType string, assetUrl string, medias []string, mainMedia string, accountId float32, projectId string, collectionId string, createdOn time.Time, modifiedOn time.Time, createdBy float32, modifiedBy float32, ) *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.


### GetTokenId

`func (o *Asset) GetTokenId() float32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Asset) GetTokenIdOk() (*float32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Asset) SetTokenId(v float32)`

SetTokenId sets TokenId field to given value.


### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *Asset) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Asset) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Asset) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetProperties

`func (o *Asset) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Asset) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Asset) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.


### GetAttributes

`func (o *Asset) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Asset) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Asset) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *Asset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Asset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Asset) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPrice

`func (o *Asset) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Asset) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Asset) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetSupply

`func (o *Asset) GetSupply() float32`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *Asset) GetSupplyOk() (*float32, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *Asset) SetSupply(v float32)`

SetSupply sets Supply field to given value.


### GetMintedAmount

`func (o *Asset) GetMintedAmount() float32`

GetMintedAmount returns the MintedAmount field if non-nil, zero value otherwise.

### GetMintedAmountOk

`func (o *Asset) GetMintedAmountOk() (*float32, bool)`

GetMintedAmountOk returns a tuple with the MintedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintedAmount

`func (o *Asset) SetMintedAmount(v float32)`

SetMintedAmount sets MintedAmount field to given value.


### GetAssetType

`func (o *Asset) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Asset) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Asset) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetAssetUrl

`func (o *Asset) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *Asset) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *Asset) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.


### GetMedias

`func (o *Asset) GetMedias() []string`

GetMedias returns the Medias field if non-nil, zero value otherwise.

### GetMediasOk

`func (o *Asset) GetMediasOk() (*[]string, bool)`

GetMediasOk returns a tuple with the Medias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedias

`func (o *Asset) SetMedias(v []string)`

SetMedias sets Medias field to given value.


### GetMainMedia

`func (o *Asset) GetMainMedia() string`

GetMainMedia returns the MainMedia field if non-nil, zero value otherwise.

### GetMainMediaOk

`func (o *Asset) GetMainMediaOk() (*string, bool)`

GetMainMediaOk returns a tuple with the MainMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainMedia

`func (o *Asset) SetMainMedia(v string)`

SetMainMedia sets MainMedia field to given value.


### GetAccountId

`func (o *Asset) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Asset) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Asset) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.


### GetProjectId

`func (o *Asset) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Asset) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Asset) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCollectionId

`func (o *Asset) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Asset) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Asset) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetCreatedOn

`func (o *Asset) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Asset) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Asset) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *Asset) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Asset) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Asset) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.


### GetCreatedBy

`func (o *Asset) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Asset) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Asset) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *Asset) GetModifiedBy() float32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Asset) GetModifiedByOk() (*float32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Asset) SetModifiedBy(v float32)`

SetModifiedBy sets ModifiedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


