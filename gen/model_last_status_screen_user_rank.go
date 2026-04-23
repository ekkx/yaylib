
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LastStatusScreenUserRank type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastStatusScreenUserRank{}

// LastStatusScreenUserRank struct for LastStatusScreenUserRank
type LastStatusScreenUserRank struct {
	UserId NullableInt64 `json:"user_id,omitempty"`
	UserRank NullableString `json:"user_rank,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LastStatusScreenUserRank LastStatusScreenUserRank

// NewLastStatusScreenUserRank instantiates a new LastStatusScreenUserRank object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastStatusScreenUserRank() *LastStatusScreenUserRank {
	this := LastStatusScreenUserRank{}
	return &this
}

// NewLastStatusScreenUserRankWithDefaults instantiates a new LastStatusScreenUserRank object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastStatusScreenUserRankWithDefaults() *LastStatusScreenUserRank {
	this := LastStatusScreenUserRank{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LastStatusScreenUserRank) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LastStatusScreenUserRank) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *LastStatusScreenUserRank) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *LastStatusScreenUserRank) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *LastStatusScreenUserRank) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *LastStatusScreenUserRank) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserRank returns the UserRank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LastStatusScreenUserRank) GetUserRank() string {
	if o == nil || IsNil(o.UserRank.Get()) {
		var ret string
		return ret
	}
	return *o.UserRank.Get()
}

// GetUserRankOk returns a tuple with the UserRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LastStatusScreenUserRank) GetUserRankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserRank.Get(), o.UserRank.IsSet()
}

// HasUserRank returns a boolean if a field has been set.
func (o *LastStatusScreenUserRank) HasUserRank() bool {
	if o != nil && o.UserRank.IsSet() {
		return true
	}

	return false
}

// SetUserRank gets a reference to the given NullableString and assigns it to the UserRank field.
func (o *LastStatusScreenUserRank) SetUserRank(v string) {
	o.UserRank.Set(&v)
}
// SetUserRankNil sets the value for UserRank to be an explicit nil
func (o *LastStatusScreenUserRank) SetUserRankNil() {
	o.UserRank.Set(nil)
}

// UnsetUserRank ensures that no value is present for UserRank, not even an explicit nil
func (o *LastStatusScreenUserRank) UnsetUserRank() {
	o.UserRank.Unset()
}

func (o LastStatusScreenUserRank) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastStatusScreenUserRank) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	if o.UserRank.IsSet() {
		toSerialize["user_rank"] = o.UserRank.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LastStatusScreenUserRank) UnmarshalJSON(data []byte) (err error) {
	varLastStatusScreenUserRank := _LastStatusScreenUserRank{}

	err = json.Unmarshal(data, &varLastStatusScreenUserRank)

	if err != nil {
		return err
	}

	*o = LastStatusScreenUserRank(varLastStatusScreenUserRank)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "user_rank")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLastStatusScreenUserRank struct {
	value *LastStatusScreenUserRank
	isSet bool
}

func (v NullableLastStatusScreenUserRank) Get() *LastStatusScreenUserRank {
	return v.value
}

func (v *NullableLastStatusScreenUserRank) Set(val *LastStatusScreenUserRank) {
	v.value = val
	v.isSet = true
}

func (v NullableLastStatusScreenUserRank) IsSet() bool {
	return v.isSet
}

func (v *NullableLastStatusScreenUserRank) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastStatusScreenUserRank(val *LastStatusScreenUserRank) *NullableLastStatusScreenUserRank {
	return &NullableLastStatusScreenUserRank{value: val, isSet: true}
}

func (v NullableLastStatusScreenUserRank) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastStatusScreenUserRank) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


