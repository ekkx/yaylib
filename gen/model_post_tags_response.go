
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTagsResponse{}

// PostTagsResponse struct for PostTagsResponse
type PostTagsResponse struct {
	Tags []PostTag `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostTagsResponse PostTagsResponse

// NewPostTagsResponse instantiates a new PostTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTagsResponse() *PostTagsResponse {
	this := PostTagsResponse{}
	return &this
}

// NewPostTagsResponseWithDefaults instantiates a new PostTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTagsResponseWithDefaults() *PostTagsResponse {
	this := PostTagsResponse{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTagsResponse) GetTags() []PostTag {
	if o == nil {
		var ret []PostTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTagsResponse) GetTagsOk() ([]PostTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PostTagsResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []PostTag and assigns it to the Tags field.
func (o *PostTagsResponse) SetTags(v []PostTag) {
	o.Tags = v
}

func (o PostTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostTagsResponse) UnmarshalJSON(data []byte) (err error) {
	varPostTagsResponse := _PostTagsResponse{}

	err = json.Unmarshal(data, &varPostTagsResponse)

	if err != nil {
		return err
	}

	*o = PostTagsResponse(varPostTagsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostTagsResponse struct {
	value *PostTagsResponse
	isSet bool
}

func (v NullablePostTagsResponse) Get() *PostTagsResponse {
	return v.value
}

func (v *NullablePostTagsResponse) Set(val *PostTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTagsResponse(val *PostTagsResponse) *NullablePostTagsResponse {
	return &NullablePostTagsResponse{value: val, isSet: true}
}

func (v NullablePostTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


