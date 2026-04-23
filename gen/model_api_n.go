
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ApiN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiN{}

// ApiN struct for ApiN
type ApiN struct {
	D NullableString `json:"d,omitempty"`
	F NullableString `json:"f,omitempty"`
	H NullableInt64 `json:"h,omitempty"`
	I NullableInt64 `json:"i,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiN ApiN

// NewApiN instantiates a new ApiN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiN() *ApiN {
	this := ApiN{}
	return &this
}

// NewApiNWithDefaults instantiates a new ApiN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNWithDefaults() *ApiN {
	this := ApiN{}
	return &this
}

// GetD returns the D field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiN) GetD() string {
	if o == nil || IsNil(o.D.Get()) {
		var ret string
		return ret
	}
	return *o.D.Get()
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiN) GetDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.D.Get(), o.D.IsSet()
}

// HasD returns a boolean if a field has been set.
func (o *ApiN) HasD() bool {
	if o != nil && o.D.IsSet() {
		return true
	}

	return false
}

// SetD gets a reference to the given NullableString and assigns it to the D field.
func (o *ApiN) SetD(v string) {
	o.D.Set(&v)
}
// SetDNil sets the value for D to be an explicit nil
func (o *ApiN) SetDNil() {
	o.D.Set(nil)
}

// UnsetD ensures that no value is present for D, not even an explicit nil
func (o *ApiN) UnsetD() {
	o.D.Unset()
}

// GetF returns the F field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiN) GetF() string {
	if o == nil || IsNil(o.F.Get()) {
		var ret string
		return ret
	}
	return *o.F.Get()
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiN) GetFOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.F.Get(), o.F.IsSet()
}

// HasF returns a boolean if a field has been set.
func (o *ApiN) HasF() bool {
	if o != nil && o.F.IsSet() {
		return true
	}

	return false
}

// SetF gets a reference to the given NullableString and assigns it to the F field.
func (o *ApiN) SetF(v string) {
	o.F.Set(&v)
}
// SetFNil sets the value for F to be an explicit nil
func (o *ApiN) SetFNil() {
	o.F.Set(nil)
}

// UnsetF ensures that no value is present for F, not even an explicit nil
func (o *ApiN) UnsetF() {
	o.F.Unset()
}

// GetH returns the H field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiN) GetH() int64 {
	if o == nil || IsNil(o.H.Get()) {
		var ret int64
		return ret
	}
	return *o.H.Get()
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiN) GetHOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.H.Get(), o.H.IsSet()
}

// HasH returns a boolean if a field has been set.
func (o *ApiN) HasH() bool {
	if o != nil && o.H.IsSet() {
		return true
	}

	return false
}

// SetH gets a reference to the given NullableInt64 and assigns it to the H field.
func (o *ApiN) SetH(v int64) {
	o.H.Set(&v)
}
// SetHNil sets the value for H to be an explicit nil
func (o *ApiN) SetHNil() {
	o.H.Set(nil)
}

// UnsetH ensures that no value is present for H, not even an explicit nil
func (o *ApiN) UnsetH() {
	o.H.Unset()
}

// GetI returns the I field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiN) GetI() int64 {
	if o == nil || IsNil(o.I.Get()) {
		var ret int64
		return ret
	}
	return *o.I.Get()
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiN) GetIOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.I.Get(), o.I.IsSet()
}

// HasI returns a boolean if a field has been set.
func (o *ApiN) HasI() bool {
	if o != nil && o.I.IsSet() {
		return true
	}

	return false
}

// SetI gets a reference to the given NullableInt64 and assigns it to the I field.
func (o *ApiN) SetI(v int64) {
	o.I.Set(&v)
}
// SetINil sets the value for I to be an explicit nil
func (o *ApiN) SetINil() {
	o.I.Set(nil)
}

// UnsetI ensures that no value is present for I, not even an explicit nil
func (o *ApiN) UnsetI() {
	o.I.Unset()
}

func (o ApiN) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.D.IsSet() {
		toSerialize["d"] = o.D.Get()
	}
	if o.F.IsSet() {
		toSerialize["f"] = o.F.Get()
	}
	if o.H.IsSet() {
		toSerialize["h"] = o.H.Get()
	}
	if o.I.IsSet() {
		toSerialize["i"] = o.I.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiN) UnmarshalJSON(data []byte) (err error) {
	varApiN := _ApiN{}

	err = json.Unmarshal(data, &varApiN)

	if err != nil {
		return err
	}

	*o = ApiN(varApiN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "d")
		delete(additionalProperties, "f")
		delete(additionalProperties, "h")
		delete(additionalProperties, "i")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiN struct {
	value *ApiN
	isSet bool
}

func (v NullableApiN) Get() *ApiN {
	return v.value
}

func (v *NullableApiN) Set(val *ApiN) {
	v.value = val
	v.isSet = true
}

func (v NullableApiN) IsSet() bool {
	return v.isSet
}

func (v *NullableApiN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiN(val *ApiN) *NullableApiN {
	return &NullableApiN{value: val, isSet: true}
}

func (v NullableApiN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


