
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FollowUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowUsersResponse{}

// FollowUsersResponse struct for FollowUsersResponse
type FollowUsersResponse struct {
	NextPageValue NullableString `json:"next_page_value,omitempty"`
	Users []RealmUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FollowUsersResponse FollowUsersResponse

// NewFollowUsersResponse instantiates a new FollowUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowUsersResponse() *FollowUsersResponse {
	this := FollowUsersResponse{}
	return &this
}

// NewFollowUsersResponseWithDefaults instantiates a new FollowUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowUsersResponseWithDefaults() *FollowUsersResponse {
	this := FollowUsersResponse{}
	return &this
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FollowUsersResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FollowUsersResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *FollowUsersResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *FollowUsersResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *FollowUsersResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *FollowUsersResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FollowUsersResponse) GetUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FollowUsersResponse) GetUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *FollowUsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []RealmUser and assigns it to the Users field.
func (o *FollowUsersResponse) SetUsers(v []RealmUser) {
	o.Users = v
}

func (o FollowUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FollowUsersResponse) UnmarshalJSON(data []byte) (err error) {
	varFollowUsersResponse := _FollowUsersResponse{}

	err = json.Unmarshal(data, &varFollowUsersResponse)

	if err != nil {
		return err
	}

	*o = FollowUsersResponse(varFollowUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFollowUsersResponse struct {
	value *FollowUsersResponse
	isSet bool
}

func (v NullableFollowUsersResponse) Get() *FollowUsersResponse {
	return v.value
}

func (v *NullableFollowUsersResponse) Set(val *FollowUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowUsersResponse(val *FollowUsersResponse) *NullableFollowUsersResponse {
	return &NullableFollowUsersResponse{value: val, isSet: true}
}

func (v NullableFollowUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


