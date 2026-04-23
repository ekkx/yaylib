
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplFeeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplFeeResponse{}

// EmplFeeResponse struct for EmplFeeResponse
type EmplFeeResponse struct {
	Result NullableEmplFee `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplFeeResponse EmplFeeResponse

// NewEmplFeeResponse instantiates a new EmplFeeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplFeeResponse() *EmplFeeResponse {
	this := EmplFeeResponse{}
	return &this
}

// NewEmplFeeResponseWithDefaults instantiates a new EmplFeeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplFeeResponseWithDefaults() *EmplFeeResponse {
	this := EmplFeeResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplFeeResponse) GetResult() EmplFee {
	if o == nil || IsNil(o.Result.Get()) {
		var ret EmplFee
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplFeeResponse) GetResultOk() (*EmplFee, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *EmplFeeResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableEmplFee and assigns it to the Result field.
func (o *EmplFeeResponse) SetResult(v EmplFee) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *EmplFeeResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *EmplFeeResponse) UnsetResult() {
	o.Result.Unset()
}

func (o EmplFeeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplFeeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplFeeResponse) UnmarshalJSON(data []byte) (err error) {
	varEmplFeeResponse := _EmplFeeResponse{}

	err = json.Unmarshal(data, &varEmplFeeResponse)

	if err != nil {
		return err
	}

	*o = EmplFeeResponse(varEmplFeeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplFeeResponse struct {
	value *EmplFeeResponse
	isSet bool
}

func (v NullableEmplFeeResponse) Get() *EmplFeeResponse {
	return v.value
}

func (v *NullableEmplFeeResponse) Set(val *EmplFeeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplFeeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplFeeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplFeeResponse(val *EmplFeeResponse) *NullableEmplFeeResponse {
	return &NullableEmplFeeResponse{value: val, isSet: true}
}

func (v NullableEmplFeeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplFeeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


