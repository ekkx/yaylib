
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MessageTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageTag{}

// MessageTag struct for MessageTag
type MessageTag struct {
	Length NullableInt32 `json:"length,omitempty"`
	Offset NullableInt32 `json:"offset,omitempty"`
	Type NullableString `json:"type,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MessageTag MessageTag

// NewMessageTag instantiates a new MessageTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageTag() *MessageTag {
	this := MessageTag{}
	return &this
}

// NewMessageTagWithDefaults instantiates a new MessageTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageTagWithDefaults() *MessageTag {
	this := MessageTag{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageTag) GetLength() int32 {
	if o == nil || IsNil(o.Length.Get()) {
		var ret int32
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageTag) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *MessageTag) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableInt32 and assigns it to the Length field.
func (o *MessageTag) SetLength(v int32) {
	o.Length.Set(&v)
}
// SetLengthNil sets the value for Length to be an explicit nil
func (o *MessageTag) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *MessageTag) UnsetLength() {
	o.Length.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageTag) GetOffset() int32 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageTag) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *MessageTag) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *MessageTag) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *MessageTag) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *MessageTag) UnsetOffset() {
	o.Offset.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageTag) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageTag) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MessageTag) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *MessageTag) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *MessageTag) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *MessageTag) UnsetType() {
	o.Type.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MessageTag) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MessageTag) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *MessageTag) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *MessageTag) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *MessageTag) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *MessageTag) UnsetUserId() {
	o.UserId.Unset()
}

func (o MessageTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Length.IsSet() {
		toSerialize["length"] = o.Length.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MessageTag) UnmarshalJSON(data []byte) (err error) {
	varMessageTag := _MessageTag{}

	err = json.Unmarshal(data, &varMessageTag)

	if err != nil {
		return err
	}

	*o = MessageTag(varMessageTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "length")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessageTag struct {
	value *MessageTag
	isSet bool
}

func (v NullableMessageTag) Get() *MessageTag {
	return v.value
}

func (v *NullableMessageTag) Set(val *MessageTag) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageTag) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageTag(val *MessageTag) *NullableMessageTag {
	return &NullableMessageTag{value: val, isSet: true}
}

func (v NullableMessageTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


