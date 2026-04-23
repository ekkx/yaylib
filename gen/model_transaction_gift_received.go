
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TransactionGiftReceived type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionGiftReceived{}

// TransactionGiftReceived struct for TransactionGiftReceived
type TransactionGiftReceived struct {
	Icon NullableString `json:"icon,omitempty"`
	Item NullableTransactionGiftReceivedItem `json:"item,omitempty"`
	ItemId NullableInt64 `json:"item_id,omitempty"`
	Quantity NullableInt32 `json:"quantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionGiftReceived TransactionGiftReceived

// NewTransactionGiftReceived instantiates a new TransactionGiftReceived object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionGiftReceived() *TransactionGiftReceived {
	this := TransactionGiftReceived{}
	return &this
}

// NewTransactionGiftReceivedWithDefaults instantiates a new TransactionGiftReceived object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionGiftReceivedWithDefaults() *TransactionGiftReceived {
	this := TransactionGiftReceived{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceived) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceived) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *TransactionGiftReceived) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *TransactionGiftReceived) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *TransactionGiftReceived) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *TransactionGiftReceived) UnsetIcon() {
	o.Icon.Unset()
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceived) GetItem() TransactionGiftReceivedItem {
	if o == nil || IsNil(o.Item.Get()) {
		var ret TransactionGiftReceivedItem
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceived) GetItemOk() (*TransactionGiftReceivedItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *TransactionGiftReceived) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableTransactionGiftReceivedItem and assigns it to the Item field.
func (o *TransactionGiftReceived) SetItem(v TransactionGiftReceivedItem) {
	o.Item.Set(&v)
}
// SetItemNil sets the value for Item to be an explicit nil
func (o *TransactionGiftReceived) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *TransactionGiftReceived) UnsetItem() {
	o.Item.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceived) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret int64
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceived) GetItemIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *TransactionGiftReceived) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableInt64 and assigns it to the ItemId field.
func (o *TransactionGiftReceived) SetItemId(v int64) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *TransactionGiftReceived) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *TransactionGiftReceived) UnsetItemId() {
	o.ItemId.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionGiftReceived) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionGiftReceived) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *TransactionGiftReceived) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt32 and assigns it to the Quantity field.
func (o *TransactionGiftReceived) SetQuantity(v int32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *TransactionGiftReceived) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *TransactionGiftReceived) UnsetQuantity() {
	o.Quantity.Unset()
}

func (o TransactionGiftReceived) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionGiftReceived) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Item.IsSet() {
		toSerialize["item"] = o.Item.Get()
	}
	if o.ItemId.IsSet() {
		toSerialize["item_id"] = o.ItemId.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransactionGiftReceived) UnmarshalJSON(data []byte) (err error) {
	varTransactionGiftReceived := _TransactionGiftReceived{}

	err = json.Unmarshal(data, &varTransactionGiftReceived)

	if err != nil {
		return err
	}

	*o = TransactionGiftReceived(varTransactionGiftReceived)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon")
		delete(additionalProperties, "item")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "quantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionGiftReceived struct {
	value *TransactionGiftReceived
	isSet bool
}

func (v NullableTransactionGiftReceived) Get() *TransactionGiftReceived {
	return v.value
}

func (v *NullableTransactionGiftReceived) Set(val *TransactionGiftReceived) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionGiftReceived) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionGiftReceived) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionGiftReceived(val *TransactionGiftReceived) *NullableTransactionGiftReceived {
	return &NullableTransactionGiftReceived{value: val, isSet: true}
}

func (v NullableTransactionGiftReceived) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionGiftReceived) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


