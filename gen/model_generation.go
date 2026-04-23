
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Generation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Generation{}

// Generation struct for Generation
type Generation struct {
	ApiValue NullableInt32 `json:"api_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Generation Generation

// NewGeneration instantiates a new Generation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneration() *Generation {
	this := Generation{}
	return &this
}

// NewGenerationWithDefaults instantiates a new Generation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerationWithDefaults() *Generation {
	this := Generation{}
	return &this
}

// GetApiValue returns the ApiValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Generation) GetApiValue() int32 {
	if o == nil || IsNil(o.ApiValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ApiValue.Get()
}

// GetApiValueOk returns a tuple with the ApiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Generation) GetApiValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiValue.Get(), o.ApiValue.IsSet()
}

// HasApiValue returns a boolean if a field has been set.
func (o *Generation) HasApiValue() bool {
	if o != nil && o.ApiValue.IsSet() {
		return true
	}

	return false
}

// SetApiValue gets a reference to the given NullableInt32 and assigns it to the ApiValue field.
func (o *Generation) SetApiValue(v int32) {
	o.ApiValue.Set(&v)
}
// SetApiValueNil sets the value for ApiValue to be an explicit nil
func (o *Generation) SetApiValueNil() {
	o.ApiValue.Set(nil)
}

// UnsetApiValue ensures that no value is present for ApiValue, not even an explicit nil
func (o *Generation) UnsetApiValue() {
	o.ApiValue.Unset()
}

func (o Generation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Generation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiValue.IsSet() {
		toSerialize["api_value"] = o.ApiValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Generation) UnmarshalJSON(data []byte) (err error) {
	varGeneration := _Generation{}

	err = json.Unmarshal(data, &varGeneration)

	if err != nil {
		return err
	}

	*o = Generation(varGeneration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGeneration struct {
	value *Generation
	isSet bool
}

func (v NullableGeneration) Get() *Generation {
	return v.value
}

func (v *NullableGeneration) Set(val *Generation) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneration) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneration(val *Generation) *NullableGeneration {
	return &NullableGeneration{value: val, isSet: true}
}

func (v NullableGeneration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


