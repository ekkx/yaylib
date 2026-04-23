
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftRewardGift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftRewardGift{}

// GiftRewardGift struct for GiftRewardGift
type GiftRewardGift struct {
	Icon NullableString `json:"icon,omitempty"`
	ReceivedCount NullableInt32 `json:"received_count,omitempty"`
	RequiredCount NullableInt32 `json:"required_count,omitempty"`
	Title NullableString `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftRewardGift GiftRewardGift

// NewGiftRewardGift instantiates a new GiftRewardGift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftRewardGift() *GiftRewardGift {
	this := GiftRewardGift{}
	return &this
}

// NewGiftRewardGiftWithDefaults instantiates a new GiftRewardGift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftRewardGiftWithDefaults() *GiftRewardGift {
	this := GiftRewardGift{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewardGift) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewardGift) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GiftRewardGift) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GiftRewardGift) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *GiftRewardGift) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GiftRewardGift) UnsetIcon() {
	o.Icon.Unset()
}

// GetReceivedCount returns the ReceivedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewardGift) GetReceivedCount() int32 {
	if o == nil || IsNil(o.ReceivedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReceivedCount.Get()
}

// GetReceivedCountOk returns a tuple with the ReceivedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewardGift) GetReceivedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivedCount.Get(), o.ReceivedCount.IsSet()
}

// HasReceivedCount returns a boolean if a field has been set.
func (o *GiftRewardGift) HasReceivedCount() bool {
	if o != nil && o.ReceivedCount.IsSet() {
		return true
	}

	return false
}

// SetReceivedCount gets a reference to the given NullableInt32 and assigns it to the ReceivedCount field.
func (o *GiftRewardGift) SetReceivedCount(v int32) {
	o.ReceivedCount.Set(&v)
}
// SetReceivedCountNil sets the value for ReceivedCount to be an explicit nil
func (o *GiftRewardGift) SetReceivedCountNil() {
	o.ReceivedCount.Set(nil)
}

// UnsetReceivedCount ensures that no value is present for ReceivedCount, not even an explicit nil
func (o *GiftRewardGift) UnsetReceivedCount() {
	o.ReceivedCount.Unset()
}

// GetRequiredCount returns the RequiredCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewardGift) GetRequiredCount() int32 {
	if o == nil || IsNil(o.RequiredCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RequiredCount.Get()
}

// GetRequiredCountOk returns a tuple with the RequiredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewardGift) GetRequiredCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredCount.Get(), o.RequiredCount.IsSet()
}

// HasRequiredCount returns a boolean if a field has been set.
func (o *GiftRewardGift) HasRequiredCount() bool {
	if o != nil && o.RequiredCount.IsSet() {
		return true
	}

	return false
}

// SetRequiredCount gets a reference to the given NullableInt32 and assigns it to the RequiredCount field.
func (o *GiftRewardGift) SetRequiredCount(v int32) {
	o.RequiredCount.Set(&v)
}
// SetRequiredCountNil sets the value for RequiredCount to be an explicit nil
func (o *GiftRewardGift) SetRequiredCountNil() {
	o.RequiredCount.Set(nil)
}

// UnsetRequiredCount ensures that no value is present for RequiredCount, not even an explicit nil
func (o *GiftRewardGift) UnsetRequiredCount() {
	o.RequiredCount.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftRewardGift) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftRewardGift) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *GiftRewardGift) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *GiftRewardGift) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *GiftRewardGift) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *GiftRewardGift) UnsetTitle() {
	o.Title.Unset()
}

func (o GiftRewardGift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftRewardGift) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.ReceivedCount.IsSet() {
		toSerialize["received_count"] = o.ReceivedCount.Get()
	}
	if o.RequiredCount.IsSet() {
		toSerialize["required_count"] = o.RequiredCount.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftRewardGift) UnmarshalJSON(data []byte) (err error) {
	varGiftRewardGift := _GiftRewardGift{}

	err = json.Unmarshal(data, &varGiftRewardGift)

	if err != nil {
		return err
	}

	*o = GiftRewardGift(varGiftRewardGift)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon")
		delete(additionalProperties, "received_count")
		delete(additionalProperties, "required_count")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftRewardGift struct {
	value *GiftRewardGift
	isSet bool
}

func (v NullableGiftRewardGift) Get() *GiftRewardGift {
	return v.value
}

func (v *NullableGiftRewardGift) Set(val *GiftRewardGift) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftRewardGift) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftRewardGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftRewardGift(val *GiftRewardGift) *NullableGiftRewardGift {
	return &NullableGiftRewardGift{value: val, isSet: true}
}

func (v NullableGiftRewardGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftRewardGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


