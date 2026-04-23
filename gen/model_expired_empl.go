
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ExpiredEmpl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpiredEmpl{}

// ExpiredEmpl struct for ExpiredEmpl
type ExpiredEmpl struct {
	ExpiringAmount NullableString `json:"expiring_amount,omitempty"`
	ExpiringDate NullableInt64 `json:"expiring_date,omitempty"`
	Soon NullableBool `json:"soon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExpiredEmpl ExpiredEmpl

// NewExpiredEmpl instantiates a new ExpiredEmpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiredEmpl() *ExpiredEmpl {
	this := ExpiredEmpl{}
	return &this
}

// NewExpiredEmplWithDefaults instantiates a new ExpiredEmpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiredEmplWithDefaults() *ExpiredEmpl {
	this := ExpiredEmpl{}
	return &this
}

// GetExpiringAmount returns the ExpiringAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpiredEmpl) GetExpiringAmount() string {
	if o == nil || IsNil(o.ExpiringAmount.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiringAmount.Get()
}

// GetExpiringAmountOk returns a tuple with the ExpiringAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpiredEmpl) GetExpiringAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiringAmount.Get(), o.ExpiringAmount.IsSet()
}

// HasExpiringAmount returns a boolean if a field has been set.
func (o *ExpiredEmpl) HasExpiringAmount() bool {
	if o != nil && o.ExpiringAmount.IsSet() {
		return true
	}

	return false
}

// SetExpiringAmount gets a reference to the given NullableString and assigns it to the ExpiringAmount field.
func (o *ExpiredEmpl) SetExpiringAmount(v string) {
	o.ExpiringAmount.Set(&v)
}
// SetExpiringAmountNil sets the value for ExpiringAmount to be an explicit nil
func (o *ExpiredEmpl) SetExpiringAmountNil() {
	o.ExpiringAmount.Set(nil)
}

// UnsetExpiringAmount ensures that no value is present for ExpiringAmount, not even an explicit nil
func (o *ExpiredEmpl) UnsetExpiringAmount() {
	o.ExpiringAmount.Unset()
}

// GetExpiringDate returns the ExpiringDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpiredEmpl) GetExpiringDate() int64 {
	if o == nil || IsNil(o.ExpiringDate.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiringDate.Get()
}

// GetExpiringDateOk returns a tuple with the ExpiringDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpiredEmpl) GetExpiringDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiringDate.Get(), o.ExpiringDate.IsSet()
}

// HasExpiringDate returns a boolean if a field has been set.
func (o *ExpiredEmpl) HasExpiringDate() bool {
	if o != nil && o.ExpiringDate.IsSet() {
		return true
	}

	return false
}

// SetExpiringDate gets a reference to the given NullableInt64 and assigns it to the ExpiringDate field.
func (o *ExpiredEmpl) SetExpiringDate(v int64) {
	o.ExpiringDate.Set(&v)
}
// SetExpiringDateNil sets the value for ExpiringDate to be an explicit nil
func (o *ExpiredEmpl) SetExpiringDateNil() {
	o.ExpiringDate.Set(nil)
}

// UnsetExpiringDate ensures that no value is present for ExpiringDate, not even an explicit nil
func (o *ExpiredEmpl) UnsetExpiringDate() {
	o.ExpiringDate.Unset()
}

// GetSoon returns the Soon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExpiredEmpl) GetSoon() bool {
	if o == nil || IsNil(o.Soon.Get()) {
		var ret bool
		return ret
	}
	return *o.Soon.Get()
}

// GetSoonOk returns a tuple with the Soon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExpiredEmpl) GetSoonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Soon.Get(), o.Soon.IsSet()
}

// HasSoon returns a boolean if a field has been set.
func (o *ExpiredEmpl) HasSoon() bool {
	if o != nil && o.Soon.IsSet() {
		return true
	}

	return false
}

// SetSoon gets a reference to the given NullableBool and assigns it to the Soon field.
func (o *ExpiredEmpl) SetSoon(v bool) {
	o.Soon.Set(&v)
}
// SetSoonNil sets the value for Soon to be an explicit nil
func (o *ExpiredEmpl) SetSoonNil() {
	o.Soon.Set(nil)
}

// UnsetSoon ensures that no value is present for Soon, not even an explicit nil
func (o *ExpiredEmpl) UnsetSoon() {
	o.Soon.Unset()
}

func (o ExpiredEmpl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpiredEmpl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiringAmount.IsSet() {
		toSerialize["expiring_amount"] = o.ExpiringAmount.Get()
	}
	if o.ExpiringDate.IsSet() {
		toSerialize["expiring_date"] = o.ExpiringDate.Get()
	}
	if o.Soon.IsSet() {
		toSerialize["soon"] = o.Soon.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExpiredEmpl) UnmarshalJSON(data []byte) (err error) {
	varExpiredEmpl := _ExpiredEmpl{}

	err = json.Unmarshal(data, &varExpiredEmpl)

	if err != nil {
		return err
	}

	*o = ExpiredEmpl(varExpiredEmpl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiring_amount")
		delete(additionalProperties, "expiring_date")
		delete(additionalProperties, "soon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpiredEmpl struct {
	value *ExpiredEmpl
	isSet bool
}

func (v NullableExpiredEmpl) Get() *ExpiredEmpl {
	return v.value
}

func (v *NullableExpiredEmpl) Set(val *ExpiredEmpl) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiredEmpl) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiredEmpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiredEmpl(val *ExpiredEmpl) *NullableExpiredEmpl {
	return &NullableExpiredEmpl{value: val, isSet: true}
}

func (v NullableExpiredEmpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiredEmpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


