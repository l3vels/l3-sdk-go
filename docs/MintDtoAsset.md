# MintDtoAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **float32** | Token ID to mint. You can provide this or asset_id | [optional] 
**AssetId** | Pointer to **string** | Asset ID to mint. You can provide this or token_id | [optional] 
**Amount** | **float32** | Amount to mint | 

## Methods

### NewMintDtoAsset

`func NewMintDtoAsset(amount float32, ) *MintDtoAsset`

NewMintDtoAsset instantiates a new MintDtoAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintDtoAssetWithDefaults

`func NewMintDtoAssetWithDefaults() *MintDtoAsset`

NewMintDtoAssetWithDefaults instantiates a new MintDtoAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *MintDtoAsset) GetTokenId() float32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *MintDtoAsset) GetTokenIdOk() (*float32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *MintDtoAsset) SetTokenId(v float32)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *MintDtoAsset) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAssetId

`func (o *MintDtoAsset) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *MintDtoAsset) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *MintDtoAsset) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *MintDtoAsset) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAmount

`func (o *MintDtoAsset) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MintDtoAsset) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MintDtoAsset) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


