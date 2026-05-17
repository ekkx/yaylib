
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftReceivedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftReceivedResponse{}

// GiftReceivedResponse struct for GiftReceivedResponse
type GiftReceivedResponse struct {
	NextPageValue NullableCursor `json:"next_page_value,omitempty"`
	ReceivedGifts []ReceivedGift `json:"received_gifts,omitempty"`
	TotalCount NullableInt32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftReceivedResponse GiftReceivedResponse

// NewGiftReceivedResponse instantiates a new GiftReceivedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftReceivedResponse() *GiftReceivedResponse {
	this := GiftReceivedResponse{}
	return &this
}

// NewGiftReceivedResponseWithDefaults instantiates a new GiftReceivedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftReceivedResponseWithDefaults() *GiftReceivedResponse {
	this := GiftReceivedResponse{}
	return &this
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReceivedResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReceivedResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GiftReceivedResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *GiftReceivedResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GiftReceivedResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GiftReceivedResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetReceivedGifts returns the ReceivedGifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReceivedResponse) GetReceivedGifts() []ReceivedGift {
	if o == nil {
		var ret []ReceivedGift
		return ret
	}
	return o.ReceivedGifts
}

// GetReceivedGiftsOk returns a tuple with the ReceivedGifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReceivedResponse) GetReceivedGiftsOk() ([]ReceivedGift, bool) {
	if o == nil || IsNil(o.ReceivedGifts) {
		return nil, false
	}
	return o.ReceivedGifts, true
}

// HasReceivedGifts returns a boolean if a field has been set.
func (o *GiftReceivedResponse) HasReceivedGifts() bool {
	if o != nil && !IsNil(o.ReceivedGifts) {
		return true
	}

	return false
}

// SetReceivedGifts gets a reference to the given []ReceivedGift and assigns it to the ReceivedGifts field.
func (o *GiftReceivedResponse) SetReceivedGifts(v []ReceivedGift) {
	o.ReceivedGifts = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftReceivedResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftReceivedResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GiftReceivedResponse) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt32 and assigns it to the TotalCount field.
func (o *GiftReceivedResponse) SetTotalCount(v int32) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *GiftReceivedResponse) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *GiftReceivedResponse) UnsetTotalCount() {
	o.TotalCount.Unset()
}

func (o GiftReceivedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftReceivedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.ReceivedGifts != nil {
		toSerialize["received_gifts"] = o.ReceivedGifts
	}
	if o.TotalCount.IsSet() {
		toSerialize["total_count"] = o.TotalCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftReceivedResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftReceivedResponse := _GiftReceivedResponse{}

	err = json.Unmarshal(data, &varGiftReceivedResponse)

	if err != nil {
		return err
	}

	*o = GiftReceivedResponse(varGiftReceivedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "received_gifts")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftReceivedResponse struct {
	value *GiftReceivedResponse
	isSet bool
}

func (v NullableGiftReceivedResponse) Get() *GiftReceivedResponse {
	return v.value
}

func (v *NullableGiftReceivedResponse) Set(val *GiftReceivedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftReceivedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftReceivedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftReceivedResponse(val *GiftReceivedResponse) *NullableGiftReceivedResponse {
	return &NullableGiftReceivedResponse{value: val, isSet: true}
}

func (v NullableGiftReceivedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftReceivedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


