
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Cooldown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cooldown{}

// Cooldown struct for Cooldown
type Cooldown struct {
	Active NullableBool `json:"active,omitempty"`
	RetryAfter NullableInt64 `json:"retry_after,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Cooldown Cooldown

// NewCooldown instantiates a new Cooldown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCooldown() *Cooldown {
	this := Cooldown{}
	return &this
}

// NewCooldownWithDefaults instantiates a new Cooldown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCooldownWithDefaults() *Cooldown {
	this := Cooldown{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cooldown) GetActive() bool {
	if o == nil || IsNil(o.Active.Get()) {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cooldown) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *Cooldown) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *Cooldown) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *Cooldown) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *Cooldown) UnsetActive() {
	o.Active.Unset()
}

// GetRetryAfter returns the RetryAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cooldown) GetRetryAfter() int64 {
	if o == nil || IsNil(o.RetryAfter.Get()) {
		var ret int64
		return ret
	}
	return *o.RetryAfter.Get()
}

// GetRetryAfterOk returns a tuple with the RetryAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cooldown) GetRetryAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryAfter.Get(), o.RetryAfter.IsSet()
}

// HasRetryAfter returns a boolean if a field has been set.
func (o *Cooldown) HasRetryAfter() bool {
	if o != nil && o.RetryAfter.IsSet() {
		return true
	}

	return false
}

// SetRetryAfter gets a reference to the given NullableInt64 and assigns it to the RetryAfter field.
func (o *Cooldown) SetRetryAfter(v int64) {
	o.RetryAfter.Set(&v)
}
// SetRetryAfterNil sets the value for RetryAfter to be an explicit nil
func (o *Cooldown) SetRetryAfterNil() {
	o.RetryAfter.Set(nil)
}

// UnsetRetryAfter ensures that no value is present for RetryAfter, not even an explicit nil
func (o *Cooldown) UnsetRetryAfter() {
	o.RetryAfter.Unset()
}

func (o Cooldown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cooldown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.RetryAfter.IsSet() {
		toSerialize["retry_after"] = o.RetryAfter.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cooldown) UnmarshalJSON(data []byte) (err error) {
	varCooldown := _Cooldown{}

	err = json.Unmarshal(data, &varCooldown)

	if err != nil {
		return err
	}

	*o = Cooldown(varCooldown)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		delete(additionalProperties, "retry_after")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCooldown struct {
	value *Cooldown
	isSet bool
}

func (v NullableCooldown) Get() *Cooldown {
	return v.value
}

func (v *NullableCooldown) Set(val *Cooldown) {
	v.value = val
	v.isSet = true
}

func (v NullableCooldown) IsSet() bool {
	return v.isSet
}

func (v *NullableCooldown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCooldown(val *Cooldown) *NullableCooldown {
	return &NullableCooldown{value: val, isSet: true}
}

func (v NullableCooldown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCooldown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


