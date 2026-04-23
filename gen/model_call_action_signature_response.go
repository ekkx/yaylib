
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CallActionSignatureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallActionSignatureResponse{}

// CallActionSignatureResponse struct for CallActionSignatureResponse
type CallActionSignatureResponse struct {
	SignaturePayload NullableSignaturePayload `json:"signature_payload,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CallActionSignatureResponse CallActionSignatureResponse

// NewCallActionSignatureResponse instantiates a new CallActionSignatureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallActionSignatureResponse() *CallActionSignatureResponse {
	this := CallActionSignatureResponse{}
	return &this
}

// NewCallActionSignatureResponseWithDefaults instantiates a new CallActionSignatureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallActionSignatureResponseWithDefaults() *CallActionSignatureResponse {
	this := CallActionSignatureResponse{}
	return &this
}

// GetSignaturePayload returns the SignaturePayload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallActionSignatureResponse) GetSignaturePayload() SignaturePayload {
	if o == nil || IsNil(o.SignaturePayload.Get()) {
		var ret SignaturePayload
		return ret
	}
	return *o.SignaturePayload.Get()
}

// GetSignaturePayloadOk returns a tuple with the SignaturePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallActionSignatureResponse) GetSignaturePayloadOk() (*SignaturePayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignaturePayload.Get(), o.SignaturePayload.IsSet()
}

// HasSignaturePayload returns a boolean if a field has been set.
func (o *CallActionSignatureResponse) HasSignaturePayload() bool {
	if o != nil && o.SignaturePayload.IsSet() {
		return true
	}

	return false
}

// SetSignaturePayload gets a reference to the given NullableSignaturePayload and assigns it to the SignaturePayload field.
func (o *CallActionSignatureResponse) SetSignaturePayload(v SignaturePayload) {
	o.SignaturePayload.Set(&v)
}
// SetSignaturePayloadNil sets the value for SignaturePayload to be an explicit nil
func (o *CallActionSignatureResponse) SetSignaturePayloadNil() {
	o.SignaturePayload.Set(nil)
}

// UnsetSignaturePayload ensures that no value is present for SignaturePayload, not even an explicit nil
func (o *CallActionSignatureResponse) UnsetSignaturePayload() {
	o.SignaturePayload.Unset()
}

func (o CallActionSignatureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallActionSignatureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SignaturePayload.IsSet() {
		toSerialize["signature_payload"] = o.SignaturePayload.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CallActionSignatureResponse) UnmarshalJSON(data []byte) (err error) {
	varCallActionSignatureResponse := _CallActionSignatureResponse{}

	err = json.Unmarshal(data, &varCallActionSignatureResponse)

	if err != nil {
		return err
	}

	*o = CallActionSignatureResponse(varCallActionSignatureResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signature_payload")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallActionSignatureResponse struct {
	value *CallActionSignatureResponse
	isSet bool
}

func (v NullableCallActionSignatureResponse) Get() *CallActionSignatureResponse {
	return v.value
}

func (v *NullableCallActionSignatureResponse) Set(val *CallActionSignatureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCallActionSignatureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCallActionSignatureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallActionSignatureResponse(val *CallActionSignatureResponse) *NullableCallActionSignatureResponse {
	return &NullableCallActionSignatureResponse{value: val, isSet: true}
}

func (v NullableCallActionSignatureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallActionSignatureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


