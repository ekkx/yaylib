
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TXNLogData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TXNLogData{}

// TXNLogData struct for TXNLogData
type TXNLogData struct {
	ContractAddress NullableString `json:"contract_address,omitempty"`
	InputsHex NullableString `json:"inputs_hex,omitempty"`
	TxUniqueId NullableString `json:"tx_unique_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TXNLogData TXNLogData

// NewTXNLogData instantiates a new TXNLogData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTXNLogData() *TXNLogData {
	this := TXNLogData{}
	return &this
}

// NewTXNLogDataWithDefaults instantiates a new TXNLogData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTXNLogDataWithDefaults() *TXNLogData {
	this := TXNLogData{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TXNLogData) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TXNLogData) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *TXNLogData) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *TXNLogData) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *TXNLogData) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *TXNLogData) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetInputsHex returns the InputsHex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TXNLogData) GetInputsHex() string {
	if o == nil || IsNil(o.InputsHex.Get()) {
		var ret string
		return ret
	}
	return *o.InputsHex.Get()
}

// GetInputsHexOk returns a tuple with the InputsHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TXNLogData) GetInputsHexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputsHex.Get(), o.InputsHex.IsSet()
}

// HasInputsHex returns a boolean if a field has been set.
func (o *TXNLogData) HasInputsHex() bool {
	if o != nil && o.InputsHex.IsSet() {
		return true
	}

	return false
}

// SetInputsHex gets a reference to the given NullableString and assigns it to the InputsHex field.
func (o *TXNLogData) SetInputsHex(v string) {
	o.InputsHex.Set(&v)
}
// SetInputsHexNil sets the value for InputsHex to be an explicit nil
func (o *TXNLogData) SetInputsHexNil() {
	o.InputsHex.Set(nil)
}

// UnsetInputsHex ensures that no value is present for InputsHex, not even an explicit nil
func (o *TXNLogData) UnsetInputsHex() {
	o.InputsHex.Unset()
}

// GetTxUniqueId returns the TxUniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TXNLogData) GetTxUniqueId() string {
	if o == nil || IsNil(o.TxUniqueId.Get()) {
		var ret string
		return ret
	}
	return *o.TxUniqueId.Get()
}

// GetTxUniqueIdOk returns a tuple with the TxUniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TXNLogData) GetTxUniqueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxUniqueId.Get(), o.TxUniqueId.IsSet()
}

// HasTxUniqueId returns a boolean if a field has been set.
func (o *TXNLogData) HasTxUniqueId() bool {
	if o != nil && o.TxUniqueId.IsSet() {
		return true
	}

	return false
}

// SetTxUniqueId gets a reference to the given NullableString and assigns it to the TxUniqueId field.
func (o *TXNLogData) SetTxUniqueId(v string) {
	o.TxUniqueId.Set(&v)
}
// SetTxUniqueIdNil sets the value for TxUniqueId to be an explicit nil
func (o *TXNLogData) SetTxUniqueIdNil() {
	o.TxUniqueId.Set(nil)
}

// UnsetTxUniqueId ensures that no value is present for TxUniqueId, not even an explicit nil
func (o *TXNLogData) UnsetTxUniqueId() {
	o.TxUniqueId.Unset()
}

func (o TXNLogData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TXNLogData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.InputsHex.IsSet() {
		toSerialize["inputs_hex"] = o.InputsHex.Get()
	}
	if o.TxUniqueId.IsSet() {
		toSerialize["tx_unique_id"] = o.TxUniqueId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TXNLogData) UnmarshalJSON(data []byte) (err error) {
	varTXNLogData := _TXNLogData{}

	err = json.Unmarshal(data, &varTXNLogData)

	if err != nil {
		return err
	}

	*o = TXNLogData(varTXNLogData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "inputs_hex")
		delete(additionalProperties, "tx_unique_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTXNLogData struct {
	value *TXNLogData
	isSet bool
}

func (v NullableTXNLogData) Get() *TXNLogData {
	return v.value
}

func (v *NullableTXNLogData) Set(val *TXNLogData) {
	v.value = val
	v.isSet = true
}

func (v NullableTXNLogData) IsSet() bool {
	return v.isSet
}

func (v *NullableTXNLogData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTXNLogData(val *TXNLogData) *NullableTXNLogData {
	return &NullableTXNLogData{value: val, isSet: true}
}

func (v NullableTXNLogData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTXNLogData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


