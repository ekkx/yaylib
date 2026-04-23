
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftReward type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftReward{}

// GiftReward struct for GiftReward
type GiftReward struct {
	EmplReward NullableInt64 `json:"empl_reward,omitempty"`
	Gifts []GiftRewardGift `json:"gifts,omitempty"`
	Id NullableInt32 `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftReward GiftReward

// NewGiftReward instantiates a new GiftReward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftReward() *GiftReward {
	this := GiftReward{}
	return &this
}

// NewGiftRewardWithDefaults instantiates a new GiftReward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftRewardWithDefaults() *GiftReward {
	this := GiftReward{}
	return &this
}

// GetEmplReward returns the EmplReward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReward) GetEmplReward() int64 {
	if o == nil || IsNil(o.EmplReward.Get()) {
		var ret int64
		return ret
	}
	return *o.EmplReward.Get()
}

// GetEmplRewardOk returns a tuple with the EmplReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReward) GetEmplRewardOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplReward.Get(), o.EmplReward.IsSet()
}

// HasEmplReward returns a boolean if a field has been set.
func (o *GiftReward) HasEmplReward() bool {
	if o != nil && o.EmplReward.IsSet() {
		return true
	}

	return false
}

// SetEmplReward gets a reference to the given NullableInt64 and assigns it to the EmplReward field.
func (o *GiftReward) SetEmplReward(v int64) {
	o.EmplReward.Set(&v)
}
// SetEmplRewardNil sets the value for EmplReward to be an explicit nil
func (o *GiftReward) SetEmplRewardNil() {
	o.EmplReward.Set(nil)
}

// UnsetEmplReward ensures that no value is present for EmplReward, not even an explicit nil
func (o *GiftReward) UnsetEmplReward() {
	o.EmplReward.Unset()
}

// GetGifts returns the Gifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReward) GetGifts() []GiftRewardGift {
	if o == nil {
		var ret []GiftRewardGift
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReward) GetGiftsOk() ([]GiftRewardGift, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *GiftReward) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given []GiftRewardGift and assigns it to the Gifts field.
func (o *GiftReward) SetGifts(v []GiftRewardGift) {
	o.Gifts = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReward) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReward) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GiftReward) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *GiftReward) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GiftReward) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GiftReward) UnsetId() {
	o.Id.Unset()
}

func (o GiftReward) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftReward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EmplReward.IsSet() {
		toSerialize["empl_reward"] = o.EmplReward.Get()
	}
	if o.Gifts != nil {
		toSerialize["gifts"] = o.Gifts
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftReward) UnmarshalJSON(data []byte) (err error) {
	varGiftReward := _GiftReward{}

	err = json.Unmarshal(data, &varGiftReward)

	if err != nil {
		return err
	}

	*o = GiftReward(varGiftReward)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "empl_reward")
		delete(additionalProperties, "gifts")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftReward struct {
	value *GiftReward
	isSet bool
}

func (v NullableGiftReward) Get() *GiftReward {
	return v.value
}

func (v *NullableGiftReward) Set(val *GiftReward) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftReward) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftReward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftReward(val *GiftReward) *NullableGiftReward {
	return &NullableGiftReward{value: val, isSet: true}
}

func (v NullableGiftReward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftReward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


