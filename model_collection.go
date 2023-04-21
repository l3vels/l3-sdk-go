/*
L3vels Api

L3vels API for Game developers

API version: 0.3
Contact: support@l3vels.xyz
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package l3vels-sdk

import (
	"encoding/json"
	"time"
)

// checks if the Collection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Collection{}

// Collection struct for Collection
type Collection struct {
	// The unique identifier for the collection entity.
	Id string `json:"id"`
	// The unique identifier that can be provided by game studio.
	UniqueId string `json:"unique_id"`
	// The name of the collection.
	Name string `json:"name"`
	// A brief description of the collection.
	Description string `json:"description"`
	// An image representing the collection's logo.
	LogoImage string `json:"logo_image"`
	// Additional images associated with the collection, such as screenshots or promotional images.
	Medias []string `json:"medias"`
	// The main or featured image associated with the game. You can set it from the dashboard as main image.
	MainMedia string `json:"main_media"`
	// A unique URL for the collection on the L3vels platform.
	Url string `json:"url"`
	// A URL link to the collection's webpage.
	WebLink string `json:"web_link"`
	// The supply represents the number of assets that are available within the Collection.
	Supply float32 `json:"supply"`
	// Custom properties that are unique to the collection and defined by the developer to categorize and filter them.
	CustomPropertyProps map[string]interface{} `json:"custom_property_props"`
	// Additional social links associated with the collection
	SocialLinks []string `json:"social_links"`
	// Custom assets fields associated with the collection.
	CustomAssetProps map[string]interface{} `json:"custom_asset_props"`
	// The category or categories that the collection belongs to.
	Categories map[string]interface{} `json:"categories"`
	// The current status of the collection. Possible values are: Draft, Active
	Status string `json:"status"`
	// The unique identifier of the account that the Collection belongs to.
	AccountId float32 `json:"account_id"`
	// The unique identifier of the project that the collection is associated with. This allows developers to organize their collections by project and helps with tracking and reporting.
	ProjectId string `json:"project_id"`
	// The date when the collection was created.
	CreatedOn time.Time `json:"created_on"`
	// The date when the collection was last modified.
	ModifiedOn time.Time `json:"modified_on"`
	// The Id of the user who created the collection.
	CreatedBy float32 `json:"created_by"`
	// The Id of the user who last modified the collection.
	ModifiedBy float32 `json:"modified_by"`
}

// NewCollection instantiates a new Collection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollection(id string, uniqueId string, name string, description string, logoImage string, medias []string, mainMedia string, url string, webLink string, supply float32, customPropertyProps map[string]interface{}, socialLinks []string, customAssetProps map[string]interface{}, categories map[string]interface{}, status string, accountId float32, projectId string, createdOn time.Time, modifiedOn time.Time, createdBy float32, modifiedBy float32) *Collection {
	this := Collection{}
	this.Id = id
	this.UniqueId = uniqueId
	this.Name = name
	this.Description = description
	this.LogoImage = logoImage
	this.Medias = medias
	this.MainMedia = mainMedia
	this.Url = url
	this.WebLink = webLink
	this.Supply = supply
	this.CustomPropertyProps = customPropertyProps
	this.SocialLinks = socialLinks
	this.CustomAssetProps = customAssetProps
	this.Categories = categories
	this.Status = status
	this.AccountId = accountId
	this.ProjectId = projectId
	this.CreatedOn = createdOn
	this.ModifiedOn = modifiedOn
	this.CreatedBy = createdBy
	this.ModifiedBy = modifiedBy
	return &this
}

// NewCollectionWithDefaults instantiates a new Collection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionWithDefaults() *Collection {
	this := Collection{}
	return &this
}

// GetId returns the Id field value
func (o *Collection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Collection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Collection) SetId(v string) {
	o.Id = v
}

// GetUniqueId returns the UniqueId field value
func (o *Collection) GetUniqueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value
// and a boolean to check if the value has been set.
func (o *Collection) GetUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueId, true
}

// SetUniqueId sets field value
func (o *Collection) SetUniqueId(v string) {
	o.UniqueId = v
}

// GetName returns the Name field value
func (o *Collection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Collection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Collection) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Collection) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Collection) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Collection) SetDescription(v string) {
	o.Description = v
}

// GetLogoImage returns the LogoImage field value
func (o *Collection) GetLogoImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoImage
}

// GetLogoImageOk returns a tuple with the LogoImage field value
// and a boolean to check if the value has been set.
func (o *Collection) GetLogoImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoImage, true
}

// SetLogoImage sets field value
func (o *Collection) SetLogoImage(v string) {
	o.LogoImage = v
}

// GetMedias returns the Medias field value
func (o *Collection) GetMedias() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Medias
}

// GetMediasOk returns a tuple with the Medias field value
// and a boolean to check if the value has been set.
func (o *Collection) GetMediasOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Medias, true
}

// SetMedias sets field value
func (o *Collection) SetMedias(v []string) {
	o.Medias = v
}

// GetMainMedia returns the MainMedia field value
func (o *Collection) GetMainMedia() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MainMedia
}

// GetMainMediaOk returns a tuple with the MainMedia field value
// and a boolean to check if the value has been set.
func (o *Collection) GetMainMediaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MainMedia, true
}

// SetMainMedia sets field value
func (o *Collection) SetMainMedia(v string) {
	o.MainMedia = v
}

// GetUrl returns the Url field value
func (o *Collection) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Collection) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Collection) SetUrl(v string) {
	o.Url = v
}

// GetWebLink returns the WebLink field value
func (o *Collection) GetWebLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebLink
}

// GetWebLinkOk returns a tuple with the WebLink field value
// and a boolean to check if the value has been set.
func (o *Collection) GetWebLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebLink, true
}

// SetWebLink sets field value
func (o *Collection) SetWebLink(v string) {
	o.WebLink = v
}

// GetSupply returns the Supply field value
func (o *Collection) GetSupply() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Supply
}

// GetSupplyOk returns a tuple with the Supply field value
// and a boolean to check if the value has been set.
func (o *Collection) GetSupplyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supply, true
}

// SetSupply sets field value
func (o *Collection) SetSupply(v float32) {
	o.Supply = v
}

// GetCustomPropertyProps returns the CustomPropertyProps field value
func (o *Collection) GetCustomPropertyProps() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomPropertyProps
}

// GetCustomPropertyPropsOk returns a tuple with the CustomPropertyProps field value
// and a boolean to check if the value has been set.
func (o *Collection) GetCustomPropertyPropsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CustomPropertyProps, true
}

// SetCustomPropertyProps sets field value
func (o *Collection) SetCustomPropertyProps(v map[string]interface{}) {
	o.CustomPropertyProps = v
}

// GetSocialLinks returns the SocialLinks field value
func (o *Collection) GetSocialLinks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SocialLinks
}

// GetSocialLinksOk returns a tuple with the SocialLinks field value
// and a boolean to check if the value has been set.
func (o *Collection) GetSocialLinksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SocialLinks, true
}

// SetSocialLinks sets field value
func (o *Collection) SetSocialLinks(v []string) {
	o.SocialLinks = v
}

// GetCustomAssetProps returns the CustomAssetProps field value
func (o *Collection) GetCustomAssetProps() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomAssetProps
}

// GetCustomAssetPropsOk returns a tuple with the CustomAssetProps field value
// and a boolean to check if the value has been set.
func (o *Collection) GetCustomAssetPropsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.CustomAssetProps, true
}

// SetCustomAssetProps sets field value
func (o *Collection) SetCustomAssetProps(v map[string]interface{}) {
	o.CustomAssetProps = v
}

// GetCategories returns the Categories field value
func (o *Collection) GetCategories() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *Collection) GetCategoriesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *Collection) SetCategories(v map[string]interface{}) {
	o.Categories = v
}

// GetStatus returns the Status field value
func (o *Collection) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Collection) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Collection) SetStatus(v string) {
	o.Status = v
}

// GetAccountId returns the AccountId field value
func (o *Collection) GetAccountId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Collection) GetAccountIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Collection) SetAccountId(v float32) {
	o.AccountId = v
}

// GetProjectId returns the ProjectId field value
func (o *Collection) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Collection) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Collection) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *Collection) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *Collection) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *Collection) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *Collection) GetModifiedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *Collection) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *Collection) SetModifiedOn(v time.Time) {
	o.ModifiedOn = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Collection) GetCreatedBy() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Collection) GetCreatedByOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Collection) SetCreatedBy(v float32) {
	o.CreatedBy = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *Collection) GetModifiedBy() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *Collection) GetModifiedByOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *Collection) SetModifiedBy(v float32) {
	o.ModifiedBy = v
}

func (o Collection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Collection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["unique_id"] = o.UniqueId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["logo_image"] = o.LogoImage
	toSerialize["medias"] = o.Medias
	toSerialize["main_media"] = o.MainMedia
	toSerialize["url"] = o.Url
	toSerialize["web_link"] = o.WebLink
	toSerialize["supply"] = o.Supply
	toSerialize["custom_property_props"] = o.CustomPropertyProps
	toSerialize["social_links"] = o.SocialLinks
	toSerialize["custom_asset_props"] = o.CustomAssetProps
	toSerialize["categories"] = o.Categories
	toSerialize["status"] = o.Status
	toSerialize["account_id"] = o.AccountId
	toSerialize["project_id"] = o.ProjectId
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["modified_by"] = o.ModifiedBy
	return toSerialize, nil
}

type NullableCollection struct {
	value *Collection
	isSet bool
}

func (v NullableCollection) Get() *Collection {
	return v.value
}

func (v *NullableCollection) Set(val *Collection) {
	v.value = val
	v.isSet = true
}

func (v NullableCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollection(val *Collection) *NullableCollection {
	return &NullableCollection{value: val, isSet: true}
}

func (v NullableCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

