
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplActivityDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplActivityDetailResponse{}

// EmplActivityDetailResponse struct for EmplActivityDetailResponse
type EmplActivityDetailResponse struct {
	Result NullableEmplActivityDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplActivityDetailResponse EmplActivityDetailResponse

// NewEmplActivityDetailResponse instantiates a new EmplActivityDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplActivityDetailResponse() *EmplActivityDetailResponse {
	this := EmplActivityDetailResponse{}
	return &this
}

// NewEmplActivityDetailResponseWithDefaults instantiates a new EmplActivityDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplActivityDetailResponseWithDefaults() *EmplActivityDetailResponse {
	this := EmplActivityDetailResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplActivityDetailResponse) GetResult() EmplActivityDTO {
	if o == nil || IsNil(o.Result.Get()) {
		var ret EmplActivityDTO
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplActivityDetailResponse) GetResultOk() (*EmplActivityDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *EmplActivityDetailResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableEmplActivityDTO and assigns it to the Result field.
func (o *EmplActivityDetailResponse) SetResult(v EmplActivityDTO) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *EmplActivityDetailResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *EmplActivityDetailResponse) UnsetResult() {
	o.Result.Unset()
}

func (o EmplActivityDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplActivityDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplActivityDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varEmplActivityDetailResponse := _EmplActivityDetailResponse{}

	err = json.Unmarshal(data, &varEmplActivityDetailResponse)

	if err != nil {
		return err
	}

	*o = EmplActivityDetailResponse(varEmplActivityDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplActivityDetailResponse struct {
	value *EmplActivityDetailResponse
	isSet bool
}

func (v NullableEmplActivityDetailResponse) Get() *EmplActivityDetailResponse {
	return v.value
}

func (v *NullableEmplActivityDetailResponse) Set(val *EmplActivityDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplActivityDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplActivityDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplActivityDetailResponse(val *EmplActivityDetailResponse) *NullableEmplActivityDetailResponse {
	return &NullableEmplActivityDetailResponse{value: val, isSet: true}
}

func (v NullableEmplActivityDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplActivityDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


