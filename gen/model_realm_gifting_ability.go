
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmGiftingAbility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmGiftingAbility{}

// RealmGiftingAbility struct for RealmGiftingAbility
type RealmGiftingAbility struct {
	CanReceive NullableBool `json:"can_receive,omitempty"`
	CanSend NullableBool `json:"can_send,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmGiftingAbility RealmGiftingAbility

// NewRealmGiftingAbility instantiates a new RealmGiftingAbility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmGiftingAbility() *RealmGiftingAbility {
	this := RealmGiftingAbility{}
	return &this
}

// NewRealmGiftingAbilityWithDefaults instantiates a new RealmGiftingAbility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmGiftingAbilityWithDefaults() *RealmGiftingAbility {
	this := RealmGiftingAbility{}
	return &this
}

// GetCanReceive returns the CanReceive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGiftingAbility) GetCanReceive() bool {
	if o == nil || IsNil(o.CanReceive.Get()) {
		var ret bool
		return ret
	}
	return *o.CanReceive.Get()
}

// GetCanReceiveOk returns a tuple with the CanReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGiftingAbility) GetCanReceiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanReceive.Get(), o.CanReceive.IsSet()
}

// HasCanReceive returns a boolean if a field has been set.
func (o *RealmGiftingAbility) HasCanReceive() bool {
	if o != nil && o.CanReceive.IsSet() {
		return true
	}

	return false
}

// SetCanReceive gets a reference to the given NullableBool and assigns it to the CanReceive field.
func (o *RealmGiftingAbility) SetCanReceive(v bool) {
	o.CanReceive.Set(&v)
}
// SetCanReceiveNil sets the value for CanReceive to be an explicit nil
func (o *RealmGiftingAbility) SetCanReceiveNil() {
	o.CanReceive.Set(nil)
}

// UnsetCanReceive ensures that no value is present for CanReceive, not even an explicit nil
func (o *RealmGiftingAbility) UnsetCanReceive() {
	o.CanReceive.Unset()
}

// GetCanSend returns the CanSend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGiftingAbility) GetCanSend() bool {
	if o == nil || IsNil(o.CanSend.Get()) {
		var ret bool
		return ret
	}
	return *o.CanSend.Get()
}

// GetCanSendOk returns a tuple with the CanSend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGiftingAbility) GetCanSendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanSend.Get(), o.CanSend.IsSet()
}

// HasCanSend returns a boolean if a field has been set.
func (o *RealmGiftingAbility) HasCanSend() bool {
	if o != nil && o.CanSend.IsSet() {
		return true
	}

	return false
}

// SetCanSend gets a reference to the given NullableBool and assigns it to the CanSend field.
func (o *RealmGiftingAbility) SetCanSend(v bool) {
	o.CanSend.Set(&v)
}
// SetCanSendNil sets the value for CanSend to be an explicit nil
func (o *RealmGiftingAbility) SetCanSendNil() {
	o.CanSend.Set(nil)
}

// UnsetCanSend ensures that no value is present for CanSend, not even an explicit nil
func (o *RealmGiftingAbility) UnsetCanSend() {
	o.CanSend.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGiftingAbility) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGiftingAbility) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *RealmGiftingAbility) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *RealmGiftingAbility) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *RealmGiftingAbility) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *RealmGiftingAbility) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmGiftingAbility) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmGiftingAbility) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *RealmGiftingAbility) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *RealmGiftingAbility) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *RealmGiftingAbility) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *RealmGiftingAbility) UnsetUserId() {
	o.UserId.Unset()
}

func (o RealmGiftingAbility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmGiftingAbility) ToMap() (map[string]interface{}, error) {
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

func (o *RealmGiftingAbility) UnmarshalJSON(data []byte) (err error) {
	varRealmGiftingAbility := _RealmGiftingAbility{}

	err = json.Unmarshal(data, &varRealmGiftingAbility)

	if err != nil {
		return err
	}

	*o = RealmGiftingAbility(varRealmGiftingAbility)

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

type NullableRealmGiftingAbility struct {
	value *RealmGiftingAbility
	isSet bool
}

func (v NullableRealmGiftingAbility) Get() *RealmGiftingAbility {
	return v.value
}

func (v *NullableRealmGiftingAbility) Set(val *RealmGiftingAbility) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmGiftingAbility) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmGiftingAbility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmGiftingAbility(val *RealmGiftingAbility) *NullableRealmGiftingAbility {
	return &NullableRealmGiftingAbility{value: val, isSet: true}
}

func (v NullableRealmGiftingAbility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmGiftingAbility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


