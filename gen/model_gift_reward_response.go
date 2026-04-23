
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftRewardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftRewardResponse{}

// GiftRewardResponse struct for GiftRewardResponse
type GiftRewardResponse struct {
	GiftRewards []GiftRewards `json:"gift_rewards,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftRewardResponse GiftRewardResponse

// NewGiftRewardResponse instantiates a new GiftRewardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftRewardResponse() *GiftRewardResponse {
	this := GiftRewardResponse{}
	return &this
}

// NewGiftRewardResponseWithDefaults instantiates a new GiftRewardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftRewardResponseWithDefaults() *GiftRewardResponse {
	this := GiftRewardResponse{}
	return &this
}

// GetGiftRewards returns the GiftRewards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewardResponse) GetGiftRewards() []GiftRewards {
	if o == nil {
		var ret []GiftRewards
		return ret
	}
	return o.GiftRewards
}

// GetGiftRewardsOk returns a tuple with the GiftRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewardResponse) GetGiftRewardsOk() ([]GiftRewards, bool) {
	if o == nil || IsNil(o.GiftRewards) {
		return nil, false
	}
	return o.GiftRewards, true
}

// HasGiftRewards returns a boolean if a field has been set.
func (o *GiftRewardResponse) HasGiftRewards() bool {
	if o != nil && !IsNil(o.GiftRewards) {
		return true
	}

	return false
}

// SetGiftRewards gets a reference to the given []GiftRewards and assigns it to the GiftRewards field.
func (o *GiftRewardResponse) SetGiftRewards(v []GiftRewards) {
	o.GiftRewards = v
}

func (o GiftRewardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftRewardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GiftRewards != nil {
		toSerialize["gift_rewards"] = o.GiftRewards
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftRewardResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftRewardResponse := _GiftRewardResponse{}

	err = json.Unmarshal(data, &varGiftRewardResponse)

	if err != nil {
		return err
	}

	*o = GiftRewardResponse(varGiftRewardResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift_rewards")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftRewardResponse struct {
	value *GiftRewardResponse
	isSet bool
}

func (v NullableGiftRewardResponse) Get() *GiftRewardResponse {
	return v.value
}

func (v *NullableGiftRewardResponse) Set(val *GiftRewardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftRewardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftRewardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftRewardResponse(val *GiftRewardResponse) *NullableGiftRewardResponse {
	return &NullableGiftRewardResponse{value: val, isSet: true}
}

func (v NullableGiftRewardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftRewardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


