
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostsResponse{}

// PostsResponse struct for PostsResponse
type PostsResponse struct {
	HasMoreHotPosts NullableBool `json:"has_more_hot_posts,omitempty"`
	NextPageValue NullableCursor `json:"next_page_value,omitempty"`
	PinnedPosts []Post `json:"pinned_posts,omitempty"`
	Posts []Post `json:"posts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostsResponse PostsResponse

// NewPostsResponse instantiates a new PostsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostsResponse() *PostsResponse {
	this := PostsResponse{}
	return &this
}

// NewPostsResponseWithDefaults instantiates a new PostsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostsResponseWithDefaults() *PostsResponse {
	this := PostsResponse{}
	return &this
}

// GetHasMoreHotPosts returns the HasMoreHotPosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostsResponse) GetHasMoreHotPosts() bool {
	if o == nil || IsNil(o.HasMoreHotPosts.Get()) {
		var ret bool
		return ret
	}
	return *o.HasMoreHotPosts.Get()
}

// GetHasMoreHotPostsOk returns a tuple with the HasMoreHotPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostsResponse) GetHasMoreHotPostsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasMoreHotPosts.Get(), o.HasMoreHotPosts.IsSet()
}

// HasHasMoreHotPosts returns a boolean if a field has been set.
func (o *PostsResponse) HasHasMoreHotPosts() bool {
	if o != nil && o.HasMoreHotPosts.IsSet() {
		return true
	}

	return false
}

// SetHasMoreHotPosts gets a reference to the given NullableBool and assigns it to the HasMoreHotPosts field.
func (o *PostsResponse) SetHasMoreHotPosts(v bool) {
	o.HasMoreHotPosts.Set(&v)
}
// SetHasMoreHotPostsNil sets the value for HasMoreHotPosts to be an explicit nil
func (o *PostsResponse) SetHasMoreHotPostsNil() {
	o.HasMoreHotPosts.Set(nil)
}

// UnsetHasMoreHotPosts ensures that no value is present for HasMoreHotPosts, not even an explicit nil
func (o *PostsResponse) UnsetHasMoreHotPosts() {
	o.HasMoreHotPosts.Unset()
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostsResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostsResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *PostsResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *PostsResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *PostsResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *PostsResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetPinnedPosts returns the PinnedPosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostsResponse) GetPinnedPosts() []Post {
	if o == nil {
		var ret []Post
		return ret
	}
	return o.PinnedPosts
}

// GetPinnedPostsOk returns a tuple with the PinnedPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostsResponse) GetPinnedPostsOk() ([]Post, bool) {
	if o == nil || IsNil(o.PinnedPosts) {
		return nil, false
	}
	return o.PinnedPosts, true
}

// HasPinnedPosts returns a boolean if a field has been set.
func (o *PostsResponse) HasPinnedPosts() bool {
	if o != nil && !IsNil(o.PinnedPosts) {
		return true
	}

	return false
}

// SetPinnedPosts gets a reference to the given []Post and assigns it to the PinnedPosts field.
func (o *PostsResponse) SetPinnedPosts(v []Post) {
	o.PinnedPosts = v
}

// GetPosts returns the Posts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostsResponse) GetPosts() []Post {
	if o == nil {
		var ret []Post
		return ret
	}
	return o.Posts
}

// GetPostsOk returns a tuple with the Posts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostsResponse) GetPostsOk() ([]Post, bool) {
	if o == nil || IsNil(o.Posts) {
		return nil, false
	}
	return o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *PostsResponse) HasPosts() bool {
	if o != nil && !IsNil(o.Posts) {
		return true
	}

	return false
}

// SetPosts gets a reference to the given []Post and assigns it to the Posts field.
func (o *PostsResponse) SetPosts(v []Post) {
	o.Posts = v
}

func (o PostsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HasMoreHotPosts.IsSet() {
		toSerialize["has_more_hot_posts"] = o.HasMoreHotPosts.Get()
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.PinnedPosts != nil {
		toSerialize["pinned_posts"] = o.PinnedPosts
	}
	if o.Posts != nil {
		toSerialize["posts"] = o.Posts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostsResponse) UnmarshalJSON(data []byte) (err error) {
	varPostsResponse := _PostsResponse{}

	err = json.Unmarshal(data, &varPostsResponse)

	if err != nil {
		return err
	}

	*o = PostsResponse(varPostsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "has_more_hot_posts")
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "pinned_posts")
		delete(additionalProperties, "posts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostsResponse struct {
	value *PostsResponse
	isSet bool
}

func (v NullablePostsResponse) Get() *PostsResponse {
	return v.value
}

func (v *NullablePostsResponse) Set(val *PostsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostsResponse(val *PostsResponse) *NullablePostsResponse {
	return &NullablePostsResponse{value: val, isSet: true}
}

func (v NullablePostsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


