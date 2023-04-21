# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the transaction entity. | 
**Status** | **string** | Transaction status in Blockchain. Can be pending, success or fail | 
**From** | **string** | Address of the sender of the transaction. | 
**To** | **string** | Address of the receiver of the transaction. It can be contract address or player address if it is a transfer transaction. | 
**ContractId** | **string** | Contract ID the transaction is associated with. | 
**ContractAddress** | **string** | Contract address where the transaction happened. | 
**Blockchain** | **string** | Main blockchain identifier: Ethereum, Polygon, etc. | 
**ChainName** | **string** | Chain name identifier: Ethereum, Goerli, Sepolia, PolygonPoS, etc. | 
**ChainId** | **float32** | Chain ID: 1 for Ethereum, 5 for Goerli, 80001 for Polygon Mumbai, etc. | 
**Environment** | **string** | Chain environment: Testnet, Mainnet, etc. | 
**TransactionHash** | **string** | Unique transaction hash in the blockchain. | 
**BlockNumber** | **float32** | Unique block number in the blockchain. | 
**Type** | **string** | Transaction type: Mint, Transfer, Award, Airdrop, etc. | 
**Method** | **string** | Function method name that was called in smart contract | 
**Events** | **[]string** | List of events that were emitted in the transaction | 
**ProjectId** | **string** | The unique identifier of the project that the transaction is associated with. This allows developers to organize their transactions by project and helps with tracking and reporting. | 
**CollectionId** | **string** | The unique identifier of the collection that the transaction is associated with. This allows developers to organize their transactions by project and helps with tracking and reporting. | 
**AccountId** | **float32** | The unique identifier of the account that the transaction belongs to. | 
**CreatedOn** | **time.Time** | The date when the collection was created. | 
**ModifiedOn** | **time.Time** | The date when the collection was last modified. | 
**CreatedBy** | **float32** | The Id of the user who created the collection. | 
**ModifiedBy** | **float32** | The Id of the user who last modified the collection. | 

## Methods

### NewTransaction

`func NewTransaction(id string, status string, from string, to string, contractId string, contractAddress string, blockchain string, chainName string, chainId float32, environment string, transactionHash string, blockNumber float32, type_ string, method string, events []string, projectId string, collectionId string, accountId float32, createdOn time.Time, modifiedOn time.Time, createdBy float32, modifiedBy float32, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFrom

`func (o *Transaction) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Transaction) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Transaction) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *Transaction) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Transaction) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Transaction) SetTo(v string)`

SetTo sets To field to given value.


### GetContractId

`func (o *Transaction) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *Transaction) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *Transaction) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetContractAddress

`func (o *Transaction) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *Transaction) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *Transaction) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetBlockchain

`func (o *Transaction) GetBlockchain() string`

GetBlockchain returns the Blockchain field if non-nil, zero value otherwise.

### GetBlockchainOk

`func (o *Transaction) GetBlockchainOk() (*string, bool)`

GetBlockchainOk returns a tuple with the Blockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchain

`func (o *Transaction) SetBlockchain(v string)`

SetBlockchain sets Blockchain field to given value.


### GetChainName

`func (o *Transaction) GetChainName() string`

GetChainName returns the ChainName field if non-nil, zero value otherwise.

### GetChainNameOk

`func (o *Transaction) GetChainNameOk() (*string, bool)`

GetChainNameOk returns a tuple with the ChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainName

`func (o *Transaction) SetChainName(v string)`

SetChainName sets ChainName field to given value.


### GetChainId

`func (o *Transaction) GetChainId() float32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *Transaction) GetChainIdOk() (*float32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *Transaction) SetChainId(v float32)`

SetChainId sets ChainId field to given value.


### GetEnvironment

`func (o *Transaction) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Transaction) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Transaction) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetTransactionHash

`func (o *Transaction) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *Transaction) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *Transaction) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.


### GetBlockNumber

`func (o *Transaction) GetBlockNumber() float32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *Transaction) GetBlockNumberOk() (*float32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *Transaction) SetBlockNumber(v float32)`

SetBlockNumber sets BlockNumber field to given value.


### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.


### GetMethod

`func (o *Transaction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *Transaction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *Transaction) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetEvents

`func (o *Transaction) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Transaction) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Transaction) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetProjectId

`func (o *Transaction) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Transaction) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Transaction) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCollectionId

`func (o *Transaction) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Transaction) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Transaction) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetAccountId

`func (o *Transaction) GetAccountId() float32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transaction) GetAccountIdOk() (*float32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transaction) SetAccountId(v float32)`

SetAccountId sets AccountId field to given value.


### GetCreatedOn

`func (o *Transaction) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Transaction) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Transaction) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.


### GetModifiedOn

`func (o *Transaction) GetModifiedOn() time.Time`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *Transaction) GetModifiedOnOk() (*time.Time, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *Transaction) SetModifiedOn(v time.Time)`

SetModifiedOn sets ModifiedOn field to given value.


### GetCreatedBy

`func (o *Transaction) GetCreatedBy() float32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Transaction) GetCreatedByOk() (*float32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Transaction) SetCreatedBy(v float32)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *Transaction) GetModifiedBy() float32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Transaction) GetModifiedByOk() (*float32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Transaction) SetModifiedBy(v float32)`

SetModifiedBy sets ModifiedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


