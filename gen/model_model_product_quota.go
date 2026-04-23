
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelProductQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelProductQuota{}

// ModelProductQuota struct for ModelProductQuota
type ModelProductQuota struct {
	Bought NullableFloat64 `json:"bought,omitempty"`
	CurrencyCode NullableString `json:"currency_code,omitempty"`
	Limit NullableFloat64 `json:"limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelProductQuota ModelProductQuota

// NewModelProductQuota instantiates a new ModelProductQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelProductQuota() *ModelProductQuota {
	this := ModelProductQuota{}
	return &this
}

// NewModelProductQuotaWithDefaults instantiates a new ModelProductQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelProductQuotaWithDefaults() *ModelProductQuota {
	this := ModelProductQuota{}
	return &this
}

// GetBought returns the Bought field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelProductQuota) GetBought() float64 {
	if o == nil || IsNil(o.Bought.Get()) {
		var ret float64
		return ret
	}
	return *o.Bought.Get()
}

// GetBoughtOk returns a tuple with the Bought field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelProductQuota) GetBoughtOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bought.Get(), o.Bought.IsSet()
}

// HasBought returns a boolean if a field has been set.
func (o *ModelProductQuota) HasBought() bool {
	if o != nil && o.Bought.IsSet() {
		return true
	}

	return false
}

// SetBought gets a reference to the given NullableFloat64 and assigns it to the Bought field.
func (o *ModelProductQuota) SetBought(v float64) {
	o.Bought.Set(&v)
}
// SetBoughtNil sets the value for Bought to be an explicit nil
func (o *ModelProductQuota) SetBoughtNil() {
	o.Bought.Set(nil)
}

// UnsetBought ensures that no value is present for Bought, not even an explicit nil
func (o *ModelProductQuota) UnsetBought() {
	o.Bought.Unset()
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelProductQuota) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyCode.Get()
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelProductQuota) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyCode.Get(), o.CurrencyCode.IsSet()
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ModelProductQuota) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given NullableString and assigns it to the CurrencyCode field.
func (o *ModelProductQuota) SetCurrencyCode(v string) {
	o.CurrencyCode.Set(&v)
}
// SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil
func (o *ModelProductQuota) SetCurrencyCodeNil() {
	o.CurrencyCode.Set(nil)
}

// UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
func (o *ModelProductQuota) UnsetCurrencyCode() {
	o.CurrencyCode.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelProductQuota) GetLimit() float64 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret float64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelProductQuota) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *ModelProductQuota) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableFloat64 and assigns it to the Limit field.
func (o *ModelProductQuota) SetLimit(v float64) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *ModelProductQuota) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *ModelProductQuota) UnsetLimit() {
	o.Limit.Unset()
}

func (o ModelProductQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelProductQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Bought.IsSet() {
		toSerialize["bought"] = o.Bought.Get()
	}
	if o.CurrencyCode.IsSet() {
		toSerialize["currency_code"] = o.CurrencyCode.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelProductQuota) UnmarshalJSON(data []byte) (err error) {
	varModelProductQuota := _ModelProductQuota{}

	err = json.Unmarshal(data, &varModelProductQuota)

	if err != nil {
		return err
	}

	*o = ModelProductQuota(varModelProductQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bought")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelProductQuota struct {
	value *ModelProductQuota
	isSet bool
}

func (v NullableModelProductQuota) Get() *ModelProductQuota {
	return v.value
}

func (v *NullableModelProductQuota) Set(val *ModelProductQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableModelProductQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableModelProductQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelProductQuota(val *ModelProductQuota) *NullableModelProductQuota {
	return &NullableModelProductQuota{value: val, isSet: true}
}

func (v NullableModelProductQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelProductQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


