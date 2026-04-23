
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupGiftHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupGiftHistory{}

// GroupGiftHistory struct for GroupGiftHistory
type GroupGiftHistory struct {
	GiftsCount []GiftCount `json:"gifts_count,omitempty"`
	ReceivedDate NullableInt64 `json:"received_date,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupGiftHistory GroupGiftHistory

// NewGroupGiftHistory instantiates a new GroupGiftHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupGiftHistory() *GroupGiftHistory {
	this := GroupGiftHistory{}
	return &this
}

// NewGroupGiftHistoryWithDefaults instantiates a new GroupGiftHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupGiftHistoryWithDefaults() *GroupGiftHistory {
	this := GroupGiftHistory{}
	return &this
}

// GetGiftsCount returns the GiftsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGiftHistory) GetGiftsCount() []GiftCount {
	if o == nil {
		var ret []GiftCount
		return ret
	}
	return o.GiftsCount
}

// GetGiftsCountOk returns a tuple with the GiftsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGiftHistory) GetGiftsCountOk() ([]GiftCount, bool) {
	if o == nil || IsNil(o.GiftsCount) {
		return nil, false
	}
	return o.GiftsCount, true
}

// HasGiftsCount returns a boolean if a field has been set.
func (o *GroupGiftHistory) HasGiftsCount() bool {
	if o != nil && !IsNil(o.GiftsCount) {
		return true
	}

	return false
}

// SetGiftsCount gets a reference to the given []GiftCount and assigns it to the GiftsCount field.
func (o *GroupGiftHistory) SetGiftsCount(v []GiftCount) {
	o.GiftsCount = v
}

// GetReceivedDate returns the ReceivedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGiftHistory) GetReceivedDate() int64 {
	if o == nil || IsNil(o.ReceivedDate.Get()) {
		var ret int64
		return ret
	}
	return *o.ReceivedDate.Get()
}

// GetReceivedDateOk returns a tuple with the ReceivedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGiftHistory) GetReceivedDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceivedDate.Get(), o.ReceivedDate.IsSet()
}

// HasReceivedDate returns a boolean if a field has been set.
func (o *GroupGiftHistory) HasReceivedDate() bool {
	if o != nil && o.ReceivedDate.IsSet() {
		return true
	}

	return false
}

// SetReceivedDate gets a reference to the given NullableInt64 and assigns it to the ReceivedDate field.
func (o *GroupGiftHistory) SetReceivedDate(v int64) {
	o.ReceivedDate.Set(&v)
}
// SetReceivedDateNil sets the value for ReceivedDate to be an explicit nil
func (o *GroupGiftHistory) SetReceivedDateNil() {
	o.ReceivedDate.Set(nil)
}

// UnsetReceivedDate ensures that no value is present for ReceivedDate, not even an explicit nil
func (o *GroupGiftHistory) UnsetReceivedDate() {
	o.ReceivedDate.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGiftHistory) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGiftHistory) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *GroupGiftHistory) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *GroupGiftHistory) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *GroupGiftHistory) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *GroupGiftHistory) UnsetUser() {
	o.User.Unset()
}

func (o GroupGiftHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupGiftHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GiftsCount != nil {
		toSerialize["gifts_count"] = o.GiftsCount
	}
	if o.ReceivedDate.IsSet() {
		toSerialize["received_date"] = o.ReceivedDate.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupGiftHistory) UnmarshalJSON(data []byte) (err error) {
	varGroupGiftHistory := _GroupGiftHistory{}

	err = json.Unmarshal(data, &varGroupGiftHistory)

	if err != nil {
		return err
	}

	*o = GroupGiftHistory(varGroupGiftHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gifts_count")
		delete(additionalProperties, "received_date")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupGiftHistory struct {
	value *GroupGiftHistory
	isSet bool
}

func (v NullableGroupGiftHistory) Get() *GroupGiftHistory {
	return v.value
}

func (v *NullableGroupGiftHistory) Set(val *GroupGiftHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupGiftHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupGiftHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupGiftHistory(val *GroupGiftHistory) *NullableGroupGiftHistory {
	return &NullableGroupGiftHistory{value: val, isSet: true}
}

func (v NullableGroupGiftHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupGiftHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


