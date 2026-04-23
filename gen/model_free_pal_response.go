
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FreePalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreePalResponse{}

// FreePalResponse struct for FreePalResponse
type FreePalResponse struct {
	Result NullablePalDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FreePalResponse FreePalResponse

// NewFreePalResponse instantiates a new FreePalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreePalResponse() *FreePalResponse {
	this := FreePalResponse{}
	return &this
}

// NewFreePalResponseWithDefaults instantiates a new FreePalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreePalResponseWithDefaults() *FreePalResponse {
	this := FreePalResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FreePalResponse) GetResult() PalDTO {
	if o == nil || IsNil(o.Result.Get()) {
		var ret PalDTO
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FreePalResponse) GetResultOk() (*PalDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *FreePalResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullablePalDTO and assigns it to the Result field.
func (o *FreePalResponse) SetResult(v PalDTO) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *FreePalResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *FreePalResponse) UnsetResult() {
	o.Result.Unset()
}

func (o FreePalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreePalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FreePalResponse) UnmarshalJSON(data []byte) (err error) {
	varFreePalResponse := _FreePalResponse{}

	err = json.Unmarshal(data, &varFreePalResponse)

	if err != nil {
		return err
	}

	*o = FreePalResponse(varFreePalResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFreePalResponse struct {
	value *FreePalResponse
	isSet bool
}

func (v NullableFreePalResponse) Get() *FreePalResponse {
	return v.value
}

func (v *NullableFreePalResponse) Set(val *FreePalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFreePalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFreePalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreePalResponse(val *FreePalResponse) *NullableFreePalResponse {
	return &NullableFreePalResponse{value: val, isSet: true}
}

func (v NullableFreePalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreePalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


