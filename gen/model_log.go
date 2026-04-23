
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Log type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Log{}

// Log struct for Log
type Log struct {
	Address NullableString `json:"address,omitempty"`
	BlockHash NullableString `json:"block_hash,omitempty"`
	BlockNumber NullableString `json:"block_number,omitempty"`
	Data NullableString `json:"data,omitempty"`
	LogIndex NullableString `json:"log_index,omitempty"`
	Removed NullableBool `json:"removed,omitempty"`
	Topics []string `json:"topics,omitempty"`
	TransactionHash NullableString `json:"transaction_hash,omitempty"`
	TransactionIndex NullableString `json:"transaction_index,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Log Log

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog() *Log {
	this := Log{}
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *Log) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *Log) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *Log) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *Log) UnsetAddress() {
	o.Address.Unset()
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash.Get()) {
		var ret string
		return ret
	}
	return *o.BlockHash.Get()
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockHash.Get(), o.BlockHash.IsSet()
}

// HasBlockHash returns a boolean if a field has been set.
func (o *Log) HasBlockHash() bool {
	if o != nil && o.BlockHash.IsSet() {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given NullableString and assigns it to the BlockHash field.
func (o *Log) SetBlockHash(v string) {
	o.BlockHash.Set(&v)
}
// SetBlockHashNil sets the value for BlockHash to be an explicit nil
func (o *Log) SetBlockHashNil() {
	o.BlockHash.Set(nil)
}

// UnsetBlockHash ensures that no value is present for BlockHash, not even an explicit nil
func (o *Log) UnsetBlockHash() {
	o.BlockHash.Unset()
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetBlockNumber() string {
	if o == nil || IsNil(o.BlockNumber.Get()) {
		var ret string
		return ret
	}
	return *o.BlockNumber.Get()
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetBlockNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockNumber.Get(), o.BlockNumber.IsSet()
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *Log) HasBlockNumber() bool {
	if o != nil && o.BlockNumber.IsSet() {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given NullableString and assigns it to the BlockNumber field.
func (o *Log) SetBlockNumber(v string) {
	o.BlockNumber.Set(&v)
}
// SetBlockNumberNil sets the value for BlockNumber to be an explicit nil
func (o *Log) SetBlockNumberNil() {
	o.BlockNumber.Set(nil)
}

// UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil
func (o *Log) UnsetBlockNumber() {
	o.BlockNumber.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *Log) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *Log) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *Log) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *Log) UnsetData() {
	o.Data.Unset()
}

// GetLogIndex returns the LogIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetLogIndex() string {
	if o == nil || IsNil(o.LogIndex.Get()) {
		var ret string
		return ret
	}
	return *o.LogIndex.Get()
}

// GetLogIndexOk returns a tuple with the LogIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetLogIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogIndex.Get(), o.LogIndex.IsSet()
}

// HasLogIndex returns a boolean if a field has been set.
func (o *Log) HasLogIndex() bool {
	if o != nil && o.LogIndex.IsSet() {
		return true
	}

	return false
}

// SetLogIndex gets a reference to the given NullableString and assigns it to the LogIndex field.
func (o *Log) SetLogIndex(v string) {
	o.LogIndex.Set(&v)
}
// SetLogIndexNil sets the value for LogIndex to be an explicit nil
func (o *Log) SetLogIndexNil() {
	o.LogIndex.Set(nil)
}

// UnsetLogIndex ensures that no value is present for LogIndex, not even an explicit nil
func (o *Log) UnsetLogIndex() {
	o.LogIndex.Unset()
}

// GetRemoved returns the Removed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetRemoved() bool {
	if o == nil || IsNil(o.Removed.Get()) {
		var ret bool
		return ret
	}
	return *o.Removed.Get()
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetRemovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Removed.Get(), o.Removed.IsSet()
}

// HasRemoved returns a boolean if a field has been set.
func (o *Log) HasRemoved() bool {
	if o != nil && o.Removed.IsSet() {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given NullableBool and assigns it to the Removed field.
func (o *Log) SetRemoved(v bool) {
	o.Removed.Set(&v)
}
// SetRemovedNil sets the value for Removed to be an explicit nil
func (o *Log) SetRemovedNil() {
	o.Removed.Set(nil)
}

// UnsetRemoved ensures that no value is present for Removed, not even an explicit nil
func (o *Log) UnsetRemoved() {
	o.Removed.Unset()
}

// GetTopics returns the Topics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetTopics() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetTopicsOk() ([]string, bool) {
	if o == nil || IsNil(o.Topics) {
		return nil, false
	}
	return o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *Log) HasTopics() bool {
	if o != nil && !IsNil(o.Topics) {
		return true
	}

	return false
}

// SetTopics gets a reference to the given []string and assigns it to the Topics field.
func (o *Log) SetTopics(v []string) {
	o.Topics = v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionHash.Get()
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionHash.Get(), o.TransactionHash.IsSet()
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *Log) HasTransactionHash() bool {
	if o != nil && o.TransactionHash.IsSet() {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given NullableString and assigns it to the TransactionHash field.
func (o *Log) SetTransactionHash(v string) {
	o.TransactionHash.Set(&v)
}
// SetTransactionHashNil sets the value for TransactionHash to be an explicit nil
func (o *Log) SetTransactionHashNil() {
	o.TransactionHash.Set(nil)
}

// UnsetTransactionHash ensures that no value is present for TransactionHash, not even an explicit nil
func (o *Log) UnsetTransactionHash() {
	o.TransactionHash.Unset()
}

// GetTransactionIndex returns the TransactionIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Log) GetTransactionIndex() string {
	if o == nil || IsNil(o.TransactionIndex.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionIndex.Get()
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Log) GetTransactionIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionIndex.Get(), o.TransactionIndex.IsSet()
}

// HasTransactionIndex returns a boolean if a field has been set.
func (o *Log) HasTransactionIndex() bool {
	if o != nil && o.TransactionIndex.IsSet() {
		return true
	}

	return false
}

// SetTransactionIndex gets a reference to the given NullableString and assigns it to the TransactionIndex field.
func (o *Log) SetTransactionIndex(v string) {
	o.TransactionIndex.Set(&v)
}
// SetTransactionIndexNil sets the value for TransactionIndex to be an explicit nil
func (o *Log) SetTransactionIndexNil() {
	o.TransactionIndex.Set(nil)
}

// UnsetTransactionIndex ensures that no value is present for TransactionIndex, not even an explicit nil
func (o *Log) UnsetTransactionIndex() {
	o.TransactionIndex.Unset()
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Log) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.BlockHash.IsSet() {
		toSerialize["block_hash"] = o.BlockHash.Get()
	}
	if o.BlockNumber.IsSet() {
		toSerialize["block_number"] = o.BlockNumber.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.LogIndex.IsSet() {
		toSerialize["log_index"] = o.LogIndex.Get()
	}
	if o.Removed.IsSet() {
		toSerialize["removed"] = o.Removed.Get()
	}
	if o.Topics != nil {
		toSerialize["topics"] = o.Topics
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

func (o *Log) UnmarshalJSON(data []byte) (err error) {
	varLog := _Log{}

	err = json.Unmarshal(data, &varLog)

	if err != nil {
		return err
	}

	*o = Log(varLog)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "block_hash")
		delete(additionalProperties, "block_number")
		delete(additionalProperties, "data")
		delete(additionalProperties, "log_index")
		delete(additionalProperties, "removed")
		delete(additionalProperties, "topics")
		delete(additionalProperties, "transaction_hash")
		delete(additionalProperties, "transaction_index")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


