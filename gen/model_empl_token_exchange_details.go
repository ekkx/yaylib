
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplTokenExchangeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplTokenExchangeDetails{}

// EmplTokenExchangeDetails struct for EmplTokenExchangeDetails
type EmplTokenExchangeDetails struct {
	ChainId NullableInt64 `json:"chain_id,omitempty"`
	MaxTokenAmount NullableString `json:"max_token_amount,omitempty"`
	MinTokenAmount NullableString `json:"min_token_amount,omitempty"`
	TokenOutAddress NullableString `json:"token_out_address,omitempty"`
	TokenSymbol NullableString `json:"token_symbol,omitempty"`
	TransactionFee NullableFloat64 `json:"transaction_fee,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplTokenExchangeDetails EmplTokenExchangeDetails

// NewEmplTokenExchangeDetails instantiates a new EmplTokenExchangeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplTokenExchangeDetails() *EmplTokenExchangeDetails {
	this := EmplTokenExchangeDetails{}
	return &this
}

// NewEmplTokenExchangeDetailsWithDefaults instantiates a new EmplTokenExchangeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplTokenExchangeDetailsWithDefaults() *EmplTokenExchangeDetails {
	this := EmplTokenExchangeDetails{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTokenExchangeDetails) GetChainId() int64 {
	if o == nil || IsNil(o.ChainId.Get()) {
		var ret int64
		return ret
	}
	return *o.ChainId.Get()
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTokenExchangeDetails) GetChainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChainId.Get(), o.ChainId.IsSet()
}

// HasChainId returns a boolean if a field has been set.
func (o *EmplTokenExchangeDetails) HasChainId() bool {
	if o != nil && o.ChainId.IsSet() {
		return true
	}

	return false
}

// SetChainId gets a reference to the given NullableInt64 and assigns it to the ChainId field.
func (o *EmplTokenExchangeDetails) SetChainId(v int64) {
	o.ChainId.Set(&v)
}
// SetChainIdNil sets the value for ChainId to be an explicit nil
func (o *EmplTokenExchangeDetails) SetChainIdNil() {
	o.ChainId.Set(nil)
}

// UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
func (o *EmplTokenExchangeDetails) UnsetChainId() {
	o.ChainId.Unset()
}

// GetMaxTokenAmount returns the MaxTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTokenExchangeDetails) GetMaxTokenAmount() string {
	if o == nil || IsNil(o.MaxTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MaxTokenAmount.Get()
}

// GetMaxTokenAmountOk returns a tuple with the MaxTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTokenExchangeDetails) GetMaxTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokenAmount.Get(), o.MaxTokenAmount.IsSet()
}

// HasMaxTokenAmount returns a boolean if a field has been set.
func (o *EmplTokenExchangeDetails) HasMaxTokenAmount() bool {
	if o != nil && o.MaxTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxTokenAmount gets a reference to the given NullableString and assigns it to the MaxTokenAmount field.
func (o *EmplTokenExchangeDetails) SetMaxTokenAmount(v string) {
	o.MaxTokenAmount.Set(&v)
}
// SetMaxTokenAmountNil sets the value for MaxTokenAmount to be an explicit nil
func (o *EmplTokenExchangeDetails) SetMaxTokenAmountNil() {
	o.MaxTokenAmount.Set(nil)
}

// UnsetMaxTokenAmount ensures that no value is present for MaxTokenAmount, not even an explicit nil
func (o *EmplTokenExchangeDetails) UnsetMaxTokenAmount() {
	o.MaxTokenAmount.Unset()
}

// GetMinTokenAmount returns the MinTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTokenExchangeDetails) GetMinTokenAmount() string {
	if o == nil || IsNil(o.MinTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MinTokenAmount.Get()
}

// GetMinTokenAmountOk returns a tuple with the MinTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTokenExchangeDetails) GetMinTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinTokenAmount.Get(), o.MinTokenAmount.IsSet()
}

// HasMinTokenAmount returns a boolean if a field has been set.
func (o *EmplTokenExchangeDetails) HasMinTokenAmount() bool {
	if o != nil && o.MinTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMinTokenAmount gets a reference to the given NullableString and assigns it to the MinTokenAmount field.
func (o *EmplTokenExchangeDetails) SetMinTokenAmount(v string) {
	o.MinTokenAmount.Set(&v)
}
// SetMinTokenAmountNil sets the value for MinTokenAmount to be an explicit nil
func (o *EmplTokenExchangeDetails) SetMinTokenAmountNil() {
	o.MinTokenAmount.Set(nil)
}

// UnsetMinTokenAmount ensures that no value is present for MinTokenAmount, not even an explicit nil
func (o *EmplTokenExchangeDetails) UnsetMinTokenAmount() {
	o.MinTokenAmount.Unset()
}

// GetTokenOutAddress returns the TokenOutAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTokenExchangeDetails) GetTokenOutAddress() string {
	if o == nil || IsNil(o.TokenOutAddress.Get()) {
		var ret string
		return ret
	}
	return *o.TokenOutAddress.Get()
}

// GetTokenOutAddressOk returns a tuple with the TokenOutAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTokenExchangeDetails) GetTokenOutAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOutAddress.Get(), o.TokenOutAddress.IsSet()
}

// HasTokenOutAddress returns a boolean if a field has been set.
func (o *EmplTokenExchangeDetails) HasTokenOutAddress() bool {
	if o != nil && o.TokenOutAddress.IsSet() {
		return true
	}

	return false
}

// SetTokenOutAddress gets a reference to the given NullableString and assigns it to the TokenOutAddress field.
func (o *EmplTokenExchangeDetails) SetTokenOutAddress(v string) {
	o.TokenOutAddress.Set(&v)
}
// SetTokenOutAddressNil sets the value for TokenOutAddress to be an explicit nil
func (o *EmplTokenExchangeDetails) SetTokenOutAddressNil() {
	o.TokenOutAddress.Set(nil)
}

// UnsetTokenOutAddress ensures that no value is present for TokenOutAddress, not even an explicit nil
func (o *EmplTokenExchangeDetails) UnsetTokenOutAddress() {
	o.TokenOutAddress.Unset()
}

// GetTokenSymbol returns the TokenSymbol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTokenExchangeDetails) GetTokenSymbol() string {
	if o == nil || IsNil(o.TokenSymbol.Get()) {
		var ret string
		return ret
	}
	return *o.TokenSymbol.Get()
}

// GetTokenSymbolOk returns a tuple with the TokenSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTokenExchangeDetails) GetTokenSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenSymbol.Get(), o.TokenSymbol.IsSet()
}

// HasTokenSymbol returns a boolean if a field has been set.
func (o *EmplTokenExchangeDetails) HasTokenSymbol() bool {
	if o != nil && o.TokenSymbol.IsSet() {
		return true
	}

	return false
}

// SetTokenSymbol gets a reference to the given NullableString and assigns it to the TokenSymbol field.
func (o *EmplTokenExchangeDetails) SetTokenSymbol(v string) {
	o.TokenSymbol.Set(&v)
}
// SetTokenSymbolNil sets the value for TokenSymbol to be an explicit nil
func (o *EmplTokenExchangeDetails) SetTokenSymbolNil() {
	o.TokenSymbol.Set(nil)
}

// UnsetTokenSymbol ensures that no value is present for TokenSymbol, not even an explicit nil
func (o *EmplTokenExchangeDetails) UnsetTokenSymbol() {
	o.TokenSymbol.Unset()
}

// GetTransactionFee returns the TransactionFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTokenExchangeDetails) GetTransactionFee() float64 {
	if o == nil || IsNil(o.TransactionFee.Get()) {
		var ret float64
		return ret
	}
	return *o.TransactionFee.Get()
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTokenExchangeDetails) GetTransactionFeeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionFee.Get(), o.TransactionFee.IsSet()
}

// HasTransactionFee returns a boolean if a field has been set.
func (o *EmplTokenExchangeDetails) HasTransactionFee() bool {
	if o != nil && o.TransactionFee.IsSet() {
		return true
	}

	return false
}

// SetTransactionFee gets a reference to the given NullableFloat64 and assigns it to the TransactionFee field.
func (o *EmplTokenExchangeDetails) SetTransactionFee(v float64) {
	o.TransactionFee.Set(&v)
}
// SetTransactionFeeNil sets the value for TransactionFee to be an explicit nil
func (o *EmplTokenExchangeDetails) SetTransactionFeeNil() {
	o.TransactionFee.Set(nil)
}

// UnsetTransactionFee ensures that no value is present for TransactionFee, not even an explicit nil
func (o *EmplTokenExchangeDetails) UnsetTransactionFee() {
	o.TransactionFee.Unset()
}

func (o EmplTokenExchangeDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplTokenExchangeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChainId.IsSet() {
		toSerialize["chain_id"] = o.ChainId.Get()
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

func (o *EmplTokenExchangeDetails) UnmarshalJSON(data []byte) (err error) {
	varEmplTokenExchangeDetails := _EmplTokenExchangeDetails{}

	err = json.Unmarshal(data, &varEmplTokenExchangeDetails)

	if err != nil {
		return err
	}

	*o = EmplTokenExchangeDetails(varEmplTokenExchangeDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chain_id")
		delete(additionalProperties, "max_token_amount")
		delete(additionalProperties, "min_token_amount")
		delete(additionalProperties, "token_out_address")
		delete(additionalProperties, "token_symbol")
		delete(additionalProperties, "transaction_fee")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplTokenExchangeDetails struct {
	value *EmplTokenExchangeDetails
	isSet bool
}

func (v NullableEmplTokenExchangeDetails) Get() *EmplTokenExchangeDetails {
	return v.value
}

func (v *NullableEmplTokenExchangeDetails) Set(val *EmplTokenExchangeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplTokenExchangeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplTokenExchangeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplTokenExchangeDetails(val *EmplTokenExchangeDetails) *NullableEmplTokenExchangeDetails {
	return &NullableEmplTokenExchangeDetails{value: val, isSet: true}
}

func (v NullableEmplTokenExchangeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplTokenExchangeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


