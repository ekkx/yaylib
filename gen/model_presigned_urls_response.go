
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PresignedUrlsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PresignedUrlsResponse{}

// PresignedUrlsResponse struct for PresignedUrlsResponse
type PresignedUrlsResponse struct {
	PresignedUrls []PresignedUrl `json:"presigned_urls,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PresignedUrlsResponse PresignedUrlsResponse

// NewPresignedUrlsResponse instantiates a new PresignedUrlsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresignedUrlsResponse() *PresignedUrlsResponse {
	this := PresignedUrlsResponse{}
	return &this
}

// NewPresignedUrlsResponseWithDefaults instantiates a new PresignedUrlsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresignedUrlsResponseWithDefaults() *PresignedUrlsResponse {
	this := PresignedUrlsResponse{}
	return &this
}

// GetPresignedUrls returns the PresignedUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PresignedUrlsResponse) GetPresignedUrls() []PresignedUrl {
	if o == nil {
		var ret []PresignedUrl
		return ret
	}
	return o.PresignedUrls
}

// GetPresignedUrlsOk returns a tuple with the PresignedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PresignedUrlsResponse) GetPresignedUrlsOk() ([]PresignedUrl, bool) {
	if o == nil || IsNil(o.PresignedUrls) {
		return nil, false
	}
	return o.PresignedUrls, true
}

// HasPresignedUrls returns a boolean if a field has been set.
func (o *PresignedUrlsResponse) HasPresignedUrls() bool {
	if o != nil && !IsNil(o.PresignedUrls) {
		return true
	}

	return false
}

// SetPresignedUrls gets a reference to the given []PresignedUrl and assigns it to the PresignedUrls field.
func (o *PresignedUrlsResponse) SetPresignedUrls(v []PresignedUrl) {
	o.PresignedUrls = v
}

func (o PresignedUrlsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PresignedUrlsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PresignedUrls != nil {
		toSerialize["presigned_urls"] = o.PresignedUrls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PresignedUrlsResponse) UnmarshalJSON(data []byte) (err error) {
	varPresignedUrlsResponse := _PresignedUrlsResponse{}

	err = json.Unmarshal(data, &varPresignedUrlsResponse)

	if err != nil {
		return err
	}

	*o = PresignedUrlsResponse(varPresignedUrlsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "presigned_urls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePresignedUrlsResponse struct {
	value *PresignedUrlsResponse
	isSet bool
}

func (v NullablePresignedUrlsResponse) Get() *PresignedUrlsResponse {
	return v.value
}

func (v *NullablePresignedUrlsResponse) Set(val *PresignedUrlsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePresignedUrlsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePresignedUrlsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresignedUrlsResponse(val *PresignedUrlsResponse) *NullablePresignedUrlsResponse {
	return &NullablePresignedUrlsResponse{value: val, isSet: true}
}

func (v NullablePresignedUrlsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresignedUrlsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


