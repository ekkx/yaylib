
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ConferenceCallBumpParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceCallBumpParams{}

// ConferenceCallBumpParams struct for ConferenceCallBumpParams
type ConferenceCallBumpParams struct {
	ParticipantLimit NullableInt32 `json:"participant_limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConferenceCallBumpParams ConferenceCallBumpParams

// NewConferenceCallBumpParams instantiates a new ConferenceCallBumpParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceCallBumpParams() *ConferenceCallBumpParams {
	this := ConferenceCallBumpParams{}
	return &this
}

// NewConferenceCallBumpParamsWithDefaults instantiates a new ConferenceCallBumpParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceCallBumpParamsWithDefaults() *ConferenceCallBumpParams {
	this := ConferenceCallBumpParams{}
	return &this
}

// GetParticipantLimit returns the ParticipantLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCallBumpParams) GetParticipantLimit() int32 {
	if o == nil || IsNil(o.ParticipantLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.ParticipantLimit.Get()
}

// GetParticipantLimitOk returns a tuple with the ParticipantLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCallBumpParams) GetParticipantLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParticipantLimit.Get(), o.ParticipantLimit.IsSet()
}

// HasParticipantLimit returns a boolean if a field has been set.
func (o *ConferenceCallBumpParams) HasParticipantLimit() bool {
	if o != nil && o.ParticipantLimit.IsSet() {
		return true
	}

	return false
}

// SetParticipantLimit gets a reference to the given NullableInt32 and assigns it to the ParticipantLimit field.
func (o *ConferenceCallBumpParams) SetParticipantLimit(v int32) {
	o.ParticipantLimit.Set(&v)
}
// SetParticipantLimitNil sets the value for ParticipantLimit to be an explicit nil
func (o *ConferenceCallBumpParams) SetParticipantLimitNil() {
	o.ParticipantLimit.Set(nil)
}

// UnsetParticipantLimit ensures that no value is present for ParticipantLimit, not even an explicit nil
func (o *ConferenceCallBumpParams) UnsetParticipantLimit() {
	o.ParticipantLimit.Unset()
}

func (o ConferenceCallBumpParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceCallBumpParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ParticipantLimit.IsSet() {
		toSerialize["participant_limit"] = o.ParticipantLimit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConferenceCallBumpParams) UnmarshalJSON(data []byte) (err error) {
	varConferenceCallBumpParams := _ConferenceCallBumpParams{}

	err = json.Unmarshal(data, &varConferenceCallBumpParams)

	if err != nil {
		return err
	}

	*o = ConferenceCallBumpParams(varConferenceCallBumpParams)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "participant_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConferenceCallBumpParams struct {
	value *ConferenceCallBumpParams
	isSet bool
}

func (v NullableConferenceCallBumpParams) Get() *ConferenceCallBumpParams {
	return v.value
}

func (v *NullableConferenceCallBumpParams) Set(val *ConferenceCallBumpParams) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceCallBumpParams) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceCallBumpParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceCallBumpParams(val *ConferenceCallBumpParams) *NullableConferenceCallBumpParams {
	return &NullableConferenceCallBumpParams{value: val, isSet: true}
}

func (v NullableConferenceCallBumpParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceCallBumpParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


