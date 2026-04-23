
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupOverallLeaderboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupOverallLeaderboard{}

// GroupOverallLeaderboard struct for GroupOverallLeaderboard
type GroupOverallLeaderboard struct {
	GroupRankings []GroupRanking `json:"group_rankings,omitempty"`
	UpdatedAtMillis NullableInt64 `json:"updated_at_millis,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupOverallLeaderboard GroupOverallLeaderboard

// NewGroupOverallLeaderboard instantiates a new GroupOverallLeaderboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupOverallLeaderboard() *GroupOverallLeaderboard {
	this := GroupOverallLeaderboard{}
	return &this
}

// NewGroupOverallLeaderboardWithDefaults instantiates a new GroupOverallLeaderboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupOverallLeaderboardWithDefaults() *GroupOverallLeaderboard {
	this := GroupOverallLeaderboard{}
	return &this
}

// GetGroupRankings returns the GroupRankings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupOverallLeaderboard) GetGroupRankings() []GroupRanking {
	if o == nil {
		var ret []GroupRanking
		return ret
	}
	return o.GroupRankings
}

// GetGroupRankingsOk returns a tuple with the GroupRankings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupOverallLeaderboard) GetGroupRankingsOk() ([]GroupRanking, bool) {
	if o == nil || IsNil(o.GroupRankings) {
		return nil, false
	}
	return o.GroupRankings, true
}

// HasGroupRankings returns a boolean if a field has been set.
func (o *GroupOverallLeaderboard) HasGroupRankings() bool {
	if o != nil && !IsNil(o.GroupRankings) {
		return true
	}

	return false
}

// SetGroupRankings gets a reference to the given []GroupRanking and assigns it to the GroupRankings field.
func (o *GroupOverallLeaderboard) SetGroupRankings(v []GroupRanking) {
	o.GroupRankings = v
}

// GetUpdatedAtMillis returns the UpdatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupOverallLeaderboard) GetUpdatedAtMillis() int64 {
	if o == nil || IsNil(o.UpdatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAtMillis.Get()
}

// GetUpdatedAtMillisOk returns a tuple with the UpdatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupOverallLeaderboard) GetUpdatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAtMillis.Get(), o.UpdatedAtMillis.IsSet()
}

// HasUpdatedAtMillis returns a boolean if a field has been set.
func (o *GroupOverallLeaderboard) HasUpdatedAtMillis() bool {
	if o != nil && o.UpdatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAtMillis gets a reference to the given NullableInt64 and assigns it to the UpdatedAtMillis field.
func (o *GroupOverallLeaderboard) SetUpdatedAtMillis(v int64) {
	o.UpdatedAtMillis.Set(&v)
}
// SetUpdatedAtMillisNil sets the value for UpdatedAtMillis to be an explicit nil
func (o *GroupOverallLeaderboard) SetUpdatedAtMillisNil() {
	o.UpdatedAtMillis.Set(nil)
}

// UnsetUpdatedAtMillis ensures that no value is present for UpdatedAtMillis, not even an explicit nil
func (o *GroupOverallLeaderboard) UnsetUpdatedAtMillis() {
	o.UpdatedAtMillis.Unset()
}

func (o GroupOverallLeaderboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupOverallLeaderboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupRankings != nil {
		toSerialize["group_rankings"] = o.GroupRankings
	}
	if o.UpdatedAtMillis.IsSet() {
		toSerialize["updated_at_millis"] = o.UpdatedAtMillis.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupOverallLeaderboard) UnmarshalJSON(data []byte) (err error) {
	varGroupOverallLeaderboard := _GroupOverallLeaderboard{}

	err = json.Unmarshal(data, &varGroupOverallLeaderboard)

	if err != nil {
		return err
	}

	*o = GroupOverallLeaderboard(varGroupOverallLeaderboard)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_rankings")
		delete(additionalProperties, "updated_at_millis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupOverallLeaderboard struct {
	value *GroupOverallLeaderboard
	isSet bool
}

func (v NullableGroupOverallLeaderboard) Get() *GroupOverallLeaderboard {
	return v.value
}

func (v *NullableGroupOverallLeaderboard) Set(val *GroupOverallLeaderboard) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupOverallLeaderboard) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupOverallLeaderboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupOverallLeaderboard(val *GroupOverallLeaderboard) *NullableGroupOverallLeaderboard {
	return &NullableGroupOverallLeaderboard{value: val, isSet: true}
}

func (v NullableGroupOverallLeaderboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupOverallLeaderboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


