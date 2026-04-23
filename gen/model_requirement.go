
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Requirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Requirement{}

// Requirement struct for Requirement
type Requirement struct {
	Gift NullableRealmGift `json:"gift,omitempty"`
	ReceivedCount NullableInt32 `json:"received_count,omitempty"`
	RequiresAmount NullableInt32 `json:"requires_amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Requirement Requirement

// NewRequirement instantiates a new Requirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequirement() *Requirement {
	this := Requirement{}
	return &this
}

// NewRequirementWithDefaults instantiates a new Requirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequirementWithDefaults() *Requirement {
	this := Requirement{}
	return &this
}

// GetGift returns the Gift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requirement) GetGift() RealmGift {
	if o == nil || IsNil(o.Gift.Get()) {
		var ret RealmGift
		return ret
	}
	return *o.Gift.Get()
}

// GetGiftOk returns a tuple with the Gift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requirement) GetGiftOk() (*RealmGift, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gift.Get(), o.Gift.IsSet()
}

// HasGift returns a boolean if a field has been set.
func (o *Requirement) HasGift() bool {
	if o != nil && o.Gift.IsSet() {
		return true
	}

	return false
}

// SetGift gets a reference to the given NullableRealmGift and assigns it to the Gift field.
func (o *Requirement) SetGift(v RealmGift) {
	o.Gift.Set(&v)
}
// SetGiftNil sets the value for Gift to be an explicit nil
func (o *Requirement) SetGiftNil() {
	o.Gift.Set(nil)
}

// UnsetGift ensures that no value is present for Gift, not even an explicit nil
func (o *Requirement) UnsetGift() {
	o.Gift.Unset()
}

// GetReceivedCount returns the ReceivedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requirement) GetReceivedCount() int32 {
	if o == nil || IsNil(o.ReceivedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReceivedCount.Get()
}

// GetReceivedCountOk returns a tuple with the ReceivedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requirement) GetReceivedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivedCount.Get(), o.ReceivedCount.IsSet()
}

// HasReceivedCount returns a boolean if a field has been set.
func (o *Requirement) HasReceivedCount() bool {
	if o != nil && o.ReceivedCount.IsSet() {
		return true
	}

	return false
}

// SetReceivedCount gets a reference to the given NullableInt32 and assigns it to the ReceivedCount field.
func (o *Requirement) SetReceivedCount(v int32) {
	o.ReceivedCount.Set(&v)
}
// SetReceivedCountNil sets the value for ReceivedCount to be an explicit nil
func (o *Requirement) SetReceivedCountNil() {
	o.ReceivedCount.Set(nil)
}

// UnsetReceivedCount ensures that no value is present for ReceivedCount, not even an explicit nil
func (o *Requirement) UnsetReceivedCount() {
	o.ReceivedCount.Unset()
}

// GetRequiresAmount returns the RequiresAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Requirement) GetRequiresAmount() int32 {
	if o == nil || IsNil(o.RequiresAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.RequiresAmount.Get()
}

// GetRequiresAmountOk returns a tuple with the RequiresAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Requirement) GetRequiresAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiresAmount.Get(), o.RequiresAmount.IsSet()
}

// HasRequiresAmount returns a boolean if a field has been set.
func (o *Requirement) HasRequiresAmount() bool {
	if o != nil && o.RequiresAmount.IsSet() {
		return true
	}

	return false
}

// SetRequiresAmount gets a reference to the given NullableInt32 and assigns it to the RequiresAmount field.
func (o *Requirement) SetRequiresAmount(v int32) {
	o.RequiresAmount.Set(&v)
}
// SetRequiresAmountNil sets the value for RequiresAmount to be an explicit nil
func (o *Requirement) SetRequiresAmountNil() {
	o.RequiresAmount.Set(nil)
}

// UnsetRequiresAmount ensures that no value is present for RequiresAmount, not even an explicit nil
func (o *Requirement) UnsetRequiresAmount() {
	o.RequiresAmount.Unset()
}

func (o Requirement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Requirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gift.IsSet() {
		toSerialize["gift"] = o.Gift.Get()
	}
	if o.ReceivedCount.IsSet() {
		toSerialize["received_count"] = o.ReceivedCount.Get()
	}
	if o.RequiresAmount.IsSet() {
		toSerialize["requires_amount"] = o.RequiresAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Requirement) UnmarshalJSON(data []byte) (err error) {
	varRequirement := _Requirement{}

	err = json.Unmarshal(data, &varRequirement)

	if err != nil {
		return err
	}

	*o = Requirement(varRequirement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift")
		delete(additionalProperties, "received_count")
		delete(additionalProperties, "requires_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequirement struct {
	value *Requirement
	isSet bool
}

func (v NullableRequirement) Get() *Requirement {
	return v.value
}

func (v *NullableRequirement) Set(val *Requirement) {
	v.value = val
	v.isSet = true
}

func (v NullableRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequirement(val *Requirement) *NullableRequirement {
	return &NullableRequirement{value: val, isSet: true}
}

func (v NullableRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


