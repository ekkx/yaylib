
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EventMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventMessage{}

// EventMessage struct for EventMessage
type EventMessage struct {
	// Unresolved Java type: jp.nanameue.yay.api.WebSocketInteractor.b
	Data map[string]interface{} `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventMessage EventMessage

// NewEventMessage instantiates a new EventMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventMessage() *EventMessage {
	this := EventMessage{}
	return &this
}

// NewEventMessageWithDefaults instantiates a new EventMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventMessageWithDefaults() *EventMessage {
	this := EventMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventMessage) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventMessage) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EventMessage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *EventMessage) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o EventMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventMessage) UnmarshalJSON(data []byte) (err error) {
	varEventMessage := _EventMessage{}

	err = json.Unmarshal(data, &varEventMessage)

	if err != nil {
		return err
	}

	*o = EventMessage(varEventMessage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventMessage struct {
	value *EventMessage
	isSet bool
}

func (v NullableEventMessage) Get() *EventMessage {
	return v.value
}

func (v *NullableEventMessage) Set(val *EventMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableEventMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableEventMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventMessage(val *EventMessage) *NullableEventMessage {
	return &NullableEventMessage{value: val, isSet: true}
}

func (v NullableEventMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


