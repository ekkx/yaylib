
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// NoreplyMode the model 'NoreplyMode'
type NoreplyMode string

// List of NoreplyMode
const (
	NOREPLYMODE_False NoreplyMode = ""
	NOREPLYMODE_True NoreplyMode = "noreply_"
)

// All allowed values of NoreplyMode enum
var AllowedNoreplyModeEnumValues = []NoreplyMode{
	"",
	"noreply_",
}

func (v *NoreplyMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NoreplyMode(value)
	*v = enumTypeValue
	return nil
}

// NewNoreplyModeFromValue returns a pointer to a valid NoreplyMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNoreplyModeFromValue(v string) (*NoreplyMode, error) {
	ev := NoreplyMode(v)
	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NoreplyMode) IsValid() bool {
	for _, existing := range AllowedNoreplyModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NoreplyMode value
func (v NoreplyMode) Ptr() *NoreplyMode {
	return &v
}

type NullableNoreplyMode struct {
	value *NoreplyMode
	isSet bool
}

func (v NullableNoreplyMode) Get() *NoreplyMode {
	return v.value
}

func (v *NullableNoreplyMode) Set(val *NoreplyMode) {
	v.value = val
	v.isSet = true
}

func (v NullableNoreplyMode) IsSet() bool {
	return v.isSet
}

func (v *NullableNoreplyMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoreplyMode(val *NoreplyMode) *NullableNoreplyMode {
	return &NullableNoreplyMode{value: val, isSet: true}
}

func (v NullableNoreplyMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoreplyMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

