
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CoinExpiration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoinExpiration{}

// CoinExpiration struct for CoinExpiration
type CoinExpiration struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	ExpiredAtSeconds NullableInt64 `json:"expired_at_seconds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CoinExpiration CoinExpiration

// NewCoinExpiration instantiates a new CoinExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoinExpiration() *CoinExpiration {
	this := CoinExpiration{}
	return &this
}

// NewCoinExpirationWithDefaults instantiates a new CoinExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoinExpirationWithDefaults() *CoinExpiration {
	this := CoinExpiration{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoinExpiration) GetAmount() float64 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoinExpiration) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *CoinExpiration) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *CoinExpiration) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *CoinExpiration) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *CoinExpiration) UnsetAmount() {
	o.Amount.Unset()
}

// GetExpiredAtSeconds returns the ExpiredAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoinExpiration) GetExpiredAtSeconds() int64 {
	if o == nil || IsNil(o.ExpiredAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiredAtSeconds.Get()
}

// GetExpiredAtSecondsOk returns a tuple with the ExpiredAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoinExpiration) GetExpiredAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAtSeconds.Get(), o.ExpiredAtSeconds.IsSet()
}

// HasExpiredAtSeconds returns a boolean if a field has been set.
func (o *CoinExpiration) HasExpiredAtSeconds() bool {
	if o != nil && o.ExpiredAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetExpiredAtSeconds gets a reference to the given NullableInt64 and assigns it to the ExpiredAtSeconds field.
func (o *CoinExpiration) SetExpiredAtSeconds(v int64) {
	o.ExpiredAtSeconds.Set(&v)
}
// SetExpiredAtSecondsNil sets the value for ExpiredAtSeconds to be an explicit nil
func (o *CoinExpiration) SetExpiredAtSecondsNil() {
	o.ExpiredAtSeconds.Set(nil)
}

// UnsetExpiredAtSeconds ensures that no value is present for ExpiredAtSeconds, not even an explicit nil
func (o *CoinExpiration) UnsetExpiredAtSeconds() {
	o.ExpiredAtSeconds.Unset()
}

func (o CoinExpiration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoinExpiration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.ExpiredAtSeconds.IsSet() {
		toSerialize["expired_at_seconds"] = o.ExpiredAtSeconds.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoinExpiration) UnmarshalJSON(data []byte) (err error) {
	varCoinExpiration := _CoinExpiration{}

	err = json.Unmarshal(data, &varCoinExpiration)

	if err != nil {
		return err
	}

	*o = CoinExpiration(varCoinExpiration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "expired_at_seconds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoinExpiration struct {
	value *CoinExpiration
	isSet bool
}

func (v NullableCoinExpiration) Get() *CoinExpiration {
	return v.value
}

func (v *NullableCoinExpiration) Set(val *CoinExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableCoinExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableCoinExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoinExpiration(val *CoinExpiration) *NullableCoinExpiration {
	return &NullableCoinExpiration{value: val, isSet: true}
}

func (v NullableCoinExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoinExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


