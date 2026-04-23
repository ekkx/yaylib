
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletEmplExpiring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletEmplExpiring{}

// Web3WalletEmplExpiring struct for Web3WalletEmplExpiring
type Web3WalletEmplExpiring struct {
	ExpiringAmount NullableFloat64 `json:"expiring_amount,omitempty"`
	ExpiringDateMillis NullableInt64 `json:"expiring_date_millis,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletEmplExpiring Web3WalletEmplExpiring

// NewWeb3WalletEmplExpiring instantiates a new Web3WalletEmplExpiring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletEmplExpiring() *Web3WalletEmplExpiring {
	this := Web3WalletEmplExpiring{}
	return &this
}

// NewWeb3WalletEmplExpiringWithDefaults instantiates a new Web3WalletEmplExpiring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletEmplExpiringWithDefaults() *Web3WalletEmplExpiring {
	this := Web3WalletEmplExpiring{}
	return &this
}

// GetExpiringAmount returns the ExpiringAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletEmplExpiring) GetExpiringAmount() float64 {
	if o == nil || IsNil(o.ExpiringAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.ExpiringAmount.Get()
}

// GetExpiringAmountOk returns a tuple with the ExpiringAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletEmplExpiring) GetExpiringAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiringAmount.Get(), o.ExpiringAmount.IsSet()
}

// HasExpiringAmount returns a boolean if a field has been set.
func (o *Web3WalletEmplExpiring) HasExpiringAmount() bool {
	if o != nil && o.ExpiringAmount.IsSet() {
		return true
	}

	return false
}

// SetExpiringAmount gets a reference to the given NullableFloat64 and assigns it to the ExpiringAmount field.
func (o *Web3WalletEmplExpiring) SetExpiringAmount(v float64) {
	o.ExpiringAmount.Set(&v)
}
// SetExpiringAmountNil sets the value for ExpiringAmount to be an explicit nil
func (o *Web3WalletEmplExpiring) SetExpiringAmountNil() {
	o.ExpiringAmount.Set(nil)
}

// UnsetExpiringAmount ensures that no value is present for ExpiringAmount, not even an explicit nil
func (o *Web3WalletEmplExpiring) UnsetExpiringAmount() {
	o.ExpiringAmount.Unset()
}

// GetExpiringDateMillis returns the ExpiringDateMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletEmplExpiring) GetExpiringDateMillis() int64 {
	if o == nil || IsNil(o.ExpiringDateMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiringDateMillis.Get()
}

// GetExpiringDateMillisOk returns a tuple with the ExpiringDateMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletEmplExpiring) GetExpiringDateMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiringDateMillis.Get(), o.ExpiringDateMillis.IsSet()
}

// HasExpiringDateMillis returns a boolean if a field has been set.
func (o *Web3WalletEmplExpiring) HasExpiringDateMillis() bool {
	if o != nil && o.ExpiringDateMillis.IsSet() {
		return true
	}

	return false
}

// SetExpiringDateMillis gets a reference to the given NullableInt64 and assigns it to the ExpiringDateMillis field.
func (o *Web3WalletEmplExpiring) SetExpiringDateMillis(v int64) {
	o.ExpiringDateMillis.Set(&v)
}
// SetExpiringDateMillisNil sets the value for ExpiringDateMillis to be an explicit nil
func (o *Web3WalletEmplExpiring) SetExpiringDateMillisNil() {
	o.ExpiringDateMillis.Set(nil)
}

// UnsetExpiringDateMillis ensures that no value is present for ExpiringDateMillis, not even an explicit nil
func (o *Web3WalletEmplExpiring) UnsetExpiringDateMillis() {
	o.ExpiringDateMillis.Unset()
}

func (o Web3WalletEmplExpiring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletEmplExpiring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiringAmount.IsSet() {
		toSerialize["expiring_amount"] = o.ExpiringAmount.Get()
	}
	if o.ExpiringDateMillis.IsSet() {
		toSerialize["expiring_date_millis"] = o.ExpiringDateMillis.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletEmplExpiring) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletEmplExpiring := _Web3WalletEmplExpiring{}

	err = json.Unmarshal(data, &varWeb3WalletEmplExpiring)

	if err != nil {
		return err
	}

	*o = Web3WalletEmplExpiring(varWeb3WalletEmplExpiring)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiring_amount")
		delete(additionalProperties, "expiring_date_millis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletEmplExpiring struct {
	value *Web3WalletEmplExpiring
	isSet bool
}

func (v NullableWeb3WalletEmplExpiring) Get() *Web3WalletEmplExpiring {
	return v.value
}

func (v *NullableWeb3WalletEmplExpiring) Set(val *Web3WalletEmplExpiring) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletEmplExpiring) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletEmplExpiring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletEmplExpiring(val *Web3WalletEmplExpiring) *NullableWeb3WalletEmplExpiring {
	return &NullableWeb3WalletEmplExpiring{value: val, isSet: true}
}

func (v NullableWeb3WalletEmplExpiring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletEmplExpiring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


