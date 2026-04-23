
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3EmplTokenExchangeConversionAndFeeDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3EmplTokenExchangeConversionAndFeeDTO{}

// Web3EmplTokenExchangeConversionAndFeeDTO struct for Web3EmplTokenExchangeConversionAndFeeDTO
type Web3EmplTokenExchangeConversionAndFeeDTO struct {
	EmplAmount NullableString `json:"empl_amount,omitempty"`
	EmplTxFee NullableString `json:"empl_tx_fee,omitempty"`
	MaxAmountOut NullableString `json:"max_amount_out,omitempty"`
	MinAmountOut NullableString `json:"min_amount_out,omitempty"`
	TokenOutAddress NullableString `json:"token_out_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3EmplTokenExchangeConversionAndFeeDTO Web3EmplTokenExchangeConversionAndFeeDTO

// NewWeb3EmplTokenExchangeConversionAndFeeDTO instantiates a new Web3EmplTokenExchangeConversionAndFeeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3EmplTokenExchangeConversionAndFeeDTO() *Web3EmplTokenExchangeConversionAndFeeDTO {
	this := Web3EmplTokenExchangeConversionAndFeeDTO{}
	return &this
}

// NewWeb3EmplTokenExchangeConversionAndFeeDTOWithDefaults instantiates a new Web3EmplTokenExchangeConversionAndFeeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3EmplTokenExchangeConversionAndFeeDTOWithDefaults() *Web3EmplTokenExchangeConversionAndFeeDTO {
	this := Web3EmplTokenExchangeConversionAndFeeDTO{}
	return &this
}

// GetEmplAmount returns the EmplAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetEmplAmount() string {
	if o == nil || IsNil(o.EmplAmount.Get()) {
		var ret string
		return ret
	}
	return *o.EmplAmount.Get()
}

// GetEmplAmountOk returns a tuple with the EmplAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetEmplAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplAmount.Get(), o.EmplAmount.IsSet()
}

// HasEmplAmount returns a boolean if a field has been set.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) HasEmplAmount() bool {
	if o != nil && o.EmplAmount.IsSet() {
		return true
	}

	return false
}

// SetEmplAmount gets a reference to the given NullableString and assigns it to the EmplAmount field.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetEmplAmount(v string) {
	o.EmplAmount.Set(&v)
}
// SetEmplAmountNil sets the value for EmplAmount to be an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetEmplAmountNil() {
	o.EmplAmount.Set(nil)
}

// UnsetEmplAmount ensures that no value is present for EmplAmount, not even an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) UnsetEmplAmount() {
	o.EmplAmount.Unset()
}

// GetEmplTxFee returns the EmplTxFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetEmplTxFee() string {
	if o == nil || IsNil(o.EmplTxFee.Get()) {
		var ret string
		return ret
	}
	return *o.EmplTxFee.Get()
}

// GetEmplTxFeeOk returns a tuple with the EmplTxFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetEmplTxFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplTxFee.Get(), o.EmplTxFee.IsSet()
}

// HasEmplTxFee returns a boolean if a field has been set.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) HasEmplTxFee() bool {
	if o != nil && o.EmplTxFee.IsSet() {
		return true
	}

	return false
}

// SetEmplTxFee gets a reference to the given NullableString and assigns it to the EmplTxFee field.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetEmplTxFee(v string) {
	o.EmplTxFee.Set(&v)
}
// SetEmplTxFeeNil sets the value for EmplTxFee to be an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetEmplTxFeeNil() {
	o.EmplTxFee.Set(nil)
}

// UnsetEmplTxFee ensures that no value is present for EmplTxFee, not even an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) UnsetEmplTxFee() {
	o.EmplTxFee.Unset()
}

// GetMaxAmountOut returns the MaxAmountOut field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetMaxAmountOut() string {
	if o == nil || IsNil(o.MaxAmountOut.Get()) {
		var ret string
		return ret
	}
	return *o.MaxAmountOut.Get()
}

// GetMaxAmountOutOk returns a tuple with the MaxAmountOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetMaxAmountOutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAmountOut.Get(), o.MaxAmountOut.IsSet()
}

// HasMaxAmountOut returns a boolean if a field has been set.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) HasMaxAmountOut() bool {
	if o != nil && o.MaxAmountOut.IsSet() {
		return true
	}

	return false
}

// SetMaxAmountOut gets a reference to the given NullableString and assigns it to the MaxAmountOut field.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetMaxAmountOut(v string) {
	o.MaxAmountOut.Set(&v)
}
// SetMaxAmountOutNil sets the value for MaxAmountOut to be an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetMaxAmountOutNil() {
	o.MaxAmountOut.Set(nil)
}

// UnsetMaxAmountOut ensures that no value is present for MaxAmountOut, not even an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) UnsetMaxAmountOut() {
	o.MaxAmountOut.Unset()
}

// GetMinAmountOut returns the MinAmountOut field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetMinAmountOut() string {
	if o == nil || IsNil(o.MinAmountOut.Get()) {
		var ret string
		return ret
	}
	return *o.MinAmountOut.Get()
}

// GetMinAmountOutOk returns a tuple with the MinAmountOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetMinAmountOutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinAmountOut.Get(), o.MinAmountOut.IsSet()
}

// HasMinAmountOut returns a boolean if a field has been set.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) HasMinAmountOut() bool {
	if o != nil && o.MinAmountOut.IsSet() {
		return true
	}

	return false
}

// SetMinAmountOut gets a reference to the given NullableString and assigns it to the MinAmountOut field.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetMinAmountOut(v string) {
	o.MinAmountOut.Set(&v)
}
// SetMinAmountOutNil sets the value for MinAmountOut to be an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetMinAmountOutNil() {
	o.MinAmountOut.Set(nil)
}

// UnsetMinAmountOut ensures that no value is present for MinAmountOut, not even an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) UnsetMinAmountOut() {
	o.MinAmountOut.Unset()
}

// GetTokenOutAddress returns the TokenOutAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetTokenOutAddress() string {
	if o == nil || IsNil(o.TokenOutAddress.Get()) {
		var ret string
		return ret
	}
	return *o.TokenOutAddress.Get()
}

// GetTokenOutAddressOk returns a tuple with the TokenOutAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) GetTokenOutAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOutAddress.Get(), o.TokenOutAddress.IsSet()
}

// HasTokenOutAddress returns a boolean if a field has been set.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) HasTokenOutAddress() bool {
	if o != nil && o.TokenOutAddress.IsSet() {
		return true
	}

	return false
}

// SetTokenOutAddress gets a reference to the given NullableString and assigns it to the TokenOutAddress field.
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetTokenOutAddress(v string) {
	o.TokenOutAddress.Set(&v)
}
// SetTokenOutAddressNil sets the value for TokenOutAddress to be an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) SetTokenOutAddressNil() {
	o.TokenOutAddress.Set(nil)
}

// UnsetTokenOutAddress ensures that no value is present for TokenOutAddress, not even an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeDTO) UnsetTokenOutAddress() {
	o.TokenOutAddress.Unset()
}

func (o Web3EmplTokenExchangeConversionAndFeeDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3EmplTokenExchangeConversionAndFeeDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EmplAmount.IsSet() {
		toSerialize["empl_amount"] = o.EmplAmount.Get()
	}
	if o.EmplTxFee.IsSet() {
		toSerialize["empl_tx_fee"] = o.EmplTxFee.Get()
	}
	if o.MaxAmountOut.IsSet() {
		toSerialize["max_amount_out"] = o.MaxAmountOut.Get()
	}
	if o.MinAmountOut.IsSet() {
		toSerialize["min_amount_out"] = o.MinAmountOut.Get()
	}
	if o.TokenOutAddress.IsSet() {
		toSerialize["token_out_address"] = o.TokenOutAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3EmplTokenExchangeConversionAndFeeDTO) UnmarshalJSON(data []byte) (err error) {
	varWeb3EmplTokenExchangeConversionAndFeeDTO := _Web3EmplTokenExchangeConversionAndFeeDTO{}

	err = json.Unmarshal(data, &varWeb3EmplTokenExchangeConversionAndFeeDTO)

	if err != nil {
		return err
	}

	*o = Web3EmplTokenExchangeConversionAndFeeDTO(varWeb3EmplTokenExchangeConversionAndFeeDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "empl_amount")
		delete(additionalProperties, "empl_tx_fee")
		delete(additionalProperties, "max_amount_out")
		delete(additionalProperties, "min_amount_out")
		delete(additionalProperties, "token_out_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3EmplTokenExchangeConversionAndFeeDTO struct {
	value *Web3EmplTokenExchangeConversionAndFeeDTO
	isSet bool
}

func (v NullableWeb3EmplTokenExchangeConversionAndFeeDTO) Get() *Web3EmplTokenExchangeConversionAndFeeDTO {
	return v.value
}

func (v *NullableWeb3EmplTokenExchangeConversionAndFeeDTO) Set(val *Web3EmplTokenExchangeConversionAndFeeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3EmplTokenExchangeConversionAndFeeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3EmplTokenExchangeConversionAndFeeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3EmplTokenExchangeConversionAndFeeDTO(val *Web3EmplTokenExchangeConversionAndFeeDTO) *NullableWeb3EmplTokenExchangeConversionAndFeeDTO {
	return &NullableWeb3EmplTokenExchangeConversionAndFeeDTO{value: val, isSet: true}
}

func (v NullableWeb3EmplTokenExchangeConversionAndFeeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3EmplTokenExchangeConversionAndFeeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


