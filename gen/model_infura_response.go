
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the InfuraResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfuraResponse{}

// InfuraResponse struct for InfuraResponse
type InfuraResponse struct {
	Id NullableInt64 `json:"id,omitempty"`
	Result interface{} `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfuraResponse InfuraResponse

// NewInfuraResponse instantiates a new InfuraResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfuraResponse() *InfuraResponse {
	this := InfuraResponse{}
	return &this
}

// NewInfuraResponseWithDefaults instantiates a new InfuraResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfuraResponseWithDefaults() *InfuraResponse {
	this := InfuraResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfuraResponse) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfuraResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InfuraResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *InfuraResponse) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InfuraResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InfuraResponse) UnsetId() {
	o.Id.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InfuraResponse) GetResult() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InfuraResponse) GetResultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return &o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InfuraResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given interface{} and assigns it to the Result field.
func (o *InfuraResponse) SetResult(v interface{}) {
	o.Result = v
}

func (o InfuraResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfuraResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InfuraResponse) UnmarshalJSON(data []byte) (err error) {
	varInfuraResponse := _InfuraResponse{}

	err = json.Unmarshal(data, &varInfuraResponse)

	if err != nil {
		return err
	}

	*o = InfuraResponse(varInfuraResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfuraResponse struct {
	value *InfuraResponse
	isSet bool
}

func (v NullableInfuraResponse) Get() *InfuraResponse {
	return v.value
}

func (v *NullableInfuraResponse) Set(val *InfuraResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInfuraResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInfuraResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfuraResponse(val *InfuraResponse) *NullableInfuraResponse {
	return &NullableInfuraResponse{value: val, isSet: true}
}

func (v NullableInfuraResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfuraResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


