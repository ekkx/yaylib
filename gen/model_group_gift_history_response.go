
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupGiftHistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupGiftHistoryResponse{}

// GroupGiftHistoryResponse struct for GroupGiftHistoryResponse
type GroupGiftHistoryResponse struct {
	GiftHistory []GroupGiftHistory `json:"gift_history,omitempty"`
	NextPageValue NullableString `json:"next_page_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupGiftHistoryResponse GroupGiftHistoryResponse

// NewGroupGiftHistoryResponse instantiates a new GroupGiftHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupGiftHistoryResponse() *GroupGiftHistoryResponse {
	this := GroupGiftHistoryResponse{}
	return &this
}

// NewGroupGiftHistoryResponseWithDefaults instantiates a new GroupGiftHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupGiftHistoryResponseWithDefaults() *GroupGiftHistoryResponse {
	this := GroupGiftHistoryResponse{}
	return &this
}

// GetGiftHistory returns the GiftHistory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGiftHistoryResponse) GetGiftHistory() []GroupGiftHistory {
	if o == nil {
		var ret []GroupGiftHistory
		return ret
	}
	return o.GiftHistory
}

// GetGiftHistoryOk returns a tuple with the GiftHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGiftHistoryResponse) GetGiftHistoryOk() ([]GroupGiftHistory, bool) {
	if o == nil || IsNil(o.GiftHistory) {
		return nil, false
	}
	return o.GiftHistory, true
}

// HasGiftHistory returns a boolean if a field has been set.
func (o *GroupGiftHistoryResponse) HasGiftHistory() bool {
	if o != nil && !IsNil(o.GiftHistory) {
		return true
	}

	return false
}

// SetGiftHistory gets a reference to the given []GroupGiftHistory and assigns it to the GiftHistory field.
func (o *GroupGiftHistoryResponse) SetGiftHistory(v []GroupGiftHistory) {
	o.GiftHistory = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupGiftHistoryResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupGiftHistoryResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GroupGiftHistoryResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *GroupGiftHistoryResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GroupGiftHistoryResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GroupGiftHistoryResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

func (o GroupGiftHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupGiftHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GiftHistory != nil {
		toSerialize["gift_history"] = o.GiftHistory
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupGiftHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupGiftHistoryResponse := _GroupGiftHistoryResponse{}

	err = json.Unmarshal(data, &varGroupGiftHistoryResponse)

	if err != nil {
		return err
	}

	*o = GroupGiftHistoryResponse(varGroupGiftHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gift_history")
		delete(additionalProperties, "next_page_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupGiftHistoryResponse struct {
	value *GroupGiftHistoryResponse
	isSet bool
}

func (v NullableGroupGiftHistoryResponse) Get() *GroupGiftHistoryResponse {
	return v.value
}

func (v *NullableGroupGiftHistoryResponse) Set(val *GroupGiftHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupGiftHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupGiftHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupGiftHistoryResponse(val *GroupGiftHistoryResponse) *NullableGroupGiftHistoryResponse {
	return &NullableGroupGiftHistoryResponse{value: val, isSet: true}
}

func (v NullableGroupGiftHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupGiftHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


