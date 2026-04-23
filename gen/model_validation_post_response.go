
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ValidationPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationPostResponse{}

// ValidationPostResponse struct for ValidationPostResponse
type ValidationPostResponse struct {
	IsAllowToPost NullableBool `json:"is_allow_to_post,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidationPostResponse ValidationPostResponse

// NewValidationPostResponse instantiates a new ValidationPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationPostResponse() *ValidationPostResponse {
	this := ValidationPostResponse{}
	return &this
}

// NewValidationPostResponseWithDefaults instantiates a new ValidationPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationPostResponseWithDefaults() *ValidationPostResponse {
	this := ValidationPostResponse{}
	return &this
}

// GetIsAllowToPost returns the IsAllowToPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ValidationPostResponse) GetIsAllowToPost() bool {
	if o == nil || IsNil(o.IsAllowToPost.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAllowToPost.Get()
}

// GetIsAllowToPostOk returns a tuple with the IsAllowToPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ValidationPostResponse) GetIsAllowToPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAllowToPost.Get(), o.IsAllowToPost.IsSet()
}

// HasIsAllowToPost returns a boolean if a field has been set.
func (o *ValidationPostResponse) HasIsAllowToPost() bool {
	if o != nil && o.IsAllowToPost.IsSet() {
		return true
	}

	return false
}

// SetIsAllowToPost gets a reference to the given NullableBool and assigns it to the IsAllowToPost field.
func (o *ValidationPostResponse) SetIsAllowToPost(v bool) {
	o.IsAllowToPost.Set(&v)
}
// SetIsAllowToPostNil sets the value for IsAllowToPost to be an explicit nil
func (o *ValidationPostResponse) SetIsAllowToPostNil() {
	o.IsAllowToPost.Set(nil)
}

// UnsetIsAllowToPost ensures that no value is present for IsAllowToPost, not even an explicit nil
func (o *ValidationPostResponse) UnsetIsAllowToPost() {
	o.IsAllowToPost.Unset()
}

func (o ValidationPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsAllowToPost.IsSet() {
		toSerialize["is_allow_to_post"] = o.IsAllowToPost.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidationPostResponse) UnmarshalJSON(data []byte) (err error) {
	varValidationPostResponse := _ValidationPostResponse{}

	err = json.Unmarshal(data, &varValidationPostResponse)

	if err != nil {
		return err
	}

	*o = ValidationPostResponse(varValidationPostResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_allow_to_post")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidationPostResponse struct {
	value *ValidationPostResponse
	isSet bool
}

func (v NullableValidationPostResponse) Get() *ValidationPostResponse {
	return v.value
}

func (v *NullableValidationPostResponse) Set(val *ValidationPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationPostResponse(val *ValidationPostResponse) *NullableValidationPostResponse {
	return &NullableValidationPostResponse{value: val, isSet: true}
}

func (v NullableValidationPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


