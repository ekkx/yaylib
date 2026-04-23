
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LoginSnsUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginSnsUserRequest{}

// LoginSnsUserRequest struct for LoginSnsUserRequest
type LoginSnsUserRequest struct {
	AccessToken NullableString `json:"access_token,omitempty"`
	AccessTokenSecret NullableString `json:"access_token_secret,omitempty"`
	ApiKey NullableString `json:"api_key,omitempty"`
	AuthorizationCode NullableString `json:"authorization_code,omitempty"`
	ConsumerKey NullableString `json:"consumer_key,omitempty"`
	ConsumerSecret NullableString `json:"consumer_secret,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Provider NullableString `json:"provider,omitempty"`
	TwoFACode NullableString `json:"two_f_a_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoginSnsUserRequest LoginSnsUserRequest

// NewLoginSnsUserRequest instantiates a new LoginSnsUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginSnsUserRequest() *LoginSnsUserRequest {
	this := LoginSnsUserRequest{}
	return &this
}

// NewLoginSnsUserRequestWithDefaults instantiates a new LoginSnsUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginSnsUserRequestWithDefaults() *LoginSnsUserRequest {
	this := LoginSnsUserRequest{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *LoginSnsUserRequest) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *LoginSnsUserRequest) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetAccessTokenSecret returns the AccessTokenSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetAccessTokenSecret() string {
	if o == nil || IsNil(o.AccessTokenSecret.Get()) {
		var ret string
		return ret
	}
	return *o.AccessTokenSecret.Get()
}

// GetAccessTokenSecretOk returns a tuple with the AccessTokenSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetAccessTokenSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessTokenSecret.Get(), o.AccessTokenSecret.IsSet()
}

// HasAccessTokenSecret returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasAccessTokenSecret() bool {
	if o != nil && o.AccessTokenSecret.IsSet() {
		return true
	}

	return false
}

// SetAccessTokenSecret gets a reference to the given NullableString and assigns it to the AccessTokenSecret field.
func (o *LoginSnsUserRequest) SetAccessTokenSecret(v string) {
	o.AccessTokenSecret.Set(&v)
}
// SetAccessTokenSecretNil sets the value for AccessTokenSecret to be an explicit nil
func (o *LoginSnsUserRequest) SetAccessTokenSecretNil() {
	o.AccessTokenSecret.Set(nil)
}

// UnsetAccessTokenSecret ensures that no value is present for AccessTokenSecret, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetAccessTokenSecret() {
	o.AccessTokenSecret.Unset()
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey.Get()) {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasApiKey() bool {
	if o != nil && o.ApiKey.IsSet() {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given NullableString and assigns it to the ApiKey field.
func (o *LoginSnsUserRequest) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}
// SetApiKeyNil sets the value for ApiKey to be an explicit nil
func (o *LoginSnsUserRequest) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetApiKey() {
	o.ApiKey.Unset()
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetAuthorizationCode() string {
	if o == nil || IsNil(o.AuthorizationCode.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode.Get()
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizationCode.Get(), o.AuthorizationCode.IsSet()
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasAuthorizationCode() bool {
	if o != nil && o.AuthorizationCode.IsSet() {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given NullableString and assigns it to the AuthorizationCode field.
func (o *LoginSnsUserRequest) SetAuthorizationCode(v string) {
	o.AuthorizationCode.Set(&v)
}
// SetAuthorizationCodeNil sets the value for AuthorizationCode to be an explicit nil
func (o *LoginSnsUserRequest) SetAuthorizationCodeNil() {
	o.AuthorizationCode.Set(nil)
}

// UnsetAuthorizationCode ensures that no value is present for AuthorizationCode, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetAuthorizationCode() {
	o.AuthorizationCode.Unset()
}

// GetConsumerKey returns the ConsumerKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetConsumerKey() string {
	if o == nil || IsNil(o.ConsumerKey.Get()) {
		var ret string
		return ret
	}
	return *o.ConsumerKey.Get()
}

// GetConsumerKeyOk returns a tuple with the ConsumerKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetConsumerKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsumerKey.Get(), o.ConsumerKey.IsSet()
}

// HasConsumerKey returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasConsumerKey() bool {
	if o != nil && o.ConsumerKey.IsSet() {
		return true
	}

	return false
}

// SetConsumerKey gets a reference to the given NullableString and assigns it to the ConsumerKey field.
func (o *LoginSnsUserRequest) SetConsumerKey(v string) {
	o.ConsumerKey.Set(&v)
}
// SetConsumerKeyNil sets the value for ConsumerKey to be an explicit nil
func (o *LoginSnsUserRequest) SetConsumerKeyNil() {
	o.ConsumerKey.Set(nil)
}

// UnsetConsumerKey ensures that no value is present for ConsumerKey, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetConsumerKey() {
	o.ConsumerKey.Unset()
}

// GetConsumerSecret returns the ConsumerSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetConsumerSecret() string {
	if o == nil || IsNil(o.ConsumerSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ConsumerSecret.Get()
}

// GetConsumerSecretOk returns a tuple with the ConsumerSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetConsumerSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsumerSecret.Get(), o.ConsumerSecret.IsSet()
}

// HasConsumerSecret returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasConsumerSecret() bool {
	if o != nil && o.ConsumerSecret.IsSet() {
		return true
	}

	return false
}

