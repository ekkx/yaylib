
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CallGiftHistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallGiftHistoryResponse{}

// CallGiftHistoryResponse struct for CallGiftHistoryResponse
type CallGiftHistoryResponse struct {
	NextPageValue NullableString `json:"next_page_value,omitempty"`
	SentGifts []GiftHistory `json:"sent_gifts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallGiftHistoryResponse CallGiftHistoryResponse

// NewCallGiftHistoryResponse instantiates a new CallGiftHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallGiftHistoryResponse() *CallGiftHistoryResponse {
	this := CallGiftHistoryResponse{}
	return &this
}

// NewCallGiftHistoryResponseWithDefaults instantiates a new CallGiftHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallGiftHistoryResponseWithDefaults() *CallGiftHistoryResponse {
	this := CallGiftHistoryResponse{}
	return &this
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallGiftHistoryResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallGiftHistoryResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *CallGiftHistoryResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *CallGiftHistoryResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *CallGiftHistoryResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *CallGiftHistoryResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetSentGifts returns the SentGifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallGiftHistoryResponse) GetSentGifts() []GiftHistory {
	if o == nil {
		var ret []GiftHistory
		return ret
	}
	return o.SentGifts
}

// GetSentGiftsOk returns a tuple with the SentGifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallGiftHistoryResponse) GetSentGiftsOk() ([]GiftHistory, bool) {
	if o == nil || IsNil(o.SentGifts) {
		return nil, false
	}
	return o.SentGifts, true
}

// HasSentGifts returns a boolean if a field has been set.
func (o *CallGiftHistoryResponse) HasSentGifts() bool {
	if o != nil && !IsNil(o.SentGifts) {
		return true
	}

	return false
}

// SetSentGifts gets a reference to the given []GiftHistory and assigns it to the SentGifts field.
func (o *CallGiftHistoryResponse) SetSentGifts(v []GiftHistory) {
	o.SentGifts = v
}

func (o CallGiftHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallGiftHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *CallGiftHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varCallGiftHistoryResponse := _CallGiftHistoryResponse{}

	err = json.Unmarshal(data, &varCallGiftHistoryResponse)

	if err != nil {
		return err
	}

	*o = CallGiftHistoryResponse(varCallGiftHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "sent_gifts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallGiftHistoryResponse struct {
	value *CallGiftHistoryResponse
	isSet bool
}

func (v NullableCallGiftHistoryResponse) Get() *CallGiftHistoryResponse {
	return v.value
}

func (v *NullableCallGiftHistoryResponse) Set(val *CallGiftHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallGiftHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallGiftHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallGiftHistoryResponse(val *CallGiftHistoryResponse) *NullableCallGiftHistoryResponse {
	return &NullableCallGiftHistoryResponse{value: val, isSet: true}
}

func (v NullableCallGiftHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallGiftHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


