
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the NftInternalList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NftInternalList{}

// NftInternalList struct for NftInternalList
type NftInternalList struct {
	Cursor NullableString `json:"cursor,omitempty"`
	Result []PalDTO `json:"result,omitempty"`
	TotalCount NullableInt32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NftInternalList NftInternalList

// NewNftInternalList instantiates a new NftInternalList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftInternalList() *NftInternalList {
	this := NftInternalList{}
	return &this
}

// NewNftInternalListWithDefaults instantiates a new NftInternalList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftInternalListWithDefaults() *NftInternalList {
	this := NftInternalList{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftInternalList) GetCursor() string {
	if o == nil || IsNil(o.Cursor.Get()) {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftInternalList) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *NftInternalList) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *NftInternalList) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *NftInternalList) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *NftInternalList) UnsetCursor() {
	o.Cursor.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftInternalList) GetResult() []PalDTO {
	if o == nil {
		var ret []PalDTO
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftInternalList) GetResultOk() ([]PalDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *NftInternalList) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []PalDTO and assigns it to the Result field.
func (o *NftInternalList) SetResult(v []PalDTO) {
	o.Result = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftInternalList) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftInternalList) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *NftInternalList) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt32 and assigns it to the TotalCount field.
func (o *NftInternalList) SetTotalCount(v int32) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *NftInternalList) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *NftInternalList) UnsetTotalCount() {
	o.TotalCount.Unset()
}

func (o NftInternalList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NftInternalList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.TotalCount.IsSet() {
		toSerialize["total_count"] = o.TotalCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NftInternalList) UnmarshalJSON(data []byte) (err error) {
	varNftInternalList := _NftInternalList{}

	err = json.Unmarshal(data, &varNftInternalList)

	if err != nil {
		return err
	}

	*o = NftInternalList(varNftInternalList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "result")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNftInternalList struct {
	value *NftInternalList
	isSet bool
}

func (v NullableNftInternalList) Get() *NftInternalList {
	return v.value
}

func (v *NullableNftInternalList) Set(val *NftInternalList) {
	v.value = val
	v.isSet = true
}

func (v NullableNftInternalList) IsSet() bool {
	return v.isSet
}

func (v *NullableNftInternalList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftInternalList(val *NftInternalList) *NullableNftInternalList {
	return &NullableNftInternalList{value: val, isSet: true}
}

func (v NullableNftInternalList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftInternalList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


