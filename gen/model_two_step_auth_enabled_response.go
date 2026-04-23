
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TwoStepAuthEnabledResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoStepAuthEnabledResponse{}

// TwoStepAuthEnabledResponse struct for TwoStepAuthEnabledResponse
type TwoStepAuthEnabledResponse struct {
	RecoveryCodes []string `json:"recovery_codes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoStepAuthEnabledResponse TwoStepAuthEnabledResponse

// NewTwoStepAuthEnabledResponse instantiates a new TwoStepAuthEnabledResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoStepAuthEnabledResponse() *TwoStepAuthEnabledResponse {
	this := TwoStepAuthEnabledResponse{}
	return &this
}

// NewTwoStepAuthEnabledResponseWithDefaults instantiates a new TwoStepAuthEnabledResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoStepAuthEnabledResponseWithDefaults() *TwoStepAuthEnabledResponse {
	this := TwoStepAuthEnabledResponse{}
	return &this
}

// GetRecoveryCodes returns the RecoveryCodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoStepAuthEnabledResponse) GetRecoveryCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RecoveryCodes
}

// GetRecoveryCodesOk returns a tuple with the RecoveryCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoStepAuthEnabledResponse) GetRecoveryCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.RecoveryCodes) {
		return nil, false
	}
	return o.RecoveryCodes, true
}

// HasRecoveryCodes returns a boolean if a field has been set.
func (o *TwoStepAuthEnabledResponse) HasRecoveryCodes() bool {
	if o != nil && !IsNil(o.RecoveryCodes) {
		return true
	}

	return false
}

// SetRecoveryCodes gets a reference to the given []string and assigns it to the RecoveryCodes field.
func (o *TwoStepAuthEnabledResponse) SetRecoveryCodes(v []string) {
	o.RecoveryCodes = v
}

func (o TwoStepAuthEnabledResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoStepAuthEnabledResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryCodes != nil {
		toSerialize["recovery_codes"] = o.RecoveryCodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoStepAuthEnabledResponse) UnmarshalJSON(data []byte) (err error) {
	varTwoStepAuthEnabledResponse := _TwoStepAuthEnabledResponse{}

	err = json.Unmarshal(data, &varTwoStepAuthEnabledResponse)

	if err != nil {
		return err
	}

	*o = TwoStepAuthEnabledResponse(varTwoStepAuthEnabledResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recovery_codes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTwoStepAuthEnabledResponse struct {
	value *TwoStepAuthEnabledResponse
	isSet bool
}

func (v NullableTwoStepAuthEnabledResponse) Get() *TwoStepAuthEnabledResponse {
	return v.value
}

func (v *NullableTwoStepAuthEnabledResponse) Set(val *TwoStepAuthEnabledResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoStepAuthEnabledResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoStepAuthEnabledResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoStepAuthEnabledResponse(val *TwoStepAuthEnabledResponse) *NullableTwoStepAuthEnabledResponse {
	return &NullableTwoStepAuthEnabledResponse{value: val, isSet: true}
}

func (v NullableTwoStepAuthEnabledResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoStepAuthEnabledResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


