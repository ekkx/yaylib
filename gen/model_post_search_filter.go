
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostSearchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSearchFilter{}

// PostSearchFilter struct for PostSearchFilter
type PostSearchFilter struct {
	MessageId NullableInt32 `json:"message_id,omitempty"`
	TitleId NullableInt32 `json:"title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostSearchFilter PostSearchFilter

// NewPostSearchFilter instantiates a new PostSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSearchFilter() *PostSearchFilter {
	this := PostSearchFilter{}
	return &this
}

// NewPostSearchFilterWithDefaults instantiates a new PostSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSearchFilterWithDefaults() *PostSearchFilter {
	this := PostSearchFilter{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSearchFilter) GetMessageId() int32 {
	if o == nil || IsNil(o.MessageId.Get()) {
		var ret int32
		return ret
	}
	return *o.MessageId.Get()
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSearchFilter) GetMessageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageId.Get(), o.MessageId.IsSet()
}

// HasMessageId returns a boolean if a field has been set.
func (o *PostSearchFilter) HasMessageId() bool {
	if o != nil && o.MessageId.IsSet() {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given NullableInt32 and assigns it to the MessageId field.
func (o *PostSearchFilter) SetMessageId(v int32) {
	o.MessageId.Set(&v)
}
// SetMessageIdNil sets the value for MessageId to be an explicit nil
func (o *PostSearchFilter) SetMessageIdNil() {
	o.MessageId.Set(nil)
}

// UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
func (o *PostSearchFilter) UnsetMessageId() {
	o.MessageId.Unset()
}

// GetTitleId returns the TitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSearchFilter) GetTitleId() int32 {
	if o == nil || IsNil(o.TitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.TitleId.Get()
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSearchFilter) GetTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleId.Get(), o.TitleId.IsSet()
}

// HasTitleId returns a boolean if a field has been set.
func (o *PostSearchFilter) HasTitleId() bool {
	if o != nil && o.TitleId.IsSet() {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given NullableInt32 and assigns it to the TitleId field.
func (o *PostSearchFilter) SetTitleId(v int32) {
	o.TitleId.Set(&v)
}
// SetTitleIdNil sets the value for TitleId to be an explicit nil
func (o *PostSearchFilter) SetTitleIdNil() {
	o.TitleId.Set(nil)
}

// UnsetTitleId ensures that no value is present for TitleId, not even an explicit nil
func (o *PostSearchFilter) UnsetTitleId() {
	o.TitleId.Unset()
}

func (o PostSearchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSearchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MessageId.IsSet() {
		toSerialize["message_id"] = o.MessageId.Get()
	}
	if o.TitleId.IsSet() {
		toSerialize["title_id"] = o.TitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostSearchFilter) UnmarshalJSON(data []byte) (err error) {
	varPostSearchFilter := _PostSearchFilter{}

	err = json.Unmarshal(data, &varPostSearchFilter)

	if err != nil {
		return err
	}

	*o = PostSearchFilter(varPostSearchFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message_id")
		delete(additionalProperties, "title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostSearchFilter struct {
	value *PostSearchFilter
	isSet bool
}

func (v NullablePostSearchFilter) Get() *PostSearchFilter {
	return v.value
}

func (v *NullablePostSearchFilter) Set(val *PostSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSearchFilter(val *PostSearchFilter) *NullablePostSearchFilter {
	return &NullablePostSearchFilter{value: val, isSet: true}
}

func (v NullablePostSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


