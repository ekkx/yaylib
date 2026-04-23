
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelSnsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSnsInfo{}

// ModelSnsInfo struct for ModelSnsInfo
type ModelSnsInfo struct {
	Biography NullableString `json:"biography,omitempty"`
	Gender NullableString `json:"gender,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	ProfileImage NullableString `json:"profile_image,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Uid NullableString `json:"uid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelSnsInfo ModelSnsInfo

// NewModelSnsInfo instantiates a new ModelSnsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSnsInfo() *ModelSnsInfo {
	this := ModelSnsInfo{}
	return &this
}

// NewModelSnsInfoWithDefaults instantiates a new ModelSnsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSnsInfoWithDefaults() *ModelSnsInfo {
	this := ModelSnsInfo{}
	return &this
}

// GetBiography returns the Biography field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSnsInfo) GetBiography() string {
	if o == nil || IsNil(o.Biography.Get()) {
		var ret string
		return ret
	}
	return *o.Biography.Get()
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSnsInfo) GetBiographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biography.Get(), o.Biography.IsSet()
}

// HasBiography returns a boolean if a field has been set.
func (o *ModelSnsInfo) HasBiography() bool {
	if o != nil && o.Biography.IsSet() {
		return true
	}

	return false
}

// SetBiography gets a reference to the given NullableString and assigns it to the Biography field.
func (o *ModelSnsInfo) SetBiography(v string) {
	o.Biography.Set(&v)
}
// SetBiographyNil sets the value for Biography to be an explicit nil
func (o *ModelSnsInfo) SetBiographyNil() {
	o.Biography.Set(nil)
}

// UnsetBiography ensures that no value is present for Biography, not even an explicit nil
func (o *ModelSnsInfo) UnsetBiography() {
	o.Biography.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSnsInfo) GetGender() string {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret string
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSnsInfo) GetGenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *ModelSnsInfo) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableString and assigns it to the Gender field.
func (o *ModelSnsInfo) SetGender(v string) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *ModelSnsInfo) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *ModelSnsInfo) UnsetGender() {
	o.Gender.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSnsInfo) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSnsInfo) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *ModelSnsInfo) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *ModelSnsInfo) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *ModelSnsInfo) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *ModelSnsInfo) UnsetNickname() {
	o.Nickname.Unset()
}

// GetProfileImage returns the ProfileImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSnsInfo) GetProfileImage() string {
	if o == nil || IsNil(o.ProfileImage.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileImage.Get()
}

// GetProfileImageOk returns a tuple with the ProfileImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSnsInfo) GetProfileImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileImage.Get(), o.ProfileImage.IsSet()
}

// HasProfileImage returns a boolean if a field has been set.
func (o *ModelSnsInfo) HasProfileImage() bool {
	if o != nil && o.ProfileImage.IsSet() {
		return true
	}

	return false
}

// SetProfileImage gets a reference to the given NullableString and assigns it to the ProfileImage field.
func (o *ModelSnsInfo) SetProfileImage(v string) {
	o.ProfileImage.Set(&v)
}
// SetProfileImageNil sets the value for ProfileImage to be an explicit nil
func (o *ModelSnsInfo) SetProfileImageNil() {
	o.ProfileImage.Set(nil)
}

// UnsetProfileImage ensures that no value is present for ProfileImage, not even an explicit nil
func (o *ModelSnsInfo) UnsetProfileImage() {
	o.ProfileImage.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSnsInfo) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSnsInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ModelSnsInfo) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ModelSnsInfo) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ModelSnsInfo) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ModelSnsInfo) UnsetType() {
	o.Type.Unset()
}

// GetUid returns the Uid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelSnsInfo) GetUid() string {
	if o == nil || IsNil(o.Uid.Get()) {
		var ret string
		return ret
	}
	return *o.Uid.Get()
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSnsInfo) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uid.Get(), o.Uid.IsSet()
}

// HasUid returns a boolean if a field has been set.
func (o *ModelSnsInfo) HasUid() bool {
	if o != nil && o.Uid.IsSet() {
		return true
	}

	return false
}

// SetUid gets a reference to the given NullableString and assigns it to the Uid field.
func (o *ModelSnsInfo) SetUid(v string) {
	o.Uid.Set(&v)
}
// SetUidNil sets the value for Uid to be an explicit nil
func (o *ModelSnsInfo) SetUidNil() {
	o.Uid.Set(nil)
}

// UnsetUid ensures that no value is present for Uid, not even an explicit nil
func (o *ModelSnsInfo) UnsetUid() {
	o.Uid.Unset()
}

func (o ModelSnsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSnsInfo) ToMap() (map[string]interface{}, error) {
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

func (o *ModelSnsInfo) UnmarshalJSON(data []byte) (err error) {
	varModelSnsInfo := _ModelSnsInfo{}

	err = json.Unmarshal(data, &varModelSnsInfo)

	if err != nil {
		return err
	}

	*o = ModelSnsInfo(varModelSnsInfo)

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

type NullableModelSnsInfo struct {
	value *ModelSnsInfo
	isSet bool
}

func (v NullableModelSnsInfo) Get() *ModelSnsInfo {
	return v.value
}

func (v *NullableModelSnsInfo) Set(val *ModelSnsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSnsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSnsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSnsInfo(val *ModelSnsInfo) *NullableModelSnsInfo {
	return &NullableModelSnsInfo{value: val, isSet: true}
}

func (v NullableModelSnsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSnsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


