# UpdateAssetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the asset. | [optional] 
**Description** | Pointer to **string** | The story of asset. | [optional] 
**Price** | Pointer to **float32** | Price of asset | [optional] 
**Supply** | Pointer to **float32** | Supply of asset | [optional] 
**AssetUrl** | Pointer to **string** | Asset URL | [optional] 
**CustomProps** | Pointer to **map[string]interface{}** | Custom props for asset. | [optional] 
**CollectionId** | **string** | Collection ID to find and update the asset in. Example: Update AK-47 asset in Weapons collection. | 
**ProjectId** | **string** | Game/project ID to update the asset in. Example: Call of Duty | 

## Methods

### NewUpdateAssetDto

`func NewUpdateAssetDto(collectionId string, projectId string, ) *UpdateAssetDto`

NewUpdateAssetDto instantiates a new UpdateAssetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAssetDtoWithDefaults

`func NewUpdateAssetDtoWithDefaults() *UpdateAssetDto`

NewUpdateAssetDtoWithDefaults instantiates a new UpdateAssetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAssetDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAssetDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAssetDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAssetDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateAssetDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAssetDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAssetDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAssetDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *UpdateAssetDto) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdateAssetDto) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdateAssetDto) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdateAssetDto) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSupply

`func (o *UpdateAssetDto) GetSupply() float32`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *UpdateAssetDto) GetSupplyOk() (*float32, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *UpdateAssetDto) SetSupply(v float32)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *UpdateAssetDto) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetAssetUrl

`func (o *UpdateAssetDto) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *UpdateAssetDto) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *UpdateAssetDto) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *UpdateAssetDto) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetCustomProps

`func (o *UpdateAssetDto) GetCustomProps() map[string]interface{}`

GetCustomProps returns the CustomProps field if non-nil, zero value otherwise.

### GetCustomPropsOk

`func (o *UpdateAssetDto) GetCustomPropsOk() (*map[string]interface{}, bool)`

GetCustomPropsOk returns a tuple with the CustomProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProps

`func (o *UpdateAssetDto) SetCustomProps(v map[string]interface{})`

SetCustomProps sets CustomProps field to given value.

### HasCustomProps

`func (o *UpdateAssetDto) HasCustomProps() bool`

HasCustomProps returns a boolean if a field has been set.

### GetCollectionId

`func (o *UpdateAssetDto) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *UpdateAssetDto) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *UpdateAssetDto) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetProjectId

`func (o *UpdateAssetDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateAssetDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateAssetDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


