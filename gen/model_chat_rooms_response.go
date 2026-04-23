
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChatRoomsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatRoomsResponse{}

// ChatRoomsResponse struct for ChatRoomsResponse
type ChatRoomsResponse struct {
	ChatRooms []RealmChatRoom `json:"chat_rooms,omitempty"`
	NextPageValue NullableInt64 `json:"next_page_value,omitempty"`
	PinnedChatRooms []RealmChatRoom `json:"pinned_chat_rooms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatRoomsResponse ChatRoomsResponse

// NewChatRoomsResponse instantiates a new ChatRoomsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatRoomsResponse() *ChatRoomsResponse {
	this := ChatRoomsResponse{}
	return &this
}

// NewChatRoomsResponseWithDefaults instantiates a new ChatRoomsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatRoomsResponseWithDefaults() *ChatRoomsResponse {
	this := ChatRoomsResponse{}
	return &this
}

// GetChatRooms returns the ChatRooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomsResponse) GetChatRooms() []RealmChatRoom {
	if o == nil {
		var ret []RealmChatRoom
		return ret
	}
	return o.ChatRooms
}

// GetChatRoomsOk returns a tuple with the ChatRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomsResponse) GetChatRoomsOk() ([]RealmChatRoom, bool) {
	if o == nil || IsNil(o.ChatRooms) {
		return nil, false
	}
	return o.ChatRooms, true
}

// HasChatRooms returns a boolean if a field has been set.
func (o *ChatRoomsResponse) HasChatRooms() bool {
	if o != nil && !IsNil(o.ChatRooms) {
		return true
	}

	return false
}

// SetChatRooms gets a reference to the given []RealmChatRoom and assigns it to the ChatRooms field.
func (o *ChatRoomsResponse) SetChatRooms(v []RealmChatRoom) {
	o.ChatRooms = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomsResponse) GetNextPageValue() int64 {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret int64
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomsResponse) GetNextPageValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *ChatRoomsResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableInt64 and assigns it to the NextPageValue field.
func (o *ChatRoomsResponse) SetNextPageValue(v int64) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *ChatRoomsResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *ChatRoomsResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetPinnedChatRooms returns the PinnedChatRooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoomsResponse) GetPinnedChatRooms() []RealmChatRoom {
	if o == nil {
		var ret []RealmChatRoom
		return ret
	}
	return o.PinnedChatRooms
}

// GetPinnedChatRoomsOk returns a tuple with the PinnedChatRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoomsResponse) GetPinnedChatRoomsOk() ([]RealmChatRoom, bool) {
	if o == nil || IsNil(o.PinnedChatRooms) {
		return nil, false
	}
	return o.PinnedChatRooms, true
}

// HasPinnedChatRooms returns a boolean if a field has been set.
func (o *ChatRoomsResponse) HasPinnedChatRooms() bool {
	if o != nil && !IsNil(o.PinnedChatRooms) {
		return true
	}

	return false
}

// SetPinnedChatRooms gets a reference to the given []RealmChatRoom and assigns it to the PinnedChatRooms field.
func (o *ChatRoomsResponse) SetPinnedChatRooms(v []RealmChatRoom) {
	o.PinnedChatRooms = v
}

func (o ChatRoomsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatRoomsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChatRooms != nil {
		toSerialize["chat_rooms"] = o.ChatRooms
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.PinnedChatRooms != nil {
		toSerialize["pinned_chat_rooms"] = o.PinnedChatRooms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatRoomsResponse) UnmarshalJSON(data []byte) (err error) {
	varChatRoomsResponse := _ChatRoomsResponse{}

	err = json.Unmarshal(data, &varChatRoomsResponse)

	if err != nil {
		return err
	}

	*o = ChatRoomsResponse(varChatRoomsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chat_rooms")
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "pinned_chat_rooms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatRoomsResponse struct {
	value *ChatRoomsResponse
	isSet bool
}

func (v NullableChatRoomsResponse) Get() *ChatRoomsResponse {
	return v.value
}

func (v *NullableChatRoomsResponse) Set(val *ChatRoomsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChatRoomsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChatRoomsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatRoomsResponse(val *ChatRoomsResponse) *NullableChatRoomsResponse {
	return &NullableChatRoomsResponse{value: val, isSet: true}
}

func (v NullableChatRoomsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatRoomsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


