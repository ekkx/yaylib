
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletPreSign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletPreSign{}

// Web3WalletPreSign struct for Web3WalletPreSign
type Web3WalletPreSign struct {
	BlockchainNetworkData NullableBCNetworkData `json:"blockchain_network_data,omitempty"`
	TransactionAndLogData NullableTXNLogData `json:"transaction_and_log_data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletPreSign Web3WalletPreSign

// NewWeb3WalletPreSign instantiates a new Web3WalletPreSign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletPreSign() *Web3WalletPreSign {
	this := Web3WalletPreSign{}
	return &this
}

// NewWeb3WalletPreSignWithDefaults instantiates a new Web3WalletPreSign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletPreSignWithDefaults() *Web3WalletPreSign {
	this := Web3WalletPreSign{}
	return &this
}

// GetBlockchainNetworkData returns the BlockchainNetworkData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSign) GetBlockchainNetworkData() BCNetworkData {
	if o == nil || IsNil(o.BlockchainNetworkData.Get()) {
		var ret BCNetworkData
		return ret
	}
	return *o.BlockchainNetworkData.Get()
}

// GetBlockchainNetworkDataOk returns a tuple with the BlockchainNetworkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSign) GetBlockchainNetworkDataOk() (*BCNetworkData, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockchainNetworkData.Get(), o.BlockchainNetworkData.IsSet()
}

// HasBlockchainNetworkData returns a boolean if a field has been set.
func (o *Web3WalletPreSign) HasBlockchainNetworkData() bool {
	if o != nil && o.BlockchainNetworkData.IsSet() {
		return true
	}

	return false
}

// SetBlockchainNetworkData gets a reference to the given NullableBCNetworkData and assigns it to the BlockchainNetworkData field.
func (o *Web3WalletPreSign) SetBlockchainNetworkData(v BCNetworkData) {
	o.BlockchainNetworkData.Set(&v)
}
// SetBlockchainNetworkDataNil sets the value for BlockchainNetworkData to be an explicit nil
func (o *Web3WalletPreSign) SetBlockchainNetworkDataNil() {
	o.BlockchainNetworkData.Set(nil)
}

// UnsetBlockchainNetworkData ensures that no value is present for BlockchainNetworkData, not even an explicit nil
func (o *Web3WalletPreSign) UnsetBlockchainNetworkData() {
	o.BlockchainNetworkData.Unset()
}

// GetTransactionAndLogData returns the TransactionAndLogData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSign) GetTransactionAndLogData() TXNLogData {
	if o == nil || IsNil(o.TransactionAndLogData.Get()) {
		var ret TXNLogData
		return ret
	}
	return *o.TransactionAndLogData.Get()
}

// GetTransactionAndLogDataOk returns a tuple with the TransactionAndLogData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSign) GetTransactionAndLogDataOk() (*TXNLogData, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionAndLogData.Get(), o.TransactionAndLogData.IsSet()
}

// HasTransactionAndLogData returns a boolean if a field has been set.
func (o *Web3WalletPreSign) HasTransactionAndLogData() bool {
	if o != nil && o.TransactionAndLogData.IsSet() {
		return true
	}

	return false
}

// SetTransactionAndLogData gets a reference to the given NullableTXNLogData and assigns it to the TransactionAndLogData field.
func (o *Web3WalletPreSign) SetTransactionAndLogData(v TXNLogData) {
	o.TransactionAndLogData.Set(&v)
}
// SetTransactionAndLogDataNil sets the value for TransactionAndLogData to be an explicit nil
func (o *Web3WalletPreSign) SetTransactionAndLogDataNil() {
	o.TransactionAndLogData.Set(nil)
}

// UnsetTransactionAndLogData ensures that no value is present for TransactionAndLogData, not even an explicit nil
func (o *Web3WalletPreSign) UnsetTransactionAndLogData() {
	o.TransactionAndLogData.Unset()
}

func (o Web3WalletPreSign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletPreSign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockchainNetworkData.IsSet() {
		toSerialize["blockchain_network_data"] = o.BlockchainNetworkData.Get()
	}
	if o.TransactionAndLogData.IsSet() {
		toSerialize["transaction_and_log_data"] = o.TransactionAndLogData.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletPreSign) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletPreSign := _Web3WalletPreSign{}

	err = json.Unmarshal(data, &varWeb3WalletPreSign)

	if err != nil {
		return err
	}

	*o = Web3WalletPreSign(varWeb3WalletPreSign)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "blockchain_network_data")
		delete(additionalProperties, "transaction_and_log_data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletPreSign struct {
	value *Web3WalletPreSign
	isSet bool
}

func (v NullableWeb3WalletPreSign) Get() *Web3WalletPreSign {
	return v.value
}

func (v *NullableWeb3WalletPreSign) Set(val *Web3WalletPreSign) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletPreSign) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletPreSign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletPreSign(val *Web3WalletPreSign) *NullableWeb3WalletPreSign {
	return &NullableWeb3WalletPreSign{value: val, isSet: true}
}

func (v NullableWeb3WalletPreSign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletPreSign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