// SetConsumerSecret gets a reference to the given NullableString and assigns it to the ConsumerSecret field.
func (o *LoginSnsUserRequest) SetConsumerSecret(v string) {
	o.ConsumerSecret.Set(&v)
}
// SetConsumerSecretNil sets the value for ConsumerSecret to be an explicit nil
func (o *LoginSnsUserRequest) SetConsumerSecretNil() {
	o.ConsumerSecret.Set(nil)
}

// UnsetConsumerSecret ensures that no value is present for ConsumerSecret, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetConsumerSecret() {
	o.ConsumerSecret.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *LoginSnsUserRequest) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *LoginSnsUserRequest) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetEmail() {
	o.Email.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetProvider() string {
	if o == nil || IsNil(o.Provider.Get()) {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasProvider() bool {
	if o != nil && o.Provider.IsSet() {
		return true
	}

	return false
}

// SetProvider gets a reference to the given NullableString and assigns it to the Provider field.
func (o *LoginSnsUserRequest) SetProvider(v string) {
	o.Provider.Set(&v)
}
// SetProviderNil sets the value for Provider to be an explicit nil
func (o *LoginSnsUserRequest) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetProvider() {
	o.Provider.Unset()
}

// GetTwoFACode returns the TwoFACode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginSnsUserRequest) GetTwoFACode() string {
	if o == nil || IsNil(o.TwoFACode.Get()) {
		var ret string
		return ret
	}
	return *o.TwoFACode.Get()
}

// GetTwoFACodeOk returns a tuple with the TwoFACode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginSnsUserRequest) GetTwoFACodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoFACode.Get(), o.TwoFACode.IsSet()
}

// HasTwoFACode returns a boolean if a field has been set.
func (o *LoginSnsUserRequest) HasTwoFACode() bool {
	if o != nil && o.TwoFACode.IsSet() {
		return true
	}

	return false
}

// SetTwoFACode gets a reference to the given NullableString and assigns it to the TwoFACode field.
func (o *LoginSnsUserRequest) SetTwoFACode(v string) {
	o.TwoFACode.Set(&v)
}
// SetTwoFACodeNil sets the value for TwoFACode to be an explicit nil
func (o *LoginSnsUserRequest) SetTwoFACodeNil() {
	o.TwoFACode.Set(nil)
}

// UnsetTwoFACode ensures that no value is present for TwoFACode, not even an explicit nil
func (o *LoginSnsUserRequest) UnsetTwoFACode() {
	o.TwoFACode.Unset()
}

func (o LoginSnsUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginSnsUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.AccessTokenSecret.IsSet() {
		toSerialize["access_token_secret"] = o.AccessTokenSecret.Get()
	}
	if o.ApiKey.IsSet() {
		toSerialize["api_key"] = o.ApiKey.Get()
	}
	if o.AuthorizationCode.IsSet() {
		toSerialize["authorization_code"] = o.AuthorizationCode.Get()
	}
	if o.ConsumerKey.IsSet() {
		toSerialize["consumer_key"] = o.ConsumerKey.Get()
	}
	if o.ConsumerSecret.IsSet() {
		toSerialize["consumer_secret"] = o.ConsumerSecret.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.TwoFACode.IsSet() {
		toSerialize["two_f_a_code"] = o.TwoFACode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoginSnsUserRequest) UnmarshalJSON(data []byte) (err error) {
	varLoginSnsUserRequest := _LoginSnsUserRequest{}

	err = json.Unmarshal(data, &varLoginSnsUserRequest)

	if err != nil {
		return err
	}

	*o = LoginSnsUserRequest(varLoginSnsUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "access_token_secret")
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "authorization_code")
		delete(additionalProperties, "consumer_key")
		delete(additionalProperties, "consumer_secret")
		delete(additionalProperties, "email")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "two_f_a_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginSnsUserRequest struct {
	value *LoginSnsUserRequest
	isSet bool
}

func (v NullableLoginSnsUserRequest) Get() *LoginSnsUserRequest {
	return v.value
}

func (v *NullableLoginSnsUserRequest) Set(val *LoginSnsUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginSnsUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginSnsUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginSnsUserRequest(val *LoginSnsUserRequest) *NullableLoginSnsUserRequest {
	return &NullableLoginSnsUserRequest{value: val, isSet: true}
}

func (v NullableLoginSnsUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginSnsUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


