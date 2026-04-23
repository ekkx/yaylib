
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RaceReward type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RaceReward{}

// RaceReward struct for RaceReward
type RaceReward struct {
	FirstPrize NullableReward `json:"first_prize,omitempty"`
	SecondPrize NullableReward `json:"second_prize,omitempty"`
	ThirdPrize NullableReward `json:"third_prize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RaceReward RaceReward

// NewRaceReward instantiates a new RaceReward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRaceReward() *RaceReward {
	this := RaceReward{}
	return &this
}

// NewRaceRewardWithDefaults instantiates a new RaceReward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRaceRewardWithDefaults() *RaceReward {
	this := RaceReward{}
	return &this
}

// GetFirstPrize returns the FirstPrize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceReward) GetFirstPrize() Reward {
	if o == nil || IsNil(o.FirstPrize.Get()) {
		var ret Reward
		return ret
	}
	return *o.FirstPrize.Get()
}

// GetFirstPrizeOk returns a tuple with the FirstPrize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceReward) GetFirstPrizeOk() (*Reward, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstPrize.Get(), o.FirstPrize.IsSet()
}

// HasFirstPrize returns a boolean if a field has been set.
func (o *RaceReward) HasFirstPrize() bool {
	if o != nil && o.FirstPrize.IsSet() {
		return true
	}

	return false
}

// SetFirstPrize gets a reference to the given NullableReward and assigns it to the FirstPrize field.
func (o *RaceReward) SetFirstPrize(v Reward) {
	o.FirstPrize.Set(&v)
}
// SetFirstPrizeNil sets the value for FirstPrize to be an explicit nil
func (o *RaceReward) SetFirstPrizeNil() {
	o.FirstPrize.Set(nil)
}

// UnsetFirstPrize ensures that no value is present for FirstPrize, not even an explicit nil
func (o *RaceReward) UnsetFirstPrize() {
	o.FirstPrize.Unset()
}

// GetSecondPrize returns the SecondPrize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceReward) GetSecondPrize() Reward {
	if o == nil || IsNil(o.SecondPrize.Get()) {
		var ret Reward
		return ret
	}
	return *o.SecondPrize.Get()
}

// GetSecondPrizeOk returns a tuple with the SecondPrize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceReward) GetSecondPrizeOk() (*Reward, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondPrize.Get(), o.SecondPrize.IsSet()
}

// HasSecondPrize returns a boolean if a field has been set.
func (o *RaceReward) HasSecondPrize() bool {
	if o != nil && o.SecondPrize.IsSet() {
		return true
	}

	return false
}

// SetSecondPrize gets a reference to the given NullableReward and assigns it to the SecondPrize field.
func (o *RaceReward) SetSecondPrize(v Reward) {
	o.SecondPrize.Set(&v)
}
// SetSecondPrizeNil sets the value for SecondPrize to be an explicit nil
func (o *RaceReward) SetSecondPrizeNil() {
	o.SecondPrize.Set(nil)
}

// UnsetSecondPrize ensures that no value is present for SecondPrize, not even an explicit nil
func (o *RaceReward) UnsetSecondPrize() {
	o.SecondPrize.Unset()
}

// GetThirdPrize returns the ThirdPrize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RaceReward) GetThirdPrize() Reward {
	if o == nil || IsNil(o.ThirdPrize.Get()) {
		var ret Reward
		return ret
	}
	return *o.ThirdPrize.Get()
}

// GetThirdPrizeOk returns a tuple with the ThirdPrize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RaceReward) GetThirdPrizeOk() (*Reward, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPrize.Get(), o.ThirdPrize.IsSet()
}

// HasThirdPrize returns a boolean if a field has been set.
func (o *RaceReward) HasThirdPrize() bool {
	if o != nil && o.ThirdPrize.IsSet() {
		return true
	}

	return false
}

// SetThirdPrize gets a reference to the given NullableReward and assigns it to the ThirdPrize field.
func (o *RaceReward) SetThirdPrize(v Reward) {
	o.ThirdPrize.Set(&v)
}
// SetThirdPrizeNil sets the value for ThirdPrize to be an explicit nil
func (o *RaceReward) SetThirdPrizeNil() {
	o.ThirdPrize.Set(nil)
}

// UnsetThirdPrize ensures that no value is present for ThirdPrize, not even an explicit nil
func (o *RaceReward) UnsetThirdPrize() {
	o.ThirdPrize.Unset()
}

func (o RaceReward) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RaceReward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstPrize.IsSet() {
		toSerialize["first_prize"] = o.FirstPrize.Get()
	}
	if o.SecondPrize.IsSet() {
		toSerialize["second_prize"] = o.SecondPrize.Get()
	}
	if o.ThirdPrize.IsSet() {
		toSerialize["third_prize"] = o.ThirdPrize.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RaceReward) UnmarshalJSON(data []byte) (err error) {
	varRaceReward := _RaceReward{}

	err = json.Unmarshal(data, &varRaceReward)

	if err != nil {
		return err
	}

	*o = RaceReward(varRaceReward)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "first_prize")
		delete(additionalProperties, "second_prize")
		delete(additionalProperties, "third_prize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRaceReward struct {
	value *RaceReward
	isSet bool
}

func (v NullableRaceReward) Get() *RaceReward {
	return v.value
}

func (v *NullableRaceReward) Set(val *RaceReward) {
	v.value = val
	v.isSet = true
}

func (v NullableRaceReward) IsSet() bool {
	return v.isSet
}

func (v *NullableRaceReward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRaceReward(val *RaceReward) *NullableRaceReward {
	return &NullableRaceReward{value: val, isSet: true}
}

func (v NullableRaceReward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRaceReward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


