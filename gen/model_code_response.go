
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeResponse{}

// CodeResponse struct for CodeResponse
type CodeResponse struct {
	Code NullableString `json:"code,omitempty"`
	NextEditableDatetime NullableInt64 `json:"next_editable_datetime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CodeResponse CodeResponse

// NewCodeResponse instantiates a new CodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeResponse() *CodeResponse {
	this := CodeResponse{}
	return &this
}

// NewCodeResponseWithDefaults instantiates a new CodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeResponseWithDefaults() *CodeResponse {
	this := CodeResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeResponse) GetCode() string {
	if o == nil || IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *CodeResponse) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *CodeResponse) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *CodeResponse) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *CodeResponse) UnsetCode() {
	o.Code.Unset()
}

// GetNextEditableDatetime returns the NextEditableDatetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodeResponse) GetNextEditableDatetime() int64 {
	if o == nil || IsNil(o.NextEditableDatetime.Get()) {
		var ret int64
		return ret
	}
	return *o.NextEditableDatetime.Get()
}

// GetNextEditableDatetimeOk returns a tuple with the NextEditableDatetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodeResponse) GetNextEditableDatetimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextEditableDatetime.Get(), o.NextEditableDatetime.IsSet()
}

// HasNextEditableDatetime returns a boolean if a field has been set.
func (o *CodeResponse) HasNextEditableDatetime() bool {
	if o != nil && o.NextEditableDatetime.IsSet() {
		return true
	}

	return false
}

// SetNextEditableDatetime gets a reference to the given NullableInt64 and assigns it to the NextEditableDatetime field.
func (o *CodeResponse) SetNextEditableDatetime(v int64) {
	o.NextEditableDatetime.Set(&v)
}
// SetNextEditableDatetimeNil sets the value for NextEditableDatetime to be an explicit nil
func (o *CodeResponse) SetNextEditableDatetimeNil() {
	o.NextEditableDatetime.Set(nil)
}

// UnsetNextEditableDatetime ensures that no value is present for NextEditableDatetime, not even an explicit nil
func (o *CodeResponse) UnsetNextEditableDatetime() {
	o.NextEditableDatetime.Unset()
}

func (o CodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.NextEditableDatetime.IsSet() {
		toSerialize["next_editable_datetime"] = o.NextEditableDatetime.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CodeResponse) UnmarshalJSON(data []byte) (err error) {
	varCodeResponse := _CodeResponse{}

	err = json.Unmarshal(data, &varCodeResponse)

	if err != nil {
		return err
	}

	*o = CodeResponse(varCodeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "next_editable_datetime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCodeResponse struct {
	value *CodeResponse
	isSet bool
}

func (v NullableCodeResponse) Get() *CodeResponse {
	return v.value
}

func (v *NullableCodeResponse) Set(val *CodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeResponse(val *CodeResponse) *NullableCodeResponse {
	return &NullableCodeResponse{value: val, isSet: true}
}

func (v NullableCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


