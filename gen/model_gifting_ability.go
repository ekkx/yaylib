
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftingAbility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftingAbility{}

// GiftingAbility struct for GiftingAbility
type GiftingAbility struct {
	CanReceive NullableBool `json:"can_receive,omitempty"`
	CanSend NullableBool `json:"can_send,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftingAbility GiftingAbility

// NewGiftingAbility instantiates a new GiftingAbility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftingAbility() *GiftingAbility {
	this := GiftingAbility{}
	return &this
}

// NewGiftingAbilityWithDefaults instantiates a new GiftingAbility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftingAbilityWithDefaults() *GiftingAbility {
	this := GiftingAbility{}
	return &this
}

// GetCanReceive returns the CanReceive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftingAbility) GetCanReceive() bool {
	if o == nil || IsNil(o.CanReceive.Get()) {
		var ret bool
		return ret
	}
	return *o.CanReceive.Get()
}

// GetCanReceiveOk returns a tuple with the CanReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftingAbility) GetCanReceiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanReceive.Get(), o.CanReceive.IsSet()
}

// HasCanReceive returns a boolean if a field has been set.
func (o *GiftingAbility) HasCanReceive() bool {
	if o != nil && o.CanReceive.IsSet() {
		return true
	}

	return false
}

// SetCanReceive gets a reference to the given NullableBool and assigns it to the CanReceive field.
func (o *GiftingAbility) SetCanReceive(v bool) {
	o.CanReceive.Set(&v)
}
// SetCanReceiveNil sets the value for CanReceive to be an explicit nil
func (o *GiftingAbility) SetCanReceiveNil() {
	o.CanReceive.Set(nil)
}

// UnsetCanReceive ensures that no value is present for CanReceive, not even an explicit nil
func (o *GiftingAbility) UnsetCanReceive() {
	o.CanReceive.Unset()
}

// GetCanSend returns the CanSend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftingAbility) GetCanSend() bool {
	if o == nil || IsNil(o.CanSend.Get()) {
		var ret bool
		return ret
	}
	return *o.CanSend.Get()
}

// GetCanSendOk returns a tuple with the CanSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftingAbility) GetCanSendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanSend.Get(), o.CanSend.IsSet()
}

// HasCanSend returns a boolean if a field has been set.
func (o *GiftingAbility) HasCanSend() bool {
	if o != nil && o.CanSend.IsSet() {
		return true
	}

	return false
}

// SetCanSend gets a reference to the given NullableBool and assigns it to the CanSend field.
func (o *GiftingAbility) SetCanSend(v bool) {
	o.CanSend.Set(&v)
}
// SetCanSendNil sets the value for CanSend to be an explicit nil
func (o *GiftingAbility) SetCanSendNil() {
	o.CanSend.Set(nil)
}

// UnsetCanSend ensures that no value is present for CanSend, not even an explicit nil
func (o *GiftingAbility) UnsetCanSend() {
	o.CanSend.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftingAbility) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftingAbility) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *GiftingAbility) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *GiftingAbility) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *GiftingAbility) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *GiftingAbility) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftingAbility) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftingAbility) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *GiftingAbility) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *GiftingAbility) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *GiftingAbility) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *GiftingAbility) UnsetUserId() {
	o.UserId.Unset()
}

func (o GiftingAbility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftingAbility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CanReceive.IsSet() {
		toSerialize["can_receive"] = o.CanReceive.Get()
	}
	if o.CanSend.IsSet() {
		toSerialize["can_send"] = o.CanSend.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftingAbility) UnmarshalJSON(data []byte) (err error) {
	varGiftingAbility := _GiftingAbility{}

	err = json.Unmarshal(data, &varGiftingAbility)

	if err != nil {
		return err
	}

	*o = GiftingAbility(varGiftingAbility)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "can_receive")
		delete(additionalProperties, "can_send")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftingAbility struct {
	value *GiftingAbility
	isSet bool
}

func (v NullableGiftingAbility) Get() *GiftingAbility {
	return v.value
}

func (v *NullableGiftingAbility) Set(val *GiftingAbility) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftingAbility) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftingAbility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftingAbility(val *GiftingAbility) *NullableGiftingAbility {
	return &NullableGiftingAbility{value: val, isSet: true}
}

func (v NullableGiftingAbility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftingAbility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


