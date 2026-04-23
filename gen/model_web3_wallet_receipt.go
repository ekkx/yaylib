
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletReceipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletReceipt{}

// Web3WalletReceipt struct for Web3WalletReceipt
type Web3WalletReceipt struct {
	BlockHash NullableString `json:"block_hash,omitempty"`
	BlockNumber NullableString `json:"block_number,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	CumulativeGasUsed NullableString `json:"cumulative_gas_used,omitempty"`
	GasUsed NullableString `json:"gas_used,omitempty"`
	Logs []Log `json:"logs,omitempty"`
	LogsBloom NullableString `json:"logs_bloom,omitempty"`
	Root NullableString `json:"root,omitempty"`
	Status NullableString `json:"status,omitempty"`
	TransactionHash NullableString `json:"transaction_hash,omitempty"`
	TransactionIndex NullableString `json:"transaction_index,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletReceipt Web3WalletReceipt

// NewWeb3WalletReceipt instantiates a new Web3WalletReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletReceipt() *Web3WalletReceipt {
	this := Web3WalletReceipt{}
	return &this
}

// NewWeb3WalletReceiptWithDefaults instantiates a new Web3WalletReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletReceiptWithDefaults() *Web3WalletReceipt {
	this := Web3WalletReceipt{}
	return &this
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash.Get()) {
		var ret string
		return ret
	}
	return *o.BlockHash.Get()
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockHash.Get(), o.BlockHash.IsSet()
}

// HasBlockHash returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasBlockHash() bool {
	if o != nil && o.BlockHash.IsSet() {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given NullableString and assigns it to the BlockHash field.
func (o *Web3WalletReceipt) SetBlockHash(v string) {
	o.BlockHash.Set(&v)
}
// SetBlockHashNil sets the value for BlockHash to be an explicit nil
func (o *Web3WalletReceipt) SetBlockHashNil() {
	o.BlockHash.Set(nil)
}

// UnsetBlockHash ensures that no value is present for BlockHash, not even an explicit nil
func (o *Web3WalletReceipt) UnsetBlockHash() {
	o.BlockHash.Unset()
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetBlockNumber() string {
	if o == nil || IsNil(o.BlockNumber.Get()) {
		var ret string
		return ret
	}
	return *o.BlockNumber.Get()
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetBlockNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockNumber.Get(), o.BlockNumber.IsSet()
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasBlockNumber() bool {
	if o != nil && o.BlockNumber.IsSet() {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given NullableString and assigns it to the BlockNumber field.
func (o *Web3WalletReceipt) SetBlockNumber(v string) {
	o.BlockNumber.Set(&v)
}
// SetBlockNumberNil sets the value for BlockNumber to be an explicit nil
func (o *Web3WalletReceipt) SetBlockNumberNil() {
	o.BlockNumber.Set(nil)
}

// UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil
func (o *Web3WalletReceipt) UnsetBlockNumber() {
	o.BlockNumber.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *Web3WalletReceipt) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *Web3WalletReceipt) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *Web3WalletReceipt) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetCumulativeGasUsed returns the CumulativeGasUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetCumulativeGasUsed() string {
	if o == nil || IsNil(o.CumulativeGasUsed.Get()) {
		var ret string
		return ret
	}
	return *o.CumulativeGasUsed.Get()
}

// GetCumulativeGasUsedOk returns a tuple with the CumulativeGasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetCumulativeGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CumulativeGasUsed.Get(), o.CumulativeGasUsed.IsSet()
}

// HasCumulativeGasUsed returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasCumulativeGasUsed() bool {
	if o != nil && o.CumulativeGasUsed.IsSet() {
		return true
	}

	return false
}

// SetCumulativeGasUsed gets a reference to the given NullableString and assigns it to the CumulativeGasUsed field.
func (o *Web3WalletReceipt) SetCumulativeGasUsed(v string) {
	o.CumulativeGasUsed.Set(&v)
}
// SetCumulativeGasUsedNil sets the value for CumulativeGasUsed to be an explicit nil
func (o *Web3WalletReceipt) SetCumulativeGasUsedNil() {
	o.CumulativeGasUsed.Set(nil)
}

// UnsetCumulativeGasUsed ensures that no value is present for CumulativeGasUsed, not even an explicit nil
func (o *Web3WalletReceipt) UnsetCumulativeGasUsed() {
	o.CumulativeGasUsed.Unset()
}

// GetGasUsed returns the GasUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetGasUsed() string {
	if o == nil || IsNil(o.GasUsed.Get()) {
		var ret string
		return ret
	}
	return *o.GasUsed.Get()
}

