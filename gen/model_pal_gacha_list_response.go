
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalGachaListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalGachaListResponse{}

// PalGachaListResponse struct for PalGachaListResponse
type PalGachaListResponse struct {
	Result []PalGacha `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalGachaListResponse PalGachaListResponse

// NewPalGachaListResponse instantiates a new PalGachaListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalGachaListResponse() *PalGachaListResponse {
	this := PalGachaListResponse{}
	return &this
}

// NewPalGachaListResponseWithDefaults instantiates a new PalGachaListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalGachaListResponseWithDefaults() *PalGachaListResponse {
	this := PalGachaListResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalGachaListResponse) GetResult() []PalGacha {
	if o == nil {
		var ret []PalGacha
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalGachaListResponse) GetResultOk() ([]PalGacha, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PalGachaListResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []PalGacha and assigns it to the Result field.
func (o *PalGachaListResponse) SetResult(v []PalGacha) {
	o.Result = v
}

func (o PalGachaListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalGachaListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalGachaListResponse) UnmarshalJSON(data []byte) (err error) {
	varPalGachaListResponse := _PalGachaListResponse{}

	err = json.Unmarshal(data, &varPalGachaListResponse)

	if err != nil {
		return err
	}

	*o = PalGachaListResponse(varPalGachaListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalGachaListResponse struct {
	value *PalGachaListResponse
	isSet bool
}

func (v NullablePalGachaListResponse) Get() *PalGachaListResponse {
	return v.value
}

func (v *NullablePalGachaListResponse) Set(val *PalGachaListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalGachaListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalGachaListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalGachaListResponse(val *PalGachaListResponse) *NullablePalGachaListResponse {
	return &NullablePalGachaListResponse{value: val, isSet: true}
}

func (v NullablePalGachaListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalGachaListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


