
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the AdditionGasPercentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionGasPercentDTO{}

// AdditionGasPercentDTO struct for AdditionGasPercentDTO
type AdditionGasPercentDTO struct {
	Fast NullableInt32 `json:"fast,omitempty"`
	Normal NullableInt32 `json:"normal,omitempty"`
	Slow NullableInt32 `json:"slow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdditionGasPercentDTO AdditionGasPercentDTO

// NewAdditionGasPercentDTO instantiates a new AdditionGasPercentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionGasPercentDTO() *AdditionGasPercentDTO {
	this := AdditionGasPercentDTO{}
	return &this
}

// NewAdditionGasPercentDTOWithDefaults instantiates a new AdditionGasPercentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionGasPercentDTOWithDefaults() *AdditionGasPercentDTO {
	this := AdditionGasPercentDTO{}
	return &this
}

// GetFast returns the Fast field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionGasPercentDTO) GetFast() int32 {
	if o == nil || IsNil(o.Fast.Get()) {
		var ret int32
		return ret
	}
	return *o.Fast.Get()
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionGasPercentDTO) GetFastOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fast.Get(), o.Fast.IsSet()
}

// HasFast returns a boolean if a field has been set.
func (o *AdditionGasPercentDTO) HasFast() bool {
	if o != nil && o.Fast.IsSet() {
		return true
	}

	return false
}

// SetFast gets a reference to the given NullableInt32 and assigns it to the Fast field.
func (o *AdditionGasPercentDTO) SetFast(v int32) {
	o.Fast.Set(&v)
}
// SetFastNil sets the value for Fast to be an explicit nil
func (o *AdditionGasPercentDTO) SetFastNil() {
	o.Fast.Set(nil)
}

// UnsetFast ensures that no value is present for Fast, not even an explicit nil
func (o *AdditionGasPercentDTO) UnsetFast() {
	o.Fast.Unset()
}

// GetNormal returns the Normal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionGasPercentDTO) GetNormal() int32 {
	if o == nil || IsNil(o.Normal.Get()) {
		var ret int32
		return ret
	}
	return *o.Normal.Get()
}

// GetNormalOk returns a tuple with the Normal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionGasPercentDTO) GetNormalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Normal.Get(), o.Normal.IsSet()
}

// HasNormal returns a boolean if a field has been set.
func (o *AdditionGasPercentDTO) HasNormal() bool {
	if o != nil && o.Normal.IsSet() {
		return true
	}

	return false
}

// SetNormal gets a reference to the given NullableInt32 and assigns it to the Normal field.
func (o *AdditionGasPercentDTO) SetNormal(v int32) {
	o.Normal.Set(&v)
}
// SetNormalNil sets the value for Normal to be an explicit nil
func (o *AdditionGasPercentDTO) SetNormalNil() {
	o.Normal.Set(nil)
}

// UnsetNormal ensures that no value is present for Normal, not even an explicit nil
func (o *AdditionGasPercentDTO) UnsetNormal() {
	o.Normal.Unset()
}

// GetSlow returns the Slow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionGasPercentDTO) GetSlow() int32 {
	if o == nil || IsNil(o.Slow.Get()) {
		var ret int32
		return ret
	}
	return *o.Slow.Get()
}

// GetSlowOk returns a tuple with the Slow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionGasPercentDTO) GetSlowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slow.Get(), o.Slow.IsSet()
}

// HasSlow returns a boolean if a field has been set.
func (o *AdditionGasPercentDTO) HasSlow() bool {
	if o != nil && o.Slow.IsSet() {
		return true
	}

	return false
}

// SetSlow gets a reference to the given NullableInt32 and assigns it to the Slow field.
func (o *AdditionGasPercentDTO) SetSlow(v int32) {
	o.Slow.Set(&v)
}
// SetSlowNil sets the value for Slow to be an explicit nil
func (o *AdditionGasPercentDTO) SetSlowNil() {
	o.Slow.Set(nil)
}

// UnsetSlow ensures that no value is present for Slow, not even an explicit nil
func (o *AdditionGasPercentDTO) UnsetSlow() {
	o.Slow.Unset()
}

func (o AdditionGasPercentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionGasPercentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Fast.IsSet() {
		toSerialize["fast"] = o.Fast.Get()
	}
	if o.Normal.IsSet() {
		toSerialize["normal"] = o.Normal.Get()
	}
	if o.Slow.IsSet() {
		toSerialize["slow"] = o.Slow.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdditionGasPercentDTO) UnmarshalJSON(data []byte) (err error) {
	varAdditionGasPercentDTO := _AdditionGasPercentDTO{}

	err = json.Unmarshal(data, &varAdditionGasPercentDTO)

	if err != nil {
		return err
	}

	*o = AdditionGasPercentDTO(varAdditionGasPercentDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fast")
		delete(additionalProperties, "normal")
		delete(additionalProperties, "slow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdditionGasPercentDTO struct {
	value *AdditionGasPercentDTO
	isSet bool
}

func (v NullableAdditionGasPercentDTO) Get() *AdditionGasPercentDTO {
	return v.value
}

func (v *NullableAdditionGasPercentDTO) Set(val *AdditionGasPercentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionGasPercentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionGasPercentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionGasPercentDTO(val *AdditionGasPercentDTO) *NullableAdditionGasPercentDTO {
	return &NullableAdditionGasPercentDTO{value: val, isSet: true}
}

func (v NullableAdditionGasPercentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionGasPercentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


