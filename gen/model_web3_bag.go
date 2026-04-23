
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3Bag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3Bag{}

// Web3Bag struct for Web3Bag
type Web3Bag struct {
	// Unresolved Java type: jp.nanameue.yay.shared.data.model.empl.Empl
	Empl map[string]interface{} `json:"empl,omitempty"`
	PalsCount NullableInt32 `json:"pals_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3Bag Web3Bag

// NewWeb3Bag instantiates a new Web3Bag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3Bag() *Web3Bag {
	this := Web3Bag{}
	return &this
}

// NewWeb3BagWithDefaults instantiates a new Web3Bag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3BagWithDefaults() *Web3Bag {
	this := Web3Bag{}
	return &this
}

// GetEmpl returns the Empl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3Bag) GetEmpl() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Empl
}

// GetEmplOk returns a tuple with the Empl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3Bag) GetEmplOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Empl) {
		return map[string]interface{}{}, false
	}
	return o.Empl, true
}

// HasEmpl returns a boolean if a field has been set.
func (o *Web3Bag) HasEmpl() bool {
	if o != nil && !IsNil(o.Empl) {
		return true
	}

	return false
}

// SetEmpl gets a reference to the given map[string]interface{} and assigns it to the Empl field.
func (o *Web3Bag) SetEmpl(v map[string]interface{}) {
	o.Empl = v
}

// GetPalsCount returns the PalsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3Bag) GetPalsCount() int32 {
	if o == nil || IsNil(o.PalsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PalsCount.Get()
}

// GetPalsCountOk returns a tuple with the PalsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3Bag) GetPalsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalsCount.Get(), o.PalsCount.IsSet()
}

// HasPalsCount returns a boolean if a field has been set.
func (o *Web3Bag) HasPalsCount() bool {
	if o != nil && o.PalsCount.IsSet() {
		return true
	}

	return false
}

// SetPalsCount gets a reference to the given NullableInt32 and assigns it to the PalsCount field.
func (o *Web3Bag) SetPalsCount(v int32) {
	o.PalsCount.Set(&v)
}
// SetPalsCountNil sets the value for PalsCount to be an explicit nil
func (o *Web3Bag) SetPalsCountNil() {
	o.PalsCount.Set(nil)
}

// UnsetPalsCount ensures that no value is present for PalsCount, not even an explicit nil
func (o *Web3Bag) UnsetPalsCount() {
	o.PalsCount.Unset()
}

func (o Web3Bag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3Bag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Empl != nil {
		toSerialize["empl"] = o.Empl
	}
	if o.PalsCount.IsSet() {
		toSerialize["pals_count"] = o.PalsCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3Bag) UnmarshalJSON(data []byte) (err error) {
	varWeb3Bag := _Web3Bag{}

	err = json.Unmarshal(data, &varWeb3Bag)

	if err != nil {
		return err
	}

	*o = Web3Bag(varWeb3Bag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "empl")
		delete(additionalProperties, "pals_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3Bag struct {
	value *Web3Bag
	isSet bool
}

func (v NullableWeb3Bag) Get() *Web3Bag {
	return v.value
}

func (v *NullableWeb3Bag) Set(val *Web3Bag) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3Bag) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3Bag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3Bag(val *Web3Bag) *NullableWeb3Bag {
	return &NullableWeb3Bag{value: val, isSet: true}
}

func (v NullableWeb3Bag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3Bag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


