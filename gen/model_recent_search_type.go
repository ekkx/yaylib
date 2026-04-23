
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RecentSearchType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentSearchType{}

// RecentSearchType struct for RecentSearchType
type RecentSearchType struct {
	ApiValue NullableString `json:"api_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecentSearchType RecentSearchType

// NewRecentSearchType instantiates a new RecentSearchType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentSearchType() *RecentSearchType {
	this := RecentSearchType{}
	return &this
}

// NewRecentSearchTypeWithDefaults instantiates a new RecentSearchType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentSearchTypeWithDefaults() *RecentSearchType {
	this := RecentSearchType{}
	return &this
}

// GetApiValue returns the ApiValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentSearchType) GetApiValue() string {
	if o == nil || IsNil(o.ApiValue.Get()) {
		var ret string
		return ret
	}
	return *o.ApiValue.Get()
}

// GetApiValueOk returns a tuple with the ApiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentSearchType) GetApiValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiValue.Get(), o.ApiValue.IsSet()
}

// HasApiValue returns a boolean if a field has been set.
func (o *RecentSearchType) HasApiValue() bool {
	if o != nil && o.ApiValue.IsSet() {
		return true
	}

	return false
}

// SetApiValue gets a reference to the given NullableString and assigns it to the ApiValue field.
func (o *RecentSearchType) SetApiValue(v string) {
	o.ApiValue.Set(&v)
}
// SetApiValueNil sets the value for ApiValue to be an explicit nil
func (o *RecentSearchType) SetApiValueNil() {
	o.ApiValue.Set(nil)
}

// UnsetApiValue ensures that no value is present for ApiValue, not even an explicit nil
func (o *RecentSearchType) UnsetApiValue() {
	o.ApiValue.Unset()
}

func (o RecentSearchType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentSearchType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiValue.IsSet() {
		toSerialize["api_value"] = o.ApiValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecentSearchType) UnmarshalJSON(data []byte) (err error) {
	varRecentSearchType := _RecentSearchType{}

	err = json.Unmarshal(data, &varRecentSearchType)

	if err != nil {
		return err
	}

	*o = RecentSearchType(varRecentSearchType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecentSearchType struct {
	value *RecentSearchType
	isSet bool
}

func (v NullableRecentSearchType) Get() *RecentSearchType {
	return v.value
}

func (v *NullableRecentSearchType) Set(val *RecentSearchType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentSearchType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentSearchType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentSearchType(val *RecentSearchType) *NullableRecentSearchType {
	return &NullableRecentSearchType{value: val, isSet: true}
}

func (v NullableRecentSearchType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentSearchType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


