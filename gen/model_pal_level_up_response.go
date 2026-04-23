
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalLevelUpResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalLevelUpResponse{}

// PalLevelUpResponse struct for PalLevelUpResponse
type PalLevelUpResponse struct {
	Result NullableLevelUpDetails `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalLevelUpResponse PalLevelUpResponse

// NewPalLevelUpResponse instantiates a new PalLevelUpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalLevelUpResponse() *PalLevelUpResponse {
	this := PalLevelUpResponse{}
	return &this
}

// NewPalLevelUpResponseWithDefaults instantiates a new PalLevelUpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalLevelUpResponseWithDefaults() *PalLevelUpResponse {
	this := PalLevelUpResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalLevelUpResponse) GetResult() LevelUpDetails {
	if o == nil || IsNil(o.Result.Get()) {
		var ret LevelUpDetails
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalLevelUpResponse) GetResultOk() (*LevelUpDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *PalLevelUpResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableLevelUpDetails and assigns it to the Result field.
func (o *PalLevelUpResponse) SetResult(v LevelUpDetails) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *PalLevelUpResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *PalLevelUpResponse) UnsetResult() {
	o.Result.Unset()
}

func (o PalLevelUpResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalLevelUpResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalLevelUpResponse) UnmarshalJSON(data []byte) (err error) {
	varPalLevelUpResponse := _PalLevelUpResponse{}

	err = json.Unmarshal(data, &varPalLevelUpResponse)

	if err != nil {
		return err
	}

	*o = PalLevelUpResponse(varPalLevelUpResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalLevelUpResponse struct {
	value *PalLevelUpResponse
	isSet bool
}

func (v NullablePalLevelUpResponse) Get() *PalLevelUpResponse {
	return v.value
}

func (v *NullablePalLevelUpResponse) Set(val *PalLevelUpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePalLevelUpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePalLevelUpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalLevelUpResponse(val *PalLevelUpResponse) *NullablePalLevelUpResponse {
	return &NullablePalLevelUpResponse{value: val, isSet: true}
}

func (v NullablePalLevelUpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalLevelUpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


