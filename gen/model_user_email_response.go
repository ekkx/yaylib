
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserEmailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEmailResponse{}

// UserEmailResponse struct for UserEmailResponse
type UserEmailResponse struct {
	Email NullableString `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserEmailResponse UserEmailResponse

// NewUserEmailResponse instantiates a new UserEmailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEmailResponse() *UserEmailResponse {
	this := UserEmailResponse{}
	return &this
}

// NewUserEmailResponseWithDefaults instantiates a new UserEmailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEmailResponseWithDefaults() *UserEmailResponse {
	this := UserEmailResponse{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserEmailResponse) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserEmailResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserEmailResponse) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserEmailResponse) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserEmailResponse) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserEmailResponse) UnsetEmail() {
	o.Email.Unset()
}

func (o UserEmailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEmailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserEmailResponse) UnmarshalJSON(data []byte) (err error) {
	varUserEmailResponse := _UserEmailResponse{}

	err = json.Unmarshal(data, &varUserEmailResponse)

	if err != nil {
		return err
	}

	*o = UserEmailResponse(varUserEmailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserEmailResponse struct {
	value *UserEmailResponse
	isSet bool
}

func (v NullableUserEmailResponse) Get() *UserEmailResponse {
	return v.value
}

func (v *NullableUserEmailResponse) Set(val *UserEmailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEmailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEmailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEmailResponse(val *UserEmailResponse) *NullableUserEmailResponse {
	return &NullableUserEmailResponse{value: val, isSet: true}
}

func (v NullableUserEmailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEmailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


