
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MaxAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaxAttribute{}

// MaxAttribute struct for MaxAttribute
type MaxAttribute struct {
	Luck NullableInt32 `json:"luck,omitempty"`
	Power NullableInt32 `json:"power,omitempty"`
	Speed NullableInt32 `json:"speed,omitempty"`
	Stamina NullableInt32 `json:"stamina,omitempty"`
	Technique NullableInt32 `json:"technique,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MaxAttribute MaxAttribute

// NewMaxAttribute instantiates a new MaxAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxAttribute() *MaxAttribute {
	this := MaxAttribute{}
	return &this
}

// NewMaxAttributeWithDefaults instantiates a new MaxAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxAttributeWithDefaults() *MaxAttribute {
	this := MaxAttribute{}
	return &this
}

// GetLuck returns the Luck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaxAttribute) GetLuck() int32 {
	if o == nil || IsNil(o.Luck.Get()) {
		var ret int32
		return ret
	}
	return *o.Luck.Get()
}

// GetLuckOk returns a tuple with the Luck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaxAttribute) GetLuckOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Luck.Get(), o.Luck.IsSet()
}

// HasLuck returns a boolean if a field has been set.
func (o *MaxAttribute) HasLuck() bool {
	if o != nil && o.Luck.IsSet() {
		return true
	}

	return false
}

// SetLuck gets a reference to the given NullableInt32 and assigns it to the Luck field.
func (o *MaxAttribute) SetLuck(v int32) {
	o.Luck.Set(&v)
}
// SetLuckNil sets the value for Luck to be an explicit nil
func (o *MaxAttribute) SetLuckNil() {
	o.Luck.Set(nil)
}

// UnsetLuck ensures that no value is present for Luck, not even an explicit nil
func (o *MaxAttribute) UnsetLuck() {
	o.Luck.Unset()
}

// GetPower returns the Power field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaxAttribute) GetPower() int32 {
	if o == nil || IsNil(o.Power.Get()) {
		var ret int32
		return ret
	}
	return *o.Power.Get()
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaxAttribute) GetPowerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Power.Get(), o.Power.IsSet()
}

// HasPower returns a boolean if a field has been set.
func (o *MaxAttribute) HasPower() bool {
	if o != nil && o.Power.IsSet() {
		return true
	}

	return false
}

// SetPower gets a reference to the given NullableInt32 and assigns it to the Power field.
func (o *MaxAttribute) SetPower(v int32) {
	o.Power.Set(&v)
}
// SetPowerNil sets the value for Power to be an explicit nil
func (o *MaxAttribute) SetPowerNil() {
	o.Power.Set(nil)
}

// UnsetPower ensures that no value is present for Power, not even an explicit nil
func (o *MaxAttribute) UnsetPower() {
	o.Power.Unset()
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaxAttribute) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret int32
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaxAttribute) GetSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *MaxAttribute) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableInt32 and assigns it to the Speed field.
func (o *MaxAttribute) SetSpeed(v int32) {
	o.Speed.Set(&v)
}
// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *MaxAttribute) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *MaxAttribute) UnsetSpeed() {
	o.Speed.Unset()
}

// GetStamina returns the Stamina field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaxAttribute) GetStamina() int32 {
	if o == nil || IsNil(o.Stamina.Get()) {
		var ret int32
		return ret
	}
	return *o.Stamina.Get()
}

// GetStaminaOk returns a tuple with the Stamina field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaxAttribute) GetStaminaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stamina.Get(), o.Stamina.IsSet()
}

// HasStamina returns a boolean if a field has been set.
func (o *MaxAttribute) HasStamina() bool {
	if o != nil && o.Stamina.IsSet() {
		return true
	}

	return false
}

// SetStamina gets a reference to the given NullableInt32 and assigns it to the Stamina field.
func (o *MaxAttribute) SetStamina(v int32) {
	o.Stamina.Set(&v)
}
// SetStaminaNil sets the value for Stamina to be an explicit nil
func (o *MaxAttribute) SetStaminaNil() {
	o.Stamina.Set(nil)
}

// UnsetStamina ensures that no value is present for Stamina, not even an explicit nil
func (o *MaxAttribute) UnsetStamina() {
	o.Stamina.Unset()
}

// GetTechnique returns the Technique field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaxAttribute) GetTechnique() int32 {
	if o == nil || IsNil(o.Technique.Get()) {
		var ret int32
		return ret
	}
	return *o.Technique.Get()
}

// GetTechniqueOk returns a tuple with the Technique field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaxAttribute) GetTechniqueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Technique.Get(), o.Technique.IsSet()
}

// HasTechnique returns a boolean if a field has been set.
func (o *MaxAttribute) HasTechnique() bool {
	if o != nil && o.Technique.IsSet() {
		return true
	}

	return false
}

// SetTechnique gets a reference to the given NullableInt32 and assigns it to the Technique field.
func (o *MaxAttribute) SetTechnique(v int32) {
	o.Technique.Set(&v)
}
// SetTechniqueNil sets the value for Technique to be an explicit nil
func (o *MaxAttribute) SetTechniqueNil() {
	o.Technique.Set(nil)
}

// UnsetTechnique ensures that no value is present for Technique, not even an explicit nil
func (o *MaxAttribute) UnsetTechnique() {
	o.Technique.Unset()
}

func (o MaxAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Luck.IsSet() {
		toSerialize["luck"] = o.Luck.Get()
	}
	if o.Power.IsSet() {
		toSerialize["power"] = o.Power.Get()
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if o.Stamina.IsSet() {
		toSerialize["stamina"] = o.Stamina.Get()
	}
	if o.Technique.IsSet() {
		toSerialize["technique"] = o.Technique.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MaxAttribute) UnmarshalJSON(data []byte) (err error) {
	varMaxAttribute := _MaxAttribute{}

	err = json.Unmarshal(data, &varMaxAttribute)

	if err != nil {
		return err
	}

	*o = MaxAttribute(varMaxAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "luck")
		delete(additionalProperties, "power")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "stamina")
		delete(additionalProperties, "technique")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMaxAttribute struct {
	value *MaxAttribute
	isSet bool
}

func (v NullableMaxAttribute) Get() *MaxAttribute {
	return v.value
}

func (v *NullableMaxAttribute) Set(val *MaxAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxAttribute(val *MaxAttribute) *NullableMaxAttribute {
	return &NullableMaxAttribute{value: val, isSet: true}
}

func (v NullableMaxAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


