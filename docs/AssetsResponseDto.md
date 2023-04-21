# AssetsResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]Asset**](Asset.md) | Number of assets to return per page | 
**Total** | **float32** | Number of total items | 
**Limit** | **float32** | Number of items to return per page | 
**Page** | **float32** | Page number | 

## Methods

### NewAssetsResponseDto

`func NewAssetsResponseDto(items []Asset, total float32, limit float32, page float32, ) *AssetsResponseDto`

NewAssetsResponseDto instantiates a new AssetsResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsResponseDtoWithDefaults

`func NewAssetsResponseDtoWithDefaults() *AssetsResponseDto`

NewAssetsResponseDtoWithDefaults instantiates a new AssetsResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AssetsResponseDto) GetItems() []Asset`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AssetsResponseDto) GetItemsOk() (*[]Asset, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AssetsResponseDto) SetItems(v []Asset)`

SetItems sets Items field to given value.


### GetTotal

`func (o *AssetsResponseDto) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AssetsResponseDto) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AssetsResponseDto) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *AssetsResponseDto) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AssetsResponseDto) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AssetsResponseDto) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetPage

`func (o *AssetsResponseDto) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AssetsResponseDto) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AssetsResponseDto) SetPage(v float32)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


