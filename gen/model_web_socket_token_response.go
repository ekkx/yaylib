
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the WebSocketTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebSocketTokenResponse{}

// WebSocketTokenResponse struct for WebSocketTokenResponse
type WebSocketTokenResponse struct {
	Token NullableString `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebSocketTokenResponse WebSocketTokenResponse

// NewWebSocketTokenResponse instantiates a new WebSocketTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebSocketTokenResponse() *WebSocketTokenResponse {
	this := WebSocketTokenResponse{}
	return &this
}

// NewWebSocketTokenResponseWithDefaults instantiates a new WebSocketTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebSocketTokenResponseWithDefaults() *WebSocketTokenResponse {
	this := WebSocketTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebSocketTokenResponse) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebSocketTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *WebSocketTokenResponse) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *WebSocketTokenResponse) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *WebSocketTokenResponse) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *WebSocketTokenResponse) UnsetToken() {
	o.Token.Unset()
}

func (o WebSocketTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebSocketTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebSocketTokenResponse) UnmarshalJSON(data []byte) (err error) {
	varWebSocketTokenResponse := _WebSocketTokenResponse{}

	err = json.Unmarshal(data, &varWebSocketTokenResponse)

	if err != nil {
		return err
	}

	*o = WebSocketTokenResponse(varWebSocketTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebSocketTokenResponse struct {
	value *WebSocketTokenResponse
	isSet bool
}

func (v NullableWebSocketTokenResponse) Get() *WebSocketTokenResponse {
	return v.value
}

func (v *NullableWebSocketTokenResponse) Set(val *WebSocketTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebSocketTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebSocketTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebSocketTokenResponse(val *WebSocketTokenResponse) *NullableWebSocketTokenResponse {
	return &NullableWebSocketTokenResponse{value: val, isSet: true}
}

func (v NullableWebSocketTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebSocketTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


