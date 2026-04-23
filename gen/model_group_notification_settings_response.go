
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupNotificationSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupNotificationSettingsResponse{}

// GroupNotificationSettingsResponse struct for GroupNotificationSettingsResponse
type GroupNotificationSettingsResponse struct {
	Setting NullableSetting `json:"setting,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupNotificationSettingsResponse GroupNotificationSettingsResponse

// NewGroupNotificationSettingsResponse instantiates a new GroupNotificationSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupNotificationSettingsResponse() *GroupNotificationSettingsResponse {
	this := GroupNotificationSettingsResponse{}
	return &this
}

// NewGroupNotificationSettingsResponseWithDefaults instantiates a new GroupNotificationSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupNotificationSettingsResponseWithDefaults() *GroupNotificationSettingsResponse {
	this := GroupNotificationSettingsResponse{}
	return &this
}

// GetSetting returns the Setting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupNotificationSettingsResponse) GetSetting() Setting {
	if o == nil || IsNil(o.Setting.Get()) {
		var ret Setting
		return ret
	}
	return *o.Setting.Get()
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupNotificationSettingsResponse) GetSettingOk() (*Setting, bool) {
	if o == nil {
		return nil, false
	}
	return o.Setting.Get(), o.Setting.IsSet()
}

// HasSetting returns a boolean if a field has been set.
func (o *GroupNotificationSettingsResponse) HasSetting() bool {
	if o != nil && o.Setting.IsSet() {
		return true
	}

	return false
}

// SetSetting gets a reference to the given NullableSetting and assigns it to the Setting field.
func (o *GroupNotificationSettingsResponse) SetSetting(v Setting) {
	o.Setting.Set(&v)
}
// SetSettingNil sets the value for Setting to be an explicit nil
func (o *GroupNotificationSettingsResponse) SetSettingNil() {
	o.Setting.Set(nil)
}

// UnsetSetting ensures that no value is present for Setting, not even an explicit nil
func (o *GroupNotificationSettingsResponse) UnsetSetting() {
	o.Setting.Unset()
}

func (o GroupNotificationSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupNotificationSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Setting.IsSet() {
		toSerialize["setting"] = o.Setting.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupNotificationSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupNotificationSettingsResponse := _GroupNotificationSettingsResponse{}

	err = json.Unmarshal(data, &varGroupNotificationSettingsResponse)

	if err != nil {
		return err
	}

	*o = GroupNotificationSettingsResponse(varGroupNotificationSettingsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "setting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupNotificationSettingsResponse struct {
	value *GroupNotificationSettingsResponse
	isSet bool
}

func (v NullableGroupNotificationSettingsResponse) Get() *GroupNotificationSettingsResponse {
	return v.value
}

func (v *NullableGroupNotificationSettingsResponse) Set(val *GroupNotificationSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupNotificationSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupNotificationSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupNotificationSettingsResponse(val *GroupNotificationSettingsResponse) *NullableGroupNotificationSettingsResponse {
	return &NullableGroupNotificationSettingsResponse{value: val, isSet: true}
}

func (v NullableGroupNotificationSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupNotificationSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


