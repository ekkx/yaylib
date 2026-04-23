
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ReceivedGift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceivedGift{}

// ReceivedGift struct for ReceivedGift
type ReceivedGift struct {
	Gift NullableGift `json:"gift,omitempty"`
	ReceivedCount NullableInt32 `json:"received_count,omitempty"`
	Senders []User `json:"senders,omitempty"`
	TotalSendersCount NullableInt32 `json:"total_senders_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReceivedGift ReceivedGift

// NewReceivedGift instantiates a new ReceivedGift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivedGift() *ReceivedGift {
	this := ReceivedGift{}
	return &this
}

// NewReceivedGiftWithDefaults instantiates a new ReceivedGift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceivedGiftWithDefaults() *ReceivedGift {
	this := ReceivedGift{}
	return &this
}

// GetGift returns the Gift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedGift) GetGift() Gift {
	if o == nil || IsNil(o.Gift.Get()) {
		var ret Gift
		return ret
	}
	return *o.Gift.Get()
}

// GetGiftOk returns a tuple with the Gift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedGift) GetGiftOk() (*Gift, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gift.Get(), o.Gift.IsSet()
}

// HasGift returns a boolean if a field has been set.
func (o *ReceivedGift) HasGift() bool {
	if o != nil && o.Gift.IsSet() {
		return true
	}

	return false
}

// SetGift gets a reference to the given NullableGift and assigns it to the Gift field.
func (o *ReceivedGift) SetGift(v Gift) {
	o.Gift.Set(&v)
}
// SetGiftNil sets the value for Gift to be an explicit nil
func (o *ReceivedGift) SetGiftNil() {
	o.Gift.Set(nil)
}

// UnsetGift ensures that no value is present for Gift, not even an explicit nil
func (o *ReceivedGift) UnsetGift() {
	o.Gift.Unset()
}

// GetReceivedCount returns the ReceivedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedGift) GetReceivedCount() int32 {
	if o == nil || IsNil(o.ReceivedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ReceivedCount.Get()
}

// GetReceivedCountOk returns a tuple with the ReceivedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedGift) GetReceivedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivedCount.Get(), o.ReceivedCount.IsSet()
}

// HasReceivedCount returns a boolean if a field has been set.
func (o *ReceivedGift) HasReceivedCount() bool {
	if o != nil && o.ReceivedCount.IsSet() {
		return true
	}

	return false
}

// SetReceivedCount gets a reference to the given NullableInt32 and assigns it to the ReceivedCount field.
func (o *ReceivedGift) SetReceivedCount(v int32) {
	o.ReceivedCount.Set(&v)
}
// SetReceivedCountNil sets the value for ReceivedCount to be an explicit nil
func (o *ReceivedGift) SetReceivedCountNil() {
	o.ReceivedCount.Set(nil)
}

// UnsetReceivedCount ensures that no value is present for ReceivedCount, not even an explicit nil
func (o *ReceivedGift) UnsetReceivedCount() {
	o.ReceivedCount.Unset()
}

// GetSenders returns the Senders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedGift) GetSenders() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedGift) GetSendersOk() ([]User, bool) {
	if o == nil || IsNil(o.Senders) {
		return nil, false
	}
	return o.Senders, true
}

// HasSenders returns a boolean if a field has been set.
func (o *ReceivedGift) HasSenders() bool {
	if o != nil && !IsNil(o.Senders) {
		return true
	}

	return false
}

// SetSenders gets a reference to the given []User and assigns it to the Senders field.
func (o *ReceivedGift) SetSenders(v []User) {
	o.Senders = v
}

// GetTotalSendersCount returns the TotalSendersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceivedGift) GetTotalSendersCount() int32 {
	if o == nil || IsNil(o.TotalSendersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalSendersCount.Get()
}

// GetTotalSendersCountOk returns a tuple with the TotalSendersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceivedGift) GetTotalSendersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalSendersCount.Get(), o.TotalSendersCount.IsSet()
}

// HasTotalSendersCount returns a boolean if a field has been set.
func (o *ReceivedGift) HasTotalSendersCount() bool {
	if o != nil && o.TotalSendersCount.IsSet() {
		return true
	}

	return false
}

// SetTotalSendersCount gets a reference to the given NullableInt32 and assigns it to the TotalSendersCount field.
func (o *ReceivedGift) SetTotalSendersCount(v int32) {
	o.TotalSendersCount.Set(&v)
}
// SetTotalSendersCountNil sets the value for TotalSendersCount to be an explicit nil
func (o *ReceivedGift) SetTotalSendersCountNil() {
	o.TotalSendersCount.Set(nil)
}

// UnsetTotalSendersCount ensures that no value is present for TotalSendersCount, not even an explicit nil
func (o *ReceivedGift) UnsetTotalSendersCount() {
	o.TotalSendersCount.Unset()
}

func (o ReceivedGift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceivedGift) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gift.IsSet() {
		toSerialize["gift"] = o.Gift.Get()
	}
	if o.ReceivedCount.IsSet() {
		toSerialize["received_count"] = o.ReceivedCount.Get()
	}
	if o.Senders != nil {
		toSerialize["senders"] = o.Senders
	}
	if o.TotalSendersCount.IsSet() {
		toSerialize["total_senders_count"] = o.TotalSendersCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReceivedGift) UnmarshalJSON(data []byte) (err error) {
	varReceivedGift := _ReceivedGift{}

	err = json.Unmarshal(data, &varReceivedGift)

	if err != nil {
		return err
	}

	*o = ReceivedGift(varReceivedGift)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift")
		delete(additionalProperties, "received_count")
		delete(additionalProperties, "senders")
		delete(additionalProperties, "total_senders_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReceivedGift struct {
	value *ReceivedGift
	isSet bool
}

func (v NullableReceivedGift) Get() *ReceivedGift {
	return v.value
}

func (v *NullableReceivedGift) Set(val *ReceivedGift) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivedGift) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivedGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivedGift(val *ReceivedGift) *NullableReceivedGift {
	return &NullableReceivedGift{value: val, isSet: true}
}

func (v NullableReceivedGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivedGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


