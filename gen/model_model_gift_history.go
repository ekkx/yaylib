
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelGiftHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGiftHistory{}

// ModelGiftHistory struct for ModelGiftHistory
type ModelGiftHistory struct {
	GiftItem NullableGift `json:"gift_item,omitempty"`
	Gifts []ReceivedGift `json:"gifts,omitempty"`
	IsSenderAnonymous NullableBool `json:"is_sender_anonymous,omitempty"`
	Receiver NullableUser `json:"receiver,omitempty"`
	Sender NullableUser `json:"sender,omitempty"`
	TransactionAtSeconds NullableInt64 `json:"transaction_at_seconds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelGiftHistory ModelGiftHistory

// NewModelGiftHistory instantiates a new ModelGiftHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGiftHistory() *ModelGiftHistory {
	this := ModelGiftHistory{}
	return &this
}

// NewModelGiftHistoryWithDefaults instantiates a new ModelGiftHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGiftHistoryWithDefaults() *ModelGiftHistory {
	this := ModelGiftHistory{}
	return &this
}

// GetGiftItem returns the GiftItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGiftHistory) GetGiftItem() Gift {
	if o == nil || IsNil(o.GiftItem.Get()) {
		var ret Gift
		return ret
	}
	return *o.GiftItem.Get()
}

// GetGiftItemOk returns a tuple with the GiftItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGiftHistory) GetGiftItemOk() (*Gift, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftItem.Get(), o.GiftItem.IsSet()
}

// HasGiftItem returns a boolean if a field has been set.
func (o *ModelGiftHistory) HasGiftItem() bool {
	if o != nil && o.GiftItem.IsSet() {
		return true
	}

	return false
}

// SetGiftItem gets a reference to the given NullableGift and assigns it to the GiftItem field.
func (o *ModelGiftHistory) SetGiftItem(v Gift) {
	o.GiftItem.Set(&v)
}
// SetGiftItemNil sets the value for GiftItem to be an explicit nil
func (o *ModelGiftHistory) SetGiftItemNil() {
	o.GiftItem.Set(nil)
}

// UnsetGiftItem ensures that no value is present for GiftItem, not even an explicit nil
func (o *ModelGiftHistory) UnsetGiftItem() {
	o.GiftItem.Unset()
}

// GetGifts returns the Gifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGiftHistory) GetGifts() []ReceivedGift {
	if o == nil {
		var ret []ReceivedGift
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGiftHistory) GetGiftsOk() ([]ReceivedGift, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *ModelGiftHistory) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given []ReceivedGift and assigns it to the Gifts field.
func (o *ModelGiftHistory) SetGifts(v []ReceivedGift) {
	o.Gifts = v
}

// GetIsSenderAnonymous returns the IsSenderAnonymous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGiftHistory) GetIsSenderAnonymous() bool {
	if o == nil || IsNil(o.IsSenderAnonymous.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSenderAnonymous.Get()
}

// GetIsSenderAnonymousOk returns a tuple with the IsSenderAnonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGiftHistory) GetIsSenderAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSenderAnonymous.Get(), o.IsSenderAnonymous.IsSet()
}

// HasIsSenderAnonymous returns a boolean if a field has been set.
func (o *ModelGiftHistory) HasIsSenderAnonymous() bool {
	if o != nil && o.IsSenderAnonymous.IsSet() {
		return true
	}

	return false
}

// SetIsSenderAnonymous gets a reference to the given NullableBool and assigns it to the IsSenderAnonymous field.
func (o *ModelGiftHistory) SetIsSenderAnonymous(v bool) {
	o.IsSenderAnonymous.Set(&v)
}
// SetIsSenderAnonymousNil sets the value for IsSenderAnonymous to be an explicit nil
func (o *ModelGiftHistory) SetIsSenderAnonymousNil() {
	o.IsSenderAnonymous.Set(nil)
}

// UnsetIsSenderAnonymous ensures that no value is present for IsSenderAnonymous, not even an explicit nil
func (o *ModelGiftHistory) UnsetIsSenderAnonymous() {
	o.IsSenderAnonymous.Unset()
}

// GetReceiver returns the Receiver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGiftHistory) GetReceiver() User {
	if o == nil || IsNil(o.Receiver.Get()) {
		var ret User
		return ret
	}
	return *o.Receiver.Get()
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGiftHistory) GetReceiverOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Receiver.Get(), o.Receiver.IsSet()
}

// HasReceiver returns a boolean if a field has been set.
func (o *ModelGiftHistory) HasReceiver() bool {
	if o != nil && o.Receiver.IsSet() {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given NullableUser and assigns it to the Receiver field.
func (o *ModelGiftHistory) SetReceiver(v User) {
	o.Receiver.Set(&v)
}
// SetReceiverNil sets the value for Receiver to be an explicit nil
func (o *ModelGiftHistory) SetReceiverNil() {
	o.Receiver.Set(nil)
}

// UnsetReceiver ensures that no value is present for Receiver, not even an explicit nil
func (o *ModelGiftHistory) UnsetReceiver() {
	o.Receiver.Unset()
}

// GetSender returns the Sender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGiftHistory) GetSender() User {
	if o == nil || IsNil(o.Sender.Get()) {
		var ret User
		return ret
	}
	return *o.Sender.Get()
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGiftHistory) GetSenderOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sender.Get(), o.Sender.IsSet()
}

// HasSender returns a boolean if a field has been set.
func (o *ModelGiftHistory) HasSender() bool {
	if o != nil && o.Sender.IsSet() {
		return true
	}

	return false
}

// SetSender gets a reference to the given NullableUser and assigns it to the Sender field.
func (o *ModelGiftHistory) SetSender(v User) {
	o.Sender.Set(&v)
}
// SetSenderNil sets the value for Sender to be an explicit nil
func (o *ModelGiftHistory) SetSenderNil() {
	o.Sender.Set(nil)
}

// UnsetSender ensures that no value is present for Sender, not even an explicit nil
func (o *ModelGiftHistory) UnsetSender() {
	o.Sender.Unset()
}

// GetTransactionAtSeconds returns the TransactionAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGiftHistory) GetTransactionAtSeconds() int64 {
	if o == nil || IsNil(o.TransactionAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.TransactionAtSeconds.Get()
}

// GetTransactionAtSecondsOk returns a tuple with the TransactionAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGiftHistory) GetTransactionAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionAtSeconds.Get(), o.TransactionAtSeconds.IsSet()
}

// HasTransactionAtSeconds returns a boolean if a field has been set.
func (o *ModelGiftHistory) HasTransactionAtSeconds() bool {
	if o != nil && o.TransactionAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetTransactionAtSeconds gets a reference to the given NullableInt64 and assigns it to the TransactionAtSeconds field.
func (o *ModelGiftHistory) SetTransactionAtSeconds(v int64) {
	o.TransactionAtSeconds.Set(&v)
}
// SetTransactionAtSecondsNil sets the value for TransactionAtSeconds to be an explicit nil
func (o *ModelGiftHistory) SetTransactionAtSecondsNil() {
	o.TransactionAtSeconds.Set(nil)
}

// UnsetTransactionAtSeconds ensures that no value is present for TransactionAtSeconds, not even an explicit nil
func (o *ModelGiftHistory) UnsetTransactionAtSeconds() {
	o.TransactionAtSeconds.Unset()
}

func (o ModelGiftHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGiftHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GiftItem.IsSet() {
		toSerialize["gift_item"] = o.GiftItem.Get()
	}
	if o.Gifts != nil {
		toSerialize["gifts"] = o.Gifts
	}
	if o.IsSenderAnonymous.IsSet() {
		toSerialize["is_sender_anonymous"] = o.IsSenderAnonymous.Get()
	}
	if o.Receiver.IsSet() {
		toSerialize["receiver"] = o.Receiver.Get()
	}
	if o.Sender.IsSet() {
		toSerialize["sender"] = o.Sender.Get()
	}
	if o.TransactionAtSeconds.IsSet() {
		toSerialize["transaction_at_seconds"] = o.TransactionAtSeconds.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelGiftHistory) UnmarshalJSON(data []byte) (err error) {
	varModelGiftHistory := _ModelGiftHistory{}

	err = json.Unmarshal(data, &varModelGiftHistory)

	if err != nil {
		return err
	}

	*o = ModelGiftHistory(varModelGiftHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift_item")
		delete(additionalProperties, "gifts")
		delete(additionalProperties, "is_sender_anonymous")
		delete(additionalProperties, "receiver")
		delete(additionalProperties, "sender")
		delete(additionalProperties, "transaction_at_seconds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelGiftHistory struct {
	value *ModelGiftHistory
	isSet bool
}

func (v NullableModelGiftHistory) Get() *ModelGiftHistory {
	return v.value
}

func (v *NullableModelGiftHistory) Set(val *ModelGiftHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGiftHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGiftHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGiftHistory(val *ModelGiftHistory) *NullableModelGiftHistory {
	return &NullableModelGiftHistory{value: val, isSet: true}
}

func (v NullableModelGiftHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGiftHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


