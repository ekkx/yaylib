
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TXNData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TXNData{}

// TXNData struct for TXNData
type TXNData struct {
	ContractAddress NullableString `json:"contract_address,omitempty"`
	InputsHex NullableString `json:"inputs_hex,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TXNData TXNData

// NewTXNData instantiates a new TXNData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTXNData() *TXNData {
	this := TXNData{}
	return &this
}

// NewTXNDataWithDefaults instantiates a new TXNData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTXNDataWithDefaults() *TXNData {
	this := TXNData{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TXNData) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TXNData) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *TXNData) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *TXNData) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *TXNData) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *TXNData) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetInputsHex returns the InputsHex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TXNData) GetInputsHex() string {
	if o == nil || IsNil(o.InputsHex.Get()) {
		var ret string
		return ret
	}
	return *o.InputsHex.Get()
}

// GetInputsHexOk returns a tuple with the InputsHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TXNData) GetInputsHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputsHex.Get(), o.InputsHex.IsSet()
}

// HasInputsHex returns a boolean if a field has been set.
func (o *TXNData) HasInputsHex() bool {
	if o != nil && o.InputsHex.IsSet() {
		return true
	}

	return false
}

// SetInputsHex gets a reference to the given NullableString and assigns it to the InputsHex field.
func (o *TXNData) SetInputsHex(v string) {
	o.InputsHex.Set(&v)
}
// SetInputsHexNil sets the value for InputsHex to be an explicit nil
func (o *TXNData) SetInputsHexNil() {
	o.InputsHex.Set(nil)
}

// UnsetInputsHex ensures that no value is present for InputsHex, not even an explicit nil
func (o *TXNData) UnsetInputsHex() {
	o.InputsHex.Unset()
}

func (o TXNData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TXNData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.InputsHex.IsSet() {
		toSerialize["inputs_hex"] = o.InputsHex.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TXNData) UnmarshalJSON(data []byte) (err error) {
	varTXNData := _TXNData{}

	err = json.Unmarshal(data, &varTXNData)

	if err != nil {
		return err
	}

	*o = TXNData(varTXNData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "inputs_hex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTXNData struct {
	value *TXNData
	isSet bool
}

func (v NullableTXNData) Get() *TXNData {
	return v.value
}

func (v *NullableTXNData) Set(val *TXNData) {
	v.value = val
	v.isSet = true
}

func (v NullableTXNData) IsSet() bool {
	return v.isSet
}

func (v *NullableTXNData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTXNData(val *TXNData) *NullableTXNData {
	return &NullableTXNData{value: val, isSet: true}
}

func (v NullableTXNData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTXNData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


