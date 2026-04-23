
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TwoStepAuthEnabled type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TwoStepAuthEnabled{}

// TwoStepAuthEnabled struct for TwoStepAuthEnabled
type TwoStepAuthEnabled struct {
	RecoveryCodeList []string `json:"recovery_code_list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TwoStepAuthEnabled TwoStepAuthEnabled

// NewTwoStepAuthEnabled instantiates a new TwoStepAuthEnabled object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTwoStepAuthEnabled() *TwoStepAuthEnabled {
	this := TwoStepAuthEnabled{}
	return &this
}

// NewTwoStepAuthEnabledWithDefaults instantiates a new TwoStepAuthEnabled object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTwoStepAuthEnabledWithDefaults() *TwoStepAuthEnabled {
	this := TwoStepAuthEnabled{}
	return &this
}

// GetRecoveryCodeList returns the RecoveryCodeList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TwoStepAuthEnabled) GetRecoveryCodeList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RecoveryCodeList
}

// GetRecoveryCodeListOk returns a tuple with the RecoveryCodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TwoStepAuthEnabled) GetRecoveryCodeListOk() ([]string, bool) {
	if o == nil || IsNil(o.RecoveryCodeList) {
		return nil, false
	}
	return o.RecoveryCodeList, true
}

// HasRecoveryCodeList returns a boolean if a field has been set.
func (o *TwoStepAuthEnabled) HasRecoveryCodeList() bool {
	if o != nil && !IsNil(o.RecoveryCodeList) {
		return true
	}

	return false
}

// SetRecoveryCodeList gets a reference to the given []string and assigns it to the RecoveryCodeList field.
func (o *TwoStepAuthEnabled) SetRecoveryCodeList(v []string) {
	o.RecoveryCodeList = v
}

func (o TwoStepAuthEnabled) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TwoStepAuthEnabled) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RecoveryCodeList != nil {
		toSerialize["recovery_code_list"] = o.RecoveryCodeList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TwoStepAuthEnabled) UnmarshalJSON(data []byte) (err error) {
	varTwoStepAuthEnabled := _TwoStepAuthEnabled{}

	err = json.Unmarshal(data, &varTwoStepAuthEnabled)

	if err != nil {
		return err
	}

	*o = TwoStepAuthEnabled(varTwoStepAuthEnabled)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recovery_code_list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTwoStepAuthEnabled struct {
	value *TwoStepAuthEnabled
	isSet bool
}

func (v NullableTwoStepAuthEnabled) Get() *TwoStepAuthEnabled {
	return v.value
}

func (v *NullableTwoStepAuthEnabled) Set(val *TwoStepAuthEnabled) {
	v.value = val
	v.isSet = true
}

func (v NullableTwoStepAuthEnabled) IsSet() bool {
	return v.isSet
}

func (v *NullableTwoStepAuthEnabled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwoStepAuthEnabled(val *TwoStepAuthEnabled) *NullableTwoStepAuthEnabled {
	return &NullableTwoStepAuthEnabled{value: val, isSet: true}
}

func (v NullableTwoStepAuthEnabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwoStepAuthEnabled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


