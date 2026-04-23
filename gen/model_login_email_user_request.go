
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LoginEmailUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginEmailUserRequest{}

// LoginEmailUserRequest struct for LoginEmailUserRequest
type LoginEmailUserRequest struct {
	ApiKey NullableString `json:"api_key,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Password NullableString `json:"password,omitempty"`
	TwoFACode NullableString `json:"two_f_a_code,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoginEmailUserRequest LoginEmailUserRequest

// NewLoginEmailUserRequest instantiates a new LoginEmailUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginEmailUserRequest() *LoginEmailUserRequest {
	this := LoginEmailUserRequest{}
	return &this
}

// NewLoginEmailUserRequestWithDefaults instantiates a new LoginEmailUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginEmailUserRequestWithDefaults() *LoginEmailUserRequest {
	this := LoginEmailUserRequest{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEmailUserRequest) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey.Get()) {
		var ret string
		return ret
	}
	return *o.ApiKey.Get()
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEmailUserRequest) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiKey.Get(), o.ApiKey.IsSet()
}

// HasApiKey returns a boolean if a field has been set.
func (o *LoginEmailUserRequest) HasApiKey() bool {
	if o != nil && o.ApiKey.IsSet() {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given NullableString and assigns it to the ApiKey field.
func (o *LoginEmailUserRequest) SetApiKey(v string) {
	o.ApiKey.Set(&v)
}
// SetApiKeyNil sets the value for ApiKey to be an explicit nil
func (o *LoginEmailUserRequest) SetApiKeyNil() {
	o.ApiKey.Set(nil)
}

// UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
func (o *LoginEmailUserRequest) UnsetApiKey() {
	o.ApiKey.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEmailUserRequest) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEmailUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *LoginEmailUserRequest) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *LoginEmailUserRequest) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *LoginEmailUserRequest) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *LoginEmailUserRequest) UnsetEmail() {
	o.Email.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEmailUserRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEmailUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *LoginEmailUserRequest) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *LoginEmailUserRequest) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *LoginEmailUserRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *LoginEmailUserRequest) UnsetPassword() {
	o.Password.Unset()
}

// GetTwoFACode returns the TwoFACode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEmailUserRequest) GetTwoFACode() string {
	if o == nil || IsNil(o.TwoFACode.Get()) {
		var ret string
		return ret
	}
	return *o.TwoFACode.Get()
}

// GetTwoFACodeOk returns a tuple with the TwoFACode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEmailUserRequest) GetTwoFACodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoFACode.Get(), o.TwoFACode.IsSet()
}

// HasTwoFACode returns a boolean if a field has been set.
func (o *LoginEmailUserRequest) HasTwoFACode() bool {
	if o != nil && o.TwoFACode.IsSet() {
		return true
	}

	return false
}

// SetTwoFACode gets a reference to the given NullableString and assigns it to the TwoFACode field.
func (o *LoginEmailUserRequest) SetTwoFACode(v string) {
	o.TwoFACode.Set(&v)
}
// SetTwoFACodeNil sets the value for TwoFACode to be an explicit nil
func (o *LoginEmailUserRequest) SetTwoFACodeNil() {
	o.TwoFACode.Set(nil)
}

// UnsetTwoFACode ensures that no value is present for TwoFACode, not even an explicit nil
func (o *LoginEmailUserRequest) UnsetTwoFACode() {
	o.TwoFACode.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEmailUserRequest) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEmailUserRequest) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *LoginEmailUserRequest) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *LoginEmailUserRequest) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *LoginEmailUserRequest) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *LoginEmailUserRequest) UnsetUuid() {
	o.Uuid.Unset()
}

func (o LoginEmailUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginEmailUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiKey.IsSet() {
		toSerialize["api_key"] = o.ApiKey.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.TwoFACode.IsSet() {
		toSerialize["two_f_a_code"] = o.TwoFACode.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoginEmailUserRequest) UnmarshalJSON(data []byte) (err error) {
	varLoginEmailUserRequest := _LoginEmailUserRequest{}

	err = json.Unmarshal(data, &varLoginEmailUserRequest)

	if err != nil {
		return err
	}

	*o = LoginEmailUserRequest(varLoginEmailUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_key")
		delete(additionalProperties, "email")
		delete(additionalProperties, "password")
		delete(additionalProperties, "two_f_a_code")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginEmailUserRequest struct {
	value *LoginEmailUserRequest
	isSet bool
}

func (v NullableLoginEmailUserRequest) Get() *LoginEmailUserRequest {
	return v.value
}

func (v *NullableLoginEmailUserRequest) Set(val *LoginEmailUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginEmailUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginEmailUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginEmailUserRequest(val *LoginEmailUserRequest) *NullableLoginEmailUserRequest {
	return &NullableLoginEmailUserRequest{value: val, isSet: true}
}

func (v NullableLoginEmailUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginEmailUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


