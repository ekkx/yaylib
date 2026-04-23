
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the WalletPresignedUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletPresignedUrlResponse{}

// WalletPresignedUrlResponse struct for WalletPresignedUrlResponse
type WalletPresignedUrlResponse struct {
	PresignedUrl NullableString `json:"presigned_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WalletPresignedUrlResponse WalletPresignedUrlResponse

// NewWalletPresignedUrlResponse instantiates a new WalletPresignedUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletPresignedUrlResponse() *WalletPresignedUrlResponse {
	this := WalletPresignedUrlResponse{}
	return &this
}

// NewWalletPresignedUrlResponseWithDefaults instantiates a new WalletPresignedUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletPresignedUrlResponseWithDefaults() *WalletPresignedUrlResponse {
	this := WalletPresignedUrlResponse{}
	return &this
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletPresignedUrlResponse) GetPresignedUrl() string {
	if o == nil || IsNil(o.PresignedUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PresignedUrl.Get()
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletPresignedUrlResponse) GetPresignedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PresignedUrl.Get(), o.PresignedUrl.IsSet()
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *WalletPresignedUrlResponse) HasPresignedUrl() bool {
	if o != nil && o.PresignedUrl.IsSet() {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given NullableString and assigns it to the PresignedUrl field.
func (o *WalletPresignedUrlResponse) SetPresignedUrl(v string) {
	o.PresignedUrl.Set(&v)
}
// SetPresignedUrlNil sets the value for PresignedUrl to be an explicit nil
func (o *WalletPresignedUrlResponse) SetPresignedUrlNil() {
	o.PresignedUrl.Set(nil)
}

// UnsetPresignedUrl ensures that no value is present for PresignedUrl, not even an explicit nil
func (o *WalletPresignedUrlResponse) UnsetPresignedUrl() {
	o.PresignedUrl.Unset()
}

func (o WalletPresignedUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletPresignedUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PresignedUrl.IsSet() {
		toSerialize["presigned_url"] = o.PresignedUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WalletPresignedUrlResponse) UnmarshalJSON(data []byte) (err error) {
	varWalletPresignedUrlResponse := _WalletPresignedUrlResponse{}

	err = json.Unmarshal(data, &varWalletPresignedUrlResponse)

	if err != nil {
		return err
	}

	*o = WalletPresignedUrlResponse(varWalletPresignedUrlResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "presigned_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalletPresignedUrlResponse struct {
	value *WalletPresignedUrlResponse
	isSet bool
}

func (v NullableWalletPresignedUrlResponse) Get() *WalletPresignedUrlResponse {
	return v.value
}

func (v *NullableWalletPresignedUrlResponse) Set(val *WalletPresignedUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletPresignedUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletPresignedUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletPresignedUrlResponse(val *WalletPresignedUrlResponse) *NullableWalletPresignedUrlResponse {
	return &NullableWalletPresignedUrlResponse{value: val, isSet: true}
}

func (v NullableWalletPresignedUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletPresignedUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


