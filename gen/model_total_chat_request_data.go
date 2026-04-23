
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TotalChatRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TotalChatRequestData{}

// TotalChatRequestData struct for TotalChatRequestData
type TotalChatRequestData struct {
	TotalCount NullableInt32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TotalChatRequestData TotalChatRequestData

// NewTotalChatRequestData instantiates a new TotalChatRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotalChatRequestData() *TotalChatRequestData {
	this := TotalChatRequestData{}
	return &this
}

// NewTotalChatRequestDataWithDefaults instantiates a new TotalChatRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalChatRequestDataWithDefaults() *TotalChatRequestData {
	this := TotalChatRequestData{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TotalChatRequestData) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TotalChatRequestData) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *TotalChatRequestData) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt32 and assigns it to the TotalCount field.
func (o *TotalChatRequestData) SetTotalCount(v int32) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *TotalChatRequestData) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *TotalChatRequestData) UnsetTotalCount() {
	o.TotalCount.Unset()
}

func (o TotalChatRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TotalChatRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount.IsSet() {
		toSerialize["total_count"] = o.TotalCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TotalChatRequestData) UnmarshalJSON(data []byte) (err error) {
	varTotalChatRequestData := _TotalChatRequestData{}

	err = json.Unmarshal(data, &varTotalChatRequestData)

	if err != nil {
		return err
	}

	*o = TotalChatRequestData(varTotalChatRequestData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTotalChatRequestData struct {
	value *TotalChatRequestData
	isSet bool
}

func (v NullableTotalChatRequestData) Get() *TotalChatRequestData {
	return v.value
}

func (v *NullableTotalChatRequestData) Set(val *TotalChatRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalChatRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalChatRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalChatRequestData(val *TotalChatRequestData) *NullableTotalChatRequestData {
	return &NullableTotalChatRequestData{value: val, isSet: true}
}

func (v NullableTotalChatRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalChatRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


