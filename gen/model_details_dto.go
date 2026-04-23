
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the DetailsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailsDTO{}

// DetailsDTO struct for DetailsDTO
type DetailsDTO struct {
	RewardAmountAdult NullableInt32 `json:"reward_amount_adult,omitempty"`
	RewardAmountMinor NullableInt32 `json:"reward_amount_minor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DetailsDTO DetailsDTO

// NewDetailsDTO instantiates a new DetailsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailsDTO() *DetailsDTO {
	this := DetailsDTO{}
	return &this
}

// NewDetailsDTOWithDefaults instantiates a new DetailsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailsDTOWithDefaults() *DetailsDTO {
	this := DetailsDTO{}
	return &this
}

// GetRewardAmountAdult returns the RewardAmountAdult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailsDTO) GetRewardAmountAdult() int32 {
	if o == nil || IsNil(o.RewardAmountAdult.Get()) {
		var ret int32
		return ret
	}
	return *o.RewardAmountAdult.Get()
}

// GetRewardAmountAdultOk returns a tuple with the RewardAmountAdult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailsDTO) GetRewardAmountAdultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardAmountAdult.Get(), o.RewardAmountAdult.IsSet()
}

// HasRewardAmountAdult returns a boolean if a field has been set.
func (o *DetailsDTO) HasRewardAmountAdult() bool {
	if o != nil && o.RewardAmountAdult.IsSet() {
		return true
	}

	return false
}

// SetRewardAmountAdult gets a reference to the given NullableInt32 and assigns it to the RewardAmountAdult field.
func (o *DetailsDTO) SetRewardAmountAdult(v int32) {
	o.RewardAmountAdult.Set(&v)
}
// SetRewardAmountAdultNil sets the value for RewardAmountAdult to be an explicit nil
func (o *DetailsDTO) SetRewardAmountAdultNil() {
	o.RewardAmountAdult.Set(nil)
}

// UnsetRewardAmountAdult ensures that no value is present for RewardAmountAdult, not even an explicit nil
func (o *DetailsDTO) UnsetRewardAmountAdult() {
	o.RewardAmountAdult.Unset()
}

// GetRewardAmountMinor returns the RewardAmountMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DetailsDTO) GetRewardAmountMinor() int32 {
	if o == nil || IsNil(o.RewardAmountMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.RewardAmountMinor.Get()
}

// GetRewardAmountMinorOk returns a tuple with the RewardAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DetailsDTO) GetRewardAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardAmountMinor.Get(), o.RewardAmountMinor.IsSet()
}

// HasRewardAmountMinor returns a boolean if a field has been set.
func (o *DetailsDTO) HasRewardAmountMinor() bool {
	if o != nil && o.RewardAmountMinor.IsSet() {
		return true
	}

	return false
}

// SetRewardAmountMinor gets a reference to the given NullableInt32 and assigns it to the RewardAmountMinor field.
func (o *DetailsDTO) SetRewardAmountMinor(v int32) {
	o.RewardAmountMinor.Set(&v)
}
// SetRewardAmountMinorNil sets the value for RewardAmountMinor to be an explicit nil
func (o *DetailsDTO) SetRewardAmountMinorNil() {
	o.RewardAmountMinor.Set(nil)
}

// UnsetRewardAmountMinor ensures that no value is present for RewardAmountMinor, not even an explicit nil
func (o *DetailsDTO) UnsetRewardAmountMinor() {
	o.RewardAmountMinor.Unset()
}

func (o DetailsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RewardAmountAdult.IsSet() {
		toSerialize["reward_amount_adult"] = o.RewardAmountAdult.Get()
	}
	if o.RewardAmountMinor.IsSet() {
		toSerialize["reward_amount_minor"] = o.RewardAmountMinor.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DetailsDTO) UnmarshalJSON(data []byte) (err error) {
	varDetailsDTO := _DetailsDTO{}

	err = json.Unmarshal(data, &varDetailsDTO)

	if err != nil {
		return err
	}

	*o = DetailsDTO(varDetailsDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reward_amount_adult")
		delete(additionalProperties, "reward_amount_minor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetailsDTO struct {
	value *DetailsDTO
	isSet bool
}

func (v NullableDetailsDTO) Get() *DetailsDTO {
	return v.value
}

func (v *NullableDetailsDTO) Set(val *DetailsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailsDTO(val *DetailsDTO) *NullableDetailsDTO {
	return &NullableDetailsDTO{value: val, isSet: true}
}

func (v NullableDetailsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


