
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Request{}

// Request struct for Request
type Request struct {
	Method NullableRequestMethod `json:"method,omitempty"`
	// Unresolved Java type: org.json.JSONArray
	Params map[string]interface{} `json:"params,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Request Request

// NewRequest instantiates a new Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequest() *Request {
	this := Request{}
	return &this
}

// NewRequestWithDefaults instantiates a new Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestWithDefaults() *Request {
	this := Request{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Request) GetMethod() RequestMethod {
	if o == nil || IsNil(o.Method.Get()) {
		var ret RequestMethod
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Request) GetMethodOk() (*RequestMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field has been set.
func (o *Request) HasMethod() bool {
	if o != nil && o.Method.IsSet() {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableRequestMethod and assigns it to the Method field.
func (o *Request) SetMethod(v RequestMethod) {
	o.Method.Set(&v)
}
// SetMethodNil sets the value for Method to be an explicit nil
func (o *Request) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *Request) UnsetMethod() {
	o.Method.Unset()
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Request) GetParams() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Request) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *Request) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *Request) SetParams(v map[string]interface{}) {
	o.Params = v
}

func (o Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Method.IsSet() {
		toSerialize["method"] = o.Method.Get()
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Request) UnmarshalJSON(data []byte) (err error) {
	varRequest := _Request{}

	err = json.Unmarshal(data, &varRequest)

	if err != nil {
		return err
	}

	*o = Request(varRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequest struct {
	value *Request
	isSet bool
}

func (v NullableRequest) Get() *Request {
	return v.value
}

func (v *NullableRequest) Set(val *Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequest(val *Request) *NullableRequest {
	return &NullableRequest{value: val, isSet: true}
}

func (v NullableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


