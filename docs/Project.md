# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the Game entity. | 
**Name** | **string** | The name of the Game. | 
**Description** | **string** | A brief description of the Game. | 
**Category** | **string** | The category or genre of the Game. | 
**LogoImage** | **string** | The logo or icon associated with the Game. | 
**Medias** | **[]string** | Additional images associated with the collection, such as screenshots or promotional images. | 
**SocialLinks** | **[]string** | Additional social links associated with the collection | 
**MainMedia** | **string** | The main or featured image associated with the Game. You can set it from the Dashboard as main image. | 
**Url** | **string** | A unique URL for the game on the L3vels platform. | 
**WebLink** | **string** | The URL for the Game&#39;s website or landing page. | 
**Discord** | **string** | The link to the Game&#39;s Discord server. | 
**Twitter** | **string** | The link to the Game&#39;s official Twitter account. | 
**Instagram** | **string** | The link to the Game&#39;s official Instagram account. | 
**ContactPhone** | **string** |  A phone number for contacting the Game&#39;s developers or support team. | 
**ContactEmail** | **string** | An email address for contacting the Game&#39;s developers or support team. | 
**Status** | **string** | The current status of the Game, such as \&quot;Draft\&quot;, \&quot;Active\&quot;, or \&quot;Inactive\&quot;. | 
**AccountId** | **float32** |  The unique identifier of the account that the Game belongs to. This allows developers to manage multiple Games across multiple accounts. | 
**CreatedOn** | **time.Time** | The date and time that the Game entity was created. | 
**ModifiedOn** | **time.Time** | The date and time that the Game entity was last modified. | 
**CreatedBy** | **float32** | The user or system that created the Game entity. | 
**ModifiedBy** | **float32** | The user or system that last modified the Game entity. | 

## Methods

### NewProject

`func NewProject(id string, name string, description string, category string, logoImage string, medias []string, socialLinks []string, mainMedia string, url string, webLink string, discord string, twitter string, instagram string, contactPhone string, contactEmail string, status string, accountId float32, createdOn time.Time, modifiedOn time.Time, createdBy float32, modifiedBy float32, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Project) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategory

`func (o *Project) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Project) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Project) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetLogoImage

`func (o *Project) GetLogoImage() string`

GetLogoImage returns the LogoImage field if non-nil, zero value otherwise.

### GetLogoImageOk

`func (o *Project) GetLogoImageOk() (*string, bool)`

GetLogoImageOk returns a tuple with the LogoImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImage

`func (o *Project) SetLogoImage(v string)`

SetLogoImage sets LogoImage field to given value.


### GetMedias

`func (o *Project) GetMedias() []string`

GetMedias returns the Medias field if non-nil, zero value otherwise.

### GetMediasOk

`func (o *Project) GetMediasOk() (*[]string, bool)`

GetMediasOk returns a tuple with the Medias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedias

`func (o *Project) SetMedias(v []string)`

SetMedias sets Medias field to given value.


### GetSocialLinks

`func (o *Project) GetSocialLinks() []string`

GetSocialLinks returns the SocialLinks field if non-nil, zero value otherwise.

### GetSocialLinksOk

`func (o *Project) GetSocialLinksOk() (*[]string, bool)`

GetSocialLinksOk returns a tuple with the SocialLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLinks

`func (o *Project) SetSocialLinks(v []string)`

SetSocialLinks sets SocialLinks field to given value.


### GetMainMedia

`func (o *Project) GetMainMedia() string`

GetMainMedia returns the MainMedia field if non-nil, zero value otherwise.

### GetMainMediaOk

`func (o *Project) GetMainMediaOk() (*string, bool)`

GetMainMediaOk returns a tuple with the MainMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainMedia

`func (o *Project) SetMainMedia(v string)`

SetMainMedia sets MainMedia field to given value.


### GetUrl

`func (o *Project) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Project) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Project) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWebLink

`func (o *Project) GetWebLink() string`

GetWebLink returns the WebLink field if non-nil, zero value otherwise.

### GetWebLinkOk

`func (o *Project) GetWebLinkOk() (*string, bool)`

GetWebLinkOk returns a tuple with the WebLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebLink

`func (o *Project) SetWebLink(v string)`

SetWebLink sets WebLink field to given value.


### GetDiscord

`func (o *Project) GetDiscord() string`

GetDiscord returns the Discord field if non-nil, zero value otherwise.

### GetDiscordOk

`func (o *Project) GetDiscordOk() (*string, bool)`

GetDiscordOk returns a tuple with the Discord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscord

`func (o *Project) SetDiscord(v string)`

SetDiscord sets Discord field to given value.


### GetTwitter

`func (o *Project) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *Project) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *Project) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.


### GetInstagram

`func (o *Project) GetInstagram() string`

GetInstagram returns the Instagram field if non-nil, zero value otherwise.

### GetInstagramOk

`func (o *Project) GetInstagramOk() (*string, bool)`

GetInstagramOk returns a tuple with the Instagram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagram

`func (o *Project) SetInstagram(v string)`

SetInstagram sets Instagram field to given value.


### GetContactPhone

`func (o *Project) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *Project) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *Project) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.


### GetContactEmail

`func (o *Project) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Project) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Project) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetStatus

`func (o *Project) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Project) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Project) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountId

`func (o *Project) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Project) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Project) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.


### GetCreatedOn

`func (o *Project) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Project) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Project) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *Project) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Project) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Project) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.


### GetCreatedBy

`func (o *Project) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Project) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Project) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *Project) GetModifiedBy() float32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Project) GetModifiedByOk() (*float32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Project) SetModifiedBy(v float32)`

SetModifiedBy sets ModifiedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


