
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftTransaction{}

// GiftTransaction struct for GiftTransaction
type GiftTransaction struct {
	Gift NullableGift `json:"gift,omitempty"`
	Quantity NullableInt32 `json:"quantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftTransaction GiftTransaction

// NewGiftTransaction instantiates a new GiftTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftTransaction() *GiftTransaction {
	this := GiftTransaction{}
	return &this
}

// NewGiftTransactionWithDefaults instantiates a new GiftTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftTransactionWithDefaults() *GiftTransaction {
	this := GiftTransaction{}
	return &this
}

// GetGift returns the Gift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftTransaction) GetGift() Gift {
	if o == nil || IsNil(o.Gift.Get()) {
		var ret Gift
		return ret
	}
	return *o.Gift.Get()
}

// GetGiftOk returns a tuple with the Gift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftTransaction) GetGiftOk() (*Gift, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gift.Get(), o.Gift.IsSet()
}

// HasGift returns a boolean if a field has been set.
func (o *GiftTransaction) HasGift() bool {
	if o != nil && o.Gift.IsSet() {
		return true
	}

	return false
}

// SetGift gets a reference to the given NullableGift and assigns it to the Gift field.
func (o *GiftTransaction) SetGift(v Gift) {
	o.Gift.Set(&v)
}
// SetGiftNil sets the value for Gift to be an explicit nil
func (o *GiftTransaction) SetGiftNil() {
	o.Gift.Set(nil)
}

// UnsetGift ensures that no value is present for Gift, not even an explicit nil
func (o *GiftTransaction) UnsetGift() {
	o.Gift.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftTransaction) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftTransaction) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *GiftTransaction) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt32 and assigns it to the Quantity field.
func (o *GiftTransaction) SetQuantity(v int32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *GiftTransaction) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *GiftTransaction) UnsetQuantity() {
	o.Quantity.Unset()
}

func (o GiftTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gift.IsSet() {
		toSerialize["gift"] = o.Gift.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftTransaction) UnmarshalJSON(data []byte) (err error) {
	varGiftTransaction := _GiftTransaction{}

	err = json.Unmarshal(data, &varGiftTransaction)

	if err != nil {
		return err
	}

	*o = GiftTransaction(varGiftTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift")
		delete(additionalProperties, "quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftTransaction struct {
	value *GiftTransaction
	isSet bool
}

func (v NullableGiftTransaction) Get() *GiftTransaction {
	return v.value
}

func (v *NullableGiftTransaction) Set(val *GiftTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftTransaction(val *GiftTransaction) *NullableGiftTransaction {
	return &NullableGiftTransaction{value: val, isSet: true}
}

func (v NullableGiftTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


