
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsResponse{}

// GroupsResponse struct for GroupsResponse
type GroupsResponse struct {
	Groups []Group `json:"groups,omitempty"`
	PinnedGroups []Group `json:"pinned_groups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupsResponse GroupsResponse

// NewGroupsResponse instantiates a new GroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsResponse() *GroupsResponse {
	this := GroupsResponse{}
	return &this
}

// NewGroupsResponseWithDefaults instantiates a new GroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsResponseWithDefaults() *GroupsResponse {
	this := GroupsResponse{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsResponse) GetGroups() []Group {
	if o == nil {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsResponse) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GroupsResponse) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *GroupsResponse) SetGroups(v []Group) {
	o.Groups = v
}

// GetPinnedGroups returns the PinnedGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsResponse) GetPinnedGroups() []Group {
	if o == nil {
		var ret []Group
		return ret
	}
	return o.PinnedGroups
}

// GetPinnedGroupsOk returns a tuple with the PinnedGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsResponse) GetPinnedGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.PinnedGroups) {
		return nil, false
	}
	return o.PinnedGroups, true
}

// HasPinnedGroups returns a boolean if a field has been set.
func (o *GroupsResponse) HasPinnedGroups() bool {
	if o != nil && !IsNil(o.PinnedGroups) {
		return true
	}

	return false
}

// SetPinnedGroups gets a reference to the given []Group and assigns it to the PinnedGroups field.
func (o *GroupsResponse) SetPinnedGroups(v []Group) {
	o.PinnedGroups = v
}

func (o GroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.PinnedGroups != nil {
		toSerialize["pinned_groups"] = o.PinnedGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupsResponse := _GroupsResponse{}

	err = json.Unmarshal(data, &varGroupsResponse)

	if err != nil {
		return err
	}

	*o = GroupsResponse(varGroupsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groups")
		delete(additionalProperties, "pinned_groups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupsResponse struct {
	value *GroupsResponse
	isSet bool
}

func (v NullableGroupsResponse) Get() *GroupsResponse {
	return v.value
}

func (v *NullableGroupsResponse) Set(val *GroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsResponse(val *GroupsResponse) *NullableGroupsResponse {
	return &NullableGroupsResponse{value: val, isSet: true}
}

func (v NullableGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


