
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelCoinExpiration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCoinExpiration{}

// ModelCoinExpiration struct for ModelCoinExpiration
type ModelCoinExpiration struct {
	Amount NullableFloat64 `json:"amount,omitempty"`
	ExpiredAtSeconds NullableInt64 `json:"expired_at_seconds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelCoinExpiration ModelCoinExpiration

// NewModelCoinExpiration instantiates a new ModelCoinExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCoinExpiration() *ModelCoinExpiration {
	this := ModelCoinExpiration{}
	return &this
}

// NewModelCoinExpirationWithDefaults instantiates a new ModelCoinExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCoinExpirationWithDefaults() *ModelCoinExpiration {
	this := ModelCoinExpiration{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCoinExpiration) GetAmount() float64 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret float64
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCoinExpiration) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *ModelCoinExpiration) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableFloat64 and assigns it to the Amount field.
func (o *ModelCoinExpiration) SetAmount(v float64) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *ModelCoinExpiration) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *ModelCoinExpiration) UnsetAmount() {
	o.Amount.Unset()
}

// GetExpiredAtSeconds returns the ExpiredAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCoinExpiration) GetExpiredAtSeconds() int64 {
	if o == nil || IsNil(o.ExpiredAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiredAtSeconds.Get()
}

// GetExpiredAtSecondsOk returns a tuple with the ExpiredAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCoinExpiration) GetExpiredAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAtSeconds.Get(), o.ExpiredAtSeconds.IsSet()
}

// HasExpiredAtSeconds returns a boolean if a field has been set.
func (o *ModelCoinExpiration) HasExpiredAtSeconds() bool {
	if o != nil && o.ExpiredAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetExpiredAtSeconds gets a reference to the given NullableInt64 and assigns it to the ExpiredAtSeconds field.
func (o *ModelCoinExpiration) SetExpiredAtSeconds(v int64) {
	o.ExpiredAtSeconds.Set(&v)
}
// SetExpiredAtSecondsNil sets the value for ExpiredAtSeconds to be an explicit nil
func (o *ModelCoinExpiration) SetExpiredAtSecondsNil() {
	o.ExpiredAtSeconds.Set(nil)
}

// UnsetExpiredAtSeconds ensures that no value is present for ExpiredAtSeconds, not even an explicit nil
func (o *ModelCoinExpiration) UnsetExpiredAtSeconds() {
	o.ExpiredAtSeconds.Unset()
}

func (o ModelCoinExpiration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCoinExpiration) ToMap() (map[string]interface{}, error) {
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

func (o *ModelCoinExpiration) UnmarshalJSON(data []byte) (err error) {
	varModelCoinExpiration := _ModelCoinExpiration{}

	err = json.Unmarshal(data, &varModelCoinExpiration)

	if err != nil {
		return err
	}

	*o = ModelCoinExpiration(varModelCoinExpiration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "expired_at_seconds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelCoinExpiration struct {
	value *ModelCoinExpiration
	isSet bool
}

func (v NullableModelCoinExpiration) Get() *ModelCoinExpiration {
	return v.value
}

func (v *NullableModelCoinExpiration) Set(val *ModelCoinExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCoinExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCoinExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCoinExpiration(val *ModelCoinExpiration) *NullableModelCoinExpiration {
	return &NullableModelCoinExpiration{value: val, isSet: true}
}

func (v NullableModelCoinExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCoinExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


