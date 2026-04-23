
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChatInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatInvitation{}

// ChatInvitation struct for ChatInvitation
type ChatInvitation struct {
	FirstUser NullableString `json:"first_user,omitempty"`
	Inviter NullableString `json:"inviter,omitempty"`
	UsersCount NullableInt32 `json:"users_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatInvitation ChatInvitation

// NewChatInvitation instantiates a new ChatInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatInvitation() *ChatInvitation {
	this := ChatInvitation{}
	return &this
}

// NewChatInvitationWithDefaults instantiates a new ChatInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatInvitationWithDefaults() *ChatInvitation {
	this := ChatInvitation{}
	return &this
}

// GetFirstUser returns the FirstUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatInvitation) GetFirstUser() string {
	if o == nil || IsNil(o.FirstUser.Get()) {
		var ret string
		return ret
	}
	return *o.FirstUser.Get()
}

// GetFirstUserOk returns a tuple with the FirstUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatInvitation) GetFirstUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstUser.Get(), o.FirstUser.IsSet()
}

// HasFirstUser returns a boolean if a field has been set.
func (o *ChatInvitation) HasFirstUser() bool {
	if o != nil && o.FirstUser.IsSet() {
		return true
	}

	return false
}

// SetFirstUser gets a reference to the given NullableString and assigns it to the FirstUser field.
func (o *ChatInvitation) SetFirstUser(v string) {
	o.FirstUser.Set(&v)
}
// SetFirstUserNil sets the value for FirstUser to be an explicit nil
func (o *ChatInvitation) SetFirstUserNil() {
	o.FirstUser.Set(nil)
}

// UnsetFirstUser ensures that no value is present for FirstUser, not even an explicit nil
func (o *ChatInvitation) UnsetFirstUser() {
	o.FirstUser.Unset()
}

// GetInviter returns the Inviter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatInvitation) GetInviter() string {
	if o == nil || IsNil(o.Inviter.Get()) {
		var ret string
		return ret
	}
	return *o.Inviter.Get()
}

// GetInviterOk returns a tuple with the Inviter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatInvitation) GetInviterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inviter.Get(), o.Inviter.IsSet()
}

// HasInviter returns a boolean if a field has been set.
func (o *ChatInvitation) HasInviter() bool {
	if o != nil && o.Inviter.IsSet() {
		return true
	}

	return false
}

// SetInviter gets a reference to the given NullableString and assigns it to the Inviter field.
func (o *ChatInvitation) SetInviter(v string) {
	o.Inviter.Set(&v)
}
// SetInviterNil sets the value for Inviter to be an explicit nil
func (o *ChatInvitation) SetInviterNil() {
	o.Inviter.Set(nil)
}

// UnsetInviter ensures that no value is present for Inviter, not even an explicit nil
func (o *ChatInvitation) UnsetInviter() {
	o.Inviter.Unset()
}

// GetUsersCount returns the UsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatInvitation) GetUsersCount() int32 {
	if o == nil || IsNil(o.UsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UsersCount.Get()
}

// GetUsersCountOk returns a tuple with the UsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatInvitation) GetUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsersCount.Get(), o.UsersCount.IsSet()
}

// HasUsersCount returns a boolean if a field has been set.
func (o *ChatInvitation) HasUsersCount() bool {
	if o != nil && o.UsersCount.IsSet() {
		return true
	}

	return false
}

// SetUsersCount gets a reference to the given NullableInt32 and assigns it to the UsersCount field.
func (o *ChatInvitation) SetUsersCount(v int32) {
	o.UsersCount.Set(&v)
}
// SetUsersCountNil sets the value for UsersCount to be an explicit nil
func (o *ChatInvitation) SetUsersCountNil() {
	o.UsersCount.Set(nil)
}

// UnsetUsersCount ensures that no value is present for UsersCount, not even an explicit nil
func (o *ChatInvitation) UnsetUsersCount() {
	o.UsersCount.Unset()
}

func (o ChatInvitation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstUser.IsSet() {
		toSerialize["first_user"] = o.FirstUser.Get()
	}
	if o.Inviter.IsSet() {
		toSerialize["inviter"] = o.Inviter.Get()
	}
	if o.UsersCount.IsSet() {
		toSerialize["users_count"] = o.UsersCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatInvitation) UnmarshalJSON(data []byte) (err error) {
	varChatInvitation := _ChatInvitation{}

	err = json.Unmarshal(data, &varChatInvitation)

	if err != nil {
		return err
	}

	*o = ChatInvitation(varChatInvitation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "first_user")
		delete(additionalProperties, "inviter")
		delete(additionalProperties, "users_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatInvitation struct {
	value *ChatInvitation
	isSet bool
}

func (v NullableChatInvitation) Get() *ChatInvitation {
	return v.value
}

func (v *NullableChatInvitation) Set(val *ChatInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableChatInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableChatInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatInvitation(val *ChatInvitation) *NullableChatInvitation {
	return &NullableChatInvitation{value: val, isSet: true}
}

func (v NullableChatInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


