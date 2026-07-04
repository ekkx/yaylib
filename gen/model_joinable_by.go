
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// JoinableBy the model 'JoinableBy'
type JoinableBy string

// List of JoinableBy
const (
	JOINABLEBY_Anyone JoinableBy = "anyone"
	JOINABLEBY_Users JoinableBy = "users"
	JOINABLEBY_FollowersOnly JoinableBy = "followers_only"
	JOINABLEBY_MutualFollowsOnly JoinableBy = "mutual_follows_only"
	JOINABLEBY_InviteOnly JoinableBy = "invite_only"
	JOINABLEBY_NoOne JoinableBy = "noone"
)

// All allowed values of JoinableBy enum
var AllowedJoinableByEnumValues = []JoinableBy{
	"anyone",
	"users",
	"followers_only",
	"mutual_follows_only",
	"invite_only",
	"noone",
}

func (v *JoinableBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JoinableBy(value)
	*v = enumTypeValue
	return nil
}

// NewJoinableByFromValue returns a pointer to a valid JoinableBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJoinableByFromValue(v string) (*JoinableBy, error) {
	ev := JoinableBy(v)
	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JoinableBy) IsValid() bool {
	for _, existing := range AllowedJoinableByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JoinableBy value
func (v JoinableBy) Ptr() *JoinableBy {
	return &v
}

type NullableJoinableBy struct {
	value *JoinableBy
	isSet bool
}

func (v NullableJoinableBy) Get() *JoinableBy {
	return v.value
}

func (v *NullableJoinableBy) Set(val *JoinableBy) {
	v.value = val
	v.isSet = true
}

func (v NullableJoinableBy) IsSet() bool {
	return v.isSet
}

func (v *NullableJoinableBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJoinableBy(val *JoinableBy) *NullableJoinableBy {
	return &NullableJoinableBy{value: val, isSet: true}
}

func (v NullableJoinableBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJoinableBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

