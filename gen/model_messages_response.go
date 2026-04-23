
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MessagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagesResponse{}

// MessagesResponse struct for MessagesResponse
type MessagesResponse struct {
	Messages []RealmMessage `json:"messages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MessagesResponse MessagesResponse

// NewMessagesResponse instantiates a new MessagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagesResponse() *MessagesResponse {
	this := MessagesResponse{}
	return &this
}

// NewMessagesResponseWithDefaults instantiates a new MessagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagesResponseWithDefaults() *MessagesResponse {
	this := MessagesResponse{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessagesResponse) GetMessages() []RealmMessage {
	if o == nil {
		var ret []RealmMessage
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessagesResponse) GetMessagesOk() ([]RealmMessage, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *MessagesResponse) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []RealmMessage and assigns it to the Messages field.
func (o *MessagesResponse) SetMessages(v []RealmMessage) {
	o.Messages = v
}

func (o MessagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MessagesResponse) UnmarshalJSON(data []byte) (err error) {
	varMessagesResponse := _MessagesResponse{}

	err = json.Unmarshal(data, &varMessagesResponse)

	if err != nil {
		return err
	}

	*o = MessagesResponse(varMessagesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "messages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessagesResponse struct {
	value *MessagesResponse
	isSet bool
}

func (v NullableMessagesResponse) Get() *MessagesResponse {
	return v.value
}

func (v *NullableMessagesResponse) Set(val *MessagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagesResponse(val *MessagesResponse) *NullableMessagesResponse {
	return &NullableMessagesResponse{value: val, isSet: true}
}

func (v NullableMessagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


