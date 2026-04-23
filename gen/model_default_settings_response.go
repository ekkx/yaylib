
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the DefaultSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultSettingsResponse{}

// DefaultSettingsResponse struct for DefaultSettingsResponse
type DefaultSettingsResponse struct {
	TimelineSettings NullableTimelineSettings `json:"timeline_settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DefaultSettingsResponse DefaultSettingsResponse

// NewDefaultSettingsResponse instantiates a new DefaultSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultSettingsResponse() *DefaultSettingsResponse {
	this := DefaultSettingsResponse{}
	return &this
}

// NewDefaultSettingsResponseWithDefaults instantiates a new DefaultSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultSettingsResponseWithDefaults() *DefaultSettingsResponse {
	this := DefaultSettingsResponse{}
	return &this
}

// GetTimelineSettings returns the TimelineSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultSettingsResponse) GetTimelineSettings() TimelineSettings {
	if o == nil || IsNil(o.TimelineSettings.Get()) {
		var ret TimelineSettings
		return ret
	}
	return *o.TimelineSettings.Get()
}

// GetTimelineSettingsOk returns a tuple with the TimelineSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultSettingsResponse) GetTimelineSettingsOk() (*TimelineSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimelineSettings.Get(), o.TimelineSettings.IsSet()
}

// HasTimelineSettings returns a boolean if a field has been set.
func (o *DefaultSettingsResponse) HasTimelineSettings() bool {
	if o != nil && o.TimelineSettings.IsSet() {
		return true
	}

	return false
}

// SetTimelineSettings gets a reference to the given NullableTimelineSettings and assigns it to the TimelineSettings field.
func (o *DefaultSettingsResponse) SetTimelineSettings(v TimelineSettings) {
	o.TimelineSettings.Set(&v)
}
// SetTimelineSettingsNil sets the value for TimelineSettings to be an explicit nil
func (o *DefaultSettingsResponse) SetTimelineSettingsNil() {
	o.TimelineSettings.Set(nil)
}

// UnsetTimelineSettings ensures that no value is present for TimelineSettings, not even an explicit nil
func (o *DefaultSettingsResponse) UnsetTimelineSettings() {
	o.TimelineSettings.Unset()
}

func (o DefaultSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TimelineSettings.IsSet() {
		toSerialize["timeline_settings"] = o.TimelineSettings.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DefaultSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	varDefaultSettingsResponse := _DefaultSettingsResponse{}

	err = json.Unmarshal(data, &varDefaultSettingsResponse)

	if err != nil {
		return err
	}

	*o = DefaultSettingsResponse(varDefaultSettingsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timeline_settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDefaultSettingsResponse struct {
	value *DefaultSettingsResponse
	isSet bool
}

func (v NullableDefaultSettingsResponse) Get() *DefaultSettingsResponse {
	return v.value
}

func (v *NullableDefaultSettingsResponse) Set(val *DefaultSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultSettingsResponse(val *DefaultSettingsResponse) *NullableDefaultSettingsResponse {
	return &NullableDefaultSettingsResponse{value: val, isSet: true}
}

func (v NullableDefaultSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


