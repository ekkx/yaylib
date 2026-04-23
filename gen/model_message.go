
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message struct for Message
type Message struct {
	ConferenceCall NullableConferenceCall `json:"conference_call,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Invitation NullableChatInvitation `json:"invitation,omitempty"`
	Text NullableString `json:"text,omitempty"`
	Type NullableType `json:"type,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage() *Message {
	this := Message{}
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetConferenceCall() ConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret ConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetConferenceCallOk() (*ConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *Message) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableConferenceCall and assigns it to the ConferenceCall field.
func (o *Message) SetConferenceCall(v ConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *Message) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *Message) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Message) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Message) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Message) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Message) UnsetId() {
	o.Id.Unset()
}

// GetInvitation returns the Invitation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetInvitation() ChatInvitation {
	if o == nil || IsNil(o.Invitation.Get()) {
		var ret ChatInvitation
		return ret
	}
	return *o.Invitation.Get()
}

// GetInvitationOk returns a tuple with the Invitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetInvitationOk() (*ChatInvitation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Invitation.Get(), o.Invitation.IsSet()
}

// HasInvitation returns a boolean if a field has been set.
func (o *Message) HasInvitation() bool {
	if o != nil && o.Invitation.IsSet() {
		return true
	}

	return false
}

// SetInvitation gets a reference to the given NullableChatInvitation and assigns it to the Invitation field.
func (o *Message) SetInvitation(v ChatInvitation) {
	o.Invitation.Set(&v)
}
// SetInvitationNil sets the value for Invitation to be an explicit nil
func (o *Message) SetInvitationNil() {
	o.Invitation.Set(nil)
}

// UnsetInvitation ensures that no value is present for Invitation, not even an explicit nil
func (o *Message) UnsetInvitation() {
	o.Invitation.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *Message) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *Message) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *Message) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *Message) UnsetText() {
	o.Text.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetType() Type {
	if o == nil || IsNil(o.Type.Get()) {
		var ret Type
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Message) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableType and assigns it to the Type field.
func (o *Message) SetType(v Type) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Message) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Message) UnsetType() {
	o.Type.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *Message) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *Message) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *Message) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *Message) UnsetUserId() {
	o.UserId.Unset()
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConferenceCall.IsSet() {
		toSerialize["conference_call"] = o.ConferenceCall.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Invitation.IsSet() {
		toSerialize["invitation"] = o.Invitation.Get()
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Message) UnmarshalJSON(data []byte) (err error) {
	varMessage := _Message{}

	err = json.Unmarshal(data, &varMessage)

	if err != nil {
		return err
	}

	*o = Message(varMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invitation")
		delete(additionalProperties, "text")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


