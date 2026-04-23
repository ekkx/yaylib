
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalFreePoolStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalFreePoolStatusResponse{}

// PalFreePoolStatusResponse struct for PalFreePoolStatusResponse
type PalFreePoolStatusResponse struct {
	Result NullablePalFreePoolStatus `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalFreePoolStatusResponse PalFreePoolStatusResponse

// NewPalFreePoolStatusResponse instantiates a new PalFreePoolStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalFreePoolStatusResponse() *PalFreePoolStatusResponse {
	this := PalFreePoolStatusResponse{}
	return &this
}

// NewPalFreePoolStatusResponseWithDefaults instantiates a new PalFreePoolStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalFreePoolStatusResponseWithDefaults() *PalFreePoolStatusResponse {
	this := PalFreePoolStatusResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalFreePoolStatusResponse) GetResult() PalFreePoolStatus {
	if o == nil || IsNil(o.Result.Get()) {
		var ret PalFreePoolStatus
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalFreePoolStatusResponse) GetResultOk() (*PalFreePoolStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *PalFreePoolStatusResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullablePalFreePoolStatus and assigns it to the Result field.
func (o *PalFreePoolStatusResponse) SetResult(v PalFreePoolStatus) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *PalFreePoolStatusResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *PalFreePoolStatusResponse) UnsetResult() {
	o.Result.Unset()
}

func (o PalFreePoolStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalFreePoolStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalFreePoolStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varPalFreePoolStatusResponse := _PalFreePoolStatusResponse{}

	err = json.Unmarshal(data, &varPalFreePoolStatusResponse)

	if err != nil {
		return err
	}

	*o = PalFreePoolStatusResponse(varPalFreePoolStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalFreePoolStatusResponse struct {
	value *PalFreePoolStatusResponse
	isSet bool
}

func (v NullablePalFreePoolStatusResponse) Get() *PalFreePoolStatusResponse {
	return v.value
}

func (v *NullablePalFreePoolStatusResponse) Set(val *PalFreePoolStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalFreePoolStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalFreePoolStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalFreePoolStatusResponse(val *PalFreePoolStatusResponse) *NullablePalFreePoolStatusResponse {
	return &NullablePalFreePoolStatusResponse{value: val, isSet: true}
}

func (v NullablePalFreePoolStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalFreePoolStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


