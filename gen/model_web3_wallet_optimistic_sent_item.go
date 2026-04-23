
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletOptimisticSentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletOptimisticSentItem{}

// Web3WalletOptimisticSentItem struct for Web3WalletOptimisticSentItem
type Web3WalletOptimisticSentItem struct {
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	CreatedAtMillis NullableInt64 `json:"created_at_millis,omitempty"`
	Id NullableString `json:"id,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletOptimisticSentItem Web3WalletOptimisticSentItem

// NewWeb3WalletOptimisticSentItem instantiates a new Web3WalletOptimisticSentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletOptimisticSentItem() *Web3WalletOptimisticSentItem {
	this := Web3WalletOptimisticSentItem{}
	return &this
}

// NewWeb3WalletOptimisticSentItemWithDefaults instantiates a new Web3WalletOptimisticSentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletOptimisticSentItemWithDefaults() *Web3WalletOptimisticSentItem {
	this := Web3WalletOptimisticSentItem{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletOptimisticSentItem) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletOptimisticSentItem) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3WalletOptimisticSentItem) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Web3WalletOptimisticSentItem) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3WalletOptimisticSentItem) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3WalletOptimisticSentItem) UnsetChainId() {
	o.ChainId.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletOptimisticSentItem) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletOptimisticSentItem) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Web3WalletOptimisticSentItem) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *Web3WalletOptimisticSentItem) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *Web3WalletOptimisticSentItem) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *Web3WalletOptimisticSentItem) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetCreatedAtMillis returns the CreatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletOptimisticSentItem) GetCreatedAtMillis() int64 {
	if o == nil || IsNil(o.CreatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtMillis.Get()
}

// GetCreatedAtMillisOk returns a tuple with the CreatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletOptimisticSentItem) GetCreatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtMillis.Get(), o.CreatedAtMillis.IsSet()
}

// HasCreatedAtMillis returns a boolean if a field has been set.
func (o *Web3WalletOptimisticSentItem) HasCreatedAtMillis() bool {
	if o != nil && o.CreatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtMillis gets a reference to the given NullableInt64 and assigns it to the CreatedAtMillis field.
func (o *Web3WalletOptimisticSentItem) SetCreatedAtMillis(v int64) {
	o.CreatedAtMillis.Set(&v)
}
// SetCreatedAtMillisNil sets the value for CreatedAtMillis to be an explicit nil
func (o *Web3WalletOptimisticSentItem) SetCreatedAtMillisNil() {
	o.CreatedAtMillis.Set(nil)
}

// UnsetCreatedAtMillis ensures that no value is present for CreatedAtMillis, not even an explicit nil
func (o *Web3WalletOptimisticSentItem) UnsetCreatedAtMillis() {
	o.CreatedAtMillis.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletOptimisticSentItem) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletOptimisticSentItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Web3WalletOptimisticSentItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *Web3WalletOptimisticSentItem) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Web3WalletOptimisticSentItem) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Web3WalletOptimisticSentItem) UnsetId() {
	o.Id.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletOptimisticSentItem) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletOptimisticSentItem) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *Web3WalletOptimisticSentItem) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *Web3WalletOptimisticSentItem) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *Web3WalletOptimisticSentItem) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *Web3WalletOptimisticSentItem) UnsetTokenId() {
	o.TokenId.Unset()
}

func (o Web3WalletOptimisticSentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletOptimisticSentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.CreatedAtMillis.IsSet() {
		toSerialize["created_at_millis"] = o.CreatedAtMillis.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletOptimisticSentItem) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletOptimisticSentItem := _Web3WalletOptimisticSentItem{}

	err = json.Unmarshal(data, &varWeb3WalletOptimisticSentItem)

	if err != nil {
		return err
	}

	*o = Web3WalletOptimisticSentItem(varWeb3WalletOptimisticSentItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "created_at_millis")
		delete(additionalProperties, "id")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletOptimisticSentItem struct {
	value *Web3WalletOptimisticSentItem
	isSet bool
}

func (v NullableWeb3WalletOptimisticSentItem) Get() *Web3WalletOptimisticSentItem {
	return v.value
}

func (v *NullableWeb3WalletOptimisticSentItem) Set(val *Web3WalletOptimisticSentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletOptimisticSentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletOptimisticSentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletOptimisticSentItem(val *Web3WalletOptimisticSentItem) *NullableWeb3WalletOptimisticSentItem {
	return &NullableWeb3WalletOptimisticSentItem{value: val, isSet: true}
}

func (v NullableWeb3WalletOptimisticSentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletOptimisticSentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


