
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSetting{}

// UserSetting struct for UserSetting
type UserSetting struct {
	NotificationChat NullableBool `json:"notification_chat,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSetting UserSetting

// NewUserSetting instantiates a new UserSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSetting() *UserSetting {
	this := UserSetting{}
	return &this
}

// NewUserSettingWithDefaults instantiates a new UserSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSettingWithDefaults() *UserSetting {
	this := UserSetting{}
	return &this
}

// GetNotificationChat returns the NotificationChat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSetting) GetNotificationChat() bool {
	if o == nil || IsNil(o.NotificationChat.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationChat.Get()
}

// GetNotificationChatOk returns a tuple with the NotificationChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSetting) GetNotificationChatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationChat.Get(), o.NotificationChat.IsSet()
}

// HasNotificationChat returns a boolean if a field has been set.
func (o *UserSetting) HasNotificationChat() bool {
	if o != nil && o.NotificationChat.IsSet() {
		return true
	}

	return false
}

// SetNotificationChat gets a reference to the given NullableBool and assigns it to the NotificationChat field.
func (o *UserSetting) SetNotificationChat(v bool) {
	o.NotificationChat.Set(&v)
}
// SetNotificationChatNil sets the value for NotificationChat to be an explicit nil
func (o *UserSetting) SetNotificationChatNil() {
	o.NotificationChat.Set(nil)
}

// UnsetNotificationChat ensures that no value is present for NotificationChat, not even an explicit nil
func (o *UserSetting) UnsetNotificationChat() {
	o.NotificationChat.Unset()
}

func (o UserSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationChat.IsSet() {
		toSerialize["notification_chat"] = o.NotificationChat.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSetting) UnmarshalJSON(data []byte) (err error) {
	varUserSetting := _UserSetting{}

	err = json.Unmarshal(data, &varUserSetting)

	if err != nil {
		return err
	}

	*o = UserSetting(varUserSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "notification_chat")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserSetting struct {
	value *UserSetting
	isSet bool
}

func (v NullableUserSetting) Get() *UserSetting {
	return v.value
}

func (v *NullableUserSetting) Set(val *UserSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSetting(val *UserSetting) *NullableUserSetting {
	return &NullableUserSetting{value: val, isSet: true}
}

func (v NullableUserSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


