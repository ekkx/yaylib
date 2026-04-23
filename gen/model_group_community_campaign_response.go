
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupCommunityCampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupCommunityCampaignResponse{}

// GroupCommunityCampaignResponse struct for GroupCommunityCampaignResponse
type GroupCommunityCampaignResponse struct {
	EligibleForApply NullableBool `json:"eligible_for_apply,omitempty"`
	ViewableGroupLeaderboard NullableBool `json:"viewable_group_leaderboard,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupCommunityCampaignResponse GroupCommunityCampaignResponse

// NewGroupCommunityCampaignResponse instantiates a new GroupCommunityCampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCommunityCampaignResponse() *GroupCommunityCampaignResponse {
	this := GroupCommunityCampaignResponse{}
	return &this
}

// NewGroupCommunityCampaignResponseWithDefaults instantiates a new GroupCommunityCampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCommunityCampaignResponseWithDefaults() *GroupCommunityCampaignResponse {
	this := GroupCommunityCampaignResponse{}
	return &this
}

// GetEligibleForApply returns the EligibleForApply field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupCommunityCampaignResponse) GetEligibleForApply() bool {
	if o == nil || IsNil(o.EligibleForApply.Get()) {
		var ret bool
		return ret
	}
	return *o.EligibleForApply.Get()
}

// GetEligibleForApplyOk returns a tuple with the EligibleForApply field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCommunityCampaignResponse) GetEligibleForApplyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EligibleForApply.Get(), o.EligibleForApply.IsSet()
}

// HasEligibleForApply returns a boolean if a field has been set.
func (o *GroupCommunityCampaignResponse) HasEligibleForApply() bool {
	if o != nil && o.EligibleForApply.IsSet() {
		return true
	}

	return false
}

// SetEligibleForApply gets a reference to the given NullableBool and assigns it to the EligibleForApply field.
func (o *GroupCommunityCampaignResponse) SetEligibleForApply(v bool) {
	o.EligibleForApply.Set(&v)
}
// SetEligibleForApplyNil sets the value for EligibleForApply to be an explicit nil
func (o *GroupCommunityCampaignResponse) SetEligibleForApplyNil() {
	o.EligibleForApply.Set(nil)
}

// UnsetEligibleForApply ensures that no value is present for EligibleForApply, not even an explicit nil
func (o *GroupCommunityCampaignResponse) UnsetEligibleForApply() {
	o.EligibleForApply.Unset()
}

// GetViewableGroupLeaderboard returns the ViewableGroupLeaderboard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupCommunityCampaignResponse) GetViewableGroupLeaderboard() bool {
	if o == nil || IsNil(o.ViewableGroupLeaderboard.Get()) {
		var ret bool
		return ret
	}
	return *o.ViewableGroupLeaderboard.Get()
}

// GetViewableGroupLeaderboardOk returns a tuple with the ViewableGroupLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCommunityCampaignResponse) GetViewableGroupLeaderboardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewableGroupLeaderboard.Get(), o.ViewableGroupLeaderboard.IsSet()
}

// HasViewableGroupLeaderboard returns a boolean if a field has been set.
func (o *GroupCommunityCampaignResponse) HasViewableGroupLeaderboard() bool {
	if o != nil && o.ViewableGroupLeaderboard.IsSet() {
		return true
	}

	return false
}

// SetViewableGroupLeaderboard gets a reference to the given NullableBool and assigns it to the ViewableGroupLeaderboard field.
func (o *GroupCommunityCampaignResponse) SetViewableGroupLeaderboard(v bool) {
	o.ViewableGroupLeaderboard.Set(&v)
}
// SetViewableGroupLeaderboardNil sets the value for ViewableGroupLeaderboard to be an explicit nil
func (o *GroupCommunityCampaignResponse) SetViewableGroupLeaderboardNil() {
	o.ViewableGroupLeaderboard.Set(nil)
}

// UnsetViewableGroupLeaderboard ensures that no value is present for ViewableGroupLeaderboard, not even an explicit nil
func (o *GroupCommunityCampaignResponse) UnsetViewableGroupLeaderboard() {
	o.ViewableGroupLeaderboard.Unset()
}

func (o GroupCommunityCampaignResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupCommunityCampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EligibleForApply.IsSet() {
		toSerialize["eligible_for_apply"] = o.EligibleForApply.Get()
	}
	if o.ViewableGroupLeaderboard.IsSet() {
		toSerialize["viewable_group_leaderboard"] = o.ViewableGroupLeaderboard.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupCommunityCampaignResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupCommunityCampaignResponse := _GroupCommunityCampaignResponse{}

	err = json.Unmarshal(data, &varGroupCommunityCampaignResponse)

	if err != nil {
		return err
	}

	*o = GroupCommunityCampaignResponse(varGroupCommunityCampaignResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eligible_for_apply")
		delete(additionalProperties, "viewable_group_leaderboard")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupCommunityCampaignResponse struct {
	value *GroupCommunityCampaignResponse
	isSet bool
}

func (v NullableGroupCommunityCampaignResponse) Get() *GroupCommunityCampaignResponse {
	return v.value
}

func (v *NullableGroupCommunityCampaignResponse) Set(val *GroupCommunityCampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCommunityCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCommunityCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCommunityCampaignResponse(val *GroupCommunityCampaignResponse) *NullableGroupCommunityCampaignResponse {
	return &NullableGroupCommunityCampaignResponse{value: val, isSet: true}
}

func (v NullableGroupCommunityCampaignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCommunityCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