// GetGasUsedOk returns a tuple with the GasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetGasUsedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasUsed.Get(), o.GasUsed.IsSet()
}

// HasGasUsed returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasGasUsed() bool {
	if o != nil && o.GasUsed.IsSet() {
		return true
	}

	return false
}

// SetGasUsed gets a reference to the given NullableString and assigns it to the GasUsed field.
func (o *Web3WalletReceipt) SetGasUsed(v string) {
	o.GasUsed.Set(&v)
}
// SetGasUsedNil sets the value for GasUsed to be an explicit nil
func (o *Web3WalletReceipt) SetGasUsedNil() {
	o.GasUsed.Set(nil)
}

// UnsetGasUsed ensures that no value is present for GasUsed, not even an explicit nil
func (o *Web3WalletReceipt) UnsetGasUsed() {
	o.GasUsed.Unset()
}

// GetLogs returns the Logs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetLogs() []Log {
	if o == nil {
		var ret []Log
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetLogsOk() ([]Log, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []Log and assigns it to the Logs field.
func (o *Web3WalletReceipt) SetLogs(v []Log) {
	o.Logs = v
}

// GetLogsBloom returns the LogsBloom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetLogsBloom() string {
	if o == nil || IsNil(o.LogsBloom.Get()) {
		var ret string
		return ret
	}
	return *o.LogsBloom.Get()
}

// GetLogsBloomOk returns a tuple with the LogsBloom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetLogsBloomOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsBloom.Get(), o.LogsBloom.IsSet()
}

// HasLogsBloom returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasLogsBloom() bool {
	if o != nil && o.LogsBloom.IsSet() {
		return true
	}

	return false
}

// SetLogsBloom gets a reference to the given NullableString and assigns it to the LogsBloom field.
func (o *Web3WalletReceipt) SetLogsBloom(v string) {
	o.LogsBloom.Set(&v)
}
// SetLogsBloomNil sets the value for LogsBloom to be an explicit nil
func (o *Web3WalletReceipt) SetLogsBloomNil() {
	o.LogsBloom.Set(nil)
}

// UnsetLogsBloom ensures that no value is present for LogsBloom, not even an explicit nil
func (o *Web3WalletReceipt) UnsetLogsBloom() {
	o.LogsBloom.Unset()
}

// GetRoot returns the Root field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetRoot() string {
	if o == nil || IsNil(o.Root.Get()) {
		var ret string
		return ret
	}
	return *o.Root.Get()
}

// GetRootOk returns a tuple with the Root field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Root.Get(), o.Root.IsSet()
}

// HasRoot returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasRoot() bool {
	if o != nil && o.Root.IsSet() {
		return true
	}

	return false
}

// SetRoot gets a reference to the given NullableString and assigns it to the Root field.
func (o *Web3WalletReceipt) SetRoot(v string) {
	o.Root.Set(&v)
}
// SetRootNil sets the value for Root to be an explicit nil
func (o *Web3WalletReceipt) SetRootNil() {
	o.Root.Set(nil)
}

