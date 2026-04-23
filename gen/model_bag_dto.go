
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BagDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BagDTO{}

// BagDTO struct for BagDTO
type BagDTO struct {
	EggsCount NullableInt32 `json:"eggs_count,omitempty"`
	Empl NullableEmplDTO `json:"empl,omitempty"`
	ExpiringEmpl NullableString `json:"expiring_empl,omitempty"`
	PalsCount NullableInt32 `json:"pals_count,omitempty"`
	RacePalAvailable NullableInt32 `json:"race_pal_available,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BagDTO BagDTO

// NewBagDTO instantiates a new BagDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBagDTO() *BagDTO {
	this := BagDTO{}
	return &this
}

// NewBagDTOWithDefaults instantiates a new BagDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBagDTOWithDefaults() *BagDTO {
	this := BagDTO{}
	return &this
}

// GetEggsCount returns the EggsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BagDTO) GetEggsCount() int32 {
	if o == nil || IsNil(o.EggsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.EggsCount.Get()
}

// GetEggsCountOk returns a tuple with the EggsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BagDTO) GetEggsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EggsCount.Get(), o.EggsCount.IsSet()
}

// HasEggsCount returns a boolean if a field has been set.
func (o *BagDTO) HasEggsCount() bool {
	if o != nil && o.EggsCount.IsSet() {
		return true
	}

	return false
}

// SetEggsCount gets a reference to the given NullableInt32 and assigns it to the EggsCount field.
func (o *BagDTO) SetEggsCount(v int32) {
	o.EggsCount.Set(&v)
}
// SetEggsCountNil sets the value for EggsCount to be an explicit nil
func (o *BagDTO) SetEggsCountNil() {
	o.EggsCount.Set(nil)
}

// UnsetEggsCount ensures that no value is present for EggsCount, not even an explicit nil
func (o *BagDTO) UnsetEggsCount() {
	o.EggsCount.Unset()
}

// GetEmpl returns the Empl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BagDTO) GetEmpl() EmplDTO {
	if o == nil || IsNil(o.Empl.Get()) {
		var ret EmplDTO
		return ret
	}
	return *o.Empl.Get()
}

// GetEmplOk returns a tuple with the Empl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BagDTO) GetEmplOk() (*EmplDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Empl.Get(), o.Empl.IsSet()
}

// HasEmpl returns a boolean if a field has been set.
func (o *BagDTO) HasEmpl() bool {
	if o != nil && o.Empl.IsSet() {
		return true
	}

	return false
}

// SetEmpl gets a reference to the given NullableEmplDTO and assigns it to the Empl field.
func (o *BagDTO) SetEmpl(v EmplDTO) {
	o.Empl.Set(&v)
}
// SetEmplNil sets the value for Empl to be an explicit nil
func (o *BagDTO) SetEmplNil() {
	o.Empl.Set(nil)
}

// UnsetEmpl ensures that no value is present for Empl, not even an explicit nil
func (o *BagDTO) UnsetEmpl() {
	o.Empl.Unset()
}

// GetExpiringEmpl returns the ExpiringEmpl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BagDTO) GetExpiringEmpl() string {
	if o == nil || IsNil(o.ExpiringEmpl.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiringEmpl.Get()
}

// GetExpiringEmplOk returns a tuple with the ExpiringEmpl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BagDTO) GetExpiringEmplOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiringEmpl.Get(), o.ExpiringEmpl.IsSet()
}

// HasExpiringEmpl returns a boolean if a field has been set.
func (o *BagDTO) HasExpiringEmpl() bool {
	if o != nil && o.ExpiringEmpl.IsSet() {
		return true
	}

	return false
}

// SetExpiringEmpl gets a reference to the given NullableString and assigns it to the ExpiringEmpl field.
func (o *BagDTO) SetExpiringEmpl(v string) {
	o.ExpiringEmpl.Set(&v)
}
// SetExpiringEmplNil sets the value for ExpiringEmpl to be an explicit nil
func (o *BagDTO) SetExpiringEmplNil() {
	o.ExpiringEmpl.Set(nil)
}

// UnsetExpiringEmpl ensures that no value is present for ExpiringEmpl, not even an explicit nil
func (o *BagDTO) UnsetExpiringEmpl() {
	o.ExpiringEmpl.Unset()
}

// GetPalsCount returns the PalsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BagDTO) GetPalsCount() int32 {
	if o == nil || IsNil(o.PalsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PalsCount.Get()
}

// GetPalsCountOk returns a tuple with the PalsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BagDTO) GetPalsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalsCount.Get(), o.PalsCount.IsSet()
}

// HasPalsCount returns a boolean if a field has been set.
func (o *BagDTO) HasPalsCount() bool {
	if o != nil && o.PalsCount.IsSet() {
		return true
	}

	return false
}

// SetPalsCount gets a reference to the given NullableInt32 and assigns it to the PalsCount field.
func (o *BagDTO) SetPalsCount(v int32) {
	o.PalsCount.Set(&v)
}
// SetPalsCountNil sets the value for PalsCount to be an explicit nil
func (o *BagDTO) SetPalsCountNil() {
	o.PalsCount.Set(nil)
}

// UnsetPalsCount ensures that no value is present for PalsCount, not even an explicit nil
func (o *BagDTO) UnsetPalsCount() {
	o.PalsCount.Unset()
}

// GetRacePalAvailable returns the RacePalAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BagDTO) GetRacePalAvailable() int32 {
	if o == nil || IsNil(o.RacePalAvailable.Get()) {
		var ret int32
		return ret
	}
	return *o.RacePalAvailable.Get()
}

// GetRacePalAvailableOk returns a tuple with the RacePalAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BagDTO) GetRacePalAvailableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RacePalAvailable.Get(), o.RacePalAvailable.IsSet()
}

// HasRacePalAvailable returns a boolean if a field has been set.
func (o *BagDTO) HasRacePalAvailable() bool {
	if o != nil && o.RacePalAvailable.IsSet() {
		return true
	}

	return false
}

// SetRacePalAvailable gets a reference to the given NullableInt32 and assigns it to the RacePalAvailable field.
func (o *BagDTO) SetRacePalAvailable(v int32) {
	o.RacePalAvailable.Set(&v)
}
// SetRacePalAvailableNil sets the value for RacePalAvailable to be an explicit nil
func (o *BagDTO) SetRacePalAvailableNil() {
	o.RacePalAvailable.Set(nil)
}

// UnsetRacePalAvailable ensures that no value is present for RacePalAvailable, not even an explicit nil
func (o *BagDTO) UnsetRacePalAvailable() {
	o.RacePalAvailable.Unset()
}

func (o BagDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BagDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EggsCount.IsSet() {
		toSerialize["eggs_count"] = o.EggsCount.Get()
	}
	if o.Empl.IsSet() {
		toSerialize["empl"] = o.Empl.Get()
	}
	if o.ExpiringEmpl.IsSet() {
		toSerialize["expiring_empl"] = o.ExpiringEmpl.Get()
	}
	if o.PalsCount.IsSet() {
		toSerialize["pals_count"] = o.PalsCount.Get()
	}
	if o.RacePalAvailable.IsSet() {
		toSerialize["race_pal_available"] = o.RacePalAvailable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BagDTO) UnmarshalJSON(data []byte) (err error) {
	varBagDTO := _BagDTO{}

	err = json.Unmarshal(data, &varBagDTO)

	if err != nil {
		return err
	}

	*o = BagDTO(varBagDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "eggs_count")
		delete(additionalProperties, "empl")
		delete(additionalProperties, "expiring_empl")
		delete(additionalProperties, "pals_count")
		delete(additionalProperties, "race_pal_available")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBagDTO struct {
	value *BagDTO
	isSet bool
}

func (v NullableBagDTO) Get() *BagDTO {
	return v.value
}

func (v *NullableBagDTO) Set(val *BagDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBagDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBagDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBagDTO(val *BagDTO) *NullableBagDTO {
	return &NullableBagDTO{value: val, isSet: true}
}

func (v NullableBagDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBagDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


