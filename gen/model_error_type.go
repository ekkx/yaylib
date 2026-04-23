
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ErrorType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorType{}

// ErrorType struct for ErrorType
type ErrorType struct {
	Code NullableInt32 `json:"code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ErrorType ErrorType

// NewErrorType instantiates a new ErrorType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorType() *ErrorType {
	this := ErrorType{}
	return &this
}

// NewErrorTypeWithDefaults instantiates a new ErrorType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorTypeWithDefaults() *ErrorType {
	this := ErrorType{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorType) GetCode() int32 {
	if o == nil || IsNil(o.Code.Get()) {
		var ret int32
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorType) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorType) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableInt32 and assigns it to the Code field.
func (o *ErrorType) SetCode(v int32) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *ErrorType) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ErrorType) UnsetCode() {
	o.Code.Unset()
}

func (o ErrorType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorType) UnmarshalJSON(data []byte) (err error) {
	varErrorType := _ErrorType{}

	err = json.Unmarshal(data, &varErrorType)

	if err != nil {
		return err
	}

	*o = ErrorType(varErrorType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorType struct {
	value *ErrorType
	isSet bool
}

func (v NullableErrorType) Get() *ErrorType {
	return v.value
}

func (v *NullableErrorType) Set(val *ErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorType(val *ErrorType) *NullableErrorType {
	return &NullableErrorType{value: val, isSet: true}
}

func (v NullableErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


