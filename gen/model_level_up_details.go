
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LevelUpDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LevelUpDetails{}

// LevelUpDetails struct for LevelUpDetails
type LevelUpDetails struct {
	Available NullableBool `json:"available,omitempty"`
	Cost NullableString `json:"cost,omitempty"`
	EmplEarnLimit NullableString `json:"empl_earn_limit,omitempty"`
	Level NullableInt32 `json:"level,omitempty"`
	Pal NullableLevelUpDetailsPal `json:"pal,omitempty"`
	StatChanges NullableStateChanges `json:"stat_changes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LevelUpDetails LevelUpDetails

// NewLevelUpDetails instantiates a new LevelUpDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLevelUpDetails() *LevelUpDetails {
	this := LevelUpDetails{}
	return &this
}

// NewLevelUpDetailsWithDefaults instantiates a new LevelUpDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLevelUpDetailsWithDefaults() *LevelUpDetails {
	this := LevelUpDetails{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUpDetails) GetAvailable() bool {
	if o == nil || IsNil(o.Available.Get()) {
		var ret bool
		return ret
	}
	return *o.Available.Get()
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUpDetails) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Available.Get(), o.Available.IsSet()
}

// HasAvailable returns a boolean if a field has been set.
func (o *LevelUpDetails) HasAvailable() bool {
	if o != nil && o.Available.IsSet() {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given NullableBool and assigns it to the Available field.
func (o *LevelUpDetails) SetAvailable(v bool) {
	o.Available.Set(&v)
}
// SetAvailableNil sets the value for Available to be an explicit nil
func (o *LevelUpDetails) SetAvailableNil() {
	o.Available.Set(nil)
}

// UnsetAvailable ensures that no value is present for Available, not even an explicit nil
func (o *LevelUpDetails) UnsetAvailable() {
	o.Available.Unset()
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUpDetails) GetCost() string {
	if o == nil || IsNil(o.Cost.Get()) {
		var ret string
		return ret
	}
	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUpDetails) GetCostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// HasCost returns a boolean if a field has been set.
func (o *LevelUpDetails) HasCost() bool {
	if o != nil && o.Cost.IsSet() {
		return true
	}

	return false
}

// SetCost gets a reference to the given NullableString and assigns it to the Cost field.
func (o *LevelUpDetails) SetCost(v string) {
	o.Cost.Set(&v)
}
// SetCostNil sets the value for Cost to be an explicit nil
func (o *LevelUpDetails) SetCostNil() {
	o.Cost.Set(nil)
}

// UnsetCost ensures that no value is present for Cost, not even an explicit nil
func (o *LevelUpDetails) UnsetCost() {
	o.Cost.Unset()
}

// GetEmplEarnLimit returns the EmplEarnLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUpDetails) GetEmplEarnLimit() string {
	if o == nil || IsNil(o.EmplEarnLimit.Get()) {
		var ret string
		return ret
	}
	return *o.EmplEarnLimit.Get()
}

// GetEmplEarnLimitOk returns a tuple with the EmplEarnLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUpDetails) GetEmplEarnLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplEarnLimit.Get(), o.EmplEarnLimit.IsSet()
}

// HasEmplEarnLimit returns a boolean if a field has been set.
func (o *LevelUpDetails) HasEmplEarnLimit() bool {
	if o != nil && o.EmplEarnLimit.IsSet() {
		return true
	}

	return false
}

// SetEmplEarnLimit gets a reference to the given NullableString and assigns it to the EmplEarnLimit field.
func (o *LevelUpDetails) SetEmplEarnLimit(v string) {
	o.EmplEarnLimit.Set(&v)
}
// SetEmplEarnLimitNil sets the value for EmplEarnLimit to be an explicit nil
func (o *LevelUpDetails) SetEmplEarnLimitNil() {
	o.EmplEarnLimit.Set(nil)
}

// UnsetEmplEarnLimit ensures that no value is present for EmplEarnLimit, not even an explicit nil
func (o *LevelUpDetails) UnsetEmplEarnLimit() {
	o.EmplEarnLimit.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUpDetails) GetLevel() int32 {
	if o == nil || IsNil(o.Level.Get()) {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUpDetails) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *LevelUpDetails) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *LevelUpDetails) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *LevelUpDetails) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *LevelUpDetails) UnsetLevel() {
	o.Level.Unset()
}

// GetPal returns the Pal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUpDetails) GetPal() LevelUpDetailsPal {
	if o == nil || IsNil(o.Pal.Get()) {
		var ret LevelUpDetailsPal
		return ret
	}
	return *o.Pal.Get()
}

// GetPalOk returns a tuple with the Pal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUpDetails) GetPalOk() (*LevelUpDetailsPal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pal.Get(), o.Pal.IsSet()
}

// HasPal returns a boolean if a field has been set.
func (o *LevelUpDetails) HasPal() bool {
	if o != nil && o.Pal.IsSet() {
		return true
	}

	return false
}

// SetPal gets a reference to the given NullableLevelUpDetailsPal and assigns it to the Pal field.
func (o *LevelUpDetails) SetPal(v LevelUpDetailsPal) {
	o.Pal.Set(&v)
}
// SetPalNil sets the value for Pal to be an explicit nil
func (o *LevelUpDetails) SetPalNil() {
	o.Pal.Set(nil)
}

// UnsetPal ensures that no value is present for Pal, not even an explicit nil
func (o *LevelUpDetails) UnsetPal() {
	o.Pal.Unset()
}

// GetStatChanges returns the StatChanges field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUpDetails) GetStatChanges() StateChanges {
	if o == nil || IsNil(o.StatChanges.Get()) {
		var ret StateChanges
		return ret
	}
	return *o.StatChanges.Get()
}

// GetStatChangesOk returns a tuple with the StatChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUpDetails) GetStatChangesOk() (*StateChanges, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatChanges.Get(), o.StatChanges.IsSet()
}

// HasStatChanges returns a boolean if a field has been set.
func (o *LevelUpDetails) HasStatChanges() bool {
	if o != nil && o.StatChanges.IsSet() {
		return true
	}

	return false
}

// SetStatChanges gets a reference to the given NullableStateChanges and assigns it to the StatChanges field.
func (o *LevelUpDetails) SetStatChanges(v StateChanges) {
	o.StatChanges.Set(&v)
}
// SetStatChangesNil sets the value for StatChanges to be an explicit nil
func (o *LevelUpDetails) SetStatChangesNil() {
	o.StatChanges.Set(nil)
}

// UnsetStatChanges ensures that no value is present for StatChanges, not even an explicit nil
func (o *LevelUpDetails) UnsetStatChanges() {
	o.StatChanges.Unset()
}

func (o LevelUpDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LevelUpDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Available.IsSet() {
		toSerialize["available"] = o.Available.Get()
	}
	if o.Cost.IsSet() {
		toSerialize["cost"] = o.Cost.Get()
	}
	if o.EmplEarnLimit.IsSet() {
		toSerialize["empl_earn_limit"] = o.EmplEarnLimit.Get()
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.Pal.IsSet() {
		toSerialize["pal"] = o.Pal.Get()
	}
	if o.StatChanges.IsSet() {
		toSerialize["stat_changes"] = o.StatChanges.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LevelUpDetails) UnmarshalJSON(data []byte) (err error) {
	varLevelUpDetails := _LevelUpDetails{}

	err = json.Unmarshal(data, &varLevelUpDetails)

	if err != nil {
		return err
	}

	*o = LevelUpDetails(varLevelUpDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "available")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "empl_earn_limit")
		delete(additionalProperties, "level")
		delete(additionalProperties, "pal")
		delete(additionalProperties, "stat_changes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLevelUpDetails struct {
	value *LevelUpDetails
	isSet bool
}

func (v NullableLevelUpDetails) Get() *LevelUpDetails {
	return v.value
}

func (v *NullableLevelUpDetails) Set(val *LevelUpDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLevelUpDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLevelUpDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLevelUpDetails(val *LevelUpDetails) *NullableLevelUpDetails {
	return &NullableLevelUpDetails{value: val, isSet: true}
}

func (v NullableLevelUpDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLevelUpDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


