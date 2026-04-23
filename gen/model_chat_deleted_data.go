
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChatDeletedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatDeletedData{}

// ChatDeletedData struct for ChatDeletedData
type ChatDeletedData struct {
	RoomId NullableInt64 `json:"room_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatDeletedData ChatDeletedData

// NewChatDeletedData instantiates a new ChatDeletedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatDeletedData() *ChatDeletedData {
	this := ChatDeletedData{}
	return &this
}

// NewChatDeletedDataWithDefaults instantiates a new ChatDeletedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatDeletedDataWithDefaults() *ChatDeletedData {
	this := ChatDeletedData{}
	return &this
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatDeletedData) GetRoomId() int64 {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret int64
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatDeletedData) GetRoomIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *ChatDeletedData) HasRoomId() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableInt64 and assigns it to the RoomId field.
func (o *ChatDeletedData) SetRoomId(v int64) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *ChatDeletedData) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *ChatDeletedData) UnsetRoomId() {
	o.RoomId.Unset()
}

func (o ChatDeletedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatDeletedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomId.IsSet() {
		toSerialize["room_id"] = o.RoomId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatDeletedData) UnmarshalJSON(data []byte) (err error) {
	varChatDeletedData := _ChatDeletedData{}

	err = json.Unmarshal(data, &varChatDeletedData)

	if err != nil {
		return err
	}

	*o = ChatDeletedData(varChatDeletedData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "room_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatDeletedData struct {
	value *ChatDeletedData
	isSet bool
}

func (v NullableChatDeletedData) Get() *ChatDeletedData {
	return v.value
}

func (v *NullableChatDeletedData) Set(val *ChatDeletedData) {
	v.value = val
	v.isSet = true
}

func (v NullableChatDeletedData) IsSet() bool {
	return v.isSet
}

func (v *NullableChatDeletedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatDeletedData(val *ChatDeletedData) *NullableChatDeletedData {
	return &NullableChatDeletedData{value: val, isSet: true}
}

func (v NullableChatDeletedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatDeletedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


