
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalListResponse{}

// PalListResponse struct for PalListResponse
type PalListResponse struct {
	Cursor NullableString `json:"cursor,omitempty"`
	Result []PalDTO `json:"result,omitempty"`
	TotalCount NullableInt32 `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalListResponse PalListResponse

// NewPalListResponse instantiates a new PalListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalListResponse() *PalListResponse {
	this := PalListResponse{}
	return &this
}

// NewPalListResponseWithDefaults instantiates a new PalListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalListResponseWithDefaults() *PalListResponse {
	this := PalListResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalListResponse) GetCursor() string {
	if o == nil || IsNil(o.Cursor.Get()) {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalListResponse) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *PalListResponse) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *PalListResponse) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *PalListResponse) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *PalListResponse) UnsetCursor() {
	o.Cursor.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalListResponse) GetResult() []PalDTO {
	if o == nil {
		var ret []PalDTO
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalListResponse) GetResultOk() ([]PalDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PalListResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []PalDTO and assigns it to the Result field.
func (o *PalListResponse) SetResult(v []PalDTO) {
	o.Result = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalListResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalCount.Get()
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalListResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCount.Get(), o.TotalCount.IsSet()
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PalListResponse) HasTotalCount() bool {
	if o != nil && o.TotalCount.IsSet() {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given NullableInt32 and assigns it to the TotalCount field.
func (o *PalListResponse) SetTotalCount(v int32) {
	o.TotalCount.Set(&v)
}
// SetTotalCountNil sets the value for TotalCount to be an explicit nil
func (o *PalListResponse) SetTotalCountNil() {
	o.TotalCount.Set(nil)
}

// UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
func (o *PalListResponse) UnsetTotalCount() {
	o.TotalCount.Unset()
}

func (o PalListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *PalListResponse) UnmarshalJSON(data []byte) (err error) {
	varPalListResponse := _PalListResponse{}

	err = json.Unmarshal(data, &varPalListResponse)

	if err != nil {
		return err
	}

	*o = PalListResponse(varPalListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "result")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalListResponse struct {
	value *PalListResponse
	isSet bool
}

func (v NullablePalListResponse) Get() *PalListResponse {
	return v.value
}

func (v *NullablePalListResponse) Set(val *PalListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalListResponse(val *PalListResponse) *NullablePalListResponse {
	return &NullablePalListResponse{value: val, isSet: true}
}

func (v NullablePalListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


