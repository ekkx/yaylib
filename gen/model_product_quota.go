
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ProductQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductQuota{}

// ProductQuota struct for ProductQuota
type ProductQuota struct {
	Bought NullableFloat64 `json:"bought,omitempty"`
	Limit NullableFloat64 `json:"limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductQuota ProductQuota

// NewProductQuota instantiates a new ProductQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductQuota() *ProductQuota {
	this := ProductQuota{}
	return &this
}

// NewProductQuotaWithDefaults instantiates a new ProductQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductQuotaWithDefaults() *ProductQuota {
	this := ProductQuota{}
	return &this
}

// GetBought returns the Bought field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQuota) GetBought() float64 {
	if o == nil || IsNil(o.Bought.Get()) {
		var ret float64
		return ret
	}
	return *o.Bought.Get()
}

// GetBoughtOk returns a tuple with the Bought field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQuota) GetBoughtOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bought.Get(), o.Bought.IsSet()
}

// HasBought returns a boolean if a field has been set.
func (o *ProductQuota) HasBought() bool {
	if o != nil && o.Bought.IsSet() {
		return true
	}

	return false
}

// SetBought gets a reference to the given NullableFloat64 and assigns it to the Bought field.
func (o *ProductQuota) SetBought(v float64) {
	o.Bought.Set(&v)
}
// SetBoughtNil sets the value for Bought to be an explicit nil
func (o *ProductQuota) SetBoughtNil() {
	o.Bought.Set(nil)
}

// UnsetBought ensures that no value is present for Bought, not even an explicit nil
func (o *ProductQuota) UnsetBought() {
	o.Bought.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductQuota) GetLimit() float64 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret float64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductQuota) GetLimitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *ProductQuota) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableFloat64 and assigns it to the Limit field.
func (o *ProductQuota) SetLimit(v float64) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *ProductQuota) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *ProductQuota) UnsetLimit() {
	o.Limit.Unset()
}

func (o ProductQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Bought.IsSet() {
		toSerialize["bought"] = o.Bought.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductQuota) UnmarshalJSON(data []byte) (err error) {
	varProductQuota := _ProductQuota{}

	err = json.Unmarshal(data, &varProductQuota)

	if err != nil {
		return err
	}

	*o = ProductQuota(varProductQuota)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bought")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductQuota struct {
	value *ProductQuota
	isSet bool
}

func (v NullableProductQuota) Get() *ProductQuota {
	return v.value
}

func (v *NullableProductQuota) Set(val *ProductQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableProductQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableProductQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductQuota(val *ProductQuota) *NullableProductQuota {
	return &NullableProductQuota{value: val, isSet: true}
}

func (v NullableProductQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


