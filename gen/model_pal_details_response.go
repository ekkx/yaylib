
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalDetailsResponse{}

// PalDetailsResponse struct for PalDetailsResponse
type PalDetailsResponse struct {
	Result NullablePalDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalDetailsResponse PalDetailsResponse

// NewPalDetailsResponse instantiates a new PalDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalDetailsResponse() *PalDetailsResponse {
	this := PalDetailsResponse{}
	return &this
}

// NewPalDetailsResponseWithDefaults instantiates a new PalDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalDetailsResponseWithDefaults() *PalDetailsResponse {
	this := PalDetailsResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDetailsResponse) GetResult() PalDTO {
	if o == nil || IsNil(o.Result.Get()) {
		var ret PalDTO
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDetailsResponse) GetResultOk() (*PalDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *PalDetailsResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullablePalDTO and assigns it to the Result field.
func (o *PalDetailsResponse) SetResult(v PalDTO) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *PalDetailsResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *PalDetailsResponse) UnsetResult() {
	o.Result.Unset()
}

func (o PalDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	varPalDetailsResponse := _PalDetailsResponse{}

	err = json.Unmarshal(data, &varPalDetailsResponse)

	if err != nil {
		return err
	}

	*o = PalDetailsResponse(varPalDetailsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalDetailsResponse struct {
	value *PalDetailsResponse
	isSet bool
}

func (v NullablePalDetailsResponse) Get() *PalDetailsResponse {
	return v.value
}

func (v *NullablePalDetailsResponse) Set(val *PalDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalDetailsResponse(val *PalDetailsResponse) *NullablePalDetailsResponse {
	return &NullablePalDetailsResponse{value: val, isSet: true}
}

func (v NullablePalDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


