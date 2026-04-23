
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageResponse{}

// MessageResponse struct for MessageResponse
type MessageResponse struct {
	ConferenceCall NullableRealmConferenceCall `json:"conference_call,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MessageResponse MessageResponse

// NewMessageResponse instantiates a new MessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageResponse() *MessageResponse {
	this := MessageResponse{}
	return &this
}

// NewMessageResponseWithDefaults instantiates a new MessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageResponseWithDefaults() *MessageResponse {
	this := MessageResponse{}
	return &this
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageResponse) GetConferenceCall() RealmConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret RealmConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageResponse) GetConferenceCallOk() (*RealmConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *MessageResponse) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableRealmConferenceCall and assigns it to the ConferenceCall field.
func (o *MessageResponse) SetConferenceCall(v RealmConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *MessageResponse) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *MessageResponse) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageResponse) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MessageResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *MessageResponse) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MessageResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MessageResponse) UnsetId() {
	o.Id.Unset()
}

func (o MessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConferenceCall.IsSet() {
		toSerialize["conference_call"] = o.ConferenceCall.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MessageResponse) UnmarshalJSON(data []byte) (err error) {
	varMessageResponse := _MessageResponse{}

	err = json.Unmarshal(data, &varMessageResponse)

	if err != nil {
		return err
	}

	*o = MessageResponse(varMessageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessageResponse struct {
	value *MessageResponse
	isSet bool
}

func (v NullableMessageResponse) Get() *MessageResponse {
	return v.value
}

func (v *NullableMessageResponse) Set(val *MessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageResponse(val *MessageResponse) *NullableMessageResponse {
	return &NullableMessageResponse{value: val, isSet: true}
}

func (v NullableMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


