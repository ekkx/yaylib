
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTag{}

// PostTag struct for PostTag
type PostTag struct {
	Id NullableInt64 `json:"id,omitempty"`
	PostHashtagsCount NullableInt32 `json:"post_hashtags_count,omitempty"`
	Tag NullableString `json:"tag,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostTag PostTag

// NewPostTag instantiates a new PostTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTag() *PostTag {
	this := PostTag{}
	return &this
}

// NewPostTagWithDefaults instantiates a new PostTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTagWithDefaults() *PostTag {
	this := PostTag{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTag) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTag) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PostTag) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *PostTag) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PostTag) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PostTag) UnsetId() {
	o.Id.Unset()
}

// GetPostHashtagsCount returns the PostHashtagsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTag) GetPostHashtagsCount() int32 {
	if o == nil || IsNil(o.PostHashtagsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PostHashtagsCount.Get()
}

// GetPostHashtagsCountOk returns a tuple with the PostHashtagsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTag) GetPostHashtagsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostHashtagsCount.Get(), o.PostHashtagsCount.IsSet()
}

// HasPostHashtagsCount returns a boolean if a field has been set.
func (o *PostTag) HasPostHashtagsCount() bool {
	if o != nil && o.PostHashtagsCount.IsSet() {
		return true
	}

	return false
}

// SetPostHashtagsCount gets a reference to the given NullableInt32 and assigns it to the PostHashtagsCount field.
func (o *PostTag) SetPostHashtagsCount(v int32) {
	o.PostHashtagsCount.Set(&v)
}
// SetPostHashtagsCountNil sets the value for PostHashtagsCount to be an explicit nil
func (o *PostTag) SetPostHashtagsCountNil() {
	o.PostHashtagsCount.Set(nil)
}

// UnsetPostHashtagsCount ensures that no value is present for PostHashtagsCount, not even an explicit nil
func (o *PostTag) UnsetPostHashtagsCount() {
	o.PostHashtagsCount.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTag) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTag) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *PostTag) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *PostTag) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *PostTag) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *PostTag) UnsetTag() {
	o.Tag.Unset()
}

func (o PostTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.PostHashtagsCount.IsSet() {
		toSerialize["post_hashtags_count"] = o.PostHashtagsCount.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostTag) UnmarshalJSON(data []byte) (err error) {
	varPostTag := _PostTag{}

	err = json.Unmarshal(data, &varPostTag)

	if err != nil {
		return err
	}

	*o = PostTag(varPostTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "post_hashtags_count")
		delete(additionalProperties, "tag")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostTag struct {
	value *PostTag
	isSet bool
}

func (v NullablePostTag) Get() *PostTag {
	return v.value
}

func (v *NullablePostTag) Set(val *PostTag) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTag) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTag(val *PostTag) *NullablePostTag {
	return &NullablePostTag{value: val, isSet: true}
}

func (v NullablePostTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


