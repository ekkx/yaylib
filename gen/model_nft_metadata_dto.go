
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the NFTMetadataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NFTMetadataDTO{}

// NFTMetadataDTO struct for NFTMetadataDTO
type NFTMetadataDTO struct {
	Attributes []NFTAttributeDTO `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NFTMetadataDTO NFTMetadataDTO

// NewNFTMetadataDTO instantiates a new NFTMetadataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFTMetadataDTO() *NFTMetadataDTO {
	this := NFTMetadataDTO{}
	return &this
}

// NewNFTMetadataDTOWithDefaults instantiates a new NFTMetadataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFTMetadataDTOWithDefaults() *NFTMetadataDTO {
	this := NFTMetadataDTO{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NFTMetadataDTO) GetAttributes() []NFTAttributeDTO {
	if o == nil {
		var ret []NFTAttributeDTO
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NFTMetadataDTO) GetAttributesOk() ([]NFTAttributeDTO, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *NFTMetadataDTO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []NFTAttributeDTO and assigns it to the Attributes field.
func (o *NFTMetadataDTO) SetAttributes(v []NFTAttributeDTO) {
	o.Attributes = v
}

func (o NFTMetadataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NFTMetadataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NFTMetadataDTO) UnmarshalJSON(data []byte) (err error) {
	varNFTMetadataDTO := _NFTMetadataDTO{}

	err = json.Unmarshal(data, &varNFTMetadataDTO)

	if err != nil {
		return err
	}

	*o = NFTMetadataDTO(varNFTMetadataDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNFTMetadataDTO struct {
	value *NFTMetadataDTO
	isSet bool
}

func (v NullableNFTMetadataDTO) Get() *NFTMetadataDTO {
	return v.value
}

func (v *NullableNFTMetadataDTO) Set(val *NFTMetadataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableNFTMetadataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableNFTMetadataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFTMetadataDTO(val *NFTMetadataDTO) *NullableNFTMetadataDTO {
	return &NullableNFTMetadataDTO{value: val, isSet: true}
}

func (v NullableNFTMetadataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFTMetadataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


