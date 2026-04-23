
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAuth{}

// UserAuth struct for UserAuth
type UserAuth struct {
	AccessToken NullableString `json:"access_token,omitempty"`
	ExpiresIn NullableInt64 `json:"expires_in,omitempty"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserAuth UserAuth

// NewUserAuth instantiates a new UserAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAuth() *UserAuth {
	this := UserAuth{}
	return &this
}

// NewUserAuthWithDefaults instantiates a new UserAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAuthWithDefaults() *UserAuth {
	this := UserAuth{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAuth) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAuth) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *UserAuth) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *UserAuth) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *UserAuth) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *UserAuth) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAuth) GetExpiresIn() int64 {
	if o == nil || IsNil(o.ExpiresIn.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresIn.Get()
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAuth) GetExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresIn.Get(), o.ExpiresIn.IsSet()
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *UserAuth) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn.IsSet() {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given NullableInt64 and assigns it to the ExpiresIn field.
func (o *UserAuth) SetExpiresIn(v int64) {
	o.ExpiresIn.Set(&v)
}
// SetExpiresInNil sets the value for ExpiresIn to be an explicit nil
func (o *UserAuth) SetExpiresInNil() {
	o.ExpiresIn.Set(nil)
}

// UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
func (o *UserAuth) UnsetExpiresIn() {
	o.ExpiresIn.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAuth) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAuth) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *UserAuth) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *UserAuth) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *UserAuth) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *UserAuth) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserAuth) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserAuth) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *UserAuth) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *UserAuth) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *UserAuth) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *UserAuth) UnsetUserId() {
	o.UserId.Unset()
}

func (o UserAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAuth) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserAuth) UnmarshalJSON(data []byte) (err error) {
	varUserAuth := _UserAuth{}

	err = json.Unmarshal(data, &varUserAuth)

	if err != nil {
		return err
	}

	*o = UserAuth(varUserAuth)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "user_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserAuth struct {
	value *UserAuth
	isSet bool
}

func (v NullableUserAuth) Get() *UserAuth {
	return v.value
}

func (v *NullableUserAuth) Set(val *UserAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuth(val *UserAuth) *NullableUserAuth {
	return &NullableUserAuth{value: val, isSet: true}
}

func (v NullableUserAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


