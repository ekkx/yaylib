
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PriceDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceDTO{}

// PriceDTO struct for PriceDTO
type PriceDTO struct {
	Jpy NullableString `json:"jpy,omitempty"`
	Usd NullableString `json:"usd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PriceDTO PriceDTO

// NewPriceDTO instantiates a new PriceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceDTO() *PriceDTO {
	this := PriceDTO{}
	return &this
}

// NewPriceDTOWithDefaults instantiates a new PriceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceDTOWithDefaults() *PriceDTO {
	this := PriceDTO{}
	return &this
}

// GetJpy returns the Jpy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceDTO) GetJpy() string {
	if o == nil || IsNil(o.Jpy.Get()) {
		var ret string
		return ret
	}
	return *o.Jpy.Get()
}

// GetJpyOk returns a tuple with the Jpy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceDTO) GetJpyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jpy.Get(), o.Jpy.IsSet()
}

// HasJpy returns a boolean if a field has been set.
func (o *PriceDTO) HasJpy() bool {
	if o != nil && o.Jpy.IsSet() {
		return true
	}

	return false
}

// SetJpy gets a reference to the given NullableString and assigns it to the Jpy field.
func (o *PriceDTO) SetJpy(v string) {
	o.Jpy.Set(&v)
}
// SetJpyNil sets the value for Jpy to be an explicit nil
func (o *PriceDTO) SetJpyNil() {
	o.Jpy.Set(nil)
}

// UnsetJpy ensures that no value is present for Jpy, not even an explicit nil
func (o *PriceDTO) UnsetJpy() {
	o.Jpy.Unset()
}

// GetUsd returns the Usd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceDTO) GetUsd() string {
	if o == nil || IsNil(o.Usd.Get()) {
		var ret string
		return ret
	}
	return *o.Usd.Get()
}

// GetUsdOk returns a tuple with the Usd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceDTO) GetUsdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usd.Get(), o.Usd.IsSet()
}

// HasUsd returns a boolean if a field has been set.
func (o *PriceDTO) HasUsd() bool {
	if o != nil && o.Usd.IsSet() {
		return true
	}

	return false
}

// SetUsd gets a reference to the given NullableString and assigns it to the Usd field.
func (o *PriceDTO) SetUsd(v string) {
	o.Usd.Set(&v)
}
// SetUsdNil sets the value for Usd to be an explicit nil
func (o *PriceDTO) SetUsdNil() {
	o.Usd.Set(nil)
}

// UnsetUsd ensures that no value is present for Usd, not even an explicit nil
func (o *PriceDTO) UnsetUsd() {
	o.Usd.Unset()
}

func (o PriceDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Jpy.IsSet() {
		toSerialize["jpy"] = o.Jpy.Get()
	}
	if o.Usd.IsSet() {
		toSerialize["usd"] = o.Usd.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PriceDTO) UnmarshalJSON(data []byte) (err error) {
	varPriceDTO := _PriceDTO{}

	err = json.Unmarshal(data, &varPriceDTO)

	if err != nil {
		return err
	}

	*o = PriceDTO(varPriceDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jpy")
		delete(additionalProperties, "usd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePriceDTO struct {
	value *PriceDTO
	isSet bool
}

func (v NullablePriceDTO) Get() *PriceDTO {
	return v.value
}

func (v *NullablePriceDTO) Set(val *PriceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceDTO(val *PriceDTO) *NullablePriceDTO {
	return &NullablePriceDTO{value: val, isSet: true}
}

func (v NullablePriceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


