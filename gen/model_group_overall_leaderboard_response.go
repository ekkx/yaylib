
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupOverallLeaderboardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupOverallLeaderboardResponse{}

// GroupOverallLeaderboardResponse struct for GroupOverallLeaderboardResponse
type GroupOverallLeaderboardResponse struct {
	GroupLeaderboard []GroupLeaderboard `json:"group_leaderboard,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupOverallLeaderboardResponse GroupOverallLeaderboardResponse

// NewGroupOverallLeaderboardResponse instantiates a new GroupOverallLeaderboardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupOverallLeaderboardResponse() *GroupOverallLeaderboardResponse {
	this := GroupOverallLeaderboardResponse{}
	return &this
}

// NewGroupOverallLeaderboardResponseWithDefaults instantiates a new GroupOverallLeaderboardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupOverallLeaderboardResponseWithDefaults() *GroupOverallLeaderboardResponse {
	this := GroupOverallLeaderboardResponse{}
	return &this
}

// GetGroupLeaderboard returns the GroupLeaderboard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupOverallLeaderboardResponse) GetGroupLeaderboard() []GroupLeaderboard {
	if o == nil {
		var ret []GroupLeaderboard
		return ret
	}
	return o.GroupLeaderboard
}

// GetGroupLeaderboardOk returns a tuple with the GroupLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupOverallLeaderboardResponse) GetGroupLeaderboardOk() ([]GroupLeaderboard, bool) {
	if o == nil || IsNil(o.GroupLeaderboard) {
		return nil, false
	}
	return o.GroupLeaderboard, true
}

// HasGroupLeaderboard returns a boolean if a field has been set.
func (o *GroupOverallLeaderboardResponse) HasGroupLeaderboard() bool {
	if o != nil && !IsNil(o.GroupLeaderboard) {
		return true
	}

	return false
}

// SetGroupLeaderboard gets a reference to the given []GroupLeaderboard and assigns it to the GroupLeaderboard field.
func (o *GroupOverallLeaderboardResponse) SetGroupLeaderboard(v []GroupLeaderboard) {
	o.GroupLeaderboard = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupOverallLeaderboardResponse) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupOverallLeaderboardResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GroupOverallLeaderboardResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *GroupOverallLeaderboardResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *GroupOverallLeaderboardResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *GroupOverallLeaderboardResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o GroupOverallLeaderboardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupOverallLeaderboardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupLeaderboard != nil {
		toSerialize["group_leaderboard"] = o.GroupLeaderboard
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupOverallLeaderboardResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupOverallLeaderboardResponse := _GroupOverallLeaderboardResponse{}

	err = json.Unmarshal(data, &varGroupOverallLeaderboardResponse)

	if err != nil {
		return err
	}

	*o = GroupOverallLeaderboardResponse(varGroupOverallLeaderboardResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_leaderboard")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupOverallLeaderboardResponse struct {
	value *GroupOverallLeaderboardResponse
	isSet bool
}

func (v NullableGroupOverallLeaderboardResponse) Get() *GroupOverallLeaderboardResponse {
	return v.value
}

func (v *NullableGroupOverallLeaderboardResponse) Set(val *GroupOverallLeaderboardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupOverallLeaderboardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupOverallLeaderboardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupOverallLeaderboardResponse(val *GroupOverallLeaderboardResponse) *NullableGroupOverallLeaderboardResponse {
	return &NullableGroupOverallLeaderboardResponse{value: val, isSet: true}
}

func (v NullableGroupOverallLeaderboardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupOverallLeaderboardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


