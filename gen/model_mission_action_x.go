
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MissionActionX type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionActionX{}

// MissionActionX struct for MissionActionX
type MissionActionX struct {
	DescriptionId NullableInt32 `json:"description_id,omitempty"`
	RedirectButtonId NullableInt32 `json:"redirect_button_id,omitempty"`
	RewardDetailsId NullableInt32 `json:"reward_details_id,omitempty"`
	TitleId NullableInt32 `json:"title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionActionX MissionActionX

// NewMissionActionX instantiates a new MissionActionX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionActionX() *MissionActionX {
	this := MissionActionX{}
	return &this
}

// NewMissionActionXWithDefaults instantiates a new MissionActionX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionActionXWithDefaults() *MissionActionX {
	this := MissionActionX{}
	return &this
}

// GetDescriptionId returns the DescriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionActionX) GetDescriptionId() int32 {
	if o == nil || IsNil(o.DescriptionId.Get()) {
		var ret int32
		return ret
	}
	return *o.DescriptionId.Get()
}

// GetDescriptionIdOk returns a tuple with the DescriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionActionX) GetDescriptionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DescriptionId.Get(), o.DescriptionId.IsSet()
}

// HasDescriptionId returns a boolean if a field has been set.
func (o *MissionActionX) HasDescriptionId() bool {
	if o != nil && o.DescriptionId.IsSet() {
		return true
	}

	return false
}

// SetDescriptionId gets a reference to the given NullableInt32 and assigns it to the DescriptionId field.
func (o *MissionActionX) SetDescriptionId(v int32) {
	o.DescriptionId.Set(&v)
}
// SetDescriptionIdNil sets the value for DescriptionId to be an explicit nil
func (o *MissionActionX) SetDescriptionIdNil() {
	o.DescriptionId.Set(nil)
}

// UnsetDescriptionId ensures that no value is present for DescriptionId, not even an explicit nil
func (o *MissionActionX) UnsetDescriptionId() {
	o.DescriptionId.Unset()
}

// GetRedirectButtonId returns the RedirectButtonId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionActionX) GetRedirectButtonId() int32 {
	if o == nil || IsNil(o.RedirectButtonId.Get()) {
		var ret int32
		return ret
	}
	return *o.RedirectButtonId.Get()
}

// GetRedirectButtonIdOk returns a tuple with the RedirectButtonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionActionX) GetRedirectButtonIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectButtonId.Get(), o.RedirectButtonId.IsSet()
}

// HasRedirectButtonId returns a boolean if a field has been set.
func (o *MissionActionX) HasRedirectButtonId() bool {
	if o != nil && o.RedirectButtonId.IsSet() {
		return true
	}

	return false
}

// SetRedirectButtonId gets a reference to the given NullableInt32 and assigns it to the RedirectButtonId field.
func (o *MissionActionX) SetRedirectButtonId(v int32) {
	o.RedirectButtonId.Set(&v)
}
// SetRedirectButtonIdNil sets the value for RedirectButtonId to be an explicit nil
func (o *MissionActionX) SetRedirectButtonIdNil() {
	o.RedirectButtonId.Set(nil)
}

// UnsetRedirectButtonId ensures that no value is present for RedirectButtonId, not even an explicit nil
func (o *MissionActionX) UnsetRedirectButtonId() {
	o.RedirectButtonId.Unset()
}

// GetRewardDetailsId returns the RewardDetailsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionActionX) GetRewardDetailsId() int32 {
	if o == nil || IsNil(o.RewardDetailsId.Get()) {
		var ret int32
		return ret
	}
	return *o.RewardDetailsId.Get()
}

// GetRewardDetailsIdOk returns a tuple with the RewardDetailsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionActionX) GetRewardDetailsIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardDetailsId.Get(), o.RewardDetailsId.IsSet()
}

// HasRewardDetailsId returns a boolean if a field has been set.
func (o *MissionActionX) HasRewardDetailsId() bool {
	if o != nil && o.RewardDetailsId.IsSet() {
		return true
	}

	return false
}

// SetRewardDetailsId gets a reference to the given NullableInt32 and assigns it to the RewardDetailsId field.
func (o *MissionActionX) SetRewardDetailsId(v int32) {
	o.RewardDetailsId.Set(&v)
}
// SetRewardDetailsIdNil sets the value for RewardDetailsId to be an explicit nil
func (o *MissionActionX) SetRewardDetailsIdNil() {
	o.RewardDetailsId.Set(nil)
}

// UnsetRewardDetailsId ensures that no value is present for RewardDetailsId, not even an explicit nil
func (o *MissionActionX) UnsetRewardDetailsId() {
	o.RewardDetailsId.Unset()
}

// GetTitleId returns the TitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionActionX) GetTitleId() int32 {
	if o == nil || IsNil(o.TitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.TitleId.Get()
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionActionX) GetTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleId.Get(), o.TitleId.IsSet()
}

// HasTitleId returns a boolean if a field has been set.
func (o *MissionActionX) HasTitleId() bool {
	if o != nil && o.TitleId.IsSet() {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given NullableInt32 and assigns it to the TitleId field.
func (o *MissionActionX) SetTitleId(v int32) {
	o.TitleId.Set(&v)
}
// SetTitleIdNil sets the value for TitleId to be an explicit nil
func (o *MissionActionX) SetTitleIdNil() {
	o.TitleId.Set(nil)
}

// UnsetTitleId ensures that no value is present for TitleId, not even an explicit nil
func (o *MissionActionX) UnsetTitleId() {
	o.TitleId.Unset()
}

func (o MissionActionX) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionActionX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DescriptionId.IsSet() {
		toSerialize["description_id"] = o.DescriptionId.Get()
	}
	if o.RedirectButtonId.IsSet() {
		toSerialize["redirect_button_id"] = o.RedirectButtonId.Get()
	}
	if o.RewardDetailsId.IsSet() {
		toSerialize["reward_details_id"] = o.RewardDetailsId.Get()
	}
	if o.TitleId.IsSet() {
		toSerialize["title_id"] = o.TitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionActionX) UnmarshalJSON(data []byte) (err error) {
	varMissionActionX := _MissionActionX{}

	err = json.Unmarshal(data, &varMissionActionX)

	if err != nil {
		return err
	}

	*o = MissionActionX(varMissionActionX)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description_id")
		delete(additionalProperties, "redirect_button_id")
		delete(additionalProperties, "reward_details_id")
		delete(additionalProperties, "title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionActionX struct {
	value *MissionActionX
	isSet bool
}

func (v NullableMissionActionX) Get() *MissionActionX {
	return v.value
}

func (v *NullableMissionActionX) Set(val *MissionActionX) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionActionX) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionActionX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionActionX(val *MissionActionX) *NullableMissionActionX {
	return &NullableMissionActionX{value: val, isSet: true}
}

func (v NullableMissionActionX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionActionX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


