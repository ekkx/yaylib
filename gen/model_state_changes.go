
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the StateChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateChanges{}

// StateChanges struct for StateChanges
type StateChanges struct {
	Attack NullableInt32 `json:"attack,omitempty"`
	Defense NullableInt32 `json:"defense,omitempty"`
	Energy NullableInt32 `json:"energy,omitempty"`
	Luck NullableInt32 `json:"luck,omitempty"`
	Power NullableInt32 `json:"power,omitempty"`
	Speed NullableInt32 `json:"speed,omitempty"`
	Stamina NullableInt32 `json:"stamina,omitempty"`
	Technique NullableInt32 `json:"technique,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StateChanges StateChanges

// NewStateChanges instantiates a new StateChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateChanges() *StateChanges {
	this := StateChanges{}
	return &this
}

// NewStateChangesWithDefaults instantiates a new StateChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateChangesWithDefaults() *StateChanges {
	this := StateChanges{}
	return &this
}

// GetAttack returns the Attack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetAttack() int32 {
	if o == nil || IsNil(o.Attack.Get()) {
		var ret int32
		return ret
	}
	return *o.Attack.Get()
}

// GetAttackOk returns a tuple with the Attack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetAttackOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attack.Get(), o.Attack.IsSet()
}

// HasAttack returns a boolean if a field has been set.
func (o *StateChanges) HasAttack() bool {
	if o != nil && o.Attack.IsSet() {
		return true
	}

	return false
}

// SetAttack gets a reference to the given NullableInt32 and assigns it to the Attack field.
func (o *StateChanges) SetAttack(v int32) {
	o.Attack.Set(&v)
}
// SetAttackNil sets the value for Attack to be an explicit nil
func (o *StateChanges) SetAttackNil() {
	o.Attack.Set(nil)
}

// UnsetAttack ensures that no value is present for Attack, not even an explicit nil
func (o *StateChanges) UnsetAttack() {
	o.Attack.Unset()
}

// GetDefense returns the Defense field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetDefense() int32 {
	if o == nil || IsNil(o.Defense.Get()) {
		var ret int32
		return ret
	}
	return *o.Defense.Get()
}

// GetDefenseOk returns a tuple with the Defense field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetDefenseOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Defense.Get(), o.Defense.IsSet()
}

// HasDefense returns a boolean if a field has been set.
func (o *StateChanges) HasDefense() bool {
	if o != nil && o.Defense.IsSet() {
		return true
	}

	return false
}

// SetDefense gets a reference to the given NullableInt32 and assigns it to the Defense field.
func (o *StateChanges) SetDefense(v int32) {
	o.Defense.Set(&v)
}
// SetDefenseNil sets the value for Defense to be an explicit nil
func (o *StateChanges) SetDefenseNil() {
	o.Defense.Set(nil)
}

// UnsetDefense ensures that no value is present for Defense, not even an explicit nil
func (o *StateChanges) UnsetDefense() {
	o.Defense.Unset()
}

// GetEnergy returns the Energy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetEnergy() int32 {
	if o == nil || IsNil(o.Energy.Get()) {
		var ret int32
		return ret
	}
	return *o.Energy.Get()
}

// GetEnergyOk returns a tuple with the Energy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetEnergyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Energy.Get(), o.Energy.IsSet()
}

// HasEnergy returns a boolean if a field has been set.
func (o *StateChanges) HasEnergy() bool {
	if o != nil && o.Energy.IsSet() {
		return true
	}

	return false
}

// SetEnergy gets a reference to the given NullableInt32 and assigns it to the Energy field.
func (o *StateChanges) SetEnergy(v int32) {
	o.Energy.Set(&v)
}
// SetEnergyNil sets the value for Energy to be an explicit nil
func (o *StateChanges) SetEnergyNil() {
	o.Energy.Set(nil)
}

// UnsetEnergy ensures that no value is present for Energy, not even an explicit nil
func (o *StateChanges) UnsetEnergy() {
	o.Energy.Unset()
}

// GetLuck returns the Luck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetLuck() int32 {
	if o == nil || IsNil(o.Luck.Get()) {
		var ret int32
		return ret
	}
	return *o.Luck.Get()
}

// GetLuckOk returns a tuple with the Luck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetLuckOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Luck.Get(), o.Luck.IsSet()
}

// HasLuck returns a boolean if a field has been set.
func (o *StateChanges) HasLuck() bool {
	if o != nil && o.Luck.IsSet() {
		return true
	}

	return false
}

// SetLuck gets a reference to the given NullableInt32 and assigns it to the Luck field.
func (o *StateChanges) SetLuck(v int32) {
	o.Luck.Set(&v)
}
// SetLuckNil sets the value for Luck to be an explicit nil
func (o *StateChanges) SetLuckNil() {
	o.Luck.Set(nil)
}

// UnsetLuck ensures that no value is present for Luck, not even an explicit nil
func (o *StateChanges) UnsetLuck() {
	o.Luck.Unset()
}

// GetPower returns the Power field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetPower() int32 {
	if o == nil || IsNil(o.Power.Get()) {
		var ret int32
		return ret
	}
	return *o.Power.Get()
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetPowerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Power.Get(), o.Power.IsSet()
}

// HasPower returns a boolean if a field has been set.
func (o *StateChanges) HasPower() bool {
	if o != nil && o.Power.IsSet() {
		return true
	}

	return false
}

// SetPower gets a reference to the given NullableInt32 and assigns it to the Power field.
func (o *StateChanges) SetPower(v int32) {
	o.Power.Set(&v)
}
// SetPowerNil sets the value for Power to be an explicit nil
func (o *StateChanges) SetPowerNil() {
	o.Power.Set(nil)
}

// UnsetPower ensures that no value is present for Power, not even an explicit nil
func (o *StateChanges) UnsetPower() {
	o.Power.Unset()
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret int32
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *StateChanges) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableInt32 and assigns it to the Speed field.
func (o *StateChanges) SetSpeed(v int32) {
	o.Speed.Set(&v)
}
// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *StateChanges) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *StateChanges) UnsetSpeed() {
	o.Speed.Unset()
}

// GetStamina returns the Stamina field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetStamina() int32 {
	if o == nil || IsNil(o.Stamina.Get()) {
		var ret int32
		return ret
	}
	return *o.Stamina.Get()
}

// GetStaminaOk returns a tuple with the Stamina field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetStaminaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stamina.Get(), o.Stamina.IsSet()
}

// HasStamina returns a boolean if a field has been set.
func (o *StateChanges) HasStamina() bool {
	if o != nil && o.Stamina.IsSet() {
		return true
	}

	return false
}

// SetStamina gets a reference to the given NullableInt32 and assigns it to the Stamina field.
func (o *StateChanges) SetStamina(v int32) {
	o.Stamina.Set(&v)
}
// SetStaminaNil sets the value for Stamina to be an explicit nil
func (o *StateChanges) SetStaminaNil() {
	o.Stamina.Set(nil)
}

// UnsetStamina ensures that no value is present for Stamina, not even an explicit nil
func (o *StateChanges) UnsetStamina() {
	o.Stamina.Unset()
}

// GetTechnique returns the Technique field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StateChanges) GetTechnique() int32 {
	if o == nil || IsNil(o.Technique.Get()) {
		var ret int32
		return ret
	}
	return *o.Technique.Get()
}

// GetTechniqueOk returns a tuple with the Technique field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StateChanges) GetTechniqueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Technique.Get(), o.Technique.IsSet()
}

// HasTechnique returns a boolean if a field has been set.
func (o *StateChanges) HasTechnique() bool {
	if o != nil && o.Technique.IsSet() {
		return true
	}

	return false
}

// SetTechnique gets a reference to the given NullableInt32 and assigns it to the Technique field.
func (o *StateChanges) SetTechnique(v int32) {
	o.Technique.Set(&v)
}
// SetTechniqueNil sets the value for Technique to be an explicit nil
func (o *StateChanges) SetTechniqueNil() {
	o.Technique.Set(nil)
}

// UnsetTechnique ensures that no value is present for Technique, not even an explicit nil
func (o *StateChanges) UnsetTechnique() {
	o.Technique.Unset()
}

func (o StateChanges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attack.IsSet() {
		toSerialize["attack"] = o.Attack.Get()
	}
	if o.Defense.IsSet() {
		toSerialize["defense"] = o.Defense.Get()
	}
	if o.Energy.IsSet() {
		toSerialize["energy"] = o.Energy.Get()
	}
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

func (o *StateChanges) UnmarshalJSON(data []byte) (err error) {
	varStateChanges := _StateChanges{}

	err = json.Unmarshal(data, &varStateChanges)

	if err != nil {
		return err
	}

	*o = StateChanges(varStateChanges)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attack")
		delete(additionalProperties, "defense")
		delete(additionalProperties, "energy")
		delete(additionalProperties, "luck")
		delete(additionalProperties, "power")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "stamina")
		delete(additionalProperties, "technique")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStateChanges struct {
	value *StateChanges
	isSet bool
}

func (v NullableStateChanges) Get() *StateChanges {
	return v.value
}

func (v *NullableStateChanges) Set(val *StateChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableStateChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableStateChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateChanges(val *StateChanges) *NullableStateChanges {
	return &NullableStateChanges{value: val, isSet: true}
}

func (v NullableStateChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


