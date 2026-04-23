
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelCoinAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCoinAmount{}

// ModelCoinAmount struct for ModelCoinAmount
type ModelCoinAmount struct {
	Free NullableFloat64 `json:"free,omitempty"`
	Paid NullableFloat64 `json:"paid,omitempty"`
	Total NullableFloat64 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelCoinAmount ModelCoinAmount

// NewModelCoinAmount instantiates a new ModelCoinAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCoinAmount() *ModelCoinAmount {
	this := ModelCoinAmount{}
	return &this
}

// NewModelCoinAmountWithDefaults instantiates a new ModelCoinAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCoinAmountWithDefaults() *ModelCoinAmount {
	this := ModelCoinAmount{}
	return &this
}

// GetFree returns the Free field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCoinAmount) GetFree() float64 {
	if o == nil || IsNil(o.Free.Get()) {
		var ret float64
		return ret
	}
	return *o.Free.Get()
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCoinAmount) GetFreeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Free.Get(), o.Free.IsSet()
}

// HasFree returns a boolean if a field has been set.
func (o *ModelCoinAmount) HasFree() bool {
	if o != nil && o.Free.IsSet() {
		return true
	}

	return false
}

// SetFree gets a reference to the given NullableFloat64 and assigns it to the Free field.
func (o *ModelCoinAmount) SetFree(v float64) {
	o.Free.Set(&v)
}
// SetFreeNil sets the value for Free to be an explicit nil
func (o *ModelCoinAmount) SetFreeNil() {
	o.Free.Set(nil)
}

// UnsetFree ensures that no value is present for Free, not even an explicit nil
func (o *ModelCoinAmount) UnsetFree() {
	o.Free.Unset()
}

// GetPaid returns the Paid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCoinAmount) GetPaid() float64 {
	if o == nil || IsNil(o.Paid.Get()) {
		var ret float64
		return ret
	}
	return *o.Paid.Get()
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCoinAmount) GetPaidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paid.Get(), o.Paid.IsSet()
}

// HasPaid returns a boolean if a field has been set.
func (o *ModelCoinAmount) HasPaid() bool {
	if o != nil && o.Paid.IsSet() {
		return true
	}

	return false
}

// SetPaid gets a reference to the given NullableFloat64 and assigns it to the Paid field.
func (o *ModelCoinAmount) SetPaid(v float64) {
	o.Paid.Set(&v)
}
// SetPaidNil sets the value for Paid to be an explicit nil
func (o *ModelCoinAmount) SetPaidNil() {
	o.Paid.Set(nil)
}

// UnsetPaid ensures that no value is present for Paid, not even an explicit nil
func (o *ModelCoinAmount) UnsetPaid() {
	o.Paid.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCoinAmount) GetTotal() float64 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret float64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCoinAmount) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *ModelCoinAmount) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableFloat64 and assigns it to the Total field.
func (o *ModelCoinAmount) SetTotal(v float64) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *ModelCoinAmount) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *ModelCoinAmount) UnsetTotal() {
	o.Total.Unset()
}

func (o ModelCoinAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCoinAmount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Free.IsSet() {
		toSerialize["free"] = o.Free.Get()
	}
	if o.Paid.IsSet() {
		toSerialize["paid"] = o.Paid.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelCoinAmount) UnmarshalJSON(data []byte) (err error) {
	varModelCoinAmount := _ModelCoinAmount{}

	err = json.Unmarshal(data, &varModelCoinAmount)

	if err != nil {
		return err
	}

	*o = ModelCoinAmount(varModelCoinAmount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "free")
		delete(additionalProperties, "paid")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelCoinAmount struct {
	value *ModelCoinAmount
	isSet bool
}

func (v NullableModelCoinAmount) Get() *ModelCoinAmount {
	return v.value
}

func (v *NullableModelCoinAmount) Set(val *ModelCoinAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCoinAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCoinAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCoinAmount(val *ModelCoinAmount) *NullableModelCoinAmount {
	return &NullableModelCoinAmount{value: val, isSet: true}
}

func (v NullableModelCoinAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCoinAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


