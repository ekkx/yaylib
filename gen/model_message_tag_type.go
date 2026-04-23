
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MessageTagType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageTagType{}

// MessageTagType struct for MessageTagType
type MessageTagType struct {
	ApiValue NullableString `json:"api_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MessageTagType MessageTagType

// NewMessageTagType instantiates a new MessageTagType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageTagType() *MessageTagType {
	this := MessageTagType{}
	return &this
}

// NewMessageTagTypeWithDefaults instantiates a new MessageTagType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageTagTypeWithDefaults() *MessageTagType {
	this := MessageTagType{}
	return &this
}

// GetApiValue returns the ApiValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageTagType) GetApiValue() string {
	if o == nil || IsNil(o.ApiValue.Get()) {
		var ret string
		return ret
	}
	return *o.ApiValue.Get()
}

// GetApiValueOk returns a tuple with the ApiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageTagType) GetApiValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiValue.Get(), o.ApiValue.IsSet()
}

// HasApiValue returns a boolean if a field has been set.
func (o *MessageTagType) HasApiValue() bool {
	if o != nil && o.ApiValue.IsSet() {
		return true
	}

	return false
}

// SetApiValue gets a reference to the given NullableString and assigns it to the ApiValue field.
func (o *MessageTagType) SetApiValue(v string) {
	o.ApiValue.Set(&v)
}
// SetApiValueNil sets the value for ApiValue to be an explicit nil
func (o *MessageTagType) SetApiValueNil() {
	o.ApiValue.Set(nil)
}

// UnsetApiValue ensures that no value is present for ApiValue, not even an explicit nil
func (o *MessageTagType) UnsetApiValue() {
	o.ApiValue.Unset()
}

func (o MessageTagType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageTagType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiValue.IsSet() {
		toSerialize["api_value"] = o.ApiValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MessageTagType) UnmarshalJSON(data []byte) (err error) {
	varMessageTagType := _MessageTagType{}

	err = json.Unmarshal(data, &varMessageTagType)

	if err != nil {
		return err
	}

	*o = MessageTagType(varMessageTagType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessageTagType struct {
	value *MessageTagType
	isSet bool
}

func (v NullableMessageTagType) Get() *MessageTagType {
	return v.value
}

func (v *NullableMessageTagType) Set(val *MessageTagType) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageTagType) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageTagType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageTagType(val *MessageTagType) *NullableMessageTagType {
	return &NullableMessageTagType{value: val, isSet: true}
}

func (v NullableMessageTagType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageTagType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


