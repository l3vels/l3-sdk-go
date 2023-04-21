# SetSaleStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Project Id | 
**CollectionId** | **string** | Collection Id | 
**SaleStatus** | **map[string]interface{}** | Sale Status of the contract | 

## Methods

### NewSetSaleStatusDto

`func NewSetSaleStatusDto(projectId string, collectionId string, saleStatus map[string]interface{}, ) *SetSaleStatusDto`

NewSetSaleStatusDto instantiates a new SetSaleStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetSaleStatusDtoWithDefaults

`func NewSetSaleStatusDtoWithDefaults() *SetSaleStatusDto`

NewSetSaleStatusDtoWithDefaults instantiates a new SetSaleStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *SetSaleStatusDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SetSaleStatusDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SetSaleStatusDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCollectionId

`func (o *SetSaleStatusDto) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *SetSaleStatusDto) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *SetSaleStatusDto) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetSaleStatus

`func (o *SetSaleStatusDto) GetSaleStatus() map[string]interface{}`

GetSaleStatus returns the SaleStatus field if non-nil, zero value otherwise.

### GetSaleStatusOk

`func (o *SetSaleStatusDto) GetSaleStatusOk() (*map[string]interface{}, bool)`

GetSaleStatusOk returns a tuple with the SaleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleStatus

`func (o *SetSaleStatusDto) SetSaleStatus(v map[string]interface{})`

SetSaleStatus sets SaleStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


