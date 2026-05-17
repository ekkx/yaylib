
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersResponse{}

// UsersResponse struct for UsersResponse
type UsersResponse struct {
	NextPageValue NullableString `json:"next_page_value,omitempty"`
	Users []RealmUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UsersResponse UsersResponse

// NewUsersResponse instantiates a new UsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersResponse() *UsersResponse {
	this := UsersResponse{}
	return &this
}

// NewUsersResponseWithDefaults instantiates a new UsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersResponseWithDefaults() *UsersResponse {
	this := UsersResponse{}
	return &this
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *UsersResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *UsersResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *UsersResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *UsersResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsersResponse) GetUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsersResponse) GetUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []RealmUser and assigns it to the Users field.
func (o *UsersResponse) SetUsers(v []RealmUser) {
	o.Users = v
}

func (o UsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersResponse) ToMap() (map[string]interface{}, error) {
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

func (o *UsersResponse) UnmarshalJSON(data []byte) (err error) {
	varUsersResponse := _UsersResponse{}

	err = json.Unmarshal(data, &varUsersResponse)

	if err != nil {
		return err
	}

	*o = UsersResponse(varUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsersResponse struct {
	value *UsersResponse
	isSet bool
}

func (v NullableUsersResponse) Get() *UsersResponse {
	return v.value
}

func (v *NullableUsersResponse) Set(val *UsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersResponse(val *UsersResponse) *NullableUsersResponse {
	return &NullableUsersResponse{value: val, isSet: true}
}

func (v NullableUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


