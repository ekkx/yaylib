
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserExternalWalletAddressDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserExternalWalletAddressDTO{}

// UserExternalWalletAddressDTO struct for UserExternalWalletAddressDTO
type UserExternalWalletAddressDTO struct {
	ExternalWalletAddress NullableString `json:"external_wallet_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserExternalWalletAddressDTO UserExternalWalletAddressDTO

// NewUserExternalWalletAddressDTO instantiates a new UserExternalWalletAddressDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserExternalWalletAddressDTO() *UserExternalWalletAddressDTO {
	this := UserExternalWalletAddressDTO{}
	return &this
}

// NewUserExternalWalletAddressDTOWithDefaults instantiates a new UserExternalWalletAddressDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserExternalWalletAddressDTOWithDefaults() *UserExternalWalletAddressDTO {
	this := UserExternalWalletAddressDTO{}
	return &this
}

// GetExternalWalletAddress returns the ExternalWalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserExternalWalletAddressDTO) GetExternalWalletAddress() string {
	if o == nil || IsNil(o.ExternalWalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalWalletAddress.Get()
}

// GetExternalWalletAddressOk returns a tuple with the ExternalWalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserExternalWalletAddressDTO) GetExternalWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalWalletAddress.Get(), o.ExternalWalletAddress.IsSet()
}

// HasExternalWalletAddress returns a boolean if a field has been set.
func (o *UserExternalWalletAddressDTO) HasExternalWalletAddress() bool {
	if o != nil && o.ExternalWalletAddress.IsSet() {
		return true
	}

	return false
}

// SetExternalWalletAddress gets a reference to the given NullableString and assigns it to the ExternalWalletAddress field.
func (o *UserExternalWalletAddressDTO) SetExternalWalletAddress(v string) {
	o.ExternalWalletAddress.Set(&v)
}
// SetExternalWalletAddressNil sets the value for ExternalWalletAddress to be an explicit nil
func (o *UserExternalWalletAddressDTO) SetExternalWalletAddressNil() {
	o.ExternalWalletAddress.Set(nil)
}

// UnsetExternalWalletAddress ensures that no value is present for ExternalWalletAddress, not even an explicit nil
func (o *UserExternalWalletAddressDTO) UnsetExternalWalletAddress() {
	o.ExternalWalletAddress.Unset()
}

func (o UserExternalWalletAddressDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserExternalWalletAddressDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalWalletAddress.IsSet() {
		toSerialize["external_wallet_address"] = o.ExternalWalletAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserExternalWalletAddressDTO) UnmarshalJSON(data []byte) (err error) {
	varUserExternalWalletAddressDTO := _UserExternalWalletAddressDTO{}

	err = json.Unmarshal(data, &varUserExternalWalletAddressDTO)

	if err != nil {
		return err
	}

	*o = UserExternalWalletAddressDTO(varUserExternalWalletAddressDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "external_wallet_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserExternalWalletAddressDTO struct {
	value *UserExternalWalletAddressDTO
	isSet bool
}

func (v NullableUserExternalWalletAddressDTO) Get() *UserExternalWalletAddressDTO {
	return v.value
}

func (v *NullableUserExternalWalletAddressDTO) Set(val *UserExternalWalletAddressDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUserExternalWalletAddressDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUserExternalWalletAddressDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserExternalWalletAddressDTO(val *UserExternalWalletAddressDTO) *NullableUserExternalWalletAddressDTO {
	return &NullableUserExternalWalletAddressDTO{value: val, isSet: true}
}

func (v NullableUserExternalWalletAddressDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserExternalWalletAddressDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


