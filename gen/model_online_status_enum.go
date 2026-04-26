
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
	"fmt"
)

// OnlineStatusEnum the model 'OnlineStatusEnum'
type OnlineStatusEnum string

// List of OnlineStatusEnum
const (
	ONLINESTATUSENUM_OFFLINE OnlineStatusEnum = "offline"
	ONLINESTATUSENUM_HIDDEN OnlineStatusEnum = "hidden"
)

// All allowed values of OnlineStatusEnum enum
var AllowedOnlineStatusEnumEnumValues = []OnlineStatusEnum{
	"offline",
	"hidden",
}

func (v *OnlineStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OnlineStatusEnum(value)
	for _, existing := range AllowedOnlineStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OnlineStatusEnum", value)
}

// NewOnlineStatusEnumFromValue returns a pointer to a valid OnlineStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOnlineStatusEnumFromValue(v string) (*OnlineStatusEnum, error) {
	ev := OnlineStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OnlineStatusEnum: valid values are %v", v, AllowedOnlineStatusEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OnlineStatusEnum) IsValid() bool {
	for _, existing := range AllowedOnlineStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OnlineStatusEnum value
func (v OnlineStatusEnum) Ptr() *OnlineStatusEnum {
	return &v
}

type NullableOnlineStatusEnum struct {
	value *OnlineStatusEnum
	isSet bool
}

func (v NullableOnlineStatusEnum) Get() *OnlineStatusEnum {
	return v.value
}

func (v *NullableOnlineStatusEnum) Set(val *OnlineStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableOnlineStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableOnlineStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnlineStatusEnum(val *OnlineStatusEnum) *NullableOnlineStatusEnum {
	return &NullableOnlineStatusEnum{value: val, isSet: true}
}

func (v NullableOnlineStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnlineStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

