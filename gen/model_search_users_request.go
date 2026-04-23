
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SearchUsersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchUsersRequest{}

// SearchUsersRequest struct for SearchUsersRequest
type SearchUsersRequest struct {
	User NullableSearchCriteria `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchUsersRequest SearchUsersRequest

// NewSearchUsersRequest instantiates a new SearchUsersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchUsersRequest() *SearchUsersRequest {
	this := SearchUsersRequest{}
	return &this
}

// NewSearchUsersRequestWithDefaults instantiates a new SearchUsersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchUsersRequestWithDefaults() *SearchUsersRequest {
	this := SearchUsersRequest{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchUsersRequest) GetUser() SearchCriteria {
	if o == nil || IsNil(o.User.Get()) {
		var ret SearchCriteria
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchUsersRequest) GetUserOk() (*SearchCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *SearchUsersRequest) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableSearchCriteria and assigns it to the User field.
func (o *SearchUsersRequest) SetUser(v SearchCriteria) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *SearchUsersRequest) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *SearchUsersRequest) UnsetUser() {
	o.User.Unset()
}

func (o SearchUsersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchUsersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchUsersRequest) UnmarshalJSON(data []byte) (err error) {
	varSearchUsersRequest := _SearchUsersRequest{}

	err = json.Unmarshal(data, &varSearchUsersRequest)

	if err != nil {
		return err
	}

	*o = SearchUsersRequest(varSearchUsersRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchUsersRequest struct {
	value *SearchUsersRequest
	isSet bool
}

func (v NullableSearchUsersRequest) Get() *SearchUsersRequest {
	return v.value
}

func (v *NullableSearchUsersRequest) Set(val *SearchUsersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchUsersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchUsersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchUsersRequest(val *SearchUsersRequest) *NullableSearchUsersRequest {
	return &NullableSearchUsersRequest{value: val, isSet: true}
}

func (v NullableSearchUsersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchUsersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


