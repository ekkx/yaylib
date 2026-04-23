
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TwoStepAuthRequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoStepAuthRequestInfo{}

// TwoStepAuthRequestInfo struct for TwoStepAuthRequestInfo
type TwoStepAuthRequestInfo struct {
	QrCodeUrl NullableString `json:"qr_code_url,omitempty"`
	SetupKey NullableString `json:"setup_key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoStepAuthRequestInfo TwoStepAuthRequestInfo

// NewTwoStepAuthRequestInfo instantiates a new TwoStepAuthRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoStepAuthRequestInfo() *TwoStepAuthRequestInfo {
	this := TwoStepAuthRequestInfo{}
	return &this
}

// NewTwoStepAuthRequestInfoWithDefaults instantiates a new TwoStepAuthRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoStepAuthRequestInfoWithDefaults() *TwoStepAuthRequestInfo {
	this := TwoStepAuthRequestInfo{}
	return &this
}

// GetQrCodeUrl returns the QrCodeUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoStepAuthRequestInfo) GetQrCodeUrl() string {
	if o == nil || IsNil(o.QrCodeUrl.Get()) {
		var ret string
		return ret
	}
	return *o.QrCodeUrl.Get()
}

// GetQrCodeUrlOk returns a tuple with the QrCodeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoStepAuthRequestInfo) GetQrCodeUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCodeUrl.Get(), o.QrCodeUrl.IsSet()
}

// HasQrCodeUrl returns a boolean if a field has been set.
func (o *TwoStepAuthRequestInfo) HasQrCodeUrl() bool {
	if o != nil && o.QrCodeUrl.IsSet() {
		return true
	}

	return false
}

// SetQrCodeUrl gets a reference to the given NullableString and assigns it to the QrCodeUrl field.
func (o *TwoStepAuthRequestInfo) SetQrCodeUrl(v string) {
	o.QrCodeUrl.Set(&v)
}
// SetQrCodeUrlNil sets the value for QrCodeUrl to be an explicit nil
func (o *TwoStepAuthRequestInfo) SetQrCodeUrlNil() {
	o.QrCodeUrl.Set(nil)
}

// UnsetQrCodeUrl ensures that no value is present for QrCodeUrl, not even an explicit nil
func (o *TwoStepAuthRequestInfo) UnsetQrCodeUrl() {
	o.QrCodeUrl.Unset()
}

// GetSetupKey returns the SetupKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoStepAuthRequestInfo) GetSetupKey() string {
	if o == nil || IsNil(o.SetupKey.Get()) {
		var ret string
		return ret
	}
	return *o.SetupKey.Get()
}

// GetSetupKeyOk returns a tuple with the SetupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoStepAuthRequestInfo) GetSetupKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SetupKey.Get(), o.SetupKey.IsSet()
}

// HasSetupKey returns a boolean if a field has been set.
func (o *TwoStepAuthRequestInfo) HasSetupKey() bool {
	if o != nil && o.SetupKey.IsSet() {
		return true
	}

	return false
}

// SetSetupKey gets a reference to the given NullableString and assigns it to the SetupKey field.
func (o *TwoStepAuthRequestInfo) SetSetupKey(v string) {
	o.SetupKey.Set(&v)
}
// SetSetupKeyNil sets the value for SetupKey to be an explicit nil
func (o *TwoStepAuthRequestInfo) SetSetupKeyNil() {
	o.SetupKey.Set(nil)
}

// UnsetSetupKey ensures that no value is present for SetupKey, not even an explicit nil
func (o *TwoStepAuthRequestInfo) UnsetSetupKey() {
	o.SetupKey.Unset()
}

func (o TwoStepAuthRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoStepAuthRequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.QrCodeUrl.IsSet() {
		toSerialize["qr_code_url"] = o.QrCodeUrl.Get()
	}
	if o.SetupKey.IsSet() {
		toSerialize["setup_key"] = o.SetupKey.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoStepAuthRequestInfo) UnmarshalJSON(data []byte) (err error) {
	varTwoStepAuthRequestInfo := _TwoStepAuthRequestInfo{}

	err = json.Unmarshal(data, &varTwoStepAuthRequestInfo)

	if err != nil {
		return err
	}

	*o = TwoStepAuthRequestInfo(varTwoStepAuthRequestInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "qr_code_url")
		delete(additionalProperties, "setup_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTwoStepAuthRequestInfo struct {
	value *TwoStepAuthRequestInfo
	isSet bool
}

func (v NullableTwoStepAuthRequestInfo) Get() *TwoStepAuthRequestInfo {
	return v.value
}

func (v *NullableTwoStepAuthRequestInfo) Set(val *TwoStepAuthRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoStepAuthRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoStepAuthRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoStepAuthRequestInfo(val *TwoStepAuthRequestInfo) *NullableTwoStepAuthRequestInfo {
	return &NullableTwoStepAuthRequestInfo{value: val, isSet: true}
}

func (v NullableTwoStepAuthRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoStepAuthRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


