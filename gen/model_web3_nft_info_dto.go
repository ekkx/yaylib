
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3NftInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3NftInfoDTO{}

// Web3NftInfoDTO struct for Web3NftInfoDTO
type Web3NftInfoDTO struct {
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	LocalizedUrls NullableLocalizedStringDTO `json:"localized_urls,omitempty"`
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3NftInfoDTO Web3NftInfoDTO

// NewWeb3NftInfoDTO instantiates a new Web3NftInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3NftInfoDTO() *Web3NftInfoDTO {
	this := Web3NftInfoDTO{}
	return &this
}

// NewWeb3NftInfoDTOWithDefaults instantiates a new Web3NftInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3NftInfoDTOWithDefaults() *Web3NftInfoDTO {
	this := Web3NftInfoDTO{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftInfoDTO) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftInfoDTO) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Web3NftInfoDTO) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Web3NftInfoDTO) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Web3NftInfoDTO) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Web3NftInfoDTO) UnsetChainId() {
	o.ChainId.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftInfoDTO) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftInfoDTO) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *Web3NftInfoDTO) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *Web3NftInfoDTO) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *Web3NftInfoDTO) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *Web3NftInfoDTO) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetLocalizedUrls returns the LocalizedUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftInfoDTO) GetLocalizedUrls() LocalizedStringDTO {
	if o == nil || IsNil(o.LocalizedUrls.Get()) {
		var ret LocalizedStringDTO
		return ret
	}
	return *o.LocalizedUrls.Get()
}

// GetLocalizedUrlsOk returns a tuple with the LocalizedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftInfoDTO) GetLocalizedUrlsOk() (*LocalizedStringDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalizedUrls.Get(), o.LocalizedUrls.IsSet()
}

// HasLocalizedUrls returns a boolean if a field has been set.
func (o *Web3NftInfoDTO) HasLocalizedUrls() bool {
	if o != nil && o.LocalizedUrls.IsSet() {
		return true
	}

	return false
}

// SetLocalizedUrls gets a reference to the given NullableLocalizedStringDTO and assigns it to the LocalizedUrls field.
func (o *Web3NftInfoDTO) SetLocalizedUrls(v LocalizedStringDTO) {
	o.LocalizedUrls.Set(&v)
}
// SetLocalizedUrlsNil sets the value for LocalizedUrls to be an explicit nil
func (o *Web3NftInfoDTO) SetLocalizedUrlsNil() {
	o.LocalizedUrls.Set(nil)
}

// UnsetLocalizedUrls ensures that no value is present for LocalizedUrls, not even an explicit nil
func (o *Web3NftInfoDTO) UnsetLocalizedUrls() {
	o.LocalizedUrls.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3NftInfoDTO) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3NftInfoDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Web3NftInfoDTO) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Web3NftInfoDTO) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Web3NftInfoDTO) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Web3NftInfoDTO) UnsetName() {
	o.Name.Unset()
}

func (o Web3NftInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3NftInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.LocalizedUrls.IsSet() {
		toSerialize["localized_urls"] = o.LocalizedUrls.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3NftInfoDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3NftInfoDTO := _Web3NftInfoDTO{}

	err = json.Unmarshal(data, &varWeb3NftInfoDTO)

	if err != nil {
		return err
	}

	*o = Web3NftInfoDTO(varWeb3NftInfoDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "localized_urls")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3NftInfoDTO struct {
	value *Web3NftInfoDTO
	isSet bool
}

func (v NullableWeb3NftInfoDTO) Get() *Web3NftInfoDTO {
	return v.value
}

func (v *NullableWeb3NftInfoDTO) Set(val *Web3NftInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3NftInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3NftInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3NftInfoDTO(val *Web3NftInfoDTO) *NullableWeb3NftInfoDTO {
	return &NullableWeb3NftInfoDTO{value: val, isSet: true}
}

func (v NullableWeb3NftInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3NftInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


