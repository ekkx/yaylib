
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TimelineSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineSettings{}

// TimelineSettings struct for TimelineSettings
type TimelineSettings struct {
	FavesFilter NullableBool `json:"faves_filter,omitempty"`
	HideHotPost NullableBool `json:"hide_hot_post,omitempty"`
	HideReplyFollowingTimeline NullableBool `json:"hide_reply_following_timeline,omitempty"`
	HideReplyPublicTimeline NullableBool `json:"hide_reply_public_timeline,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimelineSettings TimelineSettings

// NewTimelineSettings instantiates a new TimelineSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelineSettings() *TimelineSettings {
	this := TimelineSettings{}
	return &this
}

// NewTimelineSettingsWithDefaults instantiates a new TimelineSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelineSettingsWithDefaults() *TimelineSettings {
	this := TimelineSettings{}
	return &this
}

// GetFavesFilter returns the FavesFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineSettings) GetFavesFilter() bool {
	if o == nil || IsNil(o.FavesFilter.Get()) {
		var ret bool
		return ret
	}
	return *o.FavesFilter.Get()
}

// GetFavesFilterOk returns a tuple with the FavesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineSettings) GetFavesFilterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FavesFilter.Get(), o.FavesFilter.IsSet()
}

// HasFavesFilter returns a boolean if a field has been set.
func (o *TimelineSettings) HasFavesFilter() bool {
	if o != nil && o.FavesFilter.IsSet() {
		return true
	}

	return false
}

// SetFavesFilter gets a reference to the given NullableBool and assigns it to the FavesFilter field.
func (o *TimelineSettings) SetFavesFilter(v bool) {
	o.FavesFilter.Set(&v)
}
// SetFavesFilterNil sets the value for FavesFilter to be an explicit nil
func (o *TimelineSettings) SetFavesFilterNil() {
	o.FavesFilter.Set(nil)
}

// UnsetFavesFilter ensures that no value is present for FavesFilter, not even an explicit nil
func (o *TimelineSettings) UnsetFavesFilter() {
	o.FavesFilter.Unset()
}

// GetHideHotPost returns the HideHotPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineSettings) GetHideHotPost() bool {
	if o == nil || IsNil(o.HideHotPost.Get()) {
		var ret bool
		return ret
	}
	return *o.HideHotPost.Get()
}

// GetHideHotPostOk returns a tuple with the HideHotPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineSettings) GetHideHotPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideHotPost.Get(), o.HideHotPost.IsSet()
}

// HasHideHotPost returns a boolean if a field has been set.
func (o *TimelineSettings) HasHideHotPost() bool {
	if o != nil && o.HideHotPost.IsSet() {
		return true
	}

	return false
}

// SetHideHotPost gets a reference to the given NullableBool and assigns it to the HideHotPost field.
func (o *TimelineSettings) SetHideHotPost(v bool) {
	o.HideHotPost.Set(&v)
}
// SetHideHotPostNil sets the value for HideHotPost to be an explicit nil
func (o *TimelineSettings) SetHideHotPostNil() {
	o.HideHotPost.Set(nil)
}

// UnsetHideHotPost ensures that no value is present for HideHotPost, not even an explicit nil
func (o *TimelineSettings) UnsetHideHotPost() {
	o.HideHotPost.Unset()
}

// GetHideReplyFollowingTimeline returns the HideReplyFollowingTimeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineSettings) GetHideReplyFollowingTimeline() bool {
	if o == nil || IsNil(o.HideReplyFollowingTimeline.Get()) {
		var ret bool
		return ret
	}
	return *o.HideReplyFollowingTimeline.Get()
}

// GetHideReplyFollowingTimelineOk returns a tuple with the HideReplyFollowingTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineSettings) GetHideReplyFollowingTimelineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideReplyFollowingTimeline.Get(), o.HideReplyFollowingTimeline.IsSet()
}

// HasHideReplyFollowingTimeline returns a boolean if a field has been set.
func (o *TimelineSettings) HasHideReplyFollowingTimeline() bool {
	if o != nil && o.HideReplyFollowingTimeline.IsSet() {
		return true
	}

	return false
}

// SetHideReplyFollowingTimeline gets a reference to the given NullableBool and assigns it to the HideReplyFollowingTimeline field.
func (o *TimelineSettings) SetHideReplyFollowingTimeline(v bool) {
	o.HideReplyFollowingTimeline.Set(&v)
}
// SetHideReplyFollowingTimelineNil sets the value for HideReplyFollowingTimeline to be an explicit nil
func (o *TimelineSettings) SetHideReplyFollowingTimelineNil() {
	o.HideReplyFollowingTimeline.Set(nil)
}

// UnsetHideReplyFollowingTimeline ensures that no value is present for HideReplyFollowingTimeline, not even an explicit nil
func (o *TimelineSettings) UnsetHideReplyFollowingTimeline() {
	o.HideReplyFollowingTimeline.Unset()
}

// GetHideReplyPublicTimeline returns the HideReplyPublicTimeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelineSettings) GetHideReplyPublicTimeline() bool {
	if o == nil || IsNil(o.HideReplyPublicTimeline.Get()) {
		var ret bool
		return ret
	}
	return *o.HideReplyPublicTimeline.Get()
}

// GetHideReplyPublicTimelineOk returns a tuple with the HideReplyPublicTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelineSettings) GetHideReplyPublicTimelineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideReplyPublicTimeline.Get(), o.HideReplyPublicTimeline.IsSet()
}

// HasHideReplyPublicTimeline returns a boolean if a field has been set.
func (o *TimelineSettings) HasHideReplyPublicTimeline() bool {
	if o != nil && o.HideReplyPublicTimeline.IsSet() {
		return true
	}

	return false
}

// SetHideReplyPublicTimeline gets a reference to the given NullableBool and assigns it to the HideReplyPublicTimeline field.
func (o *TimelineSettings) SetHideReplyPublicTimeline(v bool) {
	o.HideReplyPublicTimeline.Set(&v)
}
// SetHideReplyPublicTimelineNil sets the value for HideReplyPublicTimeline to be an explicit nil
func (o *TimelineSettings) SetHideReplyPublicTimelineNil() {
	o.HideReplyPublicTimeline.Set(nil)
}

// UnsetHideReplyPublicTimeline ensures that no value is present for HideReplyPublicTimeline, not even an explicit nil
func (o *TimelineSettings) UnsetHideReplyPublicTimeline() {
	o.HideReplyPublicTimeline.Unset()
}

func (o TimelineSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelineSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FavesFilter.IsSet() {
		toSerialize["faves_filter"] = o.FavesFilter.Get()
	}
	if o.HideHotPost.IsSet() {
		toSerialize["hide_hot_post"] = o.HideHotPost.Get()
	}
	if o.HideReplyFollowingTimeline.IsSet() {
		toSerialize["hide_reply_following_timeline"] = o.HideReplyFollowingTimeline.Get()
	}
	if o.HideReplyPublicTimeline.IsSet() {
		toSerialize["hide_reply_public_timeline"] = o.HideReplyPublicTimeline.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimelineSettings) UnmarshalJSON(data []byte) (err error) {
	varTimelineSettings := _TimelineSettings{}

	err = json.Unmarshal(data, &varTimelineSettings)

	if err != nil {
		return err
	}

	*o = TimelineSettings(varTimelineSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "faves_filter")
		delete(additionalProperties, "hide_hot_post")
		delete(additionalProperties, "hide_reply_following_timeline")
		delete(additionalProperties, "hide_reply_public_timeline")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimelineSettings struct {
	value *TimelineSettings
	isSet bool
}

func (v NullableTimelineSettings) Get() *TimelineSettings {
	return v.value
}

func (v *NullableTimelineSettings) Set(val *TimelineSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineSettings(val *TimelineSettings) *NullableTimelineSettings {
	return &NullableTimelineSettings{value: val, isSet: true}
}

func (v NullableTimelineSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


