
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the InvitationCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitationCode{}

// InvitationCode struct for InvitationCode
type InvitationCode struct {
	PossibleToEditAtMillis NullableInt64 `json:"possible_to_edit_at_millis,omitempty"`
	Value NullableString `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvitationCode InvitationCode

// NewInvitationCode instantiates a new InvitationCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationCode() *InvitationCode {
	this := InvitationCode{}
	return &this
}

// NewInvitationCodeWithDefaults instantiates a new InvitationCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationCodeWithDefaults() *InvitationCode {
	this := InvitationCode{}
	return &this
}

// GetPossibleToEditAtMillis returns the PossibleToEditAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationCode) GetPossibleToEditAtMillis() int64 {
	if o == nil || IsNil(o.PossibleToEditAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.PossibleToEditAtMillis.Get()
}

// GetPossibleToEditAtMillisOk returns a tuple with the PossibleToEditAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationCode) GetPossibleToEditAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PossibleToEditAtMillis.Get(), o.PossibleToEditAtMillis.IsSet()
}

// HasPossibleToEditAtMillis returns a boolean if a field has been set.
func (o *InvitationCode) HasPossibleToEditAtMillis() bool {
	if o != nil && o.PossibleToEditAtMillis.IsSet() {
		return true
	}

	return false
}

// SetPossibleToEditAtMillis gets a reference to the given NullableInt64 and assigns it to the PossibleToEditAtMillis field.
func (o *InvitationCode) SetPossibleToEditAtMillis(v int64) {
	o.PossibleToEditAtMillis.Set(&v)
}
// SetPossibleToEditAtMillisNil sets the value for PossibleToEditAtMillis to be an explicit nil
func (o *InvitationCode) SetPossibleToEditAtMillisNil() {
	o.PossibleToEditAtMillis.Set(nil)
}

// UnsetPossibleToEditAtMillis ensures that no value is present for PossibleToEditAtMillis, not even an explicit nil
func (o *InvitationCode) UnsetPossibleToEditAtMillis() {
	o.PossibleToEditAtMillis.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitationCode) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitationCode) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *InvitationCode) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *InvitationCode) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *InvitationCode) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *InvitationCode) UnsetValue() {
	o.Value.Unset()
}

func (o InvitationCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitationCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PossibleToEditAtMillis.IsSet() {
		toSerialize["possible_to_edit_at_millis"] = o.PossibleToEditAtMillis.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvitationCode) UnmarshalJSON(data []byte) (err error) {
	varInvitationCode := _InvitationCode{}

	err = json.Unmarshal(data, &varInvitationCode)

	if err != nil {
		return err
	}

	*o = InvitationCode(varInvitationCode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "possible_to_edit_at_millis")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvitationCode struct {
	value *InvitationCode
	isSet bool
}

func (v NullableInvitationCode) Get() *InvitationCode {
	return v.value
}

func (v *NullableInvitationCode) Set(val *InvitationCode) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitationCode) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitationCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitationCode(val *InvitationCode) *NullableInvitationCode {
	return &NullableInvitationCode{value: val, isSet: true}
}

func (v NullableInvitationCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitationCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


