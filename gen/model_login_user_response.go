
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LoginUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginUserResponse{}

// LoginUserResponse struct for LoginUserResponse
type LoginUserResponse struct {
	AccessToken NullableString `json:"access_token,omitempty"`
	ExpiresIn NullableInt64 `json:"expires_in,omitempty"`
	IsNew NullableBool `json:"is_new,omitempty"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	SnsInfo NullableSnsInfo `json:"sns_info,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	Username NullableString `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LoginUserResponse LoginUserResponse

// NewLoginUserResponse instantiates a new LoginUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginUserResponse() *LoginUserResponse {
	this := LoginUserResponse{}
	return &this
}

// NewLoginUserResponseWithDefaults instantiates a new LoginUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginUserResponseWithDefaults() *LoginUserResponse {
	this := LoginUserResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUserResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUserResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *LoginUserResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *LoginUserResponse) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *LoginUserResponse) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *LoginUserResponse) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUserResponse) GetExpiresIn() int64 {
	if o == nil || IsNil(o.ExpiresIn.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresIn.Get()
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUserResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresIn.Get(), o.ExpiresIn.IsSet()
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *LoginUserResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn.IsSet() {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given NullableInt64 and assigns it to the ExpiresIn field.
func (o *LoginUserResponse) SetExpiresIn(v int64) {
	o.ExpiresIn.Set(&v)
}
// SetExpiresInNil sets the value for ExpiresIn to be an explicit nil
func (o *LoginUserResponse) SetExpiresInNil() {
	o.ExpiresIn.Set(nil)
}

// UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
func (o *LoginUserResponse) UnsetExpiresIn() {
	o.ExpiresIn.Unset()
}

// GetIsNew returns the IsNew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUserResponse) GetIsNew() bool {
	if o == nil || IsNil(o.IsNew.Get()) {
		var ret bool
		return ret
	}
	return *o.IsNew.Get()
}

// GetIsNewOk returns a tuple with the IsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUserResponse) GetIsNewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsNew.Get(), o.IsNew.IsSet()
}

// HasIsNew returns a boolean if a field has been set.
func (o *LoginUserResponse) HasIsNew() bool {
	if o != nil && o.IsNew.IsSet() {
		return true
	}

	return false
}

// SetIsNew gets a reference to the given NullableBool and assigns it to the IsNew field.
func (o *LoginUserResponse) SetIsNew(v bool) {
	o.IsNew.Set(&v)
}
// SetIsNewNil sets the value for IsNew to be an explicit nil
func (o *LoginUserResponse) SetIsNewNil() {
	o.IsNew.Set(nil)
}

// UnsetIsNew ensures that no value is present for IsNew, not even an explicit nil
func (o *LoginUserResponse) UnsetIsNew() {
	o.IsNew.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUserResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUserResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *LoginUserResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *LoginUserResponse) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *LoginUserResponse) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *LoginUserResponse) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetSnsInfo returns the SnsInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUserResponse) GetSnsInfo() SnsInfo {
	if o == nil || IsNil(o.SnsInfo.Get()) {
		var ret SnsInfo
		return ret
	}
	return *o.SnsInfo.Get()
}

// GetSnsInfoOk returns a tuple with the SnsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUserResponse) GetSnsInfoOk() (*SnsInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnsInfo.Get(), o.SnsInfo.IsSet()
}

// HasSnsInfo returns a boolean if a field has been set.
func (o *LoginUserResponse) HasSnsInfo() bool {
	if o != nil && o.SnsInfo.IsSet() {
		return true
	}

	return false
}

// SetSnsInfo gets a reference to the given NullableSnsInfo and assigns it to the SnsInfo field.
func (o *LoginUserResponse) SetSnsInfo(v SnsInfo) {
	o.SnsInfo.Set(&v)
}
// SetSnsInfoNil sets the value for SnsInfo to be an explicit nil
func (o *LoginUserResponse) SetSnsInfoNil() {
	o.SnsInfo.Set(nil)
}

// UnsetSnsInfo ensures that no value is present for SnsInfo, not even an explicit nil
func (o *LoginUserResponse) UnsetSnsInfo() {
	o.SnsInfo.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUserResponse) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUserResponse) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *LoginUserResponse) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *LoginUserResponse) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *LoginUserResponse) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *LoginUserResponse) UnsetUserId() {
	o.UserId.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginUserResponse) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginUserResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *LoginUserResponse) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *LoginUserResponse) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *LoginUserResponse) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *LoginUserResponse) UnsetUsername() {
	o.Username.Unset()
}

func (o LoginUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.ExpiresIn.IsSet() {
		toSerialize["expires_in"] = o.ExpiresIn.Get()
	}
	if o.IsNew.IsSet() {
		toSerialize["is_new"] = o.IsNew.Get()
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}
	if o.SnsInfo.IsSet() {
		toSerialize["sns_info"] = o.SnsInfo.Get()
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

func (o *LoginUserResponse) UnmarshalJSON(data []byte) (err error) {
	varLoginUserResponse := _LoginUserResponse{}

	err = json.Unmarshal(data, &varLoginUserResponse)

	if err != nil {
		return err
	}

	*o = LoginUserResponse(varLoginUserResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "is_new")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "sns_info")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoginUserResponse struct {
	value *LoginUserResponse
	isSet bool
}

func (v NullableLoginUserResponse) Get() *LoginUserResponse {
	return v.value
}

func (v *NullableLoginUserResponse) Set(val *LoginUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginUserResponse(val *LoginUserResponse) *NullableLoginUserResponse {
	return &NullableLoginUserResponse{value: val, isSet: true}
}

func (v NullableLoginUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


