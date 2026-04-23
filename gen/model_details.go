
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Details type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Details{}

// Details struct for Details
type Details struct {
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	Error NullableString `json:"error,omitempty"`
	MaxTokenAmount NullableString `json:"max_token_amount,omitempty"`
	MinTokenAmount NullableString `json:"min_token_amount,omitempty"`
	TokenOutAddress NullableString `json:"token_out_address,omitempty"`
	TokenOutAmount NullableString `json:"token_out_amount,omitempty"`
	TokenSymbol NullableString `json:"token_symbol,omitempty"`
	TransactionFee NullableString `json:"transaction_fee,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Details Details

// NewDetails instantiates a new Details object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetails() *Details {
	this := Details{}
	return &this
}

// NewDetailsWithDefaults instantiates a new Details object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailsWithDefaults() *Details {
	this := Details{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *Details) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *Details) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *Details) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *Details) UnsetChainId() {
	o.ChainId.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *Details) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *Details) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *Details) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *Details) UnsetError() {
	o.Error.Unset()
}

// GetMaxTokenAmount returns the MaxTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetMaxTokenAmount() string {
	if o == nil || IsNil(o.MaxTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MaxTokenAmount.Get()
}

// GetMaxTokenAmountOk returns a tuple with the MaxTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetMaxTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokenAmount.Get(), o.MaxTokenAmount.IsSet()
}

// HasMaxTokenAmount returns a boolean if a field has been set.
func (o *Details) HasMaxTokenAmount() bool {
	if o != nil && o.MaxTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxTokenAmount gets a reference to the given NullableString and assigns it to the MaxTokenAmount field.
func (o *Details) SetMaxTokenAmount(v string) {
	o.MaxTokenAmount.Set(&v)
}
// SetMaxTokenAmountNil sets the value for MaxTokenAmount to be an explicit nil
func (o *Details) SetMaxTokenAmountNil() {
	o.MaxTokenAmount.Set(nil)
}

// UnsetMaxTokenAmount ensures that no value is present for MaxTokenAmount, not even an explicit nil
func (o *Details) UnsetMaxTokenAmount() {
	o.MaxTokenAmount.Unset()
}

// GetMinTokenAmount returns the MinTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetMinTokenAmount() string {
	if o == nil || IsNil(o.MinTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MinTokenAmount.Get()
}

// GetMinTokenAmountOk returns a tuple with the MinTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetMinTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinTokenAmount.Get(), o.MinTokenAmount.IsSet()
}

// HasMinTokenAmount returns a boolean if a field has been set.
func (o *Details) HasMinTokenAmount() bool {
	if o != nil && o.MinTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMinTokenAmount gets a reference to the given NullableString and assigns it to the MinTokenAmount field.
func (o *Details) SetMinTokenAmount(v string) {
	o.MinTokenAmount.Set(&v)
}
// SetMinTokenAmountNil sets the value for MinTokenAmount to be an explicit nil
func (o *Details) SetMinTokenAmountNil() {
	o.MinTokenAmount.Set(nil)
}

// UnsetMinTokenAmount ensures that no value is present for MinTokenAmount, not even an explicit nil
func (o *Details) UnsetMinTokenAmount() {
	o.MinTokenAmount.Unset()
}

// GetTokenOutAddress returns the TokenOutAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetTokenOutAddress() string {
	if o == nil || IsNil(o.TokenOutAddress.Get()) {
		var ret string
		return ret
	}
	return *o.TokenOutAddress.Get()
}

// GetTokenOutAddressOk returns a tuple with the TokenOutAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetTokenOutAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOutAddress.Get(), o.TokenOutAddress.IsSet()
}

// HasTokenOutAddress returns a boolean if a field has been set.
func (o *Details) HasTokenOutAddress() bool {
	if o != nil && o.TokenOutAddress.IsSet() {
		return true
	}

	return false
}

// SetTokenOutAddress gets a reference to the given NullableString and assigns it to the TokenOutAddress field.
func (o *Details) SetTokenOutAddress(v string) {
	o.TokenOutAddress.Set(&v)
}
// SetTokenOutAddressNil sets the value for TokenOutAddress to be an explicit nil
func (o *Details) SetTokenOutAddressNil() {
	o.TokenOutAddress.Set(nil)
}

// UnsetTokenOutAddress ensures that no value is present for TokenOutAddress, not even an explicit nil
func (o *Details) UnsetTokenOutAddress() {
	o.TokenOutAddress.Unset()
}

// GetTokenOutAmount returns the TokenOutAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetTokenOutAmount() string {
	if o == nil || IsNil(o.TokenOutAmount.Get()) {
		var ret string
		return ret
	}
	return *o.TokenOutAmount.Get()
}

// GetTokenOutAmountOk returns a tuple with the TokenOutAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetTokenOutAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOutAmount.Get(), o.TokenOutAmount.IsSet()
}

// HasTokenOutAmount returns a boolean if a field has been set.
func (o *Details) HasTokenOutAmount() bool {
	if o != nil && o.TokenOutAmount.IsSet() {
		return true
	}

	return false
}

// SetTokenOutAmount gets a reference to the given NullableString and assigns it to the TokenOutAmount field.
func (o *Details) SetTokenOutAmount(v string) {
	o.TokenOutAmount.Set(&v)
}
// SetTokenOutAmountNil sets the value for TokenOutAmount to be an explicit nil
func (o *Details) SetTokenOutAmountNil() {
	o.TokenOutAmount.Set(nil)
}

// UnsetTokenOutAmount ensures that no value is present for TokenOutAmount, not even an explicit nil
func (o *Details) UnsetTokenOutAmount() {
	o.TokenOutAmount.Unset()
}

// GetTokenSymbol returns the TokenSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetTokenSymbol() string {
	if o == nil || IsNil(o.TokenSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.TokenSymbol.Get()
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSymbol.Get(), o.TokenSymbol.IsSet()
}

// HasTokenSymbol returns a boolean if a field has been set.
func (o *Details) HasTokenSymbol() bool {
	if o != nil && o.TokenSymbol.IsSet() {
		return true
	}

	return false
}

// SetTokenSymbol gets a reference to the given NullableString and assigns it to the TokenSymbol field.
func (o *Details) SetTokenSymbol(v string) {
	o.TokenSymbol.Set(&v)
}
// SetTokenSymbolNil sets the value for TokenSymbol to be an explicit nil
func (o *Details) SetTokenSymbolNil() {
	o.TokenSymbol.Set(nil)
}

// UnsetTokenSymbol ensures that no value is present for TokenSymbol, not even an explicit nil
func (o *Details) UnsetTokenSymbol() {
	o.TokenSymbol.Unset()
}

// GetTransactionFee returns the TransactionFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Details) GetTransactionFee() string {
	if o == nil || IsNil(o.TransactionFee.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionFee.Get()
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Details) GetTransactionFeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionFee.Get(), o.TransactionFee.IsSet()
}

// HasTransactionFee returns a boolean if a field has been set.
func (o *Details) HasTransactionFee() bool {
	if o != nil && o.TransactionFee.IsSet() {
		return true
	}

	return false
}

// SetTransactionFee gets a reference to the given NullableString and assigns it to the TransactionFee field.
func (o *Details) SetTransactionFee(v string) {
	o.TransactionFee.Set(&v)
}
// SetTransactionFeeNil sets the value for TransactionFee to be an explicit nil
func (o *Details) SetTransactionFeeNil() {
	o.TransactionFee.Set(nil)
}

// UnsetTransactionFee ensures that no value is present for TransactionFee, not even an explicit nil
func (o *Details) UnsetTransactionFee() {
	o.TransactionFee.Unset()
}

func (o Details) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Details) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.MaxTokenAmount.IsSet() {
		toSerialize["max_token_amount"] = o.MaxTokenAmount.Get()
	}
	if o.MinTokenAmount.IsSet() {
		toSerialize["min_token_amount"] = o.MinTokenAmount.Get()
	}
	if o.TokenOutAddress.IsSet() {
		toSerialize["token_out_address"] = o.TokenOutAddress.Get()
	}
	if o.TokenOutAmount.IsSet() {
		toSerialize["token_out_amount"] = o.TokenOutAmount.Get()
	}
	if o.TokenSymbol.IsSet() {
		toSerialize["token_symbol"] = o.TokenSymbol.Get()
	}
	if o.TransactionFee.IsSet() {
		toSerialize["transaction_fee"] = o.TransactionFee.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Details) UnmarshalJSON(data []byte) (err error) {
	varDetails := _Details{}

	err = json.Unmarshal(data, &varDetails)

	if err != nil {
		return err
	}

	*o = Details(varDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "error")
		delete(additionalProperties, "max_token_amount")
		delete(additionalProperties, "min_token_amount")
		delete(additionalProperties, "token_out_address")
		delete(additionalProperties, "token_out_amount")
		delete(additionalProperties, "token_symbol")
		delete(additionalProperties, "transaction_fee")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDetails struct {
	value *Details
	isSet bool
}

func (v NullableDetails) Get() *Details {
	return v.value
}

func (v *NullableDetails) Set(val *Details) {
	v.value = val
	v.isSet = true
}

func (v NullableDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetails(val *Details) *NullableDetails {
	return &NullableDetails{value: val, isSet: true}
}

func (v NullableDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


