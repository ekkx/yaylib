
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TotalChatRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TotalChatRequestResponse{}

// TotalChatRequestResponse struct for TotalChatRequestResponse
type TotalChatRequestResponse struct {
	Total NullableInt32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TotalChatRequestResponse TotalChatRequestResponse

// NewTotalChatRequestResponse instantiates a new TotalChatRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTotalChatRequestResponse() *TotalChatRequestResponse {
	this := TotalChatRequestResponse{}
	return &this
}

// NewTotalChatRequestResponseWithDefaults instantiates a new TotalChatRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTotalChatRequestResponseWithDefaults() *TotalChatRequestResponse {
	this := TotalChatRequestResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TotalChatRequestResponse) GetTotal() int32 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret int32
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TotalChatRequestResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *TotalChatRequestResponse) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableInt32 and assigns it to the Total field.
func (o *TotalChatRequestResponse) SetTotal(v int32) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *TotalChatRequestResponse) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *TotalChatRequestResponse) UnsetTotal() {
	o.Total.Unset()
}

func (o TotalChatRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TotalChatRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TotalChatRequestResponse) UnmarshalJSON(data []byte) (err error) {
	varTotalChatRequestResponse := _TotalChatRequestResponse{}

	err = json.Unmarshal(data, &varTotalChatRequestResponse)

	if err != nil {
		return err
	}

	*o = TotalChatRequestResponse(varTotalChatRequestResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTotalChatRequestResponse struct {
	value *TotalChatRequestResponse
	isSet bool
}

func (v NullableTotalChatRequestResponse) Get() *TotalChatRequestResponse {
	return v.value
}

func (v *NullableTotalChatRequestResponse) Set(val *TotalChatRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTotalChatRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTotalChatRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTotalChatRequestResponse(val *TotalChatRequestResponse) *NullableTotalChatRequestResponse {
	return &NullableTotalChatRequestResponse{value: val, isSet: true}
}

func (v NullableTotalChatRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTotalChatRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


