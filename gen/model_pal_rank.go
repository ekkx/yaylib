
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRank type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRank{}

// PalRank struct for PalRank
type PalRank struct {
	Own NullableBool `json:"own,omitempty"`
	Pal NullablePal `json:"pal,omitempty"`
	Place NullableInt32 `json:"place,omitempty"`
	YayUserName NullableString `json:"yay_user_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRank PalRank

// NewPalRank instantiates a new PalRank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRank() *PalRank {
	this := PalRank{}
	return &this
}

// NewPalRankWithDefaults instantiates a new PalRank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRankWithDefaults() *PalRank {
	this := PalRank{}
	return &this
}

// GetOwn returns the Own field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRank) GetOwn() bool {
	if o == nil || IsNil(o.Own.Get()) {
		var ret bool
		return ret
	}
	return *o.Own.Get()
}

// GetOwnOk returns a tuple with the Own field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRank) GetOwnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Own.Get(), o.Own.IsSet()
}

// HasOwn returns a boolean if a field has been set.
func (o *PalRank) HasOwn() bool {
	if o != nil && o.Own.IsSet() {
		return true
	}

	return false
}

// SetOwn gets a reference to the given NullableBool and assigns it to the Own field.
func (o *PalRank) SetOwn(v bool) {
	o.Own.Set(&v)
}
// SetOwnNil sets the value for Own to be an explicit nil
func (o *PalRank) SetOwnNil() {
	o.Own.Set(nil)
}

// UnsetOwn ensures that no value is present for Own, not even an explicit nil
func (o *PalRank) UnsetOwn() {
	o.Own.Unset()
}

// GetPal returns the Pal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRank) GetPal() Pal {
	if o == nil || IsNil(o.Pal.Get()) {
		var ret Pal
		return ret
	}
	return *o.Pal.Get()
}

// GetPalOk returns a tuple with the Pal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRank) GetPalOk() (*Pal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pal.Get(), o.Pal.IsSet()
}

// HasPal returns a boolean if a field has been set.
func (o *PalRank) HasPal() bool {
	if o != nil && o.Pal.IsSet() {
		return true
	}

	return false
}

// SetPal gets a reference to the given NullablePal and assigns it to the Pal field.
func (o *PalRank) SetPal(v Pal) {
	o.Pal.Set(&v)
}
// SetPalNil sets the value for Pal to be an explicit nil
func (o *PalRank) SetPalNil() {
	o.Pal.Set(nil)
}

// UnsetPal ensures that no value is present for Pal, not even an explicit nil
func (o *PalRank) UnsetPal() {
	o.Pal.Unset()
}

// GetPlace returns the Place field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRank) GetPlace() int32 {
	if o == nil || IsNil(o.Place.Get()) {
		var ret int32
		return ret
	}
	return *o.Place.Get()
}

// GetPlaceOk returns a tuple with the Place field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRank) GetPlaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Place.Get(), o.Place.IsSet()
}

// HasPlace returns a boolean if a field has been set.
func (o *PalRank) HasPlace() bool {
	if o != nil && o.Place.IsSet() {
		return true
	}

	return false
}

// SetPlace gets a reference to the given NullableInt32 and assigns it to the Place field.
func (o *PalRank) SetPlace(v int32) {
	o.Place.Set(&v)
}
// SetPlaceNil sets the value for Place to be an explicit nil
func (o *PalRank) SetPlaceNil() {
	o.Place.Set(nil)
}

// UnsetPlace ensures that no value is present for Place, not even an explicit nil
func (o *PalRank) UnsetPlace() {
	o.Place.Unset()
}

// GetYayUserName returns the YayUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRank) GetYayUserName() string {
	if o == nil || IsNil(o.YayUserName.Get()) {
		var ret string
		return ret
	}
	return *o.YayUserName.Get()
}

// GetYayUserNameOk returns a tuple with the YayUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRank) GetYayUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.YayUserName.Get(), o.YayUserName.IsSet()
}

// HasYayUserName returns a boolean if a field has been set.
func (o *PalRank) HasYayUserName() bool {
	if o != nil && o.YayUserName.IsSet() {
		return true
	}

	return false
}

// SetYayUserName gets a reference to the given NullableString and assigns it to the YayUserName field.
func (o *PalRank) SetYayUserName(v string) {
	o.YayUserName.Set(&v)
}
// SetYayUserNameNil sets the value for YayUserName to be an explicit nil
func (o *PalRank) SetYayUserNameNil() {
	o.YayUserName.Set(nil)
}

// UnsetYayUserName ensures that no value is present for YayUserName, not even an explicit nil
func (o *PalRank) UnsetYayUserName() {
	o.YayUserName.Unset()
}

func (o PalRank) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Own.IsSet() {
		toSerialize["own"] = o.Own.Get()
	}
	if o.Pal.IsSet() {
		toSerialize["pal"] = o.Pal.Get()
	}
	if o.Place.IsSet() {
		toSerialize["place"] = o.Place.Get()
	}
	if o.YayUserName.IsSet() {
		toSerialize["yay_user_name"] = o.YayUserName.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRank) UnmarshalJSON(data []byte) (err error) {
	varPalRank := _PalRank{}

	err = json.Unmarshal(data, &varPalRank)

	if err != nil {
		return err
	}

	*o = PalRank(varPalRank)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "own")
		delete(additionalProperties, "pal")
		delete(additionalProperties, "place")
		delete(additionalProperties, "yay_user_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRank struct {
	value *PalRank
	isSet bool
}

func (v NullablePalRank) Get() *PalRank {
	return v.value
}

func (v *NullablePalRank) Set(val *PalRank) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRank) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRank(val *PalRank) *NullablePalRank {
	return &NullablePalRank{value: val, isSet: true}
}

func (v NullablePalRank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


