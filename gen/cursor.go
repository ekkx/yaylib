// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
	"fmt"
)

// NullableCursor is a NullableString-compatible nullable string whose
// UnmarshalJSON also accepts a JSON number, coercing it to its string
// form. The pagination cursor is a string in the schema but the
// server sends a bare number on some endpoints; this keeps the typed
// decode tolerant instead of failing the whole response.
type NullableCursor struct {
	value *string
	isSet bool
}

func (v NullableCursor) Get() *string      { return v.value }
func (v *NullableCursor) Set(val *string)  { v.value = val; v.isSet = true }
func (v NullableCursor) IsSet() bool       { return v.isSet }
func (v *NullableCursor) Unset()           { v.value = nil; v.isSet = false }

func NewNullableCursor(val *string) *NullableCursor {
	return &NullableCursor{value: val, isSet: true}
}

func (v NullableCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	if string(src) == "null" {
		v.value = nil
		return nil
	}
	var s string
	if err := json.Unmarshal(src, &s); err == nil {
		v.value = &s
		return nil
	}
	var n json.Number
	if err := json.Unmarshal(src, &n); err == nil {
		str := n.String()
		v.value = &str
		return nil
	}
	return fmt.Errorf("cursor: cannot decode %s into string", src)
}
