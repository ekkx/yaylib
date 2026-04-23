
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the SearchCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCriteria{}

// SearchCriteria struct for SearchCriteria
type SearchCriteria struct {
	Biography NullableString `json:"biography,omitempty"`
	Gender NullableInt32 `json:"gender,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	Prefecture NullableString `json:"prefecture,omitempty"`
	Username NullableString `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchCriteria SearchCriteria

// NewSearchCriteria instantiates a new SearchCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCriteria() *SearchCriteria {
	this := SearchCriteria{}
	return &this
}

// NewSearchCriteriaWithDefaults instantiates a new SearchCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCriteriaWithDefaults() *SearchCriteria {
	this := SearchCriteria{}
	return &this
}

// GetBiography returns the Biography field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCriteria) GetBiography() string {
	if o == nil || IsNil(o.Biography.Get()) {
		var ret string
		return ret
	}
	return *o.Biography.Get()
}

// GetBiographyOk returns a tuple with the Biography field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCriteria) GetBiographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biography.Get(), o.Biography.IsSet()
}

// HasBiography returns a boolean if a field has been set.
func (o *SearchCriteria) HasBiography() bool {
	if o != nil && o.Biography.IsSet() {
		return true
	}

	return false
}

// SetBiography gets a reference to the given NullableString and assigns it to the Biography field.
func (o *SearchCriteria) SetBiography(v string) {
	o.Biography.Set(&v)
}
// SetBiographyNil sets the value for Biography to be an explicit nil
func (o *SearchCriteria) SetBiographyNil() {
	o.Biography.Set(nil)
}

// UnsetBiography ensures that no value is present for Biography, not even an explicit nil
func (o *SearchCriteria) UnsetBiography() {
	o.Biography.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCriteria) GetGender() int32 {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret int32
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCriteria) GetGenderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *SearchCriteria) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableInt32 and assigns it to the Gender field.
func (o *SearchCriteria) SetGender(v int32) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *SearchCriteria) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *SearchCriteria) UnsetGender() {
	o.Gender.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCriteria) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCriteria) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *SearchCriteria) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *SearchCriteria) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *SearchCriteria) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *SearchCriteria) UnsetNickname() {
	o.Nickname.Unset()
}

// GetPrefecture returns the Prefecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCriteria) GetPrefecture() string {
	if o == nil || IsNil(o.Prefecture.Get()) {
		var ret string
		return ret
	}
	return *o.Prefecture.Get()
}

// GetPrefectureOk returns a tuple with the Prefecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCriteria) GetPrefectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefecture.Get(), o.Prefecture.IsSet()
}

// HasPrefecture returns a boolean if a field has been set.
func (o *SearchCriteria) HasPrefecture() bool {
	if o != nil && o.Prefecture.IsSet() {
		return true
	}

	return false
}

// SetPrefecture gets a reference to the given NullableString and assigns it to the Prefecture field.
func (o *SearchCriteria) SetPrefecture(v string) {
	o.Prefecture.Set(&v)
}
// SetPrefectureNil sets the value for Prefecture to be an explicit nil
func (o *SearchCriteria) SetPrefectureNil() {
	o.Prefecture.Set(nil)
}

// UnsetPrefecture ensures that no value is present for Prefecture, not even an explicit nil
func (o *SearchCriteria) UnsetPrefecture() {
	o.Prefecture.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchCriteria) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchCriteria) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *SearchCriteria) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *SearchCriteria) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *SearchCriteria) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *SearchCriteria) UnsetUsername() {
	o.Username.Unset()
}

func (o SearchCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCriteria) ToMap() (map[string]interface{}, error) {
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
	if o.Prefecture.IsSet() {
		toSerialize["prefecture"] = o.Prefecture.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchCriteria) UnmarshalJSON(data []byte) (err error) {
	varSearchCriteria := _SearchCriteria{}

	err = json.Unmarshal(data, &varSearchCriteria)

	if err != nil {
		return err
	}

	*o = SearchCriteria(varSearchCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "biography")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "prefecture")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchCriteria struct {
	value *SearchCriteria
	isSet bool
}

func (v NullableSearchCriteria) Get() *SearchCriteria {
	return v.value
}

func (v *NullableSearchCriteria) Set(val *SearchCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCriteria(val *SearchCriteria) *NullableSearchCriteria {
	return &NullableSearchCriteria{value: val, isSet: true}
}

func (v NullableSearchCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


