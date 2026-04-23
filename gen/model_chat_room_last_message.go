
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChatRoomLastMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatRoomLastMessage{}

// ChatRoomLastMessage struct for ChatRoomLastMessage
type ChatRoomLastMessage struct {
	ConferenceCall NullableRealmConferenceCall `json:"conference_call,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	MessageType NullableString `json:"message_type,omitempty"`
	RoomId NullableInt64 `json:"room_id,omitempty"`
	Text NullableString `json:"text,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatRoomLastMessage ChatRoomLastMessage

// NewChatRoomLastMessage instantiates a new ChatRoomLastMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatRoomLastMessage() *ChatRoomLastMessage {
	this := ChatRoomLastMessage{}
	return &this
}

// NewChatRoomLastMessageWithDefaults instantiates a new ChatRoomLastMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatRoomLastMessageWithDefaults() *ChatRoomLastMessage {
	this := ChatRoomLastMessage{}
	return &this
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomLastMessage) GetConferenceCall() RealmConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret RealmConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomLastMessage) GetConferenceCallOk() (*RealmConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *ChatRoomLastMessage) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableRealmConferenceCall and assigns it to the ConferenceCall field.
func (o *ChatRoomLastMessage) SetConferenceCall(v RealmConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *ChatRoomLastMessage) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *ChatRoomLastMessage) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomLastMessage) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomLastMessage) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ChatRoomLastMessage) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *ChatRoomLastMessage) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ChatRoomLastMessage) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ChatRoomLastMessage) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomLastMessage) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomLastMessage) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ChatRoomLastMessage) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ChatRoomLastMessage) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ChatRoomLastMessage) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ChatRoomLastMessage) UnsetId() {
	o.Id.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomLastMessage) GetMessageType() string {
	if o == nil || IsNil(o.MessageType.Get()) {
		var ret string
		return ret
	}
	return *o.MessageType.Get()
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomLastMessage) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageType.Get(), o.MessageType.IsSet()
}

// HasMessageType returns a boolean if a field has been set.
func (o *ChatRoomLastMessage) HasMessageType() bool {
	if o != nil && o.MessageType.IsSet() {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given NullableString and assigns it to the MessageType field.
func (o *ChatRoomLastMessage) SetMessageType(v string) {
	o.MessageType.Set(&v)
}
// SetMessageTypeNil sets the value for MessageType to be an explicit nil
func (o *ChatRoomLastMessage) SetMessageTypeNil() {
	o.MessageType.Set(nil)
}

// UnsetMessageType ensures that no value is present for MessageType, not even an explicit nil
func (o *ChatRoomLastMessage) UnsetMessageType() {
	o.MessageType.Unset()
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomLastMessage) GetRoomId() int64 {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret int64
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomLastMessage) GetRoomIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *ChatRoomLastMessage) HasRoomId() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableInt64 and assigns it to the RoomId field.
func (o *ChatRoomLastMessage) SetRoomId(v int64) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *ChatRoomLastMessage) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *ChatRoomLastMessage) UnsetRoomId() {
	o.RoomId.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomLastMessage) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomLastMessage) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *ChatRoomLastMessage) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *ChatRoomLastMessage) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *ChatRoomLastMessage) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *ChatRoomLastMessage) UnsetText() {
	o.Text.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomLastMessage) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomLastMessage) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ChatRoomLastMessage) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *ChatRoomLastMessage) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ChatRoomLastMessage) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ChatRoomLastMessage) UnsetUserId() {
	o.UserId.Unset()
}

func (o ChatRoomLastMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatRoomLastMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConferenceCall.IsSet() {
		toSerialize["conference_call"] = o.ConferenceCall.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MessageType.IsSet() {
		toSerialize["message_type"] = o.MessageType.Get()
	}
	if o.RoomId.IsSet() {
		toSerialize["room_id"] = o.RoomId.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatRoomLastMessage) UnmarshalJSON(data []byte) (err error) {
	varChatRoomLastMessage := _ChatRoomLastMessage{}

	err = json.Unmarshal(data, &varChatRoomLastMessage)

	if err != nil {
		return err
	}

	*o = ChatRoomLastMessage(varChatRoomLastMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "room_id")
		delete(additionalProperties, "text")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatRoomLastMessage struct {
	value *ChatRoomLastMessage
	isSet bool
}

func (v NullableChatRoomLastMessage) Get() *ChatRoomLastMessage {
	return v.value
}

func (v *NullableChatRoomLastMessage) Set(val *ChatRoomLastMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableChatRoomLastMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableChatRoomLastMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatRoomLastMessage(val *ChatRoomLastMessage) *NullableChatRoomLastMessage {
	return &NullableChatRoomLastMessage{value: val, isSet: true}
}

func (v NullableChatRoomLastMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatRoomLastMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


