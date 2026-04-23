
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Reward type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reward{}

// Reward struct for Reward
type Reward struct {
	RewardedEmpl NullableInt32 `json:"rewarded_empl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Reward Reward

// NewReward instantiates a new Reward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReward() *Reward {
	this := Reward{}
	return &this
}

// NewRewardWithDefaults instantiates a new Reward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardWithDefaults() *Reward {
	this := Reward{}
	return &this
}

// GetRewardedEmpl returns the RewardedEmpl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Reward) GetRewardedEmpl() int32 {
	if o == nil || IsNil(o.RewardedEmpl.Get()) {
		var ret int32
		return ret
	}
	return *o.RewardedEmpl.Get()
}

// GetRewardedEmplOk returns a tuple with the RewardedEmpl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Reward) GetRewardedEmplOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardedEmpl.Get(), o.RewardedEmpl.IsSet()
}

// HasRewardedEmpl returns a boolean if a field has been set.
func (o *Reward) HasRewardedEmpl() bool {
	if o != nil && o.RewardedEmpl.IsSet() {
		return true
	}

	return false
}

// SetRewardedEmpl gets a reference to the given NullableInt32 and assigns it to the RewardedEmpl field.
func (o *Reward) SetRewardedEmpl(v int32) {
	o.RewardedEmpl.Set(&v)
}
// SetRewardedEmplNil sets the value for RewardedEmpl to be an explicit nil
func (o *Reward) SetRewardedEmplNil() {
	o.RewardedEmpl.Set(nil)
}

// UnsetRewardedEmpl ensures that no value is present for RewardedEmpl, not even an explicit nil
func (o *Reward) UnsetRewardedEmpl() {
	o.RewardedEmpl.Unset()
}

func (o Reward) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RewardedEmpl.IsSet() {
		toSerialize["rewarded_empl"] = o.RewardedEmpl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Reward) UnmarshalJSON(data []byte) (err error) {
	varReward := _Reward{}

	err = json.Unmarshal(data, &varReward)

	if err != nil {
		return err
	}

	*o = Reward(varReward)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rewarded_empl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReward struct {
	value *Reward
	isSet bool
}

func (v NullableReward) Get() *Reward {
	return v.value
}

func (v *NullableReward) Set(val *Reward) {
	v.value = val
	v.isSet = true
}

func (v NullableReward) IsSet() bool {
	return v.isSet
}

func (v *NullableReward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReward(val *Reward) *NullableReward {
	return &NullableReward{value: val, isSet: true}
}

func (v NullableReward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


