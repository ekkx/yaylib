
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the WebSocketInteractor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebSocketInteractor{}

// WebSocketInteractor struct for WebSocketInteractor
type WebSocketInteractor struct {
	// Unresolved Java type: oi.c
	D map[string]interface{} `json:"d,omitempty"`
	// Unresolved Java type: jp.nanameue.yay.api.WebSocketInteractor.a
	F map[string]interface{} `json:"f,omitempty"`
	// Unresolved Java type: jp.nanameue.yay.api.WebSocketInteractor.c
	H map[string]interface{} `json:"h,omitempty"`
	// Unresolved Java type: ei.z
	I map[string]interface{} `json:"i,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebSocketInteractor WebSocketInteractor

// NewWebSocketInteractor instantiates a new WebSocketInteractor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebSocketInteractor() *WebSocketInteractor {
	this := WebSocketInteractor{}
	return &this
}

// NewWebSocketInteractorWithDefaults instantiates a new WebSocketInteractor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebSocketInteractorWithDefaults() *WebSocketInteractor {
	this := WebSocketInteractor{}
	return &this
}

// GetD returns the D field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebSocketInteractor) GetD() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebSocketInteractor) GetDOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.D) {
		return map[string]interface{}{}, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *WebSocketInteractor) HasD() bool {
	if o != nil && !IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given map[string]interface{} and assigns it to the D field.
func (o *WebSocketInteractor) SetD(v map[string]interface{}) {
	o.D = v
}

// GetF returns the F field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebSocketInteractor) GetF() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebSocketInteractor) GetFOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.F) {
		return map[string]interface{}{}, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *WebSocketInteractor) HasF() bool {
	if o != nil && !IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given map[string]interface{} and assigns it to the F field.
func (o *WebSocketInteractor) SetF(v map[string]interface{}) {
	o.F = v
}

// GetH returns the H field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebSocketInteractor) GetH() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebSocketInteractor) GetHOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.H) {
		return map[string]interface{}{}, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *WebSocketInteractor) HasH() bool {
	if o != nil && !IsNil(o.H) {
		return true
	}

	return false
}

// SetH gets a reference to the given map[string]interface{} and assigns it to the H field.
func (o *WebSocketInteractor) SetH(v map[string]interface{}) {
	o.H = v
}

// GetI returns the I field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebSocketInteractor) GetI() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebSocketInteractor) GetIOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.I) {
		return map[string]interface{}{}, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *WebSocketInteractor) HasI() bool {
	if o != nil && !IsNil(o.I) {
		return true
	}

	return false
}

// SetI gets a reference to the given map[string]interface{} and assigns it to the I field.
func (o *WebSocketInteractor) SetI(v map[string]interface{}) {
	o.I = v
}

func (o WebSocketInteractor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebSocketInteractor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.D != nil {
		toSerialize["d"] = o.D
	}
	if o.F != nil {
		toSerialize["f"] = o.F
	}
	if o.H != nil {
		toSerialize["h"] = o.H
	}
	if o.I != nil {
		toSerialize["i"] = o.I
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebSocketInteractor) UnmarshalJSON(data []byte) (err error) {
	varWebSocketInteractor := _WebSocketInteractor{}

	err = json.Unmarshal(data, &varWebSocketInteractor)

	if err != nil {
		return err
	}

	*o = WebSocketInteractor(varWebSocketInteractor)

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

type NullableWebSocketInteractor struct {
	value *WebSocketInteractor
	isSet bool
}

func (v NullableWebSocketInteractor) Get() *WebSocketInteractor {
	return v.value
}

func (v *NullableWebSocketInteractor) Set(val *WebSocketInteractor) {
	v.value = val
	v.isSet = true
}

func (v NullableWebSocketInteractor) IsSet() bool {
	return v.isSet
}

func (v *NullableWebSocketInteractor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebSocketInteractor(val *WebSocketInteractor) *NullableWebSocketInteractor {
	return &NullableWebSocketInteractor{value: val, isSet: true}
}

func (v NullableWebSocketInteractor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebSocketInteractor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


