
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupMuteUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupMuteUsersResponse{}

// GroupMuteUsersResponse struct for GroupMuteUsersResponse
type GroupMuteUsersResponse struct {
	GroupUsers []GroupUser `json:"group_users,omitempty"`
	NextCursor NullableInt64 `json:"next_cursor,omitempty"`
	Total NullableInt32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupMuteUsersResponse GroupMuteUsersResponse

// NewGroupMuteUsersResponse instantiates a new GroupMuteUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMuteUsersResponse() *GroupMuteUsersResponse {
	this := GroupMuteUsersResponse{}
	return &this
}

// NewGroupMuteUsersResponseWithDefaults instantiates a new GroupMuteUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMuteUsersResponseWithDefaults() *GroupMuteUsersResponse {
	this := GroupMuteUsersResponse{}
	return &this
}

// GetGroupUsers returns the GroupUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupMuteUsersResponse) GetGroupUsers() []GroupUser {
	if o == nil {
		var ret []GroupUser
		return ret
	}
	return o.GroupUsers
}

// GetGroupUsersOk returns a tuple with the GroupUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMuteUsersResponse) GetGroupUsersOk() ([]GroupUser, bool) {
	if o == nil || IsNil(o.GroupUsers) {
		return nil, false
	}
	return o.GroupUsers, true
}

// HasGroupUsers returns a boolean if a field has been set.
func (o *GroupMuteUsersResponse) HasGroupUsers() bool {
	if o != nil && !IsNil(o.GroupUsers) {
		return true
	}

	return false
}

// SetGroupUsers gets a reference to the given []GroupUser and assigns it to the GroupUsers field.
func (o *GroupMuteUsersResponse) SetGroupUsers(v []GroupUser) {
	o.GroupUsers = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupMuteUsersResponse) GetNextCursor() int64 {
	if o == nil || IsNil(o.NextCursor.Get()) {
		var ret int64
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMuteUsersResponse) GetNextCursorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *GroupMuteUsersResponse) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableInt64 and assigns it to the NextCursor field.
func (o *GroupMuteUsersResponse) SetNextCursor(v int64) {
	o.NextCursor.Set(&v)
}
// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *GroupMuteUsersResponse) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *GroupMuteUsersResponse) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupMuteUsersResponse) GetTotal() int32 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret int32
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupMuteUsersResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *GroupMuteUsersResponse) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableInt32 and assigns it to the Total field.
func (o *GroupMuteUsersResponse) SetTotal(v int32) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *GroupMuteUsersResponse) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *GroupMuteUsersResponse) UnsetTotal() {
	o.Total.Unset()
}

func (o GroupMuteUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupMuteUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupUsers != nil {
		toSerialize["group_users"] = o.GroupUsers
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupMuteUsersResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupMuteUsersResponse := _GroupMuteUsersResponse{}

	err = json.Unmarshal(data, &varGroupMuteUsersResponse)

	if err != nil {
		return err
	}

	*o = GroupMuteUsersResponse(varGroupMuteUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_users")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupMuteUsersResponse struct {
	value *GroupMuteUsersResponse
	isSet bool
}

func (v NullableGroupMuteUsersResponse) Get() *GroupMuteUsersResponse {
	return v.value
}

func (v *NullableGroupMuteUsersResponse) Set(val *GroupMuteUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupMuteUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupMuteUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupMuteUsersResponse(val *GroupMuteUsersResponse) *NullableGroupMuteUsersResponse {
	return &NullableGroupMuteUsersResponse{value: val, isSet: true}
}

func (v NullableGroupMuteUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupMuteUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


