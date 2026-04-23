
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RtmTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RtmTokenResponse{}

// RtmTokenResponse struct for RtmTokenResponse
type RtmTokenResponse struct {
	Token NullableString `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RtmTokenResponse RtmTokenResponse

// NewRtmTokenResponse instantiates a new RtmTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRtmTokenResponse() *RtmTokenResponse {
	this := RtmTokenResponse{}
	return &this
}

// NewRtmTokenResponseWithDefaults instantiates a new RtmTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRtmTokenResponseWithDefaults() *RtmTokenResponse {
	this := RtmTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RtmTokenResponse) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RtmTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *RtmTokenResponse) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *RtmTokenResponse) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *RtmTokenResponse) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *RtmTokenResponse) UnsetToken() {
	o.Token.Unset()
}

func (o RtmTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RtmTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RtmTokenResponse) UnmarshalJSON(data []byte) (err error) {
	varRtmTokenResponse := _RtmTokenResponse{}

	err = json.Unmarshal(data, &varRtmTokenResponse)

	if err != nil {
		return err
	}

	*o = RtmTokenResponse(varRtmTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRtmTokenResponse struct {
	value *RtmTokenResponse
	isSet bool
}

func (v NullableRtmTokenResponse) Get() *RtmTokenResponse {
	return v.value
}

func (v *NullableRtmTokenResponse) Set(val *RtmTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRtmTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRtmTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRtmTokenResponse(val *RtmTokenResponse) *NullableRtmTokenResponse {
	return &NullableRtmTokenResponse{value: val, isSet: true}
}

func (v NullableRtmTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRtmTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


