
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FeaturesPalDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturesPalDTO{}

// FeaturesPalDTO struct for FeaturesPalDTO
type FeaturesPalDTO struct {
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	DepositwithdrawAddress NullableString `json:"depositwithdraw_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeaturesPalDTO FeaturesPalDTO

// NewFeaturesPalDTO instantiates a new FeaturesPalDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesPalDTO() *FeaturesPalDTO {
	this := FeaturesPalDTO{}
	return &this
}

// NewFeaturesPalDTOWithDefaults instantiates a new FeaturesPalDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesPalDTOWithDefaults() *FeaturesPalDTO {
	this := FeaturesPalDTO{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeaturesPalDTO) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeaturesPalDTO) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *FeaturesPalDTO) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *FeaturesPalDTO) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *FeaturesPalDTO) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *FeaturesPalDTO) UnsetChainId() {
	o.ChainId.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeaturesPalDTO) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeaturesPalDTO) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *FeaturesPalDTO) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *FeaturesPalDTO) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *FeaturesPalDTO) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *FeaturesPalDTO) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetDepositwithdrawAddress returns the DepositwithdrawAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeaturesPalDTO) GetDepositwithdrawAddress() string {
	if o == nil || IsNil(o.DepositwithdrawAddress.Get()) {
		var ret string
		return ret
	}
	return *o.DepositwithdrawAddress.Get()
}

// GetDepositwithdrawAddressOk returns a tuple with the DepositwithdrawAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeaturesPalDTO) GetDepositwithdrawAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepositwithdrawAddress.Get(), o.DepositwithdrawAddress.IsSet()
}

// HasDepositwithdrawAddress returns a boolean if a field has been set.
func (o *FeaturesPalDTO) HasDepositwithdrawAddress() bool {
	if o != nil && o.DepositwithdrawAddress.IsSet() {
		return true
	}

	return false
}

// SetDepositwithdrawAddress gets a reference to the given NullableString and assigns it to the DepositwithdrawAddress field.
func (o *FeaturesPalDTO) SetDepositwithdrawAddress(v string) {
	o.DepositwithdrawAddress.Set(&v)
}
// SetDepositwithdrawAddressNil sets the value for DepositwithdrawAddress to be an explicit nil
func (o *FeaturesPalDTO) SetDepositwithdrawAddressNil() {
	o.DepositwithdrawAddress.Set(nil)
}

// UnsetDepositwithdrawAddress ensures that no value is present for DepositwithdrawAddress, not even an explicit nil
func (o *FeaturesPalDTO) UnsetDepositwithdrawAddress() {
	o.DepositwithdrawAddress.Unset()
}

func (o FeaturesPalDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeaturesPalDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.DepositwithdrawAddress.IsSet() {
		toSerialize["depositwithdraw_address"] = o.DepositwithdrawAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeaturesPalDTO) UnmarshalJSON(data []byte) (err error) {
	varFeaturesPalDTO := _FeaturesPalDTO{}

	err = json.Unmarshal(data, &varFeaturesPalDTO)

	if err != nil {
		return err
	}

	*o = FeaturesPalDTO(varFeaturesPalDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "depositwithdraw_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeaturesPalDTO struct {
	value *FeaturesPalDTO
	isSet bool
}

func (v NullableFeaturesPalDTO) Get() *FeaturesPalDTO {
	return v.value
}

func (v *NullableFeaturesPalDTO) Set(val *FeaturesPalDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesPalDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesPalDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesPalDTO(val *FeaturesPalDTO) *NullableFeaturesPalDTO {
	return &NullableFeaturesPalDTO{value: val, isSet: true}
}

func (v NullableFeaturesPalDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesPalDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


