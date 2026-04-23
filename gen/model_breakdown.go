
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Breakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Breakdown{}

// Breakdown struct for Breakdown
type Breakdown struct {
	Activity NullableFloat64 `json:"activity,omitempty"`
	OriginalPoints NullableInt32 `json:"original_points,omitempty"`
	SpecialMission NullableFloat64 `json:"special_mission,omitempty"`
	VipMultiplier NullableFloat64 `json:"vip_multiplier,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Breakdown Breakdown

// NewBreakdown instantiates a new Breakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBreakdown() *Breakdown {
	this := Breakdown{}
	return &this
}

// NewBreakdownWithDefaults instantiates a new Breakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBreakdownWithDefaults() *Breakdown {
	this := Breakdown{}
	return &this
}

// GetActivity returns the Activity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetActivity() float64 {
	if o == nil || IsNil(o.Activity.Get()) {
		var ret float64
		return ret
	}
	return *o.Activity.Get()
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetActivityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Activity.Get(), o.Activity.IsSet()
}

// HasActivity returns a boolean if a field has been set.
func (o *Breakdown) HasActivity() bool {
	if o != nil && o.Activity.IsSet() {
		return true
	}

	return false
}

// SetActivity gets a reference to the given NullableFloat64 and assigns it to the Activity field.
func (o *Breakdown) SetActivity(v float64) {
	o.Activity.Set(&v)
}
// SetActivityNil sets the value for Activity to be an explicit nil
func (o *Breakdown) SetActivityNil() {
	o.Activity.Set(nil)
}

// UnsetActivity ensures that no value is present for Activity, not even an explicit nil
func (o *Breakdown) UnsetActivity() {
	o.Activity.Unset()
}

// GetOriginalPoints returns the OriginalPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetOriginalPoints() int32 {
	if o == nil || IsNil(o.OriginalPoints.Get()) {
		var ret int32
		return ret
	}
	return *o.OriginalPoints.Get()
}

// GetOriginalPointsOk returns a tuple with the OriginalPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetOriginalPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalPoints.Get(), o.OriginalPoints.IsSet()
}

// HasOriginalPoints returns a boolean if a field has been set.
func (o *Breakdown) HasOriginalPoints() bool {
	if o != nil && o.OriginalPoints.IsSet() {
		return true
	}

	return false
}

// SetOriginalPoints gets a reference to the given NullableInt32 and assigns it to the OriginalPoints field.
func (o *Breakdown) SetOriginalPoints(v int32) {
	o.OriginalPoints.Set(&v)
}
// SetOriginalPointsNil sets the value for OriginalPoints to be an explicit nil
func (o *Breakdown) SetOriginalPointsNil() {
	o.OriginalPoints.Set(nil)
}

// UnsetOriginalPoints ensures that no value is present for OriginalPoints, not even an explicit nil
func (o *Breakdown) UnsetOriginalPoints() {
	o.OriginalPoints.Unset()
}

// GetSpecialMission returns the SpecialMission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetSpecialMission() float64 {
	if o == nil || IsNil(o.SpecialMission.Get()) {
		var ret float64
		return ret
	}
	return *o.SpecialMission.Get()
}

// GetSpecialMissionOk returns a tuple with the SpecialMission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetSpecialMissionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecialMission.Get(), o.SpecialMission.IsSet()
}

// HasSpecialMission returns a boolean if a field has been set.
func (o *Breakdown) HasSpecialMission() bool {
	if o != nil && o.SpecialMission.IsSet() {
		return true
	}

	return false
}

// SetSpecialMission gets a reference to the given NullableFloat64 and assigns it to the SpecialMission field.
func (o *Breakdown) SetSpecialMission(v float64) {
	o.SpecialMission.Set(&v)
}
// SetSpecialMissionNil sets the value for SpecialMission to be an explicit nil
func (o *Breakdown) SetSpecialMissionNil() {
	o.SpecialMission.Set(nil)
}

// UnsetSpecialMission ensures that no value is present for SpecialMission, not even an explicit nil
func (o *Breakdown) UnsetSpecialMission() {
	o.SpecialMission.Unset()
}

// GetVipMultiplier returns the VipMultiplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Breakdown) GetVipMultiplier() float64 {
	if o == nil || IsNil(o.VipMultiplier.Get()) {
		var ret float64
		return ret
	}
	return *o.VipMultiplier.Get()
}

// GetVipMultiplierOk returns a tuple with the VipMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Breakdown) GetVipMultiplierOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipMultiplier.Get(), o.VipMultiplier.IsSet()
}

// HasVipMultiplier returns a boolean if a field has been set.
func (o *Breakdown) HasVipMultiplier() bool {
	if o != nil && o.VipMultiplier.IsSet() {
		return true
	}

	return false
}

// SetVipMultiplier gets a reference to the given NullableFloat64 and assigns it to the VipMultiplier field.
func (o *Breakdown) SetVipMultiplier(v float64) {
	o.VipMultiplier.Set(&v)
}
// SetVipMultiplierNil sets the value for VipMultiplier to be an explicit nil
func (o *Breakdown) SetVipMultiplierNil() {
	o.VipMultiplier.Set(nil)
}

// UnsetVipMultiplier ensures that no value is present for VipMultiplier, not even an explicit nil
func (o *Breakdown) UnsetVipMultiplier() {
	o.VipMultiplier.Unset()
}

func (o Breakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Breakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Activity.IsSet() {
		toSerialize["activity"] = o.Activity.Get()
	}
	if o.OriginalPoints.IsSet() {
		toSerialize["original_points"] = o.OriginalPoints.Get()
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

func (o *Breakdown) UnmarshalJSON(data []byte) (err error) {
	varBreakdown := _Breakdown{}

	err = json.Unmarshal(data, &varBreakdown)

	if err != nil {
		return err
	}

	*o = Breakdown(varBreakdown)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activity")
		delete(additionalProperties, "original_points")
		delete(additionalProperties, "special_mission")
		delete(additionalProperties, "vip_multiplier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBreakdown struct {
	value *Breakdown
	isSet bool
}

func (v NullableBreakdown) Get() *Breakdown {
	return v.value
}

func (v *NullableBreakdown) Set(val *Breakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBreakdown(val *Breakdown) *NullableBreakdown {
	return &NullableBreakdown{value: val, isSet: true}
}

func (v NullableBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


