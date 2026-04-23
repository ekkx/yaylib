
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PresignedUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresignedUrlResponse{}

// PresignedUrlResponse struct for PresignedUrlResponse
type PresignedUrlResponse struct {
	PresignedUrl NullableString `json:"presigned_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PresignedUrlResponse PresignedUrlResponse

// NewPresignedUrlResponse instantiates a new PresignedUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresignedUrlResponse() *PresignedUrlResponse {
	this := PresignedUrlResponse{}
	return &this
}

// NewPresignedUrlResponseWithDefaults instantiates a new PresignedUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresignedUrlResponseWithDefaults() *PresignedUrlResponse {
	this := PresignedUrlResponse{}
	return &this
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresignedUrlResponse) GetPresignedUrl() string {
	if o == nil || IsNil(o.PresignedUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PresignedUrl.Get()
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresignedUrlResponse) GetPresignedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresignedUrl.Get(), o.PresignedUrl.IsSet()
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *PresignedUrlResponse) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl.IsSet() {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given NullableString and assigns it to the PresignedUrl field.
func (o *PresignedUrlResponse) SetPresignedUrl(v string) {
	o.PresignedUrl.Set(&v)
}
// SetPresignedUrlNil sets the value for PresignedUrl to be an explicit nil
func (o *PresignedUrlResponse) SetPresignedUrlNil() {
	o.PresignedUrl.Set(nil)
}

// UnsetPresignedUrl ensures that no value is present for PresignedUrl, not even an explicit nil
func (o *PresignedUrlResponse) UnsetPresignedUrl() {
	o.PresignedUrl.Unset()
}

func (o PresignedUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresignedUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PresignedUrl.IsSet() {
		toSerialize["presigned_url"] = o.PresignedUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PresignedUrlResponse) UnmarshalJSON(data []byte) (err error) {
	varPresignedUrlResponse := _PresignedUrlResponse{}

	err = json.Unmarshal(data, &varPresignedUrlResponse)

	if err != nil {
		return err
	}

	*o = PresignedUrlResponse(varPresignedUrlResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "presigned_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePresignedUrlResponse struct {
	value *PresignedUrlResponse
	isSet bool
}

func (v NullablePresignedUrlResponse) Get() *PresignedUrlResponse {
	return v.value
}

func (v *NullablePresignedUrlResponse) Set(val *PresignedUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePresignedUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePresignedUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresignedUrlResponse(val *PresignedUrlResponse) *NullablePresignedUrlResponse {
	return &NullablePresignedUrlResponse{value: val, isSet: true}
}

func (v NullablePresignedUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresignedUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


