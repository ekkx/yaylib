
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateFriendShipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFriendShipsResponse{}

// CreateFriendShipsResponse struct for CreateFriendShipsResponse
type CreateFriendShipsResponse struct {
	Following NullableBool `json:"following,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateFriendShipsResponse CreateFriendShipsResponse

// NewCreateFriendShipsResponse instantiates a new CreateFriendShipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFriendShipsResponse() *CreateFriendShipsResponse {
	this := CreateFriendShipsResponse{}
	return &this
}

// NewCreateFriendShipsResponseWithDefaults instantiates a new CreateFriendShipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFriendShipsResponseWithDefaults() *CreateFriendShipsResponse {
	this := CreateFriendShipsResponse{}
	return &this
}

// GetFollowing returns the Following field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateFriendShipsResponse) GetFollowing() bool {
	if o == nil || IsNil(o.Following.Get()) {
		var ret bool
		return ret
	}
	return *o.Following.Get()
}

// GetFollowingOk returns a tuple with the Following field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFriendShipsResponse) GetFollowingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Following.Get(), o.Following.IsSet()
}

// HasFollowing returns a boolean if a field has been set.
func (o *CreateFriendShipsResponse) HasFollowing() bool {
	if o != nil && o.Following.IsSet() {
		return true
	}

	return false
}

// SetFollowing gets a reference to the given NullableBool and assigns it to the Following field.
func (o *CreateFriendShipsResponse) SetFollowing(v bool) {
	o.Following.Set(&v)
}
// SetFollowingNil sets the value for Following to be an explicit nil
func (o *CreateFriendShipsResponse) SetFollowingNil() {
	o.Following.Set(nil)
}

// UnsetFollowing ensures that no value is present for Following, not even an explicit nil
func (o *CreateFriendShipsResponse) UnsetFollowing() {
	o.Following.Unset()
}

func (o CreateFriendShipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFriendShipsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Following.IsSet() {
		toSerialize["following"] = o.Following.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFriendShipsResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateFriendShipsResponse := _CreateFriendShipsResponse{}

	err = json.Unmarshal(data, &varCreateFriendShipsResponse)

	if err != nil {
		return err
	}

	*o = CreateFriendShipsResponse(varCreateFriendShipsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "following")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFriendShipsResponse struct {
	value *CreateFriendShipsResponse
	isSet bool
}

func (v NullableCreateFriendShipsResponse) Get() *CreateFriendShipsResponse {
	return v.value
}

func (v *NullableCreateFriendShipsResponse) Set(val *CreateFriendShipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFriendShipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFriendShipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFriendShipsResponse(val *CreateFriendShipsResponse) *NullableCreateFriendShipsResponse {
	return &NullableCreateFriendShipsResponse{value: val, isSet: true}
}

func (v NullableCreateFriendShipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFriendShipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


