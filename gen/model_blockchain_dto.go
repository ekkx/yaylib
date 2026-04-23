
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BlockchainDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockchainDTO{}

// BlockchainDTO struct for BlockchainDTO
type BlockchainDTO struct {
	AdditionGasPercent NullableAdditionGasPercentDTO `json:"addition_gas_percent,omitempty"`
	BlockExplorer NullableString `json:"block_explorer,omitempty"`
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ChainName NullableString `json:"chain_name,omitempty"`
	ImageUrl NullableString `json:"image_url,omitempty"`
	NftCollections []NftCollectionDTO `json:"nft_collections,omitempty"`
	Priority NullableInt32 `json:"priority,omitempty"`
	RpcUrl NullableString `json:"rpc_url,omitempty"`
	Tokens []TokenDTO `json:"tokens,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BlockchainDTO BlockchainDTO

// NewBlockchainDTO instantiates a new BlockchainDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainDTO() *BlockchainDTO {
	this := BlockchainDTO{}
	return &this
}

// NewBlockchainDTOWithDefaults instantiates a new BlockchainDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainDTOWithDefaults() *BlockchainDTO {
	this := BlockchainDTO{}
	return &this
}

// GetAdditionGasPercent returns the AdditionGasPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetAdditionGasPercent() AdditionGasPercentDTO {
	if o == nil || IsNil(o.AdditionGasPercent.Get()) {
		var ret AdditionGasPercentDTO
		return ret
	}
	return *o.AdditionGasPercent.Get()
}

// GetAdditionGasPercentOk returns a tuple with the AdditionGasPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetAdditionGasPercentOk() (*AdditionGasPercentDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionGasPercent.Get(), o.AdditionGasPercent.IsSet()
}

// HasAdditionGasPercent returns a boolean if a field has been set.
func (o *BlockchainDTO) HasAdditionGasPercent() bool {
	if o != nil && o.AdditionGasPercent.IsSet() {
		return true
	}

	return false
}

// SetAdditionGasPercent gets a reference to the given NullableAdditionGasPercentDTO and assigns it to the AdditionGasPercent field.
func (o *BlockchainDTO) SetAdditionGasPercent(v AdditionGasPercentDTO) {
	o.AdditionGasPercent.Set(&v)
}
// SetAdditionGasPercentNil sets the value for AdditionGasPercent to be an explicit nil
func (o *BlockchainDTO) SetAdditionGasPercentNil() {
	o.AdditionGasPercent.Set(nil)
}

// UnsetAdditionGasPercent ensures that no value is present for AdditionGasPercent, not even an explicit nil
func (o *BlockchainDTO) UnsetAdditionGasPercent() {
	o.AdditionGasPercent.Unset()
}

// GetBlockExplorer returns the BlockExplorer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetBlockExplorer() string {
	if o == nil || IsNil(o.BlockExplorer.Get()) {
		var ret string
		return ret
	}
	return *o.BlockExplorer.Get()
}

// GetBlockExplorerOk returns a tuple with the BlockExplorer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetBlockExplorerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockExplorer.Get(), o.BlockExplorer.IsSet()
}

// HasBlockExplorer returns a boolean if a field has been set.
func (o *BlockchainDTO) HasBlockExplorer() bool {
	if o != nil && o.BlockExplorer.IsSet() {
		return true
	}

	return false
}

// SetBlockExplorer gets a reference to the given NullableString and assigns it to the BlockExplorer field.
func (o *BlockchainDTO) SetBlockExplorer(v string) {
	o.BlockExplorer.Set(&v)
}
// SetBlockExplorerNil sets the value for BlockExplorer to be an explicit nil
func (o *BlockchainDTO) SetBlockExplorerNil() {
	o.BlockExplorer.Set(nil)
}

// UnsetBlockExplorer ensures that no value is present for BlockExplorer, not even an explicit nil
func (o *BlockchainDTO) UnsetBlockExplorer() {
	o.BlockExplorer.Unset()
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *BlockchainDTO) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *BlockchainDTO) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *BlockchainDTO) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *BlockchainDTO) UnsetChainId() {
	o.ChainId.Unset()
}

// GetChainName returns the ChainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetChainName() string {
	if o == nil || IsNil(o.ChainName.Get()) {
		var ret string
		return ret
	}
	return *o.ChainName.Get()
}

// GetChainNameOk returns a tuple with the ChainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetChainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainName.Get(), o.ChainName.IsSet()
}

// HasChainName returns a boolean if a field has been set.
func (o *BlockchainDTO) HasChainName() bool {
	if o != nil && o.ChainName.IsSet() {
		return true
	}

	return false
}

// SetChainName gets a reference to the given NullableString and assigns it to the ChainName field.
func (o *BlockchainDTO) SetChainName(v string) {
	o.ChainName.Set(&v)
}
// SetChainNameNil sets the value for ChainName to be an explicit nil
func (o *BlockchainDTO) SetChainNameNil() {
	o.ChainName.Set(nil)
}

// UnsetChainName ensures that no value is present for ChainName, not even an explicit nil
func (o *BlockchainDTO) UnsetChainName() {
	o.ChainName.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *BlockchainDTO) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *BlockchainDTO) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *BlockchainDTO) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *BlockchainDTO) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetNftCollections returns the NftCollections field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetNftCollections() []NftCollectionDTO {
	if o == nil {
		var ret []NftCollectionDTO
		return ret
	}
	return o.NftCollections
}

// GetNftCollectionsOk returns a tuple with the NftCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetNftCollectionsOk() ([]NftCollectionDTO, bool) {
	if o == nil || IsNil(o.NftCollections) {
		return nil, false
	}
	return o.NftCollections, true
}

// HasNftCollections returns a boolean if a field has been set.
func (o *BlockchainDTO) HasNftCollections() bool {
	if o != nil && !IsNil(o.NftCollections) {
		return true
	}

	return false
}

// SetNftCollections gets a reference to the given []NftCollectionDTO and assigns it to the NftCollections field.
func (o *BlockchainDTO) SetNftCollections(v []NftCollectionDTO) {
	o.NftCollections = v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *BlockchainDTO) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *BlockchainDTO) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *BlockchainDTO) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *BlockchainDTO) UnsetPriority() {
	o.Priority.Unset()
}

// GetRpcUrl returns the RpcUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetRpcUrl() string {
	if o == nil || IsNil(o.RpcUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RpcUrl.Get()
}

// GetRpcUrlOk returns a tuple with the RpcUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetRpcUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpcUrl.Get(), o.RpcUrl.IsSet()
}

// HasRpcUrl returns a boolean if a field has been set.
func (o *BlockchainDTO) HasRpcUrl() bool {
	if o != nil && o.RpcUrl.IsSet() {
		return true
	}

	return false
}

// SetRpcUrl gets a reference to the given NullableString and assigns it to the RpcUrl field.
func (o *BlockchainDTO) SetRpcUrl(v string) {
	o.RpcUrl.Set(&v)
}
// SetRpcUrlNil sets the value for RpcUrl to be an explicit nil
func (o *BlockchainDTO) SetRpcUrlNil() {
	o.RpcUrl.Set(nil)
}

// UnsetRpcUrl ensures that no value is present for RpcUrl, not even an explicit nil
func (o *BlockchainDTO) UnsetRpcUrl() {
	o.RpcUrl.Unset()
}

// GetTokens returns the Tokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockchainDTO) GetTokens() []TokenDTO {
	if o == nil {
		var ret []TokenDTO
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockchainDTO) GetTokensOk() ([]TokenDTO, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *BlockchainDTO) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []TokenDTO and assigns it to the Tokens field.
func (o *BlockchainDTO) SetTokens(v []TokenDTO) {
	o.Tokens = v
}

func (o BlockchainDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockchainDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionGasPercent.IsSet() {
		toSerialize["addition_gas_percent"] = o.AdditionGasPercent.Get()
	}
	if o.BlockExplorer.IsSet() {
		toSerialize["block_explorer"] = o.BlockExplorer.Get()
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ChainName.IsSet() {
		toSerialize["chain_name"] = o.ChainName.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.NftCollections != nil {
		toSerialize["nft_collections"] = o.NftCollections
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.RpcUrl.IsSet() {
		toSerialize["rpc_url"] = o.RpcUrl.Get()
	}
	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BlockchainDTO) UnmarshalJSON(data []byte) (err error) {
	varBlockchainDTO := _BlockchainDTO{}

	err = json.Unmarshal(data, &varBlockchainDTO)

	if err != nil {
		return err
	}

	*o = BlockchainDTO(varBlockchainDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addition_gas_percent")
		delete(additionalProperties, "block_explorer")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "chain_name")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "nft_collections")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "rpc_url")
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBlockchainDTO struct {
	value *BlockchainDTO
	isSet bool
}

func (v NullableBlockchainDTO) Get() *BlockchainDTO {
	return v.value
}

func (v *NullableBlockchainDTO) Set(val *BlockchainDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainDTO(val *BlockchainDTO) *NullableBlockchainDTO {
	return &NullableBlockchainDTO{value: val, isSet: true}
}

func (v NullableBlockchainDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


