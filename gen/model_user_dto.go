
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the UserDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDTO{}

// UserDTO struct for UserDTO
type UserDTO struct {
	Id NullableInt64 `json:"id,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	ProfileIcon NullableString `json:"profile_icon,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserDTO UserDTO

// NewUserDTO instantiates a new UserDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDTO() *UserDTO {
	this := UserDTO{}
	return &this
}

// NewUserDTOWithDefaults instantiates a new UserDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDTOWithDefaults() *UserDTO {
	this := UserDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *UserDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *UserDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *UserDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *UserDTO) UnsetId() {
	o.Id.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDTO) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDTO) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *UserDTO) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *UserDTO) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *UserDTO) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *UserDTO) UnsetNickname() {
	o.Nickname.Unset()
}

// GetProfileIcon returns the ProfileIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDTO) GetProfileIcon() string {
	if o == nil || IsNil(o.ProfileIcon.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileIcon.Get()
}

// GetProfileIconOk returns a tuple with the ProfileIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDTO) GetProfileIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileIcon.Get(), o.ProfileIcon.IsSet()
}

// HasProfileIcon returns a boolean if a field has been set.
func (o *UserDTO) HasProfileIcon() bool {
	if o != nil && o.ProfileIcon.IsSet() {
		return true
	}

	return false
}

// SetProfileIcon gets a reference to the given NullableString and assigns it to the ProfileIcon field.
func (o *UserDTO) SetProfileIcon(v string) {
	o.ProfileIcon.Set(&v)
}
// SetProfileIconNil sets the value for ProfileIcon to be an explicit nil
func (o *UserDTO) SetProfileIconNil() {
	o.ProfileIcon.Set(nil)
}

// UnsetProfileIcon ensures that no value is present for ProfileIcon, not even an explicit nil
func (o *UserDTO) UnsetProfileIcon() {
	o.ProfileIcon.Unset()
}

func (o UserDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if o.ProfileIcon.IsSet() {
		toSerialize["profile_icon"] = o.ProfileIcon.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserDTO) UnmarshalJSON(data []byte) (err error) {
	varUserDTO := _UserDTO{}

	err = json.Unmarshal(data, &varUserDTO)

	if err != nil {
		return err
	}

	*o = UserDTO(varUserDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "profile_icon")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserDTO struct {
	value *UserDTO
	isSet bool
}

func (v NullableUserDTO) Get() *UserDTO {
	return v.value
}

func (v *NullableUserDTO) Set(val *UserDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDTO(val *UserDTO) *NullableUserDTO {
	return &NullableUserDTO{value: val, isSet: true}
}

func (v NullableUserDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


