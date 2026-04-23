
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ConferenceCallUserRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceCallUserRole{}

// ConferenceCallUserRole struct for ConferenceCallUserRole
type ConferenceCallUserRole struct {
	Id NullableInt64 `json:"id,omitempty"`
	Role NullableString `json:"role,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConferenceCallUserRole ConferenceCallUserRole

// NewConferenceCallUserRole instantiates a new ConferenceCallUserRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceCallUserRole() *ConferenceCallUserRole {
	this := ConferenceCallUserRole{}
	return &this
}

// NewConferenceCallUserRoleWithDefaults instantiates a new ConferenceCallUserRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceCallUserRoleWithDefaults() *ConferenceCallUserRole {
	this := ConferenceCallUserRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCallUserRole) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCallUserRole) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ConferenceCallUserRole) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ConferenceCallUserRole) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ConferenceCallUserRole) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ConferenceCallUserRole) UnsetId() {
	o.Id.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCallUserRole) GetRole() string {
	if o == nil || IsNil(o.Role.Get()) {
		var ret string
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCallUserRole) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *ConferenceCallUserRole) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableString and assigns it to the Role field.
func (o *ConferenceCallUserRole) SetRole(v string) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *ConferenceCallUserRole) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *ConferenceCallUserRole) UnsetRole() {
	o.Role.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCallUserRole) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCallUserRole) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ConferenceCallUserRole) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *ConferenceCallUserRole) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ConferenceCallUserRole) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ConferenceCallUserRole) UnsetUserId() {
	o.UserId.Unset()
}

func (o ConferenceCallUserRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceCallUserRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConferenceCallUserRole) UnmarshalJSON(data []byte) (err error) {
	varConferenceCallUserRole := _ConferenceCallUserRole{}

	err = json.Unmarshal(data, &varConferenceCallUserRole)

	if err != nil {
		return err
	}

	*o = ConferenceCallUserRole(varConferenceCallUserRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "role")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConferenceCallUserRole struct {
	value *ConferenceCallUserRole
	isSet bool
}

func (v NullableConferenceCallUserRole) Get() *ConferenceCallUserRole {
	return v.value
}

func (v *NullableConferenceCallUserRole) Set(val *ConferenceCallUserRole) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceCallUserRole) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceCallUserRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceCallUserRole(val *ConferenceCallUserRole) *NullableConferenceCallUserRole {
	return &NullableConferenceCallUserRole{value: val, isSet: true}
}

func (v NullableConferenceCallUserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceCallUserRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


