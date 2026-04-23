
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LoginUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginUpdateResponse{}

// LoginUpdateResponse struct for LoginUpdateResponse
type LoginUpdateResponse struct {
	AccessToken NullableString `json:"access_token,omitempty"`
	ExpiresIn NullableInt64 `json:"expires_in,omitempty"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	Username NullableString `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoginUpdateResponse LoginUpdateResponse

// NewLoginUpdateResponse instantiates a new LoginUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginUpdateResponse() *LoginUpdateResponse {
	this := LoginUpdateResponse{}
	return &this
}

// NewLoginUpdateResponseWithDefaults instantiates a new LoginUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginUpdateResponseWithDefaults() *LoginUpdateResponse {
	this := LoginUpdateResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUpdateResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUpdateResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *LoginUpdateResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *LoginUpdateResponse) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *LoginUpdateResponse) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *LoginUpdateResponse) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUpdateResponse) GetExpiresIn() int64 {
	if o == nil || IsNil(o.ExpiresIn.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresIn.Get()
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUpdateResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresIn.Get(), o.ExpiresIn.IsSet()
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *LoginUpdateResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn.IsSet() {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given NullableInt64 and assigns it to the ExpiresIn field.
func (o *LoginUpdateResponse) SetExpiresIn(v int64) {
	o.ExpiresIn.Set(&v)
}
// SetExpiresInNil sets the value for ExpiresIn to be an explicit nil
func (o *LoginUpdateResponse) SetExpiresInNil() {
	o.ExpiresIn.Set(nil)
}

// UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
func (o *LoginUpdateResponse) UnsetExpiresIn() {
	o.ExpiresIn.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUpdateResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUpdateResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *LoginUpdateResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *LoginUpdateResponse) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *LoginUpdateResponse) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *LoginUpdateResponse) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUpdateResponse) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUpdateResponse) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *LoginUpdateResponse) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *LoginUpdateResponse) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *LoginUpdateResponse) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *LoginUpdateResponse) UnsetUserId() {
	o.UserId.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUpdateResponse) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUpdateResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *LoginUpdateResponse) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *LoginUpdateResponse) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *LoginUpdateResponse) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *LoginUpdateResponse) UnsetUsername() {
	o.Username.Unset()
}

func (o LoginUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.ExpiresIn.IsSet() {
		toSerialize["expires_in"] = o.ExpiresIn.Get()
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LoginUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	varLoginUpdateResponse := _LoginUpdateResponse{}

	err = json.Unmarshal(data, &varLoginUpdateResponse)

	if err != nil {
		return err
	}

	*o = LoginUpdateResponse(varLoginUpdateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginUpdateResponse struct {
	value *LoginUpdateResponse
	isSet bool
}

func (v NullableLoginUpdateResponse) Get() *LoginUpdateResponse {
	return v.value
}

func (v *NullableLoginUpdateResponse) Set(val *LoginUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginUpdateResponse(val *LoginUpdateResponse) *NullableLoginUpdateResponse {
	return &NullableLoginUpdateResponse{value: val, isSet: true}
}

func (v NullableLoginUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


