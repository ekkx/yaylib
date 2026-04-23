
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalResponse{}

// PalResponse struct for PalResponse
type PalResponse struct {
	Result NullablePalDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalResponse PalResponse

// NewPalResponse instantiates a new PalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalResponse() *PalResponse {
	this := PalResponse{}
	return &this
}

// NewPalResponseWithDefaults instantiates a new PalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalResponseWithDefaults() *PalResponse {
	this := PalResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalResponse) GetResult() PalDTO {
	if o == nil || IsNil(o.Result.Get()) {
		var ret PalDTO
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalResponse) GetResultOk() (*PalDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *PalResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullablePalDTO and assigns it to the Result field.
func (o *PalResponse) SetResult(v PalDTO) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *PalResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *PalResponse) UnsetResult() {
	o.Result.Unset()
}

func (o PalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalResponse) UnmarshalJSON(data []byte) (err error) {
	varPalResponse := _PalResponse{}

	err = json.Unmarshal(data, &varPalResponse)

	if err != nil {
		return err
	}

	*o = PalResponse(varPalResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalResponse struct {
	value *PalResponse
	isSet bool
}

func (v NullablePalResponse) Get() *PalResponse {
	return v.value
}

func (v *NullablePalResponse) Set(val *PalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalResponse(val *PalResponse) *NullablePalResponse {
	return &NullablePalResponse{value: val, isSet: true}
}

func (v NullablePalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


