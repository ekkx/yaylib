
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SnsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnsInfo{}

// SnsInfo struct for SnsInfo
type SnsInfo struct {
	Biography NullableString `json:"biography,omitempty"`
	Gender NullableString `json:"gender,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	ProfileImage NullableString `json:"profile_image,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Uid NullableString `json:"uid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnsInfo SnsInfo

// NewSnsInfo instantiates a new SnsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnsInfo() *SnsInfo {
	this := SnsInfo{}
	return &this
}

// NewSnsInfoWithDefaults instantiates a new SnsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnsInfoWithDefaults() *SnsInfo {
	this := SnsInfo{}
	return &this
}

// GetBiography returns the Biography field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnsInfo) GetBiography() string {
	if o == nil || IsNil(o.Biography.Get()) {
		var ret string
		return ret
	}
	return *o.Biography.Get()
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsInfo) GetBiographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biography.Get(), o.Biography.IsSet()
}

// HasBiography returns a boolean if a field has been set.
func (o *SnsInfo) HasBiography() bool {
	if o != nil && o.Biography.IsSet() {
		return true
	}

	return false
}

// SetBiography gets a reference to the given NullableString and assigns it to the Biography field.
func (o *SnsInfo) SetBiography(v string) {
	o.Biography.Set(&v)
}
// SetBiographyNil sets the value for Biography to be an explicit nil
func (o *SnsInfo) SetBiographyNil() {
	o.Biography.Set(nil)
}

// UnsetBiography ensures that no value is present for Biography, not even an explicit nil
func (o *SnsInfo) UnsetBiography() {
	o.Biography.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnsInfo) GetGender() string {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret string
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsInfo) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *SnsInfo) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableString and assigns it to the Gender field.
func (o *SnsInfo) SetGender(v string) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *SnsInfo) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *SnsInfo) UnsetGender() {
	o.Gender.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnsInfo) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsInfo) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *SnsInfo) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *SnsInfo) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *SnsInfo) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *SnsInfo) UnsetNickname() {
	o.Nickname.Unset()
}

// GetProfileImage returns the ProfileImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnsInfo) GetProfileImage() string {
	if o == nil || IsNil(o.ProfileImage.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileImage.Get()
}

// GetProfileImageOk returns a tuple with the ProfileImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsInfo) GetProfileImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileImage.Get(), o.ProfileImage.IsSet()
}

// HasProfileImage returns a boolean if a field has been set.
func (o *SnsInfo) HasProfileImage() bool {
	if o != nil && o.ProfileImage.IsSet() {
		return true
	}

	return false
}

// SetProfileImage gets a reference to the given NullableString and assigns it to the ProfileImage field.
func (o *SnsInfo) SetProfileImage(v string) {
	o.ProfileImage.Set(&v)
}
// SetProfileImageNil sets the value for ProfileImage to be an explicit nil
func (o *SnsInfo) SetProfileImageNil() {
	o.ProfileImage.Set(nil)
}

// UnsetProfileImage ensures that no value is present for ProfileImage, not even an explicit nil
func (o *SnsInfo) UnsetProfileImage() {
	o.ProfileImage.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnsInfo) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *SnsInfo) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *SnsInfo) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *SnsInfo) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *SnsInfo) UnsetType() {
	o.Type.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnsInfo) GetUid() string {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret string
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnsInfo) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *SnsInfo) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableString and assigns it to the Uid field.
func (o *SnsInfo) SetUid(v string) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *SnsInfo) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *SnsInfo) UnsetUid() {
	o.Uid.Unset()
}

func (o SnsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Biography.IsSet() {
		toSerialize["biography"] = o.Biography.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if o.ProfileImage.IsSet() {
		toSerialize["profile_image"] = o.ProfileImage.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Uid.IsSet() {
		toSerialize["uid"] = o.Uid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SnsInfo) UnmarshalJSON(data []byte) (err error) {
	varSnsInfo := _SnsInfo{}

	err = json.Unmarshal(data, &varSnsInfo)

	if err != nil {
		return err
	}

	*o = SnsInfo(varSnsInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "biography")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "profile_image")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnsInfo struct {
	value *SnsInfo
	isSet bool
}

func (v NullableSnsInfo) Get() *SnsInfo {
	return v.value
}

func (v *NullableSnsInfo) Set(val *SnsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSnsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSnsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnsInfo(val *SnsInfo) *NullableSnsInfo {
	return &NullableSnsInfo{value: val, isSet: true}
}

func (v NullableSnsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


