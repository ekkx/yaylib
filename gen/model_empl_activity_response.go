
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplActivityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplActivityResponse{}

// EmplActivityResponse struct for EmplActivityResponse
type EmplActivityResponse struct {
	Cursor NullableString `json:"cursor,omitempty"`
	Result []EmplActivityDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplActivityResponse EmplActivityResponse

// NewEmplActivityResponse instantiates a new EmplActivityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplActivityResponse() *EmplActivityResponse {
	this := EmplActivityResponse{}
	return &this
}

// NewEmplActivityResponseWithDefaults instantiates a new EmplActivityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplActivityResponseWithDefaults() *EmplActivityResponse {
	this := EmplActivityResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityResponse) GetCursor() string {
	if o == nil || IsNil(o.Cursor.Get()) {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityResponse) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *EmplActivityResponse) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *EmplActivityResponse) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *EmplActivityResponse) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *EmplActivityResponse) UnsetCursor() {
	o.Cursor.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityResponse) GetResult() []EmplActivityDTO {
	if o == nil {
		var ret []EmplActivityDTO
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityResponse) GetResultOk() ([]EmplActivityDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *EmplActivityResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []EmplActivityDTO and assigns it to the Result field.
func (o *EmplActivityResponse) SetResult(v []EmplActivityDTO) {
	o.Result = v
}

func (o EmplActivityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplActivityResponse) ToMap() (map[string]interface{}, error) {
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

func (o *EmplActivityResponse) UnmarshalJSON(data []byte) (err error) {
	varEmplActivityResponse := _EmplActivityResponse{}

	err = json.Unmarshal(data, &varEmplActivityResponse)

	if err != nil {
		return err
	}

	*o = EmplActivityResponse(varEmplActivityResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplActivityResponse struct {
	value *EmplActivityResponse
	isSet bool
}

func (v NullableEmplActivityResponse) Get() *EmplActivityResponse {
	return v.value
}

func (v *NullableEmplActivityResponse) Set(val *EmplActivityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplActivityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplActivityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplActivityResponse(val *EmplActivityResponse) *NullableEmplActivityResponse {
	return &NullableEmplActivityResponse{value: val, isSet: true}
}

func (v NullableEmplActivityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplActivityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


