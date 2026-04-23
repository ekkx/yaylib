
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostResponse{}

// PostResponse struct for PostResponse
type PostResponse struct {
	Post NullablePost `json:"post,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostResponse PostResponse

// NewPostResponse instantiates a new PostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostResponse() *PostResponse {
	this := PostResponse{}
	return &this
}

// NewPostResponseWithDefaults instantiates a new PostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostResponseWithDefaults() *PostResponse {
	this := PostResponse{}
	return &this
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostResponse) GetPost() Post {
	if o == nil || IsNil(o.Post.Get()) {
		var ret Post
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostResponse) GetPostOk() (*Post, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *PostResponse) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullablePost and assigns it to the Post field.
func (o *PostResponse) SetPost(v Post) {
	o.Post.Set(&v)
}
// SetPostNil sets the value for Post to be an explicit nil
func (o *PostResponse) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *PostResponse) UnsetPost() {
	o.Post.Unset()
}

func (o PostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Post.IsSet() {
		toSerialize["post"] = o.Post.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostResponse) UnmarshalJSON(data []byte) (err error) {
	varPostResponse := _PostResponse{}

	err = json.Unmarshal(data, &varPostResponse)

	if err != nil {
		return err
	}

	*o = PostResponse(varPostResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "post")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostResponse struct {
	value *PostResponse
	isSet bool
}

func (v NullablePostResponse) Get() *PostResponse {
	return v.value
}

func (v *NullablePostResponse) Set(val *PostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostResponse(val *PostResponse) *NullablePostResponse {
	return &NullablePostResponse{value: val, isSet: true}
}

func (v NullablePostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


