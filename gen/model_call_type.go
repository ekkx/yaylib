
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// CallType the model 'CallType'
type CallType string

// List of CallType
const (
	CALLTYPE_Video CallType = "vdo"
	CALLTYPE_Voice CallType = "voice"
)

// All allowed values of CallType enum
var AllowedCallTypeEnumValues = []CallType{
	"vdo",
	"voice",
}

func (v *CallType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CallType(value)
	*v = enumTypeValue
	return nil
}

// NewCallTypeFromValue returns a pointer to a valid CallType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCallTypeFromValue(v string) (*CallType, error) {
	ev := CallType(v)
	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CallType) IsValid() bool {
	for _, existing := range AllowedCallTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CallType value
func (v CallType) Ptr() *CallType {
	return &v
}

type NullableCallType struct {
	value *CallType
	isSet bool
}

func (v NullableCallType) Get() *CallType {
	return v.value
}

func (v *NullableCallType) Set(val *CallType) {
	v.value = val
	v.isSet = true
}

func (v NullableCallType) IsSet() bool {
	return v.isSet
}

func (v *NullableCallType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallType(val *CallType) *NullableCallType {
	return &NullableCallType{value: val, isSet: true}
}

func (v NullableCallType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

