
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChannelMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelMessage{}

// ChannelMessage struct for ChannelMessage
type ChannelMessage struct {
	Identifier NullableString `json:"identifier,omitempty"`
	Message NullableEventMessage `json:"message,omitempty"`
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelMessage ChannelMessage

// NewChannelMessage instantiates a new ChannelMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelMessage() *ChannelMessage {
	this := ChannelMessage{}
	return &this
}

// NewChannelMessageWithDefaults instantiates a new ChannelMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelMessageWithDefaults() *ChannelMessage {
	this := ChannelMessage{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelMessage) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelMessage) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ChannelMessage) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *ChannelMessage) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *ChannelMessage) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *ChannelMessage) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelMessage) GetMessage() EventMessage {
	if o == nil || IsNil(o.Message.Get()) {
		var ret EventMessage
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelMessage) GetMessageOk() (*EventMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *ChannelMessage) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableEventMessage and assigns it to the Message field.
func (o *ChannelMessage) SetMessage(v EventMessage) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *ChannelMessage) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *ChannelMessage) UnsetMessage() {
	o.Message.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelMessage) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelMessage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ChannelMessage) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ChannelMessage) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ChannelMessage) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ChannelMessage) UnsetType() {
	o.Type.Unset()
}

func (o ChannelMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChannelMessage) UnmarshalJSON(data []byte) (err error) {
	varChannelMessage := _ChannelMessage{}

	err = json.Unmarshal(data, &varChannelMessage)

	if err != nil {
		return err
	}

	*o = ChannelMessage(varChannelMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "message")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelMessage struct {
	value *ChannelMessage
	isSet bool
}

func (v NullableChannelMessage) Get() *ChannelMessage {
	return v.value
}

func (v *NullableChannelMessage) Set(val *ChannelMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMessage(val *ChannelMessage) *NullableChannelMessage {
	return &NullableChannelMessage{value: val, isSet: true}
}

func (v NullableChannelMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


