
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LevelUp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LevelUp{}

// LevelUp struct for LevelUp
type LevelUp struct {
	Level NullableInt32 `json:"level,omitempty"`
	PalImageUrl NullableString `json:"pal_image_url,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	Price NullableFloat64 `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LevelUp LevelUp

// NewLevelUp instantiates a new LevelUp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLevelUp() *LevelUp {
	this := LevelUp{}
	return &this
}

// NewLevelUpWithDefaults instantiates a new LevelUp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLevelUpWithDefaults() *LevelUp {
	this := LevelUp{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUp) GetLevel() int32 {
	if o == nil || IsNil(o.Level.Get()) {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUp) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *LevelUp) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *LevelUp) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *LevelUp) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *LevelUp) UnsetLevel() {
	o.Level.Unset()
}

// GetPalImageUrl returns the PalImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUp) GetPalImageUrl() string {
	if o == nil || IsNil(o.PalImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PalImageUrl.Get()
}

// GetPalImageUrlOk returns a tuple with the PalImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUp) GetPalImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImageUrl.Get(), o.PalImageUrl.IsSet()
}

// HasPalImageUrl returns a boolean if a field has been set.
func (o *LevelUp) HasPalImageUrl() bool {
	if o != nil && o.PalImageUrl.IsSet() {
		return true
	}

	return false
}

// SetPalImageUrl gets a reference to the given NullableString and assigns it to the PalImageUrl field.
func (o *LevelUp) SetPalImageUrl(v string) {
	o.PalImageUrl.Set(&v)
}
// SetPalImageUrlNil sets the value for PalImageUrl to be an explicit nil
func (o *LevelUp) SetPalImageUrlNil() {
	o.PalImageUrl.Set(nil)
}

// UnsetPalImageUrl ensures that no value is present for PalImageUrl, not even an explicit nil
func (o *LevelUp) UnsetPalImageUrl() {
	o.PalImageUrl.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUp) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUp) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *LevelUp) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *LevelUp) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *LevelUp) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *LevelUp) UnsetPalName() {
	o.PalName.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LevelUp) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LevelUp) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *LevelUp) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *LevelUp) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *LevelUp) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *LevelUp) UnsetPrice() {
	o.Price.Unset()
}

func (o LevelUp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LevelUp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.PalImageUrl.IsSet() {
		toSerialize["pal_image_url"] = o.PalImageUrl.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LevelUp) UnmarshalJSON(data []byte) (err error) {
	varLevelUp := _LevelUp{}

	err = json.Unmarshal(data, &varLevelUp)

	if err != nil {
		return err
	}

	*o = LevelUp(varLevelUp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "level")
		delete(additionalProperties, "pal_image_url")
		delete(additionalProperties, "pal_name")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLevelUp struct {
	value *LevelUp
	isSet bool
}

func (v NullableLevelUp) Get() *LevelUp {
	return v.value
}

func (v *NullableLevelUp) Set(val *LevelUp) {
	v.value = val
	v.isSet = true
}

func (v NullableLevelUp) IsSet() bool {
	return v.isSet
}

func (v *NullableLevelUp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLevelUp(val *LevelUp) *NullableLevelUp {
	return &NullableLevelUp{value: val, isSet: true}
}

func (v NullableLevelUp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLevelUp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


