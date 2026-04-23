
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupInCircleUserLeaderboardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupInCircleUserLeaderboardResponse{}

// GroupInCircleUserLeaderboardResponse struct for GroupInCircleUserLeaderboardResponse
type GroupInCircleUserLeaderboardResponse struct {
	UserLeaderboard []UserRank `json:"user_leaderboard,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupInCircleUserLeaderboardResponse GroupInCircleUserLeaderboardResponse

// NewGroupInCircleUserLeaderboardResponse instantiates a new GroupInCircleUserLeaderboardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInCircleUserLeaderboardResponse() *GroupInCircleUserLeaderboardResponse {
	this := GroupInCircleUserLeaderboardResponse{}
	return &this
}

// NewGroupInCircleUserLeaderboardResponseWithDefaults instantiates a new GroupInCircleUserLeaderboardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInCircleUserLeaderboardResponseWithDefaults() *GroupInCircleUserLeaderboardResponse {
	this := GroupInCircleUserLeaderboardResponse{}
	return &this
}

// GetUserLeaderboard returns the UserLeaderboard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupInCircleUserLeaderboardResponse) GetUserLeaderboard() []UserRank {
	if o == nil {
		var ret []UserRank
		return ret
	}
	return o.UserLeaderboard
}

// GetUserLeaderboardOk returns a tuple with the UserLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupInCircleUserLeaderboardResponse) GetUserLeaderboardOk() ([]UserRank, bool) {
	if o == nil || IsNil(o.UserLeaderboard) {
		return nil, false
	}
	return o.UserLeaderboard, true
}

// HasUserLeaderboard returns a boolean if a field has been set.
func (o *GroupInCircleUserLeaderboardResponse) HasUserLeaderboard() bool {
	if o != nil && !IsNil(o.UserLeaderboard) {
		return true
	}

	return false
}

// SetUserLeaderboard gets a reference to the given []UserRank and assigns it to the UserLeaderboard field.
func (o *GroupInCircleUserLeaderboardResponse) SetUserLeaderboard(v []UserRank) {
	o.UserLeaderboard = v
}

func (o GroupInCircleUserLeaderboardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupInCircleUserLeaderboardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserLeaderboard != nil {
		toSerialize["user_leaderboard"] = o.UserLeaderboard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupInCircleUserLeaderboardResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupInCircleUserLeaderboardResponse := _GroupInCircleUserLeaderboardResponse{}

	err = json.Unmarshal(data, &varGroupInCircleUserLeaderboardResponse)

	if err != nil {
		return err
	}

	*o = GroupInCircleUserLeaderboardResponse(varGroupInCircleUserLeaderboardResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_leaderboard")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupInCircleUserLeaderboardResponse struct {
	value *GroupInCircleUserLeaderboardResponse
	isSet bool
}

func (v NullableGroupInCircleUserLeaderboardResponse) Get() *GroupInCircleUserLeaderboardResponse {
	return v.value
}

func (v *NullableGroupInCircleUserLeaderboardResponse) Set(val *GroupInCircleUserLeaderboardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInCircleUserLeaderboardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInCircleUserLeaderboardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInCircleUserLeaderboardResponse(val *GroupInCircleUserLeaderboardResponse) *NullableGroupInCircleUserLeaderboardResponse {
	return &NullableGroupInCircleUserLeaderboardResponse{value: val, isSet: true}
}

func (v NullableGroupInCircleUserLeaderboardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInCircleUserLeaderboardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


