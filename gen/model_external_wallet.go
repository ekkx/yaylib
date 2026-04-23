
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ExternalWallet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalWallet{}

// ExternalWallet struct for ExternalWallet
type ExternalWallet struct {
	Result NullableWeb3WalletExternalWallet `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExternalWallet ExternalWallet

// NewExternalWallet instantiates a new ExternalWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalWallet() *ExternalWallet {
	this := ExternalWallet{}
	return &this
}

// NewExternalWalletWithDefaults instantiates a new ExternalWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalWalletWithDefaults() *ExternalWallet {
	this := ExternalWallet{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalWallet) GetResult() Web3WalletExternalWallet {
	if o == nil || IsNil(o.Result.Get()) {
		var ret Web3WalletExternalWallet
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalWallet) GetResultOk() (*Web3WalletExternalWallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *ExternalWallet) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableWeb3WalletExternalWallet and assigns it to the Result field.
func (o *ExternalWallet) SetResult(v Web3WalletExternalWallet) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *ExternalWallet) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *ExternalWallet) UnsetResult() {
	o.Result.Unset()
}

func (o ExternalWallet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalWallet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalWallet) UnmarshalJSON(data []byte) (err error) {
	varExternalWallet := _ExternalWallet{}

	err = json.Unmarshal(data, &varExternalWallet)

	if err != nil {
		return err
	}

	*o = ExternalWallet(varExternalWallet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalWallet struct {
	value *ExternalWallet
	isSet bool
}

func (v NullableExternalWallet) Get() *ExternalWallet {
	return v.value
}

func (v *NullableExternalWallet) Set(val *ExternalWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalWallet(val *ExternalWallet) *NullableExternalWallet {
	return &NullableExternalWallet{value: val, isSet: true}
}

func (v NullableExternalWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


