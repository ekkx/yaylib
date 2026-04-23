
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupRanking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRanking{}

// GroupRanking struct for GroupRanking
type GroupRanking struct {
	Group NullableModelGroup `json:"group,omitempty"`
	Ranking NullableInt32 `json:"ranking,omitempty"`
	TotalScores NullableInt32 `json:"total_scores,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupRanking GroupRanking

// NewGroupRanking instantiates a new GroupRanking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRanking() *GroupRanking {
	this := GroupRanking{}
	return &this
}

// NewGroupRankingWithDefaults instantiates a new GroupRanking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRankingWithDefaults() *GroupRanking {
	this := GroupRanking{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupRanking) GetGroup() ModelGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret ModelGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupRanking) GetGroupOk() (*ModelGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *GroupRanking) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableModelGroup and assigns it to the Group field.
func (o *GroupRanking) SetGroup(v ModelGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *GroupRanking) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *GroupRanking) UnsetGroup() {
	o.Group.Unset()
}

// GetRanking returns the Ranking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupRanking) GetRanking() int32 {
	if o == nil || IsNil(o.Ranking.Get()) {
		var ret int32
		return ret
	}
	return *o.Ranking.Get()
}

// GetRankingOk returns a tuple with the Ranking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupRanking) GetRankingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ranking.Get(), o.Ranking.IsSet()
}

// HasRanking returns a boolean if a field has been set.
func (o *GroupRanking) HasRanking() bool {
	if o != nil && o.Ranking.IsSet() {
		return true
	}

	return false
}

// SetRanking gets a reference to the given NullableInt32 and assigns it to the Ranking field.
func (o *GroupRanking) SetRanking(v int32) {
	o.Ranking.Set(&v)
}
// SetRankingNil sets the value for Ranking to be an explicit nil
func (o *GroupRanking) SetRankingNil() {
	o.Ranking.Set(nil)
}

// UnsetRanking ensures that no value is present for Ranking, not even an explicit nil
func (o *GroupRanking) UnsetRanking() {
	o.Ranking.Unset()
}

// GetTotalScores returns the TotalScores field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupRanking) GetTotalScores() int32 {
	if o == nil || IsNil(o.TotalScores.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalScores.Get()
}

// GetTotalScoresOk returns a tuple with the TotalScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupRanking) GetTotalScoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalScores.Get(), o.TotalScores.IsSet()
}

// HasTotalScores returns a boolean if a field has been set.
func (o *GroupRanking) HasTotalScores() bool {
	if o != nil && o.TotalScores.IsSet() {
		return true
	}

	return false
}

// SetTotalScores gets a reference to the given NullableInt32 and assigns it to the TotalScores field.
func (o *GroupRanking) SetTotalScores(v int32) {
	o.TotalScores.Set(&v)
}
// SetTotalScoresNil sets the value for TotalScores to be an explicit nil
func (o *GroupRanking) SetTotalScoresNil() {
	o.TotalScores.Set(nil)
}

// UnsetTotalScores ensures that no value is present for TotalScores, not even an explicit nil
func (o *GroupRanking) UnsetTotalScores() {
	o.TotalScores.Unset()
}

func (o GroupRanking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRanking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Ranking.IsSet() {
		toSerialize["ranking"] = o.Ranking.Get()
	}
	if o.TotalScores.IsSet() {
		toSerialize["total_scores"] = o.TotalScores.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupRanking) UnmarshalJSON(data []byte) (err error) {
	varGroupRanking := _GroupRanking{}

	err = json.Unmarshal(data, &varGroupRanking)

	if err != nil {
		return err
	}

	*o = GroupRanking(varGroupRanking)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "ranking")
		delete(additionalProperties, "total_scores")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupRanking struct {
	value *GroupRanking
	isSet bool
}

func (v NullableGroupRanking) Get() *GroupRanking {
	return v.value
}

func (v *NullableGroupRanking) Set(val *GroupRanking) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRanking) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRanking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRanking(val *GroupRanking) *NullableGroupRanking {
	return &NullableGroupRanking{value: val, isSet: true}
}

func (v NullableGroupRanking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRanking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


