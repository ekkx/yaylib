
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3NftInfosDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3NftInfosDTO{}

// Web3NftInfosDTO struct for Web3NftInfosDTO
type Web3NftInfosDTO struct {
	NftInfos []Web3NftInfoDTO `json:"nft_infos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3NftInfosDTO Web3NftInfosDTO

// NewWeb3NftInfosDTO instantiates a new Web3NftInfosDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3NftInfosDTO() *Web3NftInfosDTO {
	this := Web3NftInfosDTO{}
	return &this
}

// NewWeb3NftInfosDTOWithDefaults instantiates a new Web3NftInfosDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3NftInfosDTOWithDefaults() *Web3NftInfosDTO {
	this := Web3NftInfosDTO{}
	return &this
}

// GetNftInfos returns the NftInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftInfosDTO) GetNftInfos() []Web3NftInfoDTO {
	if o == nil {
		var ret []Web3NftInfoDTO
		return ret
	}
	return o.NftInfos
}

// GetNftInfosOk returns a tuple with the NftInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftInfosDTO) GetNftInfosOk() ([]Web3NftInfoDTO, bool) {
	if o == nil || IsNil(o.NftInfos) {
		return nil, false
	}
	return o.NftInfos, true
}

// HasNftInfos returns a boolean if a field has been set.
func (o *Web3NftInfosDTO) HasNftInfos() bool {
	if o != nil && !IsNil(o.NftInfos) {
		return true
	}

	return false
}

// SetNftInfos gets a reference to the given []Web3NftInfoDTO and assigns it to the NftInfos field.
func (o *Web3NftInfosDTO) SetNftInfos(v []Web3NftInfoDTO) {
	o.NftInfos = v
}

func (o Web3NftInfosDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3NftInfosDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NftInfos != nil {
		toSerialize["nft_infos"] = o.NftInfos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3NftInfosDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3NftInfosDTO := _Web3NftInfosDTO{}

	err = json.Unmarshal(data, &varWeb3NftInfosDTO)

	if err != nil {
		return err
	}

	*o = Web3NftInfosDTO(varWeb3NftInfosDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nft_infos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3NftInfosDTO struct {
	value *Web3NftInfosDTO
	isSet bool
}

func (v NullableWeb3NftInfosDTO) Get() *Web3NftInfosDTO {
	return v.value
}

func (v *NullableWeb3NftInfosDTO) Set(val *Web3NftInfosDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3NftInfosDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3NftInfosDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3NftInfosDTO(val *Web3NftInfosDTO) *NullableWeb3NftInfosDTO {
	return &NullableWeb3NftInfosDTO{value: val, isSet: true}
}

func (v NullableWeb3NftInfosDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3NftInfosDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


