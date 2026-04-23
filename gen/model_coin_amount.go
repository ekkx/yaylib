
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CoinAmount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoinAmount{}

// CoinAmount struct for CoinAmount
type CoinAmount struct {
	Free NullableFloat64 `json:"free,omitempty"`
	Paid NullableFloat64 `json:"paid,omitempty"`
	Total NullableFloat64 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CoinAmount CoinAmount

// NewCoinAmount instantiates a new CoinAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoinAmount() *CoinAmount {
	this := CoinAmount{}
	return &this
}

// NewCoinAmountWithDefaults instantiates a new CoinAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoinAmountWithDefaults() *CoinAmount {
	this := CoinAmount{}
	return &this
}

// GetFree returns the Free field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoinAmount) GetFree() float64 {
	if o == nil || IsNil(o.Free.Get()) {
		var ret float64
		return ret
	}
	return *o.Free.Get()
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoinAmount) GetFreeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Free.Get(), o.Free.IsSet()
}

// HasFree returns a boolean if a field has been set.
func (o *CoinAmount) HasFree() bool {
	if o != nil && o.Free.IsSet() {
		return true
	}

	return false
}

// SetFree gets a reference to the given NullableFloat64 and assigns it to the Free field.
func (o *CoinAmount) SetFree(v float64) {
	o.Free.Set(&v)
}
// SetFreeNil sets the value for Free to be an explicit nil
func (o *CoinAmount) SetFreeNil() {
	o.Free.Set(nil)
}

// UnsetFree ensures that no value is present for Free, not even an explicit nil
func (o *CoinAmount) UnsetFree() {
	o.Free.Unset()
}

// GetPaid returns the Paid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoinAmount) GetPaid() float64 {
	if o == nil || IsNil(o.Paid.Get()) {
		var ret float64
		return ret
	}
	return *o.Paid.Get()
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoinAmount) GetPaidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Paid.Get(), o.Paid.IsSet()
}

// HasPaid returns a boolean if a field has been set.
func (o *CoinAmount) HasPaid() bool {
	if o != nil && o.Paid.IsSet() {
		return true
	}

	return false
}

// SetPaid gets a reference to the given NullableFloat64 and assigns it to the Paid field.
func (o *CoinAmount) SetPaid(v float64) {
	o.Paid.Set(&v)
}
// SetPaidNil sets the value for Paid to be an explicit nil
func (o *CoinAmount) SetPaidNil() {
	o.Paid.Set(nil)
}

// UnsetPaid ensures that no value is present for Paid, not even an explicit nil
func (o *CoinAmount) UnsetPaid() {
	o.Paid.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoinAmount) GetTotal() float64 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret float64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoinAmount) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *CoinAmount) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableFloat64 and assigns it to the Total field.
func (o *CoinAmount) SetTotal(v float64) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *CoinAmount) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *CoinAmount) UnsetTotal() {
	o.Total.Unset()
}

func (o CoinAmount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoinAmount) ToMap() (map[string]interface{}, error) {
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

func (o *CoinAmount) UnmarshalJSON(data []byte) (err error) {
	varCoinAmount := _CoinAmount{}

	err = json.Unmarshal(data, &varCoinAmount)

	if err != nil {
		return err
	}

	*o = CoinAmount(varCoinAmount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "free")
		delete(additionalProperties, "paid")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoinAmount struct {
	value *CoinAmount
	isSet bool
}

func (v NullableCoinAmount) Get() *CoinAmount {
	return v.value
}

func (v *NullableCoinAmount) Set(val *CoinAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableCoinAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableCoinAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoinAmount(val *CoinAmount) *NullableCoinAmount {
	return &NullableCoinAmount{value: val, isSet: true}
}

func (v NullableCoinAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoinAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


