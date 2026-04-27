
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// MessageType the model 'MessageType'
type MessageType string

// List of MessageType
const (
	MESSAGETYPE_Text MessageType = "text"
	MESSAGETYPE_Image MessageType = "image"
	MESSAGETYPE_EternalImage MessageType = "eternal_image"
	MESSAGETYPE_Video MessageType = "video"
	MESSAGETYPE_EternalVideo MessageType = "eternal_video"
	MESSAGETYPE_Screenshot MessageType = "screenshot_warning"
	MESSAGETYPE_Gif MessageType = "gif"
	MESSAGETYPE_Sticker MessageType = "sticker"
	MESSAGETYPE_PrivateCall MessageType = "individual_call"
	MESSAGETYPE_Deleted MessageType = "deleted"
	MESSAGETYPE_UserJoined MessageType = "user_joined"
	MESSAGETYPE_UserLeave MessageType = "user_leave"
	MESSAGETYPE_OwnerKick MessageType = "owner_kick"
)

// All allowed values of MessageType enum
var AllowedMessageTypeEnumValues = []MessageType{
	"text",
	"image",
	"eternal_image",
	"video",
	"eternal_video",
	"screenshot_warning",
	"gif",
	"sticker",
	"individual_call",
	"deleted",
	"user_joined",
	"user_leave",
	"owner_kick",
}

func (v *MessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MessageType(value)
	*v = enumTypeValue
	return nil
}

// NewMessageTypeFromValue returns a pointer to a valid MessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMessageTypeFromValue(v string) (*MessageType, error) {
	ev := MessageType(v)
	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MessageType) IsValid() bool {
	for _, existing := range AllowedMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MessageType value
func (v MessageType) Ptr() *MessageType {
	return &v
}

type NullableMessageType struct {
	value *MessageType
	isSet bool
}

func (v NullableMessageType) Get() *MessageType {
	return v.value
}

func (v *NullableMessageType) Set(val *MessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageType(val *MessageType) *NullableMessageType {
	return &NullableMessageType{value: val, isSet: true}
}

func (v NullableMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

