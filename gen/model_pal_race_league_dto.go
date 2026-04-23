
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRaceLeagueDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRaceLeagueDTO{}

// PalRaceLeagueDTO struct for PalRaceLeagueDTO
type PalRaceLeagueDTO struct {
	Accessible NullableBool `json:"accessible,omitempty"`
	EntryFee NullableInt32 `json:"entry_fee,omitempty"`
	GameEndAt NullableInt64 `json:"game_end_at,omitempty"`
	Name NullableString `json:"name,omitempty"`
	PriceReward NullableInt32 `json:"price_reward,omitempty"`
	RegistrationCloseAt NullableInt64 `json:"registration_close_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRaceLeagueDTO PalRaceLeagueDTO

// NewPalRaceLeagueDTO instantiates a new PalRaceLeagueDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRaceLeagueDTO() *PalRaceLeagueDTO {
	this := PalRaceLeagueDTO{}
	return &this
}

// NewPalRaceLeagueDTOWithDefaults instantiates a new PalRaceLeagueDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRaceLeagueDTOWithDefaults() *PalRaceLeagueDTO {
	this := PalRaceLeagueDTO{}
	return &this
}

// GetAccessible returns the Accessible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceLeagueDTO) GetAccessible() bool {
	if o == nil || IsNil(o.Accessible.Get()) {
		var ret bool
		return ret
	}
	return *o.Accessible.Get()
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceLeagueDTO) GetAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accessible.Get(), o.Accessible.IsSet()
}

// HasAccessible returns a boolean if a field has been set.
func (o *PalRaceLeagueDTO) HasAccessible() bool {
	if o != nil && o.Accessible.IsSet() {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given NullableBool and assigns it to the Accessible field.
func (o *PalRaceLeagueDTO) SetAccessible(v bool) {
	o.Accessible.Set(&v)
}
// SetAccessibleNil sets the value for Accessible to be an explicit nil
func (o *PalRaceLeagueDTO) SetAccessibleNil() {
	o.Accessible.Set(nil)
}

// UnsetAccessible ensures that no value is present for Accessible, not even an explicit nil
func (o *PalRaceLeagueDTO) UnsetAccessible() {
	o.Accessible.Unset()
}

// GetEntryFee returns the EntryFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceLeagueDTO) GetEntryFee() int32 {
	if o == nil || IsNil(o.EntryFee.Get()) {
		var ret int32
		return ret
	}
	return *o.EntryFee.Get()
}

// GetEntryFeeOk returns a tuple with the EntryFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceLeagueDTO) GetEntryFeeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryFee.Get(), o.EntryFee.IsSet()
}

// HasEntryFee returns a boolean if a field has been set.
func (o *PalRaceLeagueDTO) HasEntryFee() bool {
	if o != nil && o.EntryFee.IsSet() {
		return true
	}

	return false
}

// SetEntryFee gets a reference to the given NullableInt32 and assigns it to the EntryFee field.
func (o *PalRaceLeagueDTO) SetEntryFee(v int32) {
	o.EntryFee.Set(&v)
}
// SetEntryFeeNil sets the value for EntryFee to be an explicit nil
func (o *PalRaceLeagueDTO) SetEntryFeeNil() {
	o.EntryFee.Set(nil)
}

// UnsetEntryFee ensures that no value is present for EntryFee, not even an explicit nil
func (o *PalRaceLeagueDTO) UnsetEntryFee() {
	o.EntryFee.Unset()
}

// GetGameEndAt returns the GameEndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceLeagueDTO) GetGameEndAt() int64 {
	if o == nil || IsNil(o.GameEndAt.Get()) {
		var ret int64
		return ret
	}
	return *o.GameEndAt.Get()
}

// GetGameEndAtOk returns a tuple with the GameEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceLeagueDTO) GetGameEndAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GameEndAt.Get(), o.GameEndAt.IsSet()
}

// HasGameEndAt returns a boolean if a field has been set.
func (o *PalRaceLeagueDTO) HasGameEndAt() bool {
	if o != nil && o.GameEndAt.IsSet() {
		return true
	}

	return false
}

// SetGameEndAt gets a reference to the given NullableInt64 and assigns it to the GameEndAt field.
func (o *PalRaceLeagueDTO) SetGameEndAt(v int64) {
	o.GameEndAt.Set(&v)
}
// SetGameEndAtNil sets the value for GameEndAt to be an explicit nil
func (o *PalRaceLeagueDTO) SetGameEndAtNil() {
	o.GameEndAt.Set(nil)
}

// UnsetGameEndAt ensures that no value is present for GameEndAt, not even an explicit nil
func (o *PalRaceLeagueDTO) UnsetGameEndAt() {
	o.GameEndAt.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceLeagueDTO) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceLeagueDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PalRaceLeagueDTO) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PalRaceLeagueDTO) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PalRaceLeagueDTO) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PalRaceLeagueDTO) UnsetName() {
	o.Name.Unset()
}

// GetPriceReward returns the PriceReward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceLeagueDTO) GetPriceReward() int32 {
	if o == nil || IsNil(o.PriceReward.Get()) {
		var ret int32
		return ret
	}
	return *o.PriceReward.Get()
}

// GetPriceRewardOk returns a tuple with the PriceReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceLeagueDTO) GetPriceRewardOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceReward.Get(), o.PriceReward.IsSet()
}

// HasPriceReward returns a boolean if a field has been set.
func (o *PalRaceLeagueDTO) HasPriceReward() bool {
	if o != nil && o.PriceReward.IsSet() {
		return true
	}

	return false
}

// SetPriceReward gets a reference to the given NullableInt32 and assigns it to the PriceReward field.
func (o *PalRaceLeagueDTO) SetPriceReward(v int32) {
	o.PriceReward.Set(&v)
}
// SetPriceRewardNil sets the value for PriceReward to be an explicit nil
func (o *PalRaceLeagueDTO) SetPriceRewardNil() {
	o.PriceReward.Set(nil)
}

// UnsetPriceReward ensures that no value is present for PriceReward, not even an explicit nil
func (o *PalRaceLeagueDTO) UnsetPriceReward() {
	o.PriceReward.Unset()
}

// GetRegistrationCloseAt returns the RegistrationCloseAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceLeagueDTO) GetRegistrationCloseAt() int64 {
	if o == nil || IsNil(o.RegistrationCloseAt.Get()) {
		var ret int64
		return ret
	}
	return *o.RegistrationCloseAt.Get()
}

// GetRegistrationCloseAtOk returns a tuple with the RegistrationCloseAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceLeagueDTO) GetRegistrationCloseAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistrationCloseAt.Get(), o.RegistrationCloseAt.IsSet()
}

// HasRegistrationCloseAt returns a boolean if a field has been set.
func (o *PalRaceLeagueDTO) HasRegistrationCloseAt() bool {
	if o != nil && o.RegistrationCloseAt.IsSet() {
		return true
	}

	return false
}

// SetRegistrationCloseAt gets a reference to the given NullableInt64 and assigns it to the RegistrationCloseAt field.
func (o *PalRaceLeagueDTO) SetRegistrationCloseAt(v int64) {
	o.RegistrationCloseAt.Set(&v)
}
// SetRegistrationCloseAtNil sets the value for RegistrationCloseAt to be an explicit nil
func (o *PalRaceLeagueDTO) SetRegistrationCloseAtNil() {
	o.RegistrationCloseAt.Set(nil)
}

// UnsetRegistrationCloseAt ensures that no value is present for RegistrationCloseAt, not even an explicit nil
func (o *PalRaceLeagueDTO) UnsetRegistrationCloseAt() {
	o.RegistrationCloseAt.Unset()
}

func (o PalRaceLeagueDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRaceLeagueDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessible.IsSet() {
		toSerialize["accessible"] = o.Accessible.Get()
	}
	if o.EntryFee.IsSet() {
		toSerialize["entry_fee"] = o.EntryFee.Get()
	}
	if o.GameEndAt.IsSet() {
		toSerialize["game_end_at"] = o.GameEndAt.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.PriceReward.IsSet() {
		toSerialize["price_reward"] = o.PriceReward.Get()
	}
	if o.RegistrationCloseAt.IsSet() {
		toSerialize["registration_close_at"] = o.RegistrationCloseAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRaceLeagueDTO) UnmarshalJSON(data []byte) (err error) {
	varPalRaceLeagueDTO := _PalRaceLeagueDTO{}

	err = json.Unmarshal(data, &varPalRaceLeagueDTO)

	if err != nil {
		return err
	}

	*o = PalRaceLeagueDTO(varPalRaceLeagueDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessible")
		delete(additionalProperties, "entry_fee")
		delete(additionalProperties, "game_end_at")
		delete(additionalProperties, "name")
		delete(additionalProperties, "price_reward")
		delete(additionalProperties, "registration_close_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRaceLeagueDTO struct {
	value *PalRaceLeagueDTO
	isSet bool
}

func (v NullablePalRaceLeagueDTO) Get() *PalRaceLeagueDTO {
	return v.value
}

func (v *NullablePalRaceLeagueDTO) Set(val *PalRaceLeagueDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRaceLeagueDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRaceLeagueDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRaceLeagueDTO(val *PalRaceLeagueDTO) *NullablePalRaceLeagueDTO {
	return &NullablePalRaceLeagueDTO{value: val, isSet: true}
}

func (v NullablePalRaceLeagueDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRaceLeagueDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


