
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Gender type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gender{}

// Gender struct for Gender
type Gender struct {
	ApiIntValue NullableInt32 `json:"api_int_value,omitempty"`
	ApiStringValue NullableString `json:"api_string_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Gender Gender

// NewGender instantiates a new Gender object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGender() *Gender {
	this := Gender{}
	return &this
}

// NewGenderWithDefaults instantiates a new Gender object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenderWithDefaults() *Gender {
	this := Gender{}
	return &this
}

// GetApiIntValue returns the ApiIntValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gender) GetApiIntValue() int32 {
	if o == nil || IsNil(o.ApiIntValue.Get()) {
		var ret int32
		return ret
	}
	return *o.ApiIntValue.Get()
}

// GetApiIntValueOk returns a tuple with the ApiIntValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gender) GetApiIntValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiIntValue.Get(), o.ApiIntValue.IsSet()
}

// HasApiIntValue returns a boolean if a field has been set.
func (o *Gender) HasApiIntValue() bool {
	if o != nil && o.ApiIntValue.IsSet() {
		return true
	}

	return false
}

// SetApiIntValue gets a reference to the given NullableInt32 and assigns it to the ApiIntValue field.
func (o *Gender) SetApiIntValue(v int32) {
	o.ApiIntValue.Set(&v)
}
// SetApiIntValueNil sets the value for ApiIntValue to be an explicit nil
func (o *Gender) SetApiIntValueNil() {
	o.ApiIntValue.Set(nil)
}

// UnsetApiIntValue ensures that no value is present for ApiIntValue, not even an explicit nil
func (o *Gender) UnsetApiIntValue() {
	o.ApiIntValue.Unset()
}

// GetApiStringValue returns the ApiStringValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gender) GetApiStringValue() string {
	if o == nil || IsNil(o.ApiStringValue.Get()) {
		var ret string
		return ret
	}
	return *o.ApiStringValue.Get()
}

// GetApiStringValueOk returns a tuple with the ApiStringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gender) GetApiStringValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiStringValue.Get(), o.ApiStringValue.IsSet()
}

// HasApiStringValue returns a boolean if a field has been set.
func (o *Gender) HasApiStringValue() bool {
	if o != nil && o.ApiStringValue.IsSet() {
		return true
	}

	return false
}

// SetApiStringValue gets a reference to the given NullableString and assigns it to the ApiStringValue field.
func (o *Gender) SetApiStringValue(v string) {
	o.ApiStringValue.Set(&v)
}
// SetApiStringValueNil sets the value for ApiStringValue to be an explicit nil
func (o *Gender) SetApiStringValueNil() {
	o.ApiStringValue.Set(nil)
}

// UnsetApiStringValue ensures that no value is present for ApiStringValue, not even an explicit nil
func (o *Gender) UnsetApiStringValue() {
	o.ApiStringValue.Unset()
}

func (o Gender) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gender) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiIntValue.IsSet() {
		toSerialize["api_int_value"] = o.ApiIntValue.Get()
	}
	if o.ApiStringValue.IsSet() {
		toSerialize["api_string_value"] = o.ApiStringValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Gender) UnmarshalJSON(data []byte) (err error) {
	varGender := _Gender{}

	err = json.Unmarshal(data, &varGender)

	if err != nil {
		return err
	}

	*o = Gender(varGender)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_int_value")
		delete(additionalProperties, "api_string_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGender struct {
	value *Gender
	isSet bool
}

func (v NullableGender) Get() *Gender {
	return v.value
}

func (v *NullableGender) Set(val *Gender) {
	v.value = val
	v.isSet = true
}

func (v NullableGender) IsSet() bool {
	return v.isSet
}

func (v *NullableGender) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGender(val *Gender) *NullableGender {
	return &NullableGender{value: val, isSet: true}
}

func (v NullableGender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGender) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


