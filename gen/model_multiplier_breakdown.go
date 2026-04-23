
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MultiplierBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiplierBreakdown{}

// MultiplierBreakdown struct for MultiplierBreakdown
type MultiplierBreakdown struct {
	Activity NullableFloat64 `json:"activity,omitempty"`
	SpecialMission NullableFloat64 `json:"special_mission,omitempty"`
	VipMultiplier NullableFloat64 `json:"vip_multiplier,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultiplierBreakdown MultiplierBreakdown

// NewMultiplierBreakdown instantiates a new MultiplierBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiplierBreakdown() *MultiplierBreakdown {
	this := MultiplierBreakdown{}
	return &this
}

// NewMultiplierBreakdownWithDefaults instantiates a new MultiplierBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiplierBreakdownWithDefaults() *MultiplierBreakdown {
	this := MultiplierBreakdown{}
	return &this
}

// GetActivity returns the Activity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultiplierBreakdown) GetActivity() float64 {
	if o == nil || IsNil(o.Activity.Get()) {
		var ret float64
		return ret
	}
	return *o.Activity.Get()
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiplierBreakdown) GetActivityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Activity.Get(), o.Activity.IsSet()
}

// HasActivity returns a boolean if a field has been set.
func (o *MultiplierBreakdown) HasActivity() bool {
	if o != nil && o.Activity.IsSet() {
		return true
	}

	return false
}

// SetActivity gets a reference to the given NullableFloat64 and assigns it to the Activity field.
func (o *MultiplierBreakdown) SetActivity(v float64) {
	o.Activity.Set(&v)
}
// SetActivityNil sets the value for Activity to be an explicit nil
func (o *MultiplierBreakdown) SetActivityNil() {
	o.Activity.Set(nil)
}

// UnsetActivity ensures that no value is present for Activity, not even an explicit nil
func (o *MultiplierBreakdown) UnsetActivity() {
	o.Activity.Unset()
}

// GetSpecialMission returns the SpecialMission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultiplierBreakdown) GetSpecialMission() float64 {
	if o == nil || IsNil(o.SpecialMission.Get()) {
		var ret float64
		return ret
	}
	return *o.SpecialMission.Get()
}

// GetSpecialMissionOk returns a tuple with the SpecialMission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiplierBreakdown) GetSpecialMissionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecialMission.Get(), o.SpecialMission.IsSet()
}

// HasSpecialMission returns a boolean if a field has been set.
func (o *MultiplierBreakdown) HasSpecialMission() bool {
	if o != nil && o.SpecialMission.IsSet() {
		return true
	}

	return false
}

// SetSpecialMission gets a reference to the given NullableFloat64 and assigns it to the SpecialMission field.
func (o *MultiplierBreakdown) SetSpecialMission(v float64) {
	o.SpecialMission.Set(&v)
}
// SetSpecialMissionNil sets the value for SpecialMission to be an explicit nil
func (o *MultiplierBreakdown) SetSpecialMissionNil() {
	o.SpecialMission.Set(nil)
}

// UnsetSpecialMission ensures that no value is present for SpecialMission, not even an explicit nil
func (o *MultiplierBreakdown) UnsetSpecialMission() {
	o.SpecialMission.Unset()
}

// GetVipMultiplier returns the VipMultiplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MultiplierBreakdown) GetVipMultiplier() float64 {
	if o == nil || IsNil(o.VipMultiplier.Get()) {
		var ret float64
		return ret
	}
	return *o.VipMultiplier.Get()
}

// GetVipMultiplierOk returns a tuple with the VipMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MultiplierBreakdown) GetVipMultiplierOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipMultiplier.Get(), o.VipMultiplier.IsSet()
}

// HasVipMultiplier returns a boolean if a field has been set.
func (o *MultiplierBreakdown) HasVipMultiplier() bool {
	if o != nil && o.VipMultiplier.IsSet() {
		return true
	}

	return false
}

// SetVipMultiplier gets a reference to the given NullableFloat64 and assigns it to the VipMultiplier field.
func (o *MultiplierBreakdown) SetVipMultiplier(v float64) {
	o.VipMultiplier.Set(&v)
}
// SetVipMultiplierNil sets the value for VipMultiplier to be an explicit nil
func (o *MultiplierBreakdown) SetVipMultiplierNil() {
	o.VipMultiplier.Set(nil)
}

// UnsetVipMultiplier ensures that no value is present for VipMultiplier, not even an explicit nil
func (o *MultiplierBreakdown) UnsetVipMultiplier() {
	o.VipMultiplier.Unset()
}

func (o MultiplierBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiplierBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Activity.IsSet() {
		toSerialize["activity"] = o.Activity.Get()
	}
	if o.SpecialMission.IsSet() {
		toSerialize["special_mission"] = o.SpecialMission.Get()
	}
	if o.VipMultiplier.IsSet() {
		toSerialize["vip_multiplier"] = o.VipMultiplier.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MultiplierBreakdown) UnmarshalJSON(data []byte) (err error) {
	varMultiplierBreakdown := _MultiplierBreakdown{}

	err = json.Unmarshal(data, &varMultiplierBreakdown)

	if err != nil {
		return err
	}

	*o = MultiplierBreakdown(varMultiplierBreakdown)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activity")
		delete(additionalProperties, "special_mission")
		delete(additionalProperties, "vip_multiplier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMultiplierBreakdown struct {
	value *MultiplierBreakdown
	isSet bool
}

func (v NullableMultiplierBreakdown) Get() *MultiplierBreakdown {
	return v.value
}

func (v *NullableMultiplierBreakdown) Set(val *MultiplierBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiplierBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiplierBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiplierBreakdown(val *MultiplierBreakdown) *NullableMultiplierBreakdown {
	return &NullableMultiplierBreakdown{value: val, isSet: true}
}

func (v NullableMultiplierBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiplierBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


