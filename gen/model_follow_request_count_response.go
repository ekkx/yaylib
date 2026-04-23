
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FollowRequestCountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowRequestCountResponse{}

// FollowRequestCountResponse struct for FollowRequestCountResponse
type FollowRequestCountResponse struct {
	UsersCount NullableInt32 `json:"users_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FollowRequestCountResponse FollowRequestCountResponse

// NewFollowRequestCountResponse instantiates a new FollowRequestCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowRequestCountResponse() *FollowRequestCountResponse {
	this := FollowRequestCountResponse{}
	return &this
}

// NewFollowRequestCountResponseWithDefaults instantiates a new FollowRequestCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowRequestCountResponseWithDefaults() *FollowRequestCountResponse {
	this := FollowRequestCountResponse{}
	return &this
}

// GetUsersCount returns the UsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FollowRequestCountResponse) GetUsersCount() int32 {
	if o == nil || IsNil(o.UsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UsersCount.Get()
}

// GetUsersCountOk returns a tuple with the UsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FollowRequestCountResponse) GetUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsersCount.Get(), o.UsersCount.IsSet()
}

// HasUsersCount returns a boolean if a field has been set.
func (o *FollowRequestCountResponse) HasUsersCount() bool {
	if o != nil && o.UsersCount.IsSet() {
		return true
	}

	return false
}

// SetUsersCount gets a reference to the given NullableInt32 and assigns it to the UsersCount field.
func (o *FollowRequestCountResponse) SetUsersCount(v int32) {
	o.UsersCount.Set(&v)
}
// SetUsersCountNil sets the value for UsersCount to be an explicit nil
func (o *FollowRequestCountResponse) SetUsersCountNil() {
	o.UsersCount.Set(nil)
}

// UnsetUsersCount ensures that no value is present for UsersCount, not even an explicit nil
func (o *FollowRequestCountResponse) UnsetUsersCount() {
	o.UsersCount.Unset()
}

func (o FollowRequestCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowRequestCountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UsersCount.IsSet() {
		toSerialize["users_count"] = o.UsersCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FollowRequestCountResponse) UnmarshalJSON(data []byte) (err error) {
	varFollowRequestCountResponse := _FollowRequestCountResponse{}

	err = json.Unmarshal(data, &varFollowRequestCountResponse)

	if err != nil {
		return err
	}

	*o = FollowRequestCountResponse(varFollowRequestCountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "users_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFollowRequestCountResponse struct {
	value *FollowRequestCountResponse
	isSet bool
}

func (v NullableFollowRequestCountResponse) Get() *FollowRequestCountResponse {
	return v.value
}

func (v *NullableFollowRequestCountResponse) Set(val *FollowRequestCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowRequestCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowRequestCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowRequestCountResponse(val *FollowRequestCountResponse) *NullableFollowRequestCountResponse {
	return &NullableFollowRequestCountResponse{value: val, isSet: true}
}

func (v NullableFollowRequestCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowRequestCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


