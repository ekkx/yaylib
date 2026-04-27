
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftTransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftTransactionsResponse{}

// GiftTransactionsResponse struct for GiftTransactionsResponse
type GiftTransactionsResponse struct {
	HideGiftsReceived NullableBool `json:"hide_gifts_received,omitempty"`
	NextPageValue NullableInt64 `json:"next_page_value,omitempty"`
	SentGifts []GiftHistory `json:"sent_gifts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftTransactionsResponse GiftTransactionsResponse

// NewGiftTransactionsResponse instantiates a new GiftTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftTransactionsResponse() *GiftTransactionsResponse {
	this := GiftTransactionsResponse{}
	return &this
}

// NewGiftTransactionsResponseWithDefaults instantiates a new GiftTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftTransactionsResponseWithDefaults() *GiftTransactionsResponse {
	this := GiftTransactionsResponse{}
	return &this
}

// GetHideGiftsReceived returns the HideGiftsReceived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftTransactionsResponse) GetHideGiftsReceived() bool {
	if o == nil || IsNil(o.HideGiftsReceived.Get()) {
		var ret bool
		return ret
	}
	return *o.HideGiftsReceived.Get()
}

// GetHideGiftsReceivedOk returns a tuple with the HideGiftsReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftTransactionsResponse) GetHideGiftsReceivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideGiftsReceived.Get(), o.HideGiftsReceived.IsSet()
}

// HasHideGiftsReceived returns a boolean if a field has been set.
func (o *GiftTransactionsResponse) HasHideGiftsReceived() bool {
	if o != nil && o.HideGiftsReceived.IsSet() {
		return true
	}

	return false
}

// SetHideGiftsReceived gets a reference to the given NullableBool and assigns it to the HideGiftsReceived field.
func (o *GiftTransactionsResponse) SetHideGiftsReceived(v bool) {
	o.HideGiftsReceived.Set(&v)
}
// SetHideGiftsReceivedNil sets the value for HideGiftsReceived to be an explicit nil
func (o *GiftTransactionsResponse) SetHideGiftsReceivedNil() {
	o.HideGiftsReceived.Set(nil)
}

// UnsetHideGiftsReceived ensures that no value is present for HideGiftsReceived, not even an explicit nil
func (o *GiftTransactionsResponse) UnsetHideGiftsReceived() {
	o.HideGiftsReceived.Unset()
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftTransactionsResponse) GetNextPageValue() int64 {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret int64
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftTransactionsResponse) GetNextPageValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GiftTransactionsResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableInt64 and assigns it to the NextPageValue field.
func (o *GiftTransactionsResponse) SetNextPageValue(v int64) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GiftTransactionsResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GiftTransactionsResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetSentGifts returns the SentGifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftTransactionsResponse) GetSentGifts() []GiftHistory {
	if o == nil {
		var ret []GiftHistory
		return ret
	}
	return o.SentGifts
}

// GetSentGiftsOk returns a tuple with the SentGifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftTransactionsResponse) GetSentGiftsOk() ([]GiftHistory, bool) {
	if o == nil || IsNil(o.SentGifts) {
		return nil, false
	}
	return o.SentGifts, true
}

// HasSentGifts returns a boolean if a field has been set.
func (o *GiftTransactionsResponse) HasSentGifts() bool {
	if o != nil && !IsNil(o.SentGifts) {
		return true
	}

	return false
}

// SetSentGifts gets a reference to the given []GiftHistory and assigns it to the SentGifts field.
func (o *GiftTransactionsResponse) SetSentGifts(v []GiftHistory) {
	o.SentGifts = v
}

func (o GiftTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftTransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HideGiftsReceived.IsSet() {
		toSerialize["hide_gifts_received"] = o.HideGiftsReceived.Get()
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.SentGifts != nil {
		toSerialize["sent_gifts"] = o.SentGifts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftTransactionsResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftTransactionsResponse := _GiftTransactionsResponse{}

	err = json.Unmarshal(data, &varGiftTransactionsResponse)

	if err != nil {
		return err
	}

	*o = GiftTransactionsResponse(varGiftTransactionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hide_gifts_received")
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "sent_gifts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftTransactionsResponse struct {
	value *GiftTransactionsResponse
	isSet bool
}

func (v NullableGiftTransactionsResponse) Get() *GiftTransactionsResponse {
	return v.value
}

func (v *NullableGiftTransactionsResponse) Set(val *GiftTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftTransactionsResponse(val *GiftTransactionsResponse) *NullableGiftTransactionsResponse {
	return &NullableGiftTransactionsResponse{value: val, isSet: true}
}

func (v NullableGiftTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


