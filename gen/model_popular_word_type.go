
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PopularWordType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PopularWordType{}

// PopularWordType struct for PopularWordType
type PopularWordType struct {
	ApiValue NullableString `json:"api_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PopularWordType PopularWordType

// NewPopularWordType instantiates a new PopularWordType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPopularWordType() *PopularWordType {
	this := PopularWordType{}
	return &this
}

// NewPopularWordTypeWithDefaults instantiates a new PopularWordType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPopularWordTypeWithDefaults() *PopularWordType {
	this := PopularWordType{}
	return &this
}

// GetApiValue returns the ApiValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PopularWordType) GetApiValue() string {
	if o == nil || IsNil(o.ApiValue.Get()) {
		var ret string
		return ret
	}
	return *o.ApiValue.Get()
}

// GetApiValueOk returns a tuple with the ApiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PopularWordType) GetApiValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiValue.Get(), o.ApiValue.IsSet()
}

// HasApiValue returns a boolean if a field has been set.
func (o *PopularWordType) HasApiValue() bool {
	if o != nil && o.ApiValue.IsSet() {
		return true
	}

	return false
}

// SetApiValue gets a reference to the given NullableString and assigns it to the ApiValue field.
func (o *PopularWordType) SetApiValue(v string) {
	o.ApiValue.Set(&v)
}
// SetApiValueNil sets the value for ApiValue to be an explicit nil
func (o *PopularWordType) SetApiValueNil() {
	o.ApiValue.Set(nil)
}

// UnsetApiValue ensures that no value is present for ApiValue, not even an explicit nil
func (o *PopularWordType) UnsetApiValue() {
	o.ApiValue.Unset()
}

func (o PopularWordType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PopularWordType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiValue.IsSet() {
		toSerialize["api_value"] = o.ApiValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PopularWordType) UnmarshalJSON(data []byte) (err error) {
	varPopularWordType := _PopularWordType{}

	err = json.Unmarshal(data, &varPopularWordType)

	if err != nil {
		return err
	}

	*o = PopularWordType(varPopularWordType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePopularWordType struct {
	value *PopularWordType
	isSet bool
}

func (v NullablePopularWordType) Get() *PopularWordType {
	return v.value
}

func (v *NullablePopularWordType) Set(val *PopularWordType) {
	v.value = val
	v.isSet = true
}

func (v NullablePopularWordType) IsSet() bool {
	return v.isSet
}

func (v *NullablePopularWordType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePopularWordType(val *PopularWordType) *NullablePopularWordType {
	return &NullablePopularWordType{value: val, isSet: true}
}

func (v NullablePopularWordType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePopularWordType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


