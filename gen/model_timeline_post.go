
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TimelinePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelinePost{}

// TimelinePost struct for TimelinePost
type TimelinePost struct {
	Posts []ModelPost `json:"posts,omitempty"`
	RootPosts []ModelPost `json:"root_posts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimelinePost TimelinePost

// NewTimelinePost instantiates a new TimelinePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimelinePost() *TimelinePost {
	this := TimelinePost{}
	return &this
}

// NewTimelinePostWithDefaults instantiates a new TimelinePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimelinePostWithDefaults() *TimelinePost {
	this := TimelinePost{}
	return &this
}

// GetPosts returns the Posts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePost) GetPosts() []ModelPost {
	if o == nil {
		var ret []ModelPost
		return ret
	}
	return o.Posts
}

// GetPostsOk returns a tuple with the Posts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePost) GetPostsOk() ([]ModelPost, bool) {
	if o == nil || IsNil(o.Posts) {
		return nil, false
	}
	return o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *TimelinePost) HasPosts() bool {
	if o != nil && !IsNil(o.Posts) {
		return true
	}

	return false
}

// SetPosts gets a reference to the given []ModelPost and assigns it to the Posts field.
func (o *TimelinePost) SetPosts(v []ModelPost) {
	o.Posts = v
}

// GetRootPosts returns the RootPosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimelinePost) GetRootPosts() []ModelPost {
	if o == nil {
		var ret []ModelPost
		return ret
	}
	return o.RootPosts
}

// GetRootPostsOk returns a tuple with the RootPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimelinePost) GetRootPostsOk() ([]ModelPost, bool) {
	if o == nil || IsNil(o.RootPosts) {
		return nil, false
	}
	return o.RootPosts, true
}

// HasRootPosts returns a boolean if a field has been set.
func (o *TimelinePost) HasRootPosts() bool {
	if o != nil && !IsNil(o.RootPosts) {
		return true
	}

	return false
}

// SetRootPosts gets a reference to the given []ModelPost and assigns it to the RootPosts field.
func (o *TimelinePost) SetRootPosts(v []ModelPost) {
	o.RootPosts = v
}

func (o TimelinePost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimelinePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Posts != nil {
		toSerialize["posts"] = o.Posts
	}
	if o.RootPosts != nil {
		toSerialize["root_posts"] = o.RootPosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimelinePost) UnmarshalJSON(data []byte) (err error) {
	varTimelinePost := _TimelinePost{}

	err = json.Unmarshal(data, &varTimelinePost)

	if err != nil {
		return err
	}

	*o = TimelinePost(varTimelinePost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "posts")
		delete(additionalProperties, "root_posts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimelinePost struct {
	value *TimelinePost
	isSet bool
}

func (v NullableTimelinePost) Get() *TimelinePost {
	return v.value
}

func (v *NullableTimelinePost) Set(val *TimelinePost) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelinePost) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelinePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelinePost(val *TimelinePost) *NullableTimelinePost {
	return &NullableTimelinePost{value: val, isSet: true}
}

func (v NullableTimelinePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelinePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


