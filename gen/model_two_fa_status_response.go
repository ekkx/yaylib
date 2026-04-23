
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TwoFAStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoFAStatusResponse{}

// TwoFAStatusResponse struct for TwoFAStatusResponse
type TwoFAStatusResponse struct {
	TwoFaAuthRequired NullableTwoFaAuthRequiredDTO `json:"two_fa_auth_required,omitempty"`
	TwoFaEnabled NullableBool `json:"two_fa_enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoFAStatusResponse TwoFAStatusResponse

// NewTwoFAStatusResponse instantiates a new TwoFAStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoFAStatusResponse() *TwoFAStatusResponse {
	this := TwoFAStatusResponse{}
	return &this
}

// NewTwoFAStatusResponseWithDefaults instantiates a new TwoFAStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoFAStatusResponseWithDefaults() *TwoFAStatusResponse {
	this := TwoFAStatusResponse{}
	return &this
}

// GetTwoFaAuthRequired returns the TwoFaAuthRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoFAStatusResponse) GetTwoFaAuthRequired() TwoFaAuthRequiredDTO {
	if o == nil || IsNil(o.TwoFaAuthRequired.Get()) {
		var ret TwoFaAuthRequiredDTO
		return ret
	}
	return *o.TwoFaAuthRequired.Get()
}

// GetTwoFaAuthRequiredOk returns a tuple with the TwoFaAuthRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoFAStatusResponse) GetTwoFaAuthRequiredOk() (*TwoFaAuthRequiredDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoFaAuthRequired.Get(), o.TwoFaAuthRequired.IsSet()
}

// HasTwoFaAuthRequired returns a boolean if a field has been set.
func (o *TwoFAStatusResponse) HasTwoFaAuthRequired() bool {
	if o != nil && o.TwoFaAuthRequired.IsSet() {
		return true
	}

	return false
}

// SetTwoFaAuthRequired gets a reference to the given NullableTwoFaAuthRequiredDTO and assigns it to the TwoFaAuthRequired field.
func (o *TwoFAStatusResponse) SetTwoFaAuthRequired(v TwoFaAuthRequiredDTO) {
	o.TwoFaAuthRequired.Set(&v)
}
// SetTwoFaAuthRequiredNil sets the value for TwoFaAuthRequired to be an explicit nil
func (o *TwoFAStatusResponse) SetTwoFaAuthRequiredNil() {
	o.TwoFaAuthRequired.Set(nil)
}

// UnsetTwoFaAuthRequired ensures that no value is present for TwoFaAuthRequired, not even an explicit nil
func (o *TwoFAStatusResponse) UnsetTwoFaAuthRequired() {
	o.TwoFaAuthRequired.Unset()
}

// GetTwoFaEnabled returns the TwoFaEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoFAStatusResponse) GetTwoFaEnabled() bool {
	if o == nil || IsNil(o.TwoFaEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.TwoFaEnabled.Get()
}

// GetTwoFaEnabledOk returns a tuple with the TwoFaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoFAStatusResponse) GetTwoFaEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoFaEnabled.Get(), o.TwoFaEnabled.IsSet()
}

// HasTwoFaEnabled returns a boolean if a field has been set.
func (o *TwoFAStatusResponse) HasTwoFaEnabled() bool {
	if o != nil && o.TwoFaEnabled.IsSet() {
		return true
	}

	return false
}

// SetTwoFaEnabled gets a reference to the given NullableBool and assigns it to the TwoFaEnabled field.
func (o *TwoFAStatusResponse) SetTwoFaEnabled(v bool) {
	o.TwoFaEnabled.Set(&v)
}
// SetTwoFaEnabledNil sets the value for TwoFaEnabled to be an explicit nil
func (o *TwoFAStatusResponse) SetTwoFaEnabledNil() {
	o.TwoFaEnabled.Set(nil)
}

// UnsetTwoFaEnabled ensures that no value is present for TwoFaEnabled, not even an explicit nil
func (o *TwoFAStatusResponse) UnsetTwoFaEnabled() {
	o.TwoFaEnabled.Unset()
}

func (o TwoFAStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoFAStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.TwoFaAuthRequired.IsSet() {
		toSerialize["two_fa_auth_required"] = o.TwoFaAuthRequired.Get()
	}
	if o.TwoFaEnabled.IsSet() {
		toSerialize["two_fa_enabled"] = o.TwoFaEnabled.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoFAStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varTwoFAStatusResponse := _TwoFAStatusResponse{}

	err = json.Unmarshal(data, &varTwoFAStatusResponse)

	if err != nil {
		return err
	}

	*o = TwoFAStatusResponse(varTwoFAStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "two_fa_auth_required")
		delete(additionalProperties, "two_fa_enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTwoFAStatusResponse struct {
	value *TwoFAStatusResponse
	isSet bool
}

func (v NullableTwoFAStatusResponse) Get() *TwoFAStatusResponse {
	return v.value
}

func (v *NullableTwoFAStatusResponse) Set(val *TwoFAStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoFAStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoFAStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoFAStatusResponse(val *TwoFAStatusResponse) *NullableTwoFAStatusResponse {
	return &NullableTwoFAStatusResponse{value: val, isSet: true}
}

func (v NullableTwoFAStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoFAStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


