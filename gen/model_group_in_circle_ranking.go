
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupInCircleRanking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupInCircleRanking{}

// GroupInCircleRanking struct for GroupInCircleRanking
type GroupInCircleRanking struct {
	Gifts []string `json:"gifts,omitempty"`
	User NullableUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupInCircleRanking GroupInCircleRanking

// NewGroupInCircleRanking instantiates a new GroupInCircleRanking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInCircleRanking() *GroupInCircleRanking {
	this := GroupInCircleRanking{}
	return &this
}

// NewGroupInCircleRankingWithDefaults instantiates a new GroupInCircleRanking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInCircleRankingWithDefaults() *GroupInCircleRanking {
	this := GroupInCircleRanking{}
	return &this
}

// GetGifts returns the Gifts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupInCircleRanking) GetGifts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Gifts
}

// GetGiftsOk returns a tuple with the Gifts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupInCircleRanking) GetGiftsOk() ([]string, bool) {
	if o == nil || IsNil(o.Gifts) {
		return nil, false
	}
	return o.Gifts, true
}

// HasGifts returns a boolean if a field has been set.
func (o *GroupInCircleRanking) HasGifts() bool {
	if o != nil && !IsNil(o.Gifts) {
		return true
	}

	return false
}

// SetGifts gets a reference to the given []string and assigns it to the Gifts field.
func (o *GroupInCircleRanking) SetGifts(v []string) {
	o.Gifts = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupInCircleRanking) GetUser() User {
	if o == nil || IsNil(o.User.Get()) {
		var ret User
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupInCircleRanking) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *GroupInCircleRanking) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUser and assigns it to the User field.
func (o *GroupInCircleRanking) SetUser(v User) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *GroupInCircleRanking) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *GroupInCircleRanking) UnsetUser() {
	o.User.Unset()
}

func (o GroupInCircleRanking) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupInCircleRanking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gifts != nil {
		toSerialize["gifts"] = o.Gifts
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupInCircleRanking) UnmarshalJSON(data []byte) (err error) {
	varGroupInCircleRanking := _GroupInCircleRanking{}

	err = json.Unmarshal(data, &varGroupInCircleRanking)

	if err != nil {
		return err
	}

	*o = GroupInCircleRanking(varGroupInCircleRanking)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gifts")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupInCircleRanking struct {
	value *GroupInCircleRanking
	isSet bool
}

func (v NullableGroupInCircleRanking) Get() *GroupInCircleRanking {
	return v.value
}

func (v *NullableGroupInCircleRanking) Set(val *GroupInCircleRanking) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInCircleRanking) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInCircleRanking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInCircleRanking(val *GroupInCircleRanking) *NullableGroupInCircleRanking {
	return &NullableGroupInCircleRanking{value: val, isSet: true}
}

func (v NullableGroupInCircleRanking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInCircleRanking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


