
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the NotificationSettingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationSettingResponse{}

// NotificationSettingResponse struct for NotificationSettingResponse
type NotificationSettingResponse struct {
	Setting NullableUserSetting `json:"setting,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationSettingResponse NotificationSettingResponse

// NewNotificationSettingResponse instantiates a new NotificationSettingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSettingResponse() *NotificationSettingResponse {
	this := NotificationSettingResponse{}
	return &this
}

// NewNotificationSettingResponseWithDefaults instantiates a new NotificationSettingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSettingResponseWithDefaults() *NotificationSettingResponse {
	this := NotificationSettingResponse{}
	return &this
}

// GetSetting returns the Setting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationSettingResponse) GetSetting() UserSetting {
	if o == nil || IsNil(o.Setting.Get()) {
		var ret UserSetting
		return ret
	}
	return *o.Setting.Get()
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationSettingResponse) GetSettingOk() (*UserSetting, bool) {
	if o == nil {
		return nil, false
	}
	return o.Setting.Get(), o.Setting.IsSet()
}

// HasSetting returns a boolean if a field has been set.
func (o *NotificationSettingResponse) HasSetting() bool {
	if o != nil && o.Setting.IsSet() {
		return true
	}

	return false
}

// SetSetting gets a reference to the given NullableUserSetting and assigns it to the Setting field.
func (o *NotificationSettingResponse) SetSetting(v UserSetting) {
	o.Setting.Set(&v)
}
// SetSettingNil sets the value for Setting to be an explicit nil
func (o *NotificationSettingResponse) SetSettingNil() {
	o.Setting.Set(nil)
}

// UnsetSetting ensures that no value is present for Setting, not even an explicit nil
func (o *NotificationSettingResponse) UnsetSetting() {
	o.Setting.Unset()
}

func (o NotificationSettingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationSettingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Setting.IsSet() {
		toSerialize["setting"] = o.Setting.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationSettingResponse) UnmarshalJSON(data []byte) (err error) {
	varNotificationSettingResponse := _NotificationSettingResponse{}

	err = json.Unmarshal(data, &varNotificationSettingResponse)

	if err != nil {
		return err
	}

	*o = NotificationSettingResponse(varNotificationSettingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "setting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationSettingResponse struct {
	value *NotificationSettingResponse
	isSet bool
}

func (v NullableNotificationSettingResponse) Get() *NotificationSettingResponse {
	return v.value
}

func (v *NullableNotificationSettingResponse) Set(val *NotificationSettingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSettingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSettingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSettingResponse(val *NotificationSettingResponse) *NullableNotificationSettingResponse {
	return &NullableNotificationSettingResponse{value: val, isSet: true}
}

func (v NullableNotificationSettingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSettingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


