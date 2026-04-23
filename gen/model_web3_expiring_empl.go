
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3ExpiringEmpl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3ExpiringEmpl{}

// Web3ExpiringEmpl struct for Web3ExpiringEmpl
type Web3ExpiringEmpl struct {
	Items []Web3ExpiringEmplExpiredEmpl `json:"items,omitempty"`
	TotalAmount NullableFloat64 `json:"total_amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3ExpiringEmpl Web3ExpiringEmpl

// NewWeb3ExpiringEmpl instantiates a new Web3ExpiringEmpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3ExpiringEmpl() *Web3ExpiringEmpl {
	this := Web3ExpiringEmpl{}
	return &this
}

// NewWeb3ExpiringEmplWithDefaults instantiates a new Web3ExpiringEmpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3ExpiringEmplWithDefaults() *Web3ExpiringEmpl {
	this := Web3ExpiringEmpl{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3ExpiringEmpl) GetItems() []Web3ExpiringEmplExpiredEmpl {
	if o == nil {
		var ret []Web3ExpiringEmplExpiredEmpl
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3ExpiringEmpl) GetItemsOk() ([]Web3ExpiringEmplExpiredEmpl, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Web3ExpiringEmpl) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Web3ExpiringEmplExpiredEmpl and assigns it to the Items field.
func (o *Web3ExpiringEmpl) SetItems(v []Web3ExpiringEmplExpiredEmpl) {
	o.Items = v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3ExpiringEmpl) GetTotalAmount() float64 {
	if o == nil || IsNil(o.TotalAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalAmount.Get()
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3ExpiringEmpl) GetTotalAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalAmount.Get(), o.TotalAmount.IsSet()
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *Web3ExpiringEmpl) HasTotalAmount() bool {
	if o != nil && o.TotalAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given NullableFloat64 and assigns it to the TotalAmount field.
func (o *Web3ExpiringEmpl) SetTotalAmount(v float64) {
	o.TotalAmount.Set(&v)
}
// SetTotalAmountNil sets the value for TotalAmount to be an explicit nil
func (o *Web3ExpiringEmpl) SetTotalAmountNil() {
	o.TotalAmount.Set(nil)
}

// UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
func (o *Web3ExpiringEmpl) UnsetTotalAmount() {
	o.TotalAmount.Unset()
}

func (o Web3ExpiringEmpl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3ExpiringEmpl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.TotalAmount.IsSet() {
		toSerialize["total_amount"] = o.TotalAmount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3ExpiringEmpl) UnmarshalJSON(data []byte) (err error) {
	varWeb3ExpiringEmpl := _Web3ExpiringEmpl{}

	err = json.Unmarshal(data, &varWeb3ExpiringEmpl)

	if err != nil {
		return err
	}

	*o = Web3ExpiringEmpl(varWeb3ExpiringEmpl)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "total_amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3ExpiringEmpl struct {
	value *Web3ExpiringEmpl
	isSet bool
}

func (v NullableWeb3ExpiringEmpl) Get() *Web3ExpiringEmpl {
	return v.value
}

func (v *NullableWeb3ExpiringEmpl) Set(val *Web3ExpiringEmpl) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3ExpiringEmpl) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3ExpiringEmpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3ExpiringEmpl(val *Web3ExpiringEmpl) *NullableWeb3ExpiringEmpl {
	return &NullableWeb3ExpiringEmpl{value: val, isSet: true}
}

func (v NullableWeb3ExpiringEmpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3ExpiringEmpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


