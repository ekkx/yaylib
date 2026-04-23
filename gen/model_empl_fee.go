
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplFee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplFee{}

// EmplFee struct for EmplFee
type EmplFee struct {
	GasFee NullableString `json:"gas_fee,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplFee EmplFee

// NewEmplFee instantiates a new EmplFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplFee() *EmplFee {
	this := EmplFee{}
	return &this
}

// NewEmplFeeWithDefaults instantiates a new EmplFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplFeeWithDefaults() *EmplFee {
	this := EmplFee{}
	return &this
}

// GetGasFee returns the GasFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplFee) GetGasFee() string {
	if o == nil || IsNil(o.GasFee.Get()) {
		var ret string
		return ret
	}
	return *o.GasFee.Get()
}

// GetGasFeeOk returns a tuple with the GasFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplFee) GetGasFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GasFee.Get(), o.GasFee.IsSet()
}

// HasGasFee returns a boolean if a field has been set.
func (o *EmplFee) HasGasFee() bool {
	if o != nil && o.GasFee.IsSet() {
		return true
	}

	return false
}

// SetGasFee gets a reference to the given NullableString and assigns it to the GasFee field.
func (o *EmplFee) SetGasFee(v string) {
	o.GasFee.Set(&v)
}
// SetGasFeeNil sets the value for GasFee to be an explicit nil
func (o *EmplFee) SetGasFeeNil() {
	o.GasFee.Set(nil)
}

// UnsetGasFee ensures that no value is present for GasFee, not even an explicit nil
func (o *EmplFee) UnsetGasFee() {
	o.GasFee.Unset()
}

func (o EmplFee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplFee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GasFee.IsSet() {
		toSerialize["gas_fee"] = o.GasFee.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplFee) UnmarshalJSON(data []byte) (err error) {
	varEmplFee := _EmplFee{}

	err = json.Unmarshal(data, &varEmplFee)

	if err != nil {
		return err
	}

	*o = EmplFee(varEmplFee)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gas_fee")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplFee struct {
	value *EmplFee
	isSet bool
}

func (v NullableEmplFee) Get() *EmplFee {
	return v.value
}

func (v *NullableEmplFee) Set(val *EmplFee) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplFee) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplFee(val *EmplFee) *NullableEmplFee {
	return &NullableEmplFee{value: val, isSet: true}
}

func (v NullableEmplFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


