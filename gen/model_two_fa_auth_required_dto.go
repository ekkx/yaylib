
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TwoFaAuthRequiredDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoFaAuthRequiredDTO{}

// TwoFaAuthRequiredDTO struct for TwoFaAuthRequiredDTO
type TwoFaAuthRequiredDTO struct {
	Login NullableBool `json:"login,omitempty"`
	Wallet NullableBool `json:"wallet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoFaAuthRequiredDTO TwoFaAuthRequiredDTO

// NewTwoFaAuthRequiredDTO instantiates a new TwoFaAuthRequiredDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoFaAuthRequiredDTO() *TwoFaAuthRequiredDTO {
	this := TwoFaAuthRequiredDTO{}
	return &this
}

// NewTwoFaAuthRequiredDTOWithDefaults instantiates a new TwoFaAuthRequiredDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoFaAuthRequiredDTOWithDefaults() *TwoFaAuthRequiredDTO {
	this := TwoFaAuthRequiredDTO{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoFaAuthRequiredDTO) GetLogin() bool {
	if o == nil || IsNil(o.Login.Get()) {
		var ret bool
		return ret
	}
	return *o.Login.Get()
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoFaAuthRequiredDTO) GetLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Login.Get(), o.Login.IsSet()
}

// HasLogin returns a boolean if a field has been set.
func (o *TwoFaAuthRequiredDTO) HasLogin() bool {
	if o != nil && o.Login.IsSet() {
		return true
	}

	return false
}

// SetLogin gets a reference to the given NullableBool and assigns it to the Login field.
func (o *TwoFaAuthRequiredDTO) SetLogin(v bool) {
	o.Login.Set(&v)
}
// SetLoginNil sets the value for Login to be an explicit nil
func (o *TwoFaAuthRequiredDTO) SetLoginNil() {
	o.Login.Set(nil)
}

// UnsetLogin ensures that no value is present for Login, not even an explicit nil
func (o *TwoFaAuthRequiredDTO) UnsetLogin() {
	o.Login.Unset()
}

// GetWallet returns the Wallet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoFaAuthRequiredDTO) GetWallet() bool {
	if o == nil || IsNil(o.Wallet.Get()) {
		var ret bool
		return ret
	}
	return *o.Wallet.Get()
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoFaAuthRequiredDTO) GetWalletOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Wallet.Get(), o.Wallet.IsSet()
}

// HasWallet returns a boolean if a field has been set.
func (o *TwoFaAuthRequiredDTO) HasWallet() bool {
	if o != nil && o.Wallet.IsSet() {
		return true
	}

	return false
}

// SetWallet gets a reference to the given NullableBool and assigns it to the Wallet field.
func (o *TwoFaAuthRequiredDTO) SetWallet(v bool) {
	o.Wallet.Set(&v)
}
// SetWalletNil sets the value for Wallet to be an explicit nil
func (o *TwoFaAuthRequiredDTO) SetWalletNil() {
	o.Wallet.Set(nil)
}

// UnsetWallet ensures that no value is present for Wallet, not even an explicit nil
func (o *TwoFaAuthRequiredDTO) UnsetWallet() {
	o.Wallet.Unset()
}

func (o TwoFaAuthRequiredDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoFaAuthRequiredDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Login.IsSet() {
		toSerialize["login"] = o.Login.Get()
	}
	if o.Wallet.IsSet() {
		toSerialize["wallet"] = o.Wallet.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoFaAuthRequiredDTO) UnmarshalJSON(data []byte) (err error) {
	varTwoFaAuthRequiredDTO := _TwoFaAuthRequiredDTO{}

	err = json.Unmarshal(data, &varTwoFaAuthRequiredDTO)

	if err != nil {
		return err
	}

	*o = TwoFaAuthRequiredDTO(varTwoFaAuthRequiredDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "login")
		delete(additionalProperties, "wallet")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTwoFaAuthRequiredDTO struct {
	value *TwoFaAuthRequiredDTO
	isSet bool
}

func (v NullableTwoFaAuthRequiredDTO) Get() *TwoFaAuthRequiredDTO {
	return v.value
}

func (v *NullableTwoFaAuthRequiredDTO) Set(val *TwoFaAuthRequiredDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoFaAuthRequiredDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoFaAuthRequiredDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoFaAuthRequiredDTO(val *TwoFaAuthRequiredDTO) *NullableTwoFaAuthRequiredDTO {
	return &NullableTwoFaAuthRequiredDTO{value: val, isSet: true}
}

func (v NullableTwoFaAuthRequiredDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoFaAuthRequiredDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


