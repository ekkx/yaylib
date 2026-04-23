
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ReadAttachmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAttachmentRequest{}

// ReadAttachmentRequest struct for ReadAttachmentRequest
type ReadAttachmentRequest struct {
	AttachmentMsgIds []int64 `json:"attachment_msg_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadAttachmentRequest ReadAttachmentRequest

// NewReadAttachmentRequest instantiates a new ReadAttachmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAttachmentRequest() *ReadAttachmentRequest {
	this := ReadAttachmentRequest{}
	return &this
}

// NewReadAttachmentRequestWithDefaults instantiates a new ReadAttachmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAttachmentRequestWithDefaults() *ReadAttachmentRequest {
	this := ReadAttachmentRequest{}
	return &this
}

// GetAttachmentMsgIds returns the AttachmentMsgIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReadAttachmentRequest) GetAttachmentMsgIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.AttachmentMsgIds
}

// GetAttachmentMsgIdsOk returns a tuple with the AttachmentMsgIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReadAttachmentRequest) GetAttachmentMsgIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AttachmentMsgIds) {
		return nil, false
	}
	return o.AttachmentMsgIds, true
}

// HasAttachmentMsgIds returns a boolean if a field has been set.
func (o *ReadAttachmentRequest) HasAttachmentMsgIds() bool {
	if o != nil && !IsNil(o.AttachmentMsgIds) {
		return true
	}

	return false
}

// SetAttachmentMsgIds gets a reference to the given []int64 and assigns it to the AttachmentMsgIds field.
func (o *ReadAttachmentRequest) SetAttachmentMsgIds(v []int64) {
	o.AttachmentMsgIds = v
}

func (o ReadAttachmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAttachmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentMsgIds != nil {
		toSerialize["attachment_msg_ids"] = o.AttachmentMsgIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReadAttachmentRequest) UnmarshalJSON(data []byte) (err error) {
	varReadAttachmentRequest := _ReadAttachmentRequest{}

	err = json.Unmarshal(data, &varReadAttachmentRequest)

	if err != nil {
		return err
	}

	*o = ReadAttachmentRequest(varReadAttachmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attachment_msg_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadAttachmentRequest struct {
	value *ReadAttachmentRequest
	isSet bool
}

func (v NullableReadAttachmentRequest) Get() *ReadAttachmentRequest {
	return v.value
}

func (v *NullableReadAttachmentRequest) Set(val *ReadAttachmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAttachmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAttachmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAttachmentRequest(val *ReadAttachmentRequest) *NullableReadAttachmentRequest {
	return &NullableReadAttachmentRequest{value: val, isSet: true}
}

func (v NullableReadAttachmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAttachmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


