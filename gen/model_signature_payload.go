
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SignaturePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignaturePayload{}

// SignaturePayload struct for SignaturePayload
type SignaturePayload struct {
	Action NullableString `json:"action,omitempty"`
	CallId NullableInt64 `json:"call_id,omitempty"`
	ReceiverUuid NullableString `json:"receiver_uuid,omitempty"`
	SenderUuid NullableString `json:"sender_uuid,omitempty"`
	Signature NullableString `json:"signature,omitempty"`
	Timestamp NullableInt64 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignaturePayload SignaturePayload

// NewSignaturePayload instantiates a new SignaturePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignaturePayload() *SignaturePayload {
	this := SignaturePayload{}
	return &this
}

// NewSignaturePayloadWithDefaults instantiates a new SignaturePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignaturePayloadWithDefaults() *SignaturePayload {
	this := SignaturePayload{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignaturePayload) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignaturePayload) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *SignaturePayload) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *SignaturePayload) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *SignaturePayload) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *SignaturePayload) UnsetAction() {
	o.Action.Unset()
}

// GetCallId returns the CallId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignaturePayload) GetCallId() int64 {
	if o == nil || IsNil(o.CallId.Get()) {
		var ret int64
		return ret
	}
	return *o.CallId.Get()
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignaturePayload) GetCallIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallId.Get(), o.CallId.IsSet()
}

// HasCallId returns a boolean if a field has been set.
func (o *SignaturePayload) HasCallId() bool {
	if o != nil && o.CallId.IsSet() {
		return true
	}

	return false
}

// SetCallId gets a reference to the given NullableInt64 and assigns it to the CallId field.
func (o *SignaturePayload) SetCallId(v int64) {
	o.CallId.Set(&v)
}
// SetCallIdNil sets the value for CallId to be an explicit nil
func (o *SignaturePayload) SetCallIdNil() {
	o.CallId.Set(nil)
}

// UnsetCallId ensures that no value is present for CallId, not even an explicit nil
func (o *SignaturePayload) UnsetCallId() {
	o.CallId.Unset()
}

// GetReceiverUuid returns the ReceiverUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignaturePayload) GetReceiverUuid() string {
	if o == nil || IsNil(o.ReceiverUuid.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiverUuid.Get()
}

// GetReceiverUuidOk returns a tuple with the ReceiverUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignaturePayload) GetReceiverUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverUuid.Get(), o.ReceiverUuid.IsSet()
}

// HasReceiverUuid returns a boolean if a field has been set.
func (o *SignaturePayload) HasReceiverUuid() bool {
	if o != nil && o.ReceiverUuid.IsSet() {
		return true
	}

	return false
}

// SetReceiverUuid gets a reference to the given NullableString and assigns it to the ReceiverUuid field.
func (o *SignaturePayload) SetReceiverUuid(v string) {
	o.ReceiverUuid.Set(&v)
}
// SetReceiverUuidNil sets the value for ReceiverUuid to be an explicit nil
func (o *SignaturePayload) SetReceiverUuidNil() {
	o.ReceiverUuid.Set(nil)
}

// UnsetReceiverUuid ensures that no value is present for ReceiverUuid, not even an explicit nil
func (o *SignaturePayload) UnsetReceiverUuid() {
	o.ReceiverUuid.Unset()
}

// GetSenderUuid returns the SenderUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignaturePayload) GetSenderUuid() string {
	if o == nil || IsNil(o.SenderUuid.Get()) {
		var ret string
		return ret
	}
	return *o.SenderUuid.Get()
}

// GetSenderUuidOk returns a tuple with the SenderUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignaturePayload) GetSenderUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderUuid.Get(), o.SenderUuid.IsSet()
}

// HasSenderUuid returns a boolean if a field has been set.
func (o *SignaturePayload) HasSenderUuid() bool {
	if o != nil && o.SenderUuid.IsSet() {
		return true
	}

	return false
}

// SetSenderUuid gets a reference to the given NullableString and assigns it to the SenderUuid field.
func (o *SignaturePayload) SetSenderUuid(v string) {
	o.SenderUuid.Set(&v)
}
// SetSenderUuidNil sets the value for SenderUuid to be an explicit nil
func (o *SignaturePayload) SetSenderUuidNil() {
	o.SenderUuid.Set(nil)
}

// UnsetSenderUuid ensures that no value is present for SenderUuid, not even an explicit nil
func (o *SignaturePayload) UnsetSenderUuid() {
	o.SenderUuid.Unset()
}

// GetSignature returns the Signature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignaturePayload) GetSignature() string {
	if o == nil || IsNil(o.Signature.Get()) {
		var ret string
		return ret
	}
	return *o.Signature.Get()
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignaturePayload) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signature.Get(), o.Signature.IsSet()
}

// HasSignature returns a boolean if a field has been set.
func (o *SignaturePayload) HasSignature() bool {
	if o != nil && o.Signature.IsSet() {
		return true
	}

	return false
}

// SetSignature gets a reference to the given NullableString and assigns it to the Signature field.
func (o *SignaturePayload) SetSignature(v string) {
	o.Signature.Set(&v)
}
// SetSignatureNil sets the value for Signature to be an explicit nil
func (o *SignaturePayload) SetSignatureNil() {
	o.Signature.Set(nil)
}

// UnsetSignature ensures that no value is present for Signature, not even an explicit nil
func (o *SignaturePayload) UnsetSignature() {
	o.Signature.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SignaturePayload) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SignaturePayload) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SignaturePayload) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *SignaturePayload) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *SignaturePayload) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *SignaturePayload) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o SignaturePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignaturePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.CallId.IsSet() {
		toSerialize["call_id"] = o.CallId.Get()
	}
	if o.ReceiverUuid.IsSet() {
		toSerialize["receiver_uuid"] = o.ReceiverUuid.Get()
	}
	if o.SenderUuid.IsSet() {
		toSerialize["sender_uuid"] = o.SenderUuid.Get()
	}
	if o.Signature.IsSet() {
		toSerialize["signature"] = o.Signature.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignaturePayload) UnmarshalJSON(data []byte) (err error) {
	varSignaturePayload := _SignaturePayload{}

	err = json.Unmarshal(data, &varSignaturePayload)

	if err != nil {
		return err
	}

	*o = SignaturePayload(varSignaturePayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "call_id")
		delete(additionalProperties, "receiver_uuid")
		delete(additionalProperties, "sender_uuid")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignaturePayload struct {
	value *SignaturePayload
	isSet bool
}

func (v NullableSignaturePayload) Get() *SignaturePayload {
	return v.value
}

func (v *NullableSignaturePayload) Set(val *SignaturePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSignaturePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSignaturePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignaturePayload(val *SignaturePayload) *NullableSignaturePayload {
	return &NullableSignaturePayload{value: val, isSet: true}
}

func (v NullableSignaturePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignaturePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


