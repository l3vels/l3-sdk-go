# CreatePlayerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueId** | Pointer to **string** | You can send your generated unique ID for player if it&#39;s handy for game. If field is empty, we will generate unique ID. | [optional] 
**Name** | **string** | The name of the player. | 
**Username** | Pointer to **string** | The username of the player. | [optional] 
**Email** | Pointer to **string** | The email of the player. | [optional] 
**Avatar** | Pointer to **string** | The image URL of the player avatar. | [optional] 
**ProjectId** | **string** | Game/project ID to create the player for. Example: Create player Jack for game Fortnite. | 
**IsCreateWallet** | Pointer to **bool** | Should create wallet for player or not. | [optional] 
**CustomProps** | Pointer to **[]string** | Custom props for player. | [optional] 

## Methods

### NewCreatePlayerDto

`func NewCreatePlayerDto(name string, projectId string, ) *CreatePlayerDto`

NewCreatePlayerDto instantiates a new CreatePlayerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlayerDtoWithDefaults

`func NewCreatePlayerDtoWithDefaults() *CreatePlayerDto`

NewCreatePlayerDtoWithDefaults instantiates a new CreatePlayerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniqueId

`func (o *CreatePlayerDto) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *CreatePlayerDto) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *CreatePlayerDto) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *CreatePlayerDto) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetName

`func (o *CreatePlayerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlayerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlayerDto) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *CreatePlayerDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreatePlayerDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreatePlayerDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreatePlayerDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *CreatePlayerDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatePlayerDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatePlayerDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreatePlayerDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAvatar

`func (o *CreatePlayerDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CreatePlayerDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CreatePlayerDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CreatePlayerDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetProjectId

`func (o *CreatePlayerDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreatePlayerDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreatePlayerDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetIsCreateWallet

`func (o *CreatePlayerDto) GetIsCreateWallet() bool`

GetIsCreateWallet returns the IsCreateWallet field if non-nil, zero value otherwise.

### GetIsCreateWalletOk

`func (o *CreatePlayerDto) GetIsCreateWalletOk() (*bool, bool)`

GetIsCreateWalletOk returns a tuple with the IsCreateWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreateWallet

`func (o *CreatePlayerDto) SetIsCreateWallet(v bool)`

SetIsCreateWallet sets IsCreateWallet field to given value.

### HasIsCreateWallet

`func (o *CreatePlayerDto) HasIsCreateWallet() bool`

HasIsCreateWallet returns a boolean if a field has been set.

### GetCustomProps

`func (o *CreatePlayerDto) GetCustomProps() []string`

GetCustomProps returns the CustomProps field if non-nil, zero value otherwise.

### GetCustomPropsOk

`func (o *CreatePlayerDto) GetCustomPropsOk() (*[]string, bool)`

GetCustomPropsOk returns a tuple with the CustomProps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProps

`func (o *CreatePlayerDto) SetCustomProps(v []string)`

SetCustomProps sets CustomProps field to given value.

### HasCustomProps

`func (o *CreatePlayerDto) HasCustomProps() bool`

HasCustomProps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


