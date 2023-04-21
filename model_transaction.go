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

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction struct for Transaction
type Transaction struct {
	// The unique identifier for the transaction entity.
	Id string `json:"id"`
	// Transaction status in Blockchain. Can be pending, success or fail
	Status string `json:"status"`
	// Address of the sender of the transaction.
	From string `json:"from"`
	// Address of the receiver of the transaction. It can be contract address or player address if it is a transfer transaction.
	To string `json:"to"`
	// Contract ID the transaction is associated with.
	ContractId string `json:"contract_id"`
	// Contract address where the transaction happened.
	ContractAddress string `json:"contract_address"`
	// Main blockchain identifier: Ethereum, Polygon, etc.
	Blockchain string `json:"blockchain"`
	// Chain name identifier: Ethereum, Goerli, Sepolia, PolygonPoS, etc.
	ChainName string `json:"chain_name"`
	// Chain ID: 1 for Ethereum, 5 for Goerli, 80001 for Polygon Mumbai, etc.
	ChainId float32 `json:"chain_id"`
	// Chain environment: Testnet, Mainnet, etc.
	Environment string `json:"environment"`
	// Unique transaction hash in the blockchain.
	TransactionHash string `json:"transaction_hash"`
	// Unique block number in the blockchain.
	BlockNumber float32 `json:"block_number"`
	// Transaction type: Mint, Transfer, Award, Airdrop, etc.
	Type string `json:"type"`
	// Function method name that was called in smart contract
	Method string `json:"method"`
	// List of events that were emitted in the transaction
	Events []string `json:"events"`
	// The unique identifier of the project that the transaction is associated with. This allows developers to organize their transactions by project and helps with tracking and reporting.
	ProjectId string `json:"project_id"`
	// The unique identifier of the collection that the transaction is associated with. This allows developers to organize their transactions by project and helps with tracking and reporting.
	CollectionId string `json:"collection_id"`
	// The unique identifier of the account that the transaction belongs to.
	AccountId float32 `json:"account_id"`
	// The date when the collection was created.
	CreatedOn time.Time `json:"created_on"`
	// The date when the collection was last modified.
	ModifiedOn time.Time `json:"modified_on"`
	// The Id of the user who created the collection.
	CreatedBy float32 `json:"created_by"`
	// The Id of the user who last modified the collection.
	ModifiedBy float32 `json:"modified_by"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(id string, status string, from string, to string, contractId string, contractAddress string, blockchain string, chainName string, chainId float32, environment string, transactionHash string, blockNumber float32, type_ string, method string, events []string, projectId string, collectionId string, accountId float32, createdOn time.Time, modifiedOn time.Time, createdBy float32, modifiedBy float32) *Transaction {
	this := Transaction{}
	this.Id = id
	this.Status = status
	this.From = from
	this.To = to
	this.ContractId = contractId
	this.ContractAddress = contractAddress
	this.Blockchain = blockchain
	this.ChainName = chainName
	this.ChainId = chainId
	this.Environment = environment
	this.TransactionHash = transactionHash
	this.BlockNumber = blockNumber
	this.Type = type_
	this.Method = method
	this.Events = events
	this.ProjectId = projectId
	this.CollectionId = collectionId
	this.AccountId = accountId
	this.CreatedOn = createdOn
	this.ModifiedOn = modifiedOn
	this.CreatedBy = createdBy
	this.ModifiedBy = modifiedBy
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetId returns the Id field value
func (o *Transaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v string) {
	o.Status = v
}

// GetFrom returns the From field value
func (o *Transaction) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *Transaction) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *Transaction) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Transaction) SetTo(v string) {
	o.To = v
}

// GetContractId returns the ContractId field value
func (o *Transaction) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *Transaction) SetContractId(v string) {
	o.ContractId = v
}

// GetContractAddress returns the ContractAddress field value
func (o *Transaction) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *Transaction) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetBlockchain returns the Blockchain field value
func (o *Transaction) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetBlockchainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *Transaction) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetChainName returns the ChainName field value
func (o *Transaction) GetChainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainName
}

// GetChainNameOk returns a tuple with the ChainName field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetChainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainName, true
}

// SetChainName sets field value
func (o *Transaction) SetChainName(v string) {
	o.ChainName = v
}

// GetChainId returns the ChainId field value
func (o *Transaction) GetChainId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetChainIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *Transaction) SetChainId(v float32) {
	o.ChainId = v
}

// GetEnvironment returns the Environment field value
func (o *Transaction) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *Transaction) SetEnvironment(v string) {
	o.Environment = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *Transaction) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *Transaction) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *Transaction) GetBlockNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetBlockNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *Transaction) SetBlockNumber(v float32) {
	o.BlockNumber = v
}

// GetType returns the Type field value
func (o *Transaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transaction) SetType(v string) {
	o.Type = v
}

// GetMethod returns the Method field value
func (o *Transaction) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *Transaction) SetMethod(v string) {
	o.Method = v
}

// GetEvents returns the Events field value
func (o *Transaction) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *Transaction) SetEvents(v []string) {
	o.Events = v
}

// GetProjectId returns the ProjectId field value
func (o *Transaction) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Transaction) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCollectionId returns the CollectionId field value
func (o *Transaction) GetCollectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCollectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionId, true
}

// SetCollectionId sets field value
func (o *Transaction) SetCollectionId(v string) {
	o.CollectionId = v
}

// GetAccountId returns the AccountId field value
func (o *Transaction) GetAccountId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAccountIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Transaction) SetAccountId(v float32) {
	o.AccountId = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *Transaction) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *Transaction) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *Transaction) GetModifiedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetModifiedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *Transaction) SetModifiedOn(v time.Time) {
	o.ModifiedOn = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Transaction) GetCreatedBy() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedByOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Transaction) SetCreatedBy(v float32) {
	o.CreatedBy = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *Transaction) GetModifiedBy() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetModifiedByOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *Transaction) SetModifiedBy(v float32) {
	o.ModifiedBy = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	toSerialize["contract_id"] = o.ContractId
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["blockchain"] = o.Blockchain
	toSerialize["chain_name"] = o.ChainName
	toSerialize["chain_id"] = o.ChainId
	toSerialize["environment"] = o.Environment
	toSerialize["transaction_hash"] = o.TransactionHash
	toSerialize["block_number"] = o.BlockNumber
	toSerialize["type"] = o.Type
	toSerialize["method"] = o.Method
	toSerialize["events"] = o.Events
	toSerialize["project_id"] = o.ProjectId
	toSerialize["collection_id"] = o.CollectionId
	toSerialize["account_id"] = o.AccountId
	toSerialize["created_on"] = o.CreatedOn
	toSerialize["modified_on"] = o.ModifiedOn
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["modified_by"] = o.ModifiedBy
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

