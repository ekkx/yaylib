
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
	"fmt"
)

// PostType the model 'PostType'
type PostType string

// List of PostType
const (
	POSTTYPE_CALL PostType = "call"
	POSTTYPE_IMAGE PostType = "image"
	POSTTYPE_VIDEO PostType = "video"
	POSTTYPE_SURVEY PostType = "survey"
	POSTTYPE_REPOST PostType = "repost"
	POSTTYPE_THREAD PostType = "thread"
	POSTTYPE_SHAREABLE_GROUP PostType = "shareable_group"
	POSTTYPE_SHAREABLE_URL PostType = "shareable_url"
	POSTTYPE_YOUTUBE PostType = "youtube"
	POSTTYPE_SHAREABLE_THREAD PostType = "shareable_thread"
	POSTTYPE_SHAREPAL PostType = "sharepal"
	POSTTYPE_UNDEFINED PostType = "undefined"
)

// All allowed values of PostType enum
var AllowedPostTypeEnumValues = []PostType{
	"call",
	"image",
	"video",
	"survey",
	"repost",
	"thread",
	"shareable_group",
	"shareable_url",
	"youtube",
	"shareable_thread",
	"sharepal",
	"undefined",
}

func (v *PostType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PostType(value)
	for _, existing := range AllowedPostTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PostType", value)
}

// NewPostTypeFromValue returns a pointer to a valid PostType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPostTypeFromValue(v string) (*PostType, error) {
	ev := PostType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PostType: valid values are %v", v, AllowedPostTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PostType) IsValid() bool {
	for _, existing := range AllowedPostTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostType value
func (v PostType) Ptr() *PostType {
	return &v
}

type NullablePostType struct {
	value *PostType
	isSet bool
}

func (v NullablePostType) Get() *PostType {
	return v.value
}

func (v *NullablePostType) Set(val *PostType) {
	v.value = val
	v.isSet = true
}

func (v NullablePostType) IsSet() bool {
	return v.isSet
}

func (v *NullablePostType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostType(val *PostType) *NullablePostType {
	return &NullablePostType{value: val, isSet: true}
}

func (v NullablePostType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

