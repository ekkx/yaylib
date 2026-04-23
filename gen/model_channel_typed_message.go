
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChannelTypedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTypedMessage{}

// ChannelTypedMessage struct for ChannelTypedMessage
type ChannelTypedMessage struct {
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelTypedMessage ChannelTypedMessage

// NewChannelTypedMessage instantiates a new ChannelTypedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTypedMessage() *ChannelTypedMessage {
	this := ChannelTypedMessage{}
	return &this
}

// NewChannelTypedMessageWithDefaults instantiates a new ChannelTypedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTypedMessageWithDefaults() *ChannelTypedMessage {
	this := ChannelTypedMessage{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelTypedMessage) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelTypedMessage) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ChannelTypedMessage) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ChannelTypedMessage) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ChannelTypedMessage) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ChannelTypedMessage) UnsetType() {
	o.Type.Unset()
}

func (o ChannelTypedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTypedMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChannelTypedMessage) UnmarshalJSON(data []byte) (err error) {
	varChannelTypedMessage := _ChannelTypedMessage{}

	err = json.Unmarshal(data, &varChannelTypedMessage)

	if err != nil {
		return err
	}

	*o = ChannelTypedMessage(varChannelTypedMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelTypedMessage struct {
	value *ChannelTypedMessage
	isSet bool
}

func (v NullableChannelTypedMessage) Get() *ChannelTypedMessage {
	return v.value
}

func (v *NullableChannelTypedMessage) Set(val *ChannelTypedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTypedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTypedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTypedMessage(val *ChannelTypedMessage) *NullableChannelTypedMessage {
	return &NullableChannelTypedMessage{value: val, isSet: true}
}

func (v NullableChannelTypedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTypedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


