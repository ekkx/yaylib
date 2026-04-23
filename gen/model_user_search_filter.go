
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserSearchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSearchFilter{}

// UserSearchFilter struct for UserSearchFilter
type UserSearchFilter struct {
	TitleId NullableInt32 `json:"title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSearchFilter UserSearchFilter

// NewUserSearchFilter instantiates a new UserSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSearchFilter() *UserSearchFilter {
	this := UserSearchFilter{}
	return &this
}

// NewUserSearchFilterWithDefaults instantiates a new UserSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSearchFilterWithDefaults() *UserSearchFilter {
	this := UserSearchFilter{}
	return &this
}

// GetTitleId returns the TitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchFilter) GetTitleId() int32 {
	if o == nil || IsNil(o.TitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.TitleId.Get()
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchFilter) GetTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleId.Get(), o.TitleId.IsSet()
}

// HasTitleId returns a boolean if a field has been set.
func (o *UserSearchFilter) HasTitleId() bool {
	if o != nil && o.TitleId.IsSet() {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given NullableInt32 and assigns it to the TitleId field.
func (o *UserSearchFilter) SetTitleId(v int32) {
	o.TitleId.Set(&v)
}
// SetTitleIdNil sets the value for TitleId to be an explicit nil
func (o *UserSearchFilter) SetTitleIdNil() {
	o.TitleId.Set(nil)
}

// UnsetTitleId ensures that no value is present for TitleId, not even an explicit nil
func (o *UserSearchFilter) UnsetTitleId() {
	o.TitleId.Unset()
}

func (o UserSearchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSearchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TitleId.IsSet() {
		toSerialize["title_id"] = o.TitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSearchFilter) UnmarshalJSON(data []byte) (err error) {
	varUserSearchFilter := _UserSearchFilter{}

	err = json.Unmarshal(data, &varUserSearchFilter)

	if err != nil {
		return err
	}

	*o = UserSearchFilter(varUserSearchFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSearchFilter struct {
	value *UserSearchFilter
	isSet bool
}

func (v NullableUserSearchFilter) Get() *UserSearchFilter {
	return v.value
}

func (v *NullableUserSearchFilter) Set(val *UserSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSearchFilter(val *UserSearchFilter) *NullableUserSearchFilter {
	return &NullableUserSearchFilter{value: val, isSet: true}
}

func (v NullableUserSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


