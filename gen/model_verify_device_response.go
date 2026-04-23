
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the VerifyDeviceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyDeviceResponse{}

// VerifyDeviceResponse struct for VerifyDeviceResponse
type VerifyDeviceResponse struct {
	Verified NullableBool `json:"verified,omitempty"`
	VerifiedAt NullableString `json:"verified_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerifyDeviceResponse VerifyDeviceResponse

// NewVerifyDeviceResponse instantiates a new VerifyDeviceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyDeviceResponse() *VerifyDeviceResponse {
	this := VerifyDeviceResponse{}
	return &this
}

// NewVerifyDeviceResponseWithDefaults instantiates a new VerifyDeviceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyDeviceResponseWithDefaults() *VerifyDeviceResponse {
	this := VerifyDeviceResponse{}
	return &this
}

// GetVerified returns the Verified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyDeviceResponse) GetVerified() bool {
	if o == nil || IsNil(o.Verified.Get()) {
		var ret bool
		return ret
	}
	return *o.Verified.Get()
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyDeviceResponse) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verified.Get(), o.Verified.IsSet()
}

// HasVerified returns a boolean if a field has been set.
func (o *VerifyDeviceResponse) HasVerified() bool {
	if o != nil && o.Verified.IsSet() {
		return true
	}

	return false
}

// SetVerified gets a reference to the given NullableBool and assigns it to the Verified field.
func (o *VerifyDeviceResponse) SetVerified(v bool) {
	o.Verified.Set(&v)
}
// SetVerifiedNil sets the value for Verified to be an explicit nil
func (o *VerifyDeviceResponse) SetVerifiedNil() {
	o.Verified.Set(nil)
}

// UnsetVerified ensures that no value is present for Verified, not even an explicit nil
func (o *VerifyDeviceResponse) UnsetVerified() {
	o.Verified.Unset()
}

// GetVerifiedAt returns the VerifiedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VerifyDeviceResponse) GetVerifiedAt() string {
	if o == nil || IsNil(o.VerifiedAt.Get()) {
		var ret string
		return ret
	}
	return *o.VerifiedAt.Get()
}

// GetVerifiedAtOk returns a tuple with the VerifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VerifyDeviceResponse) GetVerifiedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerifiedAt.Get(), o.VerifiedAt.IsSet()
}

// HasVerifiedAt returns a boolean if a field has been set.
func (o *VerifyDeviceResponse) HasVerifiedAt() bool {
	if o != nil && o.VerifiedAt.IsSet() {
		return true
	}

	return false
}

// SetVerifiedAt gets a reference to the given NullableString and assigns it to the VerifiedAt field.
func (o *VerifyDeviceResponse) SetVerifiedAt(v string) {
	o.VerifiedAt.Set(&v)
}
// SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil
func (o *VerifyDeviceResponse) SetVerifiedAtNil() {
	o.VerifiedAt.Set(nil)
}

// UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil
func (o *VerifyDeviceResponse) UnsetVerifiedAt() {
	o.VerifiedAt.Unset()
}

func (o VerifyDeviceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyDeviceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Verified.IsSet() {
		toSerialize["verified"] = o.Verified.Get()
	}
	if o.VerifiedAt.IsSet() {
		toSerialize["verified_at"] = o.VerifiedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VerifyDeviceResponse) UnmarshalJSON(data []byte) (err error) {
	varVerifyDeviceResponse := _VerifyDeviceResponse{}

	err = json.Unmarshal(data, &varVerifyDeviceResponse)

	if err != nil {
		return err
	}

	*o = VerifyDeviceResponse(varVerifyDeviceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "verified")
		delete(additionalProperties, "verified_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVerifyDeviceResponse struct {
	value *VerifyDeviceResponse
	isSet bool
}

func (v NullableVerifyDeviceResponse) Get() *VerifyDeviceResponse {
	return v.value
}

func (v *NullableVerifyDeviceResponse) Set(val *VerifyDeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyDeviceResponse(val *VerifyDeviceResponse) *NullableVerifyDeviceResponse {
	return &NullableVerifyDeviceResponse{value: val, isSet: true}
}

func (v NullableVerifyDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


