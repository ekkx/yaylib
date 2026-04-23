
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupResponse{}

// GroupResponse struct for GroupResponse
type GroupResponse struct {
	Group NullableGroup `json:"group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupResponse GroupResponse

// NewGroupResponse instantiates a new GroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupResponse() *GroupResponse {
	this := GroupResponse{}
	return &this
}

// NewGroupResponseWithDefaults instantiates a new GroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupResponseWithDefaults() *GroupResponse {
	this := GroupResponse{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupResponse) GetGroup() Group {
	if o == nil || IsNil(o.Group.Get()) {
		var ret Group
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupResponse) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *GroupResponse) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableGroup and assigns it to the Group field.
func (o *GroupResponse) SetGroup(v Group) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *GroupResponse) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *GroupResponse) UnsetGroup() {
	o.Group.Unset()
}

func (o GroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupResponse := _GroupResponse{}

	err = json.Unmarshal(data, &varGroupResponse)

	if err != nil {
		return err
	}

	*o = GroupResponse(varGroupResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupResponse struct {
	value *GroupResponse
	isSet bool
}

func (v NullableGroupResponse) Get() *GroupResponse {
	return v.value
}

func (v *NullableGroupResponse) Set(val *GroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupResponse(val *GroupResponse) *NullableGroupResponse {
	return &NullableGroupResponse{value: val, isSet: true}
}

func (v NullableGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


