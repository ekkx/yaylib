
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftCardOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftCardOption{}

// GiftCardOption struct for GiftCardOption
type GiftCardOption struct {
	Empl NullableFloat64 `json:"empl,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Value NullableFloat64 `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftCardOption GiftCardOption

// NewGiftCardOption instantiates a new GiftCardOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftCardOption() *GiftCardOption {
	this := GiftCardOption{}
	return &this
}

// NewGiftCardOptionWithDefaults instantiates a new GiftCardOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftCardOptionWithDefaults() *GiftCardOption {
	this := GiftCardOption{}
	return &this
}

// GetEmpl returns the Empl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardOption) GetEmpl() float64 {
	if o == nil || IsNil(o.Empl.Get()) {
		var ret float64
		return ret
	}
	return *o.Empl.Get()
}

// GetEmplOk returns a tuple with the Empl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardOption) GetEmplOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Empl.Get(), o.Empl.IsSet()
}

// HasEmpl returns a boolean if a field has been set.
func (o *GiftCardOption) HasEmpl() bool {
	if o != nil && o.Empl.IsSet() {
		return true
	}

	return false
}

// SetEmpl gets a reference to the given NullableFloat64 and assigns it to the Empl field.
func (o *GiftCardOption) SetEmpl(v float64) {
	o.Empl.Set(&v)
}
// SetEmplNil sets the value for Empl to be an explicit nil
func (o *GiftCardOption) SetEmplNil() {
	o.Empl.Set(nil)
}

// UnsetEmpl ensures that no value is present for Empl, not even an explicit nil
func (o *GiftCardOption) UnsetEmpl() {
	o.Empl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardOption) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardOption) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GiftCardOption) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *GiftCardOption) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GiftCardOption) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GiftCardOption) UnsetId() {
	o.Id.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftCardOption) GetValue() float64 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret float64
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftCardOption) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *GiftCardOption) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableFloat64 and assigns it to the Value field.
func (o *GiftCardOption) SetValue(v float64) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *GiftCardOption) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *GiftCardOption) UnsetValue() {
	o.Value.Unset()
}

func (o GiftCardOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftCardOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Empl.IsSet() {
		toSerialize["empl"] = o.Empl.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftCardOption) UnmarshalJSON(data []byte) (err error) {
	varGiftCardOption := _GiftCardOption{}

	err = json.Unmarshal(data, &varGiftCardOption)

	if err != nil {
		return err
	}

	*o = GiftCardOption(varGiftCardOption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "empl")
		delete(additionalProperties, "id")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftCardOption struct {
	value *GiftCardOption
	isSet bool
}

func (v NullableGiftCardOption) Get() *GiftCardOption {
	return v.value
}

func (v *NullableGiftCardOption) Set(val *GiftCardOption) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftCardOption) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftCardOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftCardOption(val *GiftCardOption) *NullableGiftCardOption {
	return &NullableGiftCardOption{value: val, isSet: true}
}

func (v NullableGiftCardOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftCardOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


