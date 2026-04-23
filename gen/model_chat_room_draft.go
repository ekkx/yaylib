
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChatRoomDraft type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatRoomDraft{}

// ChatRoomDraft struct for ChatRoomDraft
type ChatRoomDraft struct {
	Id NullableInt64 `json:"id,omitempty"`
	Text NullableString `json:"text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatRoomDraft ChatRoomDraft

// NewChatRoomDraft instantiates a new ChatRoomDraft object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatRoomDraft() *ChatRoomDraft {
	this := ChatRoomDraft{}
	return &this
}

// NewChatRoomDraftWithDefaults instantiates a new ChatRoomDraft object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatRoomDraftWithDefaults() *ChatRoomDraft {
	this := ChatRoomDraft{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomDraft) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomDraft) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ChatRoomDraft) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ChatRoomDraft) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ChatRoomDraft) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ChatRoomDraft) UnsetId() {
	o.Id.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomDraft) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomDraft) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *ChatRoomDraft) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *ChatRoomDraft) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *ChatRoomDraft) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *ChatRoomDraft) UnsetText() {
	o.Text.Unset()
}

func (o ChatRoomDraft) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatRoomDraft) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatRoomDraft) UnmarshalJSON(data []byte) (err error) {
	varChatRoomDraft := _ChatRoomDraft{}

	err = json.Unmarshal(data, &varChatRoomDraft)

	if err != nil {
		return err
	}

	*o = ChatRoomDraft(varChatRoomDraft)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatRoomDraft struct {
	value *ChatRoomDraft
	isSet bool
}

func (v NullableChatRoomDraft) Get() *ChatRoomDraft {
	return v.value
}

func (v *NullableChatRoomDraft) Set(val *ChatRoomDraft) {
	v.value = val
	v.isSet = true
}

func (v NullableChatRoomDraft) IsSet() bool {
	return v.isSet
}

func (v *NullableChatRoomDraft) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatRoomDraft(val *ChatRoomDraft) *NullableChatRoomDraft {
	return &NullableChatRoomDraft{value: val, isSet: true}
}

func (v NullableChatRoomDraft) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatRoomDraft) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


