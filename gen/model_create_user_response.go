
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserResponse{}

// CreateUserResponse struct for CreateUserResponse
type CreateUserResponse struct {
	AccessToken NullableString `json:"access_token,omitempty"`
	ExpiresIn NullableInt64 `json:"expires_in,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	RefreshToken NullableString `json:"refresh_token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserResponse CreateUserResponse

// NewCreateUserResponse instantiates a new CreateUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserResponse() *CreateUserResponse {
	this := CreateUserResponse{}
	return &this
}

// NewCreateUserResponseWithDefaults instantiates a new CreateUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserResponseWithDefaults() *CreateUserResponse {
	this := CreateUserResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponse) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *CreateUserResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *CreateUserResponse) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *CreateUserResponse) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *CreateUserResponse) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponse) GetExpiresIn() int64 {
	if o == nil || IsNil(o.ExpiresIn.Get()) {
		var ret int64
		return ret
	}
	return *o.ExpiresIn.Get()
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponse) GetExpiresInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresIn.Get(), o.ExpiresIn.IsSet()
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *CreateUserResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn.IsSet() {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given NullableInt64 and assigns it to the ExpiresIn field.
func (o *CreateUserResponse) SetExpiresIn(v int64) {
	o.ExpiresIn.Set(&v)
}
// SetExpiresInNil sets the value for ExpiresIn to be an explicit nil
func (o *CreateUserResponse) SetExpiresInNil() {
	o.ExpiresIn.Set(nil)
}

// UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
func (o *CreateUserResponse) UnsetExpiresIn() {
	o.ExpiresIn.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponse) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateUserResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *CreateUserResponse) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateUserResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateUserResponse) UnsetId() {
	o.Id.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *CreateUserResponse) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *CreateUserResponse) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *CreateUserResponse) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *CreateUserResponse) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

func (o CreateUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessToken.IsSet() {
		toSerialize["access_token"] = o.AccessToken.Get()
	}
	if o.ExpiresIn.IsSet() {
		toSerialize["expires_in"] = o.ExpiresIn.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refresh_token"] = o.RefreshToken.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateUserResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateUserResponse := _CreateUserResponse{}

	err = json.Unmarshal(data, &varCreateUserResponse)

	if err != nil {
		return err
	}

	*o = CreateUserResponse(varCreateUserResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "expires_in")
		delete(additionalProperties, "id")
		delete(additionalProperties, "refresh_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateUserResponse struct {
	value *CreateUserResponse
	isSet bool
}

func (v NullableCreateUserResponse) Get() *CreateUserResponse {
	return v.value
}

func (v *NullableCreateUserResponse) Set(val *CreateUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserResponse(val *CreateUserResponse) *NullableCreateUserResponse {
	return &NullableCreateUserResponse{value: val, isSet: true}
}

func (v NullableCreateUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


