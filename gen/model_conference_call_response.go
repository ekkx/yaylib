
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ConferenceCallResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceCallResponse{}

// ConferenceCallResponse struct for ConferenceCallResponse
type ConferenceCallResponse struct {
	ConferenceCall NullableRealmConferenceCall `json:"conference_call,omitempty"`
	ConferenceCallUserUuid NullableString `json:"conference_call_user_uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConferenceCallResponse ConferenceCallResponse

// NewConferenceCallResponse instantiates a new ConferenceCallResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceCallResponse() *ConferenceCallResponse {
	this := ConferenceCallResponse{}
	return &this
}

// NewConferenceCallResponseWithDefaults instantiates a new ConferenceCallResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceCallResponseWithDefaults() *ConferenceCallResponse {
	this := ConferenceCallResponse{}
	return &this
}

// GetConferenceCall returns the ConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCallResponse) GetConferenceCall() RealmConferenceCall {
	if o == nil || IsNil(o.ConferenceCall.Get()) {
		var ret RealmConferenceCall
		return ret
	}
	return *o.ConferenceCall.Get()
}

// GetConferenceCallOk returns a tuple with the ConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCallResponse) GetConferenceCallOk() (*RealmConferenceCall, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCall.Get(), o.ConferenceCall.IsSet()
}

// HasConferenceCall returns a boolean if a field has been set.
func (o *ConferenceCallResponse) HasConferenceCall() bool {
	if o != nil && o.ConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetConferenceCall gets a reference to the given NullableRealmConferenceCall and assigns it to the ConferenceCall field.
func (o *ConferenceCallResponse) SetConferenceCall(v RealmConferenceCall) {
	o.ConferenceCall.Set(&v)
}
// SetConferenceCallNil sets the value for ConferenceCall to be an explicit nil
func (o *ConferenceCallResponse) SetConferenceCallNil() {
	o.ConferenceCall.Set(nil)
}

// UnsetConferenceCall ensures that no value is present for ConferenceCall, not even an explicit nil
func (o *ConferenceCallResponse) UnsetConferenceCall() {
	o.ConferenceCall.Unset()
}

// GetConferenceCallUserUuid returns the ConferenceCallUserUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConferenceCallResponse) GetConferenceCallUserUuid() string {
	if o == nil || IsNil(o.ConferenceCallUserUuid.Get()) {
		var ret string
		return ret
	}
	return *o.ConferenceCallUserUuid.Get()
}

// GetConferenceCallUserUuidOk returns a tuple with the ConferenceCallUserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConferenceCallResponse) GetConferenceCallUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConferenceCallUserUuid.Get(), o.ConferenceCallUserUuid.IsSet()
}

// HasConferenceCallUserUuid returns a boolean if a field has been set.
func (o *ConferenceCallResponse) HasConferenceCallUserUuid() bool {
	if o != nil && o.ConferenceCallUserUuid.IsSet() {
		return true
	}

	return false
}

// SetConferenceCallUserUuid gets a reference to the given NullableString and assigns it to the ConferenceCallUserUuid field.
func (o *ConferenceCallResponse) SetConferenceCallUserUuid(v string) {
	o.ConferenceCallUserUuid.Set(&v)
}
// SetConferenceCallUserUuidNil sets the value for ConferenceCallUserUuid to be an explicit nil
func (o *ConferenceCallResponse) SetConferenceCallUserUuidNil() {
	o.ConferenceCallUserUuid.Set(nil)
}

// UnsetConferenceCallUserUuid ensures that no value is present for ConferenceCallUserUuid, not even an explicit nil
func (o *ConferenceCallResponse) UnsetConferenceCallUserUuid() {
	o.ConferenceCallUserUuid.Unset()
}

func (o ConferenceCallResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceCallResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConferenceCall.IsSet() {
		toSerialize["conference_call"] = o.ConferenceCall.Get()
	}
	if o.ConferenceCallUserUuid.IsSet() {
		toSerialize["conference_call_user_uuid"] = o.ConferenceCallUserUuid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConferenceCallResponse) UnmarshalJSON(data []byte) (err error) {
	varConferenceCallResponse := _ConferenceCallResponse{}

	err = json.Unmarshal(data, &varConferenceCallResponse)

	if err != nil {
		return err
	}

	*o = ConferenceCallResponse(varConferenceCallResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conference_call")
		delete(additionalProperties, "conference_call_user_uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConferenceCallResponse struct {
	value *ConferenceCallResponse
	isSet bool
}

func (v NullableConferenceCallResponse) Get() *ConferenceCallResponse {
	return v.value
}

func (v *NullableConferenceCallResponse) Set(val *ConferenceCallResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceCallResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceCallResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceCallResponse(val *ConferenceCallResponse) *NullableConferenceCallResponse {
	return &NullableConferenceCallResponse{value: val, isSet: true}
}

func (v NullableConferenceCallResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceCallResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


