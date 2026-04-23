
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletPreSignApproveBCNetworkData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletPreSignApproveBCNetworkData{}

// Web3WalletPreSignApproveBCNetworkData struct for Web3WalletPreSignApproveBCNetworkData
type Web3WalletPreSignApproveBCNetworkData struct {
	AdditionGasPercent NullableWeb3WalletGasPercent `json:"addition_gas_percent,omitempty"`
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ChainUrl NullableString `json:"chain_url,omitempty"`
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletPreSignApproveBCNetworkData Web3WalletPreSignApproveBCNetworkData

// NewWeb3WalletPreSignApproveBCNetworkData instantiates a new Web3WalletPreSignApproveBCNetworkData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletPreSignApproveBCNetworkData() *Web3WalletPreSignApproveBCNetworkData {
	this := Web3WalletPreSignApproveBCNetworkData{}
	return &this
}

// NewWeb3WalletPreSignApproveBCNetworkDataWithDefaults instantiates a new Web3WalletPreSignApproveBCNetworkData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletPreSignApproveBCNetworkDataWithDefaults() *Web3WalletPreSignApproveBCNetworkData {
	this := Web3WalletPreSignApproveBCNetworkData{}
	return &this
}

// GetAdditionGasPercent returns the AdditionGasPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSignApproveBCNetworkData) GetAdditionGasPercent() Web3WalletGasPercent {
	if o == nil || IsNil(o.AdditionGasPercent.Get()) {
		var ret Web3WalletGasPercent
		return ret
	}
	return *o.AdditionGasPercent.Get()
}

// GetAdditionGasPercentOk returns a tuple with the AdditionGasPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSignApproveBCNetworkData) GetAdditionGasPercentOk() (*Web3WalletGasPercent, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionGasPercent.Get(), o.AdditionGasPercent.IsSet()
}

// HasAdditionGasPercent returns a boolean if a field has been set.
func (o *Web3WalletPreSignApproveBCNetworkData) HasAdditionGasPercent() bool {
	if o != nil && o.AdditionGasPercent.IsSet() {
		return true
	}

	return false
}

// SetAdditionGasPercent gets a reference to the given NullableWeb3WalletGasPercent and assigns it to the AdditionGasPercent field.
func (o *Web3WalletPreSignApproveBCNetworkData) SetAdditionGasPercent(v Web3WalletGasPercent) {
	o.AdditionGasPercent.Set(&v)
}
// SetAdditionGasPercentNil sets the value for AdditionGasPercent to be an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) SetAdditionGasPercentNil() {
	o.AdditionGasPercent.Set(nil)
}

// UnsetAdditionGasPercent ensures that no value is present for AdditionGasPercent, not even an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) UnsetAdditionGasPercent() {
	o.AdditionGasPercent.Unset()
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSignApproveBCNetworkData) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSignApproveBCNetworkData) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3WalletPreSignApproveBCNetworkData) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Web3WalletPreSignApproveBCNetworkData) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) UnsetChainId() {
	o.ChainId.Unset()
}

// GetChainUrl returns the ChainUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSignApproveBCNetworkData) GetChainUrl() string {
	if o == nil || IsNil(o.ChainUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ChainUrl.Get()
}

// GetChainUrlOk returns a tuple with the ChainUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSignApproveBCNetworkData) GetChainUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainUrl.Get(), o.ChainUrl.IsSet()
}

// HasChainUrl returns a boolean if a field has been set.
func (o *Web3WalletPreSignApproveBCNetworkData) HasChainUrl() bool {
	if o != nil && o.ChainUrl.IsSet() {
		return true
	}

	return false
}

// SetChainUrl gets a reference to the given NullableString and assigns it to the ChainUrl field.
func (o *Web3WalletPreSignApproveBCNetworkData) SetChainUrl(v string) {
	o.ChainUrl.Set(&v)
}
// SetChainUrlNil sets the value for ChainUrl to be an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) SetChainUrlNil() {
	o.ChainUrl.Set(nil)
}

// UnsetChainUrl ensures that no value is present for ChainUrl, not even an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) UnsetChainUrl() {
	o.ChainUrl.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletPreSignApproveBCNetworkData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletPreSignApproveBCNetworkData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Web3WalletPreSignApproveBCNetworkData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Web3WalletPreSignApproveBCNetworkData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Web3WalletPreSignApproveBCNetworkData) UnsetName() {
	o.Name.Unset()
}

func (o Web3WalletPreSignApproveBCNetworkData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletPreSignApproveBCNetworkData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionGasPercent.IsSet() {
		toSerialize["addition_gas_percent"] = o.AdditionGasPercent.Get()
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ChainUrl.IsSet() {
		toSerialize["chain_url"] = o.ChainUrl.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletPreSignApproveBCNetworkData) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletPreSignApproveBCNetworkData := _Web3WalletPreSignApproveBCNetworkData{}

	err = json.Unmarshal(data, &varWeb3WalletPreSignApproveBCNetworkData)

	if err != nil {
		return err
	}

	*o = Web3WalletPreSignApproveBCNetworkData(varWeb3WalletPreSignApproveBCNetworkData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addition_gas_percent")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "chain_url")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletPreSignApproveBCNetworkData struct {
	value *Web3WalletPreSignApproveBCNetworkData
	isSet bool
}

func (v NullableWeb3WalletPreSignApproveBCNetworkData) Get() *Web3WalletPreSignApproveBCNetworkData {
	return v.value
}

func (v *NullableWeb3WalletPreSignApproveBCNetworkData) Set(val *Web3WalletPreSignApproveBCNetworkData) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletPreSignApproveBCNetworkData) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletPreSignApproveBCNetworkData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletPreSignApproveBCNetworkData(val *Web3WalletPreSignApproveBCNetworkData) *NullableWeb3WalletPreSignApproveBCNetworkData {
	return &NullableWeb3WalletPreSignApproveBCNetworkData{value: val, isSet: true}
}

func (v NullableWeb3WalletPreSignApproveBCNetworkData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletPreSignApproveBCNetworkData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


