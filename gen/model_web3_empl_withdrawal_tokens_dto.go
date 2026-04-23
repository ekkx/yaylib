
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3EmplWithdrawalTokensDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3EmplWithdrawalTokensDTO{}

// Web3EmplWithdrawalTokensDTO struct for Web3EmplWithdrawalTokensDTO
type Web3EmplWithdrawalTokensDTO struct {
	Tokens []Web3EmplWithdrawalTokenDTO `json:"tokens,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3EmplWithdrawalTokensDTO Web3EmplWithdrawalTokensDTO

// NewWeb3EmplWithdrawalTokensDTO instantiates a new Web3EmplWithdrawalTokensDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3EmplWithdrawalTokensDTO() *Web3EmplWithdrawalTokensDTO {
	this := Web3EmplWithdrawalTokensDTO{}
	return &this
}

// NewWeb3EmplWithdrawalTokensDTOWithDefaults instantiates a new Web3EmplWithdrawalTokensDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3EmplWithdrawalTokensDTOWithDefaults() *Web3EmplWithdrawalTokensDTO {
	this := Web3EmplWithdrawalTokensDTO{}
	return &this
}

// GetTokens returns the Tokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokensDTO) GetTokens() []Web3EmplWithdrawalTokenDTO {
	if o == nil {
		var ret []Web3EmplWithdrawalTokenDTO
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokensDTO) GetTokensOk() ([]Web3EmplWithdrawalTokenDTO, bool) {
	if o == nil || IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokensDTO) HasTokens() bool {
	if o != nil && !IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []Web3EmplWithdrawalTokenDTO and assigns it to the Tokens field.
func (o *Web3EmplWithdrawalTokensDTO) SetTokens(v []Web3EmplWithdrawalTokenDTO) {
	o.Tokens = v
}

func (o Web3EmplWithdrawalTokensDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3EmplWithdrawalTokensDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3EmplWithdrawalTokensDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3EmplWithdrawalTokensDTO := _Web3EmplWithdrawalTokensDTO{}

	err = json.Unmarshal(data, &varWeb3EmplWithdrawalTokensDTO)

	if err != nil {
		return err
	}

	*o = Web3EmplWithdrawalTokensDTO(varWeb3EmplWithdrawalTokensDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tokens")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3EmplWithdrawalTokensDTO struct {
	value *Web3EmplWithdrawalTokensDTO
	isSet bool
}

func (v NullableWeb3EmplWithdrawalTokensDTO) Get() *Web3EmplWithdrawalTokensDTO {
	return v.value
}

func (v *NullableWeb3EmplWithdrawalTokensDTO) Set(val *Web3EmplWithdrawalTokensDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3EmplWithdrawalTokensDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3EmplWithdrawalTokensDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3EmplWithdrawalTokensDTO(val *Web3EmplWithdrawalTokensDTO) *NullableWeb3EmplWithdrawalTokensDTO {
	return &NullableWeb3EmplWithdrawalTokensDTO{value: val, isSet: true}
}

func (v NullableWeb3EmplWithdrawalTokensDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3EmplWithdrawalTokensDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


