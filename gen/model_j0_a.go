
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// J0A the model 'J0A'
type J0A string

// List of j0_a
const (
	J0A_PUBLIC_TIMELINE J0A = "PublicTimeline"
	J0A_FOLLOWING_TIMELINE J0A = "FollowingTimeline"
	J0A_GROUP_TIMELINE J0A = "GroupTimeline"
)

// All allowed values of J0A enum
var AllowedJ0AEnumValues = []J0A{
	"PublicTimeline",
	"FollowingTimeline",
	"GroupTimeline",
}

func (v *J0A) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := J0A(value)
	*v = enumTypeValue
	return nil
}

// NewJ0AFromValue returns a pointer to a valid J0A
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJ0AFromValue(v string) (*J0A, error) {
	ev := J0A(v)
	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v J0A) IsValid() bool {
	for _, existing := range AllowedJ0AEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to j0_a value
func (v J0A) Ptr() *J0A {
	return &v
}

type NullableJ0A struct {
	value *J0A
	isSet bool
}

func (v NullableJ0A) Get() *J0A {
	return v.value
}

func (v *NullableJ0A) Set(val *J0A) {
	v.value = val
	v.isSet = true
}

func (v NullableJ0A) IsSet() bool {
	return v.isSet
}

func (v *NullableJ0A) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJ0A(val *J0A) *NullableJ0A {
	return &NullableJ0A{value: val, isSet: true}
}

func (v NullableJ0A) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJ0A) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

