
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalGachaHistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalGachaHistoryResponse{}

// PalGachaHistoryResponse struct for PalGachaHistoryResponse
type PalGachaHistoryResponse struct {
	Cursor NullableString `json:"cursor,omitempty"`
	Result []PalGachaHistoryResponseItem `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalGachaHistoryResponse PalGachaHistoryResponse

// NewPalGachaHistoryResponse instantiates a new PalGachaHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalGachaHistoryResponse() *PalGachaHistoryResponse {
	this := PalGachaHistoryResponse{}
	return &this
}

// NewPalGachaHistoryResponseWithDefaults instantiates a new PalGachaHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalGachaHistoryResponseWithDefaults() *PalGachaHistoryResponse {
	this := PalGachaHistoryResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponse) GetCursor() string {
	if o == nil || IsNil(o.Cursor.Get()) {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponse) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *PalGachaHistoryResponse) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *PalGachaHistoryResponse) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *PalGachaHistoryResponse) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *PalGachaHistoryResponse) UnsetCursor() {
	o.Cursor.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaHistoryResponse) GetResult() []PalGachaHistoryResponseItem {
	if o == nil {
		var ret []PalGachaHistoryResponseItem
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaHistoryResponse) GetResultOk() ([]PalGachaHistoryResponseItem, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PalGachaHistoryResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []PalGachaHistoryResponseItem and assigns it to the Result field.
func (o *PalGachaHistoryResponse) SetResult(v []PalGachaHistoryResponseItem) {
	o.Result = v
}

func (o PalGachaHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalGachaHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalGachaHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varPalGachaHistoryResponse := _PalGachaHistoryResponse{}

	err = json.Unmarshal(data, &varPalGachaHistoryResponse)

	if err != nil {
		return err
	}

	*o = PalGachaHistoryResponse(varPalGachaHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalGachaHistoryResponse struct {
	value *PalGachaHistoryResponse
	isSet bool
}

func (v NullablePalGachaHistoryResponse) Get() *PalGachaHistoryResponse {
	return v.value
}

func (v *NullablePalGachaHistoryResponse) Set(val *PalGachaHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalGachaHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalGachaHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalGachaHistoryResponse(val *PalGachaHistoryResponse) *NullablePalGachaHistoryResponse {
	return &NullablePalGachaHistoryResponse{value: val, isSet: true}
}

func (v NullablePalGachaHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalGachaHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


