
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUserResponse{}

// GroupUserResponse struct for GroupUserResponse
type GroupUserResponse struct {
	GroupUser NullableGroupUser `json:"group_user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupUserResponse GroupUserResponse

// NewGroupUserResponse instantiates a new GroupUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUserResponse() *GroupUserResponse {
	this := GroupUserResponse{}
	return &this
}

// NewGroupUserResponseWithDefaults instantiates a new GroupUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUserResponseWithDefaults() *GroupUserResponse {
	this := GroupUserResponse{}
	return &this
}

// GetGroupUser returns the GroupUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUserResponse) GetGroupUser() GroupUser {
	if o == nil || IsNil(o.GroupUser.Get()) {
		var ret GroupUser
		return ret
	}
	return *o.GroupUser.Get()
}

// GetGroupUserOk returns a tuple with the GroupUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUserResponse) GetGroupUserOk() (*GroupUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupUser.Get(), o.GroupUser.IsSet()
}

// HasGroupUser returns a boolean if a field has been set.
func (o *GroupUserResponse) HasGroupUser() bool {
	if o != nil && o.GroupUser.IsSet() {
		return true
	}

	return false
}

// SetGroupUser gets a reference to the given NullableGroupUser and assigns it to the GroupUser field.
func (o *GroupUserResponse) SetGroupUser(v GroupUser) {
	o.GroupUser.Set(&v)
}
// SetGroupUserNil sets the value for GroupUser to be an explicit nil
func (o *GroupUserResponse) SetGroupUserNil() {
	o.GroupUser.Set(nil)
}

// UnsetGroupUser ensures that no value is present for GroupUser, not even an explicit nil
func (o *GroupUserResponse) UnsetGroupUser() {
	o.GroupUser.Unset()
}

func (o GroupUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupUser.IsSet() {
		toSerialize["group_user"] = o.GroupUser.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupUserResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupUserResponse := _GroupUserResponse{}

	err = json.Unmarshal(data, &varGroupUserResponse)

	if err != nil {
		return err
	}

	*o = GroupUserResponse(varGroupUserResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupUserResponse struct {
	value *GroupUserResponse
	isSet bool
}

func (v NullableGroupUserResponse) Get() *GroupUserResponse {
	return v.value
}

func (v *NullableGroupUserResponse) Set(val *GroupUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUserResponse(val *GroupUserResponse) *NullableGroupUserResponse {
	return &NullableGroupUserResponse{value: val, isSet: true}
}

func (v NullableGroupUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


