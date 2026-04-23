
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PendingEmplActivityDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingEmplActivityDetails{}

// PendingEmplActivityDetails struct for PendingEmplActivityDetails
type PendingEmplActivityDetails struct {
	MaxTokenAmount NullableString `json:"max_token_amount,omitempty"`
	MinTokenAmount NullableString `json:"min_token_amount,omitempty"`
	TokenOutAddress NullableString `json:"token_out_address,omitempty"`
	TransactionFee NullableFloat64 `json:"transaction_fee,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PendingEmplActivityDetails PendingEmplActivityDetails

// NewPendingEmplActivityDetails instantiates a new PendingEmplActivityDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingEmplActivityDetails() *PendingEmplActivityDetails {
	this := PendingEmplActivityDetails{}
	return &this
}

// NewPendingEmplActivityDetailsWithDefaults instantiates a new PendingEmplActivityDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingEmplActivityDetailsWithDefaults() *PendingEmplActivityDetails {
	this := PendingEmplActivityDetails{}
	return &this
}

// GetMaxTokenAmount returns the MaxTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingEmplActivityDetails) GetMaxTokenAmount() string {
	if o == nil || IsNil(o.MaxTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MaxTokenAmount.Get()
}

// GetMaxTokenAmountOk returns a tuple with the MaxTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingEmplActivityDetails) GetMaxTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokenAmount.Get(), o.MaxTokenAmount.IsSet()
}

// HasMaxTokenAmount returns a boolean if a field has been set.
func (o *PendingEmplActivityDetails) HasMaxTokenAmount() bool {
	if o != nil && o.MaxTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMaxTokenAmount gets a reference to the given NullableString and assigns it to the MaxTokenAmount field.
func (o *PendingEmplActivityDetails) SetMaxTokenAmount(v string) {
	o.MaxTokenAmount.Set(&v)
}
// SetMaxTokenAmountNil sets the value for MaxTokenAmount to be an explicit nil
func (o *PendingEmplActivityDetails) SetMaxTokenAmountNil() {
	o.MaxTokenAmount.Set(nil)
}

// UnsetMaxTokenAmount ensures that no value is present for MaxTokenAmount, not even an explicit nil
func (o *PendingEmplActivityDetails) UnsetMaxTokenAmount() {
	o.MaxTokenAmount.Unset()
}

// GetMinTokenAmount returns the MinTokenAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingEmplActivityDetails) GetMinTokenAmount() string {
	if o == nil || IsNil(o.MinTokenAmount.Get()) {
		var ret string
		return ret
	}
	return *o.MinTokenAmount.Get()
}

// GetMinTokenAmountOk returns a tuple with the MinTokenAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingEmplActivityDetails) GetMinTokenAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinTokenAmount.Get(), o.MinTokenAmount.IsSet()
}

// HasMinTokenAmount returns a boolean if a field has been set.
func (o *PendingEmplActivityDetails) HasMinTokenAmount() bool {
	if o != nil && o.MinTokenAmount.IsSet() {
		return true
	}

	return false
}

// SetMinTokenAmount gets a reference to the given NullableString and assigns it to the MinTokenAmount field.
func (o *PendingEmplActivityDetails) SetMinTokenAmount(v string) {
	o.MinTokenAmount.Set(&v)
}
// SetMinTokenAmountNil sets the value for MinTokenAmount to be an explicit nil
func (o *PendingEmplActivityDetails) SetMinTokenAmountNil() {
	o.MinTokenAmount.Set(nil)
}

// UnsetMinTokenAmount ensures that no value is present for MinTokenAmount, not even an explicit nil
func (o *PendingEmplActivityDetails) UnsetMinTokenAmount() {
	o.MinTokenAmount.Unset()
}

// GetTokenOutAddress returns the TokenOutAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingEmplActivityDetails) GetTokenOutAddress() string {
	if o == nil || IsNil(o.TokenOutAddress.Get()) {
		var ret string
		return ret
	}
	return *o.TokenOutAddress.Get()
}

// GetTokenOutAddressOk returns a tuple with the TokenOutAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingEmplActivityDetails) GetTokenOutAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOutAddress.Get(), o.TokenOutAddress.IsSet()
}

// HasTokenOutAddress returns a boolean if a field has been set.
func (o *PendingEmplActivityDetails) HasTokenOutAddress() bool {
	if o != nil && o.TokenOutAddress.IsSet() {
		return true
	}

	return false
}

// SetTokenOutAddress gets a reference to the given NullableString and assigns it to the TokenOutAddress field.
func (o *PendingEmplActivityDetails) SetTokenOutAddress(v string) {
	o.TokenOutAddress.Set(&v)
}
// SetTokenOutAddressNil sets the value for TokenOutAddress to be an explicit nil
func (o *PendingEmplActivityDetails) SetTokenOutAddressNil() {
	o.TokenOutAddress.Set(nil)
}

// UnsetTokenOutAddress ensures that no value is present for TokenOutAddress, not even an explicit nil
func (o *PendingEmplActivityDetails) UnsetTokenOutAddress() {
	o.TokenOutAddress.Unset()
}

// GetTransactionFee returns the TransactionFee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingEmplActivityDetails) GetTransactionFee() float64 {
	if o == nil || IsNil(o.TransactionFee.Get()) {
		var ret float64
		return ret
	}
	return *o.TransactionFee.Get()
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingEmplActivityDetails) GetTransactionFeeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionFee.Get(), o.TransactionFee.IsSet()
}

// HasTransactionFee returns a boolean if a field has been set.
func (o *PendingEmplActivityDetails) HasTransactionFee() bool {
	if o != nil && o.TransactionFee.IsSet() {
		return true
	}

	return false
}

// SetTransactionFee gets a reference to the given NullableFloat64 and assigns it to the TransactionFee field.
func (o *PendingEmplActivityDetails) SetTransactionFee(v float64) {
	o.TransactionFee.Set(&v)
}
// SetTransactionFeeNil sets the value for TransactionFee to be an explicit nil
func (o *PendingEmplActivityDetails) SetTransactionFeeNil() {
	o.TransactionFee.Set(nil)
}

// UnsetTransactionFee ensures that no value is present for TransactionFee, not even an explicit nil
func (o *PendingEmplActivityDetails) UnsetTransactionFee() {
	o.TransactionFee.Unset()
}

func (o PendingEmplActivityDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingEmplActivityDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxTokenAmount.IsSet() {
		toSerialize["max_token_amount"] = o.MaxTokenAmount.Get()
	}
	if o.MinTokenAmount.IsSet() {
		toSerialize["min_token_amount"] = o.MinTokenAmount.Get()
	}
	if o.TokenOutAddress.IsSet() {
		toSerialize["token_out_address"] = o.TokenOutAddress.Get()
	}
	if o.TransactionFee.IsSet() {
		toSerialize["transaction_fee"] = o.TransactionFee.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PendingEmplActivityDetails) UnmarshalJSON(data []byte) (err error) {
	varPendingEmplActivityDetails := _PendingEmplActivityDetails{}

	err = json.Unmarshal(data, &varPendingEmplActivityDetails)

	if err != nil {
		return err
	}

	*o = PendingEmplActivityDetails(varPendingEmplActivityDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max_token_amount")
		delete(additionalProperties, "min_token_amount")
		delete(additionalProperties, "token_out_address")
		delete(additionalProperties, "transaction_fee")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePendingEmplActivityDetails struct {
	value *PendingEmplActivityDetails
	isSet bool
}

func (v NullablePendingEmplActivityDetails) Get() *PendingEmplActivityDetails {
	return v.value
}

func (v *NullablePendingEmplActivityDetails) Set(val *PendingEmplActivityDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingEmplActivityDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingEmplActivityDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingEmplActivityDetails(val *PendingEmplActivityDetails) *NullablePendingEmplActivityDetails {
	return &NullablePendingEmplActivityDetails{value: val, isSet: true}
}

func (v NullablePendingEmplActivityDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingEmplActivityDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


