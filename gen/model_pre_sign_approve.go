
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PreSignApprove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreSignApprove{}

// PreSignApprove struct for PreSignApprove
type PreSignApprove struct {
	Result NullableWeb3WalletPreSignApprove `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreSignApprove PreSignApprove

// NewPreSignApprove instantiates a new PreSignApprove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreSignApprove() *PreSignApprove {
	this := PreSignApprove{}
	return &this
}

// NewPreSignApproveWithDefaults instantiates a new PreSignApprove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreSignApproveWithDefaults() *PreSignApprove {
	this := PreSignApprove{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PreSignApprove) GetResult() Web3WalletPreSignApprove {
	if o == nil || IsNil(o.Result.Get()) {
		var ret Web3WalletPreSignApprove
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreSignApprove) GetResultOk() (*Web3WalletPreSignApprove, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *PreSignApprove) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableWeb3WalletPreSignApprove and assigns it to the Result field.
func (o *PreSignApprove) SetResult(v Web3WalletPreSignApprove) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *PreSignApprove) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *PreSignApprove) UnsetResult() {
	o.Result.Unset()
}

func (o PreSignApprove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreSignApprove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreSignApprove) UnmarshalJSON(data []byte) (err error) {
	varPreSignApprove := _PreSignApprove{}

	err = json.Unmarshal(data, &varPreSignApprove)

	if err != nil {
		return err
	}

	*o = PreSignApprove(varPreSignApprove)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreSignApprove struct {
	value *PreSignApprove
	isSet bool
}

func (v NullablePreSignApprove) Get() *PreSignApprove {
	return v.value
}

func (v *NullablePreSignApprove) Set(val *PreSignApprove) {
	v.value = val
	v.isSet = true
}

func (v NullablePreSignApprove) IsSet() bool {
	return v.isSet
}

func (v *NullablePreSignApprove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreSignApprove(val *PreSignApprove) *NullablePreSignApprove {
	return &NullablePreSignApprove{value: val, isSet: true}
}

func (v NullablePreSignApprove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreSignApprove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


