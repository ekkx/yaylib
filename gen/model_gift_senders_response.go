
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftSendersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftSendersResponse{}

// GiftSendersResponse struct for GiftSendersResponse
type GiftSendersResponse struct {
	Senders []RealmUser `json:"senders,omitempty"`
	TotalSendersCount NullableInt32 `json:"total_senders_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftSendersResponse GiftSendersResponse

// NewGiftSendersResponse instantiates a new GiftSendersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftSendersResponse() *GiftSendersResponse {
	this := GiftSendersResponse{}
	return &this
}

// NewGiftSendersResponseWithDefaults instantiates a new GiftSendersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftSendersResponseWithDefaults() *GiftSendersResponse {
	this := GiftSendersResponse{}
	return &this
}

// GetSenders returns the Senders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftSendersResponse) GetSenders() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftSendersResponse) GetSendersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Senders) {
		return nil, false
	}
	return o.Senders, true
}

// HasSenders returns a boolean if a field has been set.
func (o *GiftSendersResponse) HasSenders() bool {
	if o != nil && !IsNil(o.Senders) {
		return true
	}

	return false
}

// SetSenders gets a reference to the given []RealmUser and assigns it to the Senders field.
func (o *GiftSendersResponse) SetSenders(v []RealmUser) {
	o.Senders = v
}

// GetTotalSendersCount returns the TotalSendersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftSendersResponse) GetTotalSendersCount() int32 {
	if o == nil || IsNil(o.TotalSendersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalSendersCount.Get()
}

// GetTotalSendersCountOk returns a tuple with the TotalSendersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftSendersResponse) GetTotalSendersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalSendersCount.Get(), o.TotalSendersCount.IsSet()
}

// HasTotalSendersCount returns a boolean if a field has been set.
func (o *GiftSendersResponse) HasTotalSendersCount() bool {
	if o != nil && o.TotalSendersCount.IsSet() {
		return true
	}

	return false
}

// SetTotalSendersCount gets a reference to the given NullableInt32 and assigns it to the TotalSendersCount field.
func (o *GiftSendersResponse) SetTotalSendersCount(v int32) {
	o.TotalSendersCount.Set(&v)
}
// SetTotalSendersCountNil sets the value for TotalSendersCount to be an explicit nil
func (o *GiftSendersResponse) SetTotalSendersCountNil() {
	o.TotalSendersCount.Set(nil)
}

// UnsetTotalSendersCount ensures that no value is present for TotalSendersCount, not even an explicit nil
func (o *GiftSendersResponse) UnsetTotalSendersCount() {
	o.TotalSendersCount.Unset()
}

func (o GiftSendersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftSendersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *GiftSendersResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftSendersResponse := _GiftSendersResponse{}

	err = json.Unmarshal(data, &varGiftSendersResponse)

	if err != nil {
		return err
	}

	*o = GiftSendersResponse(varGiftSendersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "senders")
		delete(additionalProperties, "total_senders_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftSendersResponse struct {
	value *GiftSendersResponse
	isSet bool
}

func (v NullableGiftSendersResponse) Get() *GiftSendersResponse {
	return v.value
}

func (v *NullableGiftSendersResponse) Set(val *GiftSendersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftSendersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftSendersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftSendersResponse(val *GiftSendersResponse) *NullableGiftSendersResponse {
	return &NullableGiftSendersResponse{value: val, isSet: true}
}

func (v NullableGiftSendersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftSendersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


