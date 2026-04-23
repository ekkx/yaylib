
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the NFTAttributeDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NFTAttributeDTO{}

// NFTAttributeDTO struct for NFTAttributeDTO
type NFTAttributeDTO struct {
	TraitType NullableString `json:"trait_type,omitempty"`
	Value NullableString `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NFTAttributeDTO NFTAttributeDTO

// NewNFTAttributeDTO instantiates a new NFTAttributeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTAttributeDTO() *NFTAttributeDTO {
	this := NFTAttributeDTO{}
	return &this
}

// NewNFTAttributeDTOWithDefaults instantiates a new NFTAttributeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTAttributeDTOWithDefaults() *NFTAttributeDTO {
	this := NFTAttributeDTO{}
	return &this
}

// GetTraitType returns the TraitType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NFTAttributeDTO) GetTraitType() string {
	if o == nil || IsNil(o.TraitType.Get()) {
		var ret string
		return ret
	}
	return *o.TraitType.Get()
}

// GetTraitTypeOk returns a tuple with the TraitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NFTAttributeDTO) GetTraitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraitType.Get(), o.TraitType.IsSet()
}

// HasTraitType returns a boolean if a field has been set.
func (o *NFTAttributeDTO) HasTraitType() bool {
	if o != nil && o.TraitType.IsSet() {
		return true
	}

	return false
}

// SetTraitType gets a reference to the given NullableString and assigns it to the TraitType field.
func (o *NFTAttributeDTO) SetTraitType(v string) {
	o.TraitType.Set(&v)
}
// SetTraitTypeNil sets the value for TraitType to be an explicit nil
func (o *NFTAttributeDTO) SetTraitTypeNil() {
	o.TraitType.Set(nil)
}

// UnsetTraitType ensures that no value is present for TraitType, not even an explicit nil
func (o *NFTAttributeDTO) UnsetTraitType() {
	o.TraitType.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NFTAttributeDTO) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NFTAttributeDTO) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *NFTAttributeDTO) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *NFTAttributeDTO) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *NFTAttributeDTO) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *NFTAttributeDTO) UnsetValue() {
	o.Value.Unset()
}

func (o NFTAttributeDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NFTAttributeDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TraitType.IsSet() {
		toSerialize["trait_type"] = o.TraitType.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NFTAttributeDTO) UnmarshalJSON(data []byte) (err error) {
	varNFTAttributeDTO := _NFTAttributeDTO{}

	err = json.Unmarshal(data, &varNFTAttributeDTO)

	if err != nil {
		return err
	}

	*o = NFTAttributeDTO(varNFTAttributeDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "trait_type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNFTAttributeDTO struct {
	value *NFTAttributeDTO
	isSet bool
}

func (v NullableNFTAttributeDTO) Get() *NFTAttributeDTO {
	return v.value
}

func (v *NullableNFTAttributeDTO) Set(val *NFTAttributeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTAttributeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTAttributeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTAttributeDTO(val *NFTAttributeDTO) *NullableNFTAttributeDTO {
	return &NullableNFTAttributeDTO{value: val, isSet: true}
}

func (v NullableNFTAttributeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTAttributeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


