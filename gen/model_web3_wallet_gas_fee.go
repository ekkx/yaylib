
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletGasFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletGasFee{}

// Web3WalletGasFee struct for Web3WalletGasFee
type Web3WalletGasFee struct {
	GasLimit NullableInt64 `json:"gas_limit,omitempty"`
	GasPriceWei NullableFloat64 `json:"gas_price_wei,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletGasFee Web3WalletGasFee

// NewWeb3WalletGasFee instantiates a new Web3WalletGasFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletGasFee() *Web3WalletGasFee {
	this := Web3WalletGasFee{}
	return &this
}

// NewWeb3WalletGasFeeWithDefaults instantiates a new Web3WalletGasFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletGasFeeWithDefaults() *Web3WalletGasFee {
	this := Web3WalletGasFee{}
	return &this
}

// GetGasLimit returns the GasLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletGasFee) GetGasLimit() int64 {
	if o == nil || IsNil(o.GasLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.GasLimit.Get()
}

// GetGasLimitOk returns a tuple with the GasLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletGasFee) GetGasLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasLimit.Get(), o.GasLimit.IsSet()
}

// HasGasLimit returns a boolean if a field has been set.
func (o *Web3WalletGasFee) HasGasLimit() bool {
	if o != nil && o.GasLimit.IsSet() {
		return true
	}

	return false
}

// SetGasLimit gets a reference to the given NullableInt64 and assigns it to the GasLimit field.
func (o *Web3WalletGasFee) SetGasLimit(v int64) {
	o.GasLimit.Set(&v)
}
// SetGasLimitNil sets the value for GasLimit to be an explicit nil
func (o *Web3WalletGasFee) SetGasLimitNil() {
	o.GasLimit.Set(nil)
}

// UnsetGasLimit ensures that no value is present for GasLimit, not even an explicit nil
func (o *Web3WalletGasFee) UnsetGasLimit() {
	o.GasLimit.Unset()
}

// GetGasPriceWei returns the GasPriceWei field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletGasFee) GetGasPriceWei() float64 {
	if o == nil || IsNil(o.GasPriceWei.Get()) {
		var ret float64
		return ret
	}
	return *o.GasPriceWei.Get()
}

// GetGasPriceWeiOk returns a tuple with the GasPriceWei field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletGasFee) GetGasPriceWeiOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasPriceWei.Get(), o.GasPriceWei.IsSet()
}

// HasGasPriceWei returns a boolean if a field has been set.
func (o *Web3WalletGasFee) HasGasPriceWei() bool {
	if o != nil && o.GasPriceWei.IsSet() {
		return true
	}

	return false
}

// SetGasPriceWei gets a reference to the given NullableFloat64 and assigns it to the GasPriceWei field.
func (o *Web3WalletGasFee) SetGasPriceWei(v float64) {
	o.GasPriceWei.Set(&v)
}
// SetGasPriceWeiNil sets the value for GasPriceWei to be an explicit nil
func (o *Web3WalletGasFee) SetGasPriceWeiNil() {
	o.GasPriceWei.Set(nil)
}

// UnsetGasPriceWei ensures that no value is present for GasPriceWei, not even an explicit nil
func (o *Web3WalletGasFee) UnsetGasPriceWei() {
	o.GasPriceWei.Unset()
}

func (o Web3WalletGasFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletGasFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GasLimit.IsSet() {
		toSerialize["gas_limit"] = o.GasLimit.Get()
	}
	if o.GasPriceWei.IsSet() {
		toSerialize["gas_price_wei"] = o.GasPriceWei.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletGasFee) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletGasFee := _Web3WalletGasFee{}

	err = json.Unmarshal(data, &varWeb3WalletGasFee)

	if err != nil {
		return err
	}

	*o = Web3WalletGasFee(varWeb3WalletGasFee)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gas_limit")
		delete(additionalProperties, "gas_price_wei")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletGasFee struct {
	value *Web3WalletGasFee
	isSet bool
}

func (v NullableWeb3WalletGasFee) Get() *Web3WalletGasFee {
	return v.value
}

func (v *NullableWeb3WalletGasFee) Set(val *Web3WalletGasFee) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletGasFee) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletGasFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletGasFee(val *Web3WalletGasFee) *NullableWeb3WalletGasFee {
	return &NullableWeb3WalletGasFee{value: val, isSet: true}
}

func (v NullableWeb3WalletGasFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletGasFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


