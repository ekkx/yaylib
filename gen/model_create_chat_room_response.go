
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateChatRoomResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatRoomResponse{}

// CreateChatRoomResponse struct for CreateChatRoomResponse
type CreateChatRoomResponse struct {
	RoomId NullableInt64 `json:"room_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateChatRoomResponse CreateChatRoomResponse

// NewCreateChatRoomResponse instantiates a new CreateChatRoomResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatRoomResponse() *CreateChatRoomResponse {
	this := CreateChatRoomResponse{}
	return &this
}

// NewCreateChatRoomResponseWithDefaults instantiates a new CreateChatRoomResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatRoomResponseWithDefaults() *CreateChatRoomResponse {
	this := CreateChatRoomResponse{}
	return &this
}

// GetRoomId returns the RoomId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateChatRoomResponse) GetRoomId() int64 {
	if o == nil || IsNil(o.RoomId.Get()) {
		var ret int64
		return ret
	}
	return *o.RoomId.Get()
}

// GetRoomIdOk returns a tuple with the RoomId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateChatRoomResponse) GetRoomIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoomId.Get(), o.RoomId.IsSet()
}

// HasRoomId returns a boolean if a field has been set.
func (o *CreateChatRoomResponse) HasRoomId() bool {
	if o != nil && o.RoomId.IsSet() {
		return true
	}

	return false
}

// SetRoomId gets a reference to the given NullableInt64 and assigns it to the RoomId field.
func (o *CreateChatRoomResponse) SetRoomId(v int64) {
	o.RoomId.Set(&v)
}
// SetRoomIdNil sets the value for RoomId to be an explicit nil
func (o *CreateChatRoomResponse) SetRoomIdNil() {
	o.RoomId.Set(nil)
}

// UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
func (o *CreateChatRoomResponse) UnsetRoomId() {
	o.RoomId.Unset()
}

func (o CreateChatRoomResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatRoomResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RoomId.IsSet() {
		toSerialize["room_id"] = o.RoomId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateChatRoomResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateChatRoomResponse := _CreateChatRoomResponse{}

	err = json.Unmarshal(data, &varCreateChatRoomResponse)

	if err != nil {
		return err
	}

	*o = CreateChatRoomResponse(varCreateChatRoomResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "room_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateChatRoomResponse struct {
	value *CreateChatRoomResponse
	isSet bool
}

func (v NullableCreateChatRoomResponse) Get() *CreateChatRoomResponse {
	return v.value
}

func (v *NullableCreateChatRoomResponse) Set(val *CreateChatRoomResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatRoomResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatRoomResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatRoomResponse(val *CreateChatRoomResponse) *NullableCreateChatRoomResponse {
	return &NullableCreateChatRoomResponse{value: val, isSet: true}
}

func (v NullableCreateChatRoomResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatRoomResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


