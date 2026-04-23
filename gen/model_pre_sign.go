
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PreSign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreSign{}

// PreSign struct for PreSign
type PreSign struct {
	Result NullableWeb3WalletPreSign `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreSign PreSign

// NewPreSign instantiates a new PreSign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreSign() *PreSign {
	this := PreSign{}
	return &this
}

// NewPreSignWithDefaults instantiates a new PreSign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreSignWithDefaults() *PreSign {
	this := PreSign{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreSign) GetResult() Web3WalletPreSign {
	if o == nil || IsNil(o.Result.Get()) {
		var ret Web3WalletPreSign
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreSign) GetResultOk() (*Web3WalletPreSign, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *PreSign) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableWeb3WalletPreSign and assigns it to the Result field.
func (o *PreSign) SetResult(v Web3WalletPreSign) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *PreSign) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *PreSign) UnsetResult() {
	o.Result.Unset()
}

func (o PreSign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreSign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreSign) UnmarshalJSON(data []byte) (err error) {
	varPreSign := _PreSign{}

	err = json.Unmarshal(data, &varPreSign)

	if err != nil {
		return err
	}

	*o = PreSign(varPreSign)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreSign struct {
	value *PreSign
	isSet bool
}

func (v NullablePreSign) Get() *PreSign {
	return v.value
}

func (v *NullablePreSign) Set(val *PreSign) {
	v.value = val
	v.isSet = true
}

func (v NullablePreSign) IsSet() bool {
	return v.isSet
}

func (v *NullablePreSign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreSign(val *PreSign) *NullablePreSign {
	return &NullablePreSign{value: val, isSet: true}
}

func (v NullablePreSign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreSign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


