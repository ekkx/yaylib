
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Multiplier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Multiplier{}

// Multiplier struct for Multiplier
type Multiplier struct {
	Activity NullableString `json:"activity,omitempty"`
	Base NullableString `json:"base,omitempty"`
	Date NullableString `json:"date,omitempty"`
	Mission NullableString `json:"mission,omitempty"`
	Multiplier NullableString `json:"multiplier,omitempty"`
	Vip NullableString `json:"vip,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Multiplier Multiplier

// NewMultiplier instantiates a new Multiplier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiplier() *Multiplier {
	this := Multiplier{}
	return &this
}

// NewMultiplierWithDefaults instantiates a new Multiplier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiplierWithDefaults() *Multiplier {
	this := Multiplier{}
	return &this
}

// GetActivity returns the Activity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Multiplier) GetActivity() string {
	if o == nil || IsNil(o.Activity.Get()) {
		var ret string
		return ret
	}
	return *o.Activity.Get()
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Multiplier) GetActivityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Activity.Get(), o.Activity.IsSet()
}

// HasActivity returns a boolean if a field has been set.
func (o *Multiplier) HasActivity() bool {
	if o != nil && o.Activity.IsSet() {
		return true
	}

	return false
}

// SetActivity gets a reference to the given NullableString and assigns it to the Activity field.
func (o *Multiplier) SetActivity(v string) {
	o.Activity.Set(&v)
}
// SetActivityNil sets the value for Activity to be an explicit nil
func (o *Multiplier) SetActivityNil() {
	o.Activity.Set(nil)
}

// UnsetActivity ensures that no value is present for Activity, not even an explicit nil
func (o *Multiplier) UnsetActivity() {
	o.Activity.Unset()
}

// GetBase returns the Base field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Multiplier) GetBase() string {
	if o == nil || IsNil(o.Base.Get()) {
		var ret string
		return ret
	}
	return *o.Base.Get()
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Multiplier) GetBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Base.Get(), o.Base.IsSet()
}

// HasBase returns a boolean if a field has been set.
func (o *Multiplier) HasBase() bool {
	if o != nil && o.Base.IsSet() {
		return true
	}

	return false
}

// SetBase gets a reference to the given NullableString and assigns it to the Base field.
func (o *Multiplier) SetBase(v string) {
	o.Base.Set(&v)
}
// SetBaseNil sets the value for Base to be an explicit nil
func (o *Multiplier) SetBaseNil() {
	o.Base.Set(nil)
}

// UnsetBase ensures that no value is present for Base, not even an explicit nil
func (o *Multiplier) UnsetBase() {
	o.Base.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Multiplier) GetDate() string {
	if o == nil || IsNil(o.Date.Get()) {
		var ret string
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Multiplier) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *Multiplier) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableString and assigns it to the Date field.
func (o *Multiplier) SetDate(v string) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *Multiplier) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *Multiplier) UnsetDate() {
	o.Date.Unset()
}

// GetMission returns the Mission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Multiplier) GetMission() string {
	if o == nil || IsNil(o.Mission.Get()) {
		var ret string
		return ret
	}
	return *o.Mission.Get()
}

// GetMissionOk returns a tuple with the Mission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Multiplier) GetMissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mission.Get(), o.Mission.IsSet()
}

// HasMission returns a boolean if a field has been set.
func (o *Multiplier) HasMission() bool {
	if o != nil && o.Mission.IsSet() {
		return true
	}

	return false
}

// SetMission gets a reference to the given NullableString and assigns it to the Mission field.
func (o *Multiplier) SetMission(v string) {
	o.Mission.Set(&v)
}
// SetMissionNil sets the value for Mission to be an explicit nil
func (o *Multiplier) SetMissionNil() {
	o.Mission.Set(nil)
}

// UnsetMission ensures that no value is present for Mission, not even an explicit nil
func (o *Multiplier) UnsetMission() {
	o.Mission.Unset()
}

// GetMultiplier returns the Multiplier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Multiplier) GetMultiplier() string {
	if o == nil || IsNil(o.Multiplier.Get()) {
		var ret string
		return ret
	}
	return *o.Multiplier.Get()
}

// GetMultiplierOk returns a tuple with the Multiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Multiplier) GetMultiplierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Multiplier.Get(), o.Multiplier.IsSet()
}

// HasMultiplier returns a boolean if a field has been set.
func (o *Multiplier) HasMultiplier() bool {
	if o != nil && o.Multiplier.IsSet() {
		return true
	}

	return false
}

// SetMultiplier gets a reference to the given NullableString and assigns it to the Multiplier field.
func (o *Multiplier) SetMultiplier(v string) {
	o.Multiplier.Set(&v)
}
// SetMultiplierNil sets the value for Multiplier to be an explicit nil
func (o *Multiplier) SetMultiplierNil() {
	o.Multiplier.Set(nil)
}

// UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil
func (o *Multiplier) UnsetMultiplier() {
	o.Multiplier.Unset()
}

// GetVip returns the Vip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Multiplier) GetVip() string {
	if o == nil || IsNil(o.Vip.Get()) {
		var ret string
		return ret
	}
	return *o.Vip.Get()
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Multiplier) GetVipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vip.Get(), o.Vip.IsSet()
}

// HasVip returns a boolean if a field has been set.
func (o *Multiplier) HasVip() bool {
	if o != nil && o.Vip.IsSet() {
		return true
	}

	return false
}

// SetVip gets a reference to the given NullableString and assigns it to the Vip field.
func (o *Multiplier) SetVip(v string) {
	o.Vip.Set(&v)
}
// SetVipNil sets the value for Vip to be an explicit nil
func (o *Multiplier) SetVipNil() {
	o.Vip.Set(nil)
}

// UnsetVip ensures that no value is present for Vip, not even an explicit nil
func (o *Multiplier) UnsetVip() {
	o.Vip.Unset()
}

func (o Multiplier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Multiplier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Activity.IsSet() {
		toSerialize["activity"] = o.Activity.Get()
	}
	if o.Base.IsSet() {
		toSerialize["base"] = o.Base.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Mission.IsSet() {
		toSerialize["mission"] = o.Mission.Get()
	}
	if o.Multiplier.IsSet() {
		toSerialize["multiplier"] = o.Multiplier.Get()
	}
	if o.Vip.IsSet() {
		toSerialize["vip"] = o.Vip.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Multiplier) UnmarshalJSON(data []byte) (err error) {
	varMultiplier := _Multiplier{}

	err = json.Unmarshal(data, &varMultiplier)

	if err != nil {
		return err
	}

	*o = Multiplier(varMultiplier)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activity")
		delete(additionalProperties, "base")
		delete(additionalProperties, "date")
		delete(additionalProperties, "mission")
		delete(additionalProperties, "multiplier")
		delete(additionalProperties, "vip")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMultiplier struct {
	value *Multiplier
	isSet bool
}

func (v NullableMultiplier) Get() *Multiplier {
	return v.value
}

func (v *NullableMultiplier) Set(val *Multiplier) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiplier) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiplier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiplier(val *Multiplier) *NullableMultiplier {
	return &NullableMultiplier{value: val, isSet: true}
}

func (v NullableMultiplier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiplier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


