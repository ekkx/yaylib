
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3ExpiringEmplExpiredEmpl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3ExpiringEmplExpiredEmpl{}

// Web3ExpiringEmplExpiredEmpl struct for Web3ExpiringEmplExpiredEmpl
type Web3ExpiringEmplExpiredEmpl struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	DateInMillis NullableInt64 `json:"date_in_millis,omitempty"`
	Soon NullableBool `json:"soon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3ExpiringEmplExpiredEmpl Web3ExpiringEmplExpiredEmpl

// NewWeb3ExpiringEmplExpiredEmpl instantiates a new Web3ExpiringEmplExpiredEmpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3ExpiringEmplExpiredEmpl() *Web3ExpiringEmplExpiredEmpl {
	this := Web3ExpiringEmplExpiredEmpl{}
	return &this
}

// NewWeb3ExpiringEmplExpiredEmplWithDefaults instantiates a new Web3ExpiringEmplExpiredEmpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3ExpiringEmplExpiredEmplWithDefaults() *Web3ExpiringEmplExpiredEmpl {
	this := Web3ExpiringEmplExpiredEmpl{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3ExpiringEmplExpiredEmpl) GetAmount() float64 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3ExpiringEmplExpiredEmpl) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *Web3ExpiringEmplExpiredEmpl) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *Web3ExpiringEmplExpiredEmpl) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *Web3ExpiringEmplExpiredEmpl) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *Web3ExpiringEmplExpiredEmpl) UnsetAmount() {
	o.Amount.Unset()
}

// GetDateInMillis returns the DateInMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3ExpiringEmplExpiredEmpl) GetDateInMillis() int64 {
	if o == nil || IsNil(o.DateInMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.DateInMillis.Get()
}

// GetDateInMillisOk returns a tuple with the DateInMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3ExpiringEmplExpiredEmpl) GetDateInMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateInMillis.Get(), o.DateInMillis.IsSet()
}

// HasDateInMillis returns a boolean if a field has been set.
func (o *Web3ExpiringEmplExpiredEmpl) HasDateInMillis() bool {
	if o != nil && o.DateInMillis.IsSet() {
		return true
	}

	return false
}

// SetDateInMillis gets a reference to the given NullableInt64 and assigns it to the DateInMillis field.
func (o *Web3ExpiringEmplExpiredEmpl) SetDateInMillis(v int64) {
	o.DateInMillis.Set(&v)
}
// SetDateInMillisNil sets the value for DateInMillis to be an explicit nil
func (o *Web3ExpiringEmplExpiredEmpl) SetDateInMillisNil() {
	o.DateInMillis.Set(nil)
}

// UnsetDateInMillis ensures that no value is present for DateInMillis, not even an explicit nil
func (o *Web3ExpiringEmplExpiredEmpl) UnsetDateInMillis() {
	o.DateInMillis.Unset()
}

// GetSoon returns the Soon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3ExpiringEmplExpiredEmpl) GetSoon() bool {
	if o == nil || IsNil(o.Soon.Get()) {
		var ret bool
		return ret
	}
	return *o.Soon.Get()
}

// GetSoonOk returns a tuple with the Soon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3ExpiringEmplExpiredEmpl) GetSoonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Soon.Get(), o.Soon.IsSet()
}

// HasSoon returns a boolean if a field has been set.
func (o *Web3ExpiringEmplExpiredEmpl) HasSoon() bool {
	if o != nil && o.Soon.IsSet() {
		return true
	}

	return false
}

// SetSoon gets a reference to the given NullableBool and assigns it to the Soon field.
func (o *Web3ExpiringEmplExpiredEmpl) SetSoon(v bool) {
	o.Soon.Set(&v)
}
// SetSoonNil sets the value for Soon to be an explicit nil
func (o *Web3ExpiringEmplExpiredEmpl) SetSoonNil() {
	o.Soon.Set(nil)
}

// UnsetSoon ensures that no value is present for Soon, not even an explicit nil
func (o *Web3ExpiringEmplExpiredEmpl) UnsetSoon() {
	o.Soon.Unset()
}

func (o Web3ExpiringEmplExpiredEmpl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3ExpiringEmplExpiredEmpl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.DateInMillis.IsSet() {
		toSerialize["date_in_millis"] = o.DateInMillis.Get()
	}
	if o.Soon.IsSet() {
		toSerialize["soon"] = o.Soon.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3ExpiringEmplExpiredEmpl) UnmarshalJSON(data []byte) (err error) {
	varWeb3ExpiringEmplExpiredEmpl := _Web3ExpiringEmplExpiredEmpl{}

	err = json.Unmarshal(data, &varWeb3ExpiringEmplExpiredEmpl)

	if err != nil {
		return err
	}

	*o = Web3ExpiringEmplExpiredEmpl(varWeb3ExpiringEmplExpiredEmpl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "date_in_millis")
		delete(additionalProperties, "soon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3ExpiringEmplExpiredEmpl struct {
	value *Web3ExpiringEmplExpiredEmpl
	isSet bool
}

func (v NullableWeb3ExpiringEmplExpiredEmpl) Get() *Web3ExpiringEmplExpiredEmpl {
	return v.value
}

func (v *NullableWeb3ExpiringEmplExpiredEmpl) Set(val *Web3ExpiringEmplExpiredEmpl) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3ExpiringEmplExpiredEmpl) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3ExpiringEmplExpiredEmpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3ExpiringEmplExpiredEmpl(val *Web3ExpiringEmplExpiredEmpl) *NullableWeb3ExpiringEmplExpiredEmpl {
	return &NullableWeb3ExpiringEmplExpiredEmpl{value: val, isSet: true}
}

func (v NullableWeb3ExpiringEmplExpiredEmpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3ExpiringEmplExpiredEmpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


