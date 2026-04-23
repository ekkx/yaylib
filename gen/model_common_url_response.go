
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CommonUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonUrlResponse{}

// CommonUrlResponse struct for CommonUrlResponse
type CommonUrlResponse struct {
	Url NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonUrlResponse CommonUrlResponse

// NewCommonUrlResponse instantiates a new CommonUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonUrlResponse() *CommonUrlResponse {
	this := CommonUrlResponse{}
	return &this
}

// NewCommonUrlResponseWithDefaults instantiates a new CommonUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonUrlResponseWithDefaults() *CommonUrlResponse {
	this := CommonUrlResponse{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonUrlResponse) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonUrlResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CommonUrlResponse) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CommonUrlResponse) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *CommonUrlResponse) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CommonUrlResponse) UnsetUrl() {
	o.Url.Unset()
}

func (o CommonUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonUrlResponse) UnmarshalJSON(data []byte) (err error) {
	varCommonUrlResponse := _CommonUrlResponse{}

	err = json.Unmarshal(data, &varCommonUrlResponse)

	if err != nil {
		return err
	}

	*o = CommonUrlResponse(varCommonUrlResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonUrlResponse struct {
	value *CommonUrlResponse
	isSet bool
}

func (v NullableCommonUrlResponse) Get() *CommonUrlResponse {
	return v.value
}

func (v *NullableCommonUrlResponse) Set(val *CommonUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonUrlResponse(val *CommonUrlResponse) *NullableCommonUrlResponse {
	return &NullableCommonUrlResponse{value: val, isSet: true}
}

func (v NullableCommonUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


