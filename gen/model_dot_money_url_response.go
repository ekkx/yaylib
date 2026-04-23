
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the DotMoneyUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DotMoneyUrlResponse{}

// DotMoneyUrlResponse struct for DotMoneyUrlResponse
type DotMoneyUrlResponse struct {
	Url NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DotMoneyUrlResponse DotMoneyUrlResponse

// NewDotMoneyUrlResponse instantiates a new DotMoneyUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDotMoneyUrlResponse() *DotMoneyUrlResponse {
	this := DotMoneyUrlResponse{}
	return &this
}

// NewDotMoneyUrlResponseWithDefaults instantiates a new DotMoneyUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDotMoneyUrlResponseWithDefaults() *DotMoneyUrlResponse {
	this := DotMoneyUrlResponse{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DotMoneyUrlResponse) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DotMoneyUrlResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *DotMoneyUrlResponse) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *DotMoneyUrlResponse) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *DotMoneyUrlResponse) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *DotMoneyUrlResponse) UnsetUrl() {
	o.Url.Unset()
}

func (o DotMoneyUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DotMoneyUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DotMoneyUrlResponse) UnmarshalJSON(data []byte) (err error) {
	varDotMoneyUrlResponse := _DotMoneyUrlResponse{}

	err = json.Unmarshal(data, &varDotMoneyUrlResponse)

	if err != nil {
		return err
	}

	*o = DotMoneyUrlResponse(varDotMoneyUrlResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDotMoneyUrlResponse struct {
	value *DotMoneyUrlResponse
	isSet bool
}

func (v NullableDotMoneyUrlResponse) Get() *DotMoneyUrlResponse {
	return v.value
}

func (v *NullableDotMoneyUrlResponse) Set(val *DotMoneyUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDotMoneyUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDotMoneyUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDotMoneyUrlResponse(val *DotMoneyUrlResponse) *NullableDotMoneyUrlResponse {
	return &NullableDotMoneyUrlResponse{value: val, isSet: true}
}

func (v NullableDotMoneyUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDotMoneyUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


