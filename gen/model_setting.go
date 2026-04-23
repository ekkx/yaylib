
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Setting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Setting{}

// Setting struct for Setting
type Setting struct {
	NotificationGroupJoin NullableBool `json:"notification_group_join,omitempty"`
	NotificationGroupMessageTagAll NullableBool `json:"notification_group_message_tag_all,omitempty"`
	NotificationGroupPost NullableBool `json:"notification_group_post,omitempty"`
	NotificationGroupRequest NullableBool `json:"notification_group_request,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Setting Setting

// NewSetting instantiates a new Setting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetting() *Setting {
	this := Setting{}
	return &this
}

// NewSettingWithDefaults instantiates a new Setting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingWithDefaults() *Setting {
	this := Setting{}
	return &this
}

// GetNotificationGroupJoin returns the NotificationGroupJoin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Setting) GetNotificationGroupJoin() bool {
	if o == nil || IsNil(o.NotificationGroupJoin.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupJoin.Get()
}

// GetNotificationGroupJoinOk returns a tuple with the NotificationGroupJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Setting) GetNotificationGroupJoinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupJoin.Get(), o.NotificationGroupJoin.IsSet()
}

// HasNotificationGroupJoin returns a boolean if a field has been set.
func (o *Setting) HasNotificationGroupJoin() bool {
	if o != nil && o.NotificationGroupJoin.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupJoin gets a reference to the given NullableBool and assigns it to the NotificationGroupJoin field.
func (o *Setting) SetNotificationGroupJoin(v bool) {
	o.NotificationGroupJoin.Set(&v)
}
// SetNotificationGroupJoinNil sets the value for NotificationGroupJoin to be an explicit nil
func (o *Setting) SetNotificationGroupJoinNil() {
	o.NotificationGroupJoin.Set(nil)
}

// UnsetNotificationGroupJoin ensures that no value is present for NotificationGroupJoin, not even an explicit nil
func (o *Setting) UnsetNotificationGroupJoin() {
	o.NotificationGroupJoin.Unset()
}

// GetNotificationGroupMessageTagAll returns the NotificationGroupMessageTagAll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Setting) GetNotificationGroupMessageTagAll() bool {
	if o == nil || IsNil(o.NotificationGroupMessageTagAll.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupMessageTagAll.Get()
}

// GetNotificationGroupMessageTagAllOk returns a tuple with the NotificationGroupMessageTagAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Setting) GetNotificationGroupMessageTagAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupMessageTagAll.Get(), o.NotificationGroupMessageTagAll.IsSet()
}

// HasNotificationGroupMessageTagAll returns a boolean if a field has been set.
func (o *Setting) HasNotificationGroupMessageTagAll() bool {
	if o != nil && o.NotificationGroupMessageTagAll.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupMessageTagAll gets a reference to the given NullableBool and assigns it to the NotificationGroupMessageTagAll field.
func (o *Setting) SetNotificationGroupMessageTagAll(v bool) {
	o.NotificationGroupMessageTagAll.Set(&v)
}
// SetNotificationGroupMessageTagAllNil sets the value for NotificationGroupMessageTagAll to be an explicit nil
func (o *Setting) SetNotificationGroupMessageTagAllNil() {
	o.NotificationGroupMessageTagAll.Set(nil)
}

// UnsetNotificationGroupMessageTagAll ensures that no value is present for NotificationGroupMessageTagAll, not even an explicit nil
func (o *Setting) UnsetNotificationGroupMessageTagAll() {
	o.NotificationGroupMessageTagAll.Unset()
}

// GetNotificationGroupPost returns the NotificationGroupPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Setting) GetNotificationGroupPost() bool {
	if o == nil || IsNil(o.NotificationGroupPost.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupPost.Get()
}

// GetNotificationGroupPostOk returns a tuple with the NotificationGroupPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Setting) GetNotificationGroupPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupPost.Get(), o.NotificationGroupPost.IsSet()
}

// HasNotificationGroupPost returns a boolean if a field has been set.
func (o *Setting) HasNotificationGroupPost() bool {
	if o != nil && o.NotificationGroupPost.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupPost gets a reference to the given NullableBool and assigns it to the NotificationGroupPost field.
func (o *Setting) SetNotificationGroupPost(v bool) {
	o.NotificationGroupPost.Set(&v)
}
// SetNotificationGroupPostNil sets the value for NotificationGroupPost to be an explicit nil
func (o *Setting) SetNotificationGroupPostNil() {
	o.NotificationGroupPost.Set(nil)
}

// UnsetNotificationGroupPost ensures that no value is present for NotificationGroupPost, not even an explicit nil
func (o *Setting) UnsetNotificationGroupPost() {
	o.NotificationGroupPost.Unset()
}

// GetNotificationGroupRequest returns the NotificationGroupRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Setting) GetNotificationGroupRequest() bool {
	if o == nil || IsNil(o.NotificationGroupRequest.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupRequest.Get()
}

// GetNotificationGroupRequestOk returns a tuple with the NotificationGroupRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Setting) GetNotificationGroupRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupRequest.Get(), o.NotificationGroupRequest.IsSet()
}

// HasNotificationGroupRequest returns a boolean if a field has been set.
func (o *Setting) HasNotificationGroupRequest() bool {
	if o != nil && o.NotificationGroupRequest.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupRequest gets a reference to the given NullableBool and assigns it to the NotificationGroupRequest field.
func (o *Setting) SetNotificationGroupRequest(v bool) {
	o.NotificationGroupRequest.Set(&v)
}
// SetNotificationGroupRequestNil sets the value for NotificationGroupRequest to be an explicit nil
func (o *Setting) SetNotificationGroupRequestNil() {
	o.NotificationGroupRequest.Set(nil)
}

// UnsetNotificationGroupRequest ensures that no value is present for NotificationGroupRequest, not even an explicit nil
func (o *Setting) UnsetNotificationGroupRequest() {
	o.NotificationGroupRequest.Unset()
}

func (o Setting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Setting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NotificationGroupJoin.IsSet() {
		toSerialize["notification_group_join"] = o.NotificationGroupJoin.Get()
	}
	if o.NotificationGroupMessageTagAll.IsSet() {
		toSerialize["notification_group_message_tag_all"] = o.NotificationGroupMessageTagAll.Get()
	}
	if o.NotificationGroupPost.IsSet() {
		toSerialize["notification_group_post"] = o.NotificationGroupPost.Get()
	}
	if o.NotificationGroupRequest.IsSet() {
		toSerialize["notification_group_request"] = o.NotificationGroupRequest.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Setting) UnmarshalJSON(data []byte) (err error) {
	varSetting := _Setting{}

	err = json.Unmarshal(data, &varSetting)

	if err != nil {
		return err
	}

	*o = Setting(varSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "notification_group_join")
		delete(additionalProperties, "notification_group_message_tag_all")
		delete(additionalProperties, "notification_group_post")
		delete(additionalProperties, "notification_group_request")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetting struct {
	value *Setting
	isSet bool
}

func (v NullableSetting) Get() *Setting {
	return v.value
}

func (v *NullableSetting) Set(val *Setting) {
	v.value = val
	v.isSet = true
}

func (v NullableSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetting(val *Setting) *NullableSetting {
	return &NullableSetting{value: val, isSet: true}
}

func (v NullableSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


