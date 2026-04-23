
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserMuted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMuted{}

// UserMuted struct for UserMuted
type UserMuted struct {
	IsMuted NullableBool `json:"is_muted,omitempty"`
	User NullableUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserMuted UserMuted

// NewUserMuted instantiates a new UserMuted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMuted() *UserMuted {
	this := UserMuted{}
	return &this
}

// NewUserMutedWithDefaults instantiates a new UserMuted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMutedWithDefaults() *UserMuted {
	this := UserMuted{}
	return &this
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserMuted) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMuted.Get()
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserMuted) GetIsMutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMuted.Get(), o.IsMuted.IsSet()
}

// HasIsMuted returns a boolean if a field has been set.
func (o *UserMuted) HasIsMuted() bool {
	if o != nil && o.IsMuted.IsSet() {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given NullableBool and assigns it to the IsMuted field.
func (o *UserMuted) SetIsMuted(v bool) {
	o.IsMuted.Set(&v)
}
// SetIsMutedNil sets the value for IsMuted to be an explicit nil
func (o *UserMuted) SetIsMutedNil() {
	o.IsMuted.Set(nil)
}

// UnsetIsMuted ensures that no value is present for IsMuted, not even an explicit nil
func (o *UserMuted) UnsetIsMuted() {
	o.IsMuted.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserMuted) GetUser() User {
	if o == nil || IsNil(o.User.Get()) {
		var ret User
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserMuted) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *UserMuted) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUser and assigns it to the User field.
func (o *UserMuted) SetUser(v User) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *UserMuted) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *UserMuted) UnsetUser() {
	o.User.Unset()
}

func (o UserMuted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMuted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsMuted.IsSet() {
		toSerialize["is_muted"] = o.IsMuted.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserMuted) UnmarshalJSON(data []byte) (err error) {
	varUserMuted := _UserMuted{}

	err = json.Unmarshal(data, &varUserMuted)

	if err != nil {
		return err
	}

	*o = UserMuted(varUserMuted)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_muted")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserMuted struct {
	value *UserMuted
	isSet bool
}

func (v NullableUserMuted) Get() *UserMuted {
	return v.value
}

func (v *NullableUserMuted) Set(val *UserMuted) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMuted) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMuted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMuted(val *UserMuted) *NullableUserMuted {
	return &NullableUserMuted{value: val, isSet: true}
}

func (v NullableUserMuted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMuted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


