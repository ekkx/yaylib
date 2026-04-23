
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupUpdatedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUpdatedData{}

// GroupUpdatedData struct for GroupUpdatedData
type GroupUpdatedData struct {
	GroupId NullableInt64 `json:"group_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupUpdatedData GroupUpdatedData

// NewGroupUpdatedData instantiates a new GroupUpdatedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdatedData() *GroupUpdatedData {
	this := GroupUpdatedData{}
	return &this
}

// NewGroupUpdatedDataWithDefaults instantiates a new GroupUpdatedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdatedDataWithDefaults() *GroupUpdatedData {
	this := GroupUpdatedData{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUpdatedData) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUpdatedData) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupUpdatedData) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *GroupUpdatedData) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *GroupUpdatedData) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *GroupUpdatedData) UnsetGroupId() {
	o.GroupId.Unset()
}

func (o GroupUpdatedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUpdatedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId.IsSet() {
		toSerialize["group_id"] = o.GroupId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupUpdatedData) UnmarshalJSON(data []byte) (err error) {
	varGroupUpdatedData := _GroupUpdatedData{}

	err = json.Unmarshal(data, &varGroupUpdatedData)

	if err != nil {
		return err
	}

	*o = GroupUpdatedData(varGroupUpdatedData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupUpdatedData struct {
	value *GroupUpdatedData
	isSet bool
}

func (v NullableGroupUpdatedData) Get() *GroupUpdatedData {
	return v.value
}

func (v *NullableGroupUpdatedData) Set(val *GroupUpdatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdatedData(val *GroupUpdatedData) *NullableGroupUpdatedData {
	return &NullableGroupUpdatedData{value: val, isSet: true}
}

func (v NullableGroupUpdatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


