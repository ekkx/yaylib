
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3EmplTokenExchangeConversionAndFeeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3EmplTokenExchangeConversionAndFeeResponse{}

// Web3EmplTokenExchangeConversionAndFeeResponse struct for Web3EmplTokenExchangeConversionAndFeeResponse
type Web3EmplTokenExchangeConversionAndFeeResponse struct {
	Result NullableWeb3EmplTokenExchangeConversionAndFeeDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3EmplTokenExchangeConversionAndFeeResponse Web3EmplTokenExchangeConversionAndFeeResponse

// NewWeb3EmplTokenExchangeConversionAndFeeResponse instantiates a new Web3EmplTokenExchangeConversionAndFeeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3EmplTokenExchangeConversionAndFeeResponse() *Web3EmplTokenExchangeConversionAndFeeResponse {
	this := Web3EmplTokenExchangeConversionAndFeeResponse{}
	return &this
}

// NewWeb3EmplTokenExchangeConversionAndFeeResponseWithDefaults instantiates a new Web3EmplTokenExchangeConversionAndFeeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3EmplTokenExchangeConversionAndFeeResponseWithDefaults() *Web3EmplTokenExchangeConversionAndFeeResponse {
	this := Web3EmplTokenExchangeConversionAndFeeResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplTokenExchangeConversionAndFeeResponse) GetResult() Web3EmplTokenExchangeConversionAndFeeDTO {
	if o == nil || IsNil(o.Result.Get()) {
		var ret Web3EmplTokenExchangeConversionAndFeeDTO
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplTokenExchangeConversionAndFeeResponse) GetResultOk() (*Web3EmplTokenExchangeConversionAndFeeDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *Web3EmplTokenExchangeConversionAndFeeResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableWeb3EmplTokenExchangeConversionAndFeeDTO and assigns it to the Result field.
func (o *Web3EmplTokenExchangeConversionAndFeeResponse) SetResult(v Web3EmplTokenExchangeConversionAndFeeDTO) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *Web3EmplTokenExchangeConversionAndFeeResponse) UnsetResult() {
	o.Result.Unset()
}

func (o Web3EmplTokenExchangeConversionAndFeeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3EmplTokenExchangeConversionAndFeeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3EmplTokenExchangeConversionAndFeeResponse) UnmarshalJSON(data []byte) (err error) {
	varWeb3EmplTokenExchangeConversionAndFeeResponse := _Web3EmplTokenExchangeConversionAndFeeResponse{}

	err = json.Unmarshal(data, &varWeb3EmplTokenExchangeConversionAndFeeResponse)

	if err != nil {
		return err
	}

	*o = Web3EmplTokenExchangeConversionAndFeeResponse(varWeb3EmplTokenExchangeConversionAndFeeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3EmplTokenExchangeConversionAndFeeResponse struct {
	value *Web3EmplTokenExchangeConversionAndFeeResponse
	isSet bool
}

func (v NullableWeb3EmplTokenExchangeConversionAndFeeResponse) Get() *Web3EmplTokenExchangeConversionAndFeeResponse {
	return v.value
}

func (v *NullableWeb3EmplTokenExchangeConversionAndFeeResponse) Set(val *Web3EmplTokenExchangeConversionAndFeeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3EmplTokenExchangeConversionAndFeeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3EmplTokenExchangeConversionAndFeeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3EmplTokenExchangeConversionAndFeeResponse(val *Web3EmplTokenExchangeConversionAndFeeResponse) *NullableWeb3EmplTokenExchangeConversionAndFeeResponse {
	return &NullableWeb3EmplTokenExchangeConversionAndFeeResponse{value: val, isSet: true}
}

func (v NullableWeb3EmplTokenExchangeConversionAndFeeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3EmplTokenExchangeConversionAndFeeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


