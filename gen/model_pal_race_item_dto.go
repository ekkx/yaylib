
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalRaceItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalRaceItemDTO{}

// PalRaceItemDTO struct for PalRaceItemDTO
type PalRaceItemDTO struct {
	CurrentLevel NullableInt32 `json:"current_level,omitempty"`
	FatigueLevel NullableString `json:"fatigue_level,omitempty"`
	Nickname NullableString `json:"nickname,omitempty"`
	OverallStrength NullableInt32 `json:"overall_strength,omitempty"`
	ProfileImage NullableProfileImageDTO `json:"profile_image,omitempty"`
	Rank NullableString `json:"rank,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	Type NullableString `json:"type,omitempty"`
	WinRate NullableString `json:"win_rate,omitempty"`
	WinningStreak NullableInt32 `json:"winning_streak,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalRaceItemDTO PalRaceItemDTO

// NewPalRaceItemDTO instantiates a new PalRaceItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalRaceItemDTO() *PalRaceItemDTO {
	this := PalRaceItemDTO{}
	return &this
}

// NewPalRaceItemDTOWithDefaults instantiates a new PalRaceItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalRaceItemDTOWithDefaults() *PalRaceItemDTO {
	this := PalRaceItemDTO{}
	return &this
}

// GetCurrentLevel returns the CurrentLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetCurrentLevel() int32 {
	if o == nil || IsNil(o.CurrentLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.CurrentLevel.Get()
}

// GetCurrentLevelOk returns a tuple with the CurrentLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetCurrentLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentLevel.Get(), o.CurrentLevel.IsSet()
}

// HasCurrentLevel returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasCurrentLevel() bool {
	if o != nil && o.CurrentLevel.IsSet() {
		return true
	}

	return false
}

// SetCurrentLevel gets a reference to the given NullableInt32 and assigns it to the CurrentLevel field.
func (o *PalRaceItemDTO) SetCurrentLevel(v int32) {
	o.CurrentLevel.Set(&v)
}
// SetCurrentLevelNil sets the value for CurrentLevel to be an explicit nil
func (o *PalRaceItemDTO) SetCurrentLevelNil() {
	o.CurrentLevel.Set(nil)
}

// UnsetCurrentLevel ensures that no value is present for CurrentLevel, not even an explicit nil
func (o *PalRaceItemDTO) UnsetCurrentLevel() {
	o.CurrentLevel.Unset()
}

// GetFatigueLevel returns the FatigueLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetFatigueLevel() string {
	if o == nil || IsNil(o.FatigueLevel.Get()) {
		var ret string
		return ret
	}
	return *o.FatigueLevel.Get()
}

// GetFatigueLevelOk returns a tuple with the FatigueLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetFatigueLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FatigueLevel.Get(), o.FatigueLevel.IsSet()
}

// HasFatigueLevel returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasFatigueLevel() bool {
	if o != nil && o.FatigueLevel.IsSet() {
		return true
	}

	return false
}

// SetFatigueLevel gets a reference to the given NullableString and assigns it to the FatigueLevel field.
func (o *PalRaceItemDTO) SetFatigueLevel(v string) {
	o.FatigueLevel.Set(&v)
}
// SetFatigueLevelNil sets the value for FatigueLevel to be an explicit nil
func (o *PalRaceItemDTO) SetFatigueLevelNil() {
	o.FatigueLevel.Set(nil)
}

// UnsetFatigueLevel ensures that no value is present for FatigueLevel, not even an explicit nil
func (o *PalRaceItemDTO) UnsetFatigueLevel() {
	o.FatigueLevel.Unset()
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *PalRaceItemDTO) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *PalRaceItemDTO) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *PalRaceItemDTO) UnsetNickname() {
	o.Nickname.Unset()
}

// GetOverallStrength returns the OverallStrength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetOverallStrength() int32 {
	if o == nil || IsNil(o.OverallStrength.Get()) {
		var ret int32
		return ret
	}
	return *o.OverallStrength.Get()
}

// GetOverallStrengthOk returns a tuple with the OverallStrength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetOverallStrengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverallStrength.Get(), o.OverallStrength.IsSet()
}

// HasOverallStrength returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasOverallStrength() bool {
	if o != nil && o.OverallStrength.IsSet() {
		return true
	}

	return false
}

// SetOverallStrength gets a reference to the given NullableInt32 and assigns it to the OverallStrength field.
func (o *PalRaceItemDTO) SetOverallStrength(v int32) {
	o.OverallStrength.Set(&v)
}
// SetOverallStrengthNil sets the value for OverallStrength to be an explicit nil
func (o *PalRaceItemDTO) SetOverallStrengthNil() {
	o.OverallStrength.Set(nil)
}

// UnsetOverallStrength ensures that no value is present for OverallStrength, not even an explicit nil
func (o *PalRaceItemDTO) UnsetOverallStrength() {
	o.OverallStrength.Unset()
}

// GetProfileImage returns the ProfileImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetProfileImage() ProfileImageDTO {
	if o == nil || IsNil(o.ProfileImage.Get()) {
		var ret ProfileImageDTO
		return ret
	}
	return *o.ProfileImage.Get()
}

// GetProfileImageOk returns a tuple with the ProfileImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetProfileImageOk() (*ProfileImageDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileImage.Get(), o.ProfileImage.IsSet()
}

// HasProfileImage returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasProfileImage() bool {
	if o != nil && o.ProfileImage.IsSet() {
		return true
	}

	return false
}

// SetProfileImage gets a reference to the given NullableProfileImageDTO and assigns it to the ProfileImage field.
func (o *PalRaceItemDTO) SetProfileImage(v ProfileImageDTO) {
	o.ProfileImage.Set(&v)
}
// SetProfileImageNil sets the value for ProfileImage to be an explicit nil
func (o *PalRaceItemDTO) SetProfileImageNil() {
	o.ProfileImage.Set(nil)
}

