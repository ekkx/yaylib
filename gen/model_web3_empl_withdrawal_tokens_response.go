
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3EmplWithdrawalTokensResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3EmplWithdrawalTokensResponse{}

// Web3EmplWithdrawalTokensResponse struct for Web3EmplWithdrawalTokensResponse
type Web3EmplWithdrawalTokensResponse struct {
	Result NullableWeb3EmplWithdrawalTokensDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3EmplWithdrawalTokensResponse Web3EmplWithdrawalTokensResponse

// NewWeb3EmplWithdrawalTokensResponse instantiates a new Web3EmplWithdrawalTokensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3EmplWithdrawalTokensResponse() *Web3EmplWithdrawalTokensResponse {
	this := Web3EmplWithdrawalTokensResponse{}
	return &this
}

// NewWeb3EmplWithdrawalTokensResponseWithDefaults instantiates a new Web3EmplWithdrawalTokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3EmplWithdrawalTokensResponseWithDefaults() *Web3EmplWithdrawalTokensResponse {
	this := Web3EmplWithdrawalTokensResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3EmplWithdrawalTokensResponse) GetResult() Web3EmplWithdrawalTokensDTO {
	if o == nil || IsNil(o.Result.Get()) {
		var ret Web3EmplWithdrawalTokensDTO
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3EmplWithdrawalTokensResponse) GetResultOk() (*Web3EmplWithdrawalTokensDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *Web3EmplWithdrawalTokensResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableWeb3EmplWithdrawalTokensDTO and assigns it to the Result field.
func (o *Web3EmplWithdrawalTokensResponse) SetResult(v Web3EmplWithdrawalTokensDTO) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *Web3EmplWithdrawalTokensResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *Web3EmplWithdrawalTokensResponse) UnsetResult() {
	o.Result.Unset()
}

func (o Web3EmplWithdrawalTokensResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3EmplWithdrawalTokensResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3EmplWithdrawalTokensResponse) UnmarshalJSON(data []byte) (err error) {
	varWeb3EmplWithdrawalTokensResponse := _Web3EmplWithdrawalTokensResponse{}

	err = json.Unmarshal(data, &varWeb3EmplWithdrawalTokensResponse)

	if err != nil {
		return err
	}

	*o = Web3EmplWithdrawalTokensResponse(varWeb3EmplWithdrawalTokensResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3EmplWithdrawalTokensResponse struct {
	value *Web3EmplWithdrawalTokensResponse
	isSet bool
}

func (v NullableWeb3EmplWithdrawalTokensResponse) Get() *Web3EmplWithdrawalTokensResponse {
	return v.value
}

func (v *NullableWeb3EmplWithdrawalTokensResponse) Set(val *Web3EmplWithdrawalTokensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3EmplWithdrawalTokensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3EmplWithdrawalTokensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3EmplWithdrawalTokensResponse(val *Web3EmplWithdrawalTokensResponse) *NullableWeb3EmplWithdrawalTokensResponse {
	return &NullableWeb3EmplWithdrawalTokensResponse{value: val, isSet: true}
}

func (v NullableWeb3EmplWithdrawalTokensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3EmplWithdrawalTokensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


