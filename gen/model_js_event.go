
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the JSEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JSEvent{}

// JSEvent struct for JSEvent
type JSEvent struct {
	Args NullableString `json:"args,omitempty"`
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JSEvent JSEvent

// NewJSEvent instantiates a new JSEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSEvent() *JSEvent {
	this := JSEvent{}
	return &this
}

// NewJSEventWithDefaults instantiates a new JSEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSEventWithDefaults() *JSEvent {
	this := JSEvent{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSEvent) GetArgs() string {
	if o == nil || IsNil(o.Args.Get()) {
		var ret string
		return ret
	}
	return *o.Args.Get()
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSEvent) GetArgsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Args.Get(), o.Args.IsSet()
}

// HasArgs returns a boolean if a field has been set.
func (o *JSEvent) HasArgs() bool {
	if o != nil && o.Args.IsSet() {
		return true
	}

	return false
}

// SetArgs gets a reference to the given NullableString and assigns it to the Args field.
func (o *JSEvent) SetArgs(v string) {
	o.Args.Set(&v)
}
// SetArgsNil sets the value for Args to be an explicit nil
func (o *JSEvent) SetArgsNil() {
	o.Args.Set(nil)
}

// UnsetArgs ensures that no value is present for Args, not even an explicit nil
func (o *JSEvent) UnsetArgs() {
	o.Args.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSEvent) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *JSEvent) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *JSEvent) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *JSEvent) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *JSEvent) UnsetName() {
	o.Name.Unset()
}

func (o JSEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JSEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Args.IsSet() {
		toSerialize["args"] = o.Args.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JSEvent) UnmarshalJSON(data []byte) (err error) {
	varJSEvent := _JSEvent{}

	err = json.Unmarshal(data, &varJSEvent)

	if err != nil {
		return err
	}

	*o = JSEvent(varJSEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "args")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJSEvent struct {
	value *JSEvent
	isSet bool
}

func (v NullableJSEvent) Get() *JSEvent {
	return v.value
}

func (v *NullableJSEvent) Set(val *JSEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableJSEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableJSEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSEvent(val *JSEvent) *NullableJSEvent {
	return &NullableJSEvent{value: val, isSet: true}
}

func (v NullableJSEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


