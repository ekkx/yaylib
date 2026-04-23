
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Role{}

// Role struct for Role
type Role struct {
	IsModerator NullableBool `json:"is_moderator,omitempty"`
	IsOwner NullableBool `json:"is_owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Role Role

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetIsModerator returns the IsModerator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Role) GetIsModerator() bool {
	if o == nil || IsNil(o.IsModerator.Get()) {
		var ret bool
		return ret
	}
	return *o.IsModerator.Get()
}

// GetIsModeratorOk returns a tuple with the IsModerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Role) GetIsModeratorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsModerator.Get(), o.IsModerator.IsSet()
}

// HasIsModerator returns a boolean if a field has been set.
func (o *Role) HasIsModerator() bool {
	if o != nil && o.IsModerator.IsSet() {
		return true
	}

	return false
}

// SetIsModerator gets a reference to the given NullableBool and assigns it to the IsModerator field.
func (o *Role) SetIsModerator(v bool) {
	o.IsModerator.Set(&v)
}
// SetIsModeratorNil sets the value for IsModerator to be an explicit nil
func (o *Role) SetIsModeratorNil() {
	o.IsModerator.Set(nil)
}

// UnsetIsModerator ensures that no value is present for IsModerator, not even an explicit nil
func (o *Role) UnsetIsModerator() {
	o.IsModerator.Unset()
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Role) GetIsOwner() bool {
	if o == nil || IsNil(o.IsOwner.Get()) {
		var ret bool
		return ret
	}
	return *o.IsOwner.Get()
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Role) GetIsOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsOwner.Get(), o.IsOwner.IsSet()
}

// HasIsOwner returns a boolean if a field has been set.
func (o *Role) HasIsOwner() bool {
	if o != nil && o.IsOwner.IsSet() {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given NullableBool and assigns it to the IsOwner field.
func (o *Role) SetIsOwner(v bool) {
	o.IsOwner.Set(&v)
}
// SetIsOwnerNil sets the value for IsOwner to be an explicit nil
func (o *Role) SetIsOwnerNil() {
	o.IsOwner.Set(nil)
}

// UnsetIsOwner ensures that no value is present for IsOwner, not even an explicit nil
func (o *Role) UnsetIsOwner() {
	o.IsOwner.Unset()
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsModerator.IsSet() {
		toSerialize["is_moderator"] = o.IsModerator.Get()
	}
	if o.IsOwner.IsSet() {
		toSerialize["is_owner"] = o.IsOwner.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Role) UnmarshalJSON(data []byte) (err error) {
	varRole := _Role{}

	err = json.Unmarshal(data, &varRole)

	if err != nil {
		return err
	}

	*o = Role(varRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_moderator")
		delete(additionalProperties, "is_owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


