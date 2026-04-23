
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplDetails{}

// EmplDetails struct for EmplDetails
type EmplDetails struct {
	Regular NullableFloat64 `json:"regular,omitempty"`
	Rewarded NullableFloat64 `json:"rewarded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplDetails EmplDetails

// NewEmplDetails instantiates a new EmplDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplDetails() *EmplDetails {
	this := EmplDetails{}
	return &this
}

// NewEmplDetailsWithDefaults instantiates a new EmplDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplDetailsWithDefaults() *EmplDetails {
	this := EmplDetails{}
	return &this
}

// GetRegular returns the Regular field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplDetails) GetRegular() float64 {
	if o == nil || IsNil(o.Regular.Get()) {
		var ret float64
		return ret
	}
	return *o.Regular.Get()
}

// GetRegularOk returns a tuple with the Regular field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplDetails) GetRegularOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regular.Get(), o.Regular.IsSet()
}

// HasRegular returns a boolean if a field has been set.
func (o *EmplDetails) HasRegular() bool {
	if o != nil && o.Regular.IsSet() {
		return true
	}

	return false
}

// SetRegular gets a reference to the given NullableFloat64 and assigns it to the Regular field.
func (o *EmplDetails) SetRegular(v float64) {
	o.Regular.Set(&v)
}
// SetRegularNil sets the value for Regular to be an explicit nil
func (o *EmplDetails) SetRegularNil() {
	o.Regular.Set(nil)
}

// UnsetRegular ensures that no value is present for Regular, not even an explicit nil
func (o *EmplDetails) UnsetRegular() {
	o.Regular.Unset()
}

// GetRewarded returns the Rewarded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplDetails) GetRewarded() float64 {
	if o == nil || IsNil(o.Rewarded.Get()) {
		var ret float64
		return ret
	}
	return *o.Rewarded.Get()
}

// GetRewardedOk returns a tuple with the Rewarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplDetails) GetRewardedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rewarded.Get(), o.Rewarded.IsSet()
}

// HasRewarded returns a boolean if a field has been set.
func (o *EmplDetails) HasRewarded() bool {
	if o != nil && o.Rewarded.IsSet() {
		return true
	}

	return false
}

// SetRewarded gets a reference to the given NullableFloat64 and assigns it to the Rewarded field.
func (o *EmplDetails) SetRewarded(v float64) {
	o.Rewarded.Set(&v)
}
// SetRewardedNil sets the value for Rewarded to be an explicit nil
func (o *EmplDetails) SetRewardedNil() {
	o.Rewarded.Set(nil)
}

// UnsetRewarded ensures that no value is present for Rewarded, not even an explicit nil
func (o *EmplDetails) UnsetRewarded() {
	o.Rewarded.Unset()
}

func (o EmplDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Regular.IsSet() {
		toSerialize["regular"] = o.Regular.Get()
	}
	if o.Rewarded.IsSet() {
		toSerialize["rewarded"] = o.Rewarded.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplDetails) UnmarshalJSON(data []byte) (err error) {
	varEmplDetails := _EmplDetails{}

	err = json.Unmarshal(data, &varEmplDetails)

	if err != nil {
		return err
	}

	*o = EmplDetails(varEmplDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "regular")
		delete(additionalProperties, "rewarded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplDetails struct {
	value *EmplDetails
	isSet bool
}

func (v NullableEmplDetails) Get() *EmplDetails {
	return v.value
}

func (v *NullableEmplDetails) Set(val *EmplDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplDetails(val *EmplDetails) *NullableEmplDetails {
	return &NullableEmplDetails{value: val, isSet: true}
}

func (v NullableEmplDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


