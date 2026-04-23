
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUsersResponse{}

// GroupUsersResponse struct for GroupUsersResponse
type GroupUsersResponse struct {
	GroupUsers []GroupUser `json:"group_users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupUsersResponse GroupUsersResponse

// NewGroupUsersResponse instantiates a new GroupUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUsersResponse() *GroupUsersResponse {
	this := GroupUsersResponse{}
	return &this
}

// NewGroupUsersResponseWithDefaults instantiates a new GroupUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUsersResponseWithDefaults() *GroupUsersResponse {
	this := GroupUsersResponse{}
	return &this
}

// GetGroupUsers returns the GroupUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUsersResponse) GetGroupUsers() []GroupUser {
	if o == nil {
		var ret []GroupUser
		return ret
	}
	return o.GroupUsers
}

// GetGroupUsersOk returns a tuple with the GroupUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUsersResponse) GetGroupUsersOk() ([]GroupUser, bool) {
	if o == nil || IsNil(o.GroupUsers) {
		return nil, false
	}
	return o.GroupUsers, true
}

// HasGroupUsers returns a boolean if a field has been set.
func (o *GroupUsersResponse) HasGroupUsers() bool {
	if o != nil && !IsNil(o.GroupUsers) {
		return true
	}

	return false
}

// SetGroupUsers gets a reference to the given []GroupUser and assigns it to the GroupUsers field.
func (o *GroupUsersResponse) SetGroupUsers(v []GroupUser) {
	o.GroupUsers = v
}

func (o GroupUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupUsers != nil {
		toSerialize["group_users"] = o.GroupUsers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupUsersResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupUsersResponse := _GroupUsersResponse{}

	err = json.Unmarshal(data, &varGroupUsersResponse)

	if err != nil {
		return err
	}

	*o = GroupUsersResponse(varGroupUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupUsersResponse struct {
	value *GroupUsersResponse
	isSet bool
}

func (v NullableGroupUsersResponse) Get() *GroupUsersResponse {
	return v.value
}

func (v *NullableGroupUsersResponse) Set(val *GroupUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUsersResponse(val *GroupUsersResponse) *NullableGroupUsersResponse {
	return &NullableGroupUsersResponse{value: val, isSet: true}
}

func (v NullableGroupUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


