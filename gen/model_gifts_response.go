
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftsResponse{}

// GiftsResponse struct for GiftsResponse
type GiftsResponse struct {
	Gifts []RealmGift `json:"gifts,omitempty"`
	NextPageValue NullableInt64 `json:"next_page_value,omitempty"`
	TotalCount NullableInt32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftsResponse GiftsResponse

// NewGiftsResponse instantiates a new GiftsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftsResponse() *GiftsResponse {
	this := GiftsResponse{}
	return &this
}

// NewGiftsResponseWithDefaults instantiates a new GiftsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftsResponseWithDefaults() *GiftsResponse {
	this := GiftsResponse{}
	return &this
}

// GetGifts returns the Gifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftsResponse) GetGifts() []RealmGift {
	if o == nil {
		var ret []RealmGift
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftsResponse) GetGiftsOk() ([]RealmGift, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *GiftsResponse) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given []RealmGift and assigns it to the Gifts field.
func (o *GiftsResponse) SetGifts(v []RealmGift) {
	o.Gifts = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftsResponse) GetNextPageValue() int64 {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret int64
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftsResponse) GetNextPageValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GiftsResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableInt64 and assigns it to the NextPageValue field.
func (o *GiftsResponse) SetNextPageValue(v int64) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GiftsResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GiftsResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftsResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftsResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GiftsResponse) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt32 and assigns it to the TotalCount field.
func (o *GiftsResponse) SetTotalCount(v int32) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *GiftsResponse) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *GiftsResponse) UnsetTotalCount() {
	o.TotalCount.Unset()
}

func (o GiftsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gifts != nil {
		toSerialize["gifts"] = o.Gifts
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.TotalCount.IsSet() {
		toSerialize["total_count"] = o.TotalCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftsResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftsResponse := _GiftsResponse{}

	err = json.Unmarshal(data, &varGiftsResponse)

	if err != nil {
		return err
	}

	*o = GiftsResponse(varGiftsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gifts")
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftsResponse struct {
	value *GiftsResponse
	isSet bool
}

func (v NullableGiftsResponse) Get() *GiftsResponse {
	return v.value
}

func (v *NullableGiftsResponse) Set(val *GiftsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftsResponse(val *GiftsResponse) *NullableGiftsResponse {
	return &NullableGiftsResponse{value: val, isSet: true}
}

func (v NullableGiftsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


