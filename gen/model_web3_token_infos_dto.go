
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3TokenInfosDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3TokenInfosDTO{}

// Web3TokenInfosDTO struct for Web3TokenInfosDTO
type Web3TokenInfosDTO struct {
	TokenInfos []Web3TokenInfoDTO `json:"token_infos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3TokenInfosDTO Web3TokenInfosDTO

// NewWeb3TokenInfosDTO instantiates a new Web3TokenInfosDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3TokenInfosDTO() *Web3TokenInfosDTO {
	this := Web3TokenInfosDTO{}
	return &this
}

// NewWeb3TokenInfosDTOWithDefaults instantiates a new Web3TokenInfosDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3TokenInfosDTOWithDefaults() *Web3TokenInfosDTO {
	this := Web3TokenInfosDTO{}
	return &this
}

// GetTokenInfos returns the TokenInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3TokenInfosDTO) GetTokenInfos() []Web3TokenInfoDTO {
	if o == nil {
		var ret []Web3TokenInfoDTO
		return ret
	}
	return o.TokenInfos
}

// GetTokenInfosOk returns a tuple with the TokenInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3TokenInfosDTO) GetTokenInfosOk() ([]Web3TokenInfoDTO, bool) {
	if o == nil || IsNil(o.TokenInfos) {
		return nil, false
	}
	return o.TokenInfos, true
}

// HasTokenInfos returns a boolean if a field has been set.
func (o *Web3TokenInfosDTO) HasTokenInfos() bool {
	if o != nil && !IsNil(o.TokenInfos) {
		return true
	}

	return false
}

// SetTokenInfos gets a reference to the given []Web3TokenInfoDTO and assigns it to the TokenInfos field.
func (o *Web3TokenInfosDTO) SetTokenInfos(v []Web3TokenInfoDTO) {
	o.TokenInfos = v
}

func (o Web3TokenInfosDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3TokenInfosDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TokenInfos != nil {
		toSerialize["token_infos"] = o.TokenInfos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3TokenInfosDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3TokenInfosDTO := _Web3TokenInfosDTO{}

	err = json.Unmarshal(data, &varWeb3TokenInfosDTO)

	if err != nil {
		return err
	}

	*o = Web3TokenInfosDTO(varWeb3TokenInfosDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token_infos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3TokenInfosDTO struct {
	value *Web3TokenInfosDTO
	isSet bool
}

func (v NullableWeb3TokenInfosDTO) Get() *Web3TokenInfosDTO {
	return v.value
}

func (v *NullableWeb3TokenInfosDTO) Set(val *Web3TokenInfosDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3TokenInfosDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3TokenInfosDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3TokenInfosDTO(val *Web3TokenInfosDTO) *NullableWeb3TokenInfosDTO {
	return &NullableWeb3TokenInfosDTO{value: val, isSet: true}
}

func (v NullableWeb3TokenInfosDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3TokenInfosDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


