
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FeaturesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturesDTO{}

// FeaturesDTO struct for FeaturesDTO
type FeaturesDTO struct {
	Pal NullableFeaturesPalDTO `json:"pal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeaturesDTO FeaturesDTO

// NewFeaturesDTO instantiates a new FeaturesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesDTO() *FeaturesDTO {
	this := FeaturesDTO{}
	return &this
}

// NewFeaturesDTOWithDefaults instantiates a new FeaturesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesDTOWithDefaults() *FeaturesDTO {
	this := FeaturesDTO{}
	return &this
}

// GetPal returns the Pal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeaturesDTO) GetPal() FeaturesPalDTO {
	if o == nil || IsNil(o.Pal.Get()) {
		var ret FeaturesPalDTO
		return ret
	}
	return *o.Pal.Get()
}

// GetPalOk returns a tuple with the Pal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeaturesDTO) GetPalOk() (*FeaturesPalDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pal.Get(), o.Pal.IsSet()
}

// HasPal returns a boolean if a field has been set.
func (o *FeaturesDTO) HasPal() bool {
	if o != nil && o.Pal.IsSet() {
		return true
	}

	return false
}

// SetPal gets a reference to the given NullableFeaturesPalDTO and assigns it to the Pal field.
func (o *FeaturesDTO) SetPal(v FeaturesPalDTO) {
	o.Pal.Set(&v)
}
// SetPalNil sets the value for Pal to be an explicit nil
func (o *FeaturesDTO) SetPalNil() {
	o.Pal.Set(nil)
}

// UnsetPal ensures that no value is present for Pal, not even an explicit nil
func (o *FeaturesDTO) UnsetPal() {
	o.Pal.Unset()
}

func (o FeaturesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeaturesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Pal.IsSet() {
		toSerialize["pal"] = o.Pal.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeaturesDTO) UnmarshalJSON(data []byte) (err error) {
	varFeaturesDTO := _FeaturesDTO{}

	err = json.Unmarshal(data, &varFeaturesDTO)

	if err != nil {
		return err
	}

	*o = FeaturesDTO(varFeaturesDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pal")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeaturesDTO struct {
	value *FeaturesDTO
	isSet bool
}

func (v NullableFeaturesDTO) Get() *FeaturesDTO {
	return v.value
}

func (v *NullableFeaturesDTO) Set(val *FeaturesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesDTO(val *FeaturesDTO) *NullableFeaturesDTO {
	return &NullableFeaturesDTO{value: val, isSet: true}
}

func (v NullableFeaturesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


