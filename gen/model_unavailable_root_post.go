
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UnavailableRootPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnavailableRootPost{}

// UnavailableRootPost struct for UnavailableRootPost
type UnavailableRootPost struct {
	Post NullableModelPost `json:"post,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UnavailableRootPost UnavailableRootPost

// NewUnavailableRootPost instantiates a new UnavailableRootPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnavailableRootPost() *UnavailableRootPost {
	this := UnavailableRootPost{}
	return &this
}

// NewUnavailableRootPostWithDefaults instantiates a new UnavailableRootPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnavailableRootPostWithDefaults() *UnavailableRootPost {
	this := UnavailableRootPost{}
	return &this
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnavailableRootPost) GetPost() ModelPost {
	if o == nil || IsNil(o.Post.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnavailableRootPost) GetPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *UnavailableRootPost) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullableModelPost and assigns it to the Post field.
func (o *UnavailableRootPost) SetPost(v ModelPost) {
	o.Post.Set(&v)
}
// SetPostNil sets the value for Post to be an explicit nil
func (o *UnavailableRootPost) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *UnavailableRootPost) UnsetPost() {
	o.Post.Unset()
}

func (o UnavailableRootPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnavailableRootPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Post.IsSet() {
		toSerialize["post"] = o.Post.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UnavailableRootPost) UnmarshalJSON(data []byte) (err error) {
	varUnavailableRootPost := _UnavailableRootPost{}

	err = json.Unmarshal(data, &varUnavailableRootPost)

	if err != nil {
		return err
	}

	*o = UnavailableRootPost(varUnavailableRootPost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "post")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnavailableRootPost struct {
	value *UnavailableRootPost
	isSet bool
}

func (v NullableUnavailableRootPost) Get() *UnavailableRootPost {
	return v.value
}

func (v *NullableUnavailableRootPost) Set(val *UnavailableRootPost) {
	v.value = val
	v.isSet = true
}

func (v NullableUnavailableRootPost) IsSet() bool {
	return v.isSet
}

func (v *NullableUnavailableRootPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnavailableRootPost(val *UnavailableRootPost) *NullableUnavailableRootPost {
	return &NullableUnavailableRootPost{value: val, isSet: true}
}

func (v NullableUnavailableRootPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnavailableRootPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


