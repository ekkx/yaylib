
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserInterestsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInterestsResponse{}

// UserInterestsResponse struct for UserInterestsResponse
type UserInterestsResponse struct {
	Interests []Interest `json:"interests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserInterestsResponse UserInterestsResponse

// NewUserInterestsResponse instantiates a new UserInterestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInterestsResponse() *UserInterestsResponse {
	this := UserInterestsResponse{}
	return &this
}

// NewUserInterestsResponseWithDefaults instantiates a new UserInterestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInterestsResponseWithDefaults() *UserInterestsResponse {
	this := UserInterestsResponse{}
	return &this
}

// GetInterests returns the Interests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInterestsResponse) GetInterests() []Interest {
	if o == nil {
		var ret []Interest
		return ret
	}
	return o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInterestsResponse) GetInterestsOk() ([]Interest, bool) {
	if o == nil || IsNil(o.Interests) {
		return nil, false
	}
	return o.Interests, true
}

// HasInterests returns a boolean if a field has been set.
func (o *UserInterestsResponse) HasInterests() bool {
	if o != nil && !IsNil(o.Interests) {
		return true
	}

	return false
}

// SetInterests gets a reference to the given []Interest and assigns it to the Interests field.
func (o *UserInterestsResponse) SetInterests(v []Interest) {
	o.Interests = v
}

func (o UserInterestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInterestsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Interests != nil {
		toSerialize["interests"] = o.Interests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserInterestsResponse) UnmarshalJSON(data []byte) (err error) {
	varUserInterestsResponse := _UserInterestsResponse{}

	err = json.Unmarshal(data, &varUserInterestsResponse)

	if err != nil {
		return err
	}

	*o = UserInterestsResponse(varUserInterestsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "interests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserInterestsResponse struct {
	value *UserInterestsResponse
	isSet bool
}

func (v NullableUserInterestsResponse) Get() *UserInterestsResponse {
	return v.value
}

func (v *NullableUserInterestsResponse) Set(val *UserInterestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInterestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInterestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInterestsResponse(val *UserInterestsResponse) *NullableUserInterestsResponse {
	return &NullableUserInterestsResponse{value: val, isSet: true}
}

func (v NullableUserInterestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInterestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


