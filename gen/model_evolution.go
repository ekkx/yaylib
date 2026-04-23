
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Evolution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Evolution{}

// Evolution struct for Evolution
type Evolution struct {
	// Unresolved Java type: jp.nanameue.yay.shared.data.model.blockchain.Pal.PalGrade
	Grade map[string]interface{} `json:"grade,omitempty"`
	Level NullableInt32 `json:"level,omitempty"`
	PalImageUrl NullableString `json:"pal_image_url,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Evolution Evolution

// NewEvolution instantiates a new Evolution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvolution() *Evolution {
	this := Evolution{}
	return &this
}

// NewEvolutionWithDefaults instantiates a new Evolution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvolutionWithDefaults() *Evolution {
	this := Evolution{}
	return &this
}

// GetGrade returns the Grade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Evolution) GetGrade() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Grade
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Evolution) GetGradeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Grade) {
		return map[string]interface{}{}, false
	}
	return o.Grade, true
}

// HasGrade returns a boolean if a field has been set.
func (o *Evolution) HasGrade() bool {
	if o != nil && !IsNil(o.Grade) {
		return true
	}

	return false
}

// SetGrade gets a reference to the given map[string]interface{} and assigns it to the Grade field.
func (o *Evolution) SetGrade(v map[string]interface{}) {
	o.Grade = v
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Evolution) GetLevel() int32 {
	if o == nil || IsNil(o.Level.Get()) {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Evolution) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *Evolution) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *Evolution) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *Evolution) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *Evolution) UnsetLevel() {
	o.Level.Unset()
}

// GetPalImageUrl returns the PalImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Evolution) GetPalImageUrl() string {
	if o == nil || IsNil(o.PalImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PalImageUrl.Get()
}

// GetPalImageUrlOk returns a tuple with the PalImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Evolution) GetPalImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImageUrl.Get(), o.PalImageUrl.IsSet()
}

// HasPalImageUrl returns a boolean if a field has been set.
func (o *Evolution) HasPalImageUrl() bool {
	if o != nil && o.PalImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPalImageUrl gets a reference to the given NullableString and assigns it to the PalImageUrl field.
func (o *Evolution) SetPalImageUrl(v string) {
	o.PalImageUrl.Set(&v)
}
// SetPalImageUrlNil sets the value for PalImageUrl to be an explicit nil
func (o *Evolution) SetPalImageUrlNil() {
	o.PalImageUrl.Set(nil)
}

// UnsetPalImageUrl ensures that no value is present for PalImageUrl, not even an explicit nil
func (o *Evolution) UnsetPalImageUrl() {
	o.PalImageUrl.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Evolution) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Evolution) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *Evolution) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *Evolution) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *Evolution) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *Evolution) UnsetPalName() {
	o.PalName.Unset()
}

func (o Evolution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Evolution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Grade != nil {
		toSerialize["grade"] = o.Grade
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.PalImageUrl.IsSet() {
		toSerialize["pal_image_url"] = o.PalImageUrl.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Evolution) UnmarshalJSON(data []byte) (err error) {
	varEvolution := _Evolution{}

	err = json.Unmarshal(data, &varEvolution)

	if err != nil {
		return err
	}

	*o = Evolution(varEvolution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grade")
		delete(additionalProperties, "level")
		delete(additionalProperties, "pal_image_url")
		delete(additionalProperties, "pal_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEvolution struct {
	value *Evolution
	isSet bool
}

func (v NullableEvolution) Get() *Evolution {
	return v.value
}

func (v *NullableEvolution) Set(val *Evolution) {
	v.value = val
	v.isSet = true
}

func (v NullableEvolution) IsSet() bool {
	return v.isSet
}

func (v *NullableEvolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvolution(val *Evolution) *NullableEvolution {
	return &NullableEvolution{value: val, isSet: true}
}

func (v NullableEvolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


