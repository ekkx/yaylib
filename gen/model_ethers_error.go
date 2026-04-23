
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EthersError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EthersError{}

// EthersError struct for EthersError
type EthersError struct {
	Error NullableString `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EthersError EthersError

// NewEthersError instantiates a new EthersError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEthersError() *EthersError {
	this := EthersError{}
	return &this
}

// NewEthersErrorWithDefaults instantiates a new EthersError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEthersErrorWithDefaults() *EthersError {
	this := EthersError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EthersError) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EthersError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *EthersError) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *EthersError) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *EthersError) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *EthersError) UnsetError() {
	o.Error.Unset()
}

func (o EthersError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EthersError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EthersError) UnmarshalJSON(data []byte) (err error) {
	varEthersError := _EthersError{}

	err = json.Unmarshal(data, &varEthersError)

	if err != nil {
		return err
	}

	*o = EthersError(varEthersError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEthersError struct {
	value *EthersError
	isSet bool
}

func (v NullableEthersError) Get() *EthersError {
	return v.value
}

func (v *NullableEthersError) Set(val *EthersError) {
	v.value = val
	v.isSet = true
}

func (v NullableEthersError) IsSet() bool {
	return v.isSet
}

func (v *NullableEthersError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthersError(val *EthersError) *NullableEthersError {
	return &NullableEthersError{value: val, isSet: true}
}

func (v NullableEthersError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthersError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


