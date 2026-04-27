
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// PalGrade the model 'PalGrade'
type PalGrade string

// List of PalGrade
const (
	PALGRADE_Egg PalGrade = "Egg"
	PALGRADE_Normal PalGrade = "Pal"
	PALGRADE_Super PalGrade = "Super Pal"
	PALGRADE_UltimatePal PalGrade = "Ultimate Pal"
)

// All allowed values of PalGrade enum
var AllowedPalGradeEnumValues = []PalGrade{
	"Egg",
	"Pal",
	"Super Pal",
	"Ultimate Pal",
}

func (v *PalGrade) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PalGrade(value)
	*v = enumTypeValue
	return nil
}

// NewPalGradeFromValue returns a pointer to a valid PalGrade
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPalGradeFromValue(v string) (*PalGrade, error) {
	ev := PalGrade(v)
	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PalGrade) IsValid() bool {
	for _, existing := range AllowedPalGradeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PalGrade value
func (v PalGrade) Ptr() *PalGrade {
	return &v
}

type NullablePalGrade struct {
	value *PalGrade
	isSet bool
}

func (v NullablePalGrade) Get() *PalGrade {
	return v.value
}

func (v *NullablePalGrade) Set(val *PalGrade) {
	v.value = val
	v.isSet = true
}

func (v NullablePalGrade) IsSet() bool {
	return v.isSet
}

func (v *NullablePalGrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalGrade(val *PalGrade) *NullablePalGrade {
	return &NullablePalGrade{value: val, isSet: true}
}

func (v NullablePalGrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalGrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

