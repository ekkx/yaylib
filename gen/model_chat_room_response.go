
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChatRoomResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatRoomResponse{}

// ChatRoomResponse struct for ChatRoomResponse
type ChatRoomResponse struct {
	Chat NullableRealmChatRoom `json:"chat,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatRoomResponse ChatRoomResponse

// NewChatRoomResponse instantiates a new ChatRoomResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatRoomResponse() *ChatRoomResponse {
	this := ChatRoomResponse{}
	return &this
}

// NewChatRoomResponseWithDefaults instantiates a new ChatRoomResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatRoomResponseWithDefaults() *ChatRoomResponse {
	this := ChatRoomResponse{}
	return &this
}

// GetChat returns the Chat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomResponse) GetChat() RealmChatRoom {
	if o == nil || IsNil(o.Chat.Get()) {
		var ret RealmChatRoom
		return ret
	}
	return *o.Chat.Get()
}

// GetChatOk returns a tuple with the Chat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomResponse) GetChatOk() (*RealmChatRoom, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chat.Get(), o.Chat.IsSet()
}

// HasChat returns a boolean if a field has been set.
func (o *ChatRoomResponse) HasChat() bool {
	if o != nil && o.Chat.IsSet() {
		return true
	}

	return false
}

// SetChat gets a reference to the given NullableRealmChatRoom and assigns it to the Chat field.
func (o *ChatRoomResponse) SetChat(v RealmChatRoom) {
	o.Chat.Set(&v)
}
// SetChatNil sets the value for Chat to be an explicit nil
func (o *ChatRoomResponse) SetChatNil() {
	o.Chat.Set(nil)
}

// UnsetChat ensures that no value is present for Chat, not even an explicit nil
func (o *ChatRoomResponse) UnsetChat() {
	o.Chat.Unset()
}

func (o ChatRoomResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatRoomResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Chat.IsSet() {
		toSerialize["chat"] = o.Chat.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatRoomResponse) UnmarshalJSON(data []byte) (err error) {
	varChatRoomResponse := _ChatRoomResponse{}

	err = json.Unmarshal(data, &varChatRoomResponse)

	if err != nil {
		return err
	}

	*o = ChatRoomResponse(varChatRoomResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chat")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatRoomResponse struct {
	value *ChatRoomResponse
	isSet bool
}

func (v NullableChatRoomResponse) Get() *ChatRoomResponse {
	return v.value
}

func (v *NullableChatRoomResponse) Set(val *ChatRoomResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatRoomResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatRoomResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatRoomResponse(val *ChatRoomResponse) *NullableChatRoomResponse {
	return &NullableChatRoomResponse{value: val, isSet: true}
}

func (v NullableChatRoomResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatRoomResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


