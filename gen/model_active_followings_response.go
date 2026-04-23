
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ActiveFollowingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveFollowingsResponse{}

// ActiveFollowingsResponse struct for ActiveFollowingsResponse
type ActiveFollowingsResponse struct {
	LastLoggedinAt NullableInt64 `json:"last_loggedin_at,omitempty"`
	Users []RealmUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ActiveFollowingsResponse ActiveFollowingsResponse

// NewActiveFollowingsResponse instantiates a new ActiveFollowingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveFollowingsResponse() *ActiveFollowingsResponse {
	this := ActiveFollowingsResponse{}
	return &this
}

// NewActiveFollowingsResponseWithDefaults instantiates a new ActiveFollowingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveFollowingsResponseWithDefaults() *ActiveFollowingsResponse {
	this := ActiveFollowingsResponse{}
	return &this
}

// GetLastLoggedinAt returns the LastLoggedinAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveFollowingsResponse) GetLastLoggedinAt() int64 {
	if o == nil || IsNil(o.LastLoggedinAt.Get()) {
		var ret int64
		return ret
	}
	return *o.LastLoggedinAt.Get()
}

// GetLastLoggedinAtOk returns a tuple with the LastLoggedinAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveFollowingsResponse) GetLastLoggedinAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoggedinAt.Get(), o.LastLoggedinAt.IsSet()
}

// HasLastLoggedinAt returns a boolean if a field has been set.
func (o *ActiveFollowingsResponse) HasLastLoggedinAt() bool {
	if o != nil && o.LastLoggedinAt.IsSet() {
		return true
	}

	return false
}

// SetLastLoggedinAt gets a reference to the given NullableInt64 and assigns it to the LastLoggedinAt field.
func (o *ActiveFollowingsResponse) SetLastLoggedinAt(v int64) {
	o.LastLoggedinAt.Set(&v)
}
// SetLastLoggedinAtNil sets the value for LastLoggedinAt to be an explicit nil
func (o *ActiveFollowingsResponse) SetLastLoggedinAtNil() {
	o.LastLoggedinAt.Set(nil)
}

// UnsetLastLoggedinAt ensures that no value is present for LastLoggedinAt, not even an explicit nil
func (o *ActiveFollowingsResponse) UnsetLastLoggedinAt() {
	o.LastLoggedinAt.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveFollowingsResponse) GetUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveFollowingsResponse) GetUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ActiveFollowingsResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []RealmUser and assigns it to the Users field.
func (o *ActiveFollowingsResponse) SetUsers(v []RealmUser) {
	o.Users = v
}

func (o ActiveFollowingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveFollowingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LastLoggedinAt.IsSet() {
		toSerialize["last_loggedin_at"] = o.LastLoggedinAt.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ActiveFollowingsResponse) UnmarshalJSON(data []byte) (err error) {
	varActiveFollowingsResponse := _ActiveFollowingsResponse{}

	err = json.Unmarshal(data, &varActiveFollowingsResponse)

	if err != nil {
		return err
	}

	*o = ActiveFollowingsResponse(varActiveFollowingsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "last_loggedin_at")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActiveFollowingsResponse struct {
	value *ActiveFollowingsResponse
	isSet bool
}

func (v NullableActiveFollowingsResponse) Get() *ActiveFollowingsResponse {
	return v.value
}

func (v *NullableActiveFollowingsResponse) Set(val *ActiveFollowingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveFollowingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveFollowingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveFollowingsResponse(val *ActiveFollowingsResponse) *NullableActiveFollowingsResponse {
	return &NullableActiveFollowingsResponse{value: val, isSet: true}
}

func (v NullableActiveFollowingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveFollowingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


