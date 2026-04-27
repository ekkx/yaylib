
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ApiB type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiB{}

// ApiB struct for ApiB
type ApiB struct {
	D map[string]interface{} `json:"d,omitempty"`
	F map[string]interface{} `json:"f,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiB ApiB

// NewApiB instantiates a new ApiB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiB() *ApiB {
	this := ApiB{}
	return &this
}

// NewApiBWithDefaults instantiates a new ApiB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBWithDefaults() *ApiB {
	this := ApiB{}
	return &this
}

// GetD returns the D field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiB) GetD() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiB) GetDOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.D) {
		return map[string]interface{}{}, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *ApiB) HasD() bool {
	if o != nil && !IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given map[string]interface{} and assigns it to the D field.
func (o *ApiB) SetD(v map[string]interface{}) {
	o.D = v
}

// GetF returns the F field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiB) GetF() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiB) GetFOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.F) {
		return map[string]interface{}{}, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *ApiB) HasF() bool {
	if o != nil && !IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given map[string]interface{} and assigns it to the F field.
func (o *ApiB) SetF(v map[string]interface{}) {
	o.F = v
}

func (o ApiB) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiB) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.D != nil {
		toSerialize["d"] = o.D
	}
	if o.F != nil {
		toSerialize["f"] = o.F
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiB) UnmarshalJSON(data []byte) (err error) {
	varApiB := _ApiB{}

	err = json.Unmarshal(data, &varApiB)

	if err != nil {
		return err
	}

	*o = ApiB(varApiB)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "d")
		delete(additionalProperties, "f")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiB struct {
	value *ApiB
	isSet bool
}

func (v NullableApiB) Get() *ApiB {
	return v.value
}

func (v *NullableApiB) Set(val *ApiB) {
	v.value = val
	v.isSet = true
}

func (v NullableApiB) IsSet() bool {
	return v.isSet
}

func (v *NullableApiB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiB(val *ApiB) *NullableApiB {
	return &NullableApiB{value: val, isSet: true}
}

func (v NullableApiB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


