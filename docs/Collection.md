# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the collection entity. | 
**UniqueId** | **string** | The unique identifier that can be provided by game studio. | 
**Name** | **string** | The name of the collection. | 
**Description** | **string** | A brief description of the collection. | 
**LogoImage** | **string** | An image representing the collection&#39;s logo. | 
**Medias** | **[]string** | Additional images associated with the collection, such as screenshots or promotional images. | 
**MainMedia** | **string** | The main or featured image associated with the game. You can set it from the dashboard as main image. | 
**Url** | **string** | A unique URL for the collection on the L3vels platform. | 
**WebLink** | **string** | A URL link to the collection&#39;s webpage. | 
**Supply** | **float32** | The supply represents the number of assets that are available within the Collection. | 
**CustomPropertyProps** | **map[string]interface{}** | Custom properties that are unique to the collection and defined by the developer to categorize and filter them. | 
**SocialLinks** | **[]string** | Additional social links associated with the collection | 
**CustomAssetProps** | **map[string]interface{}** | Custom assets fields associated with the collection. | 
**Categories** | **map[string]interface{}** | The category or categories that the collection belongs to. | 
**Status** | **string** | The current status of the collection. Possible values are: Draft, Active | 
**AccountId** | **float32** | The unique identifier of the account that the Collection belongs to. | 
**ProjectId** | **string** | The unique identifier of the project that the collection is associated with. This allows developers to organize their collections by project and helps with tracking and reporting. | 
**CreatedOn** | **time.Time** | The date when the collection was created. | 
**ModifiedOn** | **time.Time** | The date when the collection was last modified. | 
**CreatedBy** | **float32** | The Id of the user who created the collection. | 
**ModifiedBy** | **float32** | The Id of the user who last modified the collection. | 

## Methods

### NewCollection

`func NewCollection(id string, uniqueId string, name string, description string, logoImage string, medias []string, mainMedia string, url string, webLink string, supply float32, customPropertyProps map[string]interface{}, socialLinks []string, customAssetProps map[string]interface{}, categories map[string]interface{}, status string, accountId float32, projectId string, createdOn time.Time, modifiedOn time.Time, createdBy float32, modifiedBy float32, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Collection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Collection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Collection) SetId(v string)`

SetId sets Id field to given value.


### GetUniqueId

`func (o *Collection) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *Collection) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *Collection) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.


### GetName

`func (o *Collection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collection) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Collection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Collection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Collection) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogoImage

`func (o *Collection) GetLogoImage() string`

GetLogoImage returns the LogoImage field if non-nil, zero value otherwise.

### GetLogoImageOk

`func (o *Collection) GetLogoImageOk() (*string, bool)`

GetLogoImageOk returns a tuple with the LogoImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImage

`func (o *Collection) SetLogoImage(v string)`

SetLogoImage sets LogoImage field to given value.


### GetMedias

`func (o *Collection) GetMedias() []string`

GetMedias returns the Medias field if non-nil, zero value otherwise.

### GetMediasOk

`func (o *Collection) GetMediasOk() (*[]string, bool)`

GetMediasOk returns a tuple with the Medias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedias

`func (o *Collection) SetMedias(v []string)`

SetMedias sets Medias field to given value.


### GetMainMedia

`func (o *Collection) GetMainMedia() string`

GetMainMedia returns the MainMedia field if non-nil, zero value otherwise.

### GetMainMediaOk

`func (o *Collection) GetMainMediaOk() (*string, bool)`

GetMainMediaOk returns a tuple with the MainMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainMedia

`func (o *Collection) SetMainMedia(v string)`

SetMainMedia sets MainMedia field to given value.


### GetUrl

`func (o *Collection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Collection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Collection) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWebLink

`func (o *Collection) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *Collection) GetWebLinkOk() (*string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLink

`func (o *Collection) SetWebLink(v string)`

SetWebLink sets WebLink field to given value.


### GetSupply

`func (o *Collection) GetSupply() float32`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *Collection) GetSupplyOk() (*float32, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *Collection) SetSupply(v float32)`

SetSupply sets Supply field to given value.


### GetCustomPropertyProps

`func (o *Collection) GetCustomPropertyProps() map[string]interface{}`

GetCustomPropertyProps returns the CustomPropertyProps field if non-nil, zero value otherwise.

### GetCustomPropertyPropsOk

`func (o *Collection) GetCustomPropertyPropsOk() (*map[string]interface{}, bool)`

GetCustomPropertyPropsOk returns a tuple with the CustomPropertyProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPropertyProps

`func (o *Collection) SetCustomPropertyProps(v map[string]interface{})`

SetCustomPropertyProps sets CustomPropertyProps field to given value.


### GetSocialLinks

`func (o *Collection) GetSocialLinks() []string`

GetSocialLinks returns the SocialLinks field if non-nil, zero value otherwise.

### GetSocialLinksOk

`func (o *Collection) GetSocialLinksOk() (*[]string, bool)`

GetSocialLinksOk returns a tuple with the SocialLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLinks

`func (o *Collection) SetSocialLinks(v []string)`

SetSocialLinks sets SocialLinks field to given value.


### GetCustomAssetProps

`func (o *Collection) GetCustomAssetProps() map[string]interface{}`

GetCustomAssetProps returns the CustomAssetProps field if non-nil, zero value otherwise.

### GetCustomAssetPropsOk

`func (o *Collection) GetCustomAssetPropsOk() (*map[string]interface{}, bool)`

GetCustomAssetPropsOk returns a tuple with the CustomAssetProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAssetProps

`func (o *Collection) SetCustomAssetProps(v map[string]interface{})`

SetCustomAssetProps sets CustomAssetProps field to given value.


### GetCategories

`func (o *Collection) GetCategories() map[string]interface{}`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Collection) GetCategoriesOk() (*map[string]interface{}, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Collection) SetCategories(v map[string]interface{})`

SetCategories sets Categories field to given value.


### GetStatus

`func (o *Collection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Collection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Collection) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountId

`func (o *Collection) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Collection) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Collection) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.


### GetProjectId

`func (o *Collection) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Collection) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Collection) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedOn

`func (o *Collection) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Collection) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Collection) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *Collection) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Collection) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Collection) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.


### GetCreatedBy

`func (o *Collection) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Collection) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Collection) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *Collection) GetModifiedBy() float32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Collection) GetModifiedByOk() (*float32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Collection) SetModifiedBy(v float32)`

SetModifiedBy sets ModifiedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


