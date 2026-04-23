
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelMessageTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelMessageTag{}

// ModelMessageTag struct for ModelMessageTag
type ModelMessageTag struct {
	EndAt NullableInt32 `json:"end_at,omitempty"`
	Length NullableInt32 `json:"length,omitempty"`
	Offset NullableInt32 `json:"offset,omitempty"`
	Type NullableMessageTagType `json:"type,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelMessageTag ModelMessageTag

// NewModelMessageTag instantiates a new ModelMessageTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelMessageTag() *ModelMessageTag {
	this := ModelMessageTag{}
	return &this
}

// NewModelMessageTagWithDefaults instantiates a new ModelMessageTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelMessageTagWithDefaults() *ModelMessageTag {
	this := ModelMessageTag{}
	return &this
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelMessageTag) GetEndAt() int32 {
	if o == nil || IsNil(o.EndAt.Get()) {
		var ret int32
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelMessageTag) GetEndAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *ModelMessageTag) HasEndAt() bool {
	if o != nil && o.EndAt.IsSet() {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given NullableInt32 and assigns it to the EndAt field.
func (o *ModelMessageTag) SetEndAt(v int32) {
	o.EndAt.Set(&v)
}
// SetEndAtNil sets the value for EndAt to be an explicit nil
func (o *ModelMessageTag) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
func (o *ModelMessageTag) UnsetEndAt() {
	o.EndAt.Unset()
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelMessageTag) GetLength() int32 {
	if o == nil || IsNil(o.Length.Get()) {
		var ret int32
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelMessageTag) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *ModelMessageTag) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableInt32 and assigns it to the Length field.
func (o *ModelMessageTag) SetLength(v int32) {
	o.Length.Set(&v)
}
// SetLengthNil sets the value for Length to be an explicit nil
func (o *ModelMessageTag) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *ModelMessageTag) UnsetLength() {
	o.Length.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelMessageTag) GetOffset() int32 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int32
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelMessageTag) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *ModelMessageTag) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt32 and assigns it to the Offset field.
func (o *ModelMessageTag) SetOffset(v int32) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *ModelMessageTag) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *ModelMessageTag) UnsetOffset() {
	o.Offset.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelMessageTag) GetType() MessageTagType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret MessageTagType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelMessageTag) GetTypeOk() (*MessageTagType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ModelMessageTag) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableMessageTagType and assigns it to the Type field.
func (o *ModelMessageTag) SetType(v MessageTagType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ModelMessageTag) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ModelMessageTag) UnsetType() {
	o.Type.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelMessageTag) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelMessageTag) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ModelMessageTag) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *ModelMessageTag) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ModelMessageTag) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ModelMessageTag) UnsetUserId() {
	o.UserId.Unset()
}

func (o ModelMessageTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelMessageTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EndAt.IsSet() {
		toSerialize["end_at"] = o.EndAt.Get()
	}
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

func (o *ModelMessageTag) UnmarshalJSON(data []byte) (err error) {
	varModelMessageTag := _ModelMessageTag{}

	err = json.Unmarshal(data, &varModelMessageTag)

	if err != nil {
		return err
	}

	*o = ModelMessageTag(varModelMessageTag)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "end_at")
		delete(additionalProperties, "length")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelMessageTag struct {
	value *ModelMessageTag
	isSet bool
}

func (v NullableModelMessageTag) Get() *ModelMessageTag {
	return v.value
}

func (v *NullableModelMessageTag) Set(val *ModelMessageTag) {
	v.value = val
	v.isSet = true
}

func (v NullableModelMessageTag) IsSet() bool {
	return v.isSet
}

func (v *NullableModelMessageTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelMessageTag(val *ModelMessageTag) *NullableModelMessageTag {
	return &NullableModelMessageTag{value: val, isSet: true}
}

func (v NullableModelMessageTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelMessageTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


