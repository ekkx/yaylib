
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Callback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Callback{}

// Callback struct for Callback
type Callback struct {
	CallbackMethod map[string]interface{} `json:"callback_method,omitempty"`
	Id NullableInt32 `json:"id,omitempty"`
	ResponseString NullableString `json:"response_string,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Callback Callback

// NewCallback instantiates a new Callback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallback() *Callback {
	this := Callback{}
	return &this
}

// NewCallbackWithDefaults instantiates a new Callback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallbackWithDefaults() *Callback {
	this := Callback{}
	return &this
}

// GetCallbackMethod returns the CallbackMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Callback) GetCallbackMethod() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.CallbackMethod
}

// GetCallbackMethodOk returns a tuple with the CallbackMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Callback) GetCallbackMethodOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CallbackMethod) {
		return map[string]interface{}{}, false
	}
	return o.CallbackMethod, true
}

// HasCallbackMethod returns a boolean if a field has been set.
func (o *Callback) HasCallbackMethod() bool {
	if o != nil && !IsNil(o.CallbackMethod) {
		return true
	}

	return false
}

// SetCallbackMethod gets a reference to the given map[string]interface{} and assigns it to the CallbackMethod field.
func (o *Callback) SetCallbackMethod(v map[string]interface{}) {
	o.CallbackMethod = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Callback) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Callback) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Callback) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Callback) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Callback) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Callback) UnsetId() {
	o.Id.Unset()
}

// GetResponseString returns the ResponseString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Callback) GetResponseString() string {
	if o == nil || IsNil(o.ResponseString.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseString.Get()
}

// GetResponseStringOk returns a tuple with the ResponseString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Callback) GetResponseStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseString.Get(), o.ResponseString.IsSet()
}

// HasResponseString returns a boolean if a field has been set.
func (o *Callback) HasResponseString() bool {
	if o != nil && o.ResponseString.IsSet() {
		return true
	}

	return false
}

// SetResponseString gets a reference to the given NullableString and assigns it to the ResponseString field.
func (o *Callback) SetResponseString(v string) {
	o.ResponseString.Set(&v)
}
// SetResponseStringNil sets the value for ResponseString to be an explicit nil
func (o *Callback) SetResponseStringNil() {
	o.ResponseString.Set(nil)
}

// UnsetResponseString ensures that no value is present for ResponseString, not even an explicit nil
func (o *Callback) UnsetResponseString() {
	o.ResponseString.Unset()
}

func (o Callback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Callback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackMethod != nil {
		toSerialize["callback_method"] = o.CallbackMethod
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ResponseString.IsSet() {
		toSerialize["response_string"] = o.ResponseString.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Callback) UnmarshalJSON(data []byte) (err error) {
	varCallback := _Callback{}

	err = json.Unmarshal(data, &varCallback)

	if err != nil {
		return err
	}

	*o = Callback(varCallback)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "callback_method")
		delete(additionalProperties, "id")
		delete(additionalProperties, "response_string")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCallback struct {
	value *Callback
	isSet bool
}

func (v NullableCallback) Get() *Callback {
	return v.value
}

func (v *NullableCallback) Set(val *Callback) {
	v.value = val
	v.isSet = true
}

func (v NullableCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallback(val *Callback) *NullableCallback {
	return &NullableCallback{value: val, isSet: true}
}

func (v NullableCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


