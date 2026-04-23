
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the InvitationCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationCodeResponse{}

// InvitationCodeResponse struct for InvitationCodeResponse
type InvitationCodeResponse struct {
	InvitationCode NullableCodeResponse `json:"invitation_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvitationCodeResponse InvitationCodeResponse

// NewInvitationCodeResponse instantiates a new InvitationCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationCodeResponse() *InvitationCodeResponse {
	this := InvitationCodeResponse{}
	return &this
}

// NewInvitationCodeResponseWithDefaults instantiates a new InvitationCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationCodeResponseWithDefaults() *InvitationCodeResponse {
	this := InvitationCodeResponse{}
	return &this
}

// GetInvitationCode returns the InvitationCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationCodeResponse) GetInvitationCode() CodeResponse {
	if o == nil || IsNil(o.InvitationCode.Get()) {
		var ret CodeResponse
		return ret
	}
	return *o.InvitationCode.Get()
}

// GetInvitationCodeOk returns a tuple with the InvitationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationCodeResponse) GetInvitationCodeOk() (*CodeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvitationCode.Get(), o.InvitationCode.IsSet()
}

// HasInvitationCode returns a boolean if a field has been set.
func (o *InvitationCodeResponse) HasInvitationCode() bool {
	if o != nil && o.InvitationCode.IsSet() {
		return true
	}

	return false
}

// SetInvitationCode gets a reference to the given NullableCodeResponse and assigns it to the InvitationCode field.
func (o *InvitationCodeResponse) SetInvitationCode(v CodeResponse) {
	o.InvitationCode.Set(&v)
}
// SetInvitationCodeNil sets the value for InvitationCode to be an explicit nil
func (o *InvitationCodeResponse) SetInvitationCodeNil() {
	o.InvitationCode.Set(nil)
}

// UnsetInvitationCode ensures that no value is present for InvitationCode, not even an explicit nil
func (o *InvitationCodeResponse) UnsetInvitationCode() {
	o.InvitationCode.Unset()
}

func (o InvitationCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitationCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InvitationCode.IsSet() {
		toSerialize["invitation_code"] = o.InvitationCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvitationCodeResponse) UnmarshalJSON(data []byte) (err error) {
	varInvitationCodeResponse := _InvitationCodeResponse{}

	err = json.Unmarshal(data, &varInvitationCodeResponse)

	if err != nil {
		return err
	}

	*o = InvitationCodeResponse(varInvitationCodeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invitation_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvitationCodeResponse struct {
	value *InvitationCodeResponse
	isSet bool
}

func (v NullableInvitationCodeResponse) Get() *InvitationCodeResponse {
	return v.value
}

func (v *NullableInvitationCodeResponse) Set(val *InvitationCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationCodeResponse(val *InvitationCodeResponse) *NullableInvitationCodeResponse {
	return &NullableInvitationCodeResponse{value: val, isSet: true}
}

func (v NullableInvitationCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


