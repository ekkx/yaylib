
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the HiddenRecommendedPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HiddenRecommendedPost{}

// HiddenRecommendedPost struct for HiddenRecommendedPost
type HiddenRecommendedPost struct {
	Post NullableModelPost `json:"post,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HiddenRecommendedPost HiddenRecommendedPost

// NewHiddenRecommendedPost instantiates a new HiddenRecommendedPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHiddenRecommendedPost() *HiddenRecommendedPost {
	this := HiddenRecommendedPost{}
	return &this
}

// NewHiddenRecommendedPostWithDefaults instantiates a new HiddenRecommendedPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHiddenRecommendedPostWithDefaults() *HiddenRecommendedPost {
	this := HiddenRecommendedPost{}
	return &this
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HiddenRecommendedPost) GetPost() ModelPost {
	if o == nil || IsNil(o.Post.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HiddenRecommendedPost) GetPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *HiddenRecommendedPost) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullableModelPost and assigns it to the Post field.
func (o *HiddenRecommendedPost) SetPost(v ModelPost) {
	o.Post.Set(&v)
}
// SetPostNil sets the value for Post to be an explicit nil
func (o *HiddenRecommendedPost) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *HiddenRecommendedPost) UnsetPost() {
	o.Post.Unset()
}

func (o HiddenRecommendedPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HiddenRecommendedPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Post.IsSet() {
		toSerialize["post"] = o.Post.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HiddenRecommendedPost) UnmarshalJSON(data []byte) (err error) {
	varHiddenRecommendedPost := _HiddenRecommendedPost{}

	err = json.Unmarshal(data, &varHiddenRecommendedPost)

	if err != nil {
		return err
	}

	*o = HiddenRecommendedPost(varHiddenRecommendedPost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "post")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHiddenRecommendedPost struct {
	value *HiddenRecommendedPost
	isSet bool
}

func (v NullableHiddenRecommendedPost) Get() *HiddenRecommendedPost {
	return v.value
}

func (v *NullableHiddenRecommendedPost) Set(val *HiddenRecommendedPost) {
	v.value = val
	v.isSet = true
}

func (v NullableHiddenRecommendedPost) IsSet() bool {
	return v.isSet
}

func (v *NullableHiddenRecommendedPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHiddenRecommendedPost(val *HiddenRecommendedPost) *NullableHiddenRecommendedPost {
	return &NullableHiddenRecommendedPost{value: val, isSet: true}
}

func (v NullableHiddenRecommendedPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHiddenRecommendedPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


