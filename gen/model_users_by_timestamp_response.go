
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UsersByTimestampResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersByTimestampResponse{}

// UsersByTimestampResponse struct for UsersByTimestampResponse
type UsersByTimestampResponse struct {
	LastTimestamp NullableInt64 `json:"last_timestamp,omitempty"`
	Users []RealmUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersByTimestampResponse UsersByTimestampResponse

// NewUsersByTimestampResponse instantiates a new UsersByTimestampResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersByTimestampResponse() *UsersByTimestampResponse {
	this := UsersByTimestampResponse{}
	return &this
}

// NewUsersByTimestampResponseWithDefaults instantiates a new UsersByTimestampResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersByTimestampResponseWithDefaults() *UsersByTimestampResponse {
	this := UsersByTimestampResponse{}
	return &this
}

// GetLastTimestamp returns the LastTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersByTimestampResponse) GetLastTimestamp() int64 {
	if o == nil || IsNil(o.LastTimestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.LastTimestamp.Get()
}

// GetLastTimestampOk returns a tuple with the LastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersByTimestampResponse) GetLastTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTimestamp.Get(), o.LastTimestamp.IsSet()
}

// HasLastTimestamp returns a boolean if a field has been set.
func (o *UsersByTimestampResponse) HasLastTimestamp() bool {
	if o != nil && o.LastTimestamp.IsSet() {
		return true
	}

	return false
}

// SetLastTimestamp gets a reference to the given NullableInt64 and assigns it to the LastTimestamp field.
func (o *UsersByTimestampResponse) SetLastTimestamp(v int64) {
	o.LastTimestamp.Set(&v)
}
// SetLastTimestampNil sets the value for LastTimestamp to be an explicit nil
func (o *UsersByTimestampResponse) SetLastTimestampNil() {
	o.LastTimestamp.Set(nil)
}

// UnsetLastTimestamp ensures that no value is present for LastTimestamp, not even an explicit nil
func (o *UsersByTimestampResponse) UnsetLastTimestamp() {
	o.LastTimestamp.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersByTimestampResponse) GetUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersByTimestampResponse) GetUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UsersByTimestampResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []RealmUser and assigns it to the Users field.
func (o *UsersByTimestampResponse) SetUsers(v []RealmUser) {
	o.Users = v
}

func (o UsersByTimestampResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersByTimestampResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LastTimestamp.IsSet() {
		toSerialize["last_timestamp"] = o.LastTimestamp.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsersByTimestampResponse) UnmarshalJSON(data []byte) (err error) {
	varUsersByTimestampResponse := _UsersByTimestampResponse{}

	err = json.Unmarshal(data, &varUsersByTimestampResponse)

	if err != nil {
		return err
	}

	*o = UsersByTimestampResponse(varUsersByTimestampResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "last_timestamp")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersByTimestampResponse struct {
	value *UsersByTimestampResponse
	isSet bool
}

func (v NullableUsersByTimestampResponse) Get() *UsersByTimestampResponse {
	return v.value
}

func (v *NullableUsersByTimestampResponse) Set(val *UsersByTimestampResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersByTimestampResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersByTimestampResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersByTimestampResponse(val *UsersByTimestampResponse) *NullableUsersByTimestampResponse {
	return &NullableUsersByTimestampResponse{value: val, isSet: true}
}

func (v NullableUsersByTimestampResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersByTimestampResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


