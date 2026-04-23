
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftHistory{}

// GiftHistory struct for GiftHistory
type GiftHistory struct {
	Anonymous NullableBool `json:"anonymous,omitempty"`
	GiftItem NullableGiftSlugItem `json:"gift_item,omitempty"`
	GiftsCount []GiftCount `json:"gifts_count,omitempty"`
	Receiver NullableRealmUser `json:"receiver,omitempty"`
	Sender NullableRealmUser `json:"sender,omitempty"`
	SentAt NullableInt64 `json:"sent_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftHistory GiftHistory

// NewGiftHistory instantiates a new GiftHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftHistory() *GiftHistory {
	this := GiftHistory{}
	return &this
}

// NewGiftHistoryWithDefaults instantiates a new GiftHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftHistoryWithDefaults() *GiftHistory {
	this := GiftHistory{}
	return &this
}

// GetAnonymous returns the Anonymous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftHistory) GetAnonymous() bool {
	if o == nil || IsNil(o.Anonymous.Get()) {
		var ret bool
		return ret
	}
	return *o.Anonymous.Get()
}

// GetAnonymousOk returns a tuple with the Anonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftHistory) GetAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Anonymous.Get(), o.Anonymous.IsSet()
}

// HasAnonymous returns a boolean if a field has been set.
func (o *GiftHistory) HasAnonymous() bool {
	if o != nil && o.Anonymous.IsSet() {
		return true
	}

	return false
}

// SetAnonymous gets a reference to the given NullableBool and assigns it to the Anonymous field.
func (o *GiftHistory) SetAnonymous(v bool) {
	o.Anonymous.Set(&v)
}
// SetAnonymousNil sets the value for Anonymous to be an explicit nil
func (o *GiftHistory) SetAnonymousNil() {
	o.Anonymous.Set(nil)
}

// UnsetAnonymous ensures that no value is present for Anonymous, not even an explicit nil
func (o *GiftHistory) UnsetAnonymous() {
	o.Anonymous.Unset()
}

// GetGiftItem returns the GiftItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftHistory) GetGiftItem() GiftSlugItem {
	if o == nil || IsNil(o.GiftItem.Get()) {
		var ret GiftSlugItem
		return ret
	}
	return *o.GiftItem.Get()
}

// GetGiftItemOk returns a tuple with the GiftItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftHistory) GetGiftItemOk() (*GiftSlugItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftItem.Get(), o.GiftItem.IsSet()
}

// HasGiftItem returns a boolean if a field has been set.
func (o *GiftHistory) HasGiftItem() bool {
	if o != nil && o.GiftItem.IsSet() {
		return true
	}

	return false
}

// SetGiftItem gets a reference to the given NullableGiftSlugItem and assigns it to the GiftItem field.
func (o *GiftHistory) SetGiftItem(v GiftSlugItem) {
	o.GiftItem.Set(&v)
}
// SetGiftItemNil sets the value for GiftItem to be an explicit nil
func (o *GiftHistory) SetGiftItemNil() {
	o.GiftItem.Set(nil)
}

// UnsetGiftItem ensures that no value is present for GiftItem, not even an explicit nil
func (o *GiftHistory) UnsetGiftItem() {
	o.GiftItem.Unset()
}

// GetGiftsCount returns the GiftsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftHistory) GetGiftsCount() []GiftCount {
	if o == nil {
		var ret []GiftCount
		return ret
	}
	return o.GiftsCount
}

// GetGiftsCountOk returns a tuple with the GiftsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftHistory) GetGiftsCountOk() ([]GiftCount, bool) {
	if o == nil || IsNil(o.GiftsCount) {
		return nil, false
	}
	return o.GiftsCount, true
}

// HasGiftsCount returns a boolean if a field has been set.
func (o *GiftHistory) HasGiftsCount() bool {
	if o != nil && !IsNil(o.GiftsCount) {
		return true
	}

	return false
}

// SetGiftsCount gets a reference to the given []GiftCount and assigns it to the GiftsCount field.
func (o *GiftHistory) SetGiftsCount(v []GiftCount) {
	o.GiftsCount = v
}

// GetReceiver returns the Receiver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftHistory) GetReceiver() RealmUser {
	if o == nil || IsNil(o.Receiver.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.Receiver.Get()
}

// GetReceiverOk returns a tuple with the Receiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftHistory) GetReceiverOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Receiver.Get(), o.Receiver.IsSet()
}

// HasReceiver returns a boolean if a field has been set.
func (o *GiftHistory) HasReceiver() bool {
	if o != nil && o.Receiver.IsSet() {
		return true
	}

	return false
}

// SetReceiver gets a reference to the given NullableRealmUser and assigns it to the Receiver field.
func (o *GiftHistory) SetReceiver(v RealmUser) {
	o.Receiver.Set(&v)
}
// SetReceiverNil sets the value for Receiver to be an explicit nil
func (o *GiftHistory) SetReceiverNil() {
	o.Receiver.Set(nil)
}

// UnsetReceiver ensures that no value is present for Receiver, not even an explicit nil
func (o *GiftHistory) UnsetReceiver() {
	o.Receiver.Unset()
}

// GetSender returns the Sender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftHistory) GetSender() RealmUser {
	if o == nil || IsNil(o.Sender.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.Sender.Get()
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftHistory) GetSenderOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sender.Get(), o.Sender.IsSet()
}

// HasSender returns a boolean if a field has been set.
func (o *GiftHistory) HasSender() bool {
	if o != nil && o.Sender.IsSet() {
		return true
	}

	return false
}

// SetSender gets a reference to the given NullableRealmUser and assigns it to the Sender field.
func (o *GiftHistory) SetSender(v RealmUser) {
	o.Sender.Set(&v)
}
// SetSenderNil sets the value for Sender to be an explicit nil
func (o *GiftHistory) SetSenderNil() {
	o.Sender.Set(nil)
}

// UnsetSender ensures that no value is present for Sender, not even an explicit nil
func (o *GiftHistory) UnsetSender() {
	o.Sender.Unset()
}

// GetSentAt returns the SentAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftHistory) GetSentAt() int64 {
	if o == nil || IsNil(o.SentAt.Get()) {
		var ret int64
		return ret
	}
	return *o.SentAt.Get()
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftHistory) GetSentAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SentAt.Get(), o.SentAt.IsSet()
}

// HasSentAt returns a boolean if a field has been set.
func (o *GiftHistory) HasSentAt() bool {
	if o != nil && o.SentAt.IsSet() {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given NullableInt64 and assigns it to the SentAt field.
func (o *GiftHistory) SetSentAt(v int64) {
	o.SentAt.Set(&v)
}
// SetSentAtNil sets the value for SentAt to be an explicit nil
func (o *GiftHistory) SetSentAtNil() {
	o.SentAt.Set(nil)
}

// UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
func (o *GiftHistory) UnsetSentAt() {
	o.SentAt.Unset()
}

func (o GiftHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Anonymous.IsSet() {
		toSerialize["anonymous"] = o.Anonymous.Get()
	}
	if o.GiftItem.IsSet() {
		toSerialize["gift_item"] = o.GiftItem.Get()
	}
	if o.GiftsCount != nil {
		toSerialize["gifts_count"] = o.GiftsCount
	}
	if o.Receiver.IsSet() {
		toSerialize["receiver"] = o.Receiver.Get()
	}
	if o.Sender.IsSet() {
		toSerialize["sender"] = o.Sender.Get()
	}
	if o.SentAt.IsSet() {
		toSerialize["sent_at"] = o.SentAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftHistory) UnmarshalJSON(data []byte) (err error) {
	varGiftHistory := _GiftHistory{}

	err = json.Unmarshal(data, &varGiftHistory)

	if err != nil {
		return err
	}

	*o = GiftHistory(varGiftHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "anonymous")
		delete(additionalProperties, "gift_item")
		delete(additionalProperties, "gifts_count")
		delete(additionalProperties, "receiver")
		delete(additionalProperties, "sender")
		delete(additionalProperties, "sent_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftHistory struct {
	value *GiftHistory
	isSet bool
}

func (v NullableGiftHistory) Get() *GiftHistory {
	return v.value
}

func (v *NullableGiftHistory) Set(val *GiftHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftHistory(val *GiftHistory) *NullableGiftHistory {
	return &NullableGiftHistory{value: val, isSet: true}
}

func (v NullableGiftHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


