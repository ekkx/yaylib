
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the IdCheckerPresignedUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdCheckerPresignedUrlResponse{}

// IdCheckerPresignedUrlResponse struct for IdCheckerPresignedUrlResponse
type IdCheckerPresignedUrlResponse struct {
	PresignedUrl NullableString `json:"presigned_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdCheckerPresignedUrlResponse IdCheckerPresignedUrlResponse

// NewIdCheckerPresignedUrlResponse instantiates a new IdCheckerPresignedUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdCheckerPresignedUrlResponse() *IdCheckerPresignedUrlResponse {
	this := IdCheckerPresignedUrlResponse{}
	return &this
}

// NewIdCheckerPresignedUrlResponseWithDefaults instantiates a new IdCheckerPresignedUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdCheckerPresignedUrlResponseWithDefaults() *IdCheckerPresignedUrlResponse {
	this := IdCheckerPresignedUrlResponse{}
	return &this
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdCheckerPresignedUrlResponse) GetPresignedUrl() string {
	if o == nil || IsNil(o.PresignedUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PresignedUrl.Get()
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdCheckerPresignedUrlResponse) GetPresignedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresignedUrl.Get(), o.PresignedUrl.IsSet()
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *IdCheckerPresignedUrlResponse) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl.IsSet() {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given NullableString and assigns it to the PresignedUrl field.
func (o *IdCheckerPresignedUrlResponse) SetPresignedUrl(v string) {
	o.PresignedUrl.Set(&v)
}
// SetPresignedUrlNil sets the value for PresignedUrl to be an explicit nil
func (o *IdCheckerPresignedUrlResponse) SetPresignedUrlNil() {
	o.PresignedUrl.Set(nil)
}

// UnsetPresignedUrl ensures that no value is present for PresignedUrl, not even an explicit nil
func (o *IdCheckerPresignedUrlResponse) UnsetPresignedUrl() {
	o.PresignedUrl.Unset()
}

func (o IdCheckerPresignedUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdCheckerPresignedUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PresignedUrl.IsSet() {
		toSerialize["presigned_url"] = o.PresignedUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdCheckerPresignedUrlResponse) UnmarshalJSON(data []byte) (err error) {
	varIdCheckerPresignedUrlResponse := _IdCheckerPresignedUrlResponse{}

	err = json.Unmarshal(data, &varIdCheckerPresignedUrlResponse)

	if err != nil {
		return err
	}

	*o = IdCheckerPresignedUrlResponse(varIdCheckerPresignedUrlResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "presigned_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdCheckerPresignedUrlResponse struct {
	value *IdCheckerPresignedUrlResponse
	isSet bool
}

func (v NullableIdCheckerPresignedUrlResponse) Get() *IdCheckerPresignedUrlResponse {
	return v.value
}

func (v *NullableIdCheckerPresignedUrlResponse) Set(val *IdCheckerPresignedUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdCheckerPresignedUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdCheckerPresignedUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdCheckerPresignedUrlResponse(val *IdCheckerPresignedUrlResponse) *NullableIdCheckerPresignedUrlResponse {
	return &NullableIdCheckerPresignedUrlResponse{value: val, isSet: true}
}

func (v NullableIdCheckerPresignedUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdCheckerPresignedUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


