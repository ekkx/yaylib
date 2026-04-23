
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChannelCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelCommand{}

// ChannelCommand struct for ChannelCommand
type ChannelCommand struct {
	Command NullableString `json:"command,omitempty"`
	Identifier NullableString `json:"identifier,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelCommand ChannelCommand

// NewChannelCommand instantiates a new ChannelCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelCommand() *ChannelCommand {
	this := ChannelCommand{}
	return &this
}

// NewChannelCommandWithDefaults instantiates a new ChannelCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelCommandWithDefaults() *ChannelCommand {
	this := ChannelCommand{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelCommand) GetCommand() string {
	if o == nil || IsNil(o.Command.Get()) {
		var ret string
		return ret
	}
	return *o.Command.Get()
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelCommand) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Command.Get(), o.Command.IsSet()
}

// HasCommand returns a boolean if a field has been set.
func (o *ChannelCommand) HasCommand() bool {
	if o != nil && o.Command.IsSet() {
		return true
	}

	return false
}

// SetCommand gets a reference to the given NullableString and assigns it to the Command field.
func (o *ChannelCommand) SetCommand(v string) {
	o.Command.Set(&v)
}
// SetCommandNil sets the value for Command to be an explicit nil
func (o *ChannelCommand) SetCommandNil() {
	o.Command.Set(nil)
}

// UnsetCommand ensures that no value is present for Command, not even an explicit nil
func (o *ChannelCommand) UnsetCommand() {
	o.Command.Unset()
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelCommand) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelCommand) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ChannelCommand) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *ChannelCommand) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *ChannelCommand) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *ChannelCommand) UnsetIdentifier() {
	o.Identifier.Unset()
}

func (o ChannelCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Command.IsSet() {
		toSerialize["command"] = o.Command.Get()
	}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChannelCommand) UnmarshalJSON(data []byte) (err error) {
	varChannelCommand := _ChannelCommand{}

	err = json.Unmarshal(data, &varChannelCommand)

	if err != nil {
		return err
	}

	*o = ChannelCommand(varChannelCommand)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "command")
		delete(additionalProperties, "identifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelCommand struct {
	value *ChannelCommand
	isSet bool
}

func (v NullableChannelCommand) Get() *ChannelCommand {
	return v.value
}

func (v *NullableChannelCommand) Set(val *ChannelCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelCommand(val *ChannelCommand) *NullableChannelCommand {
	return &NullableChannelCommand{value: val, isSet: true}
}

func (v NullableChannelCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


