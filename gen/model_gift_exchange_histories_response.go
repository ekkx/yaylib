
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftExchangeHistoriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftExchangeHistoriesResponse{}

// GiftExchangeHistoriesResponse struct for GiftExchangeHistoriesResponse
type GiftExchangeHistoriesResponse struct {
	GiftExchangeHistories []GiftExchangeHistory `json:"gift_exchange_histories,omitempty"`
	NextPageValue NullableString `json:"next_page_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftExchangeHistoriesResponse GiftExchangeHistoriesResponse

// NewGiftExchangeHistoriesResponse instantiates a new GiftExchangeHistoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftExchangeHistoriesResponse() *GiftExchangeHistoriesResponse {
	this := GiftExchangeHistoriesResponse{}
	return &this
}

// NewGiftExchangeHistoriesResponseWithDefaults instantiates a new GiftExchangeHistoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftExchangeHistoriesResponseWithDefaults() *GiftExchangeHistoriesResponse {
	this := GiftExchangeHistoriesResponse{}
	return &this
}

// GetGiftExchangeHistories returns the GiftExchangeHistories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistoriesResponse) GetGiftExchangeHistories() []GiftExchangeHistory {
	if o == nil {
		var ret []GiftExchangeHistory
		return ret
	}
	return o.GiftExchangeHistories
}

// GetGiftExchangeHistoriesOk returns a tuple with the GiftExchangeHistories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistoriesResponse) GetGiftExchangeHistoriesOk() ([]GiftExchangeHistory, bool) {
	if o == nil || IsNil(o.GiftExchangeHistories) {
		return nil, false
	}
	return o.GiftExchangeHistories, true
}

// HasGiftExchangeHistories returns a boolean if a field has been set.
func (o *GiftExchangeHistoriesResponse) HasGiftExchangeHistories() bool {
	if o != nil && !IsNil(o.GiftExchangeHistories) {
		return true
	}

	return false
}

// SetGiftExchangeHistories gets a reference to the given []GiftExchangeHistory and assigns it to the GiftExchangeHistories field.
func (o *GiftExchangeHistoriesResponse) SetGiftExchangeHistories(v []GiftExchangeHistory) {
	o.GiftExchangeHistories = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistoriesResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistoriesResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GiftExchangeHistoriesResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *GiftExchangeHistoriesResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GiftExchangeHistoriesResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GiftExchangeHistoriesResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

func (o GiftExchangeHistoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftExchangeHistoriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GiftExchangeHistories != nil {
		toSerialize["gift_exchange_histories"] = o.GiftExchangeHistories
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftExchangeHistoriesResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftExchangeHistoriesResponse := _GiftExchangeHistoriesResponse{}

	err = json.Unmarshal(data, &varGiftExchangeHistoriesResponse)

	if err != nil {
		return err
	}

	*o = GiftExchangeHistoriesResponse(varGiftExchangeHistoriesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift_exchange_histories")
		delete(additionalProperties, "next_page_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftExchangeHistoriesResponse struct {
	value *GiftExchangeHistoriesResponse
	isSet bool
}

func (v NullableGiftExchangeHistoriesResponse) Get() *GiftExchangeHistoriesResponse {
	return v.value
}

func (v *NullableGiftExchangeHistoriesResponse) Set(val *GiftExchangeHistoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftExchangeHistoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftExchangeHistoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftExchangeHistoriesResponse(val *GiftExchangeHistoriesResponse) *NullableGiftExchangeHistoriesResponse {
	return &NullableGiftExchangeHistoriesResponse{value: val, isSet: true}
}

func (v NullableGiftExchangeHistoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftExchangeHistoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