// UnsetProfileImage ensures that no value is present for ProfileImage, not even an explicit nil
func (o *PalRaceItemDTO) UnsetProfileImage() {
	o.ProfileImage.Unset()
}

// GetRank returns the Rank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetRank() string {
	if o == nil || IsNil(o.Rank.Get()) {
		var ret string
		return ret
	}
	return *o.Rank.Get()
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetRankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rank.Get(), o.Rank.IsSet()
}

// HasRank returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasRank() bool {
	if o != nil && o.Rank.IsSet() {
		return true
	}

	return false
}

// SetRank gets a reference to the given NullableString and assigns it to the Rank field.
func (o *PalRaceItemDTO) SetRank(v string) {
	o.Rank.Set(&v)
}
// SetRankNil sets the value for Rank to be an explicit nil
func (o *PalRaceItemDTO) SetRankNil() {
	o.Rank.Set(nil)
}

// UnsetRank ensures that no value is present for Rank, not even an explicit nil
func (o *PalRaceItemDTO) UnsetRank() {
	o.Rank.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *PalRaceItemDTO) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *PalRaceItemDTO) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *PalRaceItemDTO) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *PalRaceItemDTO) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *PalRaceItemDTO) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PalRaceItemDTO) UnsetType() {
	o.Type.Unset()
}

// GetWinRate returns the WinRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetWinRate() string {
	if o == nil || IsNil(o.WinRate.Get()) {
		var ret string
		return ret
	}
	return *o.WinRate.Get()
}

// GetWinRateOk returns a tuple with the WinRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetWinRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WinRate.Get(), o.WinRate.IsSet()
}

// HasWinRate returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasWinRate() bool {
	if o != nil && o.WinRate.IsSet() {
		return true
	}

	return false
}

// SetWinRate gets a reference to the given NullableString and assigns it to the WinRate field.
func (o *PalRaceItemDTO) SetWinRate(v string) {
	o.WinRate.Set(&v)
}
// SetWinRateNil sets the value for WinRate to be an explicit nil
func (o *PalRaceItemDTO) SetWinRateNil() {
	o.WinRate.Set(nil)
}

// UnsetWinRate ensures that no value is present for WinRate, not even an explicit nil
func (o *PalRaceItemDTO) UnsetWinRate() {
	o.WinRate.Unset()
}

// GetWinningStreak returns the WinningStreak field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalRaceItemDTO) GetWinningStreak() int32 {
	if o == nil || IsNil(o.WinningStreak.Get()) {
		var ret int32
		return ret
	}
	return *o.WinningStreak.Get()
}

// GetWinningStreakOk returns a tuple with the WinningStreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalRaceItemDTO) GetWinningStreakOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WinningStreak.Get(), o.WinningStreak.IsSet()
}

// HasWinningStreak returns a boolean if a field has been set.
func (o *PalRaceItemDTO) HasWinningStreak() bool {
	if o != nil && o.WinningStreak.IsSet() {
		return true
	}

	return false
}

// SetWinningStreak gets a reference to the given NullableInt32 and assigns it to the WinningStreak field.
func (o *PalRaceItemDTO) SetWinningStreak(v int32) {
	o.WinningStreak.Set(&v)
}
// SetWinningStreakNil sets the value for WinningStreak to be an explicit nil
func (o *PalRaceItemDTO) SetWinningStreakNil() {
	o.WinningStreak.Set(nil)
}

// UnsetWinningStreak ensures that no value is present for WinningStreak, not even an explicit nil
func (o *PalRaceItemDTO) UnsetWinningStreak() {
	o.WinningStreak.Unset()
}

func (o PalRaceItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalRaceItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentLevel.IsSet() {
		toSerialize["current_level"] = o.CurrentLevel.Get()
	}
	if o.FatigueLevel.IsSet() {
		toSerialize["fatigue_level"] = o.FatigueLevel.Get()
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if o.OverallStrength.IsSet() {
		toSerialize["overall_strength"] = o.OverallStrength.Get()
	}
	if o.ProfileImage.IsSet() {
		toSerialize["profile_image"] = o.ProfileImage.Get()
	}
	if o.Rank.IsSet() {
		toSerialize["rank"] = o.Rank.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.WinRate.IsSet() {
		toSerialize["win_rate"] = o.WinRate.Get()
	}
	if o.WinningStreak.IsSet() {
		toSerialize["winning_streak"] = o.WinningStreak.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalRaceItemDTO) UnmarshalJSON(data []byte) (err error) {
	varPalRaceItemDTO := _PalRaceItemDTO{}

	err = json.Unmarshal(data, &varPalRaceItemDTO)

	if err != nil {
		return err
	}

	*o = PalRaceItemDTO(varPalRaceItemDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current_level")
		delete(additionalProperties, "fatigue_level")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "overall_strength")
		delete(additionalProperties, "profile_image")
		delete(additionalProperties, "rank")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "win_rate")
		delete(additionalProperties, "winning_streak")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalRaceItemDTO struct {
	value *PalRaceItemDTO
	isSet bool
}

func (v NullablePalRaceItemDTO) Get() *PalRaceItemDTO {
	return v.value
}

func (v *NullablePalRaceItemDTO) Set(val *PalRaceItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePalRaceItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePalRaceItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalRaceItemDTO(val *PalRaceItemDTO) *NullablePalRaceItemDTO {
	return &NullablePalRaceItemDTO{value: val, isSet: true}
}

func (v NullablePalRaceItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalRaceItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


