
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelWeb3WalletPreSign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelWeb3WalletPreSign{}

// ModelWeb3WalletPreSign struct for ModelWeb3WalletPreSign
type ModelWeb3WalletPreSign struct {
	AdditionGasPercent NullableModelWeb3WalletGasPercent `json:"addition_gas_percent,omitempty"`
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	ChainUrl NullableString `json:"chain_url,omitempty"`
	ContractAddress NullableString `json:"contract_address,omitempty"`
	InputsHex NullableString `json:"inputs_hex,omitempty"`
	Name NullableString `json:"name,omitempty"`
	TxUniqueId NullableString `json:"tx_unique_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelWeb3WalletPreSign ModelWeb3WalletPreSign

// NewModelWeb3WalletPreSign instantiates a new ModelWeb3WalletPreSign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelWeb3WalletPreSign() *ModelWeb3WalletPreSign {
	this := ModelWeb3WalletPreSign{}
	return &this
}

// NewModelWeb3WalletPreSignWithDefaults instantiates a new ModelWeb3WalletPreSign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelWeb3WalletPreSignWithDefaults() *ModelWeb3WalletPreSign {
	this := ModelWeb3WalletPreSign{}
	return &this
}

// GetAdditionGasPercent returns the AdditionGasPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletPreSign) GetAdditionGasPercent() ModelWeb3WalletGasPercent {
	if o == nil || IsNil(o.AdditionGasPercent.Get()) {
		var ret ModelWeb3WalletGasPercent
		return ret
	}
	return *o.AdditionGasPercent.Get()
}

// GetAdditionGasPercentOk returns a tuple with the AdditionGasPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletPreSign) GetAdditionGasPercentOk() (*ModelWeb3WalletGasPercent, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdditionGasPercent.Get(), o.AdditionGasPercent.IsSet()
}

// HasAdditionGasPercent returns a boolean if a field has been set.
func (o *ModelWeb3WalletPreSign) HasAdditionGasPercent() bool {
	if o != nil && o.AdditionGasPercent.IsSet() {
		return true
	}

	return false
}

// SetAdditionGasPercent gets a reference to the given NullableModelWeb3WalletGasPercent and assigns it to the AdditionGasPercent field.
func (o *ModelWeb3WalletPreSign) SetAdditionGasPercent(v ModelWeb3WalletGasPercent) {
	o.AdditionGasPercent.Set(&v)
}
// SetAdditionGasPercentNil sets the value for AdditionGasPercent to be an explicit nil
func (o *ModelWeb3WalletPreSign) SetAdditionGasPercentNil() {
	o.AdditionGasPercent.Set(nil)
}

// UnsetAdditionGasPercent ensures that no value is present for AdditionGasPercent, not even an explicit nil
func (o *ModelWeb3WalletPreSign) UnsetAdditionGasPercent() {
	o.AdditionGasPercent.Unset()
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletPreSign) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletPreSign) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *ModelWeb3WalletPreSign) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *ModelWeb3WalletPreSign) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *ModelWeb3WalletPreSign) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *ModelWeb3WalletPreSign) UnsetChainId() {
	o.ChainId.Unset()
}

// GetChainUrl returns the ChainUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletPreSign) GetChainUrl() string {
	if o == nil || IsNil(o.ChainUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ChainUrl.Get()
}

// GetChainUrlOk returns a tuple with the ChainUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletPreSign) GetChainUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainUrl.Get(), o.ChainUrl.IsSet()
}

// HasChainUrl returns a boolean if a field has been set.
func (o *ModelWeb3WalletPreSign) HasChainUrl() bool {
	if o != nil && o.ChainUrl.IsSet() {
		return true
	}

	return false
}

// SetChainUrl gets a reference to the given NullableString and assigns it to the ChainUrl field.
func (o *ModelWeb3WalletPreSign) SetChainUrl(v string) {
	o.ChainUrl.Set(&v)
}
// SetChainUrlNil sets the value for ChainUrl to be an explicit nil
func (o *ModelWeb3WalletPreSign) SetChainUrlNil() {
	o.ChainUrl.Set(nil)
}

// UnsetChainUrl ensures that no value is present for ChainUrl, not even an explicit nil
func (o *ModelWeb3WalletPreSign) UnsetChainUrl() {
	o.ChainUrl.Unset()
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletPreSign) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletPreSign) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *ModelWeb3WalletPreSign) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *ModelWeb3WalletPreSign) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *ModelWeb3WalletPreSign) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *ModelWeb3WalletPreSign) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetInputsHex returns the InputsHex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletPreSign) GetInputsHex() string {
	if o == nil || IsNil(o.InputsHex.Get()) {
		var ret string
		return ret
	}
	return *o.InputsHex.Get()
}

// GetInputsHexOk returns a tuple with the InputsHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletPreSign) GetInputsHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputsHex.Get(), o.InputsHex.IsSet()
}

// HasInputsHex returns a boolean if a field has been set.
func (o *ModelWeb3WalletPreSign) HasInputsHex() bool {
	if o != nil && o.InputsHex.IsSet() {
		return true
	}

	return false
}

// SetInputsHex gets a reference to the given NullableString and assigns it to the InputsHex field.
func (o *ModelWeb3WalletPreSign) SetInputsHex(v string) {
	o.InputsHex.Set(&v)
}
// SetInputsHexNil sets the value for InputsHex to be an explicit nil
func (o *ModelWeb3WalletPreSign) SetInputsHexNil() {
	o.InputsHex.Set(nil)
}

// UnsetInputsHex ensures that no value is present for InputsHex, not even an explicit nil
func (o *ModelWeb3WalletPreSign) UnsetInputsHex() {
	o.InputsHex.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletPreSign) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletPreSign) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelWeb3WalletPreSign) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelWeb3WalletPreSign) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelWeb3WalletPreSign) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelWeb3WalletPreSign) UnsetName() {
	o.Name.Unset()
}

// GetTxUniqueId returns the TxUniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelWeb3WalletPreSign) GetTxUniqueId() string {
	if o == nil || IsNil(o.TxUniqueId.Get()) {
		var ret string
		return ret
	}
	return *o.TxUniqueId.Get()
}

// GetTxUniqueIdOk returns a tuple with the TxUniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelWeb3WalletPreSign) GetTxUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxUniqueId.Get(), o.TxUniqueId.IsSet()
}

// HasTxUniqueId returns a boolean if a field has been set.
func (o *ModelWeb3WalletPreSign) HasTxUniqueId() bool {
	if o != nil && o.TxUniqueId.IsSet() {
		return true
	}

	return false
}

// SetTxUniqueId gets a reference to the given NullableString and assigns it to the TxUniqueId field.
func (o *ModelWeb3WalletPreSign) SetTxUniqueId(v string) {
	o.TxUniqueId.Set(&v)
}
// SetTxUniqueIdNil sets the value for TxUniqueId to be an explicit nil
func (o *ModelWeb3WalletPreSign) SetTxUniqueIdNil() {
	o.TxUniqueId.Set(nil)
}

// UnsetTxUniqueId ensures that no value is present for TxUniqueId, not even an explicit nil
func (o *ModelWeb3WalletPreSign) UnsetTxUniqueId() {
	o.TxUniqueId.Unset()
}

func (o ModelWeb3WalletPreSign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelWeb3WalletPreSign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionGasPercent.IsSet() {
		toSerialize["addition_gas_percent"] = o.AdditionGasPercent.Get()
	}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.ChainUrl.IsSet() {
		toSerialize["chain_url"] = o.ChainUrl.Get()
	}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.InputsHex.IsSet() {
		toSerialize["inputs_hex"] = o.InputsHex.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.TxUniqueId.IsSet() {
		toSerialize["tx_unique_id"] = o.TxUniqueId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelWeb3WalletPreSign) UnmarshalJSON(data []byte) (err error) {
	varModelWeb3WalletPreSign := _ModelWeb3WalletPreSign{}

	err = json.Unmarshal(data, &varModelWeb3WalletPreSign)

	if err != nil {
		return err
	}

	*o = ModelWeb3WalletPreSign(varModelWeb3WalletPreSign)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addition_gas_percent")
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "chain_url")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "inputs_hex")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tx_unique_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelWeb3WalletPreSign struct {
	value *ModelWeb3WalletPreSign
	isSet bool
}

func (v NullableModelWeb3WalletPreSign) Get() *ModelWeb3WalletPreSign {
	return v.value
}

func (v *NullableModelWeb3WalletPreSign) Set(val *ModelWeb3WalletPreSign) {
	v.value = val
	v.isSet = true
}

func (v NullableModelWeb3WalletPreSign) IsSet() bool {
	return v.isSet
}

func (v *NullableModelWeb3WalletPreSign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelWeb3WalletPreSign(val *ModelWeb3WalletPreSign) *NullableModelWeb3WalletPreSign {
	return &NullableModelWeb3WalletPreSign{value: val, isSet: true}
}

func (v NullableModelWeb3WalletPreSign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelWeb3WalletPreSign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


