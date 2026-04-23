
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the InfuraRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfuraRequest{}

// InfuraRequest struct for InfuraRequest
type InfuraRequest struct {
	Id NullableInt64 `json:"id,omitempty"`
	Jsonrpc NullableString `json:"jsonrpc,omitempty"`
	Method NullableString `json:"method,omitempty"`
	Params []interface{} `json:"params,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfuraRequest InfuraRequest

// NewInfuraRequest instantiates a new InfuraRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfuraRequest() *InfuraRequest {
	this := InfuraRequest{}
	return &this
}

// NewInfuraRequestWithDefaults instantiates a new InfuraRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfuraRequestWithDefaults() *InfuraRequest {
	this := InfuraRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfuraRequest) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfuraRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InfuraRequest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *InfuraRequest) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InfuraRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InfuraRequest) UnsetId() {
	o.Id.Unset()
}

// GetJsonrpc returns the Jsonrpc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfuraRequest) GetJsonrpc() string {
	if o == nil || IsNil(o.Jsonrpc.Get()) {
		var ret string
		return ret
	}
	return *o.Jsonrpc.Get()
}

// GetJsonrpcOk returns a tuple with the Jsonrpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfuraRequest) GetJsonrpcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jsonrpc.Get(), o.Jsonrpc.IsSet()
}

// HasJsonrpc returns a boolean if a field has been set.
func (o *InfuraRequest) HasJsonrpc() bool {
	if o != nil && o.Jsonrpc.IsSet() {
		return true
	}

	return false
}

// SetJsonrpc gets a reference to the given NullableString and assigns it to the Jsonrpc field.
func (o *InfuraRequest) SetJsonrpc(v string) {
	o.Jsonrpc.Set(&v)
}
// SetJsonrpcNil sets the value for Jsonrpc to be an explicit nil
func (o *InfuraRequest) SetJsonrpcNil() {
	o.Jsonrpc.Set(nil)
}

// UnsetJsonrpc ensures that no value is present for Jsonrpc, not even an explicit nil
func (o *InfuraRequest) UnsetJsonrpc() {
	o.Jsonrpc.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfuraRequest) GetMethod() string {
	if o == nil || IsNil(o.Method.Get()) {
		var ret string
		return ret
	}
	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfuraRequest) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// HasMethod returns a boolean if a field has been set.
func (o *InfuraRequest) HasMethod() bool {
	if o != nil && o.Method.IsSet() {
		return true
	}

	return false
}

// SetMethod gets a reference to the given NullableString and assigns it to the Method field.
func (o *InfuraRequest) SetMethod(v string) {
	o.Method.Set(&v)
}
// SetMethodNil sets the value for Method to be an explicit nil
func (o *InfuraRequest) SetMethodNil() {
	o.Method.Set(nil)
}

// UnsetMethod ensures that no value is present for Method, not even an explicit nil
func (o *InfuraRequest) UnsetMethod() {
	o.Method.Unset()
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfuraRequest) GetParams() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfuraRequest) GetParamsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *InfuraRequest) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []interface{} and assigns it to the Params field.
func (o *InfuraRequest) SetParams(v []interface{}) {
	o.Params = v
}

func (o InfuraRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfuraRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Jsonrpc.IsSet() {
		toSerialize["jsonrpc"] = o.Jsonrpc.Get()
	}
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

func (o *InfuraRequest) UnmarshalJSON(data []byte) (err error) {
	varInfuraRequest := _InfuraRequest{}

	err = json.Unmarshal(data, &varInfuraRequest)

	if err != nil {
		return err
	}

	*o = InfuraRequest(varInfuraRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "jsonrpc")
		delete(additionalProperties, "method")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfuraRequest struct {
	value *InfuraRequest
	isSet bool
}

func (v NullableInfuraRequest) Get() *InfuraRequest {
	return v.value
}

func (v *NullableInfuraRequest) Set(val *InfuraRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInfuraRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInfuraRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfuraRequest(val *InfuraRequest) *NullableInfuraRequest {
	return &NullableInfuraRequest{value: val, isSet: true}
}

func (v NullableInfuraRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfuraRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


