
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelConferenceCallUserRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelConferenceCallUserRole{}

// ModelConferenceCallUserRole struct for ModelConferenceCallUserRole
type ModelConferenceCallUserRole struct {
	Id NullableInt64 `json:"id,omitempty"`
	// Unresolved Java type: jp.nanameue.agora.f.a
	Role map[string]interface{} `json:"role,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelConferenceCallUserRole ModelConferenceCallUserRole

// NewModelConferenceCallUserRole instantiates a new ModelConferenceCallUserRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelConferenceCallUserRole() *ModelConferenceCallUserRole {
	this := ModelConferenceCallUserRole{}
	return &this
}

// NewModelConferenceCallUserRoleWithDefaults instantiates a new ModelConferenceCallUserRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelConferenceCallUserRoleWithDefaults() *ModelConferenceCallUserRole {
	this := ModelConferenceCallUserRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelConferenceCallUserRole) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelConferenceCallUserRole) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelConferenceCallUserRole) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelConferenceCallUserRole) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelConferenceCallUserRole) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelConferenceCallUserRole) UnsetId() {
	o.Id.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelConferenceCallUserRole) GetRole() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelConferenceCallUserRole) GetRoleOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Role) {
		return map[string]interface{}{}, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ModelConferenceCallUserRole) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given map[string]interface{} and assigns it to the Role field.
func (o *ModelConferenceCallUserRole) SetRole(v map[string]interface{}) {
	o.Role = v
}

func (o ModelConferenceCallUserRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelConferenceCallUserRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelConferenceCallUserRole) UnmarshalJSON(data []byte) (err error) {
	varModelConferenceCallUserRole := _ModelConferenceCallUserRole{}

	err = json.Unmarshal(data, &varModelConferenceCallUserRole)

	if err != nil {
		return err
	}

	*o = ModelConferenceCallUserRole(varModelConferenceCallUserRole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelConferenceCallUserRole struct {
	value *ModelConferenceCallUserRole
	isSet bool
}

func (v NullableModelConferenceCallUserRole) Get() *ModelConferenceCallUserRole {
	return v.value
}

func (v *NullableModelConferenceCallUserRole) Set(val *ModelConferenceCallUserRole) {
	v.value = val
	v.isSet = true
}

func (v NullableModelConferenceCallUserRole) IsSet() bool {
	return v.isSet
}

func (v *NullableModelConferenceCallUserRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelConferenceCallUserRole(val *ModelConferenceCallUserRole) *NullableModelConferenceCallUserRole {
	return &NullableModelConferenceCallUserRole{value: val, isSet: true}
}

func (v NullableModelConferenceCallUserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelConferenceCallUserRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


