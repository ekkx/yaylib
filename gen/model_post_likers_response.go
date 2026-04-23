
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostLikersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLikersResponse{}

// PostLikersResponse struct for PostLikersResponse
type PostLikersResponse struct {
	LastId NullableInt64 `json:"last_id,omitempty"`
	Users []RealmUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostLikersResponse PostLikersResponse

// NewPostLikersResponse instantiates a new PostLikersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLikersResponse() *PostLikersResponse {
	this := PostLikersResponse{}
	return &this
}

// NewPostLikersResponseWithDefaults instantiates a new PostLikersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLikersResponseWithDefaults() *PostLikersResponse {
	this := PostLikersResponse{}
	return &this
}

// GetLastId returns the LastId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLikersResponse) GetLastId() int64 {
	if o == nil || IsNil(o.LastId.Get()) {
		var ret int64
		return ret
	}
	return *o.LastId.Get()
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLikersResponse) GetLastIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastId.Get(), o.LastId.IsSet()
}

// HasLastId returns a boolean if a field has been set.
func (o *PostLikersResponse) HasLastId() bool {
	if o != nil && o.LastId.IsSet() {
		return true
	}

	return false
}

// SetLastId gets a reference to the given NullableInt64 and assigns it to the LastId field.
func (o *PostLikersResponse) SetLastId(v int64) {
	o.LastId.Set(&v)
}
// SetLastIdNil sets the value for LastId to be an explicit nil
func (o *PostLikersResponse) SetLastIdNil() {
	o.LastId.Set(nil)
}

// UnsetLastId ensures that no value is present for LastId, not even an explicit nil
func (o *PostLikersResponse) UnsetLastId() {
	o.LastId.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostLikersResponse) GetUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostLikersResponse) GetUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *PostLikersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []RealmUser and assigns it to the Users field.
func (o *PostLikersResponse) SetUsers(v []RealmUser) {
	o.Users = v
}

func (o PostLikersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLikersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LastId.IsSet() {
		toSerialize["last_id"] = o.LastId.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostLikersResponse) UnmarshalJSON(data []byte) (err error) {
	varPostLikersResponse := _PostLikersResponse{}

	err = json.Unmarshal(data, &varPostLikersResponse)

	if err != nil {
		return err
	}

	*o = PostLikersResponse(varPostLikersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "last_id")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostLikersResponse struct {
	value *PostLikersResponse
	isSet bool
}

func (v NullablePostLikersResponse) Get() *PostLikersResponse {
	return v.value
}

func (v *NullablePostLikersResponse) Set(val *PostLikersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLikersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLikersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLikersResponse(val *PostLikersResponse) *NullablePostLikersResponse {
	return &NullablePostLikersResponse{value: val, isSet: true}
}

func (v NullablePostLikersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLikersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


