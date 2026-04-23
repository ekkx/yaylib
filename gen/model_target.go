
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Target type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Target{}

// Target struct for Target
type Target struct {
	FollowedBy NullableBool `json:"followed_by,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Target Target

// NewTarget instantiates a new Target object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTarget() *Target {
	this := Target{}
	return &this
}

// NewTargetWithDefaults instantiates a new Target object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetWithDefaults() *Target {
	this := Target{}
	return &this
}

// GetFollowedBy returns the FollowedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Target) GetFollowedBy() bool {
	if o == nil || IsNil(o.FollowedBy.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowedBy.Get()
}

// GetFollowedByOk returns a tuple with the FollowedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Target) GetFollowedByOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowedBy.Get(), o.FollowedBy.IsSet()
}

// HasFollowedBy returns a boolean if a field has been set.
func (o *Target) HasFollowedBy() bool {
	if o != nil && o.FollowedBy.IsSet() {
		return true
	}

	return false
}

// SetFollowedBy gets a reference to the given NullableBool and assigns it to the FollowedBy field.
func (o *Target) SetFollowedBy(v bool) {
	o.FollowedBy.Set(&v)
}
// SetFollowedByNil sets the value for FollowedBy to be an explicit nil
func (o *Target) SetFollowedByNil() {
	o.FollowedBy.Set(nil)
}

// UnsetFollowedBy ensures that no value is present for FollowedBy, not even an explicit nil
func (o *Target) UnsetFollowedBy() {
	o.FollowedBy.Unset()
}

func (o Target) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Target) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FollowedBy.IsSet() {
		toSerialize["followed_by"] = o.FollowedBy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Target) UnmarshalJSON(data []byte) (err error) {
	varTarget := _Target{}

	err = json.Unmarshal(data, &varTarget)

	if err != nil {
		return err
	}

	*o = Target(varTarget)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "followed_by")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTarget struct {
	value *Target
	isSet bool
}

func (v NullableTarget) Get() *Target {
	return v.value
}

func (v *NullableTarget) Set(val *Target) {
	v.value = val
	v.isSet = true
}

func (v NullableTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTarget(val *Target) *NullableTarget {
	return &NullableTarget{value: val, isSet: true}
}

func (v NullableTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


