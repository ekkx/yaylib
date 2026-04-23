
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserWrapper{}

// UserWrapper struct for UserWrapper
type UserWrapper struct {
	Id NullableInt64 `json:"id,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserWrapper UserWrapper

// NewUserWrapper instantiates a new UserWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserWrapper() *UserWrapper {
	this := UserWrapper{}
	return &this
}

// NewUserWrapperWithDefaults instantiates a new UserWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWrapperWithDefaults() *UserWrapper {
	this := UserWrapper{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserWrapper) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserWrapper) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UserWrapper) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *UserWrapper) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UserWrapper) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UserWrapper) UnsetId() {
	o.Id.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserWrapper) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserWrapper) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *UserWrapper) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *UserWrapper) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *UserWrapper) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *UserWrapper) UnsetUser() {
	o.User.Unset()
}

func (o UserWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserWrapper) UnmarshalJSON(data []byte) (err error) {
	varUserWrapper := _UserWrapper{}

	err = json.Unmarshal(data, &varUserWrapper)

	if err != nil {
		return err
	}

	*o = UserWrapper(varUserWrapper)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserWrapper struct {
	value *UserWrapper
	isSet bool
}

func (v NullableUserWrapper) Get() *UserWrapper {
	return v.value
}

func (v *NullableUserWrapper) Set(val *UserWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableUserWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableUserWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserWrapper(val *UserWrapper) *NullableUserWrapper {
	return &NullableUserWrapper{value: val, isSet: true}
}

func (v NullableUserWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


