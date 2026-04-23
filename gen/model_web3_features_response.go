
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3FeaturesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3FeaturesResponse{}

// Web3FeaturesResponse struct for Web3FeaturesResponse
type Web3FeaturesResponse struct {
	Result NullableFeaturesResultResponse `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3FeaturesResponse Web3FeaturesResponse

// NewWeb3FeaturesResponse instantiates a new Web3FeaturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3FeaturesResponse() *Web3FeaturesResponse {
	this := Web3FeaturesResponse{}
	return &this
}

// NewWeb3FeaturesResponseWithDefaults instantiates a new Web3FeaturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3FeaturesResponseWithDefaults() *Web3FeaturesResponse {
	this := Web3FeaturesResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3FeaturesResponse) GetResult() FeaturesResultResponse {
	if o == nil || IsNil(o.Result.Get()) {
		var ret FeaturesResultResponse
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3FeaturesResponse) GetResultOk() (*FeaturesResultResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *Web3FeaturesResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableFeaturesResultResponse and assigns it to the Result field.
func (o *Web3FeaturesResponse) SetResult(v FeaturesResultResponse) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *Web3FeaturesResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *Web3FeaturesResponse) UnsetResult() {
	o.Result.Unset()
}

func (o Web3FeaturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3FeaturesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3FeaturesResponse) UnmarshalJSON(data []byte) (err error) {
	varWeb3FeaturesResponse := _Web3FeaturesResponse{}

	err = json.Unmarshal(data, &varWeb3FeaturesResponse)

	if err != nil {
		return err
	}

	*o = Web3FeaturesResponse(varWeb3FeaturesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3FeaturesResponse struct {
	value *Web3FeaturesResponse
	isSet bool
}

func (v NullableWeb3FeaturesResponse) Get() *Web3FeaturesResponse {
	return v.value
}

func (v *NullableWeb3FeaturesResponse) Set(val *Web3FeaturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3FeaturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3FeaturesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3FeaturesResponse(val *Web3FeaturesResponse) *NullableWeb3FeaturesResponse {
	return &NullableWeb3FeaturesResponse{value: val, isSet: true}
}

func (v NullableWeb3FeaturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3FeaturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


