
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Relationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Relationship{}

// Relationship struct for Relationship
type Relationship struct {
	Target NullableTarget `json:"target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Relationship Relationship

// NewRelationship instantiates a new Relationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationship() *Relationship {
	this := Relationship{}
	return &this
}

// NewRelationshipWithDefaults instantiates a new Relationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWithDefaults() *Relationship {
	this := Relationship{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Relationship) GetTarget() Target {
	if o == nil || IsNil(o.Target.Get()) {
		var ret Target
		return ret
	}
	return *o.Target.Get()
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Relationship) GetTargetOk() (*Target, bool) {
	if o == nil {
		return nil, false
	}
	return o.Target.Get(), o.Target.IsSet()
}

// HasTarget returns a boolean if a field has been set.
func (o *Relationship) HasTarget() bool {
	if o != nil && o.Target.IsSet() {
		return true
	}

	return false
}

// SetTarget gets a reference to the given NullableTarget and assigns it to the Target field.
func (o *Relationship) SetTarget(v Target) {
	o.Target.Set(&v)
}
// SetTargetNil sets the value for Target to be an explicit nil
func (o *Relationship) SetTargetNil() {
	o.Target.Set(nil)
}

// UnsetTarget ensures that no value is present for Target, not even an explicit nil
func (o *Relationship) UnsetTarget() {
	o.Target.Unset()
}

func (o Relationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Relationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Target.IsSet() {
		toSerialize["target"] = o.Target.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Relationship) UnmarshalJSON(data []byte) (err error) {
	varRelationship := _Relationship{}

	err = json.Unmarshal(data, &varRelationship)

	if err != nil {
		return err
	}

	*o = Relationship(varRelationship)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRelationship struct {
	value *Relationship
	isSet bool
}

func (v NullableRelationship) Get() *Relationship {
	return v.value
}

func (v *NullableRelationship) Set(val *Relationship) {
	v.value = val
	v.isSet = true
}

func (v NullableRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelationship(val *Relationship) *NullableRelationship {
	return &NullableRelationship{value: val, isSet: true}
}

func (v NullableRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


