
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TwoStepAuthRequestInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoStepAuthRequestInfoResponse{}

// TwoStepAuthRequestInfoResponse struct for TwoStepAuthRequestInfoResponse
type TwoStepAuthRequestInfoResponse struct {
	QrCode NullableString `json:"qr_code,omitempty"`
	Secret NullableString `json:"secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoStepAuthRequestInfoResponse TwoStepAuthRequestInfoResponse

// NewTwoStepAuthRequestInfoResponse instantiates a new TwoStepAuthRequestInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoStepAuthRequestInfoResponse() *TwoStepAuthRequestInfoResponse {
	this := TwoStepAuthRequestInfoResponse{}
	return &this
}

// NewTwoStepAuthRequestInfoResponseWithDefaults instantiates a new TwoStepAuthRequestInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoStepAuthRequestInfoResponseWithDefaults() *TwoStepAuthRequestInfoResponse {
	this := TwoStepAuthRequestInfoResponse{}
	return &this
}

// GetQrCode returns the QrCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoStepAuthRequestInfoResponse) GetQrCode() string {
	if o == nil || IsNil(o.QrCode.Get()) {
		var ret string
		return ret
	}
	return *o.QrCode.Get()
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoStepAuthRequestInfoResponse) GetQrCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCode.Get(), o.QrCode.IsSet()
}

// HasQrCode returns a boolean if a field has been set.
func (o *TwoStepAuthRequestInfoResponse) HasQrCode() bool {
	if o != nil && o.QrCode.IsSet() {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given NullableString and assigns it to the QrCode field.
func (o *TwoStepAuthRequestInfoResponse) SetQrCode(v string) {
	o.QrCode.Set(&v)
}
// SetQrCodeNil sets the value for QrCode to be an explicit nil
func (o *TwoStepAuthRequestInfoResponse) SetQrCodeNil() {
	o.QrCode.Set(nil)
}

// UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil
func (o *TwoStepAuthRequestInfoResponse) UnsetQrCode() {
	o.QrCode.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoStepAuthRequestInfoResponse) GetSecret() string {
	if o == nil || IsNil(o.Secret.Get()) {
		var ret string
		return ret
	}
	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoStepAuthRequestInfoResponse) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// HasSecret returns a boolean if a field has been set.
func (o *TwoStepAuthRequestInfoResponse) HasSecret() bool {
	if o != nil && o.Secret.IsSet() {
		return true
	}

	return false
}

// SetSecret gets a reference to the given NullableString and assigns it to the Secret field.
func (o *TwoStepAuthRequestInfoResponse) SetSecret(v string) {
	o.Secret.Set(&v)
}
// SetSecretNil sets the value for Secret to be an explicit nil
func (o *TwoStepAuthRequestInfoResponse) SetSecretNil() {
	o.Secret.Set(nil)
}

// UnsetSecret ensures that no value is present for Secret, not even an explicit nil
func (o *TwoStepAuthRequestInfoResponse) UnsetSecret() {
	o.Secret.Unset()
}

func (o TwoStepAuthRequestInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoStepAuthRequestInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.QrCode.IsSet() {
		toSerialize["qr_code"] = o.QrCode.Get()
	}
	if o.Secret.IsSet() {
		toSerialize["secret"] = o.Secret.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoStepAuthRequestInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varTwoStepAuthRequestInfoResponse := _TwoStepAuthRequestInfoResponse{}

	err = json.Unmarshal(data, &varTwoStepAuthRequestInfoResponse)

	if err != nil {
		return err
	}

	*o = TwoStepAuthRequestInfoResponse(varTwoStepAuthRequestInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "qr_code")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTwoStepAuthRequestInfoResponse struct {
	value *TwoStepAuthRequestInfoResponse
	isSet bool
}

func (v NullableTwoStepAuthRequestInfoResponse) Get() *TwoStepAuthRequestInfoResponse {
	return v.value
}

func (v *NullableTwoStepAuthRequestInfoResponse) Set(val *TwoStepAuthRequestInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoStepAuthRequestInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoStepAuthRequestInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoStepAuthRequestInfoResponse(val *TwoStepAuthRequestInfoResponse) *NullableTwoStepAuthRequestInfoResponse {
	return &NullableTwoStepAuthRequestInfoResponse{value: val, isSet: true}
}

func (v NullableTwoStepAuthRequestInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoStepAuthRequestInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


