
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletPreSignApprove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletPreSignApprove{}

// Web3WalletPreSignApprove struct for Web3WalletPreSignApprove
type Web3WalletPreSignApprove struct {
	BlockchainNetworkData NullableWeb3WalletPreSignApproveBCNetworkData `json:"blockchain_network_data,omitempty"`
	TransactionData NullableTXNData `json:"transaction_data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletPreSignApprove Web3WalletPreSignApprove

// NewWeb3WalletPreSignApprove instantiates a new Web3WalletPreSignApprove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletPreSignApprove() *Web3WalletPreSignApprove {
	this := Web3WalletPreSignApprove{}
	return &this
}

// NewWeb3WalletPreSignApproveWithDefaults instantiates a new Web3WalletPreSignApprove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletPreSignApproveWithDefaults() *Web3WalletPreSignApprove {
	this := Web3WalletPreSignApprove{}
	return &this
}

// GetBlockchainNetworkData returns the BlockchainNetworkData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSignApprove) GetBlockchainNetworkData() Web3WalletPreSignApproveBCNetworkData {
	if o == nil || IsNil(o.BlockchainNetworkData.Get()) {
		var ret Web3WalletPreSignApproveBCNetworkData
		return ret
	}
	return *o.BlockchainNetworkData.Get()
}

// GetBlockchainNetworkDataOk returns a tuple with the BlockchainNetworkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSignApprove) GetBlockchainNetworkDataOk() (*Web3WalletPreSignApproveBCNetworkData, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockchainNetworkData.Get(), o.BlockchainNetworkData.IsSet()
}

// HasBlockchainNetworkData returns a boolean if a field has been set.
func (o *Web3WalletPreSignApprove) HasBlockchainNetworkData() bool {
	if o != nil && o.BlockchainNetworkData.IsSet() {
		return true
	}

	return false
}

// SetBlockchainNetworkData gets a reference to the given NullableWeb3WalletPreSignApproveBCNetworkData and assigns it to the BlockchainNetworkData field.
func (o *Web3WalletPreSignApprove) SetBlockchainNetworkData(v Web3WalletPreSignApproveBCNetworkData) {
	o.BlockchainNetworkData.Set(&v)
}
// SetBlockchainNetworkDataNil sets the value for BlockchainNetworkData to be an explicit nil
func (o *Web3WalletPreSignApprove) SetBlockchainNetworkDataNil() {
	o.BlockchainNetworkData.Set(nil)
}

// UnsetBlockchainNetworkData ensures that no value is present for BlockchainNetworkData, not even an explicit nil
func (o *Web3WalletPreSignApprove) UnsetBlockchainNetworkData() {
	o.BlockchainNetworkData.Unset()
}

// GetTransactionData returns the TransactionData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSignApprove) GetTransactionData() TXNData {
	if o == nil || IsNil(o.TransactionData.Get()) {
		var ret TXNData
		return ret
	}
	return *o.TransactionData.Get()
}

// GetTransactionDataOk returns a tuple with the TransactionData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSignApprove) GetTransactionDataOk() (*TXNData, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionData.Get(), o.TransactionData.IsSet()
}

// HasTransactionData returns a boolean if a field has been set.
func (o *Web3WalletPreSignApprove) HasTransactionData() bool {
	if o != nil && o.TransactionData.IsSet() {
		return true
	}

	return false
}

// SetTransactionData gets a reference to the given NullableTXNData and assigns it to the TransactionData field.
func (o *Web3WalletPreSignApprove) SetTransactionData(v TXNData) {
	o.TransactionData.Set(&v)
}
// SetTransactionDataNil sets the value for TransactionData to be an explicit nil
func (o *Web3WalletPreSignApprove) SetTransactionDataNil() {
	o.TransactionData.Set(nil)
}

// UnsetTransactionData ensures that no value is present for TransactionData, not even an explicit nil
func (o *Web3WalletPreSignApprove) UnsetTransactionData() {
	o.TransactionData.Unset()
}

func (o Web3WalletPreSignApprove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletPreSignApprove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockchainNetworkData.IsSet() {
		toSerialize["blockchain_network_data"] = o.BlockchainNetworkData.Get()
	}
	if o.TransactionData.IsSet() {
		toSerialize["transaction_data"] = o.TransactionData.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletPreSignApprove) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletPreSignApprove := _Web3WalletPreSignApprove{}

	err = json.Unmarshal(data, &varWeb3WalletPreSignApprove)

	if err != nil {
		return err
	}

	*o = Web3WalletPreSignApprove(varWeb3WalletPreSignApprove)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "blockchain_network_data")
		delete(additionalProperties, "transaction_data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletPreSignApprove struct {
	value *Web3WalletPreSignApprove
	isSet bool
}

func (v NullableWeb3WalletPreSignApprove) Get() *Web3WalletPreSignApprove {
	return v.value
}

func (v *NullableWeb3WalletPreSignApprove) Set(val *Web3WalletPreSignApprove) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletPreSignApprove) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletPreSignApprove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletPreSignApprove(val *Web3WalletPreSignApprove) *NullableWeb3WalletPreSignApprove {
	return &NullableWeb3WalletPreSignApprove{value: val, isSet: true}
}

func (v NullableWeb3WalletPreSignApprove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletPreSignApprove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