// UnsetRoot ensures that no value is present for Root, not even an explicit nil
func (o *Web3WalletReceipt) UnsetRoot() {
	o.Root.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Web3WalletReceipt) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Web3WalletReceipt) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Web3WalletReceipt) UnsetStatus() {
	o.Status.Unset()
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionHash.Get()
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionHash.Get(), o.TransactionHash.IsSet()
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasTransactionHash() bool {
	if o != nil && o.TransactionHash.IsSet() {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given NullableString and assigns it to the TransactionHash field.
func (o *Web3WalletReceipt) SetTransactionHash(v string) {
	o.TransactionHash.Set(&v)
}
// SetTransactionHashNil sets the value for TransactionHash to be an explicit nil
func (o *Web3WalletReceipt) SetTransactionHashNil() {
	o.TransactionHash.Set(nil)
}

// UnsetTransactionHash ensures that no value is present for TransactionHash, not even an explicit nil
func (o *Web3WalletReceipt) UnsetTransactionHash() {
	o.TransactionHash.Unset()
}

// GetTransactionIndex returns the TransactionIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletReceipt) GetTransactionIndex() string {
	if o == nil || IsNil(o.TransactionIndex.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionIndex.Get()
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletReceipt) GetTransactionIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionIndex.Get(), o.TransactionIndex.IsSet()
}

// HasTransactionIndex returns a boolean if a field has been set.
func (o *Web3WalletReceipt) HasTransactionIndex() bool {
	if o != nil && o.TransactionIndex.IsSet() {
		return true
	}

	return false
}

// SetTransactionIndex gets a reference to the given NullableString and assigns it to the TransactionIndex field.
func (o *Web3WalletReceipt) SetTransactionIndex(v string) {
	o.TransactionIndex.Set(&v)
}
// SetTransactionIndexNil sets the value for TransactionIndex to be an explicit nil
func (o *Web3WalletReceipt) SetTransactionIndexNil() {
	o.TransactionIndex.Set(nil)
}

// UnsetTransactionIndex ensures that no value is present for TransactionIndex, not even an explicit nil
func (o *Web3WalletReceipt) UnsetTransactionIndex() {
	o.TransactionIndex.Unset()
}

func (o Web3WalletReceipt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletReceipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockHash.IsSet() {
		toSerialize["block_hash"] = o.BlockHash.Get()
	}
	if o.BlockNumber.IsSet() {
		toSerialize["block_number"] = o.BlockNumber.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.CumulativeGasUsed.IsSet() {
		toSerialize["cumulative_gas_used"] = o.CumulativeGasUsed.Get()
	}
	if o.GasUsed.IsSet() {
		toSerialize["gas_used"] = o.GasUsed.Get()
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	if o.LogsBloom.IsSet() {
		toSerialize["logs_bloom"] = o.LogsBloom.Get()
	}
	if o.Root.IsSet() {
		toSerialize["root"] = o.Root.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.TransactionHash.IsSet() {
		toSerialize["transaction_hash"] = o.TransactionHash.Get()
	}
	if o.TransactionIndex.IsSet() {
		toSerialize["transaction_index"] = o.TransactionIndex.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletReceipt) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletReceipt := _Web3WalletReceipt{}

	err = json.Unmarshal(data, &varWeb3WalletReceipt)

	if err != nil {
		return err
	}

	*o = Web3WalletReceipt(varWeb3WalletReceipt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "block_hash")
		delete(additionalProperties, "block_number")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "cumulative_gas_used")
		delete(additionalProperties, "gas_used")
		delete(additionalProperties, "logs")
		delete(additionalProperties, "logs_bloom")
		delete(additionalProperties, "root")
		delete(additionalProperties, "status")
		delete(additionalProperties, "transaction_hash")
		delete(additionalProperties, "transaction_index")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletReceipt struct {
	value *Web3WalletReceipt
	isSet bool
}

func (v NullableWeb3WalletReceipt) Get() *Web3WalletReceipt {
	return v.value
}

func (v *NullableWeb3WalletReceipt) Set(val *Web3WalletReceipt) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletReceipt(val *Web3WalletReceipt) *NullableWeb3WalletReceipt {
	return &NullableWeb3WalletReceipt{value: val, isSet: true}
}

func (v NullableWeb3WalletReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


