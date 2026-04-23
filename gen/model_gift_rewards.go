
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftRewards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftRewards{}

// GiftRewards struct for GiftRewards
type GiftRewards struct {
	Id NullableInt32 `json:"id,omitempty"`
	Requirements []Requirement `json:"requirements,omitempty"`
	RewardedAmount NullableInt64 `json:"rewarded_amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftRewards GiftRewards

// NewGiftRewards instantiates a new GiftRewards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftRewards() *GiftRewards {
	this := GiftRewards{}
	return &this
}

// NewGiftRewardsWithDefaults instantiates a new GiftRewards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftRewardsWithDefaults() *GiftRewards {
	this := GiftRewards{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewards) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewards) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GiftRewards) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *GiftRewards) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GiftRewards) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GiftRewards) UnsetId() {
	o.Id.Unset()
}

// GetRequirements returns the Requirements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewards) GetRequirements() []Requirement {
	if o == nil {
		var ret []Requirement
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewards) GetRequirementsOk() ([]Requirement, bool) {
	if o == nil || IsNil(o.Requirements) {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *GiftRewards) HasRequirements() bool {
	if o != nil && !IsNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given []Requirement and assigns it to the Requirements field.
func (o *GiftRewards) SetRequirements(v []Requirement) {
	o.Requirements = v
}

// GetRewardedAmount returns the RewardedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewards) GetRewardedAmount() int64 {
	if o == nil || IsNil(o.RewardedAmount.Get()) {
		var ret int64
		return ret
	}
	return *o.RewardedAmount.Get()
}

// GetRewardedAmountOk returns a tuple with the RewardedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewards) GetRewardedAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardedAmount.Get(), o.RewardedAmount.IsSet()
}

// HasRewardedAmount returns a boolean if a field has been set.
func (o *GiftRewards) HasRewardedAmount() bool {
	if o != nil && o.RewardedAmount.IsSet() {
		return true
	}

	return false
}

// SetRewardedAmount gets a reference to the given NullableInt64 and assigns it to the RewardedAmount field.
func (o *GiftRewards) SetRewardedAmount(v int64) {
	o.RewardedAmount.Set(&v)
}
// SetRewardedAmountNil sets the value for RewardedAmount to be an explicit nil
func (o *GiftRewards) SetRewardedAmountNil() {
	o.RewardedAmount.Set(nil)
}

// UnsetRewardedAmount ensures that no value is present for RewardedAmount, not even an explicit nil
func (o *GiftRewards) UnsetRewardedAmount() {
	o.RewardedAmount.Unset()
}

func (o GiftRewards) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftRewards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	if o.RewardedAmount.IsSet() {
		toSerialize["rewarded_amount"] = o.RewardedAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftRewards) UnmarshalJSON(data []byte) (err error) {
	varGiftRewards := _GiftRewards{}

	err = json.Unmarshal(data, &varGiftRewards)

	if err != nil {
		return err
	}

	*o = GiftRewards(varGiftRewards)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "requirements")
		delete(additionalProperties, "rewarded_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftRewards struct {
	value *GiftRewards
	isSet bool
}

func (v NullableGiftRewards) Get() *GiftRewards {
	return v.value
}

func (v *NullableGiftRewards) Set(val *GiftRewards) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftRewards) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftRewards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftRewards(val *GiftRewards) *NullableGiftRewards {
	return &NullableGiftRewards{value: val, isSet: true}
}

func (v NullableGiftRewards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftRewards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


