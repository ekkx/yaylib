
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

// Error struct for Error
type Error struct {
	// Unresolved Java type: jp.nanameue.yay.data.model.ErrorAction
	Action map[string]interface{} `json:"action,omitempty"`
	OriginalCode NullableInt32 `json:"original_code,omitempty"`
	// Unresolved Java type: java.lang.Throwable
	Throwable map[string]interface{} `json:"throwable,omitempty"`
	Type NullableErrorType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Error Error

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetAction() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetActionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Action) {
		return map[string]interface{}{}, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Error) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given map[string]interface{} and assigns it to the Action field.
func (o *Error) SetAction(v map[string]interface{}) {
	o.Action = v
}

// GetOriginalCode returns the OriginalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetOriginalCode() int32 {
	if o == nil || IsNil(o.OriginalCode.Get()) {
		var ret int32
		return ret
	}
	return *o.OriginalCode.Get()
}

// GetOriginalCodeOk returns a tuple with the OriginalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetOriginalCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalCode.Get(), o.OriginalCode.IsSet()
}

// HasOriginalCode returns a boolean if a field has been set.
func (o *Error) HasOriginalCode() bool {
	if o != nil && o.OriginalCode.IsSet() {
		return true
	}

	return false
}

// SetOriginalCode gets a reference to the given NullableInt32 and assigns it to the OriginalCode field.
func (o *Error) SetOriginalCode(v int32) {
	o.OriginalCode.Set(&v)
}
// SetOriginalCodeNil sets the value for OriginalCode to be an explicit nil
func (o *Error) SetOriginalCodeNil() {
	o.OriginalCode.Set(nil)
}

// UnsetOriginalCode ensures that no value is present for OriginalCode, not even an explicit nil
func (o *Error) UnsetOriginalCode() {
	o.OriginalCode.Unset()
}

// GetThrowable returns the Throwable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetThrowable() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Throwable
}

// GetThrowableOk returns a tuple with the Throwable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetThrowableOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Throwable) {
		return map[string]interface{}{}, false
	}
	return o.Throwable, true
}

// HasThrowable returns a boolean if a field has been set.
func (o *Error) HasThrowable() bool {
	if o != nil && !IsNil(o.Throwable) {
		return true
	}

	return false
}

// SetThrowable gets a reference to the given map[string]interface{} and assigns it to the Throwable field.
func (o *Error) SetThrowable(v map[string]interface{}) {
	o.Throwable = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetType() ErrorType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret ErrorType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetTypeOk() (*ErrorType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Error) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableErrorType and assigns it to the Type field.
func (o *Error) SetType(v ErrorType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Error) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Error) UnsetType() {
	o.Type.Unset()
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.OriginalCode.IsSet() {
		toSerialize["original_code"] = o.OriginalCode.Get()
	}
	if o.Throwable != nil {
		toSerialize["throwable"] = o.Throwable
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Error) UnmarshalJSON(data []byte) (err error) {
	varError := _Error{}

	err = json.Unmarshal(data, &varError)

	if err != nil {
		return err
	}

	*o = Error(varError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "original_code")
		delete(additionalProperties, "throwable")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


