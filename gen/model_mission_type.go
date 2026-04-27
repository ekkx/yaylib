
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// MissionType the model 'MissionType'
type MissionType string

// List of MissionType
const (
	MISSIONTYPE_Daily MissionType = "daily"
	MISSIONTYPE_Special MissionType = "special"
	MISSIONTYPE_Endless MissionType = "unlimited"
	MISSIONTYPE_Normal MissionType = "once"
)

// All allowed values of MissionType enum
var AllowedMissionTypeEnumValues = []MissionType{
	"daily",
	"special",
	"unlimited",
	"once",
}

func (v *MissionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MissionType(value)
	*v = enumTypeValue
	return nil
}

// NewMissionTypeFromValue returns a pointer to a valid MissionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMissionTypeFromValue(v string) (*MissionType, error) {
	ev := MissionType(v)
	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MissionType) IsValid() bool {
	for _, existing := range AllowedMissionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MissionType value
func (v MissionType) Ptr() *MissionType {
	return &v
}

type NullableMissionType struct {
	value *MissionType
	isSet bool
}

func (v NullableMissionType) Get() *MissionType {
	return v.value
}

func (v *NullableMissionType) Set(val *MissionType) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionType) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionType(val *MissionType) *NullableMissionType {
	return &NullableMissionType{value: val, isSet: true}
}

func (v NullableMissionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

