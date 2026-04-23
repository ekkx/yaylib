
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreatePostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePostResponse{}

// CreatePostResponse struct for CreatePostResponse
type CreatePostResponse struct {
	ConferenceCall NullableRealmConferenceCall `json:"conference_call,omitempty"`
	Post NullablePost `json:"post,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreatePostResponse CreatePostResponse

// NewCreatePostResponse instantiates a new CreatePostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePostResponse() *CreatePostResponse {
	this := CreatePostResponse{}
	return &this
}

// NewCreatePostResponseWithDefaults instantiates a new CreatePostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePostResponseWithDefaults() *CreatePostResponse {
	this := CreatePostResponse{}
	return &this
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePostResponse) GetConferenceCall() RealmConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret RealmConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePostResponse) GetConferenceCallOk() (*RealmConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *CreatePostResponse) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableRealmConferenceCall and assigns it to the ConferenceCall field.
func (o *CreatePostResponse) SetConferenceCall(v RealmConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *CreatePostResponse) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *CreatePostResponse) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePostResponse) GetPost() Post {
	if o == nil || IsNil(o.Post.Get()) {
		var ret Post
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePostResponse) GetPostOk() (*Post, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *CreatePostResponse) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullablePost and assigns it to the Post field.
func (o *CreatePostResponse) SetPost(v Post) {
	o.Post.Set(&v)
}
// SetPostNil sets the value for Post to be an explicit nil
func (o *CreatePostResponse) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *CreatePostResponse) UnsetPost() {
	o.Post.Unset()
}

func (o CreatePostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConferenceCall.IsSet() {
		toSerialize["conference_call"] = o.ConferenceCall.Get()
	}
	if o.Post.IsSet() {
		toSerialize["post"] = o.Post.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreatePostResponse) UnmarshalJSON(data []byte) (err error) {
	varCreatePostResponse := _CreatePostResponse{}

	err = json.Unmarshal(data, &varCreatePostResponse)

	if err != nil {
		return err
	}

	*o = CreatePostResponse(varCreatePostResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "post")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePostResponse struct {
	value *CreatePostResponse
	isSet bool
}

func (v NullableCreatePostResponse) Get() *CreatePostResponse {
	return v.value
}

func (v *NullableCreatePostResponse) Set(val *CreatePostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePostResponse(val *CreatePostResponse) *NullableCreatePostResponse {
	return &NullableCreatePostResponse{value: val, isSet: true}
}

func (v NullableCreatePostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


