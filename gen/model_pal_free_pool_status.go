
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalFreePoolStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalFreePoolStatus{}

// PalFreePoolStatus struct for PalFreePoolStatus
type PalFreePoolStatus struct {
	Cooldown NullableCooldown `json:"cooldown,omitempty"`
	PalsAvailable NullableBool `json:"pals_available,omitempty"`
	Qualifications []Qualification `json:"qualifications,omitempty"`
	Qualified NullableBool `json:"qualified,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalFreePoolStatus PalFreePoolStatus

// NewPalFreePoolStatus instantiates a new PalFreePoolStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalFreePoolStatus() *PalFreePoolStatus {
	this := PalFreePoolStatus{}
	return &this
}

// NewPalFreePoolStatusWithDefaults instantiates a new PalFreePoolStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalFreePoolStatusWithDefaults() *PalFreePoolStatus {
	this := PalFreePoolStatus{}
	return &this
}

// GetCooldown returns the Cooldown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalFreePoolStatus) GetCooldown() Cooldown {
	if o == nil || IsNil(o.Cooldown.Get()) {
		var ret Cooldown
		return ret
	}
	return *o.Cooldown.Get()
}

// GetCooldownOk returns a tuple with the Cooldown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalFreePoolStatus) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cooldown.Get(), o.Cooldown.IsSet()
}

// HasCooldown returns a boolean if a field has been set.
func (o *PalFreePoolStatus) HasCooldown() bool {
	if o != nil && o.Cooldown.IsSet() {
		return true
	}

	return false
}

// SetCooldown gets a reference to the given NullableCooldown and assigns it to the Cooldown field.
func (o *PalFreePoolStatus) SetCooldown(v Cooldown) {
	o.Cooldown.Set(&v)
}
// SetCooldownNil sets the value for Cooldown to be an explicit nil
func (o *PalFreePoolStatus) SetCooldownNil() {
	o.Cooldown.Set(nil)
}

// UnsetCooldown ensures that no value is present for Cooldown, not even an explicit nil
func (o *PalFreePoolStatus) UnsetCooldown() {
	o.Cooldown.Unset()
}

// GetPalsAvailable returns the PalsAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalFreePoolStatus) GetPalsAvailable() bool {
	if o == nil || IsNil(o.PalsAvailable.Get()) {
		var ret bool
		return ret
	}
	return *o.PalsAvailable.Get()
}

// GetPalsAvailableOk returns a tuple with the PalsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalFreePoolStatus) GetPalsAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalsAvailable.Get(), o.PalsAvailable.IsSet()
}

// HasPalsAvailable returns a boolean if a field has been set.
func (o *PalFreePoolStatus) HasPalsAvailable() bool {
	if o != nil && o.PalsAvailable.IsSet() {
		return true
	}

	return false
}

// SetPalsAvailable gets a reference to the given NullableBool and assigns it to the PalsAvailable field.
func (o *PalFreePoolStatus) SetPalsAvailable(v bool) {
	o.PalsAvailable.Set(&v)
}
// SetPalsAvailableNil sets the value for PalsAvailable to be an explicit nil
func (o *PalFreePoolStatus) SetPalsAvailableNil() {
	o.PalsAvailable.Set(nil)
}

// UnsetPalsAvailable ensures that no value is present for PalsAvailable, not even an explicit nil
func (o *PalFreePoolStatus) UnsetPalsAvailable() {
	o.PalsAvailable.Unset()
}

// GetQualifications returns the Qualifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalFreePoolStatus) GetQualifications() []Qualification {
	if o == nil {
		var ret []Qualification
		return ret
	}
	return o.Qualifications
}

// GetQualificationsOk returns a tuple with the Qualifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalFreePoolStatus) GetQualificationsOk() ([]Qualification, bool) {
	if o == nil || IsNil(o.Qualifications) {
		return nil, false
	}
	return o.Qualifications, true
}

// HasQualifications returns a boolean if a field has been set.
func (o *PalFreePoolStatus) HasQualifications() bool {
	if o != nil && !IsNil(o.Qualifications) {
		return true
	}

	return false
}

// SetQualifications gets a reference to the given []Qualification and assigns it to the Qualifications field.
func (o *PalFreePoolStatus) SetQualifications(v []Qualification) {
	o.Qualifications = v
}

// GetQualified returns the Qualified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalFreePoolStatus) GetQualified() bool {
	if o == nil || IsNil(o.Qualified.Get()) {
		var ret bool
		return ret
	}
	return *o.Qualified.Get()
}

// GetQualifiedOk returns a tuple with the Qualified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalFreePoolStatus) GetQualifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Qualified.Get(), o.Qualified.IsSet()
}

// HasQualified returns a boolean if a field has been set.
func (o *PalFreePoolStatus) HasQualified() bool {
	if o != nil && o.Qualified.IsSet() {
		return true
	}

	return false
}

// SetQualified gets a reference to the given NullableBool and assigns it to the Qualified field.
func (o *PalFreePoolStatus) SetQualified(v bool) {
	o.Qualified.Set(&v)
}
// SetQualifiedNil sets the value for Qualified to be an explicit nil
func (o *PalFreePoolStatus) SetQualifiedNil() {
	o.Qualified.Set(nil)
}

// UnsetQualified ensures that no value is present for Qualified, not even an explicit nil
func (o *PalFreePoolStatus) UnsetQualified() {
	o.Qualified.Unset()
}

func (o PalFreePoolStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalFreePoolStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cooldown.IsSet() {
		toSerialize["cooldown"] = o.Cooldown.Get()
	}
	if o.PalsAvailable.IsSet() {
		toSerialize["pals_available"] = o.PalsAvailable.Get()
	}
	if o.Qualifications != nil {
		toSerialize["qualifications"] = o.Qualifications
	}
	if o.Qualified.IsSet() {
		toSerialize["qualified"] = o.Qualified.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalFreePoolStatus) UnmarshalJSON(data []byte) (err error) {
	varPalFreePoolStatus := _PalFreePoolStatus{}

	err = json.Unmarshal(data, &varPalFreePoolStatus)

	if err != nil {
		return err
	}

	*o = PalFreePoolStatus(varPalFreePoolStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cooldown")
		delete(additionalProperties, "pals_available")
		delete(additionalProperties, "qualifications")
		delete(additionalProperties, "qualified")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalFreePoolStatus struct {
	value *PalFreePoolStatus
	isSet bool
}

func (v NullablePalFreePoolStatus) Get() *PalFreePoolStatus {
	return v.value
}

func (v *NullablePalFreePoolStatus) Set(val *PalFreePoolStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePalFreePoolStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePalFreePoolStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalFreePoolStatus(val *PalFreePoolStatus) *NullablePalFreePoolStatus {
	return &NullablePalFreePoolStatus{value: val, isSet: true}
}

func (v NullablePalFreePoolStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalFreePoolStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


