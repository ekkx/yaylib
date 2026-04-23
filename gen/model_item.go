
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Item type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item{}

// Item struct for Item
type Item struct {
	Gift NullableRealmGift `json:"gift,omitempty"`
	Quantity NullableInt32 `json:"quantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Item Item

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem() *Item {
	this := Item{}
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	return &this
}

// GetGift returns the Gift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Item) GetGift() RealmGift {
	if o == nil || IsNil(o.Gift.Get()) {
		var ret RealmGift
		return ret
	}
	return *o.Gift.Get()
}

// GetGiftOk returns a tuple with the Gift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Item) GetGiftOk() (*RealmGift, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gift.Get(), o.Gift.IsSet()
}

// HasGift returns a boolean if a field has been set.
func (o *Item) HasGift() bool {
	if o != nil && o.Gift.IsSet() {
		return true
	}

	return false
}

// SetGift gets a reference to the given NullableRealmGift and assigns it to the Gift field.
func (o *Item) SetGift(v RealmGift) {
	o.Gift.Set(&v)
}
// SetGiftNil sets the value for Gift to be an explicit nil
func (o *Item) SetGiftNil() {
	o.Gift.Set(nil)
}

// UnsetGift ensures that no value is present for Gift, not even an explicit nil
func (o *Item) UnsetGift() {
	o.Gift.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Item) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Item) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *Item) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt32 and assigns it to the Quantity field.
func (o *Item) SetQuantity(v int32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *Item) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *Item) UnsetQuantity() {
	o.Quantity.Unset()
}

func (o Item) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Item) ToMap() (map[string]interface{}, error) {
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

func (o *Item) UnmarshalJSON(data []byte) (err error) {
	varItem := _Item{}

	err = json.Unmarshal(data, &varItem)

	if err != nil {
		return err
	}

	*o = Item(varItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift")
		delete(additionalProperties, "quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


